// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// PostgresqlStorageInstanceRole PostgreSQL replica role normalized from metrics labels.
type PostgresqlStorageInstanceRole string

// List of PostgresqlStorageInstanceRole.
const (
	PostgresqlStorageInstanceRolePrimary PostgresqlStorageInstanceRole = "primary"
	PostgresqlStorageInstanceRoleStandby PostgresqlStorageInstanceRole = "standby"
	PostgresqlStorageInstanceRoleUnknown PostgresqlStorageInstanceRole = "unknown"
)

var allowedPostgresqlStorageInstanceRoleEnumValues = []PostgresqlStorageInstanceRole{
	PostgresqlStorageInstanceRolePrimary,
	PostgresqlStorageInstanceRoleStandby,
	PostgresqlStorageInstanceRoleUnknown,
}

// GetAllowedValues returns the list of possible values.
func (v *PostgresqlStorageInstanceRole) GetAllowedValues() []PostgresqlStorageInstanceRole {
	return allowedPostgresqlStorageInstanceRoleEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PostgresqlStorageInstanceRole) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PostgresqlStorageInstanceRole(value)
	return nil
}

// NewPostgresqlStorageInstanceRoleFromValue returns a pointer to a valid PostgresqlStorageInstanceRole
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPostgresqlStorageInstanceRoleFromValue(v string) (*PostgresqlStorageInstanceRole, error) {
	ev := PostgresqlStorageInstanceRole(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PostgresqlStorageInstanceRole: valid values are %v", v, allowedPostgresqlStorageInstanceRoleEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PostgresqlStorageInstanceRole) IsValid() bool {
	for _, existing := range allowedPostgresqlStorageInstanceRoleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PostgresqlStorageInstanceRole value.
func (v PostgresqlStorageInstanceRole) Ptr() *PostgresqlStorageInstanceRole {
	return &v
}
