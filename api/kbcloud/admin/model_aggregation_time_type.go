// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// AggregationTimeType The type of aggregation time
type AggregationTimeType string

// List of AggregationTimeType.
const (
	AggregationTimeTypeYearly    AggregationTimeType = "yearly"
	AggregationTimeTypeMonthly   AggregationTimeType = "monthly"
	AggregationTimeTypeQuarterly AggregationTimeType = "quarterly"
	AggregationTimeTypeDaily     AggregationTimeType = "daily"
)

var allowedAggregationTimeTypeEnumValues = []AggregationTimeType{
	AggregationTimeTypeYearly,
	AggregationTimeTypeMonthly,
	AggregationTimeTypeQuarterly,
	AggregationTimeTypeDaily,
}

// GetAllowedValues returns the list of possible values.
func (v *AggregationTimeType) GetAllowedValues() []AggregationTimeType {
	return allowedAggregationTimeTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AggregationTimeType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AggregationTimeType(value)
	return nil
}

// NewAggregationTimeTypeFromValue returns a pointer to a valid AggregationTimeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAggregationTimeTypeFromValue(v string) (*AggregationTimeType, error) {
	ev := AggregationTimeType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AggregationTimeType: valid values are %v", v, allowedAggregationTimeTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AggregationTimeType) IsValid() bool {
	for _, existing := range allowedAggregationTimeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AggregationTimeType value.
func (v AggregationTimeType) Ptr() *AggregationTimeType {
	return &v
}

// NullableAggregationTimeType handles when a null is used for AggregationTimeType.
type NullableAggregationTimeType struct {
	value *AggregationTimeType
	isSet bool
}

// Get returns the associated value.
func (v NullableAggregationTimeType) Get() *AggregationTimeType {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableAggregationTimeType) Set(val *AggregationTimeType) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableAggregationTimeType) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableAggregationTimeType) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableAggregationTimeType initializes the struct as if Set has been called.
func NewNullableAggregationTimeType(val *AggregationTimeType) *NullableAggregationTimeType {
	return &NullableAggregationTimeType{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableAggregationTimeType) MarshalJSON() ([]byte, error) {
	return common.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableAggregationTimeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return common.Unmarshal(src, &v.value)
}
