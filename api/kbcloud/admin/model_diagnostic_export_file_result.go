// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticExportFileResult Collection result for one file included in a redacted diagnostic package.
type DiagnosticExportFileResult string

// List of DiagnosticExportFileResult.
const (
	DiagnosticExportFileResultSuccess       DiagnosticExportFileResult = "success"
	DiagnosticExportFileResultFailed        DiagnosticExportFileResult = "failed"
	DiagnosticExportFileResultSkipped       DiagnosticExportFileResult = "skipped"
	DiagnosticExportFileResultNotApplicable DiagnosticExportFileResult = "not_applicable"
)

var allowedDiagnosticExportFileResultEnumValues = []DiagnosticExportFileResult{
	DiagnosticExportFileResultSuccess,
	DiagnosticExportFileResultFailed,
	DiagnosticExportFileResultSkipped,
	DiagnosticExportFileResultNotApplicable,
}

// GetAllowedValues returns the list of possible values.
func (v *DiagnosticExportFileResult) GetAllowedValues() []DiagnosticExportFileResult {
	return allowedDiagnosticExportFileResultEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DiagnosticExportFileResult) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DiagnosticExportFileResult(value)
	return nil
}

// NewDiagnosticExportFileResultFromValue returns a pointer to a valid DiagnosticExportFileResult
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDiagnosticExportFileResultFromValue(v string) (*DiagnosticExportFileResult, error) {
	ev := DiagnosticExportFileResult(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DiagnosticExportFileResult: valid values are %v", v, allowedDiagnosticExportFileResultEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DiagnosticExportFileResult) IsValid() bool {
	for _, existing := range allowedDiagnosticExportFileResultEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DiagnosticExportFileResult value.
func (v DiagnosticExportFileResult) Ptr() *DiagnosticExportFileResult {
	return &v
}
