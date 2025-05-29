// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ClusterValidationPolicy Cluster operation validation policy, such as create, hscale, vscale, etc.
type ClusterValidationPolicy string

// List of ClusterValidationPolicy.
const (
	ClusterValidationPolicyDisabled         ClusterValidationPolicy = "Disabled"
	ClusterValidationPolicyValidateOnly     ClusterValidationPolicy = "ValidateOnly"
	ClusterValidationPolicyForbiddenExecute ClusterValidationPolicy = "ForbiddenExecute"
)

var allowedClusterValidationPolicyEnumValues = []ClusterValidationPolicy{
	ClusterValidationPolicyDisabled,
	ClusterValidationPolicyValidateOnly,
	ClusterValidationPolicyForbiddenExecute,
}

// GetAllowedValues returns the list of possible values.
func (v *ClusterValidationPolicy) GetAllowedValues() []ClusterValidationPolicy {
	return allowedClusterValidationPolicyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ClusterValidationPolicy) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ClusterValidationPolicy(value)
	return nil
}

// NewClusterValidationPolicyFromValue returns a pointer to a valid ClusterValidationPolicy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewClusterValidationPolicyFromValue(v string) (*ClusterValidationPolicy, error) {
	ev := ClusterValidationPolicy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ClusterValidationPolicy: valid values are %v", v, allowedClusterValidationPolicyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ClusterValidationPolicy) IsValid() bool {
	for _, existing := range allowedClusterValidationPolicyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ClusterValidationPolicy value.
func (v ClusterValidationPolicy) Ptr() *ClusterValidationPolicy {
	return &v
}
