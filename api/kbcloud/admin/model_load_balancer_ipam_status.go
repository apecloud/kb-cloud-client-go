// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// LoadBalancerIpamStatus Whether the ipam is enabled.
type LoadBalancerIpamStatus string

// List of LoadBalancerIpamStatus.
const (
	LOADBALANCERIPAMSTATUS_ENABLED  LoadBalancerIpamStatus = "Enabled"
	LOADBALANCERIPAMSTATUS_DISABLED LoadBalancerIpamStatus = "Disabled"
)

var allowedLoadBalancerIpamStatusEnumValues = []LoadBalancerIpamStatus{
	LOADBALANCERIPAMSTATUS_ENABLED,
	LOADBALANCERIPAMSTATUS_DISABLED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LoadBalancerIpamStatus) GetAllowedValues() []LoadBalancerIpamStatus {
	return allowedLoadBalancerIpamStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LoadBalancerIpamStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LoadBalancerIpamStatus(value)
	return nil
}

// NewLoadBalancerIpamStatusFromValue returns a pointer to a valid LoadBalancerIpamStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLoadBalancerIpamStatusFromValue(v string) (*LoadBalancerIpamStatus, error) {
	ev := LoadBalancerIpamStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LoadBalancerIpamStatus: valid values are %v", v, allowedLoadBalancerIpamStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LoadBalancerIpamStatus) IsValid() bool {
	for _, existing := range allowedLoadBalancerIpamStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LoadBalancerIpamStatus value.
func (v LoadBalancerIpamStatus) Ptr() *LoadBalancerIpamStatus {
	return &v
}
