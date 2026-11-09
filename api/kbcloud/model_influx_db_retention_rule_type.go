// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type InfluxDBRetentionRuleType string

// List of InfluxDBRetentionRuleType.
const (
	InfluxDBRetentionRuleTypeExpire InfluxDBRetentionRuleType = "expire"
)

var allowedInfluxDBRetentionRuleTypeEnumValues = []InfluxDBRetentionRuleType{
	InfluxDBRetentionRuleTypeExpire,
}

// GetAllowedValues returns the list of possible values.
func (v *InfluxDBRetentionRuleType) GetAllowedValues() []InfluxDBRetentionRuleType {
	return allowedInfluxDBRetentionRuleTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *InfluxDBRetentionRuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = InfluxDBRetentionRuleType(value)
	return nil
}

// NewInfluxDBRetentionRuleTypeFromValue returns a pointer to a valid InfluxDBRetentionRuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewInfluxDBRetentionRuleTypeFromValue(v string) (*InfluxDBRetentionRuleType, error) {
	ev := InfluxDBRetentionRuleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for InfluxDBRetentionRuleType: valid values are %v", v, allowedInfluxDBRetentionRuleTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v InfluxDBRetentionRuleType) IsValid() bool {
	for _, existing := range allowedInfluxDBRetentionRuleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InfluxDBRetentionRuleType value.
func (v InfluxDBRetentionRuleType) Ptr() *InfluxDBRetentionRuleType {
	return &v
}
