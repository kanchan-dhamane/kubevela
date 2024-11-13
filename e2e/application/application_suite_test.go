/*
Copyright 2021 The KubeVela Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package e2e

import (
	"testing"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	k8s "github.com/gruntwork-io/terratest/modules/k8s"

	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/sirupsen/logrus"
)

func TestApplication(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Application Suite")
}

var _ = ginkgo.BeforeSuite(func() {
	logrus.Debug("Deploy the componentdefinition")
	deployComponentDefition()

})

var _ = ginkgo.AfterSuite(func() {
	deleteComponentDefintion()
	logrus.Debug("delete the componentdefinition")
})

func deployComponentDefition() {
	logrus.Debug("Deploying jspolicies")
	k8sOptions := &k8s.KubectlOptions{Logger: logger.Discard}
	_err := k8s.KubectlApplyE(ginkgo.GinkgoT(), k8sOptions, "./componentdefinitions")
	gomega.Expect(_err).NotTo(gomega.HaveOccurred())
	logrus.Debug("Deployed jspolicies")
}

func deleteComponentDefintion() {
	logrus.Debug("Deleting jspolicies")
	k8sOptions := &k8s.KubectlOptions{Logger: logger.Discard}
	_err := k8s.KubectlDeleteE(ginkgo.GinkgoT(), k8sOptions, "./componentdefinitions")
	gomega.Expect(_err).NotTo(gomega.HaveOccurred())
	logrus.Debug("Deleted jspolicies")
}
