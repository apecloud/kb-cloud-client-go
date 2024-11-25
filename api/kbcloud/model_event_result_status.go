// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
)

// EventResultStatus result status of the operation event
type EventResultStatus string

// List of EventResultStatus.
const (
	EventResultStatusSuccess EventResultStatus = "success"
	EventResultStatusFailed  EventResultStatus = "failed"
)

var allowedEventResultStatusEnumValues = []EventResultStatus{
	EventResultStatusSuccess,
	EventResultStatusFailed,
}

// GetAllowedValues returns the list of possible values.
func (v *EventResultStatus) GetAllowedValues() []EventResultStatus {
	return allowedEventResultStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EventResultStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EventResultStatus(value)
	return nil
}

// NewEventResultStatusFromValue returns a pointer to a valid EventResultStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEventResultStatusFromValue(v string) (*EventResultStatus, error) {
	ev := EventResultStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EventResultStatus: valid values are %v", v, allowedEventResultStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EventResultStatus) IsValid() bool {
	for _, existing := range allowedEventResultStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventResultStatus value.
func (v EventResultStatus) Ptr() *EventResultStatus {
	return &v
}
