// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// TopologyType Topology level at which pod placement constraints are evaluated.
type TopologyType string

// List of TopologyType.
const (
	TopologyTypeZone TopologyType = "Zone"
	TopologyTypeNode TopologyType = "Node"
)

var allowedTopologyTypeEnumValues = []TopologyType{
	TopologyTypeZone,
	TopologyTypeNode,
}

// GetAllowedValues returns the list of possible values.
func (v *TopologyType) GetAllowedValues() []TopologyType {
	return allowedTopologyTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TopologyType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TopologyType(value)
	return nil
}

// NewTopologyTypeFromValue returns a pointer to a valid TopologyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTopologyTypeFromValue(v string) (*TopologyType, error) {
	ev := TopologyType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TopologyType: valid values are %v", v, allowedTopologyTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TopologyType) IsValid() bool {
	for _, existing := range allowedTopologyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TopologyType value.
func (v TopologyType) Ptr() *TopologyType {
	return &v
}
