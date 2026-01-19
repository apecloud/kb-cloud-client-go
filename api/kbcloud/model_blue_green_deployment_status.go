// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// BlueGreenDeploymentStatus The status of a blue-green deployment.
type BlueGreenDeploymentStatus string

// List of BlueGreenDeploymentStatus.
const (
	BlueGreenDeploymentStatusMaintenance   BlueGreenDeploymentStatus = "Maintenance"
	BlueGreenDeploymentStatusSynchronizing BlueGreenDeploymentStatus = "Synchronizing"
	BlueGreenDeploymentStatusSynced        BlueGreenDeploymentStatus = "Synced"
	BlueGreenDeploymentStatusInterrupted   BlueGreenDeploymentStatus = "Interrupted"
	BlueGreenDeploymentStatusSwitched      BlueGreenDeploymentStatus = "Switched"
	BlueGreenDeploymentStatusFailed        BlueGreenDeploymentStatus = "Failed"
)

var allowedBlueGreenDeploymentStatusEnumValues = []BlueGreenDeploymentStatus{
	BlueGreenDeploymentStatusMaintenance,
	BlueGreenDeploymentStatusSynchronizing,
	BlueGreenDeploymentStatusSynced,
	BlueGreenDeploymentStatusInterrupted,
	BlueGreenDeploymentStatusSwitched,
	BlueGreenDeploymentStatusFailed,
}

// GetAllowedValues returns the list of possible values.
func (v *BlueGreenDeploymentStatus) GetAllowedValues() []BlueGreenDeploymentStatus {
	return allowedBlueGreenDeploymentStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *BlueGreenDeploymentStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = BlueGreenDeploymentStatus(value)
	return nil
}

// NewBlueGreenDeploymentStatusFromValue returns a pointer to a valid BlueGreenDeploymentStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewBlueGreenDeploymentStatusFromValue(v string) (*BlueGreenDeploymentStatus, error) {
	ev := BlueGreenDeploymentStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for BlueGreenDeploymentStatus: valid values are %v", v, allowedBlueGreenDeploymentStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v BlueGreenDeploymentStatus) IsValid() bool {
	for _, existing := range allowedBlueGreenDeploymentStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BlueGreenDeploymentStatus value.
func (v BlueGreenDeploymentStatus) Ptr() *BlueGreenDeploymentStatus {
	return &v
}
