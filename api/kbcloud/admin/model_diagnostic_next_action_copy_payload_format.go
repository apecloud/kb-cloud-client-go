// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticNextActionCopyPayloadFormat Payload format. Sprint 19 only uses plain_text.
type DiagnosticNextActionCopyPayloadFormat string

// List of DiagnosticNextActionCopyPayloadFormat.
const (
	DiagnosticNextActionCopyPayloadFormatPlainText DiagnosticNextActionCopyPayloadFormat = "plain_text"
)

var allowedDiagnosticNextActionCopyPayloadFormatEnumValues = []DiagnosticNextActionCopyPayloadFormat{
	DiagnosticNextActionCopyPayloadFormatPlainText,
}

// GetAllowedValues returns the list of possible values.
func (v *DiagnosticNextActionCopyPayloadFormat) GetAllowedValues() []DiagnosticNextActionCopyPayloadFormat {
	return allowedDiagnosticNextActionCopyPayloadFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DiagnosticNextActionCopyPayloadFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DiagnosticNextActionCopyPayloadFormat(value)
	return nil
}

// NewDiagnosticNextActionCopyPayloadFormatFromValue returns a pointer to a valid DiagnosticNextActionCopyPayloadFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDiagnosticNextActionCopyPayloadFormatFromValue(v string) (*DiagnosticNextActionCopyPayloadFormat, error) {
	ev := DiagnosticNextActionCopyPayloadFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DiagnosticNextActionCopyPayloadFormat: valid values are %v", v, allowedDiagnosticNextActionCopyPayloadFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DiagnosticNextActionCopyPayloadFormat) IsValid() bool {
	for _, existing := range allowedDiagnosticNextActionCopyPayloadFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DiagnosticNextActionCopyPayloadFormat value.
func (v DiagnosticNextActionCopyPayloadFormat) Ptr() *DiagnosticNextActionCopyPayloadFormat {
	return &v
}
