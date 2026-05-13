// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticTaskStatus Diagnostic task status.
type DiagnosticTaskStatus string

// List of DiagnosticTaskStatus.
const (
	DiagnosticTaskStatusPending   DiagnosticTaskStatus = "pending"
	DiagnosticTaskStatusRunning   DiagnosticTaskStatus = "running"
	DiagnosticTaskStatusSucceeded DiagnosticTaskStatus = "succeeded"
	DiagnosticTaskStatusFailed    DiagnosticTaskStatus = "failed"
	DiagnosticTaskStatusCanceled  DiagnosticTaskStatus = "canceled"
)

var allowedDiagnosticTaskStatusEnumValues = []DiagnosticTaskStatus{
	DiagnosticTaskStatusPending,
	DiagnosticTaskStatusRunning,
	DiagnosticTaskStatusSucceeded,
	DiagnosticTaskStatusFailed,
	DiagnosticTaskStatusCanceled,
}

// GetAllowedValues returns the list of possible values.
func (v *DiagnosticTaskStatus) GetAllowedValues() []DiagnosticTaskStatus {
	return allowedDiagnosticTaskStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DiagnosticTaskStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DiagnosticTaskStatus(value)
	return nil
}

// NewDiagnosticTaskStatusFromValue returns a pointer to a valid DiagnosticTaskStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDiagnosticTaskStatusFromValue(v string) (*DiagnosticTaskStatus, error) {
	ev := DiagnosticTaskStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DiagnosticTaskStatus: valid values are %v", v, allowedDiagnosticTaskStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DiagnosticTaskStatus) IsValid() bool {
	for _, existing := range allowedDiagnosticTaskStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DiagnosticTaskStatus value.
func (v DiagnosticTaskStatus) Ptr() *DiagnosticTaskStatus {
	return &v
}
