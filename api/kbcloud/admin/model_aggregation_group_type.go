// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// AggregationGroupType The type of aggregation group
type AggregationGroupType string

// List of AggregationGroupType.
const (
	AggregationGroupTypeGlobal       AggregationGroupType = "global"
	AggregationGroupTypeOrganization AggregationGroupType = "organization"
	AggregationGroupTypeProject      AggregationGroupType = "project"
	AggregationGroupTypeCluster      AggregationGroupType = "cluster"
)

var allowedAggregationGroupTypeEnumValues = []AggregationGroupType{
	AggregationGroupTypeGlobal,
	AggregationGroupTypeOrganization,
	AggregationGroupTypeProject,
	AggregationGroupTypeCluster,
}

// GetAllowedValues returns the list of possible values.
func (v *AggregationGroupType) GetAllowedValues() []AggregationGroupType {
	return allowedAggregationGroupTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AggregationGroupType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AggregationGroupType(value)
	return nil
}

// NewAggregationGroupTypeFromValue returns a pointer to a valid AggregationGroupType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAggregationGroupTypeFromValue(v string) (*AggregationGroupType, error) {
	ev := AggregationGroupType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AggregationGroupType: valid values are %v", v, allowedAggregationGroupTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AggregationGroupType) IsValid() bool {
	for _, existing := range allowedAggregationGroupTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AggregationGroupType value.
func (v AggregationGroupType) Ptr() *AggregationGroupType {
	return &v
}

// NullableAggregationGroupType handles when a null is used for AggregationGroupType.
type NullableAggregationGroupType struct {
	value *AggregationGroupType
	isSet bool
}

// Get returns the associated value.
func (v NullableAggregationGroupType) Get() *AggregationGroupType {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableAggregationGroupType) Set(val *AggregationGroupType) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableAggregationGroupType) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableAggregationGroupType) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableAggregationGroupType initializes the struct as if Set has been called.
func NewNullableAggregationGroupType(val *AggregationGroupType) *NullableAggregationGroupType {
	return &NullableAggregationGroupType{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableAggregationGroupType) MarshalJSON() ([]byte, error) {
	return common.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableAggregationGroupType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return common.Unmarshal(src, &v.value)
}
