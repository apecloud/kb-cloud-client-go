// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
)

// EnvironmentConditionStatus Status is the status of the condition. Can be True, False, Unknown.
type EnvironmentConditionStatus string

// List of EnvironmentConditionStatus.
const (
	EnvironmentConditionStatusTrue    EnvironmentConditionStatus = "True"
	EnvironmentConditionStatusFalse   EnvironmentConditionStatus = "False"
	EnvironmentConditionStatusUnknown EnvironmentConditionStatus = "Unknown"
)

var allowedEnvironmentConditionStatusEnumValues = []EnvironmentConditionStatus{
	EnvironmentConditionStatusTrue,
	EnvironmentConditionStatusFalse,
	EnvironmentConditionStatusUnknown,
}

// GetAllowedValues returns the list of possible values.
func (v *EnvironmentConditionStatus) GetAllowedValues() []EnvironmentConditionStatus {
	return allowedEnvironmentConditionStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EnvironmentConditionStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EnvironmentConditionStatus(value)
	return nil
}

// NewEnvironmentConditionStatusFromValue returns a pointer to a valid EnvironmentConditionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEnvironmentConditionStatusFromValue(v string) (*EnvironmentConditionStatus, error) {
	ev := EnvironmentConditionStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EnvironmentConditionStatus: valid values are %v", v, allowedEnvironmentConditionStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EnvironmentConditionStatus) IsValid() bool {
	for _, existing := range allowedEnvironmentConditionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnvironmentConditionStatus value.
func (v EnvironmentConditionStatus) Ptr() *EnvironmentConditionStatus {
	return &v
}
