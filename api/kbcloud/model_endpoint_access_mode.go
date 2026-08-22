// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EndpointAccessMode Stable connection path identity for direct database and pooled application endpoints.
type EndpointAccessMode string

// List of EndpointAccessMode.
const (
	EndpointAccessModeDirect EndpointAccessMode = "direct"
	EndpointAccessModePooled EndpointAccessMode = "pooled"
)

var allowedEndpointAccessModeEnumValues = []EndpointAccessMode{
	EndpointAccessModeDirect,
	EndpointAccessModePooled,
}

// GetAllowedValues returns the list of possible values.
func (v *EndpointAccessMode) GetAllowedValues() []EndpointAccessMode {
	return allowedEndpointAccessModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EndpointAccessMode) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EndpointAccessMode(value)
	return nil
}

// NewEndpointAccessModeFromValue returns a pointer to a valid EndpointAccessMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEndpointAccessModeFromValue(v string) (*EndpointAccessMode, error) {
	ev := EndpointAccessMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EndpointAccessMode: valid values are %v", v, allowedEndpointAccessModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EndpointAccessMode) IsValid() bool {
	for _, existing := range allowedEndpointAccessModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EndpointAccessMode value.
func (v EndpointAccessMode) Ptr() *EndpointAccessMode {
	return &v
}
