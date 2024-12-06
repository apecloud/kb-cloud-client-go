// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
)

// EventSource event source
type EventSource string

// List of EventSource.
const (
	EventSourceUser   EventSource = "user"
	EventSourceSystem EventSource = "system"
)

var allowedEventSourceEnumValues = []EventSource{
	EventSourceUser,
	EventSourceSystem,
}

// GetAllowedValues returns the list of possible values.
func (v *EventSource) GetAllowedValues() []EventSource {
	return allowedEventSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EventSource) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EventSource(value)
	return nil
}

// NewEventSourceFromValue returns a pointer to a valid EventSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEventSourceFromValue(v string) (*EventSource, error) {
	ev := EventSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EventSource: valid values are %v", v, allowedEventSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EventSource) IsValid() bool {
	for _, existing := range allowedEventSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventSource value.
func (v EventSource) Ptr() *EventSource {
	return &v
}
