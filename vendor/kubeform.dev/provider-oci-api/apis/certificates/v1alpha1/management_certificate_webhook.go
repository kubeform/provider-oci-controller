/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	"fmt"
	"strings"

	base "kubeform.dev/apimachinery/api/v1alpha1"
	"kubeform.dev/apimachinery/pkg/util"

	jsoniter "github.com/json-iterator/go"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

func (r *ManagementCertificate) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:verbs=create;update;delete,path=/validate-certificates-oci-kubeform-com-v1alpha1-managementcertificate,mutating=false,failurePolicy=fail,groups=certificates.oci.kubeform.com,resources=managementcertificates,versions=v1alpha1,name=managementcertificate.certificates.oci.kubeform.io,sideEffects=None,admissionReviewVersions=v1

var _ webhook.Validator = &ManagementCertificate{}

var managementcertificateForceNewList = map[string]bool{
	"/certificate_config/*/certificate_profile_type":               true,
	"/certificate_config/*/issuer_certificate_authority_id":        true,
	"/certificate_config/*/key_algorithm":                          true,
	"/certificate_config/*/signature_algorithm":                    true,
	"/certificate_config/*/subject/*/common_name":                  true,
	"/certificate_config/*/subject/*/country":                      true,
	"/certificate_config/*/subject/*/distinguished_name_qualifier": true,
	"/certificate_config/*/subject/*/domain_component":             true,
	"/certificate_config/*/subject/*/generation_qualifier":         true,
	"/certificate_config/*/subject/*/given_name":                   true,
	"/certificate_config/*/subject/*/initials":                     true,
	"/certificate_config/*/subject/*/locality_name":                true,
	"/certificate_config/*/subject/*/organization":                 true,
	"/certificate_config/*/subject/*/organizational_unit":          true,
	"/certificate_config/*/subject/*/pseudonym":                    true,
	"/certificate_config/*/subject/*/serial_number":                true,
	"/certificate_config/*/subject/*/state_or_province_name":       true,
	"/certificate_config/*/subject/*/street":                       true,
	"/certificate_config/*/subject/*/surname":                      true,
	"/certificate_config/*/subject/*/title":                        true,
	"/certificate_config/*/subject/*/user_id":                      true,
	"/certificate_config/*/subject_alternative_names/*/type":       true,
	"/certificate_config/*/subject_alternative_names/*/value":      true,
	"/name": true,
}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *ManagementCertificate) ValidateCreate() error {
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *ManagementCertificate) ValidateUpdate(old runtime.Object) error {
	if r.Spec.Resource.ID == "" {
		return nil
	}
	newObj := r.Spec.Resource
	res := old.(*ManagementCertificate)
	oldObj := res.Spec.Resource

	jsnitr := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		TagKey:                 "tf",
		ValidateJsonRawMessage: true,
		TypeEncoders:           GetEncoder(),
		TypeDecoders:           GetDecoder(),
	}.Froze()

	byteNew, err := jsnitr.Marshal(newObj)
	if err != nil {
		return err
	}
	tempNew := make(map[string]interface{})
	err = jsnitr.Unmarshal(byteNew, &tempNew)
	if err != nil {
		return err
	}

	byteOld, err := jsnitr.Marshal(oldObj)
	if err != nil {
		return err
	}
	tempOld := make(map[string]interface{})
	err = jsnitr.Unmarshal(byteOld, &tempOld)
	if err != nil {
		return err
	}

	for key, _ := range managementcertificateForceNewList {
		keySplit := strings.Split(key, "/*")
		length := len(keySplit)
		checkIfAnyDif := false
		util.CheckIfAnyDifference("", keySplit, 0, length, &checkIfAnyDif, tempOld, tempOld, tempNew)
		util.CheckIfAnyDifference("", keySplit, 0, length, &checkIfAnyDif, tempNew, tempOld, tempNew)

		if checkIfAnyDif && r.Spec.UpdatePolicy == base.UpdatePolicyDoNotDestroy {
			return fmt.Errorf(`managementcertificate "%v/%v" immutable field can't be updated. To update, change spec.updatePolicy to Destroy`, r.Namespace, r.Name)
		}
	}
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *ManagementCertificate) ValidateDelete() error {
	if r.Spec.TerminationPolicy == base.TerminationPolicyDoNotTerminate {
		return fmt.Errorf(`managementcertificate "%v/%v" can't be terminated. To delete, change spec.terminationPolicy to Delete`, r.Namespace, r.Name)
	}
	return nil
}
