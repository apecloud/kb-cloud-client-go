// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type EndpointStatusEnum string

// List of EndpointStatusEnum.
const (
	EndpointStatusEnumValid   EndpointStatusEnum = "valid"
	EndpointStatusEnumInvalid EndpointStatusEnum = "invalid"
)

var allowedEndpointStatusEnumEnumValues = []EndpointStatusEnum{
	EndpointStatusEnumValid,
	EndpointStatusEnumInvalid,
}

// GetAllowedValues returns the list of possible values.
func (v *EndpointStatusEnum) GetAllowedValues() []EndpointStatusEnum {
	return allowedEndpointStatusEnumEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EndpointStatusEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EndpointStatusEnum(value)
	return nil
}

// NewEndpointStatusEnumFromValue returns a pointer to a valid EndpointStatusEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEndpointStatusEnumFromValue(v string) (*EndpointStatusEnum, error) {
	ev := EndpointStatusEnum(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EndpointStatusEnum: valid values are %v", v, allowedEndpointStatusEnumEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EndpointStatusEnum) IsValid() bool {
	for _, existing := range allowedEndpointStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EndpointStatusEnum value.
func (v EndpointStatusEnum) Ptr() *EndpointStatusEnum {
	return &v
}
