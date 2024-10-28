// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// AccountListRoleType The user role name, should be one of [ROOT, SUPERUSER, BASICUSER].
type AccountListRoleType string

// List of AccountListRoleType.
const (
	ACCOUNTLISTROLETYPE_SUPERUSER AccountListRoleType = "SUPERUSER"
	ACCOUNTLISTROLETYPE_BASICUSER AccountListRoleType = "BASICUSER"
	ACCOUNTLISTROLETYPE_ROOT      AccountListRoleType = "ROOT"
)

var allowedAccountListRoleTypeEnumValues = []AccountListRoleType{
	ACCOUNTLISTROLETYPE_SUPERUSER,
	ACCOUNTLISTROLETYPE_BASICUSER,
	ACCOUNTLISTROLETYPE_ROOT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AccountListRoleType) GetAllowedValues() []AccountListRoleType {
	return allowedAccountListRoleTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AccountListRoleType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AccountListRoleType(value)
	return nil
}

// NewAccountListRoleTypeFromValue returns a pointer to a valid AccountListRoleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAccountListRoleTypeFromValue(v string) (*AccountListRoleType, error) {
	ev := AccountListRoleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AccountListRoleType: valid values are %v", v, allowedAccountListRoleTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AccountListRoleType) IsValid() bool {
	for _, existing := range allowedAccountListRoleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccountListRoleType value.
func (v AccountListRoleType) Ptr() *AccountListRoleType {
	return &v
}
