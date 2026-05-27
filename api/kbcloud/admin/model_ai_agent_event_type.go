// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentEventType string

// List of AiAgentEventType.
const (
	AiAgentEventTypeConversationReady         AiAgentEventType = "conversation_ready"
	AiAgentEventTypeTurnStarted               AiAgentEventType = "turn_started"
	AiAgentEventTypeAssistantDelta            AiAgentEventType = "assistant_delta"
	AiAgentEventTypeAssistantInterim          AiAgentEventType = "assistant_interim"
	AiAgentEventTypeAssistantMessageCompleted AiAgentEventType = "assistant_message_completed"
	AiAgentEventTypeAgentReasoning            AiAgentEventType = "agent_reasoning"
	AiAgentEventTypeAgentThinking             AiAgentEventType = "agent_thinking"
	AiAgentEventTypeAgentStep                 AiAgentEventType = "agent_step"
	AiAgentEventTypeToolConfirmationRequested AiAgentEventType = "tool_confirmation_requested"
	AiAgentEventTypeToolConfirmationDecided   AiAgentEventType = "tool_confirmation_decided"
	AiAgentEventTypeToolGenerating            AiAgentEventType = "tool_generating"
	AiAgentEventTypeToolStarted               AiAgentEventType = "tool_started"
	AiAgentEventTypeToolCompleted             AiAgentEventType = "tool_completed"
	AiAgentEventTypeTurnCompleted             AiAgentEventType = "turn_completed"
	AiAgentEventTypeTurnFailed                AiAgentEventType = "turn_failed"
	AiAgentEventTypeTurnCancelled             AiAgentEventType = "turn_cancelled"
	AiAgentEventTypeRuntimeStatusChanged      AiAgentEventType = "runtime_status_changed"
)

var allowedAiAgentEventTypeEnumValues = []AiAgentEventType{
	AiAgentEventTypeConversationReady,
	AiAgentEventTypeTurnStarted,
	AiAgentEventTypeAssistantDelta,
	AiAgentEventTypeAssistantInterim,
	AiAgentEventTypeAssistantMessageCompleted,
	AiAgentEventTypeAgentReasoning,
	AiAgentEventTypeAgentThinking,
	AiAgentEventTypeAgentStep,
	AiAgentEventTypeToolConfirmationRequested,
	AiAgentEventTypeToolConfirmationDecided,
	AiAgentEventTypeToolGenerating,
	AiAgentEventTypeToolStarted,
	AiAgentEventTypeToolCompleted,
	AiAgentEventTypeTurnCompleted,
	AiAgentEventTypeTurnFailed,
	AiAgentEventTypeTurnCancelled,
	AiAgentEventTypeRuntimeStatusChanged,
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
