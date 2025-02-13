// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EventQueryFilterType Console do not support orgName and operator.
type EventQueryFilterType string

// List of EventQueryFilterType.
const (
	EventQueryFilterTypeOrgName      EventQueryFilterType = "orgName"
	EventQueryFilterTypeOperator     EventQueryFilterType = "operator"
	EventQueryFilterTypeEventName    EventQueryFilterType = "eventName"
	EventQueryFilterTypeResourceName EventQueryFilterType = "resourceName"
	EventQueryFilterTypeResourceType EventQueryFilterType = "resourceType"
)

var allowedEventQueryFilterTypeEnumValues = []EventQueryFilterType{
	EventQueryFilterTypeOrgName,
	EventQueryFilterTypeOperator,
	EventQueryFilterTypeEventName,
	EventQueryFilterTypeResourceName,
	EventQueryFilterTypeResourceType,
}

// GetAllowedValues returns the list of possible values.
func (v *EventQueryFilterType) GetAllowedValues() []EventQueryFilterType {
	return allowedEventQueryFilterTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EventQueryFilterType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EventQueryFilterType(value)
	return nil
}

// NewEventQueryFilterTypeFromValue returns a pointer to a valid EventQueryFilterType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEventQueryFilterTypeFromValue(v string) (*EventQueryFilterType, error) {
	ev := EventQueryFilterType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EventQueryFilterType: valid values are %v", v, allowedEventQueryFilterTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EventQueryFilterType) IsValid() bool {
	for _, existing := range allowedEventQueryFilterTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventQueryFilterType value.
func (v EventQueryFilterType) Ptr() *EventQueryFilterType {
	return &v
}
