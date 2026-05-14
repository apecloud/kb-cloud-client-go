// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentContinueReportRequestMode string

// List of AiAgentContinueReportRequestMode.
const (
	AiAgentContinueReportRequestModeRegenerateFromExistingEvidence     AiAgentContinueReportRequestMode = "regenerate_from_existing_evidence"
	AiAgentContinueReportRequestModeContinueWithMoreEvidence           AiAgentContinueReportRequestMode = "continue_with_more_evidence"
	AiAgentContinueReportRequestModeCreateActionPlanFromRecommendation AiAgentContinueReportRequestMode = "create_action_plan_from_recommendation"
)

var allowedAiAgentContinueReportRequestModeEnumValues = []AiAgentContinueReportRequestMode{
	AiAgentContinueReportRequestModeRegenerateFromExistingEvidence,
	AiAgentContinueReportRequestModeContinueWithMoreEvidence,
	AiAgentContinueReportRequestModeCreateActionPlanFromRecommendation,
}

// GetAllowedValues returns the list of possible values.
func (v *AiAgentContinueReportRequestMode) GetAllowedValues() []AiAgentContinueReportRequestMode {
	return allowedAiAgentContinueReportRequestModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiAgentContinueReportRequestMode) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiAgentContinueReportRequestMode(value)
	return nil
}

// NewAiAgentContinueReportRequestModeFromValue returns a pointer to a valid AiAgentContinueReportRequestMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiAgentContinueReportRequestModeFromValue(v string) (*AiAgentContinueReportRequestMode, error) {
	ev := AiAgentContinueReportRequestMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiAgentContinueReportRequestMode: valid values are %v", v, allowedAiAgentContinueReportRequestModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiAgentContinueReportRequestMode) IsValid() bool {
	for _, existing := range allowedAiAgentContinueReportRequestModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiAgentContinueReportRequestMode value.
func (v AiAgentContinueReportRequestMode) Ptr() *AiAgentContinueReportRequestMode {
	return &v
}
