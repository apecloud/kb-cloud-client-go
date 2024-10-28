// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EngineOptionsMetricsQueryType Specifies the type of metrics query to be performed.
// 'instant' for a single point in time, 'range' for a time range.
type EngineOptionsMetricsQueryType string

// List of EngineOptionsMetricsQueryType.
const (
	ENGINEOPTIONSMETRICSQUERYTYPE_INSTANT EngineOptionsMetricsQueryType = "instant"
	ENGINEOPTIONSMETRICSQUERYTYPE_RANGE   EngineOptionsMetricsQueryType = "range"
)

var allowedEngineOptionsMetricsQueryTypeEnumValues = []EngineOptionsMetricsQueryType{
	ENGINEOPTIONSMETRICSQUERYTYPE_INSTANT,
	ENGINEOPTIONSMETRICSQUERYTYPE_RANGE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EngineOptionsMetricsQueryType) GetAllowedValues() []EngineOptionsMetricsQueryType {
	return allowedEngineOptionsMetricsQueryTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EngineOptionsMetricsQueryType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EngineOptionsMetricsQueryType(value)
	return nil
}

// NewEngineOptionsMetricsQueryTypeFromValue returns a pointer to a valid EngineOptionsMetricsQueryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEngineOptionsMetricsQueryTypeFromValue(v string) (*EngineOptionsMetricsQueryType, error) {
	ev := EngineOptionsMetricsQueryType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EngineOptionsMetricsQueryType: valid values are %v", v, allowedEngineOptionsMetricsQueryTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EngineOptionsMetricsQueryType) IsValid() bool {
	for _, existing := range allowedEngineOptionsMetricsQueryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EngineOptionsMetricsQueryType value.
func (v EngineOptionsMetricsQueryType) Ptr() *EngineOptionsMetricsQueryType {
	return &v
}
