// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ImportSourceCategory Data source category
type ImportSourceCategory string

// List of ImportSourceCategory.
const (
	ImportSourceCategoryEngine ImportSourceCategory = "engine"
)

var allowedImportSourceCategoryEnumValues = []ImportSourceCategory{
	ImportSourceCategoryEngine,
}

// GetAllowedValues returns the list of possible values.
func (v *ImportSourceCategory) GetAllowedValues() []ImportSourceCategory {
	return allowedImportSourceCategoryEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ImportSourceCategory) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ImportSourceCategory(value)
	return nil
}

// NewImportSourceCategoryFromValue returns a pointer to a valid ImportSourceCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewImportSourceCategoryFromValue(v string) (*ImportSourceCategory, error) {
	ev := ImportSourceCategory(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ImportSourceCategory: valid values are %v", v, allowedImportSourceCategoryEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ImportSourceCategory) IsValid() bool {
	for _, existing := range allowedImportSourceCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ImportSourceCategory value.
func (v ImportSourceCategory) Ptr() *ImportSourceCategory {
	return &v
}
