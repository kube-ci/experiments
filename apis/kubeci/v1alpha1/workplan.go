package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ResourceKindWorkplan = "Workplan"
	ResourceWorkplans    = "workplans"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Workplan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WorkplanSpec   `json:"spec,omitempty"`
	Status WorkplanStatus `json:"status,omitempty"`
}

type Task struct { // analogous to a single pod
	SerialSteps   []Step // analogous to init-containers
	ParallelSteps []Step // analogous to sidecar-containers
}

type WorkplanSpec struct {
	Workflow     string       `json:"workflow,omitempty"`
	Tasks        []Task       `json:"tasks,omitempty"`
	TriggeredFor TriggeredFor `json:"triggeredFor"`
	// set container environment variables from configmaps and secrets
	EnvFrom []corev1.EnvFromSource `json:"envFrom,omitempty"`
}

type WorkplanStatus struct {
	Phase     string `json:"phase"`
	Reason    string `json:"reason"`
	TaskIndex int    `json:"taskIndex"`
}

type TriggeredFor struct {
	ObjectReference    ObjectReference    `json:"objectReference,omitempty"`
	ResourceGeneration ResourceGeneration `json:"resourceGeneration,omitempty"`
}

type ObjectReference struct {
	Kind       string `json:"kind,omitempty"`
	APIVersion string `json:"apiVersion,omitempty"`
	Namespace  string `json:"namespace,omitempty"`
	Name       string `json:"name,omitempty"`
}

type ResourceGeneration struct {
	Generation int64  `json:"generation,omitempty"`
	Hash       string `json:"hash,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type WorkplanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Workplan `json:"items"`
}
