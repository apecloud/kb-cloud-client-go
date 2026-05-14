// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentFeedbackRequestReasonsItem string

// List of AiAgentFeedbackRequestReasonsItem.
const (
	AiAgentFeedbackRequestReasonsItemHelpful            AiAgentFeedbackRequestReasonsItem = "helpful"
	AiAgentFeedbackRequestReasonsItemNotHelpful         AiAgentFeedbackRequestReasonsItem = "not_helpful"
	AiAgentFeedbackRequestReasonsItemWrongRootCause     AiAgentFeedbackRequestReasonsItem = "wrong_root_cause"
	AiAgentFeedbackRequestReasonsItemMissingEvidence    AiAgentFeedbackRequestReasonsItem = "missing_evidence"
	AiAgentFeedbackRequestReasonsItemUnsafeAction       AiAgentFeedbackRequestReasonsItem = "unsafe_action"
	AiAgentFeedbackRequestReasonsItemUnclearExplanation AiAgentFeedbackRequestReasonsItem = "unclear_explanation"
)

var allowedAiAgentFeedbackRequestReasonsItemEnumValues = []AiAgentFeedbackRequestReasonsItem{
	AiAgentFeedbackRequestReasonsItemHelpful,
	AiAgentFeedbackRequestReasonsItemNotHelpful,
	AiAgentFeedbackRequestReasonsItemWrongRootCause,
	AiAgentFeedbackRequestReasonsItemMissingEvidence,
	AiAgentFeedbackRequestReasonsItemUnsafeAction,
	AiAgentFeedbackRequestReasonsItemUnclearExplanation,
}

// GetAllowedValues returns the list of possible values.
func (v *AiAgentFeedbackRequestReasonsItem) GetAllowedValues() []AiAgentFeedbackRequestReasonsItem {
	return allowedAiAgentFeedbackRequestReasonsItemEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiAgentFeedbackRequestReasonsItem) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiAgentFeedbackRequestReasonsItem(value)
	return nil
}

// NewAiAgentFeedbackRequestReasonsItemFromValue returns a pointer to a valid AiAgentFeedbackRequestReasonsItem
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiAgentFeedbackRequestReasonsItemFromValue(v string) (*AiAgentFeedbackRequestReasonsItem, error) {
	ev := AiAgentFeedbackRequestReasonsItem(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiAgentFeedbackRequestReasonsItem: valid values are %v", v, allowedAiAgentFeedbackRequestReasonsItemEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiAgentFeedbackRequestReasonsItem) IsValid() bool {
	for _, existing := range allowedAiAgentFeedbackRequestReasonsItemEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiAgentFeedbackRequestReasonsItem value.
func (v AiAgentFeedbackRequestReasonsItem) Ptr() *AiAgentFeedbackRequestReasonsItem {
	return &v
}
