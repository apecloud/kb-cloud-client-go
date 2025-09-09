// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// InspectionScriptCategoryV2 Specifies the category of the inspection script.
type InspectionScriptCategoryV2 string

// List of InspectionScriptCategoryV2.
const (
	InspectionScriptCategoryV2Performance  InspectionScriptCategoryV2 = "performance"
	InspectionScriptCategoryV2Availability InspectionScriptCategoryV2 = "availability"
	InspectionScriptCategoryV2Security     InspectionScriptCategoryV2 = "security"
	InspectionScriptCategoryV2Backup       InspectionScriptCategoryV2 = "backup"
	InspectionScriptCategoryV2Other        InspectionScriptCategoryV2 = "other"
)

var allowedInspectionScriptCategoryV2EnumValues = []InspectionScriptCategoryV2{
	InspectionScriptCategoryV2Performance,
	InspectionScriptCategoryV2Availability,
	InspectionScriptCategoryV2Security,
	InspectionScriptCategoryV2Backup,
	InspectionScriptCategoryV2Other,
}

// GetAllowedValues returns the list of possible values.
func (v *InspectionScriptCategoryV2) GetAllowedValues() []InspectionScriptCategoryV2 {
	return allowedInspectionScriptCategoryV2EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *InspectionScriptCategoryV2) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = InspectionScriptCategoryV2(value)
	return nil
}

// NewInspectionScriptCategoryV2FromValue returns a pointer to a valid InspectionScriptCategoryV2
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewInspectionScriptCategoryV2FromValue(v string) (*InspectionScriptCategoryV2, error) {
	ev := InspectionScriptCategoryV2(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for InspectionScriptCategoryV2: valid values are %v", v, allowedInspectionScriptCategoryV2EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v InspectionScriptCategoryV2) IsValid() bool {
	for _, existing := range allowedInspectionScriptCategoryV2EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InspectionScriptCategoryV2 value.
func (v InspectionScriptCategoryV2) Ptr() *InspectionScriptCategoryV2 {
	return &v
}
