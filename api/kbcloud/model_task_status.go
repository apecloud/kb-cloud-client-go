// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// TaskStatus Current status of the task
type TaskStatus string

// List of TaskStatus.
const (
	TaskStatusPending TaskStatus = "Pending"
	TaskStatusRunning TaskStatus = "Running"
	TaskStatusStopped TaskStatus = "Stopped"
	TaskStatusSucceed TaskStatus = "Succeed"
	TaskStatusFailed  TaskStatus = "Failed"
)

var allowedTaskStatusEnumValues = []TaskStatus{
	TaskStatusPending,
	TaskStatusRunning,
	TaskStatusStopped,
	TaskStatusSucceed,
	TaskStatusFailed,
}

// GetAllowedValues returns the list of possible values.
func (v *TaskStatus) GetAllowedValues() []TaskStatus {
	return allowedTaskStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TaskStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TaskStatus(value)
	return nil
}

// NewTaskStatusFromValue returns a pointer to a valid TaskStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTaskStatusFromValue(v string) (*TaskStatus, error) {
	ev := TaskStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TaskStatus: valid values are %v", v, allowedTaskStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TaskStatus) IsValid() bool {
	for _, existing := range allowedTaskStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TaskStatus value.
func (v TaskStatus) Ptr() *TaskStatus {
	return &v
}
