// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentConfidence string

// List of AiAgentConfidence.
const (
	AiAgentConfidenceHigh   AiAgentConfidence = "high"
	AiAgentConfidenceMedium AiAgentConfidence = "medium"
	AiAgentConfidenceLow    AiAgentConfidence = "low"
	AiAgentConfidenceNon    AiAgentConfidence = "none"
)

var allowedAiAgentConfidenceEnumValues = []AiAgentConfidence{
	AiAgentConfidenceHigh,
	AiAgentConfidenceMedium,
	AiAgentConfidenceLow,
	AiAgentConfidenceNon,
}

// GetAllowedValues returns the list of possible values.
func (v *AiAgentConfidence) GetAllowedValues() []AiAgentConfidence {
	return allowedAiAgentConfidenceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiAgentConfidence) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiAgentConfidence(value)
	return nil
}

// NewAiAgentConfidenceFromValue returns a pointer to a valid AiAgentConfidence
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiAgentConfidenceFromValue(v string) (*AiAgentConfidence, error) {
	ev := AiAgentConfidence(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiAgentConfidence: valid values are %v", v, allowedAiAgentConfidenceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiAgentConfidence) IsValid() bool {
	for _, existing := range allowedAiAgentConfidenceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiAgentConfidence value.
func (v AiAgentConfidence) Ptr() *AiAgentConfidence {
	return &v
}
