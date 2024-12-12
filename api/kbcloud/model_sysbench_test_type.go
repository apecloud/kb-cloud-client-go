// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
)

// SysbenchTestType Test type for sysbench
type SysbenchTestType string

// List of SysbenchTestType.
const (
	SysbenchTestTypeOltpDelete         SysbenchTestType = "oltp_delete"
	SysbenchTestTypeOltpInsert         SysbenchTestType = "oltp_insert"
	SysbenchTestTypeOltpPointSelect    SysbenchTestType = "oltp_point_select"
	SysbenchTestTypeOltpReadOnly       SysbenchTestType = "oltp_read_only"
	SysbenchTestTypeOltpReadWrite      SysbenchTestType = "oltp_read_write"
	SysbenchTestTypeOltpUpdateIndex    SysbenchTestType = "oltp_update_index"
	SysbenchTestTypeOltpUpdateNonIndex SysbenchTestType = "oltp_update_non_index"
	SysbenchTestTypeOltpWriteOnly      SysbenchTestType = "oltp_write_only"
	SysbenchTestTypeSelectRandomPoints SysbenchTestType = "select_random_points"
	SysbenchTestTypeSelectRandomRanges SysbenchTestType = "select_random_ranges"
	SysbenchTestTypeBulkInsert         SysbenchTestType = "bulk_insert"
	SysbenchTestTypeOltpReadWritePct   SysbenchTestType = "oltp_read_write_pct"
)

var allowedSysbenchTestTypeEnumValues = []SysbenchTestType{
	SysbenchTestTypeOltpDelete,
	SysbenchTestTypeOltpInsert,
	SysbenchTestTypeOltpPointSelect,
	SysbenchTestTypeOltpReadOnly,
	SysbenchTestTypeOltpReadWrite,
	SysbenchTestTypeOltpUpdateIndex,
	SysbenchTestTypeOltpUpdateNonIndex,
	SysbenchTestTypeOltpWriteOnly,
	SysbenchTestTypeSelectRandomPoints,
	SysbenchTestTypeSelectRandomRanges,
	SysbenchTestTypeBulkInsert,
	SysbenchTestTypeOltpReadWritePct,
}

// GetAllowedValues returns the list of possible values.
func (v *SysbenchTestType) GetAllowedValues() []SysbenchTestType {
	return allowedSysbenchTestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SysbenchTestType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SysbenchTestType(value)
	return nil
}

// NewSysbenchTestTypeFromValue returns a pointer to a valid SysbenchTestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSysbenchTestTypeFromValue(v string) (*SysbenchTestType, error) {
	ev := SysbenchTestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SysbenchTestType: valid values are %v", v, allowedSysbenchTestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SysbenchTestType) IsValid() bool {
	for _, existing := range allowedSysbenchTestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SysbenchTestType value.
func (v SysbenchTestType) Ptr() *SysbenchTestType {
	return &v
}
