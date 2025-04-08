// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AccessLevel string

// List of AccessLevel.
const (
	AccessLevelPlatform     AccessLevel = "Platform"
	AccessLevelOrganization AccessLevel = "Organization"
)

var allowedAccessLevelEnumValues = []AccessLevel{
	AccessLevelPlatform,
	AccessLevelOrganization,
}

// GetAllowedValues returns the list of possible values.
func (v *AccessLevel) GetAllowedValues() []AccessLevel {
	return allowedAccessLevelEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AccessLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AccessLevel(value)
	return nil
}

// NewAccessLevelFromValue returns a pointer to a valid AccessLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAccessLevelFromValue(v string) (*AccessLevel, error) {
	ev := AccessLevel(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AccessLevel: valid values are %v", v, allowedAccessLevelEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AccessLevel) IsValid() bool {
	for _, existing := range allowedAccessLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccessLevel value.
func (v AccessLevel) Ptr() *AccessLevel {
	return &v
}
