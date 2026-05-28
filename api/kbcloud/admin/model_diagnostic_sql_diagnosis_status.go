// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticSQLDiagnosisStatus SQL diagnosis result status.
type DiagnosticSQLDiagnosisStatus string

// List of DiagnosticSQLDiagnosisStatus.
const (
	DiagnosticSQLDiagnosisStatusSucceeded DiagnosticSQLDiagnosisStatus = "succeeded"
	DiagnosticSQLDiagnosisStatusFailed    DiagnosticSQLDiagnosisStatus = "failed"
	DiagnosticSQLDiagnosisStatusRejected  DiagnosticSQLDiagnosisStatus = "rejected"
)

var allowedDiagnosticSQLDiagnosisStatusEnumValues = []DiagnosticSQLDiagnosisStatus{
	DiagnosticSQLDiagnosisStatusSucceeded,
	DiagnosticSQLDiagnosisStatusFailed,
	DiagnosticSQLDiagnosisStatusRejected,
}

// GetAllowedValues returns the list of possible values.
func (v *DiagnosticSQLDiagnosisStatus) GetAllowedValues() []DiagnosticSQLDiagnosisStatus {
	return allowedDiagnosticSQLDiagnosisStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DiagnosticSQLDiagnosisStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DiagnosticSQLDiagnosisStatus(value)
	return nil
}

// NewDiagnosticSQLDiagnosisStatusFromValue returns a pointer to a valid DiagnosticSQLDiagnosisStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDiagnosticSQLDiagnosisStatusFromValue(v string) (*DiagnosticSQLDiagnosisStatus, error) {
	ev := DiagnosticSQLDiagnosisStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DiagnosticSQLDiagnosisStatus: valid values are %v", v, allowedDiagnosticSQLDiagnosisStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DiagnosticSQLDiagnosisStatus) IsValid() bool {
	for _, existing := range allowedDiagnosticSQLDiagnosisStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DiagnosticSQLDiagnosisStatus value.
func (v DiagnosticSQLDiagnosisStatus) Ptr() *DiagnosticSQLDiagnosisStatus {
	return &v
}
