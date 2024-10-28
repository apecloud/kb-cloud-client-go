// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// LoadBalancerAvailableType Whether the loadbalancer is available in the environment.
type LoadBalancerAvailableType string

// List of LoadBalancerAvailableType.
const (
	LOADBALANCERAVAILABLETYPE_AVAILABLE   LoadBalancerAvailableType = "Available"
	LOADBALANCERAVAILABLETYPE_UNAVAILABLE LoadBalancerAvailableType = "Unavailable"
	LOADBALANCERAVAILABLETYPE_CHECKING    LoadBalancerAvailableType = "Checking"
	LOADBALANCERAVAILABLETYPE_UNKNOWN     LoadBalancerAvailableType = "Unknown"
)

var allowedLoadBalancerAvailableTypeEnumValues = []LoadBalancerAvailableType{
	LOADBALANCERAVAILABLETYPE_AVAILABLE,
	LOADBALANCERAVAILABLETYPE_UNAVAILABLE,
	LOADBALANCERAVAILABLETYPE_CHECKING,
	LOADBALANCERAVAILABLETYPE_UNKNOWN,
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
