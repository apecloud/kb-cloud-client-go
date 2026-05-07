// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type InitOptionType string

// List of InitOptionType.
const (
	InitOptionTypeTimezone InitOptionType = "timezone"
	InitOptionTypeEnum     InitOptionType = "enum"
	InitOptionTypeString   InitOptionType = "string"
)

var allowedInitOptionTypeEnumValues = []InitOptionType{
	InitOptionTypeTimezone,
	InitOptionTypeEnum,
	InitOptionTypeString,
}

// GetAllowedValues returns the list of possible values.
func (v *InitOptionType) GetAllowedValues() []InitOptionType {
	return allowedInitOptionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *InitOptionType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = InitOptionType(value)
	return nil
}

// NewInitOptionTypeFromValue returns a pointer to a valid InitOptionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewInitOptionTypeFromValue(v string) (*InitOptionType, error) {
	ev := InitOptionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for InitOptionType: valid values are %v", v, allowedInitOptionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v InitOptionType) IsValid() bool {
	for _, existing := range allowedInitOptionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InitOptionType value.
func (v InitOptionType) Ptr() *InitOptionType {
	return &v
}
