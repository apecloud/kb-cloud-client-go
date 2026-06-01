// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlRealtimeSessionsReportType string

// List of PostgresqlRealtimeSessionsReportType.
const (
	PostgresqlRealtimeSessionsReportTypePostgresqlRealtimeSessions PostgresqlRealtimeSessionsReportType = "postgresql_realtime_sessions"
)

var allowedPostgresqlRealtimeSessionsReportTypeEnumValues = []PostgresqlRealtimeSessionsReportType{
	PostgresqlRealtimeSessionsReportTypePostgresqlRealtimeSessions,
}

// GetAllowedValues returns the list of possible values.
func (v *PostgresqlRealtimeSessionsReportType) GetAllowedValues() []PostgresqlRealtimeSessionsReportType {
	return allowedPostgresqlRealtimeSessionsReportTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PostgresqlRealtimeSessionsReportType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PostgresqlRealtimeSessionsReportType(value)
	return nil
}

// NewPostgresqlRealtimeSessionsReportTypeFromValue returns a pointer to a valid PostgresqlRealtimeSessionsReportType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPostgresqlRealtimeSessionsReportTypeFromValue(v string) (*PostgresqlRealtimeSessionsReportType, error) {
	ev := PostgresqlRealtimeSessionsReportType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PostgresqlRealtimeSessionsReportType: valid values are %v", v, allowedPostgresqlRealtimeSessionsReportTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PostgresqlRealtimeSessionsReportType) IsValid() bool {
	for _, existing := range allowedPostgresqlRealtimeSessionsReportTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PostgresqlRealtimeSessionsReportType value.
func (v PostgresqlRealtimeSessionsReportType) Ptr() *PostgresqlRealtimeSessionsReportType {
	return &v
}
