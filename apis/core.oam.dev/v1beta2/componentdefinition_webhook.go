package v1beta2

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

var ComponentDefinitionlog = logf.Log.WithName("ComponentDefinition-resource")

func (r *ComponentDefinition) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

var _ webhook.Defaulter = &ComponentDefinition{}

var _ webhook.Validator = &ComponentDefinition{}

func (r *ComponentDefinition) Default() {

}

func (r *ComponentDefinition) ValidateCreate() error {
	ComponentDefinitionlog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *ComponentDefinition) ValidateUpdate(old runtime.Object) error {
	ComponentDefinitionlog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *ComponentDefinition) ValidateDelete() error {
	ComponentDefinitionlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
