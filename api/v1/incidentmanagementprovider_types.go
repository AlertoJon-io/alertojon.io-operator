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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// IncidentManagementProviderSpec defines the desired state of IncidentManagementProvider
type IncidentManagementProviderSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Specifies IncidentManagementProvider.
	// Valid values are:
	// - "PagerDuty" (default): allows CronJobs to run concurrently;
	// - "Opsgenie": IncidentManagement provider
	Type string `json:"type,omitempty"`

	// Specifies api key of the IncidentManagementProvider.
	// +optional
	Apikey string `json:"apikey,omitempty"`

	// Specifies secret api key for the IncidentManagementProvider.
	// +optional
	ApiKeySecretRef *ApiKeySecretKeySpec `json:"apiKeySecretRef,omitempty"`
}

type ApiKeySecretKeySpec struct {
	// Specifies the name of the secret containing the api key.
	SecretName string `json:"secretName,omitempty"`
	// Specifies the key of the secret containing the api key.
	SecretKey string `json:"secretKey,omitempty"`
}

// IncidentManagementProviderStatus defines the observed state of IncidentManagementProvider
type IncidentManagementProviderStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// IncidentManagementProvider is the Schema for the incidentmanagementproviders API
type IncidentManagementProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IncidentManagementProviderSpec   `json:"spec,omitempty"`
	Status IncidentManagementProviderStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// IncidentManagementProviderList contains a list of IncidentManagementProvider
type IncidentManagementProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IncidentManagementProvider `json:"items"`
}

func init() {
	SchemeBuilder.Register(&IncidentManagementProvider{}, &IncidentManagementProviderList{})
}
