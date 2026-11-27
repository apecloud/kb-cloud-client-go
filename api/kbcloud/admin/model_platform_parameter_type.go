// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// PlatformParameterType type of platformParameter value. Values are always transferred as strings;
// the type declares how the value must be interpreted and validated:
// - string: free-form text, may be constrained by enum
// - int: base-10 integer, may be constrained by min/max
// - boolean: only "true" or "false"
// - json: a JSON object, editable as structured text
type PlatformParameterType string

// List of PlatformParameterType.
const (
	PlatformParameterTypeString  PlatformParameterType = "string"
	PlatformParameterTypeInt     PlatformParameterType = "int"
	PlatformParameterTypeBoolean PlatformParameterType = "boolean"
	PlatformParameterTypeJson    PlatformParameterType = "json"
)

var allowedPlatformParameterTypeEnumValues = []PlatformParameterType{
	PlatformParameterTypeString,
	PlatformParameterTypeInt,
	PlatformParameterTypeBoolean,
	PlatformParameterTypeJson,
}

// GetAllowedValues returns the list of possible values.
func (v *PlatformParameterType) GetAllowedValues() []PlatformParameterType {
	return allowedPlatformParameterTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PlatformParameterType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PlatformParameterType(value)
	return nil
}

// NewPlatformParameterTypeFromValue returns a pointer to a valid PlatformParameterType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPlatformParameterTypeFromValue(v string) (*PlatformParameterType, error) {
	ev := PlatformParameterType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PlatformParameterType: valid values are %v", v, allowedPlatformParameterTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PlatformParameterType) IsValid() bool {
	for _, existing := range allowedPlatformParameterTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PlatformParameterType value.
func (v PlatformParameterType) Ptr() *PlatformParameterType {
	return &v
}
