// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EnvironmentDeletePolicy Environment delete policy to protect environment from false delete
type EnvironmentDeletePolicy string

// List of EnvironmentDeletePolicy.
const (
	ENVIRONMENTDELETEPOLICY_DELETE      EnvironmentDeletePolicy = "Delete"
	ENVIRONMENTDELETEPOLICY_DONOTDELETE EnvironmentDeletePolicy = "DoNotDelete"
)

var allowedEnvironmentDeletePolicyEnumValues = []EnvironmentDeletePolicy{
	ENVIRONMENTDELETEPOLICY_DELETE,
	ENVIRONMENTDELETEPOLICY_DONOTDELETE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EnvironmentDeletePolicy) GetAllowedValues() []EnvironmentDeletePolicy {
	return allowedEnvironmentDeletePolicyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EnvironmentDeletePolicy) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EnvironmentDeletePolicy(value)
	return nil
}

// NewEnvironmentDeletePolicyFromValue returns a pointer to a valid EnvironmentDeletePolicy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEnvironmentDeletePolicyFromValue(v string) (*EnvironmentDeletePolicy, error) {
	ev := EnvironmentDeletePolicy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EnvironmentDeletePolicy: valid values are %v", v, allowedEnvironmentDeletePolicyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EnvironmentDeletePolicy) IsValid() bool {
	for _, existing := range allowedEnvironmentDeletePolicyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnvironmentDeletePolicy value.
func (v EnvironmentDeletePolicy) Ptr() *EnvironmentDeletePolicy {
	return &v
}
