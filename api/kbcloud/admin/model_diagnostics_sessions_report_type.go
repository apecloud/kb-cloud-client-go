// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DiagnosticsSessionsReportType string

// List of DiagnosticsSessionsReportType.
const (
	DiagnosticsSessionsReportTypePostgresqlSessions DiagnosticsSessionsReportType = "postgresql_sessions"
)

var allowedDiagnosticsSessionsReportTypeEnumValues = []DiagnosticsSessionsReportType{
	DiagnosticsSessionsReportTypePostgresqlSessions,
}

// GetAllowedValues returns the list of possible values.
func (v *DiagnosticsSessionsReportType) GetAllowedValues() []DiagnosticsSessionsReportType {
	return allowedDiagnosticsSessionsReportTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DiagnosticsSessionsReportType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DiagnosticsSessionsReportType(value)
	return nil
}

// NewDiagnosticsSessionsReportTypeFromValue returns a pointer to a valid DiagnosticsSessionsReportType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDiagnosticsSessionsReportTypeFromValue(v string) (*DiagnosticsSessionsReportType, error) {
	ev := DiagnosticsSessionsReportType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DiagnosticsSessionsReportType: valid values are %v", v, allowedDiagnosticsSessionsReportTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DiagnosticsSessionsReportType) IsValid() bool {
	for _, existing := range allowedDiagnosticsSessionsReportTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DiagnosticsSessionsReportType value.
func (v DiagnosticsSessionsReportType) Ptr() *DiagnosticsSessionsReportType {
	return &v
}
