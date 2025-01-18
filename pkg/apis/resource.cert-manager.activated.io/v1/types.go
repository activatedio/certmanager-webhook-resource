package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ACMEDNSChallenge struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ACMEDNSChallengeSpec   `json:"spec,omitempty"`
	Status            ACMEDNSChallengeStatus `json:"status,omitempty"`
}

type ACMEDNSChallengeSpec struct {
	DNSName      string `json:"dnsName"`
	Key          string `json:"key"`
	ResolvedFQDN string `json:"resolvedFQDN,omitempty"`
	ResolvedZone string `json:"resolvedZone,omitempty"`
}

type ACMEDNSChallengeStatus struct {
}
