// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type MssqlSQLOrder string

// List of MssqlSQLOrder.
const (
	MssqlSQLOrderTotalElapsedMs MssqlSQLOrder = "totalElapsedMs"
	MssqlSQLOrderMeanElapsedMs  MssqlSQLOrder = "meanElapsedMs"
	MssqlSQLOrderMaxElapsedMs   MssqlSQLOrder = "maxElapsedMs"
	MssqlSQLOrderTotalCpuMs     MssqlSQLOrder = "totalCpuMs"
	MssqlSQLOrderMeanCpuMs      MssqlSQLOrder = "meanCpuMs"
	MssqlSQLOrderLogicalReads   MssqlSQLOrder = "logicalReads"
	MssqlSQLOrderPhysicalReads  MssqlSQLOrder = "physicalReads"
	MssqlSQLOrderLogicalWrites  MssqlSQLOrder = "logicalWrites"
	MssqlSQLOrderExecutions     MssqlSQLOrder = "executions"
)

var allowedMssqlSQLOrderEnumValues = []MssqlSQLOrder{
	MssqlSQLOrderTotalElapsedMs,
	MssqlSQLOrderMeanElapsedMs,
	MssqlSQLOrderMaxElapsedMs,
	MssqlSQLOrderTotalCpuMs,
	MssqlSQLOrderMeanCpuMs,
	MssqlSQLOrderLogicalReads,
	MssqlSQLOrderPhysicalReads,
	MssqlSQLOrderLogicalWrites,
	MssqlSQLOrderExecutions,
}

// GetAllowedValues returns the list of possible values.
func (v *MssqlSQLOrder) GetAllowedValues() []MssqlSQLOrder {
	return allowedMssqlSQLOrderEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MssqlSQLOrder) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MssqlSQLOrder(value)
	return nil
}

// NewMssqlSQLOrderFromValue returns a pointer to a valid MssqlSQLOrder
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMssqlSQLOrderFromValue(v string) (*MssqlSQLOrder, error) {
	ev := MssqlSQLOrder(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MssqlSQLOrder: valid values are %v", v, allowedMssqlSQLOrderEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MssqlSQLOrder) IsValid() bool {
	for _, existing := range allowedMssqlSQLOrderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MssqlSQLOrder value.
func (v MssqlSQLOrder) Ptr() *MssqlSQLOrder {
	return &v
}
