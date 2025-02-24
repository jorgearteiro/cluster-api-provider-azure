/*
Copyright 2021 The Kubernetes Authors.

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

package v1beta1

import (
	"reflect"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// AzureMachineTemplateImmutableMsg ...
const AzureMachineTemplateImmutableMsg = "AzureMachineTemplate spec.template.spec field is immutable. Please create new resource instead. ref doc: https://cluster-api.sigs.k8s.io/tasks/change-machine-template.html"

// log is for logging in this package.
var machinetemplatelog = logf.Log.WithName("azuremachinetemplate-resource")

// SetupWebhookWithManager sets up and registers the webhook with the manager.
func (r *AzureMachineTemplate) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// +kubebuilder:webhook:verbs=create;update,path=/mutate-infrastructure-cluster-x-k8s-io-v1beta1-azuremachinetemplate,mutating=true,failurePolicy=fail,matchPolicy=Equivalent,groups=infrastructure.cluster.x-k8s.io,resources=azuremachinetemplates,versions=v1beta1,name=default.azuremachinetemplate.infrastructure.cluster.x-k8s.io,sideEffects=None,admissionReviewVersions=v1;v1beta1
// +kubebuilder:webhook:verbs=create;update,path=/validate-infrastructure-cluster-x-k8s-io-v1beta1-azuremachinetemplate,mutating=false,failurePolicy=fail,matchPolicy=Equivalent,groups=infrastructure.cluster.x-k8s.io,resources=azuremachinetemplates,versions=v1beta1,name=validation.azuremachinetemplate.infrastructure.cluster.x-k8s.io,sideEffects=None,admissionReviewVersions=v1;v1beta1

var _ webhook.Defaulter = &AzureMachineTemplate{}
var _ webhook.Validator = &AzureMachineTemplate{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type.
func (r *AzureMachineTemplate) ValidateCreate() error {
	machinetemplatelog.Info("validate create", "name", r.Name)
	spec := r.Spec.Template.Spec

	if allErrs := ValidateAzureMachineSpec(spec); len(allErrs) > 0 {
		return apierrors.NewInvalid(GroupVersion.WithKind("AzureMachineTemplate").GroupKind(), r.Name, allErrs)
	}
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type.
func (r *AzureMachineTemplate) ValidateUpdate(oldRaw runtime.Object) error {
	machinetemplatelog.Info("validate update", "name", r.Name)
	var allErrs field.ErrorList
	old := oldRaw.(*AzureMachineTemplate)

	if !reflect.DeepEqual(r.Spec.Template.Spec, old.Spec.Template.Spec) {
		allErrs = append(allErrs,
			field.Invalid(field.NewPath("AzureMachineTemplate", "spec", "template", "spec"), r, AzureMachineTemplateImmutableMsg),
		)
	}

	if len(allErrs) == 0 {
		return nil
	}
	return apierrors.NewInvalid(GroupVersion.WithKind("AzureMachineTemplate").GroupKind(), r.Name, allErrs)
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type.
func (r *AzureMachineTemplate) ValidateDelete() error {
	return nil
}

// Default implements webhookutil.defaulter so a webhook will be registered for the type.
func (r *AzureMachineTemplate) Default() {
	machinetemplatelog.Info("default", "name", r.Name)
	r.Spec.Template.Spec.SetDefaults(machinetemplatelog)
}
