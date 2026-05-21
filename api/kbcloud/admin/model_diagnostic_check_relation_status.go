// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticCheckRelationStatus Whether this check has proven relation to the current database object.
type DiagnosticCheckRelationStatus string

// List of DiagnosticCheckRelationStatus.
const (
	DiagnosticCheckRelationStatusRelated   DiagnosticCheckRelationStatus = "related"
	DiagnosticCheckRelationStatusUnknown   DiagnosticCheckRelationStatus = "unknown"
	DiagnosticCheckRelationStatusUnrelated DiagnosticCheckRelationStatus = "unrelated"
)

var allowedDiagnosticCheckRelationStatusEnumValues = []DiagnosticCheckRelationStatus{
	DiagnosticCheckRelationStatusRelated,
	DiagnosticCheckRelationStatusUnknown,
	DiagnosticCheckRelationStatusUnrelated,
}

// GetAllowedValues returns the list of possible values.
func (v *DiagnosticCheckRelationStatus) GetAllowedValues() []DiagnosticCheckRelationStatus {
	return allowedDiagnosticCheckRelationStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DiagnosticCheckRelationStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DiagnosticCheckRelationStatus(value)
	return nil
}

// NewDiagnosticCheckRelationStatusFromValue returns a pointer to a valid DiagnosticCheckRelationStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDiagnosticCheckRelationStatusFromValue(v string) (*DiagnosticCheckRelationStatus, error) {
	ev := DiagnosticCheckRelationStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DiagnosticCheckRelationStatus: valid values are %v", v, allowedDiagnosticCheckRelationStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DiagnosticCheckRelationStatus) IsValid() bool {
	for _, existing := range allowedDiagnosticCheckRelationStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DiagnosticCheckRelationStatus value.
func (v DiagnosticCheckRelationStatus) Ptr() *DiagnosticCheckRelationStatus {
	return &v
}
