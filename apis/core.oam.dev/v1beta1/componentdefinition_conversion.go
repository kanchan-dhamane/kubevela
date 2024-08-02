package v1beta1

import (
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	"encoding/json"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/oam-dev/kubevela/apis/core.oam.dev/v1beta2"
)

var ComponentDefinitionlog = logf.Log.WithName("ComponentDefinition-resource")

// ConvertTo translates v1lbeta1 to v1beta2
func (src *ComponentDefinition) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1beta2.ComponentDefinition)

	ComponentDefinitionlog.Info("src: %v dst: %v", src.APIVersion, dst.APIVersion)
	// if src.APIVersion == dst.APIVersion {
	// 	return nil
	// }
	restored := &v1beta2.ComponentDefinition{}
	ok, err := UnmarshalData(src, restored)
	if err != nil {
		return err
	}
	if ok {
		dst.ObjectMeta = src.ObjectMeta
		dst.Spec.Versions = restored.Spec.Versions
		return nil
	}

	ComponentDefinitionlog.Info("Convert to v1beta2--- ----------------")

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
	// ComponentDefinitionlog.Info("Source:")
	// ComponentDefinitionlog.Info(src)

	dst.ObjectMeta = src.ObjectMeta
	if err := MarshalData(src, dst); err != nil {
		return nil
	}
	componentDefinitionSpec := &ComponentDefinitionSpec{}
	componentDefinitionSpec.Workload = src.Spec.Versions[0].Workload
	componentDefinitionSpec.ChildResourceKinds = src.Spec.Versions[0].ChildResourceKinds
	componentDefinitionSpec.RevisionLabel = src.Spec.Versions[0].RevisionLabel
	componentDefinitionSpec.Status = src.Spec.Versions[0].Status
	componentDefinitionSpec.Schematic = src.Spec.Versions[0].Schematic
	componentDefinitionSpec.Extension = src.Spec.Versions[0].Extension
	dst.Spec = *componentDefinitionSpec

	return nil
}

const DataAnnotation = "kd-is-an-idiot"

// MarshalData stores the source object as json data in the destination object annotations map.
// It ignores the metadata of the source object.
func MarshalData(src metav1.Object, dst metav1.Object) error {
	u, err := runtime.DefaultUnstructuredConverter.ToUnstructured(src)
	if err != nil {
		// ComponentDefinitionlog.Info("Marshalling Error: %v", err)
		return err
	}
	delete(u, "metadata")

	data, err := json.Marshal(u)
	if err != nil {
		// ComponentDefinitionlog.Info("Marshalling Error - 2: %v", err)
		return err
	}

	ComponentDefinitionlog.Info("Got past marshalling")
	annotations := dst.GetAnnotations()
	if annotations == nil {
		annotations = map[string]string{}
	}
	ComponentDefinitionlog.Info("Got past marshalling - 2")
	annotations[DataAnnotation] = string(data)
	dst.SetAnnotations(annotations)
	// ComponentDefinitionlog.Info("annotation %v", dst.GetAnnotations())
	return nil
}

// UnmarshalData tries to retrieve the data from the annotation and unmarshals it into the object passed as input.
func UnmarshalData(from metav1.Object, to interface{}) (bool, error) {
	annotations := from.GetAnnotations()
	data, ok := annotations[DataAnnotation]
	ComponentDefinitionlog.Info("Got annotation %s", data)
	if !ok {
		ComponentDefinitionlog.Info("Did not find the annotation %s", DataAnnotation)
		return false, nil
	}
	if err := json.Unmarshal([]byte(data), to); err != nil {
		return false, err
	}
	delete(annotations, DataAnnotation)
	from.SetAnnotations(annotations)
	return true, nil
}
