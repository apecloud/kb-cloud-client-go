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
	AiAgentEventTypeCloudConversationReady         AiAgentEventType = "cloud.conversation.ready"
	AiAgentEventTypeCloudStreamConnected           AiAgentEventType = "cloud.stream.connected"
	AiAgentEventTypeCloudTurnStarted               AiAgentEventType = "cloud.turn.started"
	AiAgentEventTypeMessageDelta                   AiAgentEventType = "message.delta"
	AiAgentEventTypeCloudAssistantMessageCompleted AiAgentEventType = "cloud.assistant.message.completed"
	AiAgentEventTypeReasoningAvailable             AiAgentEventType = "reasoning.available"
	AiAgentEventTypeAgentStep                      AiAgentEventType = "agent.step"
	AiAgentEventTypeToolGenerating                 AiAgentEventType = "tool.generating"
	AiAgentEventTypeToolStarted                    AiAgentEventType = "tool.started"
	AiAgentEventTypeToolCompleted                  AiAgentEventType = "tool.completed"
	AiAgentEventTypeApprovalRequest                AiAgentEventType = "approval.request"
	AiAgentEventTypeApprovalResponded              AiAgentEventType = "approval.responded"
	AiAgentEventTypeRunCompleted                   AiAgentEventType = "run.completed"
	AiAgentEventTypeRunFailed                      AiAgentEventType = "run.failed"
	AiAgentEventTypeRunCancelled                   AiAgentEventType = "run.cancelled"
)

var allowedAiAgentEventTypeEnumValues = []AiAgentEventType{
	AiAgentEventTypeCloudConversationReady,
	AiAgentEventTypeCloudStreamConnected,
	AiAgentEventTypeCloudTurnStarted,
	AiAgentEventTypeMessageDelta,
	AiAgentEventTypeCloudAssistantMessageCompleted,
	AiAgentEventTypeReasoningAvailable,
	AiAgentEventTypeAgentStep,
	AiAgentEventTypeToolGenerating,
	AiAgentEventTypeToolStarted,
	AiAgentEventTypeToolCompleted,
	AiAgentEventTypeApprovalRequest,
	AiAgentEventTypeApprovalResponded,
	AiAgentEventTypeRunCompleted,
	AiAgentEventTypeRunFailed,
	AiAgentEventTypeRunCancelled,
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
