// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentReportContinueMode string

// List of AiAgentReportContinueMode.
const (
	AiAgentReportContinueModeRegenerateFromExistingEvidence     AiAgentReportContinueMode = "regenerate_from_existing_evidence"
	AiAgentReportContinueModeContinueWithMoreEvidence           AiAgentReportContinueMode = "continue_with_more_evidence"
	AiAgentReportContinueModeCreateActionPlanFromRecommendation AiAgentReportContinueMode = "create_action_plan_from_recommendation"
)

var allowedAiAgentReportContinueModeEnumValues = []AiAgentReportContinueMode{
	AiAgentReportContinueModeRegenerateFromExistingEvidence,
	AiAgentReportContinueModeContinueWithMoreEvidence,
	AiAgentReportContinueModeCreateActionPlanFromRecommendation,
}

// GetAllowedValues returns the list of possible values.
func (v *AiAgentReportContinueMode) GetAllowedValues() []AiAgentReportContinueMode {
	return allowedAiAgentReportContinueModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiAgentReportContinueMode) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiAgentReportContinueMode(value)
	return nil
}

// NewAiAgentReportContinueModeFromValue returns a pointer to a valid AiAgentReportContinueMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiAgentReportContinueModeFromValue(v string) (*AiAgentReportContinueMode, error) {
	ev := AiAgentReportContinueMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiAgentReportContinueMode: valid values are %v", v, allowedAiAgentReportContinueModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiAgentReportContinueMode) IsValid() bool {
	for _, existing := range allowedAiAgentReportContinueModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiAgentReportContinueMode value.
func (v AiAgentReportContinueMode) Ptr() *AiAgentReportContinueMode {
	return &v
}
