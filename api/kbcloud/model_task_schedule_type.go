// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// TaskScheduleType Schedule type for task execution
type TaskScheduleType string

// List of TaskScheduleType.
const (
	TaskScheduleTypeImmediate         TaskScheduleType = "immediate"
	TaskScheduleTypeScheduled         TaskScheduleType = "scheduled"
	TaskScheduleTypeMaintenanceWindow TaskScheduleType = "maintenanceWindow"
)

var allowedTaskScheduleTypeEnumValues = []TaskScheduleType{
	TaskScheduleTypeImmediate,
	TaskScheduleTypeScheduled,
	TaskScheduleTypeMaintenanceWindow,
}

// GetAllowedValues returns the list of possible values.
func (v *TaskScheduleType) GetAllowedValues() []TaskScheduleType {
	return allowedTaskScheduleTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TaskScheduleType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TaskScheduleType(value)
	return nil
}

// NewTaskScheduleTypeFromValue returns a pointer to a valid TaskScheduleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTaskScheduleTypeFromValue(v string) (*TaskScheduleType, error) {
	ev := TaskScheduleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TaskScheduleType: valid values are %v", v, allowedTaskScheduleTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TaskScheduleType) IsValid() bool {
	for _, existing := range allowedTaskScheduleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TaskScheduleType value.
func (v TaskScheduleType) Ptr() *TaskScheduleType {
	return &v
}
