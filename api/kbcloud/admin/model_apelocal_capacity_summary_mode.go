// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ApelocalCapacitySummaryMode Conflicting shared quotas cannot form a single logical total.
type ApelocalCapacitySummaryMode string

// List of ApelocalCapacitySummaryMode.
const (
	ApelocalCapacitySummaryModeFinite    ApelocalCapacitySummaryMode = "finite"
	ApelocalCapacitySummaryModeUnlimited ApelocalCapacitySummaryMode = "unlimited"
	ApelocalCapacitySummaryModeConflict  ApelocalCapacitySummaryMode = "conflict"
)

var allowedApelocalCapacitySummaryModeEnumValues = []ApelocalCapacitySummaryMode{
	ApelocalCapacitySummaryModeFinite,
	ApelocalCapacitySummaryModeUnlimited,
	ApelocalCapacitySummaryModeConflict,
}

// GetAllowedValues returns the list of possible values.
func (v *ApelocalCapacitySummaryMode) GetAllowedValues() []ApelocalCapacitySummaryMode {
	return allowedApelocalCapacitySummaryModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ApelocalCapacitySummaryMode) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ApelocalCapacitySummaryMode(value)
	return nil
}

// NewApelocalCapacitySummaryModeFromValue returns a pointer to a valid ApelocalCapacitySummaryMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewApelocalCapacitySummaryModeFromValue(v string) (*ApelocalCapacitySummaryMode, error) {
	ev := ApelocalCapacitySummaryMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ApelocalCapacitySummaryMode: valid values are %v", v, allowedApelocalCapacitySummaryModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ApelocalCapacitySummaryMode) IsValid() bool {
	for _, existing := range allowedApelocalCapacitySummaryModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApelocalCapacitySummaryMode value.
func (v ApelocalCapacitySummaryMode) Ptr() *ApelocalCapacitySummaryMode {
	return &v
}
