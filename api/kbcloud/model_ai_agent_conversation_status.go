// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentConversationStatus string

// List of AiAgentConversationStatus.
const (
	AiAgentConversationStatusActive    AiAgentConversationStatus = "active"
	AiAgentConversationStatusCompleted AiAgentConversationStatus = "completed"
	AiAgentConversationStatusFailed    AiAgentConversationStatus = "failed"
	AiAgentConversationStatusCancelled AiAgentConversationStatus = "cancelled"
	AiAgentConversationStatusArchived  AiAgentConversationStatus = "archived"
)

var allowedAiAgentConversationStatusEnumValues = []AiAgentConversationStatus{
	AiAgentConversationStatusActive,
	AiAgentConversationStatusCompleted,
	AiAgentConversationStatusFailed,
	AiAgentConversationStatusCancelled,
	AiAgentConversationStatusArchived,
}

// GetAllowedValues returns the list of possible values.
func (v *AiAgentConversationStatus) GetAllowedValues() []AiAgentConversationStatus {
	return allowedAiAgentConversationStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiAgentConversationStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiAgentConversationStatus(value)
	return nil
}

// NewAiAgentConversationStatusFromValue returns a pointer to a valid AiAgentConversationStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiAgentConversationStatusFromValue(v string) (*AiAgentConversationStatus, error) {
	ev := AiAgentConversationStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiAgentConversationStatus: valid values are %v", v, allowedAiAgentConversationStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiAgentConversationStatus) IsValid() bool {
	for _, existing := range allowedAiAgentConversationStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiAgentConversationStatus value.
func (v AiAgentConversationStatus) Ptr() *AiAgentConversationStatus {
	return &v
}
