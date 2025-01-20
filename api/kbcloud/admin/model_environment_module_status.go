// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EnvironmentModuleStatus Status of environment module
type EnvironmentModuleStatus string

// List of EnvironmentModuleStatus.
const (
	EnvironmentModuleStatusRunning  EnvironmentModuleStatus = "Running"
	EnvironmentModuleStatusUpdating EnvironmentModuleStatus = "Updating"
	EnvironmentModuleStatusError    EnvironmentModuleStatus = "Error"
	EnvironmentModuleStatusDisabled EnvironmentModuleStatus = "Disabled"
)

var allowedEnvironmentModuleStatusEnumValues = []EnvironmentModuleStatus{
	EnvironmentModuleStatusRunning,
	EnvironmentModuleStatusUpdating,
	EnvironmentModuleStatusError,
	EnvironmentModuleStatusDisabled,
}

// GetAllowedValues returns the list of possible values.
func (v *EnvironmentModuleStatus) GetAllowedValues() []EnvironmentModuleStatus {
	return allowedEnvironmentModuleStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EnvironmentModuleStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EnvironmentModuleStatus(value)
	return nil
}

// NewEnvironmentModuleStatusFromValue returns a pointer to a valid EnvironmentModuleStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEnvironmentModuleStatusFromValue(v string) (*EnvironmentModuleStatus, error) {
	ev := EnvironmentModuleStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EnvironmentModuleStatus: valid values are %v", v, allowedEnvironmentModuleStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EnvironmentModuleStatus) IsValid() bool {
	for _, existing := range allowedEnvironmentModuleStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnvironmentModuleStatus value.
func (v EnvironmentModuleStatus) Ptr() *EnvironmentModuleStatus {
	return &v
}
