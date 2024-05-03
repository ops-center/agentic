/*
Copyright 2024.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// RunbookSpec defines the desired state of Runbook
type RunbookSpec struct {
	DiagnosticChecks []DiagnosticCheck `json:"diagnosticChecks,omitempty"`

	RemedyActions []RemedyAction `json:"remedyActions"`
}

type DiagnosticCheck struct {
	Type DiagnosticCheckType   `json:"type"`
	Args *runtime.RawExtension `json:"args"`
}

type DiagnosticCheckType string

const (
	CollectLog DiagnosticCheckType = "CollectLog"
)

type RemedyAction struct {
	Type RemedyActionType      `json:"type"`
	Args *runtime.RawExtension `json:"args"`
}

type RemedyActionType string

const (
	RestartPod RemedyActionType = "RestartPod"
)

// RunbookStatus defines the observed state of Runbook
type RunbookStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:scope=Cluster

// Runbook is the Schema for the runbooks API
type Runbook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RunbookSpec   `json:"spec,omitempty"`
	Status RunbookStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// RunbookList contains a list of Runbook
type RunbookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Runbook `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Runbook{}, &RunbookList{})
}
