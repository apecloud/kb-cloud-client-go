// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// TaskRetryMode retry reruns all steps; resume preserves successful steps.
type TaskRetryMode string

// List of TaskRetryMode.
const (
	TaskRetryModeRetry  TaskRetryMode = "retry"
	TaskRetryModeResume TaskRetryMode = "resume"
)

var allowedTaskRetryModeEnumValues = []TaskRetryMode{
	TaskRetryModeRetry,
	TaskRetryModeResume,
}

// GetAllowedValues returns the list of possible values.
func (v *TaskRetryMode) GetAllowedValues() []TaskRetryMode {
	return allowedTaskRetryModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TaskRetryMode) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TaskRetryMode(value)
	return nil
}

// NewTaskRetryModeFromValue returns a pointer to a valid TaskRetryMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTaskRetryModeFromValue(v string) (*TaskRetryMode, error) {
	ev := TaskRetryMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TaskRetryMode: valid values are %v", v, allowedTaskRetryModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TaskRetryMode) IsValid() bool {
	for _, existing := range allowedTaskRetryModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TaskRetryMode value.
func (v TaskRetryMode) Ptr() *TaskRetryMode {
	return &v
}
