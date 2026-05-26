// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentContinueMode string

// List of AiAgentContinueMode.
const (
	AiAgentContinueModeContinueInvestigation AiAgentContinueMode = "continue_investigation"
	AiAgentContinueModeAnswerFollowup        AiAgentContinueMode = "answer_followup"
	AiAgentContinueModeNarrowScope           AiAgentContinueMode = "narrow_scope"
	AiAgentContinueModeClarifyContext        AiAgentContinueMode = "clarify_context"
)

var allowedAiAgentContinueModeEnumValues = []AiAgentContinueMode{
	AiAgentContinueModeContinueInvestigation,
	AiAgentContinueModeAnswerFollowup,
	AiAgentContinueModeNarrowScope,
	AiAgentContinueModeClarifyContext,
}

// GetAllowedValues returns the list of possible values.
func (v *AiAgentContinueMode) GetAllowedValues() []AiAgentContinueMode {
	return allowedAiAgentContinueModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiAgentContinueMode) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiAgentContinueMode(value)
	return nil
}

// NewAiAgentContinueModeFromValue returns a pointer to a valid AiAgentContinueMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiAgentContinueModeFromValue(v string) (*AiAgentContinueMode, error) {
	ev := AiAgentContinueMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiAgentContinueMode: valid values are %v", v, allowedAiAgentContinueModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiAgentContinueMode) IsValid() bool {
	for _, existing := range allowedAiAgentContinueModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiAgentContinueMode value.
func (v AiAgentContinueMode) Ptr() *AiAgentContinueMode {
	return &v
}
