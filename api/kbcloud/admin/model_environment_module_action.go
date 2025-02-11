// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EnvironmentModuleAction Action to perform on the environment module:
// - Enable: Enable the module
// - Disable: Disable the module
// - Upgrade: Upgrade the module to specified version
type EnvironmentModuleAction string

// List of EnvironmentModuleAction.
const (
	EnvironmentModuleActionEnable  EnvironmentModuleAction = "Enable"
	EnvironmentModuleActionDisable EnvironmentModuleAction = "Disable"
	EnvironmentModuleActionUpgrade EnvironmentModuleAction = "Upgrade"
)

var allowedEnvironmentModuleActionEnumValues = []EnvironmentModuleAction{
	EnvironmentModuleActionEnable,
	EnvironmentModuleActionDisable,
	EnvironmentModuleActionUpgrade,
}

// GetAllowedValues returns the list of possible values.
func (v *EnvironmentModuleAction) GetAllowedValues() []EnvironmentModuleAction {
	return allowedEnvironmentModuleActionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EnvironmentModuleAction) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EnvironmentModuleAction(value)
	return nil
}

// NewEnvironmentModuleActionFromValue returns a pointer to a valid EnvironmentModuleAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEnvironmentModuleActionFromValue(v string) (*EnvironmentModuleAction, error) {
	ev := EnvironmentModuleAction(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EnvironmentModuleAction: valid values are %v", v, allowedEnvironmentModuleActionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EnvironmentModuleAction) IsValid() bool {
	for _, existing := range allowedEnvironmentModuleActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnvironmentModuleAction value.
func (v EnvironmentModuleAction) Ptr() *EnvironmentModuleAction {
	return &v
}
