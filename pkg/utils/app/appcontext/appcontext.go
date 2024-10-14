/*
Copyright 2022 The KubeVela Authors.

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

package appcontext

import (
	"strconv"

	"github.com/oam-dev/kubevela/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela/pkg/oam"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateFunctionalContext(app *v1beta1.Application) map[string]string {
	var fctx = make(map[string]string)
	fctx["autoUpdate"] = strconv.FormatBool(metav1.HasAnnotation(app.ObjectMeta, oam.AnnotationAutoUpdate))
	return fctx
}
