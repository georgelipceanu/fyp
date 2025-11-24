package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// WorkloadTarget selects the workloads a policy applies to.
type WorkloadTarget struct {
	// NamespaceSelector optionally restricts to namespaces with these labels.
	// If empty, all namespaces are considered.
	// +optional
	NamespaceSelector *metav1.LabelSelector `json:"namespaceSelector,omitempty"`

	// MatchLabels selects workloads by labels.
	// This maps directly to Pod.metadata.labels for matching.
	// +optional
	MatchLabels map[string]string `json:"matchLabels,omitempty"`
}

// SchedulerHints contains per-workload preferences for the scheduler plugin.
type SchedulerHints struct {
	// UtilisationWeight is an integer 0–100.
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	UtilisationWeight *int32 `json:"utilisationWeight,omitempty"`

	// CarbonWeight is an integer 0–100.
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	CarbonWeight *int32 `json:"carbonWeight,omitempty"`

	// PreferBinPack indicates if the scheduler should favour filling nodes (bin-packing).
	// +optional
	PreferBinPack *bool `json:"preferBinPack,omitempty"`
}

// WorkloadPolicySpec describes the sustainability preferences for a set of workloads.
// In the diagram: "WorkloadPolicy • maxCarbonIntensity • schedulerHints".
type WorkloadPolicySpec struct {
	// Target defines which workloads are governed by this policy.
	Target WorkloadTarget `json:"target"`

	// MaxCarbonIntensity is an optional upper bound (gCO2/kWh) for eligible nodes.
	// +kubebuilder:validation:Minimum=0
	// +optional
	MaxCarbonIntensity *int32 `json:"maxCarbonIntensity,omitempty"`

	// SchedulerHints define how the CarbonBinPack plugin should weigh utilisation vs carbon.
	// +optional
	SchedulerHints *SchedulerHints `json:"schedulerHints,omitempty"`

	// Enforcement controls whether the policy is a soft preference or a hard requirement.
	// Allowed values: "soft", "hard".
	// +kubebuilder:validation:Enum=soft;hard
	// +kubebuilder:default=soft
	Enforcement string `json:"enforcement,omitempty"`
}

// WorkloadPolicyStatus reports how effectively the policy is applied.
type WorkloadPolicyStatus struct {
	// Enforced indicates whether hints/constraints have been applied successfully.
	// +optional
	Enforced bool `json:"enforced,omitempty"`

	// MatchedWorkloads is the number of workloads currently targeted by this policy.
	// +optional
	MatchedWorkloads int32 `json:"matchedWorkloads,omitempty"`

	// Conditions captures reconciliation state, e.g. "Ready", "PartiallyEnforced".
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// WorkloadPolicy declares carbon-aware scheduling preferences for a set of workloads.
type WorkloadPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WorkloadPolicySpec   `json:"spec,omitempty"`
	Status WorkloadPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkloadPolicyList contains a list of WorkloadPolicy.
type WorkloadPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorkloadPolicy `json:"items"`
}
