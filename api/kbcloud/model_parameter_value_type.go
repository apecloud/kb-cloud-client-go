// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ParameterValueType string

// List of ParameterValueType.
const (
	ParameterValueTypeString  ParameterValueType = "string"
	ParameterValueTypeInteger ParameterValueType = "integer"
)

var allowedParameterValueTypeEnumValues = []ParameterValueType{
	ParameterValueTypeString,
	ParameterValueTypeInteger,
}

// GetAllowedValues returns the list of possible values.
func (v *ParameterValueType) GetAllowedValues() []ParameterValueType {
	return allowedParameterValueTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ParameterValueType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ParameterValueType(value)
	return nil
}

// NewParameterValueTypeFromValue returns a pointer to a valid ParameterValueType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewParameterValueTypeFromValue(v string) (*ParameterValueType, error) {
	ev := ParameterValueType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ParameterValueType: valid values are %v", v, allowedParameterValueTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ParameterValueType) IsValid() bool {
	for _, existing := range allowedParameterValueTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ParameterValueType value.
func (v ParameterValueType) Ptr() *ParameterValueType {
	return &v
}
