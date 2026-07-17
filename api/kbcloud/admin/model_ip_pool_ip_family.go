// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type IpPoolIPFamily string

// List of IpPoolIPFamily.
const (
	IpPoolIPFamilyIPv4 IpPoolIPFamily = "IPv4"
	IpPoolIPFamilyIPv6 IpPoolIPFamily = "IPv6"
)

var allowedIpPoolIPFamilyEnumValues = []IpPoolIPFamily{
	IpPoolIPFamilyIPv4,
	IpPoolIPFamilyIPv6,
}

// GetAllowedValues returns the list of possible values.
func (v *IpPoolIPFamily) GetAllowedValues() []IpPoolIPFamily {
	return allowedIpPoolIPFamilyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IpPoolIPFamily) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IpPoolIPFamily(value)
	return nil
}

// NewIpPoolIPFamilyFromValue returns a pointer to a valid IpPoolIPFamily
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIpPoolIPFamilyFromValue(v string) (*IpPoolIPFamily, error) {
	ev := IpPoolIPFamily(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IpPoolIPFamily: valid values are %v", v, allowedIpPoolIPFamilyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IpPoolIPFamily) IsValid() bool {
	for _, existing := range allowedIpPoolIPFamilyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IpPoolIPFamily value.
func (v IpPoolIPFamily) Ptr() *IpPoolIPFamily {
	return &v
}
