// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
)

// EndpointType Type of endpoint
type EndpointType string

// List of EndpointType.
const (
	EndpointTypeClusterIp    EndpointType = "ClusterIP"
	EndpointTypeNodePort     EndpointType = "NodePort"
	EndpointTypeLoadBalancer EndpointType = "LoadBalancer"
	EndpointTypeFixedPodIp   EndpointType = "FixedPodIP"
	EndpointTypeHostNetwork  EndpointType = "HostNetwork"
)

var allowedEndpointTypeEnumValues = []EndpointType{
	EndpointTypeClusterIp,
	EndpointTypeNodePort,
	EndpointTypeLoadBalancer,
	EndpointTypeFixedPodIp,
	EndpointTypeHostNetwork,
}

// GetAllowedValues returns the list of possible values.
func (v *EndpointType) GetAllowedValues() []EndpointType {
	return allowedEndpointTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EndpointType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EndpointType(value)
	return nil
}

// NewEndpointTypeFromValue returns a pointer to a valid EndpointType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEndpointTypeFromValue(v string) (*EndpointType, error) {
	ev := EndpointType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EndpointType: valid values are %v", v, allowedEndpointTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EndpointType) IsValid() bool {
	for _, existing := range allowedEndpointTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EndpointType value.
func (v EndpointType) Ptr() *EndpointType {
	return &v
}
