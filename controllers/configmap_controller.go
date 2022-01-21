///*
//Copyright 2022.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//*/
//
package controllers

//
//import (
//	"context"
//	v1 "k8s.io/api/core/v1"
//	"k8s.io/apimachinery/pkg/runtime"
//	ctrl "sigs.k8s.io/controller-runtime"
//	"sigs.k8s.io/controller-runtime/pkg/client"
//	"sigs.k8s.io/controller-runtime/pkg/log"
//)
//
//// ConfigmapReconciler reconciles a Configmap object
//type ConfigmapReconciler struct {
//	client.Client
//	Scheme *runtime.Scheme
//}
//
////+kubebuilder:rbac:groups="",resources=configmaps,verbs=get;list;watch
//
//// Reconcile is part of the main kubernetes reconciliation loop which aims to
//// move the current state of the cluster closer to the desired state.
//// TODO(user): Modify the Reconcile function to compare the state specified by
//// the Pagerduty object against the actual cluster state, and then
//// perform operations to make the cluster state reflect the state specified by
//// the user.
////
//// For more details, check Reconcile and its Result here:
//// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.10.0/pkg/reconcile
//func (r *ConfigmapReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
//	logger := log.FromContext(ctx)
//	logger.Info("Reconciling Configmap")
//
//	configMap := &v1.ConfigMap{}
//	err := r.Get(ctx, req.NamespacedName, configMap)
//	if err != nil {
//		return ctrl.Result{}, client.IgnoreNotFound(err)
//	}
//	logger.Info("The data from the configmap %v\n", configMap.Data)
//	return ctrl.Result{}, nil
//}
//
//// SetupWithManager sets up the controller with the Manager.
//func (r *ConfigmapReconciler) SetupWithManager(mgr ctrl.Manager) error {
//	return ctrl.NewControllerManagedBy(mgr).
//		For(&v1.ConfigMap{}).
//		Complete(r)
//}
