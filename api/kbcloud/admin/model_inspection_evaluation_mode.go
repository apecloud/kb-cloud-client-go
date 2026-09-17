// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// InspectionEvaluationMode Empty legacy values mean check. Display collects information without scoring. Manual requires type=manual and presents scriptExpr verbatim as instructions; it executes no SQL or templates, returns severity=unknown, and retains rule importance.
type InspectionEvaluationMode string

// List of InspectionEvaluationMode.
const (
	InspectionEvaluationModeCheck   InspectionEvaluationMode = "check"
	InspectionEvaluationModeDisplay InspectionEvaluationMode = "display"
	InspectionEvaluationModeManual  InspectionEvaluationMode = "manual"
)

var allowedInspectionEvaluationModeEnumValues = []InspectionEvaluationMode{
	InspectionEvaluationModeCheck,
	InspectionEvaluationModeDisplay,
	InspectionEvaluationModeManual,
}

// GetAllowedValues returns the list of possible values.
func (v *InspectionEvaluationMode) GetAllowedValues() []InspectionEvaluationMode {
	return allowedInspectionEvaluationModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *InspectionEvaluationMode) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = InspectionEvaluationMode(value)
	return nil
}

// NewInspectionEvaluationModeFromValue returns a pointer to a valid InspectionEvaluationMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewInspectionEvaluationModeFromValue(v string) (*InspectionEvaluationMode, error) {
	ev := InspectionEvaluationMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for InspectionEvaluationMode: valid values are %v", v, allowedInspectionEvaluationModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v InspectionEvaluationMode) IsValid() bool {
	for _, existing := range allowedInspectionEvaluationModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InspectionEvaluationMode value.
func (v InspectionEvaluationMode) Ptr() *InspectionEvaluationMode {
	return &v
}
