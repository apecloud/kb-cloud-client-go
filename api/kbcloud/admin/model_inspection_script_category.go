// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// InspectionScriptCategory Current inspection category vocabulary. Existing legacy categories are normalized by the backend before they are returned.
type InspectionScriptCategory string

// List of InspectionScriptCategory.
const (
	InspectionScriptCategoryAvailability     InspectionScriptCategory = "availability"
	InspectionScriptCategoryPerformance      InspectionScriptCategory = "performance"
	InspectionScriptCategoryConnections      InspectionScriptCategory = "connections"
	InspectionScriptCategoryResource         InspectionScriptCategory = "resource"
	InspectionScriptCategoryCapacity         InspectionScriptCategory = "capacity"
	InspectionScriptCategoryVacuumWraparound InspectionScriptCategory = "vacuumWraparound"
	InspectionScriptCategoryBackup           InspectionScriptCategory = "backup"
	InspectionScriptCategorySecurity         InspectionScriptCategory = "security"
)

var allowedInspectionScriptCategoryEnumValues = []InspectionScriptCategory{
	InspectionScriptCategoryAvailability,
	InspectionScriptCategoryPerformance,
	InspectionScriptCategoryConnections,
	InspectionScriptCategoryResource,
	InspectionScriptCategoryCapacity,
	InspectionScriptCategoryVacuumWraparound,
	InspectionScriptCategoryBackup,
	InspectionScriptCategorySecurity,
}

// GetAllowedValues returns the list of possible values.
func (v *InspectionScriptCategory) GetAllowedValues() []InspectionScriptCategory {
	return allowedInspectionScriptCategoryEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *InspectionScriptCategory) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = InspectionScriptCategory(value)
	return nil
}

// NewInspectionScriptCategoryFromValue returns a pointer to a valid InspectionScriptCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewInspectionScriptCategoryFromValue(v string) (*InspectionScriptCategory, error) {
	ev := InspectionScriptCategory(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for InspectionScriptCategory: valid values are %v", v, allowedInspectionScriptCategoryEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v InspectionScriptCategory) IsValid() bool {
	for _, existing := range allowedInspectionScriptCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InspectionScriptCategory value.
func (v InspectionScriptCategory) Ptr() *InspectionScriptCategory {
	return &v
}
