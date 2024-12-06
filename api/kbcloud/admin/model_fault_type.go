// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
)

type FaultType string

// List of FaultType.
const (
	FaultTypePodChaos     FaultType = "PodChaos"
	FaultTypeNetworkChaos FaultType = "NetworkChaos"
	FaultTypeIoChaos      FaultType = "IoChaos"
	FaultTypeTimeChaos    FaultType = "TimeChaos"
	FaultTypeStressChaos  FaultType = "StressChaos"
)

var allowedFaultTypeEnumValues = []FaultType{
	FaultTypePodChaos,
	FaultTypeNetworkChaos,
	FaultTypeIoChaos,
	FaultTypeTimeChaos,
	FaultTypeStressChaos,
}

// GetAllowedValues returns the list of possible values.
func (v *FaultType) GetAllowedValues() []FaultType {
	return allowedFaultTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FaultType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FaultType(value)
	return nil
}

// NewFaultTypeFromValue returns a pointer to a valid FaultType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFaultTypeFromValue(v string) (*FaultType, error) {
	ev := FaultType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FaultType: valid values are %v", v, allowedFaultTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FaultType) IsValid() bool {
	for _, existing := range allowedFaultTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FaultType value.
func (v FaultType) Ptr() *FaultType {
	return &v
}
