// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticCPUAnalysisState Business-result axis for PostgreSQL readonly CPU analysis. with_data means
// at least one CPU-related evidence partition was collected. no_data means
// the collector ran but did not find CPU risk evidence in the current window.
// collection_failed and source_missing keep unavailable data sources explicit.
type DiagnosticCPUAnalysisState string

// List of DiagnosticCPUAnalysisState.
const (
	DiagnosticCPUAnalysisStateSourcePresentWithData         DiagnosticCPUAnalysisState = "source_present_with_data"
	DiagnosticCPUAnalysisStateSourcePresentNoData           DiagnosticCPUAnalysisState = "source_present_no_data"
	DiagnosticCPUAnalysisStateSourcePresentCollectionFailed DiagnosticCPUAnalysisState = "source_present_collection_failed"
	DiagnosticCPUAnalysisStateSourceMissing                 DiagnosticCPUAnalysisState = "source_missing"
)

var allowedDiagnosticCPUAnalysisStateEnumValues = []DiagnosticCPUAnalysisState{
	DiagnosticCPUAnalysisStateSourcePresentWithData,
	DiagnosticCPUAnalysisStateSourcePresentNoData,
	DiagnosticCPUAnalysisStateSourcePresentCollectionFailed,
	DiagnosticCPUAnalysisStateSourceMissing,
}

// GetAllowedValues returns the list of possible values.
func (v *DiagnosticCPUAnalysisState) GetAllowedValues() []DiagnosticCPUAnalysisState {
	return allowedDiagnosticCPUAnalysisStateEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DiagnosticCPUAnalysisState) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DiagnosticCPUAnalysisState(value)
	return nil
}

// NewDiagnosticCPUAnalysisStateFromValue returns a pointer to a valid DiagnosticCPUAnalysisState
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDiagnosticCPUAnalysisStateFromValue(v string) (*DiagnosticCPUAnalysisState, error) {
	ev := DiagnosticCPUAnalysisState(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DiagnosticCPUAnalysisState: valid values are %v", v, allowedDiagnosticCPUAnalysisStateEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DiagnosticCPUAnalysisState) IsValid() bool {
	for _, existing := range allowedDiagnosticCPUAnalysisStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DiagnosticCPUAnalysisState value.
func (v DiagnosticCPUAnalysisState) Ptr() *DiagnosticCPUAnalysisState {
	return &v
}
