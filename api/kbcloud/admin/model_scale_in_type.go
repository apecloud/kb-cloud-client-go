// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ScaleInType Type of scale in strategy
type ScaleInType string

// List of ScaleInType.
const (
	ScaleInTypeParallel ScaleInType = "Parallel"
	ScaleInTypeSerial   ScaleInType = "Serial"
)

var allowedScaleInTypeEnumValues = []ScaleInType{
	ScaleInTypeParallel,
	ScaleInTypeSerial,
}

// GetAllowedValues returns the list of possible values.
func (v *ScaleInType) GetAllowedValues() []ScaleInType {
	return allowedScaleInTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScaleInType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScaleInType(value)
	return nil
}

// NewScaleInTypeFromValue returns a pointer to a valid ScaleInType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScaleInTypeFromValue(v string) (*ScaleInType, error) {
	ev := ScaleInType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScaleInType: valid values are %v", v, allowedScaleInTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScaleInType) IsValid() bool {
	for _, existing := range allowedScaleInTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScaleInType value.
func (v ScaleInType) Ptr() *ScaleInType {
	return &v
}
