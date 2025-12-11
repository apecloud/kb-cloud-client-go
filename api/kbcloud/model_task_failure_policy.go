// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// TaskFailurePolicy Policy to handle failures
type TaskFailurePolicy string

// List of TaskFailurePolicy.
const (
	TaskFailurePolicyInterrupt TaskFailurePolicy = "interrupt"
	TaskFailurePolicyIgnore    TaskFailurePolicy = "ignore"
)

var allowedTaskFailurePolicyEnumValues = []TaskFailurePolicy{
	TaskFailurePolicyInterrupt,
	TaskFailurePolicyIgnore,
}

// GetAllowedValues returns the list of possible values.
func (v *TaskFailurePolicy) GetAllowedValues() []TaskFailurePolicy {
	return allowedTaskFailurePolicyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TaskFailurePolicy) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TaskFailurePolicy(value)
	return nil
}

// NewTaskFailurePolicyFromValue returns a pointer to a valid TaskFailurePolicy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTaskFailurePolicyFromValue(v string) (*TaskFailurePolicy, error) {
	ev := TaskFailurePolicy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TaskFailurePolicy: valid values are %v", v, allowedTaskFailurePolicyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TaskFailurePolicy) IsValid() bool {
	for _, existing := range allowedTaskFailurePolicyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TaskFailurePolicy value.
func (v TaskFailurePolicy) Ptr() *TaskFailurePolicy {
	return &v
}
