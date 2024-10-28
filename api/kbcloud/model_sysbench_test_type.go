// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// SysbenchTestType Test type for sysbench
type SysbenchTestType string

// List of SysbenchTestType.
const (
	SYSBENCHTESTTYPE_OLTP_DELETE           SysbenchTestType = "oltp_delete"
	SYSBENCHTESTTYPE_OLTP_INSERT           SysbenchTestType = "oltp_insert"
	SYSBENCHTESTTYPE_OLTP_POINT_SELECT     SysbenchTestType = "oltp_point_select"
	SYSBENCHTESTTYPE_OLTP_READ_ONLY        SysbenchTestType = "oltp_read_only"
	SYSBENCHTESTTYPE_OLTP_READ_WRITE       SysbenchTestType = "oltp_read_write"
	SYSBENCHTESTTYPE_OLTP_UPDATE_INDEX     SysbenchTestType = "oltp_update_index"
	SYSBENCHTESTTYPE_OLTP_UPDATE_NON_INDEX SysbenchTestType = "oltp_update_non_index"
	SYSBENCHTESTTYPE_OLTP_WRITE_ONLY       SysbenchTestType = "oltp_write_only"
	SYSBENCHTESTTYPE_SELECT_RANDOM_POINTS  SysbenchTestType = "select_random_points"
	SYSBENCHTESTTYPE_SELECT_RANDOM_RANGES  SysbenchTestType = "select_random_ranges"
	SYSBENCHTESTTYPE_BULK_INSERT           SysbenchTestType = "bulk_insert"
	SYSBENCHTESTTYPE_OLTP_READ_WRITE_PCT   SysbenchTestType = "oltp_read_write_pct"
)

var allowedSysbenchTestTypeEnumValues = []SysbenchTestType{
	SYSBENCHTESTTYPE_OLTP_DELETE,
	SYSBENCHTESTTYPE_OLTP_INSERT,
	SYSBENCHTESTTYPE_OLTP_POINT_SELECT,
	SYSBENCHTESTTYPE_OLTP_READ_ONLY,
	SYSBENCHTESTTYPE_OLTP_READ_WRITE,
	SYSBENCHTESTTYPE_OLTP_UPDATE_INDEX,
	SYSBENCHTESTTYPE_OLTP_UPDATE_NON_INDEX,
	SYSBENCHTESTTYPE_OLTP_WRITE_ONLY,
	SYSBENCHTESTTYPE_SELECT_RANDOM_POINTS,
	SYSBENCHTESTTYPE_SELECT_RANDOM_RANGES,
	SYSBENCHTESTTYPE_BULK_INSERT,
	SYSBENCHTESTTYPE_OLTP_READ_WRITE_PCT,
}

// GetAllowedValues reeturns the list of possible values.
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
