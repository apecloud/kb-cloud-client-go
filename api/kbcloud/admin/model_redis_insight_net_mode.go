// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type RedisInsightNetMode string

// List of RedisInsightNetMode.
const (
	RedisInsightNetModePrivate RedisInsightNetMode = "PRIVATE"
	RedisInsightNetModePublic  RedisInsightNetMode = "PUBLIC"
)

var allowedRedisInsightNetModeEnumValues = []RedisInsightNetMode{
	RedisInsightNetModePrivate,
	RedisInsightNetModePublic,
}

// GetAllowedValues returns the list of possible values.
func (v *RedisInsightNetMode) GetAllowedValues() []RedisInsightNetMode {
	return allowedRedisInsightNetModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RedisInsightNetMode) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RedisInsightNetMode(value)
	return nil
}

// NewRedisInsightNetModeFromValue returns a pointer to a valid RedisInsightNetMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRedisInsightNetModeFromValue(v string) (*RedisInsightNetMode, error) {
	ev := RedisInsightNetMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RedisInsightNetMode: valid values are %v", v, allowedRedisInsightNetModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RedisInsightNetMode) IsValid() bool {
	for _, existing := range allowedRedisInsightNetModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RedisInsightNetMode value.
func (v RedisInsightNetMode) Ptr() *RedisInsightNetMode {
	return &v
}
