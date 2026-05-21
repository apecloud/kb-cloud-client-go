// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticDBPerformanceMetricReasonCode Structured reason code for unavailable or degraded database performance metrics.
type DiagnosticDBPerformanceMetricReasonCode string

// List of DiagnosticDBPerformanceMetricReasonCode.
const (
	DiagnosticDBPerformanceMetricReasonCodeDataSourceNotIntegrated DiagnosticDBPerformanceMetricReasonCode = "data_source_not_integrated"
	DiagnosticDBPerformanceMetricReasonCodeEngineNotSupported      DiagnosticDBPerformanceMetricReasonCode = "engine_not_supported"
	DiagnosticDBPerformanceMetricReasonCodeCollectionError         DiagnosticDBPerformanceMetricReasonCode = "collection_error"
)

var allowedDiagnosticDBPerformanceMetricReasonCodeEnumValues = []DiagnosticDBPerformanceMetricReasonCode{
	DiagnosticDBPerformanceMetricReasonCodeDataSourceNotIntegrated,
	DiagnosticDBPerformanceMetricReasonCodeEngineNotSupported,
	DiagnosticDBPerformanceMetricReasonCodeCollectionError,
}

// GetAllowedValues returns the list of possible values.
func (v *DiagnosticDBPerformanceMetricReasonCode) GetAllowedValues() []DiagnosticDBPerformanceMetricReasonCode {
	return allowedDiagnosticDBPerformanceMetricReasonCodeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DiagnosticDBPerformanceMetricReasonCode) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DiagnosticDBPerformanceMetricReasonCode(value)
	return nil
}

// NewDiagnosticDBPerformanceMetricReasonCodeFromValue returns a pointer to a valid DiagnosticDBPerformanceMetricReasonCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDiagnosticDBPerformanceMetricReasonCodeFromValue(v string) (*DiagnosticDBPerformanceMetricReasonCode, error) {
	ev := DiagnosticDBPerformanceMetricReasonCode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DiagnosticDBPerformanceMetricReasonCode: valid values are %v", v, allowedDiagnosticDBPerformanceMetricReasonCodeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DiagnosticDBPerformanceMetricReasonCode) IsValid() bool {
	for _, existing := range allowedDiagnosticDBPerformanceMetricReasonCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DiagnosticDBPerformanceMetricReasonCode value.
func (v DiagnosticDBPerformanceMetricReasonCode) Ptr() *DiagnosticDBPerformanceMetricReasonCode {
	return &v
}
