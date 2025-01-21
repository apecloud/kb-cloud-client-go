// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// TimeOrderBy Specifies the order of the time range for the query.
type TimeOrderBy string

// List of TimeOrderBy.
const (
	TimeOrderByStart TimeOrderBy = "start"
	TimeOrderByEnd   TimeOrderBy = "end"
)

var allowedTimeOrderByEnumValues = []TimeOrderBy{
	TimeOrderByStart,
	TimeOrderByEnd,
}

// GetAllowedValues returns the list of possible values.
func (v *TimeOrderBy) GetAllowedValues() []TimeOrderBy {
	return allowedTimeOrderByEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TimeOrderBy) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TimeOrderBy(value)
	return nil
}

// NewTimeOrderByFromValue returns a pointer to a valid TimeOrderBy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTimeOrderByFromValue(v string) (*TimeOrderBy, error) {
	ev := TimeOrderBy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TimeOrderBy: valid values are %v", v, allowedTimeOrderByEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TimeOrderBy) IsValid() bool {
	for _, existing := range allowedTimeOrderByEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TimeOrderBy value.
func (v TimeOrderBy) Ptr() *TimeOrderBy {
	return &v
}
