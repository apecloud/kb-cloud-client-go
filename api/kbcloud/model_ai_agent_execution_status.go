// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentExecutionStatus string

// List of AiAgentExecutionStatus.
const (
	AiAgentExecutionStatusQueued    AiAgentExecutionStatus = "queued"
	AiAgentExecutionStatusRunning   AiAgentExecutionStatus = "running"
	AiAgentExecutionStatusCompleted AiAgentExecutionStatus = "completed"
	AiAgentExecutionStatusFailed    AiAgentExecutionStatus = "failed"
	AiAgentExecutionStatusCancelled AiAgentExecutionStatus = "cancelled"
	AiAgentExecutionStatusExpired   AiAgentExecutionStatus = "expired"
)

var allowedAiAgentExecutionStatusEnumValues = []AiAgentExecutionStatus{
	AiAgentExecutionStatusQueued,
	AiAgentExecutionStatusRunning,
	AiAgentExecutionStatusCompleted,
	AiAgentExecutionStatusFailed,
	AiAgentExecutionStatusCancelled,
	AiAgentExecutionStatusExpired,
}

// GetAllowedValues returns the list of possible values.
func (v *AiAgentExecutionStatus) GetAllowedValues() []AiAgentExecutionStatus {
	return allowedAiAgentExecutionStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiAgentExecutionStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiAgentExecutionStatus(value)
	return nil
}

// NewAiAgentExecutionStatusFromValue returns a pointer to a valid AiAgentExecutionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiAgentExecutionStatusFromValue(v string) (*AiAgentExecutionStatus, error) {
	ev := AiAgentExecutionStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiAgentExecutionStatus: valid values are %v", v, allowedAiAgentExecutionStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiAgentExecutionStatus) IsValid() bool {
	for _, existing := range allowedAiAgentExecutionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiAgentExecutionStatus value.
func (v AiAgentExecutionStatus) Ptr() *AiAgentExecutionStatus {
	return &v
}
