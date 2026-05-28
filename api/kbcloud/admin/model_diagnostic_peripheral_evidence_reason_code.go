// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticPeripheralEvidenceReasonCode Stable reason code for unavailable or degraded peripheral evidence.
type DiagnosticPeripheralEvidenceReasonCode string

// List of DiagnosticPeripheralEvidenceReasonCode.
const (
	DiagnosticPeripheralEvidenceReasonCodeNon                   DiagnosticPeripheralEvidenceReasonCode = "none"
	DiagnosticPeripheralEvidenceReasonCodeNoRelevantEvidence    DiagnosticPeripheralEvidenceReasonCode = "no_relevant_evidence"
	DiagnosticPeripheralEvidenceReasonCodePermissionDenied      DiagnosticPeripheralEvidenceReasonCode = "permission_denied"
	DiagnosticPeripheralEvidenceReasonCodeNoLogInWindow         DiagnosticPeripheralEvidenceReasonCode = "no_log_in_window"
	DiagnosticPeripheralEvidenceReasonCodeDatasourceUnavailable DiagnosticPeripheralEvidenceReasonCode = "datasource_unavailable"
	DiagnosticPeripheralEvidenceReasonCodeCollectionError       DiagnosticPeripheralEvidenceReasonCode = "collection_error"
)

var allowedDiagnosticPeripheralEvidenceReasonCodeEnumValues = []DiagnosticPeripheralEvidenceReasonCode{
	DiagnosticPeripheralEvidenceReasonCodeNon,
	DiagnosticPeripheralEvidenceReasonCodeNoRelevantEvidence,
	DiagnosticPeripheralEvidenceReasonCodePermissionDenied,
	DiagnosticPeripheralEvidenceReasonCodeNoLogInWindow,
	DiagnosticPeripheralEvidenceReasonCodeDatasourceUnavailable,
	DiagnosticPeripheralEvidenceReasonCodeCollectionError,
}

// GetAllowedValues returns the list of possible values.
func (v *DiagnosticPeripheralEvidenceReasonCode) GetAllowedValues() []DiagnosticPeripheralEvidenceReasonCode {
	return allowedDiagnosticPeripheralEvidenceReasonCodeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DiagnosticPeripheralEvidenceReasonCode) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DiagnosticPeripheralEvidenceReasonCode(value)
	return nil
}

// NewDiagnosticPeripheralEvidenceReasonCodeFromValue returns a pointer to a valid DiagnosticPeripheralEvidenceReasonCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDiagnosticPeripheralEvidenceReasonCodeFromValue(v string) (*DiagnosticPeripheralEvidenceReasonCode, error) {
	ev := DiagnosticPeripheralEvidenceReasonCode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DiagnosticPeripheralEvidenceReasonCode: valid values are %v", v, allowedDiagnosticPeripheralEvidenceReasonCodeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DiagnosticPeripheralEvidenceReasonCode) IsValid() bool {
	for _, existing := range allowedDiagnosticPeripheralEvidenceReasonCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DiagnosticPeripheralEvidenceReasonCode value.
func (v DiagnosticPeripheralEvidenceReasonCode) Ptr() *DiagnosticPeripheralEvidenceReasonCode {
	return &v
}
