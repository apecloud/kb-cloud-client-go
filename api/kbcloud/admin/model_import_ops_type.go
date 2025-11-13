// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ImportOpsType Import task operation type
type ImportOpsType string

// List of ImportOpsType.
const (
	ImportOpsTypePause  ImportOpsType = "pause"
	ImportOpsTypeResume ImportOpsType = "resume"
)

var allowedImportOpsTypeEnumValues = []ImportOpsType{
	ImportOpsTypePause,
	ImportOpsTypeResume,
}

// GetAllowedValues returns the list of possible values.
func (v *ImportOpsType) GetAllowedValues() []ImportOpsType {
	return allowedImportOpsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ImportOpsType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ImportOpsType(value)
	return nil
}

// NewImportOpsTypeFromValue returns a pointer to a valid ImportOpsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewImportOpsTypeFromValue(v string) (*ImportOpsType, error) {
	ev := ImportOpsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ImportOpsType: valid values are %v", v, allowedImportOpsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ImportOpsType) IsValid() bool {
	for _, existing := range allowedImportOpsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ImportOpsType value.
func (v ImportOpsType) Ptr() *ImportOpsType {
	return &v
}
