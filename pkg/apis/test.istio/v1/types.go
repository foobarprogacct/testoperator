package v1

import (
        metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TestIstioOperator is a specification for a TestIstioOperator resource
type TestIstioOperator struct {
        metav1.TypeMeta   `json:",inline"`
        metav1.ObjectMeta `json:"metadata,omitempty"`

        Spec   TestIstioOperatorSpec   `json:"spec,omitempty"`
        Status TestIstioOperatorStatus `json:"status,omitempty"`
}

// +k8s:depcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type TestIstioOperatorList struct {
        metav1.TypeMeta `json:",inline"`
        metav1.ListMeta `json:"metadata,omitempty"`
        Items           []TestIstioOperator `json:"items"`
}

// TestIstioOperatorSpec defines the desired state of TestIstioOperator
type TestIstioOperatorSpec struct {
        // INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
        // Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
        // Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
        Config  string `json:"config"`
        Profile string `json:"profile"`
}

// TestIstioOperatorStatus defines the observed state of TestIstioOperator
type TestIstioOperatorStatus struct {
        // INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
        // Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
        // Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
        Status string
}
