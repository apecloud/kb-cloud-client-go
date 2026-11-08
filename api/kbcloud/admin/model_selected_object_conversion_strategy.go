// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// SelectedObjectConversionStrategy Strategy for converting the selected tree to parameters. levelMapping maps every configured level and applies skipWhenParentLevelsSelectedAll. wholeObject merges duplicate branches and maps only whole selected nodes (selectedAll or leaves), omitting covered descendants. Defaults to levelMapping when omitted.
type SelectedObjectConversionStrategy string

// List of SelectedObjectConversionStrategy.
const (
	SelectedObjectConversionStrategyLevelMapping SelectedObjectConversionStrategy = "levelMapping"
	SelectedObjectConversionStrategyWholeObject  SelectedObjectConversionStrategy = "wholeObject"
)

var allowedSelectedObjectConversionStrategyEnumValues = []SelectedObjectConversionStrategy{
	SelectedObjectConversionStrategyLevelMapping,
	SelectedObjectConversionStrategyWholeObject,
}

// GetAllowedValues returns the list of possible values.
func (v *SelectedObjectConversionStrategy) GetAllowedValues() []SelectedObjectConversionStrategy {
	return allowedSelectedObjectConversionStrategyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SelectedObjectConversionStrategy) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SelectedObjectConversionStrategy(value)
	return nil
}

// NewSelectedObjectConversionStrategyFromValue returns a pointer to a valid SelectedObjectConversionStrategy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSelectedObjectConversionStrategyFromValue(v string) (*SelectedObjectConversionStrategy, error) {
	ev := SelectedObjectConversionStrategy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SelectedObjectConversionStrategy: valid values are %v", v, allowedSelectedObjectConversionStrategyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SelectedObjectConversionStrategy) IsValid() bool {
	for _, existing := range allowedSelectedObjectConversionStrategyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SelectedObjectConversionStrategy value.
func (v SelectedObjectConversionStrategy) Ptr() *SelectedObjectConversionStrategy {
	return &v
}
