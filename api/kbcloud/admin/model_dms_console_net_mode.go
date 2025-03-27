// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DMSConsoleNetMode string

// List of DMSConsoleNetMode.
const (
	DMSConsoleNetModePrivate DMSConsoleNetMode = "PRIVATE"
	DMSConsoleNetModePublic  DMSConsoleNetMode = "PUBLIC"
)

var allowedDMSConsoleNetModeEnumValues = []DMSConsoleNetMode{
	DMSConsoleNetModePrivate,
	DMSConsoleNetModePublic,
}

// GetAllowedValues returns the list of possible values.
func (v *DMSConsoleNetMode) GetAllowedValues() []DMSConsoleNetMode {
	return allowedDMSConsoleNetModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DMSConsoleNetMode) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DMSConsoleNetMode(value)
	return nil
}

// NewDMSConsoleNetModeFromValue returns a pointer to a valid DMSConsoleNetMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDMSConsoleNetModeFromValue(v string) (*DMSConsoleNetMode, error) {
	ev := DMSConsoleNetMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DMSConsoleNetMode: valid values are %v", v, allowedDMSConsoleNetModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DMSConsoleNetMode) IsValid() bool {
	for _, existing := range allowedDMSConsoleNetModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DMSConsoleNetMode value.
func (v DMSConsoleNetMode) Ptr() *DMSConsoleNetMode {
	return &v
}
