// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// InspectionTaskType Specifies the type of inspection task.
type InspectionTaskType string

// List of InspectionTaskType.
const (
	InspectionTaskTypeCluster InspectionTaskType = "cluster"
	InspectionTaskTypeNode    InspectionTaskType = "node"
)

var allowedInspectionTaskTypeEnumValues = []InspectionTaskType{
	InspectionTaskTypeCluster,
	InspectionTaskTypeNode,
}

// GetAllowedValues returns the list of possible values.
func (v *InspectionTaskType) GetAllowedValues() []InspectionTaskType {
	return allowedInspectionTaskTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *InspectionTaskType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = InspectionTaskType(value)
	return nil
}

// NewInspectionTaskTypeFromValue returns a pointer to a valid InspectionTaskType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewInspectionTaskTypeFromValue(v string) (*InspectionTaskType, error) {
	ev := InspectionTaskType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for InspectionTaskType: valid values are %v", v, allowedInspectionTaskTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v InspectionTaskType) IsValid() bool {
	for _, existing := range allowedInspectionTaskTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InspectionTaskType value.
func (v InspectionTaskType) Ptr() *InspectionTaskType {
	return &v
}
