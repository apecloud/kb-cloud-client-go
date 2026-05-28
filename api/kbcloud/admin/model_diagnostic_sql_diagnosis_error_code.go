// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticSQLDiagnosisErrorCode Stable error code for SQL diagnosis failures and safety rejections.
type DiagnosticSQLDiagnosisErrorCode string

// List of DiagnosticSQLDiagnosisErrorCode.
const (
	DiagnosticSQLDiagnosisErrorCodeEmptySql                    DiagnosticSQLDiagnosisErrorCode = "empty_sql"
	DiagnosticSQLDiagnosisErrorCodeNonSelectRejected           DiagnosticSQLDiagnosisErrorCode = "non_select_rejected"
	DiagnosticSQLDiagnosisErrorCodeParameterizedSqlNeedsValues DiagnosticSQLDiagnosisErrorCode = "parameterized_sql_needs_values"
	DiagnosticSQLDiagnosisErrorCodeMissingParameterValues      DiagnosticSQLDiagnosisErrorCode = "missing_parameter_values"
	DiagnosticSQLDiagnosisErrorCodeParameterCountMismatch      DiagnosticSQLDiagnosisErrorCode = "parameter_count_mismatch"
	DiagnosticSQLDiagnosisErrorCodePermissionDenied            DiagnosticSQLDiagnosisErrorCode = "permission_denied"
	DiagnosticSQLDiagnosisErrorCodeTimeout                     DiagnosticSQLDiagnosisErrorCode = "timeout"
	DiagnosticSQLDiagnosisErrorCodeParseFailed                 DiagnosticSQLDiagnosisErrorCode = "parse_failed"
	DiagnosticSQLDiagnosisErrorCodeExplainFailed               DiagnosticSQLDiagnosisErrorCode = "explain_failed"
)

var allowedDiagnosticSQLDiagnosisErrorCodeEnumValues = []DiagnosticSQLDiagnosisErrorCode{
	DiagnosticSQLDiagnosisErrorCodeEmptySql,
	DiagnosticSQLDiagnosisErrorCodeNonSelectRejected,
	DiagnosticSQLDiagnosisErrorCodeParameterizedSqlNeedsValues,
	DiagnosticSQLDiagnosisErrorCodeMissingParameterValues,
	DiagnosticSQLDiagnosisErrorCodeParameterCountMismatch,
	DiagnosticSQLDiagnosisErrorCodePermissionDenied,
	DiagnosticSQLDiagnosisErrorCodeTimeout,
	DiagnosticSQLDiagnosisErrorCodeParseFailed,
	DiagnosticSQLDiagnosisErrorCodeExplainFailed,
}

// GetAllowedValues returns the list of possible values.
func (v *DiagnosticSQLDiagnosisErrorCode) GetAllowedValues() []DiagnosticSQLDiagnosisErrorCode {
	return allowedDiagnosticSQLDiagnosisErrorCodeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DiagnosticSQLDiagnosisErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DiagnosticSQLDiagnosisErrorCode(value)
	return nil
}

// NewDiagnosticSQLDiagnosisErrorCodeFromValue returns a pointer to a valid DiagnosticSQLDiagnosisErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDiagnosticSQLDiagnosisErrorCodeFromValue(v string) (*DiagnosticSQLDiagnosisErrorCode, error) {
	ev := DiagnosticSQLDiagnosisErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DiagnosticSQLDiagnosisErrorCode: valid values are %v", v, allowedDiagnosticSQLDiagnosisErrorCodeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DiagnosticSQLDiagnosisErrorCode) IsValid() bool {
	for _, existing := range allowedDiagnosticSQLDiagnosisErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DiagnosticSQLDiagnosisErrorCode value.
func (v DiagnosticSQLDiagnosisErrorCode) Ptr() *DiagnosticSQLDiagnosisErrorCode {
	return &v
}
