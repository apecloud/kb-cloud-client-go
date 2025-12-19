// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// InspectionSupportedEngines Specifies the supported engines for the inspection task.
type InspectionSupportedEngines string

// List of InspectionSupportedEngines.
const (
	InspectionSupportedEnginesNode     InspectionSupportedEngines = "node"
	InspectionSupportedEnginesMysql    InspectionSupportedEngines = "mysql"
	InspectionSupportedEnginesRedis    InspectionSupportedEngines = "redis"
	InspectionSupportedEnginesDamengdb InspectionSupportedEngines = "damengdb"
	InspectionSupportedEnginesKafka    InspectionSupportedEngines = "kafka"
)

var allowedInspectionSupportedEnginesEnumValues = []InspectionSupportedEngines{
	InspectionSupportedEnginesNode,
	InspectionSupportedEnginesMysql,
	InspectionSupportedEnginesRedis,
	InspectionSupportedEnginesDamengdb,
	InspectionSupportedEnginesKafka,
}

// GetAllowedValues returns the list of possible values.
func (v *InspectionSupportedEngines) GetAllowedValues() []InspectionSupportedEngines {
	return allowedInspectionSupportedEnginesEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *InspectionSupportedEngines) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = InspectionSupportedEngines(value)
	return nil
}

// NewInspectionSupportedEnginesFromValue returns a pointer to a valid InspectionSupportedEngines
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewInspectionSupportedEnginesFromValue(v string) (*InspectionSupportedEngines, error) {
	ev := InspectionSupportedEngines(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for InspectionSupportedEngines: valid values are %v", v, allowedInspectionSupportedEnginesEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v InspectionSupportedEngines) IsValid() bool {
	for _, existing := range allowedInspectionSupportedEnginesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InspectionSupportedEngines value.
func (v InspectionSupportedEngines) Ptr() *InspectionSupportedEngines {
	return &v
}
