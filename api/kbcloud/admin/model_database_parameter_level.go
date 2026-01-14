// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DatabaseParameterLevel The level of the parameter
type DatabaseParameterLevel string

// List of DatabaseParameterLevel.
const (
	DatabaseParameterLevelNormal    DatabaseParameterLevel = "normal"
	DatabaseParameterLevelImportant DatabaseParameterLevel = "important"
	DatabaseParameterLevelProtected DatabaseParameterLevel = "protected"
)

var allowedDatabaseParameterLevelEnumValues = []DatabaseParameterLevel{
	DatabaseParameterLevelNormal,
	DatabaseParameterLevelImportant,
	DatabaseParameterLevelProtected,
}

// GetAllowedValues returns the list of possible values.
func (v *DatabaseParameterLevel) GetAllowedValues() []DatabaseParameterLevel {
	return allowedDatabaseParameterLevelEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DatabaseParameterLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DatabaseParameterLevel(value)
	return nil
}

// NewDatabaseParameterLevelFromValue returns a pointer to a valid DatabaseParameterLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDatabaseParameterLevelFromValue(v string) (*DatabaseParameterLevel, error) {
	ev := DatabaseParameterLevel(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DatabaseParameterLevel: valid values are %v", v, allowedDatabaseParameterLevelEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DatabaseParameterLevel) IsValid() bool {
	for _, existing := range allowedDatabaseParameterLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DatabaseParameterLevel value.
func (v DatabaseParameterLevel) Ptr() *DatabaseParameterLevel {
	return &v
}
