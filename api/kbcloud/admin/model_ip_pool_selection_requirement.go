// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type IpPoolSelectionRequirement string

// List of IpPoolSelectionRequirement.
const (
	IpPoolSelectionRequirementRequired    IpPoolSelectionRequirement = "required"
	IpPoolSelectionRequirementOptional    IpPoolSelectionRequirement = "optional"
	IpPoolSelectionRequirementUnsupported IpPoolSelectionRequirement = "unsupported"
)

var allowedIpPoolSelectionRequirementEnumValues = []IpPoolSelectionRequirement{
	IpPoolSelectionRequirementRequired,
	IpPoolSelectionRequirementOptional,
	IpPoolSelectionRequirementUnsupported,
}

// GetAllowedValues returns the list of possible values.
func (v *IpPoolSelectionRequirement) GetAllowedValues() []IpPoolSelectionRequirement {
	return allowedIpPoolSelectionRequirementEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IpPoolSelectionRequirement) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IpPoolSelectionRequirement(value)
	return nil
}

// NewIpPoolSelectionRequirementFromValue returns a pointer to a valid IpPoolSelectionRequirement
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIpPoolSelectionRequirementFromValue(v string) (*IpPoolSelectionRequirement, error) {
	ev := IpPoolSelectionRequirement(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IpPoolSelectionRequirement: valid values are %v", v, allowedIpPoolSelectionRequirementEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IpPoolSelectionRequirement) IsValid() bool {
	for _, existing := range allowedIpPoolSelectionRequirementEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IpPoolSelectionRequirement value.
func (v IpPoolSelectionRequirement) Ptr() *IpPoolSelectionRequirement {
	return &v
}
