// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EngineActionType engine action type
type EngineActionType string

// List of EngineActionType.
const (
	EngineActionTypeEnable  EngineActionType = "enable"
	EngineActionTypeDisable EngineActionType = "disable"
	EngineActionTypeUpgrade EngineActionType = "upgrade"
)

var allowedEngineActionTypeEnumValues = []EngineActionType{
	EngineActionTypeEnable,
	EngineActionTypeDisable,
	EngineActionTypeUpgrade,
}

// GetAllowedValues returns the list of possible values.
func (v *EngineActionType) GetAllowedValues() []EngineActionType {
	return allowedEngineActionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EngineActionType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EngineActionType(value)
	return nil
}

// NewEngineActionTypeFromValue returns a pointer to a valid EngineActionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEngineActionTypeFromValue(v string) (*EngineActionType, error) {
	ev := EngineActionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EngineActionType: valid values are %v", v, allowedEngineActionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EngineActionType) IsValid() bool {
	for _, existing := range allowedEngineActionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EngineActionType value.
func (v EngineActionType) Ptr() *EngineActionType {
	return &v
}
