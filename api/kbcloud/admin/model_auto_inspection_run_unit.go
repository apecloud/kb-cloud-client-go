// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// AutoInspectionRunUnit Specifies the unit of time for the auto inspection schedule.
type AutoInspectionRunUnit string

// List of AutoInspectionRunUnit.
const (
	AUTOINSPECTIONRUNUNIT_HOUR  AutoInspectionRunUnit = "hour"
	AUTOINSPECTIONRUNUNIT_DAY   AutoInspectionRunUnit = "day"
	AUTOINSPECTIONRUNUNIT_WEEK  AutoInspectionRunUnit = "week"
	AUTOINSPECTIONRUNUNIT_MONTH AutoInspectionRunUnit = "month"
)

var allowedAutoInspectionRunUnitEnumValues = []AutoInspectionRunUnit{
	AUTOINSPECTIONRUNUNIT_HOUR,
	AUTOINSPECTIONRUNUNIT_DAY,
	AUTOINSPECTIONRUNUNIT_WEEK,
	AUTOINSPECTIONRUNUNIT_MONTH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AutoInspectionRunUnit) GetAllowedValues() []AutoInspectionRunUnit {
	return allowedAutoInspectionRunUnitEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AutoInspectionRunUnit) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AutoInspectionRunUnit(value)
	return nil
}

// NewAutoInspectionRunUnitFromValue returns a pointer to a valid AutoInspectionRunUnit
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAutoInspectionRunUnitFromValue(v string) (*AutoInspectionRunUnit, error) {
	ev := AutoInspectionRunUnit(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AutoInspectionRunUnit: valid values are %v", v, allowedAutoInspectionRunUnitEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AutoInspectionRunUnit) IsValid() bool {
	for _, existing := range allowedAutoInspectionRunUnitEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AutoInspectionRunUnit value.
func (v AutoInspectionRunUnit) Ptr() *AutoInspectionRunUnit {
	return &v
}
