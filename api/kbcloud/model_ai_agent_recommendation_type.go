// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentRecommendationType string

// List of AiAgentRecommendationType.
const (
	AiAgentRecommendationTypeReadOnly            AiAgentRecommendationType = "read_only"
	AiAgentRecommendationTypeActionPlanAvailable AiAgentRecommendationType = "action_plan_available"
)

var allowedAiAgentRecommendationTypeEnumValues = []AiAgentRecommendationType{
	AiAgentRecommendationTypeReadOnly,
	AiAgentRecommendationTypeActionPlanAvailable,
}

// GetAllowedValues returns the list of possible values.
func (v *AiAgentRecommendationType) GetAllowedValues() []AiAgentRecommendationType {
	return allowedAiAgentRecommendationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiAgentRecommendationType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiAgentRecommendationType(value)
	return nil
}

// NewAiAgentRecommendationTypeFromValue returns a pointer to a valid AiAgentRecommendationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiAgentRecommendationTypeFromValue(v string) (*AiAgentRecommendationType, error) {
	ev := AiAgentRecommendationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiAgentRecommendationType: valid values are %v", v, allowedAiAgentRecommendationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiAgentRecommendationType) IsValid() bool {
	for _, existing := range allowedAiAgentRecommendationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiAgentRecommendationType value.
func (v AiAgentRecommendationType) Ptr() *AiAgentRecommendationType {
	return &v
}
