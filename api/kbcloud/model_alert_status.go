// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
)

type AlertStatus string

// List of AlertStatus.
const (
	AlertStatusResolved AlertStatus = "resolved"
	AlertStatusFiring   AlertStatus = "firing"
)

var allowedAlertStatusEnumValues = []AlertStatus{
	AlertStatusResolved,
	AlertStatusFiring,
}

// GetAllowedValues returns the list of possible values.
func (v *AlertStatus) GetAllowedValues() []AlertStatus {
	return allowedAlertStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AlertStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AlertStatus(value)
	return nil
}

// NewAlertStatusFromValue returns a pointer to a valid AlertStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAlertStatusFromValue(v string) (*AlertStatus, error) {
	ev := AlertStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AlertStatus: valid values are %v", v, allowedAlertStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AlertStatus) IsValid() bool {
	for _, existing := range allowedAlertStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AlertStatus value.
func (v AlertStatus) Ptr() *AlertStatus {
	return &v
}
