// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentEventType string

// List of AiAgentEventType.
const (
	AiAgentEventTypeRunStarted           AiAgentEventType = "run_started"
	AiAgentEventTypeThinking             AiAgentEventType = "thinking"
	AiAgentEventTypeAssistantDelta       AiAgentEventType = "assistant_delta"
	AiAgentEventTypeStageUpdate          AiAgentEventType = "stage_update"
	AiAgentEventTypeEvidenceUpdate       AiAgentEventType = "evidence_update"
	AiAgentEventTypeReportReady          AiAgentEventType = "report_ready"
	AiAgentEventTypeApprovalRequired     AiAgentEventType = "approval_required"
	AiAgentEventTypeRunCompleted         AiAgentEventType = "run_completed"
	AiAgentEventTypeRunFailed            AiAgentEventType = "run_failed"
	AiAgentEventTypeRunCancelled         AiAgentEventType = "run_cancelled"
	AiAgentEventTypePlanCreated          AiAgentEventType = "plan_created"
	AiAgentEventTypePlanUpdated          AiAgentEventType = "plan_updated"
	AiAgentEventTypeToolRequestStarted   AiAgentEventType = "tool_request_started"
	AiAgentEventTypeToolRequestCompleted AiAgentEventType = "tool_request_completed"
	AiAgentEventTypeToolRequestFailed    AiAgentEventType = "tool_request_failed"
	AiAgentEventTypeToolRequestRejected  AiAgentEventType = "tool_request_rejected"
	AiAgentEventTypeNeedsUserInput       AiAgentEventType = "needs_user_input"
	AiAgentEventTypeBudgetUpdate         AiAgentEventType = "budget_update"
	AiAgentEventTypeRunPaused            AiAgentEventType = "run_paused"
	AiAgentEventTypeRunResumed           AiAgentEventType = "run_resumed"
)

var allowedAiAgentEventTypeEnumValues = []AiAgentEventType{
	AiAgentEventTypeRunStarted,
	AiAgentEventTypeThinking,
	AiAgentEventTypeAssistantDelta,
	AiAgentEventTypeStageUpdate,
	AiAgentEventTypeEvidenceUpdate,
	AiAgentEventTypeReportReady,
	AiAgentEventTypeApprovalRequired,
	AiAgentEventTypeRunCompleted,
	AiAgentEventTypeRunFailed,
	AiAgentEventTypeRunCancelled,
	AiAgentEventTypePlanCreated,
	AiAgentEventTypePlanUpdated,
	AiAgentEventTypeToolRequestStarted,
	AiAgentEventTypeToolRequestCompleted,
	AiAgentEventTypeToolRequestFailed,
	AiAgentEventTypeToolRequestRejected,
	AiAgentEventTypeNeedsUserInput,
	AiAgentEventTypeBudgetUpdate,
	AiAgentEventTypeRunPaused,
	AiAgentEventTypeRunResumed,
}

// GetAllowedValues returns the list of possible values.
func (v *AiAgentEventType) GetAllowedValues() []AiAgentEventType {
	return allowedAiAgentEventTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiAgentEventType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiAgentEventType(value)
	return nil
}

// NewAiAgentEventTypeFromValue returns a pointer to a valid AiAgentEventType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiAgentEventTypeFromValue(v string) (*AiAgentEventType, error) {
	ev := AiAgentEventType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiAgentEventType: valid values are %v", v, allowedAiAgentEventTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiAgentEventType) IsValid() bool {
	for _, existing := range allowedAiAgentEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiAgentEventType value.
func (v AiAgentEventType) Ptr() *AiAgentEventType {
	return &v
}
