// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ElasticsearchIndexSortBy string

// List of ElasticsearchIndexSortBy.
const (
	ElasticsearchIndexSortByName            ElasticsearchIndexSortBy = "name"
	ElasticsearchIndexSortByDocuments       ElasticsearchIndexSortBy = "documents"
	ElasticsearchIndexSortByTotalStoreBytes ElasticsearchIndexSortBy = "totalStoreBytes"
	ElasticsearchIndexSortBySegments        ElasticsearchIndexSortBy = "segments"
)

var allowedElasticsearchIndexSortByEnumValues = []ElasticsearchIndexSortBy{
	ElasticsearchIndexSortByName,
	ElasticsearchIndexSortByDocuments,
	ElasticsearchIndexSortByTotalStoreBytes,
	ElasticsearchIndexSortBySegments,
}

// GetAllowedValues returns the list of possible values.
func (v *ElasticsearchIndexSortBy) GetAllowedValues() []ElasticsearchIndexSortBy {
	return allowedElasticsearchIndexSortByEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ElasticsearchIndexSortBy) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ElasticsearchIndexSortBy(value)
	return nil
}

// NewElasticsearchIndexSortByFromValue returns a pointer to a valid ElasticsearchIndexSortBy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewElasticsearchIndexSortByFromValue(v string) (*ElasticsearchIndexSortBy, error) {
	ev := ElasticsearchIndexSortBy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ElasticsearchIndexSortBy: valid values are %v", v, allowedElasticsearchIndexSortByEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ElasticsearchIndexSortBy) IsValid() bool {
	for _, existing := range allowedElasticsearchIndexSortByEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ElasticsearchIndexSortBy value.
func (v ElasticsearchIndexSortBy) Ptr() *ElasticsearchIndexSortBy {
	return &v
}
