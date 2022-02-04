package objects

import (
	alertproviderv1 "alertojon.io/operator/api/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const secretName = "pagerduty-api-token"
const secretApiTokenKeyName = "api_token"
const deploymentNamePagerduty = "pagerduty-deployment"
const envVarNamePagerdutyApiToken = "PAGERDUTY_API_TOKEN"

func CreatePagerDutySecret(alertProvider *alertproviderv1.Pagerduty) *corev1.Secret {
	secret := &corev1.Secret{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Secret",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      secretName,
			Namespace: alertProvider.Namespace,
		},
		Data: map[string][]byte{
			secretApiTokenKeyName: []byte(alertProvider.Spec.ApiToken),
		},
		Type: "Opaque",
	}
	return secret
}

func CreatePagerDutyDeployment(provider *alertproviderv1.Pagerduty) *appsv1.Deployment {
	// TODO:: inorder to support scaling and replicas use provider.Spec.Replicas and delete the following line.
	replicas := int32(1)
	deployment := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      deploymentNamePagerduty,
			Namespace: provider.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": deploymentNamePagerduty,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": deploymentNamePagerduty,
					},
				},
				Spec: corev1.PodSpec{
					ServiceAccountName: "pagerduty-operator",
					ImagePullSecrets: []corev1.LocalObjectReference{
						{
							Name: "alertojon-io", //TODO:: change to the image pull secret name to be extracted from pagerduty CR
						},
					},
					Containers: []corev1.Container{
						{
							Name:  deploymentNamePagerduty,
							Image: "registry.digitalocean.com/alertojon-io/pagetduty-operator:latest",
							Env: []corev1.EnvVar{
								{
									Name: envVarNamePagerdutyApiToken,
									ValueFrom: &corev1.EnvVarSource{
										SecretKeyRef: &corev1.SecretKeySelector{
											LocalObjectReference: corev1.LocalObjectReference{
												Name: secretName,
											},
											Key: secretApiTokenKeyName,
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
