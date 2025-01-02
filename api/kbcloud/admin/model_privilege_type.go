// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// PrivilegeType The type of privilege.
type PrivilegeType string

// List of PrivilegeType.
const (
	PrivilegeTypeDbadmin   PrivilegeType = "DBADMIN"
	PrivilegeTypeReadwrite PrivilegeType = "READWRITE"
	PrivilegeTypeReadonly  PrivilegeType = "READONLY"
	PrivilegeTypeDdlonly   PrivilegeType = "DDLONLY"
	PrivilegeTypeDmlonly   PrivilegeType = "DMLONLY"
)

var allowedPrivilegeTypeEnumValues = []PrivilegeType{
	PrivilegeTypeDbadmin,
	PrivilegeTypeReadwrite,
	PrivilegeTypeReadonly,
	PrivilegeTypeDdlonly,
	PrivilegeTypeDmlonly,
}

// GetAllowedValues returns the list of possible values.
func (v *PrivilegeType) GetAllowedValues() []PrivilegeType {
	return allowedPrivilegeTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PrivilegeType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PrivilegeType(value)
	return nil
}

// NewPrivilegeTypeFromValue returns a pointer to a valid PrivilegeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPrivilegeTypeFromValue(v string) (*PrivilegeType, error) {
	ev := PrivilegeType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PrivilegeType: valid values are %v", v, allowedPrivilegeTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PrivilegeType) IsValid() bool {
	for _, existing := range allowedPrivilegeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PrivilegeType value.
func (v PrivilegeType) Ptr() *PrivilegeType {
	return &v
}
