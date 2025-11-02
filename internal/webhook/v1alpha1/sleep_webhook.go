/*
Copyright 2025.

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

package v1alpha1

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	mydomainv1alpha1 "github.com/sp-yduck/admission-sleep-webhook/api/v1alpha1"
)

// nolint:unused
// log is for logging in this package.
var sleeplog = logf.Log.WithName("sleep-resource")

// SetupSleepWebhookWithManager registers the webhook for Sleep in the manager.
func SetupSleepWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).For(&mydomainv1alpha1.Sleep{}).
		WithValidator(&SleepCustomValidator{}).
		WithDefaulter(&SleepCustomDefaulter{}).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate-my-domain-v1alpha1-sleep,mutating=true,failurePolicy=fail,sideEffects=None,groups=my.domain,resources=sleeps,verbs=create;update,versions=v1alpha1,name=msleep-v1alpha1.kb.io,admissionReviewVersions=v1

// SleepCustomDefaulter struct is responsible for setting default values on the custom resource of the
// Kind Sleep when those are created or updated.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as it is used only for temporary operations and does not need to be deeply copied.
type SleepCustomDefaulter struct {
	// TODO(user): Add more fields as needed for defaulting
}

var _ webhook.CustomDefaulter = &SleepCustomDefaulter{}

// Default implements webhook.CustomDefaulter so a webhook will be registered for the Kind Sleep.
func (d *SleepCustomDefaulter) Default(_ context.Context, obj runtime.Object) error {
	sleep, ok := obj.(*mydomainv1alpha1.Sleep)

	if !ok {
		return fmt.Errorf("expected an Sleep object but got %T", obj)
	}
	sleeplog.Info("Defaulting for Sleep", "name", sleep.GetName())

	// TODO(user): fill in your defaulting logic.

	return nil
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// NOTE: The 'path' attribute must follow a specific pattern and should not be modified directly here.
// Modifying the path for an invalid path can cause API server errors; failing to locate the webhook.
// +kubebuilder:webhook:path=/validate-my-domain-v1alpha1-sleep,mutating=false,failurePolicy=fail,sideEffects=None,groups=my.domain,resources=sleeps,verbs=create;update,versions=v1alpha1,name=vsleep-v1alpha1.kb.io,admissionReviewVersions=v1

// SleepCustomValidator struct is responsible for validating the Sleep resource
// when it is created, updated, or deleted.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as this struct is used only for temporary operations and does not need to be deeply copied.
type SleepCustomValidator struct {
	// TODO(user): Add more fields as needed for validation
}

var _ webhook.CustomValidator = &SleepCustomValidator{}

// ValidateCreate implements webhook.CustomValidator so a webhook will be registered for the type Sleep.
func (v *SleepCustomValidator) ValidateCreate(_ context.Context, obj runtime.Object) (admission.Warnings, error) {
	sleep, ok := obj.(*mydomainv1alpha1.Sleep)
	if !ok {
		return nil, fmt.Errorf("expected a Sleep object but got %T", obj)
	}
	sleeplog.Info("Validation for Sleep upon creation", "name", sleep.GetName())

	// TODO(user): fill in your validation logic upon object creation.

	return nil, nil
}

// ValidateUpdate implements webhook.CustomValidator so a webhook will be registered for the type Sleep.
func (v *SleepCustomValidator) ValidateUpdate(_ context.Context, oldObj, newObj runtime.Object) (admission.Warnings, error) {
	sleep, ok := newObj.(*mydomainv1alpha1.Sleep)
	if !ok {
		return nil, fmt.Errorf("expected a Sleep object for the newObj but got %T", newObj)
	}
	sleeplog.Info("Validation for Sleep upon update", "name", sleep.GetName())

	// TODO(user): fill in your validation logic upon object update.

	return nil, nil
}

// ValidateDelete implements webhook.CustomValidator so a webhook will be registered for the type Sleep.
func (v *SleepCustomValidator) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	sleep, ok := obj.(*mydomainv1alpha1.Sleep)
	if !ok {
		return nil, fmt.Errorf("expected a Sleep object but got %T", obj)
	}
	sleeplog.Info("Validation for Sleep upon deletion", "name", sleep.GetName())

	// TODO(user): fill in your validation logic upon object deletion.

	return nil, nil
}
