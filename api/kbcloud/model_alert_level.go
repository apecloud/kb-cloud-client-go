// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AlertLevel string

// List of AlertLevel.
const (
	AlertLevelGlobal       AlertLevel = "global"
	AlertLevelOrganization AlertLevel = "organization"
)

var allowedAlertLevelEnumValues = []AlertLevel{
	AlertLevelGlobal,
	AlertLevelOrganization,
}

// GetAllowedValues returns the list of possible values.
func (v *AlertLevel) GetAllowedValues() []AlertLevel {
	return allowedAlertLevelEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AlertLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AlertLevel(value)
	return nil
}

// NewAlertLevelFromValue returns a pointer to a valid AlertLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAlertLevelFromValue(v string) (*AlertLevel, error) {
	ev := AlertLevel(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AlertLevel: valid values are %v", v, allowedAlertLevelEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AlertLevel) IsValid() bool {
	for _, existing := range allowedAlertLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AlertLevel value.
func (v AlertLevel) Ptr() *AlertLevel {
	return &v
}
