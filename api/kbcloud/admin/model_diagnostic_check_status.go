// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticCheckStatus Result status for one executed diagnostic check.
type DiagnosticCheckStatus string

// List of DiagnosticCheckStatus.
const (
	DiagnosticCheckStatusPass             DiagnosticCheckStatus = "pass"
	DiagnosticCheckStatusWarning          DiagnosticCheckStatus = "warning"
	DiagnosticCheckStatusFail             DiagnosticCheckStatus = "fail"
	DiagnosticCheckStatusCollectionFailed DiagnosticCheckStatus = "collection_failed"
	DiagnosticCheckStatusNotApplicable    DiagnosticCheckStatus = "not_applicable"
)

var allowedDiagnosticCheckStatusEnumValues = []DiagnosticCheckStatus{
	DiagnosticCheckStatusPass,
	DiagnosticCheckStatusWarning,
	DiagnosticCheckStatusFail,
	DiagnosticCheckStatusCollectionFailed,
	DiagnosticCheckStatusNotApplicable,
}

// GetAllowedValues returns the list of possible values.
func (v *DiagnosticCheckStatus) GetAllowedValues() []DiagnosticCheckStatus {
	return allowedDiagnosticCheckStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DiagnosticCheckStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DiagnosticCheckStatus(value)
	return nil
}

// NewDiagnosticCheckStatusFromValue returns a pointer to a valid DiagnosticCheckStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDiagnosticCheckStatusFromValue(v string) (*DiagnosticCheckStatus, error) {
	ev := DiagnosticCheckStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DiagnosticCheckStatus: valid values are %v", v, allowedDiagnosticCheckStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DiagnosticCheckStatus) IsValid() bool {
	for _, existing := range allowedDiagnosticCheckStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DiagnosticCheckStatus value.
func (v DiagnosticCheckStatus) Ptr() *DiagnosticCheckStatus {
	return &v
}
