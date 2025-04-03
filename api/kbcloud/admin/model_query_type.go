// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// QueryType 查询类型
type QueryType string

// List of QueryType.
const (
	QueryTypeAnalyze QueryType = "Analyze"
	QueryTypeBi      QueryType = "BI"
)

var allowedQueryTypeEnumValues = []QueryType{
	QueryTypeAnalyze,
	QueryTypeBi,
}

// GetAllowedValues returns the list of possible values.
func (v *QueryType) GetAllowedValues() []QueryType {
	return allowedQueryTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *QueryType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = QueryType(value)
	return nil
}

// NewQueryTypeFromValue returns a pointer to a valid QueryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewQueryTypeFromValue(v string) (*QueryType, error) {
	ev := QueryType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for QueryType: valid values are %v", v, allowedQueryTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v QueryType) IsValid() bool {
	for _, existing := range allowedQueryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to QueryType value.
func (v QueryType) Ptr() *QueryType {
	return &v
}
