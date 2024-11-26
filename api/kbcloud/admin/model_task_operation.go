// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
)

// TaskOperation Type of scaling operation
type TaskOperation string

// List of TaskOperation.
const (
	TaskOperationScaleOut TaskOperation = "ScaleOut"
	TaskOperationScaleIn  TaskOperation = "ScaleIn"
)

var allowedTaskOperationEnumValues = []TaskOperation{
	TaskOperationScaleOut,
	TaskOperationScaleIn,
}

// GetAllowedValues returns the list of possible values.
func (v *TaskOperation) GetAllowedValues() []TaskOperation {
	return allowedTaskOperationEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TaskOperation) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TaskOperation(value)
	return nil
}

// NewTaskOperationFromValue returns a pointer to a valid TaskOperation
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTaskOperationFromValue(v string) (*TaskOperation, error) {
	ev := TaskOperation(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TaskOperation: valid values are %v", v, allowedTaskOperationEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TaskOperation) IsValid() bool {
	for _, existing := range allowedTaskOperationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TaskOperation value.
func (v TaskOperation) Ptr() *TaskOperation {
	return &v
}
