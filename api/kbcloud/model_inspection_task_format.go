// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// InspectionTaskFormat Specifies the format of the task in the aggregate task result.
type InspectionTaskFormat string

// List of InspectionTaskFormat.
const (
	InspectionTaskFormatJson InspectionTaskFormat = "json"
	InspectionTaskFormatPdf  InspectionTaskFormat = "pdf"
)

var allowedInspectionTaskFormatEnumValues = []InspectionTaskFormat{
	InspectionTaskFormatJson,
	InspectionTaskFormatPdf,
}

// GetAllowedValues returns the list of possible values.
func (v *InspectionTaskFormat) GetAllowedValues() []InspectionTaskFormat {
	return allowedInspectionTaskFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *InspectionTaskFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = InspectionTaskFormat(value)
	return nil
}

// NewInspectionTaskFormatFromValue returns a pointer to a valid InspectionTaskFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewInspectionTaskFormatFromValue(v string) (*InspectionTaskFormat, error) {
	ev := InspectionTaskFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for InspectionTaskFormat: valid values are %v", v, allowedInspectionTaskFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v InspectionTaskFormat) IsValid() bool {
	for _, existing := range allowedInspectionTaskFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InspectionTaskFormat value.
func (v InspectionTaskFormat) Ptr() *InspectionTaskFormat {
	return &v
}
