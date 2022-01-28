package logic

//
//import (
//	appsv1 "k8s.io/api/apps/v1"
//	corev1 "k8s.io/api/core/v1"
//	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//
//	alertproviderv1 "alertojon.io/operator/api/v1"
//)
//
//// create *appsv1.Deployment object that gets an alertproviderv1 object as input and returns a Deployment object
//func (f *Factory) CreateDeployment(alertprovider *alertproviderv1.AlertProvider) *appsv1.Deployment {
//	deployment := &appsv1.Deployment{
//		ObjectMeta: metav1.ObjectMeta{
//			Name:      alertprovider.Name,
//			Namespace: alertprovider.Namespace,
//		},
//		Spec: appsv1.DeploymentSpec{
//			Replicas: &alertprovider.Spec.Replicas,
//			Selector: &metav1.LabelSelector{
//				MatchLabels: map[string]string{
//					"app": alertprovider.Name,
//				},
//			},
//			Template: corev1.PodTemplateSpec{
//				ObjectMeta: metav1.ObjectMeta{
//					Labels: map[string]string{
//						"app": alertprovider.Name,
//					},
//				},
//				Spec: corev1.PodSpec{
//					Containers: []corev1.Container{
//						{
//							Name:  alertprovider.Name,
//							Image: alertprovider.Spec.Image,
//							Ports: []corev1.ContainerPort{
//								{
//									ContainerPort: alertprovider.Spec.Port,
//								},
//							},
//						},
//					},
//				},
//			},
//		},
//	}
//	return deployment
//}
