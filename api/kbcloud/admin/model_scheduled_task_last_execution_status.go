// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ScheduledTaskLastExecutionStatus Most recently observable execution status
type ScheduledTaskLastExecutionStatus string

// List of ScheduledTaskLastExecutionStatus.
const (
	ScheduledTaskLastExecutionStatusSuccess   ScheduledTaskLastExecutionStatus = "success"
	ScheduledTaskLastExecutionStatusFailed    ScheduledTaskLastExecutionStatus = "failed"
	ScheduledTaskLastExecutionStatusRunning   ScheduledTaskLastExecutionStatus = "running"
	ScheduledTaskLastExecutionStatusPending   ScheduledTaskLastExecutionStatus = "pending"
	ScheduledTaskLastExecutionStatusCancelled ScheduledTaskLastExecutionStatus = "cancelled"
	ScheduledTaskLastExecutionStatusUnknown   ScheduledTaskLastExecutionStatus = "unknown"
)

var allowedScheduledTaskLastExecutionStatusEnumValues = []ScheduledTaskLastExecutionStatus{
	ScheduledTaskLastExecutionStatusSuccess,
	ScheduledTaskLastExecutionStatusFailed,
	ScheduledTaskLastExecutionStatusRunning,
	ScheduledTaskLastExecutionStatusPending,
	ScheduledTaskLastExecutionStatusCancelled,
	ScheduledTaskLastExecutionStatusUnknown,
}

// GetAllowedValues returns the list of possible values.
func (v *ScheduledTaskLastExecutionStatus) GetAllowedValues() []ScheduledTaskLastExecutionStatus {
	return allowedScheduledTaskLastExecutionStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScheduledTaskLastExecutionStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScheduledTaskLastExecutionStatus(value)
	return nil
}

// NewScheduledTaskLastExecutionStatusFromValue returns a pointer to a valid ScheduledTaskLastExecutionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScheduledTaskLastExecutionStatusFromValue(v string) (*ScheduledTaskLastExecutionStatus, error) {
	ev := ScheduledTaskLastExecutionStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScheduledTaskLastExecutionStatus: valid values are %v", v, allowedScheduledTaskLastExecutionStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScheduledTaskLastExecutionStatus) IsValid() bool {
	for _, existing := range allowedScheduledTaskLastExecutionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScheduledTaskLastExecutionStatus value.
func (v ScheduledTaskLastExecutionStatus) Ptr() *ScheduledTaskLastExecutionStatus {
	return &v
}
