// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type EventType string

// List of EventType.
const (
	EventTypeDml EventType = "dml"
	EventTypeDdl EventType = "ddl"
	EventTypeDcl EventType = "dcl"
)

var allowedEventTypeEnumValues = []EventType{
	EventTypeDml,
	EventTypeDdl,
	EventTypeDcl,
}

// GetAllowedValues returns the list of possible values.
func (v *EventType) GetAllowedValues() []EventType {
	return allowedEventTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EventType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EventType(value)
	return nil
}

// NewEventTypeFromValue returns a pointer to a valid EventType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEventTypeFromValue(v string) (*EventType, error) {
	ev := EventType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EventType: valid values are %v", v, allowedEventTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EventType) IsValid() bool {
	for _, existing := range allowedEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventType value.
func (v EventType) Ptr() *EventType {
	return &v
}
