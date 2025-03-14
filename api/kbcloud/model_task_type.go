// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// TaskType Type of task operation
type TaskType string

// List of TaskType.
const (
	TaskTypeScaleOut                       TaskType = "ScaleOut"
	TaskTypeScaleIn                        TaskType = "ScaleIn"
	TaskTypeDeleteDisasterRecoveryCluster  TaskType = "DeleteDisasterRecoveryCluster"
	TaskTypeCreateDisasterRecoveryCluster  TaskType = "CreateDisasterRecoveryCluster"
	TaskTypePromoteDisasterRecoveryCluster TaskType = "PromoteDisasterRecoveryCluster"
)

var allowedTaskTypeEnumValues = []TaskType{
	TaskTypeScaleOut,
	TaskTypeScaleIn,
	TaskTypeDeleteDisasterRecoveryCluster,
	TaskTypeCreateDisasterRecoveryCluster,
	TaskTypePromoteDisasterRecoveryCluster,
}

// GetAllowedValues returns the list of possible values.
func (v *TaskType) GetAllowedValues() []TaskType {
	return allowedTaskTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TaskType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TaskType(value)
	return nil
}

// NewTaskTypeFromValue returns a pointer to a valid TaskType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTaskTypeFromValue(v string) (*TaskType, error) {
	ev := TaskType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TaskType: valid values are %v", v, allowedTaskTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TaskType) IsValid() bool {
	for _, existing := range allowedTaskTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TaskType value.
func (v TaskType) Ptr() *TaskType {
	return &v
}
