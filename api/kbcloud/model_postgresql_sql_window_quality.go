// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlSQLWindowQuality string

// List of PostgresqlSQLWindowQuality.
const (
	PostgresqlSQLWindowQualityComplete PostgresqlSQLWindowQuality = "complete"
	PostgresqlSQLWindowQualityPartial  PostgresqlSQLWindowQuality = "partial"
	PostgresqlSQLWindowQualityEmpty    PostgresqlSQLWindowQuality = "empty"
)

var allowedPostgresqlSQLWindowQualityEnumValues = []PostgresqlSQLWindowQuality{
	PostgresqlSQLWindowQualityComplete,
	PostgresqlSQLWindowQualityPartial,
	PostgresqlSQLWindowQualityEmpty,
}

// GetAllowedValues returns the list of possible values.
func (v *PostgresqlSQLWindowQuality) GetAllowedValues() []PostgresqlSQLWindowQuality {
	return allowedPostgresqlSQLWindowQualityEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PostgresqlSQLWindowQuality) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PostgresqlSQLWindowQuality(value)
	return nil
}

// NewPostgresqlSQLWindowQualityFromValue returns a pointer to a valid PostgresqlSQLWindowQuality
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPostgresqlSQLWindowQualityFromValue(v string) (*PostgresqlSQLWindowQuality, error) {
	ev := PostgresqlSQLWindowQuality(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PostgresqlSQLWindowQuality: valid values are %v", v, allowedPostgresqlSQLWindowQualityEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PostgresqlSQLWindowQuality) IsValid() bool {
	for _, existing := range allowedPostgresqlSQLWindowQualityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PostgresqlSQLWindowQuality value.
func (v PostgresqlSQLWindowQuality) Ptr() *PostgresqlSQLWindowQuality {
	return &v
}
