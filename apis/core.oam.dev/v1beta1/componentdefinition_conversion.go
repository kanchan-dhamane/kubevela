package v1beta1

import (
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/oam-dev/kubevela/apis/core.oam.dev/v1beta2"
)

var ComponentDefinitionlog = logf.Log.WithName("ComponentDefinition-resource")

// convert v1lbeta1 to v1beta2
func (src *ComponentDefinition) ConvertTo(dstRaw conversion.Hub) error {
	ComponentDefinitionlog.Info("ConvertTo to v1beta2--- ----------------")
	dst := dstRaw.(*v1beta2.ComponentDefinition)
	version := "1.6.7"
	dst.ObjectMeta = src.ObjectMeta
	componentDefinitionSpecTest := &v1beta2.ComponentDefinitionSpecTest{}
	componentDefinitionSpecTest.Version = version
	componentDefinitionSpecTest.Workload = src.Spec.Workload
	componentDefinitionSpecTest.ChildResourceKinds = src.Spec.ChildResourceKinds
	componentDefinitionSpecTest.RevisionLabel = src.Spec.RevisionLabel
	componentDefinitionSpecTest.PodSpecPath = src.Spec.PodSpecPath
	componentDefinitionSpecTest.Status = src.Spec.Status
	componentDefinitionSpecTest.Schematic = src.Spec.Schematic
	componentDefinitionSpecTest.Extension = src.Spec.Extension
	dst.Spec.Versions = []v1beta2.ComponentDefinitionSpecTest{*componentDefinitionSpecTest}

	return nil
}

// convert v1beta2 to v1beta1
func (dst *ComponentDefinition) ConvertFrom(srcRaw conversion.Hub) error {
	ComponentDefinitionlog.Info("ConvertFrom to v1beta2--- ------------------")
	src := srcRaw.(*v1beta2.ComponentDefinition)

	if len(src.Spec.Versions) >= 1 {
		ComponentDefinitionlog.Info(src.Spec.Versions[0].Version)
	}

	componentDefinitionSpec := &ComponentDefinitionSpec{}
	componentDefinitionSpec.Workload = src.Spec.Versions[0].Workload
	componentDefinitionSpec.ChildResourceKinds = src.Spec.Versions[0].ChildResourceKinds
	componentDefinitionSpec.RevisionLabel = src.Spec.Versions[0].RevisionLabel
	componentDefinitionSpec.Status = src.Spec.Versions[0].Status
	componentDefinitionSpec.Schematic = src.Spec.Versions[0].Schematic
	componentDefinitionSpec.Extension = src.Spec.Versions[0].Extension
	dst.Spec = *componentDefinitionSpec
	dst.ObjectMeta = src.ObjectMeta
	return nil
}
