// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// AutoInspectionResourceType Specifies the type of the resource for the auto inspection.
type AutoInspectionResourceType string

// List of AutoInspectionResourceType.
const (
	AutoInspectionResourceTypeOrganization    AutoInspectionResourceType = "organization"
	AutoInspectionResourceTypeCluster         AutoInspectionResourceType = "cluster"
	AutoInspectionResourceTypeEnvironment     AutoInspectionResourceType = "environment"
	AutoInspectionResourceTypeEnvironmentNode AutoInspectionResourceType = "environment_node"
)

var allowedAutoInspectionResourceTypeEnumValues = []AutoInspectionResourceType{
	AutoInspectionResourceTypeOrganization,
	AutoInspectionResourceTypeCluster,
	AutoInspectionResourceTypeEnvironment,
	AutoInspectionResourceTypeEnvironmentNode,
}

// GetAllowedValues returns the list of possible values.
func (v *AutoInspectionResourceType) GetAllowedValues() []AutoInspectionResourceType {
	return allowedAutoInspectionResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AutoInspectionResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AutoInspectionResourceType(value)
	return nil
}

// NewAutoInspectionResourceTypeFromValue returns a pointer to a valid AutoInspectionResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAutoInspectionResourceTypeFromValue(v string) (*AutoInspectionResourceType, error) {
	ev := AutoInspectionResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AutoInspectionResourceType: valid values are %v", v, allowedAutoInspectionResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AutoInspectionResourceType) IsValid() bool {
	for _, existing := range allowedAutoInspectionResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AutoInspectionResourceType value.
func (v AutoInspectionResourceType) Ptr() *AutoInspectionResourceType {
	return &v
}
