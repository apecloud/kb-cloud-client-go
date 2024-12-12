// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
)

type NetworkMode string

// List of NetworkMode.
const (
	NetworkModeHostNetwork NetworkMode = "HostNetwork"
	NetworkModeNodePort    NetworkMode = "NodePort"
	NetworkModeFixedPodIp  NetworkMode = "FixedPodIP"
)

var allowedNetworkModeEnumValues = []NetworkMode{
	NetworkModeHostNetwork,
	NetworkModeNodePort,
	NetworkModeFixedPodIp,
}

// GetAllowedValues returns the list of possible values.
func (v *NetworkMode) GetAllowedValues() []NetworkMode {
	return allowedNetworkModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *NetworkMode) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NetworkMode(value)
	return nil
}

// NewNetworkModeFromValue returns a pointer to a valid NetworkMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewNetworkModeFromValue(v string) (*NetworkMode, error) {
	ev := NetworkMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NetworkMode: valid values are %v", v, allowedNetworkModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v NetworkMode) IsValid() bool {
	for _, existing := range allowedNetworkModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NetworkMode value.
func (v NetworkMode) Ptr() *NetworkMode {
	return &v
}
