// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type CdcConfigFormatType string

// List of CdcConfigFormatType.
const (
	CdcConfigFormatTypeIni        CdcConfigFormatType = "Ini"
	CdcConfigFormatTypeToml       CdcConfigFormatType = "Toml"
	CdcConfigFormatTypeProperties CdcConfigFormatType = "Properties"
)

var allowedCdcConfigFormatTypeEnumValues = []CdcConfigFormatType{
	CdcConfigFormatTypeIni,
	CdcConfigFormatTypeToml,
	CdcConfigFormatTypeProperties,
}

// GetAllowedValues returns the list of possible values.
func (v *CdcConfigFormatType) GetAllowedValues() []CdcConfigFormatType {
	return allowedCdcConfigFormatTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CdcConfigFormatType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CdcConfigFormatType(value)
	return nil
}

// NewCdcConfigFormatTypeFromValue returns a pointer to a valid CdcConfigFormatType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCdcConfigFormatTypeFromValue(v string) (*CdcConfigFormatType, error) {
	ev := CdcConfigFormatType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CdcConfigFormatType: valid values are %v", v, allowedCdcConfigFormatTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CdcConfigFormatType) IsValid() bool {
	for _, existing := range allowedCdcConfigFormatTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CdcConfigFormatType value.
func (v CdcConfigFormatType) Ptr() *CdcConfigFormatType {
	return &v
}
