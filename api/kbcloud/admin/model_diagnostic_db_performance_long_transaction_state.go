// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticDBPerformanceLongTransactionState Business-result axis for the PostgreSQL long-transaction metric. Live
// collection produces the three states source_present_with_data,
// source_present_no_data, and source_present_collection_failed.
// source_missing is the compatibility state used when the collector is not
// enabled or no PostgreSQL pod is available.
type DiagnosticDBPerformanceLongTransactionState string

// List of DiagnosticDBPerformanceLongTransactionState.
const (
	DiagnosticDBPerformanceLongTransactionStateSourcePresentWithData         DiagnosticDBPerformanceLongTransactionState = "source_present_with_data"
	DiagnosticDBPerformanceLongTransactionStateSourcePresentNoData           DiagnosticDBPerformanceLongTransactionState = "source_present_no_data"
	DiagnosticDBPerformanceLongTransactionStateSourcePresentCollectionFailed DiagnosticDBPerformanceLongTransactionState = "source_present_collection_failed"
	DiagnosticDBPerformanceLongTransactionStateSourceMissing                 DiagnosticDBPerformanceLongTransactionState = "source_missing"
)

var allowedDiagnosticDBPerformanceLongTransactionStateEnumValues = []DiagnosticDBPerformanceLongTransactionState{
	DiagnosticDBPerformanceLongTransactionStateSourcePresentWithData,
	DiagnosticDBPerformanceLongTransactionStateSourcePresentNoData,
	DiagnosticDBPerformanceLongTransactionStateSourcePresentCollectionFailed,
	DiagnosticDBPerformanceLongTransactionStateSourceMissing,
}

// GetAllowedValues returns the list of possible values.
func (v *DiagnosticDBPerformanceLongTransactionState) GetAllowedValues() []DiagnosticDBPerformanceLongTransactionState {
	return allowedDiagnosticDBPerformanceLongTransactionStateEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DiagnosticDBPerformanceLongTransactionState) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DiagnosticDBPerformanceLongTransactionState(value)
	return nil
}

// NewDiagnosticDBPerformanceLongTransactionStateFromValue returns a pointer to a valid DiagnosticDBPerformanceLongTransactionState
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDiagnosticDBPerformanceLongTransactionStateFromValue(v string) (*DiagnosticDBPerformanceLongTransactionState, error) {
	ev := DiagnosticDBPerformanceLongTransactionState(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DiagnosticDBPerformanceLongTransactionState: valid values are %v", v, allowedDiagnosticDBPerformanceLongTransactionStateEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DiagnosticDBPerformanceLongTransactionState) IsValid() bool {
	for _, existing := range allowedDiagnosticDBPerformanceLongTransactionStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DiagnosticDBPerformanceLongTransactionState value.
func (v DiagnosticDBPerformanceLongTransactionState) Ptr() *DiagnosticDBPerformanceLongTransactionState {
	return &v
}
