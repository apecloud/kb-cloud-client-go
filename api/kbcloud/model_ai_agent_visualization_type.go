// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentVisualizationType string

// List of AiAgentVisualizationType.
const (
	AiAgentVisualizationTypeLine  AiAgentVisualizationType = "line"
	AiAgentVisualizationTypeBar   AiAgentVisualizationType = "bar"
	AiAgentVisualizationTypeTable AiAgentVisualizationType = "table"
	AiAgentVisualizationTypeStat  AiAgentVisualizationType = "stat"
)

var allowedAiAgentVisualizationTypeEnumValues = []AiAgentVisualizationType{
	AiAgentVisualizationTypeLine,
	AiAgentVisualizationTypeBar,
	AiAgentVisualizationTypeTable,
	AiAgentVisualizationTypeStat,
}

// GetAllowedValues returns the list of possible values.
func (v *AiAgentVisualizationType) GetAllowedValues() []AiAgentVisualizationType {
	return allowedAiAgentVisualizationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiAgentVisualizationType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiAgentVisualizationType(value)
	return nil
}

// NewAiAgentVisualizationTypeFromValue returns a pointer to a valid AiAgentVisualizationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiAgentVisualizationTypeFromValue(v string) (*AiAgentVisualizationType, error) {
	ev := AiAgentVisualizationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiAgentVisualizationType: valid values are %v", v, allowedAiAgentVisualizationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiAgentVisualizationType) IsValid() bool {
	for _, existing := range allowedAiAgentVisualizationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiAgentVisualizationType value.
func (v AiAgentVisualizationType) Ptr() *AiAgentVisualizationType {
	return &v
}
