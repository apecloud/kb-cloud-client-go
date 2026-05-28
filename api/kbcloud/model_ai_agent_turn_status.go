// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentTurnStatus string

// List of AiAgentTurnStatus.
const (
	AiAgentTurnStatusRunning            AiAgentTurnStatus = "running"
	AiAgentTurnStatusWaitingForApproval AiAgentTurnStatus = "waiting_for_approval"
	AiAgentTurnStatusCompleted          AiAgentTurnStatus = "completed"
	AiAgentTurnStatusFailed             AiAgentTurnStatus = "failed"
	AiAgentTurnStatusCancelled          AiAgentTurnStatus = "cancelled"
)

var allowedAiAgentTurnStatusEnumValues = []AiAgentTurnStatus{
	AiAgentTurnStatusRunning,
	AiAgentTurnStatusWaitingForApproval,
	AiAgentTurnStatusCompleted,
	AiAgentTurnStatusFailed,
	AiAgentTurnStatusCancelled,
}

// GetAllowedValues returns the list of possible values.
func (v *AiAgentTurnStatus) GetAllowedValues() []AiAgentTurnStatus {
	return allowedAiAgentTurnStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiAgentTurnStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiAgentTurnStatus(value)
	return nil
}

// NewAiAgentTurnStatusFromValue returns a pointer to a valid AiAgentTurnStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiAgentTurnStatusFromValue(v string) (*AiAgentTurnStatus, error) {
	ev := AiAgentTurnStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiAgentTurnStatus: valid values are %v", v, allowedAiAgentTurnStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiAgentTurnStatus) IsValid() bool {
	for _, existing := range allowedAiAgentTurnStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiAgentTurnStatus value.
func (v AiAgentTurnStatus) Ptr() *AiAgentTurnStatus {
	return &v
}
