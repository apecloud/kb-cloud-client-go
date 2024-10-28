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
	ENVIRONMENTSTATE_PENDING      EnvironmentState = "PENDING"
	ENVIRONMENTSTATE_REGISTERED   EnvironmentState = "REGISTERED"
	ENVIRONMENTSTATE_PROVISIONING EnvironmentState = "PROVISIONING"
	ENVIRONMENTSTATE_NOTREADY     EnvironmentState = "NOTREADY"
	ENVIRONMENTSTATE_READY        EnvironmentState = "READY"
	ENVIRONMENTSTATE_WARNING      EnvironmentState = "WARNING"
	ENVIRONMENTSTATE_UNREACHABLE  EnvironmentState = "UNREACHABLE"
	ENVIRONMENTSTATE_DELETING     EnvironmentState = "DELETING"
	ENVIRONMENTSTATE_OUTOFSTOCK   EnvironmentState = "OUTOFSTOCK"
	ENVIRONMENTSTATE_UPDATING     EnvironmentState = "UPDATING"
)

var allowedEnvironmentStateEnumValues = []EnvironmentState{
	ENVIRONMENTSTATE_PENDING,
	ENVIRONMENTSTATE_REGISTERED,
	ENVIRONMENTSTATE_PROVISIONING,
	ENVIRONMENTSTATE_NOTREADY,
	ENVIRONMENTSTATE_READY,
	ENVIRONMENTSTATE_WARNING,
	ENVIRONMENTSTATE_UNREACHABLE,
	ENVIRONMENTSTATE_DELETING,
	ENVIRONMENTSTATE_OUTOFSTOCK,
	ENVIRONMENTSTATE_UPDATING,
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
