// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
)

// PlatformParameterCategory platformParameter category, representing the platformParameter group, such as 'recycle bin'
type PlatformParameterCategory string

// List of PlatformParameterCategory.
const (
	PlatformParameterCategoryRecycleBin PlatformParameterCategory = "recycle-bin"
)

var allowedPlatformParameterCategoryEnumValues = []PlatformParameterCategory{
	PlatformParameterCategoryRecycleBin,
}

// GetAllowedValues returns the list of possible values.
func (v *PlatformParameterCategory) GetAllowedValues() []PlatformParameterCategory {
	return allowedPlatformParameterCategoryEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PlatformParameterCategory) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PlatformParameterCategory(value)
	return nil
}

// NewPlatformParameterCategoryFromValue returns a pointer to a valid PlatformParameterCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPlatformParameterCategoryFromValue(v string) (*PlatformParameterCategory, error) {
	ev := PlatformParameterCategory(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PlatformParameterCategory: valid values are %v", v, allowedPlatformParameterCategoryEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PlatformParameterCategory) IsValid() bool {
	for _, existing := range allowedPlatformParameterCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PlatformParameterCategory value.
func (v PlatformParameterCategory) Ptr() *PlatformParameterCategory {
	return &v
}
