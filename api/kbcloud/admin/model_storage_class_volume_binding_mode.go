// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// StorageClassVolumeBindingMode volumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound. Defaults to Immediate.
type StorageClassVolumeBindingMode string

// List of StorageClassVolumeBindingMode.
const (
	STORAGECLASSVOLUMEBINDINGMODE_IMMEDIATE            StorageClassVolumeBindingMode = "Immediate"
	STORAGECLASSVOLUMEBINDINGMODE_WAITFORFIRSTCONSUMER StorageClassVolumeBindingMode = "WaitForFirstConsumer"
)

var allowedStorageClassVolumeBindingModeEnumValues = []StorageClassVolumeBindingMode{
	STORAGECLASSVOLUMEBINDINGMODE_IMMEDIATE,
	STORAGECLASSVOLUMEBINDINGMODE_WAITFORFIRSTCONSUMER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *StorageClassVolumeBindingMode) GetAllowedValues() []StorageClassVolumeBindingMode {
	return allowedStorageClassVolumeBindingModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *StorageClassVolumeBindingMode) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = StorageClassVolumeBindingMode(value)
	return nil
}

// NewStorageClassVolumeBindingModeFromValue returns a pointer to a valid StorageClassVolumeBindingMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewStorageClassVolumeBindingModeFromValue(v string) (*StorageClassVolumeBindingMode, error) {
	ev := StorageClassVolumeBindingMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for StorageClassVolumeBindingMode: valid values are %v", v, allowedStorageClassVolumeBindingModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v StorageClassVolumeBindingMode) IsValid() bool {
	for _, existing := range allowedStorageClassVolumeBindingModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StorageClassVolumeBindingMode value.
func (v StorageClassVolumeBindingMode) Ptr() *StorageClassVolumeBindingMode {
	return &v
}
