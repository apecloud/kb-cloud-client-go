// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticCPUAnalysisRiskLevel CPU risk level derived from evidence thresholds, not from an automatic fix decision.
type DiagnosticCPUAnalysisRiskLevel string

// List of DiagnosticCPUAnalysisRiskLevel.
const (
	DiagnosticCPUAnalysisRiskLevelNormal   DiagnosticCPUAnalysisRiskLevel = "normal"
	DiagnosticCPUAnalysisRiskLevelWarning  DiagnosticCPUAnalysisRiskLevel = "warning"
	DiagnosticCPUAnalysisRiskLevelCritical DiagnosticCPUAnalysisRiskLevel = "critical"
)

var allowedDiagnosticCPUAnalysisRiskLevelEnumValues = []DiagnosticCPUAnalysisRiskLevel{
	DiagnosticCPUAnalysisRiskLevelNormal,
	DiagnosticCPUAnalysisRiskLevelWarning,
	DiagnosticCPUAnalysisRiskLevelCritical,
}

// GetAllowedValues returns the list of possible values.
func (v *DiagnosticCPUAnalysisRiskLevel) GetAllowedValues() []DiagnosticCPUAnalysisRiskLevel {
	return allowedDiagnosticCPUAnalysisRiskLevelEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DiagnosticCPUAnalysisRiskLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DiagnosticCPUAnalysisRiskLevel(value)
	return nil
}

// NewDiagnosticCPUAnalysisRiskLevelFromValue returns a pointer to a valid DiagnosticCPUAnalysisRiskLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDiagnosticCPUAnalysisRiskLevelFromValue(v string) (*DiagnosticCPUAnalysisRiskLevel, error) {
	ev := DiagnosticCPUAnalysisRiskLevel(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DiagnosticCPUAnalysisRiskLevel: valid values are %v", v, allowedDiagnosticCPUAnalysisRiskLevelEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DiagnosticCPUAnalysisRiskLevel) IsValid() bool {
	for _, existing := range allowedDiagnosticCPUAnalysisRiskLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DiagnosticCPUAnalysisRiskLevel value.
func (v DiagnosticCPUAnalysisRiskLevel) Ptr() *DiagnosticCPUAnalysisRiskLevel {
	return &v
}
