// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentRunStatus string

// List of AiAgentRunStatus.
const (
	AiAgentRunStatusPlanning           AiAgentRunStatus = "planning"
	AiAgentRunStatusCollectingEvidence AiAgentRunStatus = "collecting_evidence"
	AiAgentRunStatusAnalyzing          AiAgentRunStatus = "analyzing"
	AiAgentRunStatusAwaitingApproval   AiAgentRunStatus = "awaiting_approval"
	AiAgentRunStatusExecuting          AiAgentRunStatus = "executing"
	AiAgentRunStatusSummarizing        AiAgentRunStatus = "summarizing"
	AiAgentRunStatusWaitingForUser     AiAgentRunStatus = "waiting_for_user"
	AiAgentRunStatusPaused             AiAgentRunStatus = "paused"
	AiAgentRunStatusCompleted          AiAgentRunStatus = "completed"
	AiAgentRunStatusFailed             AiAgentRunStatus = "failed"
	AiAgentRunStatusCancelled          AiAgentRunStatus = "cancelled"
)

var allowedAiAgentRunStatusEnumValues = []AiAgentRunStatus{
	AiAgentRunStatusPlanning,
	AiAgentRunStatusCollectingEvidence,
	AiAgentRunStatusAnalyzing,
	AiAgentRunStatusAwaitingApproval,
	AiAgentRunStatusExecuting,
	AiAgentRunStatusSummarizing,
	AiAgentRunStatusWaitingForUser,
	AiAgentRunStatusPaused,
	AiAgentRunStatusCompleted,
	AiAgentRunStatusFailed,
	AiAgentRunStatusCancelled,
}

// GetAllowedValues returns the list of possible values.
func (v *AiAgentRunStatus) GetAllowedValues() []AiAgentRunStatus {
	return allowedAiAgentRunStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiAgentRunStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiAgentRunStatus(value)
	return nil
}

// NewAiAgentRunStatusFromValue returns a pointer to a valid AiAgentRunStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiAgentRunStatusFromValue(v string) (*AiAgentRunStatus, error) {
	ev := AiAgentRunStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiAgentRunStatus: valid values are %v", v, allowedAiAgentRunStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiAgentRunStatus) IsValid() bool {
	for _, existing := range allowedAiAgentRunStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiAgentRunStatus value.
func (v AiAgentRunStatus) Ptr() *AiAgentRunStatus {
	return &v
}
