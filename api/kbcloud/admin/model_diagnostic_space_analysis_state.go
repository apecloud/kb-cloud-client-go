// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticSpaceAnalysisState Business-result axis for PostgreSQL readonly space analysis. It separates
// "collector ran and found Top objects" from "collector ran but there was no
// user object", "collector failed", and the compatibility state where no
// supported data source is available.
type DiagnosticSpaceAnalysisState string

// List of DiagnosticSpaceAnalysisState.
const (
	DiagnosticSpaceAnalysisStateSourcePresentWithData         DiagnosticSpaceAnalysisState = "source_present_with_data"
	DiagnosticSpaceAnalysisStateSourcePresentNoData           DiagnosticSpaceAnalysisState = "source_present_no_data"
	DiagnosticSpaceAnalysisStateSourcePresentCollectionFailed DiagnosticSpaceAnalysisState = "source_present_collection_failed"
	DiagnosticSpaceAnalysisStateSourceMissing                 DiagnosticSpaceAnalysisState = "source_missing"
)

var allowedDiagnosticSpaceAnalysisStateEnumValues = []DiagnosticSpaceAnalysisState{
	DiagnosticSpaceAnalysisStateSourcePresentWithData,
	DiagnosticSpaceAnalysisStateSourcePresentNoData,
	DiagnosticSpaceAnalysisStateSourcePresentCollectionFailed,
	DiagnosticSpaceAnalysisStateSourceMissing,
}

// GetAllowedValues returns the list of possible values.
func (v *DiagnosticSpaceAnalysisState) GetAllowedValues() []DiagnosticSpaceAnalysisState {
	return allowedDiagnosticSpaceAnalysisStateEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DiagnosticSpaceAnalysisState) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DiagnosticSpaceAnalysisState(value)
	return nil
}

// NewDiagnosticSpaceAnalysisStateFromValue returns a pointer to a valid DiagnosticSpaceAnalysisState
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDiagnosticSpaceAnalysisStateFromValue(v string) (*DiagnosticSpaceAnalysisState, error) {
	ev := DiagnosticSpaceAnalysisState(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DiagnosticSpaceAnalysisState: valid values are %v", v, allowedDiagnosticSpaceAnalysisStateEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DiagnosticSpaceAnalysisState) IsValid() bool {
	for _, existing := range allowedDiagnosticSpaceAnalysisStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DiagnosticSpaceAnalysisState value.
func (v DiagnosticSpaceAnalysisState) Ptr() *DiagnosticSpaceAnalysisState {
	return &v
}
