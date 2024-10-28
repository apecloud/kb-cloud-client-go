// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EnvironmentType Type of this environment
type EnvironmentType string

// List of EnvironmentType.
const (
	ENVIRONMENTTYPE_PUBLIC  EnvironmentType = "public"
	ENVIRONMENTTYPE_PRIVATE EnvironmentType = "private"
)

var allowedEnvironmentTypeEnumValues = []EnvironmentType{
	ENVIRONMENTTYPE_PUBLIC,
	ENVIRONMENTTYPE_PRIVATE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EnvironmentType) GetAllowedValues() []EnvironmentType {
	return allowedEnvironmentTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EnvironmentType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EnvironmentType(value)
	return nil
}

// NewEnvironmentTypeFromValue returns a pointer to a valid EnvironmentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEnvironmentTypeFromValue(v string) (*EnvironmentType, error) {
	ev := EnvironmentType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EnvironmentType: valid values are %v", v, allowedEnvironmentTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EnvironmentType) IsValid() bool {
	for _, existing := range allowedEnvironmentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnvironmentType value.
func (v EnvironmentType) Ptr() *EnvironmentType {
	return &v
}
