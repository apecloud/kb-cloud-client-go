// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// StorageCapacityMode CSI capacity policy; unlimited does not represent physical capacity.
type StorageCapacityMode string

// List of StorageCapacityMode.
const (
	StorageCapacityModeAsIs       StorageCapacityMode = "asIs"
	StorageCapacityModeMultiplier StorageCapacityMode = "multiplier"
	StorageCapacityModeFixed      StorageCapacityMode = "fixed"
	StorageCapacityModeUnlimited  StorageCapacityMode = "unlimited"
)

var allowedStorageCapacityModeEnumValues = []StorageCapacityMode{
	StorageCapacityModeAsIs,
	StorageCapacityModeMultiplier,
	StorageCapacityModeFixed,
	StorageCapacityModeUnlimited,
}

// GetAllowedValues returns the list of possible values.
func (v *StorageCapacityMode) GetAllowedValues() []StorageCapacityMode {
	return allowedStorageCapacityModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *StorageCapacityMode) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = StorageCapacityMode(value)
	return nil
}

// NewStorageCapacityModeFromValue returns a pointer to a valid StorageCapacityMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewStorageCapacityModeFromValue(v string) (*StorageCapacityMode, error) {
	ev := StorageCapacityMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for StorageCapacityMode: valid values are %v", v, allowedStorageCapacityModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v StorageCapacityMode) IsValid() bool {
	for _, existing := range allowedStorageCapacityModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StorageCapacityMode value.
func (v StorageCapacityMode) Ptr() *StorageCapacityMode {
	return &v
}
