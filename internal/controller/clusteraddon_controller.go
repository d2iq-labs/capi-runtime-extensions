// Copyright 2023 D2iQ, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	clusteraddonsv1alpha1 "github.com/d2iq-labs/capi-runtime-extensions/api/v1alpha1"
)

// ClusterAddonReconciler reconciles a ClusterAddon object.
type ClusterAddonReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//nolint:lll // This is a long line.
//+kubebuilder:rbac:groups=clusteraddons.labs.d2iq.io,resources=clusteraddons,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=clusteraddons.labs.d2iq.io,resources=clusteraddons/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=clusteraddons.labs.d2iq.io,resources=clusteraddons/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the ClusterAddon object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *ClusterAddonReconciler) Reconcile(
	ctx context.Context,
	req ctrl.Request,
) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ClusterAddonReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&clusteraddonsv1alpha1.ClusterAddon{}).
		Complete(r)
}
