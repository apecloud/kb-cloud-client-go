// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticSQLDiagnosisParameterType Minimal value type used to render a PostgreSQL literal safely for EXPLAIN.
type DiagnosticSQLDiagnosisParameterType string

// List of DiagnosticSQLDiagnosisParameterType.
const (
	DiagnosticSQLDiagnosisParameterTypeString  DiagnosticSQLDiagnosisParameterType = "string"
	DiagnosticSQLDiagnosisParameterTypeNumber  DiagnosticSQLDiagnosisParameterType = "number"
	DiagnosticSQLDiagnosisParameterTypeBoolean DiagnosticSQLDiagnosisParameterType = "boolean"
	DiagnosticSQLDiagnosisParameterTypeNull    DiagnosticSQLDiagnosisParameterType = "null"
)

var allowedDiagnosticSQLDiagnosisParameterTypeEnumValues = []DiagnosticSQLDiagnosisParameterType{
	DiagnosticSQLDiagnosisParameterTypeString,
	DiagnosticSQLDiagnosisParameterTypeNumber,
	DiagnosticSQLDiagnosisParameterTypeBoolean,
	DiagnosticSQLDiagnosisParameterTypeNull,
}

// GetAllowedValues returns the list of possible values.
func (v *DiagnosticSQLDiagnosisParameterType) GetAllowedValues() []DiagnosticSQLDiagnosisParameterType {
	return allowedDiagnosticSQLDiagnosisParameterTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DiagnosticSQLDiagnosisParameterType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DiagnosticSQLDiagnosisParameterType(value)
	return nil
}

// NewDiagnosticSQLDiagnosisParameterTypeFromValue returns a pointer to a valid DiagnosticSQLDiagnosisParameterType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDiagnosticSQLDiagnosisParameterTypeFromValue(v string) (*DiagnosticSQLDiagnosisParameterType, error) {
	ev := DiagnosticSQLDiagnosisParameterType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DiagnosticSQLDiagnosisParameterType: valid values are %v", v, allowedDiagnosticSQLDiagnosisParameterTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DiagnosticSQLDiagnosisParameterType) IsValid() bool {
	for _, existing := range allowedDiagnosticSQLDiagnosisParameterTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DiagnosticSQLDiagnosisParameterType value.
func (v DiagnosticSQLDiagnosisParameterType) Ptr() *DiagnosticSQLDiagnosisParameterType {
	return &v
}
