// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// SchedulingPolicyType * `HardAntiAffinity` - Strictly enforce pod anti-affinity across nodes. Pods will not be scheduled when the anti-affinity constraints cannot be satisfied.
// * `SoftAntiAffinity` - Prefer to spread pods across nodes using anti-affinity, but allow scheduling on the same node when constraints cannot be fully satisfied.
// * `Disabled` - Do not apply pod anti-affinity constraints on nodes.
type SchedulingPolicyType string

// List of SchedulingPolicyType.
const (
	SchedulingPolicyTypeHardAntiAffinity SchedulingPolicyType = "HardAntiAffinity"
	SchedulingPolicyTypeSoftAntiAffinity SchedulingPolicyType = "SoftAntiAffinity"
	SchedulingPolicyTypeDisabled         SchedulingPolicyType = "Disabled"
)

var allowedSchedulingPolicyTypeEnumValues = []SchedulingPolicyType{
	SchedulingPolicyTypeHardAntiAffinity,
	SchedulingPolicyTypeSoftAntiAffinity,
	SchedulingPolicyTypeDisabled,
}

// GetAllowedValues returns the list of possible values.
func (v *SchedulingPolicyType) GetAllowedValues() []SchedulingPolicyType {
	return allowedSchedulingPolicyTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SchedulingPolicyType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SchedulingPolicyType(value)
	return nil
}

// NewSchedulingPolicyTypeFromValue returns a pointer to a valid SchedulingPolicyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSchedulingPolicyTypeFromValue(v string) (*SchedulingPolicyType, error) {
	ev := SchedulingPolicyType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SchedulingPolicyType: valid values are %v", v, allowedSchedulingPolicyTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SchedulingPolicyType) IsValid() bool {
	for _, existing := range allowedSchedulingPolicyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SchedulingPolicyType value.
func (v SchedulingPolicyType) Ptr() *SchedulingPolicyType {
	return &v
}
