package v1beta1

import (
	"fmt"

	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

func (src *ComponentDefinition) ConvertTo(dstRaw conversion.Hub) error {
	fmt.Println("ConvertTo to v1beta2---")
	return nil
}

func (src *ComponentDefinition) ConvertFrom(dstRaw conversion.Hub) error {
	fmt.Println("ConvertFrom to v1beta2---")
	return nil
}
