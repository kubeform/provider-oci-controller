/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kubeform. DO NOT EDIT.

package ons

import (
	"context"

	"github.com/go-logr/logr"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	auditlib "go.bytebuilders.dev/audit/lib"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/klog/v2"
	meta_util "kmodules.xyz/client-go/meta"
	onsv1alpha1 "kubeform.dev/provider-oci-api/apis/ons/v1alpha1"
	"kubeform.dev/provider-oci-controller/controllers"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// NotificationTopicReconciler reconciles a NotificationTopic object
type NotificationTopicReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme

	Gvk      schema.GroupVersionKind // GVK of the Resource
	Provider *tfschema.Provider      // returns a *schema.Provider from the provider package
	Resource *tfschema.Resource      // returns *schema.Resource
	TypeName string                  // resource type
}

// +kubebuilder:rbac:groups=ons.oci.kubeform.com,resources=notificationtopics,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=ons.oci.kubeform.com,resources=notificationtopics/status,verbs=get;update;patch

func (r *NotificationTopicReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.Log.WithValues("notificationtopic", req.NamespacedName)

	var unstructuredObj unstructured.Unstructured
	unstructuredObj.SetGroupVersionKind(r.Gvk)

	if err := r.Get(ctx, req.NamespacedName, &unstructuredObj); err != nil {
		log.Error(err, "unable to fetch NotificationTopic")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	rClient := r.Client
	provider := r.Provider
	res := r.Resource
	gv := r.Gvk.GroupVersion()
	tName := r.TypeName
	jsonit := controllers.GetJSONItr(onsv1alpha1.GetEncoder(), onsv1alpha1.GetDecoder())
	err := controllers.StartProcess(rClient, provider, ctx, res, gv, &unstructuredObj, tName, jsonit)
	return ctrl.Result{}, err
}

func (r *NotificationTopicReconciler) SetupWithManager(ctx context.Context, mgr ctrl.Manager, auditor *auditlib.EventPublisher, restrictToNamespace string) error {
	if auditor != nil {
		if err := auditor.SetupWithManager(ctx, mgr, &onsv1alpha1.NotificationTopic{}); err != nil {
			klog.Error(err, "unable to set up auditor", onsv1alpha1.NotificationTopic{}.APIVersion, onsv1alpha1.NotificationTopic{}.Kind)
			return err
		}
	}

	return ctrl.NewControllerManagedBy(mgr).
		For(&onsv1alpha1.NotificationTopic{}).
		WithEventFilter(predicate.Funcs{
			CreateFunc: func(e event.CreateEvent) bool {
				return !meta_util.MustAlreadyReconciled(e.Object)
			},
			UpdateFunc: func(e event.UpdateEvent) bool {
				return (e.ObjectNew.(metav1.Object)).GetDeletionTimestamp() != nil || !meta_util.MustAlreadyReconciled(e.ObjectNew)
			},
		}).
		WithEventFilter(predicate.NewPredicateFuncs(func(e client.Object) bool {
			if restrictToNamespace != "" && e.GetNamespace() != restrictToNamespace {
				klog.Infof("Only %s namespace is supported for Kubeform Community. Please upgrade to Kubeform Enterprise to use any namespace.", restrictToNamespace)
				return false
			}
			return true
		})).
		Complete(r)
}
