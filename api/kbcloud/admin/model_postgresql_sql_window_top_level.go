// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlSQLWindowTopLevel string

// List of PostgresqlSQLWindowTopLevel.
const (
	PostgresqlSQLWindowTopLevelTrue    PostgresqlSQLWindowTopLevel = "true"
	PostgresqlSQLWindowTopLevelFalse   PostgresqlSQLWindowTopLevel = "false"
	PostgresqlSQLWindowTopLevelUnknown PostgresqlSQLWindowTopLevel = "unknown"
)

var allowedPostgresqlSQLWindowTopLevelEnumValues = []PostgresqlSQLWindowTopLevel{
	PostgresqlSQLWindowTopLevelTrue,
	PostgresqlSQLWindowTopLevelFalse,
	PostgresqlSQLWindowTopLevelUnknown,
}

// GetAllowedValues returns the list of possible values.
func (v *PostgresqlSQLWindowTopLevel) GetAllowedValues() []PostgresqlSQLWindowTopLevel {
	return allowedPostgresqlSQLWindowTopLevelEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PostgresqlSQLWindowTopLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PostgresqlSQLWindowTopLevel(value)
	return nil
}

// NewPostgresqlSQLWindowTopLevelFromValue returns a pointer to a valid PostgresqlSQLWindowTopLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPostgresqlSQLWindowTopLevelFromValue(v string) (*PostgresqlSQLWindowTopLevel, error) {
	ev := PostgresqlSQLWindowTopLevel(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PostgresqlSQLWindowTopLevel: valid values are %v", v, allowedPostgresqlSQLWindowTopLevelEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PostgresqlSQLWindowTopLevel) IsValid() bool {
	for _, existing := range allowedPostgresqlSQLWindowTopLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PostgresqlSQLWindowTopLevel value.
func (v PostgresqlSQLWindowTopLevel) Ptr() *PostgresqlSQLWindowTopLevel {
	return &v
}
