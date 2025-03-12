// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EventFilterType Console do not support orgName.
type EventFilterType string

// List of EventFilterType.
const (
	EventFilterTypeOrgName      EventFilterType = "orgName"
	EventFilterTypeOperator     EventFilterType = "operator"
	EventFilterTypeEventName    EventFilterType = "eventName"
	EventFilterTypeResourceName EventFilterType = "resourceName"
	EventFilterTypeResourceType EventFilterType = "resourceType"
)

var allowedEventFilterTypeEnumValues = []EventFilterType{
	EventFilterTypeOrgName,
	EventFilterTypeOperator,
	EventFilterTypeEventName,
	EventFilterTypeResourceName,
	EventFilterTypeResourceType,
}

// GetAllowedValues returns the list of possible values.
func (v *EventFilterType) GetAllowedValues() []EventFilterType {
	return allowedEventFilterTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EventFilterType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EventFilterType(value)
	return nil
}

// NewEventFilterTypeFromValue returns a pointer to a valid EventFilterType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEventFilterTypeFromValue(v string) (*EventFilterType, error) {
	ev := EventFilterType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EventFilterType: valid values are %v", v, allowedEventFilterTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EventFilterType) IsValid() bool {
	for _, existing := range allowedEventFilterTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventFilterType value.
func (v EventFilterType) Ptr() *EventFilterType {
	return &v
}
