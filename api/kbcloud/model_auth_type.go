// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AuthType string

// List of AuthType.
const (
	AuthTypeBasic AuthType = "basic"
)

var allowedAuthTypeEnumValues = []AuthType{
	AuthTypeBasic,
}

// GetAllowedValues returns the list of possible values.
func (v *AuthType) GetAllowedValues() []AuthType {
	return allowedAuthTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AuthType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AuthType(value)
	return nil
}

// NewAuthTypeFromValue returns a pointer to a valid AuthType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAuthTypeFromValue(v string) (*AuthType, error) {
	ev := AuthType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AuthType: valid values are %v", v, allowedAuthTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AuthType) IsValid() bool {
	for _, existing := range allowedAuthTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AuthType value.
func (v AuthType) Ptr() *AuthType {
	return &v
}
