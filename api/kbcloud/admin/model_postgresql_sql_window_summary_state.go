// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlSQLWindowSummaryState string

// List of PostgresqlSQLWindowSummaryState.
const (
	PostgresqlSQLWindowSummaryStateRedacted    PostgresqlSQLWindowSummaryState = "redacted"
	PostgresqlSQLWindowSummaryStateTruncated   PostgresqlSQLWindowSummaryState = "truncated"
	PostgresqlSQLWindowSummaryStateUnavailable PostgresqlSQLWindowSummaryState = "unavailable"
)

var allowedPostgresqlSQLWindowSummaryStateEnumValues = []PostgresqlSQLWindowSummaryState{
	PostgresqlSQLWindowSummaryStateRedacted,
	PostgresqlSQLWindowSummaryStateTruncated,
	PostgresqlSQLWindowSummaryStateUnavailable,
}

// GetAllowedValues returns the list of possible values.
func (v *PostgresqlSQLWindowSummaryState) GetAllowedValues() []PostgresqlSQLWindowSummaryState {
	return allowedPostgresqlSQLWindowSummaryStateEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PostgresqlSQLWindowSummaryState) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PostgresqlSQLWindowSummaryState(value)
	return nil
}

// NewPostgresqlSQLWindowSummaryStateFromValue returns a pointer to a valid PostgresqlSQLWindowSummaryState
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPostgresqlSQLWindowSummaryStateFromValue(v string) (*PostgresqlSQLWindowSummaryState, error) {
	ev := PostgresqlSQLWindowSummaryState(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PostgresqlSQLWindowSummaryState: valid values are %v", v, allowedPostgresqlSQLWindowSummaryStateEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PostgresqlSQLWindowSummaryState) IsValid() bool {
	for _, existing := range allowedPostgresqlSQLWindowSummaryStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PostgresqlSQLWindowSummaryState value.
func (v PostgresqlSQLWindowSummaryState) Ptr() *PostgresqlSQLWindowSummaryState {
	return &v
}
