// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package kbcloud

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)



type ClassSeries string

// List of ClassSeries.
const (
	ClassSeriesGeneral ClassSeries = "general"
	ClassSeriesExclusive ClassSeries = "exclusive"
)

var allowedClassSeriesEnumValues = []ClassSeries{
	ClassSeriesGeneral,
	ClassSeriesExclusive,
}

// GetAllowedValues returns the list of possible values.
func (v *ClassSeries) GetAllowedValues() []ClassSeries {
	return allowedClassSeriesEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ClassSeries) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ClassSeries(value)
	return nil
}

// NewClassSeriesFromValue returns a pointer to a valid ClassSeries
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewClassSeriesFromValue(v string) (*ClassSeries, error) {
	ev := ClassSeries(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ClassSeries: valid values are %v", v, allowedClassSeriesEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ClassSeries) IsValid() bool {
	for _, existing := range allowedClassSeriesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ClassSeries value.
func (v ClassSeries) Ptr() *ClassSeries {
	return &v
}
