// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// NodeStatus Status of the node
type NodeStatus string

// List of NodeStatus.
const (
	NodeStatusReady            NodeStatus = "Ready"
	NodeStatusWarning          NodeStatus = "Warning"
	NodeStatusUnderMaintenance NodeStatus = "UnderMaintenance"
	NodeStatusUnschedulable    NodeStatus = "Unschedulable"
)

var allowedNodeStatusEnumValues = []NodeStatus{
	NodeStatusReady,
	NodeStatusWarning,
	NodeStatusUnderMaintenance,
	NodeStatusUnschedulable,
}

// GetAllowedValues returns the list of possible values.
func (v *NodeStatus) GetAllowedValues() []NodeStatus {
	return allowedNodeStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *NodeStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NodeStatus(value)
	return nil
}

// NewNodeStatusFromValue returns a pointer to a valid NodeStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewNodeStatusFromValue(v string) (*NodeStatus, error) {
	ev := NodeStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NodeStatus: valid values are %v", v, allowedNodeStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v NodeStatus) IsValid() bool {
	for _, existing := range allowedNodeStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NodeStatus value.
func (v NodeStatus) Ptr() *NodeStatus {
	return &v
}
