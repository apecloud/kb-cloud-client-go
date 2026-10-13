// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ComponentAvailabilityZonePlacement Optional AZ selection target in a component mode. component selects one AZ for the component; instanceTemplate selects one AZ for each declared template. This declares the UI target; it does not require an AZ in create requests.
type ComponentAvailabilityZonePlacement string

// List of ComponentAvailabilityZonePlacement.
const (
	ComponentAvailabilityZonePlacementComponent        ComponentAvailabilityZonePlacement = "component"
	ComponentAvailabilityZonePlacementInstanceTemplate ComponentAvailabilityZonePlacement = "instanceTemplate"
)

var allowedComponentAvailabilityZonePlacementEnumValues = []ComponentAvailabilityZonePlacement{
	ComponentAvailabilityZonePlacementComponent,
	ComponentAvailabilityZonePlacementInstanceTemplate,
}

// GetAllowedValues returns the list of possible values.
func (v *ComponentAvailabilityZonePlacement) GetAllowedValues() []ComponentAvailabilityZonePlacement {
	return allowedComponentAvailabilityZonePlacementEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ComponentAvailabilityZonePlacement) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ComponentAvailabilityZonePlacement(value)
	return nil
}

// NewComponentAvailabilityZonePlacementFromValue returns a pointer to a valid ComponentAvailabilityZonePlacement
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewComponentAvailabilityZonePlacementFromValue(v string) (*ComponentAvailabilityZonePlacement, error) {
	ev := ComponentAvailabilityZonePlacement(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ComponentAvailabilityZonePlacement: valid values are %v", v, allowedComponentAvailabilityZonePlacementEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ComponentAvailabilityZonePlacement) IsValid() bool {
	for _, existing := range allowedComponentAvailabilityZonePlacementEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ComponentAvailabilityZonePlacement value.
func (v ComponentAvailabilityZonePlacement) Ptr() *ComponentAvailabilityZonePlacement {
	return &v
}
