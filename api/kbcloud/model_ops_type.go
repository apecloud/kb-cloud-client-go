// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type OpsType string

// List of OpsType.
const (
	OpsTypeVerticalScaling   OpsType = "VerticalScaling"
	OpsTypeHorizontalScaling OpsType = "HorizontalScaling"
	OpsTypeVolumeExpansion   OpsType = "VolumeExpansion"
	OpsTypeUpgrade           OpsType = "Upgrade"
	OpsTypeReconfiguring     OpsType = "Reconfiguring"
	OpsTypeSwitchover        OpsType = "Switchover"
	OpsTypeRestart           OpsType = "Restart"
	OpsTypeStop              OpsType = "Stop"
	OpsTypeStart             OpsType = "Start"
	OpsTypeExpose            OpsType = "Expose"
	OpsTypeBackup            OpsType = "Backup"
	OpsTypeRestore           OpsType = "Restore"
	OpsTypeRebuildInstance   OpsType = "RebuildInstance"
	OpsTypeCustom            OpsType = "Custom"
	OpsTypeUpdateLicense     OpsType = "UpdateLicense"
)

var allowedOpsTypeEnumValues = []OpsType{
	OpsTypeVerticalScaling,
	OpsTypeHorizontalScaling,
	OpsTypeVolumeExpansion,
	OpsTypeUpgrade,
	OpsTypeReconfiguring,
	OpsTypeSwitchover,
	OpsTypeRestart,
	OpsTypeStop,
	OpsTypeStart,
	OpsTypeExpose,
	OpsTypeBackup,
	OpsTypeRestore,
	OpsTypeRebuildInstance,
	OpsTypeCustom,
	OpsTypeUpdateLicense,
}

// GetAllowedValues returns the list of possible values.
func (v *OpsType) GetAllowedValues() []OpsType {
	return allowedOpsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OpsType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OpsType(value)
	return nil
}

// NewOpsTypeFromValue returns a pointer to a valid OpsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOpsTypeFromValue(v string) (*OpsType, error) {
	ev := OpsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OpsType: valid values are %v", v, allowedOpsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OpsType) IsValid() bool {
	for _, existing := range allowedOpsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OpsType value.
func (v OpsType) Ptr() *OpsType {
	return &v
}
