// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type SqlAuditSwitchOptionType string

// List of SqlAuditSwitchOptionType.
const (
	SqlAuditSwitchOptionTypeSql       SqlAuditSwitchOptionType = "sql"
	SqlAuditSwitchOptionTypeParameter SqlAuditSwitchOptionType = "parameter"
)

var allowedSqlAuditSwitchOptionTypeEnumValues = []SqlAuditSwitchOptionType{
	SqlAuditSwitchOptionTypeSql,
	SqlAuditSwitchOptionTypeParameter,
}

// GetAllowedValues returns the list of possible values.
func (v *SqlAuditSwitchOptionType) GetAllowedValues() []SqlAuditSwitchOptionType {
	return allowedSqlAuditSwitchOptionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SqlAuditSwitchOptionType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SqlAuditSwitchOptionType(value)
	return nil
}

// NewSqlAuditSwitchOptionTypeFromValue returns a pointer to a valid SqlAuditSwitchOptionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSqlAuditSwitchOptionTypeFromValue(v string) (*SqlAuditSwitchOptionType, error) {
	ev := SqlAuditSwitchOptionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SqlAuditSwitchOptionType: valid values are %v", v, allowedSqlAuditSwitchOptionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SqlAuditSwitchOptionType) IsValid() bool {
	for _, existing := range allowedSqlAuditSwitchOptionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SqlAuditSwitchOptionType value.
func (v SqlAuditSwitchOptionType) Ptr() *SqlAuditSwitchOptionType {
	return &v
}
