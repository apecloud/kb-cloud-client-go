// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ImportFieldType Import field type
type ImportFieldType string

// List of ImportFieldType.
const (
	String  ImportFieldType = "string"
	Integer ImportFieldType = "integer"
	Number  ImportFieldType = "number"
	Boolean ImportFieldType = "boolean"
	Enum    ImportFieldType = "enum"
)

var allowedImportFieldTypeEnumValues = []ImportFieldType{
	String,
	Integer,
	Number,
	Boolean,
	Enum,
}

// GetAllowedValues returns the list of possible values.
func (v *ImportFieldType) GetAllowedValues() []ImportFieldType {
	return allowedImportFieldTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ImportFieldType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ImportFieldType(value)
	return nil
}

// NewImportFieldTypeFromValue returns a pointer to a valid ImportFieldType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewImportFieldTypeFromValue(v string) (*ImportFieldType, error) {
	ev := ImportFieldType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ImportFieldType: valid values are %v", v, allowedImportFieldTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ImportFieldType) IsValid() bool {
	for _, existing := range allowedImportFieldTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ImportFieldType value.
func (v ImportFieldType) Ptr() *ImportFieldType {
	return &v
}
