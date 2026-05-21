// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticDBPerformanceSlowSQLState Business-result axis for the slow SQL metric. Sprint 9 v0.1 four-state contract.
// Separate from dataQuality (collection-quality axis) so the Console can render the
// correct main text without overloading dataQuality. Each enum value pairs with a
// fixed (dataQuality, reasonCode, count) tuple enforced by backend Validate.
type DiagnosticDBPerformanceSlowSQLState string

// List of DiagnosticDBPerformanceSlowSQLState.
const (
	DiagnosticDBPerformanceSlowSQLStateSourcePresentWithData         DiagnosticDBPerformanceSlowSQLState = "source_present_with_data"
	DiagnosticDBPerformanceSlowSQLStateSourcePresentNoDataInWindow   DiagnosticDBPerformanceSlowSQLState = "source_present_no_data_in_window"
	DiagnosticDBPerformanceSlowSQLStateSourcePresentCollectionFailed DiagnosticDBPerformanceSlowSQLState = "source_present_collection_failed"
	DiagnosticDBPerformanceSlowSQLStateSourceMissing                 DiagnosticDBPerformanceSlowSQLState = "source_missing"
)

var allowedDiagnosticDBPerformanceSlowSQLStateEnumValues = []DiagnosticDBPerformanceSlowSQLState{
	DiagnosticDBPerformanceSlowSQLStateSourcePresentWithData,
	DiagnosticDBPerformanceSlowSQLStateSourcePresentNoDataInWindow,
	DiagnosticDBPerformanceSlowSQLStateSourcePresentCollectionFailed,
	DiagnosticDBPerformanceSlowSQLStateSourceMissing,
}

// GetAllowedValues returns the list of possible values.
func (v *DiagnosticDBPerformanceSlowSQLState) GetAllowedValues() []DiagnosticDBPerformanceSlowSQLState {
	return allowedDiagnosticDBPerformanceSlowSQLStateEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DiagnosticDBPerformanceSlowSQLState) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DiagnosticDBPerformanceSlowSQLState(value)
	return nil
}

// NewDiagnosticDBPerformanceSlowSQLStateFromValue returns a pointer to a valid DiagnosticDBPerformanceSlowSQLState
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDiagnosticDBPerformanceSlowSQLStateFromValue(v string) (*DiagnosticDBPerformanceSlowSQLState, error) {
	ev := DiagnosticDBPerformanceSlowSQLState(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DiagnosticDBPerformanceSlowSQLState: valid values are %v", v, allowedDiagnosticDBPerformanceSlowSQLStateEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DiagnosticDBPerformanceSlowSQLState) IsValid() bool {
	for _, existing := range allowedDiagnosticDBPerformanceSlowSQLStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DiagnosticDBPerformanceSlowSQLState value.
func (v DiagnosticDBPerformanceSlowSQLState) Ptr() *DiagnosticDBPerformanceSlowSQLState {
	return &v
}
