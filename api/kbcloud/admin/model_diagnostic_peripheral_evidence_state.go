// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticPeripheralEvidenceState Collection-result axis for the Sprint 24 K8s / Operator / Event explanation layer.
// with_data means at least one external evidence explanation is available.
// The remaining states explain why explanations are empty or incomplete without
// inventing a database-internal conclusion.
type DiagnosticPeripheralEvidenceState string

// List of DiagnosticPeripheralEvidenceState.
const (
	DiagnosticPeripheralEvidenceStateWithData              DiagnosticPeripheralEvidenceState = "with_data"
	DiagnosticPeripheralEvidenceStateNoData                DiagnosticPeripheralEvidenceState = "no_data"
	DiagnosticPeripheralEvidenceStateNoPermission          DiagnosticPeripheralEvidenceState = "no_permission"
	DiagnosticPeripheralEvidenceStateNoLogWindow           DiagnosticPeripheralEvidenceState = "no_log_window"
	DiagnosticPeripheralEvidenceStateDatasourceUnavailable DiagnosticPeripheralEvidenceState = "datasource_unavailable"
	DiagnosticPeripheralEvidenceStateCollectionFailed      DiagnosticPeripheralEvidenceState = "collection_failed"
)

var allowedDiagnosticPeripheralEvidenceStateEnumValues = []DiagnosticPeripheralEvidenceState{
	DiagnosticPeripheralEvidenceStateWithData,
	DiagnosticPeripheralEvidenceStateNoData,
	DiagnosticPeripheralEvidenceStateNoPermission,
	DiagnosticPeripheralEvidenceStateNoLogWindow,
	DiagnosticPeripheralEvidenceStateDatasourceUnavailable,
	DiagnosticPeripheralEvidenceStateCollectionFailed,
}

// GetAllowedValues returns the list of possible values.
func (v *DiagnosticPeripheralEvidenceState) GetAllowedValues() []DiagnosticPeripheralEvidenceState {
	return allowedDiagnosticPeripheralEvidenceStateEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DiagnosticPeripheralEvidenceState) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DiagnosticPeripheralEvidenceState(value)
	return nil
}

// NewDiagnosticPeripheralEvidenceStateFromValue returns a pointer to a valid DiagnosticPeripheralEvidenceState
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDiagnosticPeripheralEvidenceStateFromValue(v string) (*DiagnosticPeripheralEvidenceState, error) {
	ev := DiagnosticPeripheralEvidenceState(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DiagnosticPeripheralEvidenceState: valid values are %v", v, allowedDiagnosticPeripheralEvidenceStateEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DiagnosticPeripheralEvidenceState) IsValid() bool {
	for _, existing := range allowedDiagnosticPeripheralEvidenceStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DiagnosticPeripheralEvidenceState value.
func (v DiagnosticPeripheralEvidenceState) Ptr() *DiagnosticPeripheralEvidenceState {
	return &v
}
