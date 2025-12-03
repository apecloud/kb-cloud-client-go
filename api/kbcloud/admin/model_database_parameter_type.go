// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DatabaseParameterType The type of the parameter value
type DatabaseParameterType string

// List of DatabaseParameterType.
const (
	DatabaseParameterTypeString  DatabaseParameterType = "string"
	DatabaseParameterTypeInteger DatabaseParameterType = "integer"
	DatabaseParameterTypeNumber  DatabaseParameterType = "number"
	DatabaseParameterTypeBoolean DatabaseParameterType = "boolean"
	DatabaseParameterTypeArray   DatabaseParameterType = "array"
)

var allowedDatabaseParameterTypeEnumValues = []DatabaseParameterType{
	DatabaseParameterTypeString,
	DatabaseParameterTypeInteger,
	DatabaseParameterTypeNumber,
	DatabaseParameterTypeBoolean,
	DatabaseParameterTypeArray,
}

// GetAllowedValues returns the list of possible values.
func (v *DatabaseParameterType) GetAllowedValues() []DatabaseParameterType {
	return allowedDatabaseParameterTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DatabaseParameterType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DatabaseParameterType(value)
	return nil
}

// NewDatabaseParameterTypeFromValue returns a pointer to a valid DatabaseParameterType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDatabaseParameterTypeFromValue(v string) (*DatabaseParameterType, error) {
	ev := DatabaseParameterType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DatabaseParameterType: valid values are %v", v, allowedDatabaseParameterTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DatabaseParameterType) IsValid() bool {
	for _, existing := range allowedDatabaseParameterTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DatabaseParameterType value.
func (v DatabaseParameterType) Ptr() *DatabaseParameterType {
	return &v
}
