// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticCPUAnalysisPrimaryCause Best-effort evidence bucket for the first screen. This is not a root-cause
// proof and must be read with cannotProve.
type DiagnosticCPUAnalysisPrimaryCause string

// List of DiagnosticCPUAnalysisPrimaryCause.
const (
	DiagnosticCPUAnalysisPrimaryCauseSqlPressure          DiagnosticCPUAnalysisPrimaryCause = "sql_pressure"
	DiagnosticCPUAnalysisPrimaryCauseSessionPressure      DiagnosticCPUAnalysisPrimaryCause = "session_pressure"
	DiagnosticCPUAnalysisPrimaryCauseResourceLimit        DiagnosticCPUAnalysisPrimaryCause = "resource_limit"
	DiagnosticCPUAnalysisPrimaryCauseBackgroundActivity   DiagnosticCPUAnalysisPrimaryCause = "background_activity"
	DiagnosticCPUAnalysisPrimaryCauseMixed                DiagnosticCPUAnalysisPrimaryCause = "mixed"
	DiagnosticCPUAnalysisPrimaryCauseEvidenceInsufficient DiagnosticCPUAnalysisPrimaryCause = "evidence_insufficient"
)

var allowedDiagnosticCPUAnalysisPrimaryCauseEnumValues = []DiagnosticCPUAnalysisPrimaryCause{
	DiagnosticCPUAnalysisPrimaryCauseSqlPressure,
	DiagnosticCPUAnalysisPrimaryCauseSessionPressure,
	DiagnosticCPUAnalysisPrimaryCauseResourceLimit,
	DiagnosticCPUAnalysisPrimaryCauseBackgroundActivity,
	DiagnosticCPUAnalysisPrimaryCauseMixed,
	DiagnosticCPUAnalysisPrimaryCauseEvidenceInsufficient,
}

// GetAllowedValues returns the list of possible values.
func (v *DiagnosticCPUAnalysisPrimaryCause) GetAllowedValues() []DiagnosticCPUAnalysisPrimaryCause {
	return allowedDiagnosticCPUAnalysisPrimaryCauseEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DiagnosticCPUAnalysisPrimaryCause) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DiagnosticCPUAnalysisPrimaryCause(value)
	return nil
}

// NewDiagnosticCPUAnalysisPrimaryCauseFromValue returns a pointer to a valid DiagnosticCPUAnalysisPrimaryCause
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDiagnosticCPUAnalysisPrimaryCauseFromValue(v string) (*DiagnosticCPUAnalysisPrimaryCause, error) {
	ev := DiagnosticCPUAnalysisPrimaryCause(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DiagnosticCPUAnalysisPrimaryCause: valid values are %v", v, allowedDiagnosticCPUAnalysisPrimaryCauseEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DiagnosticCPUAnalysisPrimaryCause) IsValid() bool {
	for _, existing := range allowedDiagnosticCPUAnalysisPrimaryCauseEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DiagnosticCPUAnalysisPrimaryCause value.
func (v DiagnosticCPUAnalysisPrimaryCause) Ptr() *DiagnosticCPUAnalysisPrimaryCause {
	return &v
}
