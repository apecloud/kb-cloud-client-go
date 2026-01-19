// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ClusterType Describes the type of cluster
type ClusterType string

// List of ClusterType.
const (
	ClusterTypeNormal              ClusterType = "Normal"
	ClusterTypeDisasterRecovery    ClusterType = "DisasterRecovery"
	ClusterTypeBlueGreenDeployment ClusterType = "BlueGreenDeployment"
)

var allowedClusterTypeEnumValues = []ClusterType{
	ClusterTypeNormal,
	ClusterTypeDisasterRecovery,
	ClusterTypeBlueGreenDeployment,
}

// GetAllowedValues returns the list of possible values.
func (v *ClusterType) GetAllowedValues() []ClusterType {
	return allowedClusterTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ClusterType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ClusterType(value)
	return nil
}

// NewClusterTypeFromValue returns a pointer to a valid ClusterType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewClusterTypeFromValue(v string) (*ClusterType, error) {
	ev := ClusterType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ClusterType: valid values are %v", v, allowedClusterTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ClusterType) IsValid() bool {
	for _, existing := range allowedClusterTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ClusterType value.
func (v ClusterType) Ptr() *ClusterType {
	return &v
}

// NullableClusterType handles when a null is used for ClusterType.
type NullableClusterType struct {
	value *ClusterType
	isSet bool
}

// Get returns the associated value.
func (v NullableClusterType) Get() *ClusterType {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableClusterType) Set(val *ClusterType) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableClusterType) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableClusterType) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableClusterType initializes the struct as if Set has been called.
func NewNullableClusterType(val *ClusterType) *NullableClusterType {
	return &NullableClusterType{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableClusterType) MarshalJSON() ([]byte, error) {
	return common.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableClusterType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return common.Unmarshal(src, &v.value)
}
