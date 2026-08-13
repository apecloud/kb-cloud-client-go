// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DmsExecutionPlanRawFormat string

// List of DmsExecutionPlanRawFormat.
const (
	DmsExecutionPlanRawFormatJson  DmsExecutionPlanRawFormat = "json"
	DmsExecutionPlanRawFormatXml   DmsExecutionPlanRawFormat = "xml"
	DmsExecutionPlanRawFormatText  DmsExecutionPlanRawFormat = "text"
	DmsExecutionPlanRawFormatTable DmsExecutionPlanRawFormat = "table"
)

var allowedDmsExecutionPlanRawFormatEnumValues = []DmsExecutionPlanRawFormat{
	DmsExecutionPlanRawFormatJson,
	DmsExecutionPlanRawFormatXml,
	DmsExecutionPlanRawFormatText,
	DmsExecutionPlanRawFormatTable,
}

// GetAllowedValues returns the list of possible values.
func (v *DmsExecutionPlanRawFormat) GetAllowedValues() []DmsExecutionPlanRawFormat {
	return allowedDmsExecutionPlanRawFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DmsExecutionPlanRawFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DmsExecutionPlanRawFormat(value)
	return nil
}

// NewDmsExecutionPlanRawFormatFromValue returns a pointer to a valid DmsExecutionPlanRawFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDmsExecutionPlanRawFormatFromValue(v string) (*DmsExecutionPlanRawFormat, error) {
	ev := DmsExecutionPlanRawFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DmsExecutionPlanRawFormat: valid values are %v", v, allowedDmsExecutionPlanRawFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DmsExecutionPlanRawFormat) IsValid() bool {
	for _, existing := range allowedDmsExecutionPlanRawFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DmsExecutionPlanRawFormat value.
func (v DmsExecutionPlanRawFormat) Ptr() *DmsExecutionPlanRawFormat {
	return &v
}
