// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package admin

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)



// LoadBalancerStatus Status of the load balancer
type LoadBalancerStatus string

// List of LoadBalancerStatus.
const (
	LoadBalancerStatusEnabling LoadBalancerStatus = "Enabling"
	LoadBalancerStatusEnabled LoadBalancerStatus = "Enabled"
	LoadBalancerStatusDisabling LoadBalancerStatus = "Disabling"
	LoadBalancerStatusDisabled LoadBalancerStatus = "Disabled"
	LoadBalancerStatusFailed LoadBalancerStatus = "Failed"
)

var allowedLoadBalancerStatusEnumValues = []LoadBalancerStatus{
	LoadBalancerStatusEnabling,
	LoadBalancerStatusEnabled,
	LoadBalancerStatusDisabling,
	LoadBalancerStatusDisabled,
	LoadBalancerStatusFailed,
}

// GetAllowedValues returns the list of possible values.
func (v *LoadBalancerStatus) GetAllowedValues() []LoadBalancerStatus {
	return allowedLoadBalancerStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LoadBalancerStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LoadBalancerStatus(value)
	return nil
}

// NewLoadBalancerStatusFromValue returns a pointer to a valid LoadBalancerStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLoadBalancerStatusFromValue(v string) (*LoadBalancerStatus, error) {
	ev := LoadBalancerStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LoadBalancerStatus: valid values are %v", v, allowedLoadBalancerStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LoadBalancerStatus) IsValid() bool {
	for _, existing := range allowedLoadBalancerStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LoadBalancerStatus value.
func (v LoadBalancerStatus) Ptr() *LoadBalancerStatus {
	return &v
}
