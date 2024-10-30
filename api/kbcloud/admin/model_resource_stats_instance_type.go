// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package admin

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)



// ResourceStatsInstanceType Type of the instance, choose cluster or system
type ResourceStatsInstanceType string

// List of ResourceStatsInstanceType.
const (
	ResourceStatsInstanceTypeCluster ResourceStatsInstanceType = "cluster"
	ResourceStatsInstanceTypeSystem ResourceStatsInstanceType = "system"
)

var allowedResourceStatsInstanceTypeEnumValues = []ResourceStatsInstanceType{
	ResourceStatsInstanceTypeCluster,
	ResourceStatsInstanceTypeSystem,
}

// GetAllowedValues returns the list of possible values.
func (v *ResourceStatsInstanceType) GetAllowedValues() []ResourceStatsInstanceType {
	return allowedResourceStatsInstanceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ResourceStatsInstanceType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ResourceStatsInstanceType(value)
	return nil
}

// NewResourceStatsInstanceTypeFromValue returns a pointer to a valid ResourceStatsInstanceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewResourceStatsInstanceTypeFromValue(v string) (*ResourceStatsInstanceType, error) {
	ev := ResourceStatsInstanceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ResourceStatsInstanceType: valid values are %v", v, allowedResourceStatsInstanceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ResourceStatsInstanceType) IsValid() bool {
	for _, existing := range allowedResourceStatsInstanceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ResourceStatsInstanceType value.
func (v ResourceStatsInstanceType) Ptr() *ResourceStatsInstanceType {
	return &v
}
