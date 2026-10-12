// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ModeAvailabilityZonePlacement AZ selection granularity. Omitted means cluster placement. deploymentUnit requires an explicit AZ for every deployed component or instance template at the granularity declared by its mode component.
type ModeAvailabilityZonePlacement string

// List of ModeAvailabilityZonePlacement.
const (
	ModeAvailabilityZonePlacementCluster        ModeAvailabilityZonePlacement = "cluster"
	ModeAvailabilityZonePlacementDeploymentUnit ModeAvailabilityZonePlacement = "deploymentUnit"
)

var allowedModeAvailabilityZonePlacementEnumValues = []ModeAvailabilityZonePlacement{
	ModeAvailabilityZonePlacementCluster,
	ModeAvailabilityZonePlacementDeploymentUnit,
}

// GetAllowedValues returns the list of possible values.
func (v *ModeAvailabilityZonePlacement) GetAllowedValues() []ModeAvailabilityZonePlacement {
	return allowedModeAvailabilityZonePlacementEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ModeAvailabilityZonePlacement) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ModeAvailabilityZonePlacement(value)
	return nil
}

// NewModeAvailabilityZonePlacementFromValue returns a pointer to a valid ModeAvailabilityZonePlacement
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewModeAvailabilityZonePlacementFromValue(v string) (*ModeAvailabilityZonePlacement, error) {
	ev := ModeAvailabilityZonePlacement(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ModeAvailabilityZonePlacement: valid values are %v", v, allowedModeAvailabilityZonePlacementEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ModeAvailabilityZonePlacement) IsValid() bool {
	for _, existing := range allowedModeAvailabilityZonePlacementEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ModeAvailabilityZonePlacement value.
func (v ModeAvailabilityZonePlacement) Ptr() *ModeAvailabilityZonePlacement {
	return &v
}
