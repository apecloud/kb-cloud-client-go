// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// InstanceResourceStatsType Type of the instance, choose cluster or system
type InstanceResourceStatsType string

// List of InstanceResourceStatsType.
const (
	InstanceResourceStatsTypeCluster InstanceResourceStatsType = "cluster"
	InstanceResourceStatsTypeSystem  InstanceResourceStatsType = "system"
)

var allowedInstanceResourceStatsTypeEnumValues = []InstanceResourceStatsType{
	InstanceResourceStatsTypeCluster,
	InstanceResourceStatsTypeSystem,
}

// GetAllowedValues returns the list of possible values.
func (v *InstanceResourceStatsType) GetAllowedValues() []InstanceResourceStatsType {
	return allowedInstanceResourceStatsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *InstanceResourceStatsType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = InstanceResourceStatsType(value)
	return nil
}

// NewInstanceResourceStatsTypeFromValue returns a pointer to a valid InstanceResourceStatsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewInstanceResourceStatsTypeFromValue(v string) (*InstanceResourceStatsType, error) {
	ev := InstanceResourceStatsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for InstanceResourceStatsType: valid values are %v", v, allowedInstanceResourceStatsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v InstanceResourceStatsType) IsValid() bool {
	for _, existing := range allowedInstanceResourceStatsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InstanceResourceStatsType value.
func (v InstanceResourceStatsType) Ptr() *InstanceResourceStatsType {
	return &v
}
