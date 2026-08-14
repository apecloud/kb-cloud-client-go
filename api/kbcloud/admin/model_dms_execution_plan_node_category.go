// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DmsExecutionPlanNodeCategory string

// List of DmsExecutionPlanNodeCategory.
const (
	DmsExecutionPlanNodeCategoryResult      DmsExecutionPlanNodeCategory = "result"
	DmsExecutionPlanNodeCategorySort        DmsExecutionPlanNodeCategory = "sort"
	DmsExecutionPlanNodeCategoryJoin        DmsExecutionPlanNodeCategory = "join"
	DmsExecutionPlanNodeCategoryTableScan   DmsExecutionPlanNodeCategory = "table_scan"
	DmsExecutionPlanNodeCategoryIndexScan   DmsExecutionPlanNodeCategory = "index_scan"
	DmsExecutionPlanNodeCategoryLookup      DmsExecutionPlanNodeCategory = "lookup"
	DmsExecutionPlanNodeCategoryMaterialize DmsExecutionPlanNodeCategory = "materialize"
	DmsExecutionPlanNodeCategoryAggregate   DmsExecutionPlanNodeCategory = "aggregate"
	DmsExecutionPlanNodeCategoryExchange    DmsExecutionPlanNodeCategory = "exchange"
	DmsExecutionPlanNodeCategoryModify      DmsExecutionPlanNodeCategory = "modify"
	DmsExecutionPlanNodeCategoryOther       DmsExecutionPlanNodeCategory = "other"
)

var allowedDmsExecutionPlanNodeCategoryEnumValues = []DmsExecutionPlanNodeCategory{
	DmsExecutionPlanNodeCategoryResult,
	DmsExecutionPlanNodeCategorySort,
	DmsExecutionPlanNodeCategoryJoin,
	DmsExecutionPlanNodeCategoryTableScan,
	DmsExecutionPlanNodeCategoryIndexScan,
	DmsExecutionPlanNodeCategoryLookup,
	DmsExecutionPlanNodeCategoryMaterialize,
	DmsExecutionPlanNodeCategoryAggregate,
	DmsExecutionPlanNodeCategoryExchange,
	DmsExecutionPlanNodeCategoryModify,
	DmsExecutionPlanNodeCategoryOther,
}

// GetAllowedValues returns the list of possible values.
func (v *DmsExecutionPlanNodeCategory) GetAllowedValues() []DmsExecutionPlanNodeCategory {
	return allowedDmsExecutionPlanNodeCategoryEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DmsExecutionPlanNodeCategory) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DmsExecutionPlanNodeCategory(value)
	return nil
}

// NewDmsExecutionPlanNodeCategoryFromValue returns a pointer to a valid DmsExecutionPlanNodeCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDmsExecutionPlanNodeCategoryFromValue(v string) (*DmsExecutionPlanNodeCategory, error) {
	ev := DmsExecutionPlanNodeCategory(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DmsExecutionPlanNodeCategory: valid values are %v", v, allowedDmsExecutionPlanNodeCategoryEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DmsExecutionPlanNodeCategory) IsValid() bool {
	for _, existing := range allowedDmsExecutionPlanNodeCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DmsExecutionPlanNodeCategory value.
func (v DmsExecutionPlanNodeCategory) Ptr() *DmsExecutionPlanNodeCategory {
	return &v
}
