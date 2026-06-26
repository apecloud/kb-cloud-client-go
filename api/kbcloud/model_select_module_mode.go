// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type SelectModuleMode string

// List of SelectModuleMode.
const (
	SelectModuleModeSingle   SelectModuleMode = "single"
	SelectModuleModeMultiple SelectModuleMode = "multiple"
)

var allowedSelectModuleModeEnumValues = []SelectModuleMode{
	SelectModuleModeSingle,
	SelectModuleModeMultiple,
}

// GetAllowedValues returns the list of possible values.
func (v *SelectModuleMode) GetAllowedValues() []SelectModuleMode {
	return allowedSelectModuleModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SelectModuleMode) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SelectModuleMode(value)
	return nil
}

// NewSelectModuleModeFromValue returns a pointer to a valid SelectModuleMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSelectModuleModeFromValue(v string) (*SelectModuleMode, error) {
	ev := SelectModuleMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SelectModuleMode: valid values are %v", v, allowedSelectModuleModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SelectModuleMode) IsValid() bool {
	for _, existing := range allowedSelectModuleModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SelectModuleMode value.
func (v SelectModuleMode) Ptr() *SelectModuleMode {
	return &v
}
