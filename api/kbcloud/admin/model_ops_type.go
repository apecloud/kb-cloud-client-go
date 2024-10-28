// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// NODESCRIPTION OpsType
type OpsType string

// List of OpsType.
const (
	OPSTYPE_VERTICALSCALING   OpsType = "VerticalScaling"
	OPSTYPE_HORIZONTALSCALING OpsType = "HorizontalScaling"
	OPSTYPE_VOLUMEEXPANSION   OpsType = "VolumeExpansion"
	OPSTYPE_UPGRADE           OpsType = "Upgrade"
	OPSTYPE_RECONFIGURING     OpsType = "Reconfiguring"
	OPSTYPE_SWITCHOVER        OpsType = "Switchover"
	OPSTYPE_RESTART           OpsType = "Restart"
	OPSTYPE_STOP              OpsType = "Stop"
	OPSTYPE_START             OpsType = "Start"
	OPSTYPE_EXPOSE            OpsType = "Expose"
	OPSTYPE_DATASCRIPT        OpsType = "DataScript"
)

var allowedOpsTypeEnumValues = []OpsType{
	OPSTYPE_VERTICALSCALING,
	OPSTYPE_HORIZONTALSCALING,
	OPSTYPE_VOLUMEEXPANSION,
	OPSTYPE_UPGRADE,
	OPSTYPE_RECONFIGURING,
	OPSTYPE_SWITCHOVER,
	OPSTYPE_RESTART,
	OPSTYPE_STOP,
	OPSTYPE_START,
	OPSTYPE_EXPOSE,
	OPSTYPE_DATASCRIPT,
}

// GetAllowedValues reeturns the list of possible values.
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
