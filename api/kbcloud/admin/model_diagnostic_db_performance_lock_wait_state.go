// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticDBPerformanceLockWaitState Business-result axis for the PostgreSQL lock-wait metric. Live collection
// produces the three states source_present_with_data, source_present_no_data,
// and source_present_collection_failed. source_missing is the compatibility
// state used when the collector is not enabled or no PostgreSQL pod is
// available.
type DiagnosticDBPerformanceLockWaitState string

// List of DiagnosticDBPerformanceLockWaitState.
const (
	DiagnosticDBPerformanceLockWaitStateSourcePresentWithData         DiagnosticDBPerformanceLockWaitState = "source_present_with_data"
	DiagnosticDBPerformanceLockWaitStateSourcePresentNoData           DiagnosticDBPerformanceLockWaitState = "source_present_no_data"
	DiagnosticDBPerformanceLockWaitStateSourcePresentCollectionFailed DiagnosticDBPerformanceLockWaitState = "source_present_collection_failed"
	DiagnosticDBPerformanceLockWaitStateSourceMissing                 DiagnosticDBPerformanceLockWaitState = "source_missing"
)

var allowedDiagnosticDBPerformanceLockWaitStateEnumValues = []DiagnosticDBPerformanceLockWaitState{
	DiagnosticDBPerformanceLockWaitStateSourcePresentWithData,
	DiagnosticDBPerformanceLockWaitStateSourcePresentNoData,
	DiagnosticDBPerformanceLockWaitStateSourcePresentCollectionFailed,
	DiagnosticDBPerformanceLockWaitStateSourceMissing,
}

// GetAllowedValues returns the list of possible values.
func (v *DiagnosticDBPerformanceLockWaitState) GetAllowedValues() []DiagnosticDBPerformanceLockWaitState {
	return allowedDiagnosticDBPerformanceLockWaitStateEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DiagnosticDBPerformanceLockWaitState) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DiagnosticDBPerformanceLockWaitState(value)
	return nil
}

// NewDiagnosticDBPerformanceLockWaitStateFromValue returns a pointer to a valid DiagnosticDBPerformanceLockWaitState
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDiagnosticDBPerformanceLockWaitStateFromValue(v string) (*DiagnosticDBPerformanceLockWaitState, error) {
	ev := DiagnosticDBPerformanceLockWaitState(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DiagnosticDBPerformanceLockWaitState: valid values are %v", v, allowedDiagnosticDBPerformanceLockWaitStateEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DiagnosticDBPerformanceLockWaitState) IsValid() bool {
	for _, existing := range allowedDiagnosticDBPerformanceLockWaitStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DiagnosticDBPerformanceLockWaitState value.
func (v DiagnosticDBPerformanceLockWaitState) Ptr() *DiagnosticDBPerformanceLockWaitState {
	return &v
}
