// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EngineOptionsServicePattern service name pattern, e.g. ClusterName-ComponentName or .ClusterName`
type EngineOptionsServicePattern string

// List of EngineOptionsServicePattern.
const (
	ENGINEOPTIONSSERVICEPATTERN_CLUSTER_COMPONENT         EngineOptionsServicePattern = "cluster-component"
	ENGINEOPTIONSSERVICEPATTERN_CLUSTER                   EngineOptionsServicePattern = "cluster"
	ENGINEOPTIONSSERVICEPATTERN_CLUSTER_COMPONENT_SERVICE EngineOptionsServicePattern = "cluster-component-service"
)

var allowedEngineOptionsServicePatternEnumValues = []EngineOptionsServicePattern{
	ENGINEOPTIONSSERVICEPATTERN_CLUSTER_COMPONENT,
	ENGINEOPTIONSSERVICEPATTERN_CLUSTER,
	ENGINEOPTIONSSERVICEPATTERN_CLUSTER_COMPONENT_SERVICE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EngineOptionsServicePattern) GetAllowedValues() []EngineOptionsServicePattern {
	return allowedEngineOptionsServicePatternEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EngineOptionsServicePattern) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EngineOptionsServicePattern(value)
	return nil
}

// NewEngineOptionsServicePatternFromValue returns a pointer to a valid EngineOptionsServicePattern
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEngineOptionsServicePatternFromValue(v string) (*EngineOptionsServicePattern, error) {
	ev := EngineOptionsServicePattern(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EngineOptionsServicePattern: valid values are %v", v, allowedEngineOptionsServicePatternEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EngineOptionsServicePattern) IsValid() bool {
	for _, existing := range allowedEngineOptionsServicePatternEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EngineOptionsServicePattern value.
func (v EngineOptionsServicePattern) Ptr() *EngineOptionsServicePattern {
	return &v
}
