// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ElasticsearchSortOrder string

// List of ElasticsearchSortOrder.
const (
	ElasticsearchSortOrderAsc  ElasticsearchSortOrder = "asc"
	ElasticsearchSortOrderDesc ElasticsearchSortOrder = "desc"
)

var allowedElasticsearchSortOrderEnumValues = []ElasticsearchSortOrder{
	ElasticsearchSortOrderAsc,
	ElasticsearchSortOrderDesc,
}

// GetAllowedValues returns the list of possible values.
func (v *ElasticsearchSortOrder) GetAllowedValues() []ElasticsearchSortOrder {
	return allowedElasticsearchSortOrderEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ElasticsearchSortOrder) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ElasticsearchSortOrder(value)
	return nil
}

// NewElasticsearchSortOrderFromValue returns a pointer to a valid ElasticsearchSortOrder
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewElasticsearchSortOrderFromValue(v string) (*ElasticsearchSortOrder, error) {
	ev := ElasticsearchSortOrder(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ElasticsearchSortOrder: valid values are %v", v, allowedElasticsearchSortOrderEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ElasticsearchSortOrder) IsValid() bool {
	for _, existing := range allowedElasticsearchSortOrderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ElasticsearchSortOrder value.
func (v ElasticsearchSortOrder) Ptr() *ElasticsearchSortOrder {
	return &v
}
