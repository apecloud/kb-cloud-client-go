// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlSQLWindowInstanceStatus string

// List of PostgresqlSQLWindowInstanceStatus.
const (
	PostgresqlSQLWindowInstanceStatusAvailable   PostgresqlSQLWindowInstanceStatus = "available"
	PostgresqlSQLWindowInstanceStatusUnavailable PostgresqlSQLWindowInstanceStatus = "unavailable"
)

var allowedPostgresqlSQLWindowInstanceStatusEnumValues = []PostgresqlSQLWindowInstanceStatus{
	PostgresqlSQLWindowInstanceStatusAvailable,
	PostgresqlSQLWindowInstanceStatusUnavailable,
}

// GetAllowedValues returns the list of possible values.
func (v *PostgresqlSQLWindowInstanceStatus) GetAllowedValues() []PostgresqlSQLWindowInstanceStatus {
	return allowedPostgresqlSQLWindowInstanceStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PostgresqlSQLWindowInstanceStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PostgresqlSQLWindowInstanceStatus(value)
	return nil
}

// NewPostgresqlSQLWindowInstanceStatusFromValue returns a pointer to a valid PostgresqlSQLWindowInstanceStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPostgresqlSQLWindowInstanceStatusFromValue(v string) (*PostgresqlSQLWindowInstanceStatus, error) {
	ev := PostgresqlSQLWindowInstanceStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PostgresqlSQLWindowInstanceStatus: valid values are %v", v, allowedPostgresqlSQLWindowInstanceStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PostgresqlSQLWindowInstanceStatus) IsValid() bool {
	for _, existing := range allowedPostgresqlSQLWindowInstanceStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PostgresqlSQLWindowInstanceStatus value.
func (v PostgresqlSQLWindowInstanceStatus) Ptr() *PostgresqlSQLWindowInstanceStatus {
	return &v
}
