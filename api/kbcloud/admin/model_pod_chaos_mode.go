// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// PodChaosMode specify the mode of chaos
type PodChaosMode string

// List of PodChaosMode.
const (
	PODCHAOSMODE_ONE                PodChaosMode = "one"
	PODCHAOSMODE_ALL                PodChaosMode = "all"
	PODCHAOSMODE_FIXED              PodChaosMode = "fixed"
	PODCHAOSMODE_FIXED_PERCENT      PodChaosMode = "fixed-percent"
	PODCHAOSMODE_RANDOM_MAX_PERCENT PodChaosMode = "random-max-percent"
)

var allowedPodChaosModeEnumValues = []PodChaosMode{
	PODCHAOSMODE_ONE,
	PODCHAOSMODE_ALL,
	PODCHAOSMODE_FIXED,
	PODCHAOSMODE_FIXED_PERCENT,
	PODCHAOSMODE_RANDOM_MAX_PERCENT,
}

// GetAllowedValues returns the list of possible values.
func (v *PodChaosMode) GetAllowedValues() []PodChaosMode {
	return allowedPodChaosModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PodChaosMode) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PodChaosMode(value)
	return nil
}

// NewPodChaosModeFromValue returns a pointer to a valid PodChaosMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPodChaosModeFromValue(v string) (*PodChaosMode, error) {
	ev := PodChaosMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PodChaosMode: valid values are %v", v, allowedPodChaosModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PodChaosMode) IsValid() bool {
	for _, existing := range allowedPodChaosModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PodChaosMode value.
func (v PodChaosMode) Ptr() *PodChaosMode {
	return &v
}
