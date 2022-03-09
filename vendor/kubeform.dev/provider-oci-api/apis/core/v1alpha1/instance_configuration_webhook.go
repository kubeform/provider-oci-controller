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

func (r *InstanceConfiguration) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:verbs=create;update;delete,path=/validate-core-oci-kubeform-com-v1alpha1-instanceconfiguration,mutating=false,failurePolicy=fail,groups=core.oci.kubeform.com,resources=instanceconfigurations,versions=v1alpha1,name=instanceconfiguration.core.oci.kubeform.io,sideEffects=None,admissionReviewVersions=v1

var _ webhook.Validator = &InstanceConfiguration{}

var instanceconfigurationForceNewList = map[string]bool{
	"/instance_details/*/block_volumes/*/attach_details/*/device":                                                 true,
	"/instance_details/*/block_volumes/*/attach_details/*/display_name":                                           true,
	"/instance_details/*/block_volumes/*/attach_details/*/is_pv_encryption_in_transit_enabled":                    true,
	"/instance_details/*/block_volumes/*/attach_details/*/is_read_only":                                           true,
	"/instance_details/*/block_volumes/*/attach_details/*/is_shareable":                                           true,
	"/instance_details/*/block_volumes/*/attach_details/*/type":                                                   true,
	"/instance_details/*/block_volumes/*/attach_details/*/use_chap":                                               true,
	"/instance_details/*/block_volumes/*/create_details/*/availability_domain":                                    true,
	"/instance_details/*/block_volumes/*/create_details/*/backup_policy_id":                                       true,
	"/instance_details/*/block_volumes/*/create_details/*/compartment_id":                                         true,
	"/instance_details/*/block_volumes/*/create_details/*/defined_tags":                                           true,
	"/instance_details/*/block_volumes/*/create_details/*/display_name":                                           true,
	"/instance_details/*/block_volumes/*/create_details/*/freeform_tags":                                          true,
	"/instance_details/*/block_volumes/*/create_details/*/kms_key_id":                                             true,
	"/instance_details/*/block_volumes/*/create_details/*/size_in_gbs":                                            true,
	"/instance_details/*/block_volumes/*/create_details/*/source_details/*/id":                                    true,
	"/instance_details/*/block_volumes/*/create_details/*/source_details/*/type":                                  true,
	"/instance_details/*/block_volumes/*/create_details/*/vpus_per_gb":                                            true,
	"/instance_details/*/block_volumes/*/volume_id":                                                               true,
	"/instance_details/*/instance_type":                                                                           true,
	"/instance_details/*/launch_details/*/agent_config/*/are_all_plugins_disabled":                                true,
	"/instance_details/*/launch_details/*/agent_config/*/is_management_disabled":                                  true,
	"/instance_details/*/launch_details/*/agent_config/*/is_monitoring_disabled":                                  true,
	"/instance_details/*/launch_details/*/agent_config/*/plugins_config/*/desired_state":                          true,
	"/instance_details/*/launch_details/*/agent_config/*/plugins_config/*/name":                                   true,
	"/instance_details/*/launch_details/*/availability_config/*/recovery_action":                                  true,
	"/instance_details/*/launch_details/*/availability_domain":                                                    true,
	"/instance_details/*/launch_details/*/capacity_reservation_id":                                                true,
	"/instance_details/*/launch_details/*/compartment_id":                                                         true,
	"/instance_details/*/launch_details/*/create_vnic_details/*/assign_private_dns_record":                        true,
	"/instance_details/*/launch_details/*/create_vnic_details/*/assign_public_ip":                                 true,
	"/instance_details/*/launch_details/*/create_vnic_details/*/defined_tags":                                     true,
	"/instance_details/*/launch_details/*/create_vnic_details/*/display_name":                                     true,
	"/instance_details/*/launch_details/*/create_vnic_details/*/freeform_tags":                                    true,
	"/instance_details/*/launch_details/*/create_vnic_details/*/hostname_label":                                   true,
	"/instance_details/*/launch_details/*/create_vnic_details/*/nsg_ids":                                          true,
	"/instance_details/*/launch_details/*/create_vnic_details/*/private_ip":                                       true,
	"/instance_details/*/launch_details/*/create_vnic_details/*/skip_source_dest_check":                           true,
	"/instance_details/*/launch_details/*/create_vnic_details/*/subnet_id":                                        true,
	"/instance_details/*/launch_details/*/dedicated_vm_host_id":                                                   true,
	"/instance_details/*/launch_details/*/defined_tags":                                                           true,
	"/instance_details/*/launch_details/*/display_name":                                                           true,
	"/instance_details/*/launch_details/*/extended_metadata":                                                      true,
	"/instance_details/*/launch_details/*/fault_domain":                                                           true,
	"/instance_details/*/launch_details/*/freeform_tags":                                                          true,
	"/instance_details/*/launch_details/*/instance_options/*/are_legacy_imds_endpoints_disabled":                  true,
	"/instance_details/*/launch_details/*/ipxe_script":                                                            true,
	"/instance_details/*/launch_details/*/is_pv_encryption_in_transit_enabled":                                    true,
	"/instance_details/*/launch_details/*/launch_mode":                                                            true,
	"/instance_details/*/launch_details/*/launch_options/*/boot_volume_type":                                      true,
	"/instance_details/*/launch_details/*/launch_options/*/firmware":                                              true,
	"/instance_details/*/launch_details/*/launch_options/*/is_consistent_volume_naming_enabled":                   true,
	"/instance_details/*/launch_details/*/launch_options/*/is_pv_encryption_in_transit_enabled":                   true,
	"/instance_details/*/launch_details/*/launch_options/*/network_type":                                          true,
	"/instance_details/*/launch_details/*/launch_options/*/remote_data_volume_type":                               true,
	"/instance_details/*/launch_details/*/metadata":                                                               true,
	"/instance_details/*/launch_details/*/platform_config/*/is_measured_boot_enabled":                             true,
	"/instance_details/*/launch_details/*/platform_config/*/is_secure_boot_enabled":                               true,
	"/instance_details/*/launch_details/*/platform_config/*/is_trusted_platform_module_enabled":                   true,
	"/instance_details/*/launch_details/*/platform_config/*/numa_nodes_per_socket":                                true,
	"/instance_details/*/launch_details/*/platform_config/*/type":                                                 true,
	"/instance_details/*/launch_details/*/preemptible_instance_config/*/preemption_action/*/preserve_boot_volume": true,
	"/instance_details/*/launch_details/*/preemptible_instance_config/*/preemption_action/*/type":                 true,
	"/instance_details/*/launch_details/*/preferred_maintenance_action":                                           true,
	"/instance_details/*/launch_details/*/shape":                                                                  true,
	"/instance_details/*/launch_details/*/shape_config/*/baseline_ocpu_utilization":                               true,
	"/instance_details/*/launch_details/*/shape_config/*/memory_in_gbs":                                           true,
	"/instance_details/*/launch_details/*/shape_config/*/ocpus":                                                   true,
	"/instance_details/*/launch_details/*/source_details/*/boot_volume_id":                                        true,
	"/instance_details/*/launch_details/*/source_details/*/boot_volume_size_in_gbs":                               true,
	"/instance_details/*/launch_details/*/source_details/*/image_id":                                              true,
	"/instance_details/*/launch_details/*/source_details/*/source_type":                                           true,
	"/instance_details/*/secondary_vnics/*/create_vnic_details/*/assign_private_dns_record":                       true,
	"/instance_details/*/secondary_vnics/*/create_vnic_details/*/assign_public_ip":                                true,
	"/instance_details/*/secondary_vnics/*/create_vnic_details/*/defined_tags":                                    true,
	"/instance_details/*/secondary_vnics/*/create_vnic_details/*/display_name":                                    true,
	"/instance_details/*/secondary_vnics/*/create_vnic_details/*/freeform_tags":                                   true,
	"/instance_details/*/secondary_vnics/*/create_vnic_details/*/hostname_label":                                  true,
	"/instance_details/*/secondary_vnics/*/create_vnic_details/*/nsg_ids":                                         true,
	"/instance_details/*/secondary_vnics/*/create_vnic_details/*/private_ip":                                      true,
	"/instance_details/*/secondary_vnics/*/create_vnic_details/*/skip_source_dest_check":                          true,
	"/instance_details/*/secondary_vnics/*/create_vnic_details/*/subnet_id":                                       true,
	"/instance_details/*/secondary_vnics/*/display_name":                                                          true,
	"/instance_details/*/secondary_vnics/*/nic_index":                                                             true,
	"/instance_id": true,
	"/source":      true,
}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *InstanceConfiguration) ValidateCreate() error {
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *InstanceConfiguration) ValidateUpdate(old runtime.Object) error {
	if r.Spec.Resource.ID == "" {
		return nil
	}
	newObj := r.Spec.Resource
	res := old.(*InstanceConfiguration)
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

	for key, _ := range instanceconfigurationForceNewList {
		keySplit := strings.Split(key, "/*")
		length := len(keySplit)
		checkIfAnyDif := false
		util.CheckIfAnyDifference("", keySplit, 0, length, &checkIfAnyDif, tempOld, tempOld, tempNew)
		util.CheckIfAnyDifference("", keySplit, 0, length, &checkIfAnyDif, tempNew, tempOld, tempNew)

		if checkIfAnyDif && r.Spec.UpdatePolicy == base.UpdatePolicyDoNotDestroy {
			return fmt.Errorf(`instanceconfiguration "%v/%v" immutable field can't be updated. To update, change spec.updatePolicy to Destroy`, r.Namespace, r.Name)
		}
	}
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *InstanceConfiguration) ValidateDelete() error {
	if r.Spec.TerminationPolicy == base.TerminationPolicyDoNotTerminate {
		return fmt.Errorf(`instanceconfiguration "%v/%v" can't be terminated. To delete, change spec.terminationPolicy to Delete`, r.Namespace, r.Name)
	}
	return nil
}
