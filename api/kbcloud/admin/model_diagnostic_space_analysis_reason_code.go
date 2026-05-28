// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticSpaceAnalysisReasonCode Stable reason code for unavailable or degraded space analysis.
type DiagnosticSpaceAnalysisReasonCode string

// List of DiagnosticSpaceAnalysisReasonCode.
const (
	DiagnosticSpaceAnalysisReasonCodeNon                     DiagnosticSpaceAnalysisReasonCode = "none"
	DiagnosticSpaceAnalysisReasonCodeNoSpaceObjects          DiagnosticSpaceAnalysisReasonCode = "no_space_objects"
	DiagnosticSpaceAnalysisReasonCodeDataSourceNotIntegrated DiagnosticSpaceAnalysisReasonCode = "data_source_not_integrated"
	DiagnosticSpaceAnalysisReasonCodePermissionDenied        DiagnosticSpaceAnalysisReasonCode = "permission_denied"
	DiagnosticSpaceAnalysisReasonCodeQueryTimeout            DiagnosticSpaceAnalysisReasonCode = "query_timeout"
	DiagnosticSpaceAnalysisReasonCodePrimaryPodNotFound      DiagnosticSpaceAnalysisReasonCode = "primary_pod_not_found"
	DiagnosticSpaceAnalysisReasonCodeCollectionError         DiagnosticSpaceAnalysisReasonCode = "collection_error"
	DiagnosticSpaceAnalysisReasonCodeUnexpectedOutput        DiagnosticSpaceAnalysisReasonCode = "unexpected_output"
)

var allowedDiagnosticSpaceAnalysisReasonCodeEnumValues = []DiagnosticSpaceAnalysisReasonCode{
	DiagnosticSpaceAnalysisReasonCodeNon,
	DiagnosticSpaceAnalysisReasonCodeNoSpaceObjects,
	DiagnosticSpaceAnalysisReasonCodeDataSourceNotIntegrated,
	DiagnosticSpaceAnalysisReasonCodePermissionDenied,
	DiagnosticSpaceAnalysisReasonCodeQueryTimeout,
	DiagnosticSpaceAnalysisReasonCodePrimaryPodNotFound,
	DiagnosticSpaceAnalysisReasonCodeCollectionError,
	DiagnosticSpaceAnalysisReasonCodeUnexpectedOutput,
}

// GetAllowedValues returns the list of possible values.
func (v *DiagnosticSpaceAnalysisReasonCode) GetAllowedValues() []DiagnosticSpaceAnalysisReasonCode {
	return allowedDiagnosticSpaceAnalysisReasonCodeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DiagnosticSpaceAnalysisReasonCode) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DiagnosticSpaceAnalysisReasonCode(value)
	return nil
}

// NewDiagnosticSpaceAnalysisReasonCodeFromValue returns a pointer to a valid DiagnosticSpaceAnalysisReasonCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDiagnosticSpaceAnalysisReasonCodeFromValue(v string) (*DiagnosticSpaceAnalysisReasonCode, error) {
	ev := DiagnosticSpaceAnalysisReasonCode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DiagnosticSpaceAnalysisReasonCode: valid values are %v", v, allowedDiagnosticSpaceAnalysisReasonCodeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DiagnosticSpaceAnalysisReasonCode) IsValid() bool {
	for _, existing := range allowedDiagnosticSpaceAnalysisReasonCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DiagnosticSpaceAnalysisReasonCode value.
func (v DiagnosticSpaceAnalysisReasonCode) Ptr() *DiagnosticSpaceAnalysisReasonCode {
	return &v
}
