// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// AccountRoleType Role name should be one of [SUPERUSER, BASICUSER].
type AccountRoleType string

// List of AccountRoleType.
const (
	ACCOUNTROLETYPE_SUPERUSER AccountRoleType = "SUPERUSER"
	ACCOUNTROLETYPE_BASICUSER AccountRoleType = "BASICUSER"
)

var allowedAccountRoleTypeEnumValues = []AccountRoleType{
	ACCOUNTROLETYPE_SUPERUSER,
	ACCOUNTROLETYPE_BASICUSER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AccountRoleType) GetAllowedValues() []AccountRoleType {
	return allowedAccountRoleTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AccountRoleType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AccountRoleType(value)
	return nil
}

// NewAccountRoleTypeFromValue returns a pointer to a valid AccountRoleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAccountRoleTypeFromValue(v string) (*AccountRoleType, error) {
	ev := AccountRoleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AccountRoleType: valid values are %v", v, allowedAccountRoleTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AccountRoleType) IsValid() bool {
	for _, existing := range allowedAccountRoleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccountRoleType value.
func (v AccountRoleType) Ptr() *AccountRoleType {
	return &v
}
