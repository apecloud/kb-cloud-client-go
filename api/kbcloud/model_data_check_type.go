// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DataCheckType Type of check to run. A standalone check runs exactly one type.
type DataCheckType string

// List of DataCheckType.
const (
	DataCheckTypeStruct   DataCheckType = "struct"
	DataCheckTypeSnapshot DataCheckType = "snapshot"
)

var allowedDataCheckTypeEnumValues = []DataCheckType{
	DataCheckTypeStruct,
	DataCheckTypeSnapshot,
}

// GetAllowedValues returns the list of possible values.
func (v *DataCheckType) GetAllowedValues() []DataCheckType {
	return allowedDataCheckTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DataCheckType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DataCheckType(value)
	return nil
}

// NewDataCheckTypeFromValue returns a pointer to a valid DataCheckType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDataCheckTypeFromValue(v string) (*DataCheckType, error) {
	ev := DataCheckType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DataCheckType: valid values are %v", v, allowedDataCheckTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DataCheckType) IsValid() bool {
	for _, existing := range allowedDataCheckTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataCheckType value.
func (v DataCheckType) Ptr() *DataCheckType {
	return &v
}
