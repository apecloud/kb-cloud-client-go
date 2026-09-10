// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ElasticsearchHotThreadsType string

// List of ElasticsearchHotThreadsType.
const (
	ElasticsearchHotThreadsTypeCpu   ElasticsearchHotThreadsType = "cpu"
	ElasticsearchHotThreadsTypeWait  ElasticsearchHotThreadsType = "wait"
	ElasticsearchHotThreadsTypeBlock ElasticsearchHotThreadsType = "block"
)

var allowedElasticsearchHotThreadsTypeEnumValues = []ElasticsearchHotThreadsType{
	ElasticsearchHotThreadsTypeCpu,
	ElasticsearchHotThreadsTypeWait,
	ElasticsearchHotThreadsTypeBlock,
}

// GetAllowedValues returns the list of possible values.
func (v *ElasticsearchHotThreadsType) GetAllowedValues() []ElasticsearchHotThreadsType {
	return allowedElasticsearchHotThreadsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ElasticsearchHotThreadsType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ElasticsearchHotThreadsType(value)
	return nil
}

// NewElasticsearchHotThreadsTypeFromValue returns a pointer to a valid ElasticsearchHotThreadsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewElasticsearchHotThreadsTypeFromValue(v string) (*ElasticsearchHotThreadsType, error) {
	ev := ElasticsearchHotThreadsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ElasticsearchHotThreadsType: valid values are %v", v, allowedElasticsearchHotThreadsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ElasticsearchHotThreadsType) IsValid() bool {
	for _, existing := range allowedElasticsearchHotThreadsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ElasticsearchHotThreadsType value.
func (v ElasticsearchHotThreadsType) Ptr() *ElasticsearchHotThreadsType {
	return &v
}
