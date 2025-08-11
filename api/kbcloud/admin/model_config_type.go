// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ConfigType the type of config
type ConfigType string

// List of ConfigType.
const (
	ConfigTypeString  ConfigType = "string"
	ConfigTypeInteger ConfigType = "integer"
	ConfigTypeFloat   ConfigType = "float"
)

var allowedConfigTypeEnumValues = []ConfigType{
	ConfigTypeString,
	ConfigTypeInteger,
	ConfigTypeFloat,
}

// GetAllowedValues returns the list of possible values.
func (v *ConfigType) GetAllowedValues() []ConfigType {
	return allowedConfigTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ConfigType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ConfigType(value)
	return nil
}

// NewConfigTypeFromValue returns a pointer to a valid ConfigType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewConfigTypeFromValue(v string) (*ConfigType, error) {
	ev := ConfigType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ConfigType: valid values are %v", v, allowedConfigTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ConfigType) IsValid() bool {
	for _, existing := range allowedConfigTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConfigType value.
func (v ConfigType) Ptr() *ConfigType {
	return &v
}
