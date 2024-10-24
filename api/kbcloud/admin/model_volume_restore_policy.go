// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// VolumeRestorePolicy the volume claim restore policy, support values: [Serial, Parallel]
type VolumeRestorePolicy string

// List of VolumeRestorePolicy.
const (
	VOLUMERESTOREPOLICY_SERIAL   VolumeRestorePolicy = "Serial"
	VOLUMERESTOREPOLICY_PARALLEL VolumeRestorePolicy = "Parallel"
)

var allowedVolumeRestorePolicyEnumValues = []VolumeRestorePolicy{
	VOLUMERESTOREPOLICY_SERIAL,
	VOLUMERESTOREPOLICY_PARALLEL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *VolumeRestorePolicy) GetAllowedValues() []VolumeRestorePolicy {
	return allowedVolumeRestorePolicyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *VolumeRestorePolicy) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = VolumeRestorePolicy(value)
	return nil
}

// NewVolumeRestorePolicyFromValue returns a pointer to a valid VolumeRestorePolicy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewVolumeRestorePolicyFromValue(v string) (*VolumeRestorePolicy, error) {
	ev := VolumeRestorePolicy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for VolumeRestorePolicy: valid values are %v", v, allowedVolumeRestorePolicyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v VolumeRestorePolicy) IsValid() bool {
	for _, existing := range allowedVolumeRestorePolicyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VolumeRestorePolicy value.
func (v VolumeRestorePolicy) Ptr() *VolumeRestorePolicy {
	return &v
}
