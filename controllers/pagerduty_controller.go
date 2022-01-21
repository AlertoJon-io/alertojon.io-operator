/*
Copyright 2022.

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
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	alertproviderv1 "alertojon.io/operator/api/v1"
)

// PagerdutyReconciler reconciles a Pagerduty object
type PagerdutyReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=alertprovider.alertojon.io,resources=pagerduties,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=alertprovider.alertojon.io,resources=pagerduties/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=alertprovider.alertojon.io,resources=pagerduties/finalizers,verbs=update

//+kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core,resources=secrets,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Pagerduty object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.10.0/pkg/reconcile
func (receiver *PagerdutyReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)
	logger.WithValues("namespace", req.Namespace, "name", req.Name)

	// Fetch the Pagerduty kind
	alertProvider := &alertproviderv1.Pagerduty{}
	err := receiver.Get(ctx, req.NamespacedName, alertProvider)
	if err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	secret := receiver.createPagerDutySecretContainingApiKey(alertProvider)
	deployment := receiver.createPagerDutyDeployment(alertProvider)

	err = receiver.Client.Create(ctx, secret)
	if err != nil {
		return ctrl.Result{}, err
	}

	err = receiver.Client.Create(ctx, deployment)
	if err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (receiver *PagerdutyReconciler) createPagerDutySecretContainingApiKey(alertProvider *alertproviderv1.Pagerduty) *corev1.Secret {
	secret := &corev1.Secret{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Secret",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "pagerduty-api-token",
			Namespace: alertProvider.Namespace,
		},
		Data: map[string][]byte{
			"api_token": []byte(alertProvider.Spec.ApiToken),
		},
		Type: "Opaque",
	}
	return secret
}

func (receiver *PagerdutyReconciler) createPagerDutyDeployment(provider *alertproviderv1.Pagerduty) *appsv1.Deployment {
	// TODO:: inorder to support scaling and replicas use provider.Spec.Replicas and delete the following line.
	replicas := int32(1)
	deployment := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "pagerduty-deployment",
			Namespace: provider.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "pagerduty-deployment",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "pagerduty-deployment",
					},
				},
				Spec: corev1.PodSpec{
					ServiceAccountName: "pagerduty-operator",
					Containers: []corev1.Container{
						{
							Name:  "pagerduty-deployment",
							Image: "registry.digitalocean.com/alertojon-io/pagetduty-operator:latest",
							Env: []corev1.EnvVar{
								{
									Name: "PAGERDUTY_API_TOKEN",
									ValueFrom: &corev1.EnvVarSource{
										SecretKeyRef: &corev1.SecretKeySelector{
											LocalObjectReference: corev1.LocalObjectReference{
												Name: "pagerduty-api-token",
											},
											Key: "api_token",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	return deployment
}

// SetupWithManager sets up the controller with the Manager.
func (receiver *PagerdutyReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&alertproviderv1.Pagerduty{}).
		Complete(receiver)
}
