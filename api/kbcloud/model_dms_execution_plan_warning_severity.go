// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DmsExecutionPlanWarningSeverity string

// List of DmsExecutionPlanWarningSeverity.
const (
	DmsExecutionPlanWarningSeverityInfo     DmsExecutionPlanWarningSeverity = "info"
	DmsExecutionPlanWarningSeverityWarning  DmsExecutionPlanWarningSeverity = "warning"
	DmsExecutionPlanWarningSeverityCritical DmsExecutionPlanWarningSeverity = "critical"
)

var allowedDmsExecutionPlanWarningSeverityEnumValues = []DmsExecutionPlanWarningSeverity{
	DmsExecutionPlanWarningSeverityInfo,
	DmsExecutionPlanWarningSeverityWarning,
	DmsExecutionPlanWarningSeverityCritical,
}

// GetAllowedValues returns the list of possible values.
func (v *DmsExecutionPlanWarningSeverity) GetAllowedValues() []DmsExecutionPlanWarningSeverity {
	return allowedDmsExecutionPlanWarningSeverityEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DmsExecutionPlanWarningSeverity) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DmsExecutionPlanWarningSeverity(value)
	return nil
}

// NewDmsExecutionPlanWarningSeverityFromValue returns a pointer to a valid DmsExecutionPlanWarningSeverity
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDmsExecutionPlanWarningSeverityFromValue(v string) (*DmsExecutionPlanWarningSeverity, error) {
	ev := DmsExecutionPlanWarningSeverity(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DmsExecutionPlanWarningSeverity: valid values are %v", v, allowedDmsExecutionPlanWarningSeverityEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DmsExecutionPlanWarningSeverity) IsValid() bool {
	for _, existing := range allowedDmsExecutionPlanWarningSeverityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DmsExecutionPlanWarningSeverity value.
func (v DmsExecutionPlanWarningSeverity) Ptr() *DmsExecutionPlanWarningSeverity {
	return &v
}
