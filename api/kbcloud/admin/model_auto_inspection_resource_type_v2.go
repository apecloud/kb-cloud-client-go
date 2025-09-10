// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// AutoInspectionResourceTypeV2 Specifies the type of the resource for the auto inspection.
type AutoInspectionResourceTypeV2 string

// List of AutoInspectionResourceTypeV2.
const (
	AutoInspectionResourceTypeV2Organization    AutoInspectionResourceTypeV2 = "organization"
	AutoInspectionResourceTypeV2Cluster         AutoInspectionResourceTypeV2 = "cluster"
	AutoInspectionResourceTypeV2Environment     AutoInspectionResourceTypeV2 = "environment"
	AutoInspectionResourceTypeV2EnvironmentNode AutoInspectionResourceTypeV2 = "environment_node"
)

var allowedAutoInspectionResourceTypeV2EnumValues = []AutoInspectionResourceTypeV2{
	AutoInspectionResourceTypeV2Organization,
	AutoInspectionResourceTypeV2Cluster,
	AutoInspectionResourceTypeV2Environment,
	AutoInspectionResourceTypeV2EnvironmentNode,
}

// GetAllowedValues returns the list of possible values.
func (v *AutoInspectionResourceTypeV2) GetAllowedValues() []AutoInspectionResourceTypeV2 {
	return allowedAutoInspectionResourceTypeV2EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AutoInspectionResourceTypeV2) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AutoInspectionResourceTypeV2(value)
	return nil
}

// NewAutoInspectionResourceTypeV2FromValue returns a pointer to a valid AutoInspectionResourceTypeV2
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAutoInspectionResourceTypeV2FromValue(v string) (*AutoInspectionResourceTypeV2, error) {
	ev := AutoInspectionResourceTypeV2(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AutoInspectionResourceTypeV2: valid values are %v", v, allowedAutoInspectionResourceTypeV2EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AutoInspectionResourceTypeV2) IsValid() bool {
	for _, existing := range allowedAutoInspectionResourceTypeV2EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AutoInspectionResourceTypeV2 value.
func (v AutoInspectionResourceTypeV2) Ptr() *AutoInspectionResourceTypeV2 {
	return &v
}
