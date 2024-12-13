// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
)

// NodeOperationType operation for node, either `add` or `del`
type NodeOperationType string

// List of NodeOperationType.
const (
	NodeOperationTypeAdd NodeOperationType = "add"
	NodeOperationTypeDel NodeOperationType = "del"
)

var allowedNodeOperationTypeEnumValues = []NodeOperationType{
	NodeOperationTypeAdd,
	NodeOperationTypeDel,
}

// GetAllowedValues returns the list of possible values.
func (v *NodeOperationType) GetAllowedValues() []NodeOperationType {
	return allowedNodeOperationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *NodeOperationType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NodeOperationType(value)
	return nil
}

// NewNodeOperationTypeFromValue returns a pointer to a valid NodeOperationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewNodeOperationTypeFromValue(v string) (*NodeOperationType, error) {
	ev := NodeOperationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NodeOperationType: valid values are %v", v, allowedNodeOperationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v NodeOperationType) IsValid() bool {
	for _, existing := range allowedNodeOperationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NodeOperationType value.
func (v NodeOperationType) Ptr() *NodeOperationType {
	return &v
}
