// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EnvironmentState Output only. State of the Environment resource
type EnvironmentState string

// List of EnvironmentState.
const (
	EnvironmentStatePending      EnvironmentState = "PENDING"
	EnvironmentStateRegistered   EnvironmentState = "REGISTERED"
	EnvironmentStateProvisioning EnvironmentState = "PROVISIONING"
	EnvironmentStateNotready     EnvironmentState = "NOTREADY"
	EnvironmentStateReady        EnvironmentState = "READY"
	EnvironmentStateWarning      EnvironmentState = "WARNING"
	EnvironmentStateUnreachable  EnvironmentState = "UNREACHABLE"
	EnvironmentStateDeleting     EnvironmentState = "DELETING"
	EnvironmentStateOutofstock   EnvironmentState = "OUTOFSTOCK"
	EnvironmentStateUpdating     EnvironmentState = "UPDATING"
)

var allowedEnvironmentStateEnumValues = []EnvironmentState{
	EnvironmentStatePending,
	EnvironmentStateRegistered,
	EnvironmentStateProvisioning,
	EnvironmentStateNotready,
	EnvironmentStateReady,
	EnvironmentStateWarning,
	EnvironmentStateUnreachable,
	EnvironmentStateDeleting,
	EnvironmentStateOutofstock,
	EnvironmentStateUpdating,
}

// GetAllowedValues returns the list of possible values.
func (v *EnvironmentState) GetAllowedValues() []EnvironmentState {
	return allowedEnvironmentStateEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EnvironmentState) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EnvironmentState(value)
	return nil
}

// NewEnvironmentStateFromValue returns a pointer to a valid EnvironmentState
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEnvironmentStateFromValue(v string) (*EnvironmentState, error) {
	ev := EnvironmentState(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EnvironmentState: valid values are %v", v, allowedEnvironmentStateEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EnvironmentState) IsValid() bool {
	for _, existing := range allowedEnvironmentStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnvironmentState value.
func (v EnvironmentState) Ptr() *EnvironmentState {
	return &v
}
