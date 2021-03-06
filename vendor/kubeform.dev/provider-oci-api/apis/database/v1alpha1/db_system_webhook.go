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

func (r *DbSystem) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:verbs=create;update;delete,path=/validate-database-oci-kubeform-com-v1alpha1-dbsystem,mutating=false,failurePolicy=fail,groups=database.oci.kubeform.com,resources=dbsystems,versions=v1alpha1,name=dbsystem.database.oci.kubeform.io,sideEffects=None,admissionReviewVersions=v1

var _ webhook.Validator = &DbSystem{}

var dbsystemForceNewList = map[string]bool{
	"/availability_domain":                             true,
	"/backup_subnet_id":                                true,
	"/cluster_name":                                    true,
	"/data_storage_percentage":                         true,
	"/database_edition":                                true,
	"/db_home/*/database/*/backup_id":                  true,
	"/db_home/*/database/*/character_set":              true,
	"/db_home/*/database/*/database_id":                true,
	"/db_home/*/database/*/database_software_image_id": true,
	"/db_home/*/database/*/db_backup_config/*/backup_destination_details/*/id":   true,
	"/db_home/*/database/*/db_backup_config/*/backup_destination_details/*/type": true,
	"/db_home/*/database/*/db_domain":                                            true,
	"/db_home/*/database/*/db_name":                                              true,
	"/db_home/*/database/*/db_workload":                                          true,
	"/db_home/*/database/*/ncharacter_set":                                       true,
	"/db_home/*/database/*/pdb_name":                                             true,
	"/db_home/*/database/*/time_stamp_for_point_in_time_recovery":                true,
	"/db_home/*/database_software_image_id":                                      true,
	"/db_home/*/db_version":                                                      true,
	"/db_home/*/defined_tags":                                                    true,
	"/db_home/*/display_name":                                                    true,
	"/db_home/*/freeform_tags":                                                   true,
	"/db_system_options/*/storage_management":                                    true,
	"/disk_redundancy":                                                           true,
	"/display_name":                                                              true,
	"/domain":                                                                    true,
	"/fault_domains":                                                             true,
	"/hostname":                                                                  true,
	"/kms_key_id":                                                                true,
	"/kms_key_version_id":                                                        true,
	"/node_count":                                                                true,
	"/private_ip":                                                                true,
	"/source":                                                                    true,
	"/source_db_system_id":                                                       true,
	"/sparse_diskgroup":                                                          true,
	"/subnet_id":                                                                 true,
	"/time_zone":                                                                 true,
}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *DbSystem) ValidateCreate() error {
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *DbSystem) ValidateUpdate(old runtime.Object) error {
	if r.Spec.Resource.ID == "" {
		return nil
	}
	newObj := r.Spec.Resource
	res := old.(*DbSystem)
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

	for key, _ := range dbsystemForceNewList {
		keySplit := strings.Split(key, "/*")
		length := len(keySplit)
		checkIfAnyDif := false
		util.CheckIfAnyDifference("", keySplit, 0, length, &checkIfAnyDif, tempOld, tempOld, tempNew)
		util.CheckIfAnyDifference("", keySplit, 0, length, &checkIfAnyDif, tempNew, tempOld, tempNew)

		if checkIfAnyDif && r.Spec.UpdatePolicy == base.UpdatePolicyDoNotDestroy {
			return fmt.Errorf(`dbsystem "%v/%v" immutable field can't be updated. To update, change spec.updatePolicy to Destroy`, r.Namespace, r.Name)
		}
	}
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *DbSystem) ValidateDelete() error {
	if r.Spec.TerminationPolicy == base.TerminationPolicyDoNotTerminate {
		return fmt.Errorf(`dbsystem "%v/%v" can't be terminated. To delete, change spec.terminationPolicy to Delete`, r.Namespace, r.Name)
	}
	return nil
}
