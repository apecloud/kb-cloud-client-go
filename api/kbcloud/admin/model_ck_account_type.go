// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// CkAccountType The account role type.
type CkAccountType string

// List of CkAccountType.
const (
	CkAccountTypeSuperuser CkAccountType = "SUPERUSER"
	CkAccountTypeBasicuser CkAccountType = "BASICUSER"
	CkAccountTypeRoot      CkAccountType = "ROOT"
)

var allowedCkAccountTypeEnumValues = []CkAccountType{
	CkAccountTypeSuperuser,
	CkAccountTypeBasicuser,
	CkAccountTypeRoot,
}

// GetAllowedValues returns the list of possible values.
func (v *CkAccountType) GetAllowedValues() []CkAccountType {
	return allowedCkAccountTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CkAccountType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CkAccountType(value)
	return nil
}

// NewCkAccountTypeFromValue returns a pointer to a valid CkAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCkAccountTypeFromValue(v string) (*CkAccountType, error) {
	ev := CkAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CkAccountType: valid values are %v", v, allowedCkAccountTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CkAccountType) IsValid() bool {
	for _, existing := range allowedCkAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CkAccountType value.
func (v CkAccountType) Ptr() *CkAccountType {
	return &v
}
