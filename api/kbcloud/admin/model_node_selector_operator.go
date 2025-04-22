// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type NodeSelectorOperator string

// List of NodeSelectorOperator.
const (
	NodeSelectorOperatorIn           NodeSelectorOperator = "In"
	NodeSelectorOperatorNotIn        NodeSelectorOperator = "NotIn"
	NodeSelectorOperatorExists       NodeSelectorOperator = "Exists"
	NodeSelectorOperatorDoesNotExist NodeSelectorOperator = "DoesNotExist"
)

var allowedNodeSelectorOperatorEnumValues = []NodeSelectorOperator{
	NodeSelectorOperatorIn,
	NodeSelectorOperatorNotIn,
	NodeSelectorOperatorExists,
	NodeSelectorOperatorDoesNotExist,
}

// GetAllowedValues returns the list of possible values.
func (v *NodeSelectorOperator) GetAllowedValues() []NodeSelectorOperator {
	return allowedNodeSelectorOperatorEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *NodeSelectorOperator) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NodeSelectorOperator(value)
	return nil
}

// NewNodeSelectorOperatorFromValue returns a pointer to a valid NodeSelectorOperator
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewNodeSelectorOperatorFromValue(v string) (*NodeSelectorOperator, error) {
	ev := NodeSelectorOperator(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NodeSelectorOperator: valid values are %v", v, allowedNodeSelectorOperatorEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v NodeSelectorOperator) IsValid() bool {
	for _, existing := range allowedNodeSelectorOperatorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NodeSelectorOperator value.
func (v NodeSelectorOperator) Ptr() *NodeSelectorOperator {
	return &v
}
