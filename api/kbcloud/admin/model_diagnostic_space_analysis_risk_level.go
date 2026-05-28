// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticSpaceAnalysisRiskLevel Object space risk level derived only from percentOfTotal.
type DiagnosticSpaceAnalysisRiskLevel string

// List of DiagnosticSpaceAnalysisRiskLevel.
const (
	DiagnosticSpaceAnalysisRiskLevelNormal   DiagnosticSpaceAnalysisRiskLevel = "normal"
	DiagnosticSpaceAnalysisRiskLevelWarning  DiagnosticSpaceAnalysisRiskLevel = "warning"
	DiagnosticSpaceAnalysisRiskLevelCritical DiagnosticSpaceAnalysisRiskLevel = "critical"
)

var allowedDiagnosticSpaceAnalysisRiskLevelEnumValues = []DiagnosticSpaceAnalysisRiskLevel{
	DiagnosticSpaceAnalysisRiskLevelNormal,
	DiagnosticSpaceAnalysisRiskLevelWarning,
	DiagnosticSpaceAnalysisRiskLevelCritical,
}

// GetAllowedValues returns the list of possible values.
func (v *DiagnosticSpaceAnalysisRiskLevel) GetAllowedValues() []DiagnosticSpaceAnalysisRiskLevel {
	return allowedDiagnosticSpaceAnalysisRiskLevelEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DiagnosticSpaceAnalysisRiskLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DiagnosticSpaceAnalysisRiskLevel(value)
	return nil
}

// NewDiagnosticSpaceAnalysisRiskLevelFromValue returns a pointer to a valid DiagnosticSpaceAnalysisRiskLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDiagnosticSpaceAnalysisRiskLevelFromValue(v string) (*DiagnosticSpaceAnalysisRiskLevel, error) {
	ev := DiagnosticSpaceAnalysisRiskLevel(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DiagnosticSpaceAnalysisRiskLevel: valid values are %v", v, allowedDiagnosticSpaceAnalysisRiskLevelEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DiagnosticSpaceAnalysisRiskLevel) IsValid() bool {
	for _, existing := range allowedDiagnosticSpaceAnalysisRiskLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DiagnosticSpaceAnalysisRiskLevel value.
func (v DiagnosticSpaceAnalysisRiskLevel) Ptr() *DiagnosticSpaceAnalysisRiskLevel {
	return &v
}
