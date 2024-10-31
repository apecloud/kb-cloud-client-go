// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EventResourceType Type of the resource
type EventResourceType string

// List of EventResourceType.
const (
	EventResourceTypeCluster EventResourceType = "cluster"
	EventResourceTypeRole    EventResourceType = "role"
)

var allowedEventResourceTypeEnumValues = []EventResourceType{
	EventResourceTypeCluster,
	EventResourceTypeRole,
}

// GetAllowedValues returns the list of possible values.
func (v *EventResourceType) GetAllowedValues() []EventResourceType {
	return allowedEventResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EventResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EventResourceType(value)
	return nil
}

// NewEventResourceTypeFromValue returns a pointer to a valid EventResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEventResourceTypeFromValue(v string) (*EventResourceType, error) {
	ev := EventResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EventResourceType: valid values are %v", v, allowedEventResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EventResourceType) IsValid() bool {
	for _, existing := range allowedEventResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventResourceType value.
func (v EventResourceType) Ptr() *EventResourceType {
	return &v
}
