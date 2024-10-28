// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// NetworkChaosMode Specifies the mode of network chaos to be applied. This determines how the chaos is distributed among the target pods.
type NetworkChaosMode string

// List of NetworkChaosMode.
const (
	NETWORKCHAOSMODE_ONE                NetworkChaosMode = "one"
	NETWORKCHAOSMODE_ALL                NetworkChaosMode = "all"
	NETWORKCHAOSMODE_FIXED              NetworkChaosMode = "fixed"
	NETWORKCHAOSMODE_FIXED_PERCENT      NetworkChaosMode = "fixed-percent"
	NETWORKCHAOSMODE_RANDOM_MAX_PERCENT NetworkChaosMode = "random-max-percent"
)

var allowedNetworkChaosModeEnumValues = []NetworkChaosMode{
	NETWORKCHAOSMODE_ONE,
	NETWORKCHAOSMODE_ALL,
	NETWORKCHAOSMODE_FIXED,
	NETWORKCHAOSMODE_FIXED_PERCENT,
	NETWORKCHAOSMODE_RANDOM_MAX_PERCENT,
}

// GetAllowedValues returns the list of possible values.
func (v *NetworkChaosMode) GetAllowedValues() []NetworkChaosMode {
	return allowedNetworkChaosModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *NetworkChaosMode) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NetworkChaosMode(value)
	return nil
}

// NewNetworkChaosModeFromValue returns a pointer to a valid NetworkChaosMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewNetworkChaosModeFromValue(v string) (*NetworkChaosMode, error) {
	ev := NetworkChaosMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NetworkChaosMode: valid values are %v", v, allowedNetworkChaosModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v NetworkChaosMode) IsValid() bool {
	for _, existing := range allowedNetworkChaosModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NetworkChaosMode value.
func (v NetworkChaosMode) Ptr() *NetworkChaosMode {
	return &v
}
