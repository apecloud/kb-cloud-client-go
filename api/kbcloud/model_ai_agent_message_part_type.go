// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentMessagePartType string

// List of AiAgentMessagePartType.
const (
	AiAgentMessagePartTypeText                   AiAgentMessagePartType = "text"
	AiAgentMessagePartTypeReasoningSummary       AiAgentMessagePartType = "reasoning_summary"
	AiAgentMessagePartTypeDiagnosisSummary       AiAgentMessagePartType = "diagnosis_summary"
	AiAgentMessagePartTypeToolObservationSummary AiAgentMessagePartType = "tool_observation_summary"
)

var allowedAiAgentMessagePartTypeEnumValues = []AiAgentMessagePartType{
	AiAgentMessagePartTypeText,
	AiAgentMessagePartTypeReasoningSummary,
	AiAgentMessagePartTypeDiagnosisSummary,
	AiAgentMessagePartTypeToolObservationSummary,
}

// GetAllowedValues returns the list of possible values.
func (v *AiAgentMessagePartType) GetAllowedValues() []AiAgentMessagePartType {
	return allowedAiAgentMessagePartTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiAgentMessagePartType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiAgentMessagePartType(value)
	return nil
}

// NewAiAgentMessagePartTypeFromValue returns a pointer to a valid AiAgentMessagePartType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiAgentMessagePartTypeFromValue(v string) (*AiAgentMessagePartType, error) {
	ev := AiAgentMessagePartType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiAgentMessagePartType: valid values are %v", v, allowedAiAgentMessagePartTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiAgentMessagePartType) IsValid() bool {
	for _, existing := range allowedAiAgentMessagePartTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiAgentMessagePartType value.
func (v AiAgentMessagePartType) Ptr() *AiAgentMessagePartType {
	return &v
}
