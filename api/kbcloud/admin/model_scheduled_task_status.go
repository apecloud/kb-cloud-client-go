// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ScheduledTaskStatus Current normalized status of the scheduled task
type ScheduledTaskStatus string

// List of ScheduledTaskStatus.
const (
	ScheduledTaskStatusActive   ScheduledTaskStatus = "active"
	ScheduledTaskStatusDisabled ScheduledTaskStatus = "disabled"
	ScheduledTaskStatusError    ScheduledTaskStatus = "error"
	ScheduledTaskStatusUnknown  ScheduledTaskStatus = "unknown"
)

var allowedScheduledTaskStatusEnumValues = []ScheduledTaskStatus{
	ScheduledTaskStatusActive,
	ScheduledTaskStatusDisabled,
	ScheduledTaskStatusError,
	ScheduledTaskStatusUnknown,
}

// GetAllowedValues returns the list of possible values.
func (v *ScheduledTaskStatus) GetAllowedValues() []ScheduledTaskStatus {
	return allowedScheduledTaskStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScheduledTaskStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScheduledTaskStatus(value)
	return nil
}

// NewScheduledTaskStatusFromValue returns a pointer to a valid ScheduledTaskStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScheduledTaskStatusFromValue(v string) (*ScheduledTaskStatus, error) {
	ev := ScheduledTaskStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScheduledTaskStatus: valid values are %v", v, allowedScheduledTaskStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScheduledTaskStatus) IsValid() bool {
	for _, existing := range allowedScheduledTaskStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScheduledTaskStatus value.
func (v ScheduledTaskStatus) Ptr() *ScheduledTaskStatus {
	return &v
}
