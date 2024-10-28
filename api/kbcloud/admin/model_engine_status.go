// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EngineStatus engine addon status in K8s
type EngineStatus string

// List of EngineStatus.
const (
	ENGINESTATUS_DISABLED  EngineStatus = "Disabled"
	ENGINESTATUS_ENABLED   EngineStatus = "Enabled"
	ENGINESTATUS_FAILED    EngineStatus = "Failed"
	ENGINESTATUS_ENABLING  EngineStatus = "Enabling"
	ENGINESTATUS_DISABLING EngineStatus = "Disabling"
)

var allowedEngineStatusEnumValues = []EngineStatus{
	ENGINESTATUS_DISABLED,
	ENGINESTATUS_ENABLED,
	ENGINESTATUS_FAILED,
	ENGINESTATUS_ENABLING,
	ENGINESTATUS_DISABLING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EngineStatus) GetAllowedValues() []EngineStatus {
	return allowedEngineStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EngineStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EngineStatus(value)
	return nil
}

// NewEngineStatusFromValue returns a pointer to a valid EngineStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEngineStatusFromValue(v string) (*EngineStatus, error) {
	ev := EngineStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EngineStatus: valid values are %v", v, allowedEngineStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EngineStatus) IsValid() bool {
	for _, existing := range allowedEngineStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EngineStatus value.
func (v EngineStatus) Ptr() *EngineStatus {
	return &v
}
