// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type RedisOperationStatus string

// List of RedisOperationStatus.
const (
	RedisOperationStatusSuccess RedisOperationStatus = "success"
	RedisOperationStatusError   RedisOperationStatus = "error"
	RedisOperationStatusBlocked RedisOperationStatus = "blocked"
)

var allowedRedisOperationStatusEnumValues = []RedisOperationStatus{
	RedisOperationStatusSuccess,
	RedisOperationStatusError,
	RedisOperationStatusBlocked,
}

// GetAllowedValues returns the list of possible values.
func (v *RedisOperationStatus) GetAllowedValues() []RedisOperationStatus {
	return allowedRedisOperationStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RedisOperationStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RedisOperationStatus(value)
	return nil
}

// NewRedisOperationStatusFromValue returns a pointer to a valid RedisOperationStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRedisOperationStatusFromValue(v string) (*RedisOperationStatus, error) {
	ev := RedisOperationStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RedisOperationStatus: valid values are %v", v, allowedRedisOperationStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RedisOperationStatus) IsValid() bool {
	for _, existing := range allowedRedisOperationStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RedisOperationStatus value.
func (v RedisOperationStatus) Ptr() *RedisOperationStatus {
	return &v
}
