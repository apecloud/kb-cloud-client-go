// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticCheckSeverity User-facing severity and priority bucket for one diagnostic check.
type DiagnosticCheckSeverity string

// List of DiagnosticCheckSeverity.
const (
	DiagnosticCheckSeverityCritical         DiagnosticCheckSeverity = "critical"
	DiagnosticCheckSeveritySevere           DiagnosticCheckSeverity = "severe"
	DiagnosticCheckSeverityWarning          DiagnosticCheckSeverity = "warning"
	DiagnosticCheckSeverityNormal           DiagnosticCheckSeverity = "normal"
	DiagnosticCheckSeverityNotApplicable    DiagnosticCheckSeverity = "not_applicable"
	DiagnosticCheckSeverityCollectionFailed DiagnosticCheckSeverity = "collection_failed"
)

var allowedDiagnosticCheckSeverityEnumValues = []DiagnosticCheckSeverity{
	DiagnosticCheckSeverityCritical,
	DiagnosticCheckSeveritySevere,
	DiagnosticCheckSeverityWarning,
	DiagnosticCheckSeverityNormal,
	DiagnosticCheckSeverityNotApplicable,
	DiagnosticCheckSeverityCollectionFailed,
}

// GetAllowedValues returns the list of possible values.
func (v *DiagnosticCheckSeverity) GetAllowedValues() []DiagnosticCheckSeverity {
	return allowedDiagnosticCheckSeverityEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DiagnosticCheckSeverity) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DiagnosticCheckSeverity(value)
	return nil
}

// NewDiagnosticCheckSeverityFromValue returns a pointer to a valid DiagnosticCheckSeverity
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDiagnosticCheckSeverityFromValue(v string) (*DiagnosticCheckSeverity, error) {
	ev := DiagnosticCheckSeverity(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DiagnosticCheckSeverity: valid values are %v", v, allowedDiagnosticCheckSeverityEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DiagnosticCheckSeverity) IsValid() bool {
	for _, existing := range allowedDiagnosticCheckSeverityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DiagnosticCheckSeverity value.
func (v DiagnosticCheckSeverity) Ptr() *DiagnosticCheckSeverity {
	return &v
}
