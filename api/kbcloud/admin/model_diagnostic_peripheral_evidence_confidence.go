// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticPeripheralEvidenceConfidence Confidence of the explanation based only on the collected external evidence.
type DiagnosticPeripheralEvidenceConfidence string

// List of DiagnosticPeripheralEvidenceConfidence.
const (
	DiagnosticPeripheralEvidenceConfidenceHigh   DiagnosticPeripheralEvidenceConfidence = "high"
	DiagnosticPeripheralEvidenceConfidenceMedium DiagnosticPeripheralEvidenceConfidence = "medium"
	DiagnosticPeripheralEvidenceConfidenceLow    DiagnosticPeripheralEvidenceConfidence = "low"
)

var allowedDiagnosticPeripheralEvidenceConfidenceEnumValues = []DiagnosticPeripheralEvidenceConfidence{
	DiagnosticPeripheralEvidenceConfidenceHigh,
	DiagnosticPeripheralEvidenceConfidenceMedium,
	DiagnosticPeripheralEvidenceConfidenceLow,
}

// GetAllowedValues returns the list of possible values.
func (v *DiagnosticPeripheralEvidenceConfidence) GetAllowedValues() []DiagnosticPeripheralEvidenceConfidence {
	return allowedDiagnosticPeripheralEvidenceConfidenceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DiagnosticPeripheralEvidenceConfidence) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DiagnosticPeripheralEvidenceConfidence(value)
	return nil
}

// NewDiagnosticPeripheralEvidenceConfidenceFromValue returns a pointer to a valid DiagnosticPeripheralEvidenceConfidence
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDiagnosticPeripheralEvidenceConfidenceFromValue(v string) (*DiagnosticPeripheralEvidenceConfidence, error) {
	ev := DiagnosticPeripheralEvidenceConfidence(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DiagnosticPeripheralEvidenceConfidence: valid values are %v", v, allowedDiagnosticPeripheralEvidenceConfidenceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DiagnosticPeripheralEvidenceConfidence) IsValid() bool {
	for _, existing := range allowedDiagnosticPeripheralEvidenceConfidenceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DiagnosticPeripheralEvidenceConfidence value.
func (v DiagnosticPeripheralEvidenceConfidence) Ptr() *DiagnosticPeripheralEvidenceConfidence {
	return &v
}
