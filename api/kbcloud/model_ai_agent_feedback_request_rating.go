// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentFeedbackRequestRating string

// List of AiAgentFeedbackRequestRating.
const (
	AiAgentFeedbackRequestRatingHelpful    AiAgentFeedbackRequestRating = "helpful"
	AiAgentFeedbackRequestRatingNotHelpful AiAgentFeedbackRequestRating = "not_helpful"
)

var allowedAiAgentFeedbackRequestRatingEnumValues = []AiAgentFeedbackRequestRating{
	AiAgentFeedbackRequestRatingHelpful,
	AiAgentFeedbackRequestRatingNotHelpful,
}

// GetAllowedValues returns the list of possible values.
func (v *AiAgentFeedbackRequestRating) GetAllowedValues() []AiAgentFeedbackRequestRating {
	return allowedAiAgentFeedbackRequestRatingEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiAgentFeedbackRequestRating) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiAgentFeedbackRequestRating(value)
	return nil
}

// NewAiAgentFeedbackRequestRatingFromValue returns a pointer to a valid AiAgentFeedbackRequestRating
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiAgentFeedbackRequestRatingFromValue(v string) (*AiAgentFeedbackRequestRating, error) {
	ev := AiAgentFeedbackRequestRating(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiAgentFeedbackRequestRating: valid values are %v", v, allowedAiAgentFeedbackRequestRatingEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiAgentFeedbackRequestRating) IsValid() bool {
	for _, existing := range allowedAiAgentFeedbackRequestRatingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiAgentFeedbackRequestRating value.
func (v AiAgentFeedbackRequestRating) Ptr() *AiAgentFeedbackRequestRating {
	return &v
}
