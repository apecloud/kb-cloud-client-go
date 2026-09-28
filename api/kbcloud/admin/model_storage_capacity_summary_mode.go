// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// StorageCapacitySummaryMode Conflicting shared quotas cannot form a single logical total.
type StorageCapacitySummaryMode string

// List of StorageCapacitySummaryMode.
const (
	StorageCapacitySummaryModeFinite    StorageCapacitySummaryMode = "finite"
	StorageCapacitySummaryModeUnlimited StorageCapacitySummaryMode = "unlimited"
	StorageCapacitySummaryModeConflict  StorageCapacitySummaryMode = "conflict"
)

var allowedStorageCapacitySummaryModeEnumValues = []StorageCapacitySummaryMode{
	StorageCapacitySummaryModeFinite,
	StorageCapacitySummaryModeUnlimited,
	StorageCapacitySummaryModeConflict,
}

// GetAllowedValues returns the list of possible values.
func (v *StorageCapacitySummaryMode) GetAllowedValues() []StorageCapacitySummaryMode {
	return allowedStorageCapacitySummaryModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *StorageCapacitySummaryMode) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = StorageCapacitySummaryMode(value)
	return nil
}

// NewStorageCapacitySummaryModeFromValue returns a pointer to a valid StorageCapacitySummaryMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewStorageCapacitySummaryModeFromValue(v string) (*StorageCapacitySummaryMode, error) {
	ev := StorageCapacitySummaryMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for StorageCapacitySummaryMode: valid values are %v", v, allowedStorageCapacitySummaryModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v StorageCapacitySummaryMode) IsValid() bool {
	for _, existing := range allowedStorageCapacitySummaryModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StorageCapacitySummaryMode value.
func (v StorageCapacitySummaryMode) Ptr() *StorageCapacitySummaryMode {
	return &v
}
