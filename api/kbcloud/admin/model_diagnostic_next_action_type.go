// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticNextActionType Safe operation type for a diagnostic next action. Console uses this to decide whether to render a clickable action, a copy action, a rerun entry, or a manual suggestion.
type DiagnosticNextActionType string

// List of DiagnosticNextActionType.
const (
	DiagnosticNextActionTypeOpenDetail         DiagnosticNextActionType = "open_detail"
	DiagnosticNextActionTypeJumpToEvidence     DiagnosticNextActionType = "jump_to_evidence"
	DiagnosticNextActionTypeCopyContext        DiagnosticNextActionType = "copy_context"
	DiagnosticNextActionTypeRerunDiagnosis     DiagnosticNextActionType = "rerun_diagnosis"
	DiagnosticNextActionTypeExternalManualStep DiagnosticNextActionType = "external_manual_step"
)

var allowedDiagnosticNextActionTypeEnumValues = []DiagnosticNextActionType{
	DiagnosticNextActionTypeOpenDetail,
	DiagnosticNextActionTypeJumpToEvidence,
	DiagnosticNextActionTypeCopyContext,
	DiagnosticNextActionTypeRerunDiagnosis,
	DiagnosticNextActionTypeExternalManualStep,
}

// GetAllowedValues returns the list of possible values.
func (v *DiagnosticNextActionType) GetAllowedValues() []DiagnosticNextActionType {
	return allowedDiagnosticNextActionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DiagnosticNextActionType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DiagnosticNextActionType(value)
	return nil
}

// NewDiagnosticNextActionTypeFromValue returns a pointer to a valid DiagnosticNextActionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDiagnosticNextActionTypeFromValue(v string) (*DiagnosticNextActionType, error) {
	ev := DiagnosticNextActionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DiagnosticNextActionType: valid values are %v", v, allowedDiagnosticNextActionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DiagnosticNextActionType) IsValid() bool {
	for _, existing := range allowedDiagnosticNextActionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DiagnosticNextActionType value.
func (v DiagnosticNextActionType) Ptr() *DiagnosticNextActionType {
	return &v
}
