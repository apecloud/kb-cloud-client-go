// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package data

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)



// RedisPrivilegeType Redis user privileges
type RedisPrivilegeType string

// List of RedisPrivilegeType.
const (
	RedisPrivilegeTypeReadwrite RedisPrivilegeType = "READWRITE"
	RedisPrivilegeTypeReadonly RedisPrivilegeType = "READONLY"
	RedisPrivilegeTypeUnknown RedisPrivilegeType = "UNKNOWN"
)

var allowedRedisPrivilegeTypeEnumValues = []RedisPrivilegeType{
	RedisPrivilegeTypeReadwrite,
	RedisPrivilegeTypeReadonly,
	RedisPrivilegeTypeUnknown,
}

// GetAllowedValues returns the list of possible values.
func (v *RedisPrivilegeType) GetAllowedValues() []RedisPrivilegeType {
	return allowedRedisPrivilegeTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RedisPrivilegeType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RedisPrivilegeType(value)
	return nil
}

// NewRedisPrivilegeTypeFromValue returns a pointer to a valid RedisPrivilegeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRedisPrivilegeTypeFromValue(v string) (*RedisPrivilegeType, error) {
	ev := RedisPrivilegeType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RedisPrivilegeType: valid values are %v", v, allowedRedisPrivilegeTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RedisPrivilegeType) IsValid() bool {
	for _, existing := range allowedRedisPrivilegeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RedisPrivilegeType value.
func (v RedisPrivilegeType) Ptr() *RedisPrivilegeType {
	return &v
}
