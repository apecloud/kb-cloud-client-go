// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// AutoInspectionLevel Specifies the level of the auto inspection.
type AutoInspectionLevel string

// List of AutoInspectionLevel.
const (
	AutoInspectionLevelOrg     AutoInspectionLevel = "org"
	AutoInspectionLevelCluster AutoInspectionLevel = "cluster"
)

var allowedAutoInspectionLevelEnumValues = []AutoInspectionLevel{
	AutoInspectionLevelOrg,
	AutoInspectionLevelCluster,
}

// GetAllowedValues returns the list of possible values.
func (v *AutoInspectionLevel) GetAllowedValues() []AutoInspectionLevel {
	return allowedAutoInspectionLevelEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AutoInspectionLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AutoInspectionLevel(value)
	return nil
}

// NewAutoInspectionLevelFromValue returns a pointer to a valid AutoInspectionLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAutoInspectionLevelFromValue(v string) (*AutoInspectionLevel, error) {
	ev := AutoInspectionLevel(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AutoInspectionLevel: valid values are %v", v, allowedAutoInspectionLevelEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AutoInspectionLevel) IsValid() bool {
	for _, existing := range allowedAutoInspectionLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AutoInspectionLevel value.
func (v AutoInspectionLevel) Ptr() *AutoInspectionLevel {
	return &v
}
