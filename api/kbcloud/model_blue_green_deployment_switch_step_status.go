// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// BlueGreenDeploymentSwitchStepStatus The status of a step in a blue-green deployment switch.
type BlueGreenDeploymentSwitchStepStatus string

// List of BlueGreenDeploymentSwitchStepStatus.
const (
	BlueGreenDeploymentSwitchStepStatusPending BlueGreenDeploymentSwitchStepStatus = "Pending"
	BlueGreenDeploymentSwitchStepStatusRunning BlueGreenDeploymentSwitchStepStatus = "Running"
	BlueGreenDeploymentSwitchStepStatusSucceed BlueGreenDeploymentSwitchStepStatus = "Succeed"
	BlueGreenDeploymentSwitchStepStatusFailed  BlueGreenDeploymentSwitchStepStatus = "Failed"
)

var allowedBlueGreenDeploymentSwitchStepStatusEnumValues = []BlueGreenDeploymentSwitchStepStatus{
	BlueGreenDeploymentSwitchStepStatusPending,
	BlueGreenDeploymentSwitchStepStatusRunning,
	BlueGreenDeploymentSwitchStepStatusSucceed,
	BlueGreenDeploymentSwitchStepStatusFailed,
}

// GetAllowedValues returns the list of possible values.
func (v *BlueGreenDeploymentSwitchStepStatus) GetAllowedValues() []BlueGreenDeploymentSwitchStepStatus {
	return allowedBlueGreenDeploymentSwitchStepStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *BlueGreenDeploymentSwitchStepStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = BlueGreenDeploymentSwitchStepStatus(value)
	return nil
}

// NewBlueGreenDeploymentSwitchStepStatusFromValue returns a pointer to a valid BlueGreenDeploymentSwitchStepStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewBlueGreenDeploymentSwitchStepStatusFromValue(v string) (*BlueGreenDeploymentSwitchStepStatus, error) {
	ev := BlueGreenDeploymentSwitchStepStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for BlueGreenDeploymentSwitchStepStatus: valid values are %v", v, allowedBlueGreenDeploymentSwitchStepStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v BlueGreenDeploymentSwitchStepStatus) IsValid() bool {
	for _, existing := range allowedBlueGreenDeploymentSwitchStepStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BlueGreenDeploymentSwitchStepStatus value.
func (v BlueGreenDeploymentSwitchStepStatus) Ptr() *BlueGreenDeploymentSwitchStepStatus {
	return &v
}
