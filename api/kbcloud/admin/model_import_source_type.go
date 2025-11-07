// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ImportSourceType Import source type
type ImportSourceType string

// List of ImportSourceType.
const (
	ImportSourceTypeMysql ImportSourceType = "mysql"
)

var allowedImportSourceTypeEnumValues = []ImportSourceType{
	ImportSourceTypeMysql,
}

// GetAllowedValues returns the list of possible values.
func (v *ImportSourceType) GetAllowedValues() []ImportSourceType {
	return allowedImportSourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ImportSourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ImportSourceType(value)
	return nil
}

// NewImportSourceTypeFromValue returns a pointer to a valid ImportSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewImportSourceTypeFromValue(v string) (*ImportSourceType, error) {
	ev := ImportSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ImportSourceType: valid values are %v", v, allowedImportSourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ImportSourceType) IsValid() bool {
	for _, existing := range allowedImportSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ImportSourceType value.
func (v ImportSourceType) Ptr() *ImportSourceType {
	return &v
}
