// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentFeedbackRating string

// List of AiAgentFeedbackRating.
const (
	AiAgentFeedbackRatingHelpful    AiAgentFeedbackRating = "helpful"
	AiAgentFeedbackRatingNotHelpful AiAgentFeedbackRating = "not_helpful"
)

var allowedAiAgentFeedbackRatingEnumValues = []AiAgentFeedbackRating{
	AiAgentFeedbackRatingHelpful,
	AiAgentFeedbackRatingNotHelpful,
}

// GetAllowedValues returns the list of possible values.
func (v *AiAgentFeedbackRating) GetAllowedValues() []AiAgentFeedbackRating {
	return allowedAiAgentFeedbackRatingEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiAgentFeedbackRating) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiAgentFeedbackRating(value)
	return nil
}

// NewAiAgentFeedbackRatingFromValue returns a pointer to a valid AiAgentFeedbackRating
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiAgentFeedbackRatingFromValue(v string) (*AiAgentFeedbackRating, error) {
	ev := AiAgentFeedbackRating(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiAgentFeedbackRating: valid values are %v", v, allowedAiAgentFeedbackRatingEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiAgentFeedbackRating) IsValid() bool {
	for _, existing := range allowedAiAgentFeedbackRatingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiAgentFeedbackRating value.
func (v AiAgentFeedbackRating) Ptr() *AiAgentFeedbackRating {
	return &v
}
