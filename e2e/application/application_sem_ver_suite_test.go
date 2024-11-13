package e2e


import (
	context2 "context"
	"fmt"
	"strings"
	"time"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/oam-dev/kubevela/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela/e2e"
	"github.com/oam-dev/kubevela/pkg/utils/common"
)

var UsecorrectVersionOfComponentDefinition = func(context string, applicationName, workloadType, envName string) bool {
	return ginkgo.It(context+": should get status of the service", func() {
		ginkgo.By("init new k8s client")
		k8sclient, err := common.NewK8sClient()
		gomega.Expect(err).NotTo(gomega.HaveOccurred())

		ginkgo.By("check Application reconciled ready")
		app := &v1beta1.Application{}
		gomega.Eventually(func() bool {
			_ = k8sclient.Get(context2.Background(), client.ObjectKey{Name: applicationName, Namespace: "default"}, app)
			return app.Status.LatestRevision != nil
		}, 180*time.Second, 1*time.Second).Should(gomega.BeTrue())

		cli := fmt.Sprintf("vela status %s", applicationName)
		output, err := e2e.LongTimeExec(cli, 120*time.Second)
		gomega.Expect(err).NotTo(gomega.HaveOccurred())
		gomega.Expect(strings.ToLower(output)).To(gomega.ContainSubstring("healthy"))
		// TODO(zzxwill) need to check workloadType after app status is refined
	})
}