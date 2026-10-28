// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ObjectPrivilegeSubjectType Native authorization subject kind. Does not create accounts or roles, or change memberships.
type ObjectPrivilegeSubjectType string

// List of ObjectPrivilegeSubjectType.
const (
	ObjectPrivilegeSubjectTypeAccount ObjectPrivilegeSubjectType = "account"
	ObjectPrivilegeSubjectTypeRole    ObjectPrivilegeSubjectType = "role"
)

var allowedObjectPrivilegeSubjectTypeEnumValues = []ObjectPrivilegeSubjectType{
	ObjectPrivilegeSubjectTypeAccount,
	ObjectPrivilegeSubjectTypeRole,
}

// GetAllowedValues returns the list of possible values.
func (v *ObjectPrivilegeSubjectType) GetAllowedValues() []ObjectPrivilegeSubjectType {
	return allowedObjectPrivilegeSubjectTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObjectPrivilegeSubjectType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObjectPrivilegeSubjectType(value)
	return nil
}

// NewObjectPrivilegeSubjectTypeFromValue returns a pointer to a valid ObjectPrivilegeSubjectType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObjectPrivilegeSubjectTypeFromValue(v string) (*ObjectPrivilegeSubjectType, error) {
	ev := ObjectPrivilegeSubjectType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObjectPrivilegeSubjectType: valid values are %v", v, allowedObjectPrivilegeSubjectTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObjectPrivilegeSubjectType) IsValid() bool {
	for _, existing := range allowedObjectPrivilegeSubjectTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObjectPrivilegeSubjectType value.
func (v ObjectPrivilegeSubjectType) Ptr() *ObjectPrivilegeSubjectType {
	return &v
}
