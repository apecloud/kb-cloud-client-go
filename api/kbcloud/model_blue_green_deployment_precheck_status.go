// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// BlueGreenDeploymentPrecheckStatus The status of a step in a blue-green deployment precheck.
type BlueGreenDeploymentPrecheckStatus string

// List of BlueGreenDeploymentPrecheckStatus.
const (
	BlueGreenDeploymentPrecheckStatusPending BlueGreenDeploymentPrecheckStatus = "Pending"
	BlueGreenDeploymentPrecheckStatusRunning BlueGreenDeploymentPrecheckStatus = "Running"
	BlueGreenDeploymentPrecheckStatusSucceed BlueGreenDeploymentPrecheckStatus = "Succeed"
	BlueGreenDeploymentPrecheckStatusWarning BlueGreenDeploymentPrecheckStatus = "Warning"
	BlueGreenDeploymentPrecheckStatusFailed  BlueGreenDeploymentPrecheckStatus = "Failed"
)

var allowedBlueGreenDeploymentPrecheckStatusEnumValues = []BlueGreenDeploymentPrecheckStatus{
	BlueGreenDeploymentPrecheckStatusPending,
	BlueGreenDeploymentPrecheckStatusRunning,
	BlueGreenDeploymentPrecheckStatusSucceed,
	BlueGreenDeploymentPrecheckStatusWarning,
	BlueGreenDeploymentPrecheckStatusFailed,
}

// GetAllowedValues returns the list of possible values.
func (v *BlueGreenDeploymentPrecheckStatus) GetAllowedValues() []BlueGreenDeploymentPrecheckStatus {
	return allowedBlueGreenDeploymentPrecheckStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *BlueGreenDeploymentPrecheckStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = BlueGreenDeploymentPrecheckStatus(value)
	return nil
}

// NewBlueGreenDeploymentPrecheckStatusFromValue returns a pointer to a valid BlueGreenDeploymentPrecheckStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewBlueGreenDeploymentPrecheckStatusFromValue(v string) (*BlueGreenDeploymentPrecheckStatus, error) {
	ev := BlueGreenDeploymentPrecheckStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for BlueGreenDeploymentPrecheckStatus: valid values are %v", v, allowedBlueGreenDeploymentPrecheckStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v BlueGreenDeploymentPrecheckStatus) IsValid() bool {
	for _, existing := range allowedBlueGreenDeploymentPrecheckStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BlueGreenDeploymentPrecheckStatus value.
func (v BlueGreenDeploymentPrecheckStatus) Ptr() *BlueGreenDeploymentPrecheckStatus {
	return &v
}
