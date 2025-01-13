// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package data

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type OffsetResetStrategy string

// List of OffsetResetStrategy.
const (
	OffsetResetStrategyEarliest  OffsetResetStrategy = "EARLIEST"
	OffsetResetStrategyLatest    OffsetResetStrategy = "LATEST"
	OffsetResetStrategyTimestamp OffsetResetStrategy = "TIMESTAMP"
	OffsetResetStrategyAbsolute  OffsetResetStrategy = "ABSOLUTE"
)

var allowedOffsetResetStrategyEnumValues = []OffsetResetStrategy{
	OffsetResetStrategyEarliest,
	OffsetResetStrategyLatest,
	OffsetResetStrategyTimestamp,
	OffsetResetStrategyAbsolute,
}

// GetAllowedValues returns the list of possible values.
func (v *OffsetResetStrategy) GetAllowedValues() []OffsetResetStrategy {
	return allowedOffsetResetStrategyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OffsetResetStrategy) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OffsetResetStrategy(value)
	return nil
}

// NewOffsetResetStrategyFromValue returns a pointer to a valid OffsetResetStrategy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOffsetResetStrategyFromValue(v string) (*OffsetResetStrategy, error) {
	ev := OffsetResetStrategy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OffsetResetStrategy: valid values are %v", v, allowedOffsetResetStrategyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OffsetResetStrategy) IsValid() bool {
	for _, existing := range allowedOffsetResetStrategyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OffsetResetStrategy value.
func (v OffsetResetStrategy) Ptr() *OffsetResetStrategy {
	return &v
}
