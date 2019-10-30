/*

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

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	"sigs.k8s.io/cluster-api/util"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	testv1alpha3 "github.com/liztio/proj/api/v1alpha3"
)

// TestClusterReconciler reconciles a TestCluster object
type TestClusterReconciler struct {
	client.Client
	Log logr.Logger
}

// +kubebuilder:rbac:groups=test.github.com/liztio,resources=testclusters,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=test.github.com/liztio,resources=testclusters/status,verbs=get;update;patch

func (r *TestClusterReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	ctx := context.Background()
	_ = r.Log.WithValues("mailguncluster", req.NamespacedName)

	var tCluster infrastructurev1alpha3.TestCluster
	if err := r.Get(ctx, req.NamespacedName, &mgCluster); err != nil {
		// 	apierrors "k8s.io/apimachinery/pkg/api/errors"
		if apierrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	cluster, err := util.GetOwnerMachine(ctx, r.Client, tCluster.ObjectMeta)
	if err != nil {
		return ctrl.Result{}, err
	}

	// your logic here

	return ctrl.Result{}, nil
}

func (r *TestClusterReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&testv1alpha3.TestCluster{}).
		Complete(r)
}
