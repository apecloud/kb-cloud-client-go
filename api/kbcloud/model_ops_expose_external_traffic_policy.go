// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// OpsExposeExternalTrafficPolicy Specifies the external traffic policy for the exposed LoadBalancer or NodePort service.
type OpsExposeExternalTrafficPolicy string

// List of OpsExposeExternalTrafficPolicy.
const (
	OpsExposeExternalTrafficPolicyLocal   OpsExposeExternalTrafficPolicy = "Local"
	OpsExposeExternalTrafficPolicyCluster OpsExposeExternalTrafficPolicy = "Cluster"
)

var allowedOpsExposeExternalTrafficPolicyEnumValues = []OpsExposeExternalTrafficPolicy{
	OpsExposeExternalTrafficPolicyLocal,
	OpsExposeExternalTrafficPolicyCluster,
}

// GetAllowedValues returns the list of possible values.
func (v *OpsExposeExternalTrafficPolicy) GetAllowedValues() []OpsExposeExternalTrafficPolicy {
	return allowedOpsExposeExternalTrafficPolicyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OpsExposeExternalTrafficPolicy) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OpsExposeExternalTrafficPolicy(value)
	return nil
}

// NewOpsExposeExternalTrafficPolicyFromValue returns a pointer to a valid OpsExposeExternalTrafficPolicy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOpsExposeExternalTrafficPolicyFromValue(v string) (*OpsExposeExternalTrafficPolicy, error) {
	ev := OpsExposeExternalTrafficPolicy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OpsExposeExternalTrafficPolicy: valid values are %v", v, allowedOpsExposeExternalTrafficPolicyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OpsExposeExternalTrafficPolicy) IsValid() bool {
	for _, existing := range allowedOpsExposeExternalTrafficPolicyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OpsExposeExternalTrafficPolicy value.
func (v OpsExposeExternalTrafficPolicy) Ptr() *OpsExposeExternalTrafficPolicy {
	return &v
}
