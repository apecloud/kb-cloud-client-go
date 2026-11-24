// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ApelocalCapacityMode CSI capacity policy; unlimited does not represent physical capacity.
type ApelocalCapacityMode string

// List of ApelocalCapacityMode.
const (
	ApelocalCapacityModeAsIs       ApelocalCapacityMode = "asIs"
	ApelocalCapacityModeMultiplier ApelocalCapacityMode = "multiplier"
	ApelocalCapacityModeFixed      ApelocalCapacityMode = "fixed"
	ApelocalCapacityModeUnlimited  ApelocalCapacityMode = "unlimited"
)

var allowedApelocalCapacityModeEnumValues = []ApelocalCapacityMode{
	ApelocalCapacityModeAsIs,
	ApelocalCapacityModeMultiplier,
	ApelocalCapacityModeFixed,
	ApelocalCapacityModeUnlimited,
}

// GetAllowedValues returns the list of possible values.
func (v *ApelocalCapacityMode) GetAllowedValues() []ApelocalCapacityMode {
	return allowedApelocalCapacityModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ApelocalCapacityMode) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ApelocalCapacityMode(value)
	return nil
}

// NewApelocalCapacityModeFromValue returns a pointer to a valid ApelocalCapacityMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewApelocalCapacityModeFromValue(v string) (*ApelocalCapacityMode, error) {
	ev := ApelocalCapacityMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ApelocalCapacityMode: valid values are %v", v, allowedApelocalCapacityModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ApelocalCapacityMode) IsValid() bool {
	for _, existing := range allowedApelocalCapacityModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApelocalCapacityMode value.
func (v ApelocalCapacityMode) Ptr() *ApelocalCapacityMode {
	return &v
}
