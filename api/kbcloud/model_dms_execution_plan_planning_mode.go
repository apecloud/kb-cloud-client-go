// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DmsExecutionPlanPlanningMode string

// List of DmsExecutionPlanPlanningMode.
const (
	DmsExecutionPlanPlanningModeStandard DmsExecutionPlanPlanningMode = "standard"
	DmsExecutionPlanPlanningModeGeneric  DmsExecutionPlanPlanningMode = "generic"
)

var allowedDmsExecutionPlanPlanningModeEnumValues = []DmsExecutionPlanPlanningMode{
	DmsExecutionPlanPlanningModeStandard,
	DmsExecutionPlanPlanningModeGeneric,
}

// GetAllowedValues returns the list of possible values.
func (v *DmsExecutionPlanPlanningMode) GetAllowedValues() []DmsExecutionPlanPlanningMode {
	return allowedDmsExecutionPlanPlanningModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DmsExecutionPlanPlanningMode) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DmsExecutionPlanPlanningMode(value)
	return nil
}

// NewDmsExecutionPlanPlanningModeFromValue returns a pointer to a valid DmsExecutionPlanPlanningMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDmsExecutionPlanPlanningModeFromValue(v string) (*DmsExecutionPlanPlanningMode, error) {
	ev := DmsExecutionPlanPlanningMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DmsExecutionPlanPlanningMode: valid values are %v", v, allowedDmsExecutionPlanPlanningModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DmsExecutionPlanPlanningMode) IsValid() bool {
	for _, existing := range allowedDmsExecutionPlanPlanningModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DmsExecutionPlanPlanningMode value.
func (v DmsExecutionPlanPlanningMode) Ptr() *DmsExecutionPlanPlanningMode {
	return &v
}
