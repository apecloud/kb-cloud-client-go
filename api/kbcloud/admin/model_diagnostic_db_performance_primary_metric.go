// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticDBPerformancePrimaryMetric Metric key the user should inspect first. One of slowSql, longTransaction, lockWait, none, or collectionFailed.
type DiagnosticDBPerformancePrimaryMetric string

// List of DiagnosticDBPerformancePrimaryMetric.
const (
	DiagnosticDBPerformancePrimaryMetricSlowSql          DiagnosticDBPerformancePrimaryMetric = "slowSql"
	DiagnosticDBPerformancePrimaryMetricLongTransaction  DiagnosticDBPerformancePrimaryMetric = "longTransaction"
	DiagnosticDBPerformancePrimaryMetricLockWait         DiagnosticDBPerformancePrimaryMetric = "lockWait"
	DiagnosticDBPerformancePrimaryMetricNon              DiagnosticDBPerformancePrimaryMetric = "none"
	DiagnosticDBPerformancePrimaryMetricCollectionFailed DiagnosticDBPerformancePrimaryMetric = "collectionFailed"
)

var allowedDiagnosticDBPerformancePrimaryMetricEnumValues = []DiagnosticDBPerformancePrimaryMetric{
	DiagnosticDBPerformancePrimaryMetricSlowSql,
	DiagnosticDBPerformancePrimaryMetricLongTransaction,
	DiagnosticDBPerformancePrimaryMetricLockWait,
	DiagnosticDBPerformancePrimaryMetricNon,
	DiagnosticDBPerformancePrimaryMetricCollectionFailed,
}

// GetAllowedValues returns the list of possible values.
func (v *DiagnosticDBPerformancePrimaryMetric) GetAllowedValues() []DiagnosticDBPerformancePrimaryMetric {
	return allowedDiagnosticDBPerformancePrimaryMetricEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DiagnosticDBPerformancePrimaryMetric) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DiagnosticDBPerformancePrimaryMetric(value)
	return nil
}

// NewDiagnosticDBPerformancePrimaryMetricFromValue returns a pointer to a valid DiagnosticDBPerformancePrimaryMetric
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDiagnosticDBPerformancePrimaryMetricFromValue(v string) (*DiagnosticDBPerformancePrimaryMetric, error) {
	ev := DiagnosticDBPerformancePrimaryMetric(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DiagnosticDBPerformancePrimaryMetric: valid values are %v", v, allowedDiagnosticDBPerformancePrimaryMetricEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DiagnosticDBPerformancePrimaryMetric) IsValid() bool {
	for _, existing := range allowedDiagnosticDBPerformancePrimaryMetricEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DiagnosticDBPerformancePrimaryMetric value.
func (v DiagnosticDBPerformancePrimaryMetric) Ptr() *DiagnosticDBPerformancePrimaryMetric {
	return &v
}
