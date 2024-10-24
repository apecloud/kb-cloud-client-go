// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// LoadBalancerType Type of the load balancer
type LoadBalancerType string

// List of LoadBalancerType.
const (
	LOADBALANCERTYPE_F5      LoadBalancerType = "f5"
	LOADBALANCERTYPE_METALLB LoadBalancerType = "metallb"
)

var allowedLoadBalancerTypeEnumValues = []LoadBalancerType{
	LOADBALANCERTYPE_F5,
	LOADBALANCERTYPE_METALLB,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LoadBalancerType) GetAllowedValues() []LoadBalancerType {
	return allowedLoadBalancerTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LoadBalancerType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LoadBalancerType(value)
	return nil
}

// NewLoadBalancerTypeFromValue returns a pointer to a valid LoadBalancerType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLoadBalancerTypeFromValue(v string) (*LoadBalancerType, error) {
	ev := LoadBalancerType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LoadBalancerType: valid values are %v", v, allowedLoadBalancerTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LoadBalancerType) IsValid() bool {
	for _, existing := range allowedLoadBalancerTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LoadBalancerType value.
func (v LoadBalancerType) Ptr() *LoadBalancerType {
	return &v
}
