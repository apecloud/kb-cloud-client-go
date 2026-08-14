// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DmsExecutionPlanMode string

// List of DmsExecutionPlanMode.
const (
	DmsExecutionPlanModeEstimate  DmsExecutionPlanMode = "estimate"
	DmsExecutionPlanModeAnalyze   DmsExecutionPlanMode = "analyze"
	DmsExecutionPlanModeActual    DmsExecutionPlanMode = "actual"
	DmsExecutionPlanModeAutotrace DmsExecutionPlanMode = "autotrace"
)

var allowedDmsExecutionPlanModeEnumValues = []DmsExecutionPlanMode{
	DmsExecutionPlanModeEstimate,
	DmsExecutionPlanModeAnalyze,
	DmsExecutionPlanModeActual,
	DmsExecutionPlanModeAutotrace,
}

// GetAllowedValues returns the list of possible values.
func (v *DmsExecutionPlanMode) GetAllowedValues() []DmsExecutionPlanMode {
	return allowedDmsExecutionPlanModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DmsExecutionPlanMode) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DmsExecutionPlanMode(value)
	return nil
}

// NewDmsExecutionPlanModeFromValue returns a pointer to a valid DmsExecutionPlanMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDmsExecutionPlanModeFromValue(v string) (*DmsExecutionPlanMode, error) {
	ev := DmsExecutionPlanMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DmsExecutionPlanMode: valid values are %v", v, allowedDmsExecutionPlanModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DmsExecutionPlanMode) IsValid() bool {
	for _, existing := range allowedDmsExecutionPlanModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DmsExecutionPlanMode value.
func (v DmsExecutionPlanMode) Ptr() *DmsExecutionPlanMode {
	return &v
}
