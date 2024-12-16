// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
)

// LoadBalancerAvailableType Whether the loadbalancer is available in the environment.
type LoadBalancerAvailableType string

// List of LoadBalancerAvailableType.
const (
	LoadBalancerAvailableTypeAvailable   LoadBalancerAvailableType = "Available"
	LoadBalancerAvailableTypeUnavailable LoadBalancerAvailableType = "Unavailable"
	LoadBalancerAvailableTypeChecking    LoadBalancerAvailableType = "Checking"
	LoadBalancerAvailableTypeUnknown     LoadBalancerAvailableType = "Unknown"
)

var allowedLoadBalancerAvailableTypeEnumValues = []LoadBalancerAvailableType{
	LoadBalancerAvailableTypeAvailable,
	LoadBalancerAvailableTypeUnavailable,
	LoadBalancerAvailableTypeChecking,
	LoadBalancerAvailableTypeUnknown,
}

// GetAllowedValues returns the list of possible values.
func (v *LoadBalancerAvailableType) GetAllowedValues() []LoadBalancerAvailableType {
	return allowedLoadBalancerAvailableTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LoadBalancerAvailableType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LoadBalancerAvailableType(value)
	return nil
}

// NewLoadBalancerAvailableTypeFromValue returns a pointer to a valid LoadBalancerAvailableType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLoadBalancerAvailableTypeFromValue(v string) (*LoadBalancerAvailableType, error) {
	ev := LoadBalancerAvailableType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LoadBalancerAvailableType: valid values are %v", v, allowedLoadBalancerAvailableTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LoadBalancerAvailableType) IsValid() bool {
	for _, existing := range allowedLoadBalancerAvailableTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LoadBalancerAvailableType value.
func (v LoadBalancerAvailableType) Ptr() *LoadBalancerAvailableType {
	return &v
}
