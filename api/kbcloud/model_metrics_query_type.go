// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type MetricsQueryType string

// List of MetricsQueryType.
const (
	MetricsQueryTypeInstant MetricsQueryType = "instant"
	MetricsQueryTypeRange   MetricsQueryType = "range"
)

var allowedMetricsQueryTypeEnumValues = []MetricsQueryType{
	MetricsQueryTypeInstant,
	MetricsQueryTypeRange,
}

// GetAllowedValues returns the list of possible values.
func (v *MetricsQueryType) GetAllowedValues() []MetricsQueryType {
	return allowedMetricsQueryTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MetricsQueryType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MetricsQueryType(value)
	return nil
}

// NewMetricsQueryTypeFromValue returns a pointer to a valid MetricsQueryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMetricsQueryTypeFromValue(v string) (*MetricsQueryType, error) {
	ev := MetricsQueryType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MetricsQueryType: valid values are %v", v, allowedMetricsQueryTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MetricsQueryType) IsValid() bool {
	for _, existing := range allowedMetricsQueryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MetricsQueryType value.
func (v MetricsQueryType) Ptr() *MetricsQueryType {
	return &v
}
