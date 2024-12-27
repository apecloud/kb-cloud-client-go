// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// SortType Represents the type of sorting.
type SortType string

// List of SortType.
const (
	SortTypeAsc  SortType = "asc"
	SortTypeDesc SortType = "desc"
)

var allowedSortTypeEnumValues = []SortType{
	SortTypeAsc,
	SortTypeDesc,
}

// GetAllowedValues returns the list of possible values.
func (v *SortType) GetAllowedValues() []SortType {
	return allowedSortTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SortType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SortType(value)
	return nil
}

// NewSortTypeFromValue returns a pointer to a valid SortType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSortTypeFromValue(v string) (*SortType, error) {
	ev := SortType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SortType: valid values are %v", v, allowedSortTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SortType) IsValid() bool {
	for _, existing := range allowedSortTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SortType value.
func (v SortType) Ptr() *SortType {
	return &v
}
