// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticCPUAnalysisReasonCode Stable reason code for unavailable or degraded CPU analysis.
type DiagnosticCPUAnalysisReasonCode string

// List of DiagnosticCPUAnalysisReasonCode.
const (
	DiagnosticCPUAnalysisReasonCodeNon                     DiagnosticCPUAnalysisReasonCode = "none"
	DiagnosticCPUAnalysisReasonCodeNoCpuRiskEvidence       DiagnosticCPUAnalysisReasonCode = "no_cpu_risk_evidence"
	DiagnosticCPUAnalysisReasonCodeDataSourceNotIntegrated DiagnosticCPUAnalysisReasonCode = "data_source_not_integrated"
	DiagnosticCPUAnalysisReasonCodePermissionDenied        DiagnosticCPUAnalysisReasonCode = "permission_denied"
	DiagnosticCPUAnalysisReasonCodeQueryTimeout            DiagnosticCPUAnalysisReasonCode = "query_timeout"
	DiagnosticCPUAnalysisReasonCodePrimaryPodNotFound      DiagnosticCPUAnalysisReasonCode = "primary_pod_not_found"
	DiagnosticCPUAnalysisReasonCodeMetricsUnavailable      DiagnosticCPUAnalysisReasonCode = "metrics_unavailable"
	DiagnosticCPUAnalysisReasonCodeCollectionError         DiagnosticCPUAnalysisReasonCode = "collection_error"
	DiagnosticCPUAnalysisReasonCodeUnexpectedOutput        DiagnosticCPUAnalysisReasonCode = "unexpected_output"
)

var allowedDiagnosticCPUAnalysisReasonCodeEnumValues = []DiagnosticCPUAnalysisReasonCode{
	DiagnosticCPUAnalysisReasonCodeNon,
	DiagnosticCPUAnalysisReasonCodeNoCpuRiskEvidence,
	DiagnosticCPUAnalysisReasonCodeDataSourceNotIntegrated,
	DiagnosticCPUAnalysisReasonCodePermissionDenied,
	DiagnosticCPUAnalysisReasonCodeQueryTimeout,
	DiagnosticCPUAnalysisReasonCodePrimaryPodNotFound,
	DiagnosticCPUAnalysisReasonCodeMetricsUnavailable,
	DiagnosticCPUAnalysisReasonCodeCollectionError,
	DiagnosticCPUAnalysisReasonCodeUnexpectedOutput,
}

// GetAllowedValues returns the list of possible values.
func (v *DiagnosticCPUAnalysisReasonCode) GetAllowedValues() []DiagnosticCPUAnalysisReasonCode {
	return allowedDiagnosticCPUAnalysisReasonCodeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DiagnosticCPUAnalysisReasonCode) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DiagnosticCPUAnalysisReasonCode(value)
	return nil
}

// NewDiagnosticCPUAnalysisReasonCodeFromValue returns a pointer to a valid DiagnosticCPUAnalysisReasonCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDiagnosticCPUAnalysisReasonCodeFromValue(v string) (*DiagnosticCPUAnalysisReasonCode, error) {
	ev := DiagnosticCPUAnalysisReasonCode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DiagnosticCPUAnalysisReasonCode: valid values are %v", v, allowedDiagnosticCPUAnalysisReasonCodeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DiagnosticCPUAnalysisReasonCode) IsValid() bool {
	for _, existing := range allowedDiagnosticCPUAnalysisReasonCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DiagnosticCPUAnalysisReasonCode value.
func (v DiagnosticCPUAnalysisReasonCode) Ptr() *DiagnosticCPUAnalysisReasonCode {
	return &v
}
