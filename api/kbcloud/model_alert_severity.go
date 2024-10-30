// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package kbcloud

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)



type AlertSeverity string

// List of AlertSeverity.
const (
	AlertSeverityCritical AlertSeverity = "critical"
	AlertSeverityWarning AlertSeverity = "warning"
	AlertSeverityInfo AlertSeverity = "info"
)

var allowedAlertSeverityEnumValues = []AlertSeverity{
	AlertSeverityCritical,
	AlertSeverityWarning,
	AlertSeverityInfo,
}

// GetAllowedValues returns the list of possible values.
func (v *AlertSeverity) GetAllowedValues() []AlertSeverity {
	return allowedAlertSeverityEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AlertSeverity) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AlertSeverity(value)
	return nil
}

// NewAlertSeverityFromValue returns a pointer to a valid AlertSeverity
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAlertSeverityFromValue(v string) (*AlertSeverity, error) {
	ev := AlertSeverity(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AlertSeverity: valid values are %v", v, allowedAlertSeverityEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AlertSeverity) IsValid() bool {
	for _, existing := range allowedAlertSeverityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AlertSeverity value.
func (v AlertSeverity) Ptr() *AlertSeverity {
	return &v
}
