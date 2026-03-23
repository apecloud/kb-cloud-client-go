// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// OrgAggregationGroupType The type of aggregation group for organization-scoped bills
type OrgAggregationGroupType string

// List of OrgAggregationGroupType.
const (
	OrgAggregationGroupTypeOrganization OrgAggregationGroupType = "organization"
	OrgAggregationGroupTypeCluster      OrgAggregationGroupType = "cluster"
)

var allowedOrgAggregationGroupTypeEnumValues = []OrgAggregationGroupType{
	OrgAggregationGroupTypeOrganization,
	OrgAggregationGroupTypeCluster,
}

// GetAllowedValues returns the list of possible values.
func (v *OrgAggregationGroupType) GetAllowedValues() []OrgAggregationGroupType {
	return allowedOrgAggregationGroupTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OrgAggregationGroupType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OrgAggregationGroupType(value)
	return nil
}

// NewOrgAggregationGroupTypeFromValue returns a pointer to a valid OrgAggregationGroupType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOrgAggregationGroupTypeFromValue(v string) (*OrgAggregationGroupType, error) {
	ev := OrgAggregationGroupType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OrgAggregationGroupType: valid values are %v", v, allowedOrgAggregationGroupTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OrgAggregationGroupType) IsValid() bool {
	for _, existing := range allowedOrgAggregationGroupTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrgAggregationGroupType value.
func (v OrgAggregationGroupType) Ptr() *OrgAggregationGroupType {
	return &v
}

// NullableOrgAggregationGroupType handles when a null is used for OrgAggregationGroupType.
type NullableOrgAggregationGroupType struct {
	value *OrgAggregationGroupType
	isSet bool
}

// Get returns the associated value.
func (v NullableOrgAggregationGroupType) Get() *OrgAggregationGroupType {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableOrgAggregationGroupType) Set(val *OrgAggregationGroupType) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableOrgAggregationGroupType) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableOrgAggregationGroupType) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableOrgAggregationGroupType initializes the struct as if Set has been called.
func NewNullableOrgAggregationGroupType(val *OrgAggregationGroupType) *NullableOrgAggregationGroupType {
	return &NullableOrgAggregationGroupType{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableOrgAggregationGroupType) MarshalJSON() ([]byte, error) {
	return common.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableOrgAggregationGroupType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return common.Unmarshal(src, &v.value)
}
