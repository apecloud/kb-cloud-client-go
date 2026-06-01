// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlRealtimeSessionsReportStatus string

// List of PostgresqlRealtimeSessionsReportStatus.
const (
	PostgresqlRealtimeSessionsReportStatusWithData          PostgresqlRealtimeSessionsReportStatus = "with_data"
	PostgresqlRealtimeSessionsReportStatusNoData            PostgresqlRealtimeSessionsReportStatus = "no_data"
	PostgresqlRealtimeSessionsReportStatusUnsupportedEngine PostgresqlRealtimeSessionsReportStatus = "unsupported_engine"
	PostgresqlRealtimeSessionsReportStatusPermissionDenied  PostgresqlRealtimeSessionsReportStatus = "permission_denied"
	PostgresqlRealtimeSessionsReportStatusCollectionFailed  PostgresqlRealtimeSessionsReportStatus = "collection_failed"
)

var allowedPostgresqlRealtimeSessionsReportStatusEnumValues = []PostgresqlRealtimeSessionsReportStatus{
	PostgresqlRealtimeSessionsReportStatusWithData,
	PostgresqlRealtimeSessionsReportStatusNoData,
	PostgresqlRealtimeSessionsReportStatusUnsupportedEngine,
	PostgresqlRealtimeSessionsReportStatusPermissionDenied,
	PostgresqlRealtimeSessionsReportStatusCollectionFailed,
}

// GetAllowedValues returns the list of possible values.
func (v *PostgresqlRealtimeSessionsReportStatus) GetAllowedValues() []PostgresqlRealtimeSessionsReportStatus {
	return allowedPostgresqlRealtimeSessionsReportStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PostgresqlRealtimeSessionsReportStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PostgresqlRealtimeSessionsReportStatus(value)
	return nil
}

// NewPostgresqlRealtimeSessionsReportStatusFromValue returns a pointer to a valid PostgresqlRealtimeSessionsReportStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPostgresqlRealtimeSessionsReportStatusFromValue(v string) (*PostgresqlRealtimeSessionsReportStatus, error) {
	ev := PostgresqlRealtimeSessionsReportStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PostgresqlRealtimeSessionsReportStatus: valid values are %v", v, allowedPostgresqlRealtimeSessionsReportStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PostgresqlRealtimeSessionsReportStatus) IsValid() bool {
	for _, existing := range allowedPostgresqlRealtimeSessionsReportStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PostgresqlRealtimeSessionsReportStatus value.
func (v PostgresqlRealtimeSessionsReportStatus) Ptr() *PostgresqlRealtimeSessionsReportStatus {
	return &v
}
