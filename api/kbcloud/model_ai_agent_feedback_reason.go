// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentFeedbackReason string

// List of AiAgentFeedbackReason.
const (
	AiAgentFeedbackReasonHelpful            AiAgentFeedbackReason = "helpful"
	AiAgentFeedbackReasonNotHelpful         AiAgentFeedbackReason = "not_helpful"
	AiAgentFeedbackReasonWrongRootCause     AiAgentFeedbackReason = "wrong_root_cause"
	AiAgentFeedbackReasonMissingEvidence    AiAgentFeedbackReason = "missing_evidence"
	AiAgentFeedbackReasonUnsafeAction       AiAgentFeedbackReason = "unsafe_action"
	AiAgentFeedbackReasonUnclearExplanation AiAgentFeedbackReason = "unclear_explanation"
)

var allowedAiAgentFeedbackReasonEnumValues = []AiAgentFeedbackReason{
	AiAgentFeedbackReasonHelpful,
	AiAgentFeedbackReasonNotHelpful,
	AiAgentFeedbackReasonWrongRootCause,
	AiAgentFeedbackReasonMissingEvidence,
	AiAgentFeedbackReasonUnsafeAction,
	AiAgentFeedbackReasonUnclearExplanation,
}

// GetAllowedValues returns the list of possible values.
func (v *AiAgentFeedbackReason) GetAllowedValues() []AiAgentFeedbackReason {
	return allowedAiAgentFeedbackReasonEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiAgentFeedbackReason) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiAgentFeedbackReason(value)
	return nil
}

// NewAiAgentFeedbackReasonFromValue returns a pointer to a valid AiAgentFeedbackReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiAgentFeedbackReasonFromValue(v string) (*AiAgentFeedbackReason, error) {
	ev := AiAgentFeedbackReason(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiAgentFeedbackReason: valid values are %v", v, allowedAiAgentFeedbackReasonEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiAgentFeedbackReason) IsValid() bool {
	for _, existing := range allowedAiAgentFeedbackReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiAgentFeedbackReason value.
func (v AiAgentFeedbackReason) Ptr() *AiAgentFeedbackReason {
	return &v
}
