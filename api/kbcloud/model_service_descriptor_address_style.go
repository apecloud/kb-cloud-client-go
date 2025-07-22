// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ServiceDescriptorAddressStyle specify the style that will be used in servicedescriptor.
// "hostport" will ask user to provide both host and port.
// "endpoint" will ask user to provide an endpoint.
// When using a serviceSelector, this option will not be effective.
type ServiceDescriptorAddressStyle string

// List of ServiceDescriptorAddressStyle.
const (
	ServiceDescriptorAddressStyleHostport ServiceDescriptorAddressStyle = "hostport"
	ServiceDescriptorAddressStyleEndpoint ServiceDescriptorAddressStyle = "endpoint"
)

var allowedServiceDescriptorAddressStyleEnumValues = []ServiceDescriptorAddressStyle{
	ServiceDescriptorAddressStyleHostport,
	ServiceDescriptorAddressStyleEndpoint,
}

// GetAllowedValues returns the list of possible values.
func (v *ServiceDescriptorAddressStyle) GetAllowedValues() []ServiceDescriptorAddressStyle {
	return allowedServiceDescriptorAddressStyleEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ServiceDescriptorAddressStyle) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ServiceDescriptorAddressStyle(value)
	return nil
}

// NewServiceDescriptorAddressStyleFromValue returns a pointer to a valid ServiceDescriptorAddressStyle
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewServiceDescriptorAddressStyleFromValue(v string) (*ServiceDescriptorAddressStyle, error) {
	ev := ServiceDescriptorAddressStyle(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ServiceDescriptorAddressStyle: valid values are %v", v, allowedServiceDescriptorAddressStyleEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ServiceDescriptorAddressStyle) IsValid() bool {
	for _, existing := range allowedServiceDescriptorAddressStyleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceDescriptorAddressStyle value.
func (v ServiceDescriptorAddressStyle) Ptr() *ServiceDescriptorAddressStyle {
	return &v
}
