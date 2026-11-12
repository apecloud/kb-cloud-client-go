// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ScheduledTaskEnableStatus Whether the schedule is immutable, disabled, or enabled
type ScheduledTaskEnableStatus string

// List of ScheduledTaskEnableStatus.
const (
	ScheduledTaskEnableStatusImmutable ScheduledTaskEnableStatus = "immutable"
	ScheduledTaskEnableStatusDisabled  ScheduledTaskEnableStatus = "disabled"
	ScheduledTaskEnableStatusEnabled   ScheduledTaskEnableStatus = "enabled"
)

var allowedScheduledTaskEnableStatusEnumValues = []ScheduledTaskEnableStatus{
	ScheduledTaskEnableStatusImmutable,
	ScheduledTaskEnableStatusDisabled,
	ScheduledTaskEnableStatusEnabled,
}

// GetAllowedValues returns the list of possible values.
func (v *ScheduledTaskEnableStatus) GetAllowedValues() []ScheduledTaskEnableStatus {
	return allowedScheduledTaskEnableStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScheduledTaskEnableStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScheduledTaskEnableStatus(value)
	return nil
}

// NewScheduledTaskEnableStatusFromValue returns a pointer to a valid ScheduledTaskEnableStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScheduledTaskEnableStatusFromValue(v string) (*ScheduledTaskEnableStatus, error) {
	ev := ScheduledTaskEnableStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScheduledTaskEnableStatus: valid values are %v", v, allowedScheduledTaskEnableStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScheduledTaskEnableStatus) IsValid() bool {
	for _, existing := range allowedScheduledTaskEnableStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScheduledTaskEnableStatus value.
func (v ScheduledTaskEnableStatus) Ptr() *ScheduledTaskEnableStatus {
	return &v
}
