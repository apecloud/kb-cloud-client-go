// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
)

// EndpointNetworkType Network type of endpoint
type EndpointNetworkType string

// List of EndpointNetworkType.
const (
	EndpointNetworkTypeInternet EndpointNetworkType = "Internet"
	EndpointNetworkTypeIntranet EndpointNetworkType = "Intranet"
)

var allowedEndpointNetworkTypeEnumValues = []EndpointNetworkType{
	EndpointNetworkTypeInternet,
	EndpointNetworkTypeIntranet,
}

// GetAllowedValues returns the list of possible values.
func (v *EndpointNetworkType) GetAllowedValues() []EndpointNetworkType {
	return allowedEndpointNetworkTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EndpointNetworkType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EndpointNetworkType(value)
	return nil
}

// NewEndpointNetworkTypeFromValue returns a pointer to a valid EndpointNetworkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEndpointNetworkTypeFromValue(v string) (*EndpointNetworkType, error) {
	ev := EndpointNetworkType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EndpointNetworkType: valid values are %v", v, allowedEndpointNetworkTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EndpointNetworkType) IsValid() bool {
	for _, existing := range allowedEndpointNetworkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EndpointNetworkType value.
func (v EndpointNetworkType) Ptr() *EndpointNetworkType {
	return &v
}
