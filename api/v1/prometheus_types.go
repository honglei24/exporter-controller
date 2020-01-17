/*
Copyright 2020 honglei.

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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PrometheusSpec defines the desired state of Prometheus
type PrometheusSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	//Deploy appsv1.DeploymentSpec `json:"deploy,omitempty"`

	// ConfigMap prometheus config
	ConfigMap corev1.ConfigMap `json:"config_map,omitempty"`
}

// PrometheusStatus defines the observed state of Prometheus
type PrometheusStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// Prometheus is the Schema for the prometheus API
type Prometheus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PrometheusSpec   `json:"spec,omitempty"`
	Status PrometheusStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PrometheusList contains a list of Prometheus
type PrometheusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Prometheus `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Prometheus{}, &PrometheusList{})
}
