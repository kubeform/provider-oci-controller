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

func (r *AnomalyDetectionDataAsset) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:verbs=create;update;delete,path=/validate-ai-oci-kubeform-com-v1alpha1-anomalydetectiondataasset,mutating=false,failurePolicy=fail,groups=ai.oci.kubeform.com,resources=anomalydetectiondataassets,versions=v1alpha1,name=anomalydetectiondataasset.ai.oci.kubeform.io,sideEffects=None,admissionReviewVersions=v1

var _ webhook.Validator = &AnomalyDetectionDataAsset{}

var anomalydetectiondataassetForceNewList = map[string]bool{
	"/data_source_details/*/atp_user_name":                                    true,
	"/data_source_details/*/bucket":                                           true,
	"/data_source_details/*/cwallet_file_secret_id":                           true,
	"/data_source_details/*/data_source_type":                                 true,
	"/data_source_details/*/database_name":                                    true,
	"/data_source_details/*/ewallet_file_secret_id":                           true,
	"/data_source_details/*/key_store_file_secret_id":                         true,
	"/data_source_details/*/measurement_name":                                 true,
	"/data_source_details/*/namespace":                                        true,
	"/data_source_details/*/object":                                           true,
	"/data_source_details/*/ojdbc_file_secret_id":                             true,
	"/data_source_details/*/table_name":                                       true,
	"/data_source_details/*/tnsnames_file_secret_id":                          true,
	"/data_source_details/*/truststore_file_secret_id":                        true,
	"/data_source_details/*/url":                                              true,
	"/data_source_details/*/user_name":                                        true,
	"/data_source_details/*/version_specific_details/*/bucket":                true,
	"/data_source_details/*/version_specific_details/*/database_name":         true,
	"/data_source_details/*/version_specific_details/*/influx_version":        true,
	"/data_source_details/*/version_specific_details/*/organization_name":     true,
	"/data_source_details/*/version_specific_details/*/retention_policy_name": true,
	"/private_endpoint_id":                                                    true,
	"/project_id":                                                             true,
}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *AnomalyDetectionDataAsset) ValidateCreate() error {
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *AnomalyDetectionDataAsset) ValidateUpdate(old runtime.Object) error {
	if r.Spec.Resource.ID == "" {
		return nil
	}
	newObj := r.Spec.Resource
	res := old.(*AnomalyDetectionDataAsset)
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

	for key := range anomalydetectiondataassetForceNewList {
		keySplit := strings.Split(key, "/*")
		length := len(keySplit)
		checkIfAnyDif := false
		util.CheckIfAnyDifference("", keySplit, 0, length, &checkIfAnyDif, tempOld, tempOld, tempNew)
		util.CheckIfAnyDifference("", keySplit, 0, length, &checkIfAnyDif, tempNew, tempOld, tempNew)

		if checkIfAnyDif && r.Spec.UpdatePolicy == base.UpdatePolicyDoNotDestroy {
			return fmt.Errorf(`anomalydetectiondataasset "%v/%v" immutable field can't be updated. To update, change spec.updatePolicy to Destroy`, r.Namespace, r.Name)
		}
	}
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *AnomalyDetectionDataAsset) ValidateDelete() error {
	if r.Spec.TerminationPolicy == base.TerminationPolicyDoNotTerminate {
		return fmt.Errorf(`anomalydetectiondataasset "%v/%v" can't be terminated. To delete, change spec.terminationPolicy to Delete`, r.Namespace, r.Name)
	}
	return nil
}
