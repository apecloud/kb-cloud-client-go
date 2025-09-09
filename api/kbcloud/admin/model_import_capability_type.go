// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ImportCapabilityType Import capability type
type ImportCapabilityType string

// List of ImportCapabilityType.
const (
	ImportCapabilityTypeStructure ImportCapabilityType = "structure"
	ImportCapabilityTypeSnapshot  ImportCapabilityType = "snapshot"
	ImportCapabilityTypeCdc       ImportCapabilityType = "cdc"
)

var allowedImportCapabilityTypeEnumValues = []ImportCapabilityType{
	ImportCapabilityTypeStructure,
	ImportCapabilityTypeSnapshot,
	ImportCapabilityTypeCdc,
}

// GetAllowedValues returns the list of possible values.
func (v *ImportCapabilityType) GetAllowedValues() []ImportCapabilityType {
	return allowedImportCapabilityTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ImportCapabilityType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ImportCapabilityType(value)
	return nil
}

// NewImportCapabilityTypeFromValue returns a pointer to a valid ImportCapabilityType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewImportCapabilityTypeFromValue(v string) (*ImportCapabilityType, error) {
	ev := ImportCapabilityType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ImportCapabilityType: valid values are %v", v, allowedImportCapabilityTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ImportCapabilityType) IsValid() bool {
	for _, existing := range allowedImportCapabilityTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ImportCapabilityType value.
func (v ImportCapabilityType) Ptr() *ImportCapabilityType {
	return &v
}
