/*
 Copyright 2021. The KubeVela Authors.

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

package v1beta2

import (
	"reflect"

	"k8s.io/apimachinery/pkg/runtime/schema"
	k8sscheme "k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/scheme"

	"github.com/oam-dev/kubevela/apis/core.oam.dev/common"
)

// Package type metadata.
const (
	Group   = common.Group
	Version = "v1beta2"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: Group, Version: Version}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}

	// AddToScheme is a global function that registers this API group & version to a scheme
	AddToScheme = SchemeBuilder.AddToScheme
)

// ComponentDefinition type metadata.
var (
	ComponentDefinitionKind             = reflect.TypeOf(ComponentDefinition{}).Name()
	ComponentDefinitionGroupKind        = schema.GroupKind{Group: Group, Kind: ComponentDefinitionKind}.String()
	ComponentDefinitionKindAPIVersion   = ComponentDefinitionKind + "." + SchemeGroupVersion.String()
	ComponentDefinitionGroupVersionKind = SchemeGroupVersion.WithKind(ComponentDefinitionKind)
)

func init() {
	SchemeBuilder.Register(&ComponentDefinition{}, &ComponentDefinitionList{})
	_ = SchemeBuilder.AddToScheme(k8sscheme.Scheme)
}

// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}