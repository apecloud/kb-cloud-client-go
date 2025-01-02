// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// YcsbRedisMode Redis mode
type YcsbRedisMode string

// List of YcsbRedisMode.
const (
	YcsbRedisModeSingle   YcsbRedisMode = "single"
	YcsbRedisModeSentinel YcsbRedisMode = "sentinel"
	YcsbRedisModeCluster  YcsbRedisMode = "cluster"
)

var allowedYcsbRedisModeEnumValues = []YcsbRedisMode{
	YcsbRedisModeSingle,
	YcsbRedisModeSentinel,
	YcsbRedisModeCluster,
}

// GetAllowedValues returns the list of possible values.
func (v *YcsbRedisMode) GetAllowedValues() []YcsbRedisMode {
	return allowedYcsbRedisModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *YcsbRedisMode) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = YcsbRedisMode(value)
	return nil
}

// NewYcsbRedisModeFromValue returns a pointer to a valid YcsbRedisMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewYcsbRedisModeFromValue(v string) (*YcsbRedisMode, error) {
	ev := YcsbRedisMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for YcsbRedisMode: valid values are %v", v, allowedYcsbRedisModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v YcsbRedisMode) IsValid() bool {
	for _, existing := range allowedYcsbRedisModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to YcsbRedisMode value.
func (v YcsbRedisMode) Ptr() *YcsbRedisMode {
	return &v
}
