// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DmsGenerateDdlOperationType operation type
type DmsGenerateDdlOperationType string

// List of DmsGenerateDdlOperationType.
const (
	DmsGenerateDdlOperationTypeTable DmsGenerateDdlOperationType = "table"
	DmsGenerateDdlOperationTypeView  DmsGenerateDdlOperationType = "view"
)

var allowedDmsGenerateDdlOperationTypeEnumValues = []DmsGenerateDdlOperationType{
	DmsGenerateDdlOperationTypeTable,
	DmsGenerateDdlOperationTypeView,
}

// GetAllowedValues returns the list of possible values.
func (v *DmsGenerateDdlOperationType) GetAllowedValues() []DmsGenerateDdlOperationType {
	return allowedDmsGenerateDdlOperationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DmsGenerateDdlOperationType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DmsGenerateDdlOperationType(value)
	return nil
}

// NewDmsGenerateDdlOperationTypeFromValue returns a pointer to a valid DmsGenerateDdlOperationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDmsGenerateDdlOperationTypeFromValue(v string) (*DmsGenerateDdlOperationType, error) {
	ev := DmsGenerateDdlOperationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DmsGenerateDdlOperationType: valid values are %v", v, allowedDmsGenerateDdlOperationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DmsGenerateDdlOperationType) IsValid() bool {
	for _, existing := range allowedDmsGenerateDdlOperationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DmsGenerateDdlOperationType value.
func (v DmsGenerateDdlOperationType) Ptr() *DmsGenerateDdlOperationType {
	return &v
}
