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

package budget

import (
	"context"

	"github.com/go-logr/logr"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	auditlib "go.bytebuilders.dev/audit/lib"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/klog/v2"
	meta_util "kmodules.xyz/client-go/meta"
	budgetv1alpha1 "kubeform.dev/provider-oci-api/apis/budget/v1alpha1"
	"kubeform.dev/provider-oci-controller/controllers"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// BudgetReconciler reconciles a Budget object
type BudgetReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme

	Gvk              schema.GroupVersionKind // GVK of the Resource
	Provider         *tfschema.Provider      // returns a *schema.Provider from the provider package
	Resource         *tfschema.Resource      // returns *schema.Resource
	TypeName         string                  // resource type
	WatchOnlyDefault bool
}

// +kubebuilder:rbac:groups=budget.oci.kubeform.com,resources=budgets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=budget.oci.kubeform.com,resources=budgets/status,verbs=get;update;patch

func (r *BudgetReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.Log.WithValues("budget", req.NamespacedName)

	if r.WatchOnlyDefault && req.Namespace != v1.NamespaceDefault {
		log.Info("Only default namespace is supported for Kubeform Community, Please upgrade to Kubeform Enterprise to use any namespace.")
		return ctrl.Result{}, nil
	}

	var unstructuredObj unstructured.Unstructured
	unstructuredObj.SetGroupVersionKind(r.Gvk)

	if err := r.Get(ctx, req.NamespacedName, &unstructuredObj); err != nil {
		log.Error(err, "unable to fetch Budget")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	rClient := r.Client
	provider := r.Provider
	res := r.Resource
	gv := r.Gvk.GroupVersion()
	tName := r.TypeName
	jsonit := controllers.GetJSONItr(budgetv1alpha1.GetEncoder(), budgetv1alpha1.GetDecoder())
	err := controllers.StartProcess(rClient, provider, ctx, res, gv, &unstructuredObj, tName, jsonit)
	return ctrl.Result{}, err
}

func (r *BudgetReconciler) SetupWithManager(ctx context.Context, mgr ctrl.Manager, auditor *auditlib.EventPublisher) error {
	if auditor != nil {
		if err := auditor.SetupWithManager(ctx, mgr, &budgetv1alpha1.Budget{}); err != nil {
			klog.Error(err, "unable to set up auditor", budgetv1alpha1.Budget{}.APIVersion, budgetv1alpha1.Budget{}.Kind)
			return err
		}
	}

	return ctrl.NewControllerManagedBy(mgr).
		For(&budgetv1alpha1.Budget{}).
		WithEventFilter(predicate.Funcs{
			CreateFunc: func(e event.CreateEvent) bool {
				return !meta_util.MustAlreadyReconciled(e.Object)
			},
			UpdateFunc: func(e event.UpdateEvent) bool {
				return (e.ObjectNew.(metav1.Object)).GetDeletionTimestamp() != nil || !meta_util.MustAlreadyReconciled(e.ObjectNew)
			},
		}).
		Owns(&v1.Secret{}).
		Complete(r)
}
