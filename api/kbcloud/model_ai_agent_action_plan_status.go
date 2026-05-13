// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentActionPlanStatus string

// List of AiAgentActionPlanStatus.
const (
	AiAgentActionPlanStatusPendingConfirmation AiAgentActionPlanStatus = "pending_confirmation"
	AiAgentActionPlanStatusApproved            AiAgentActionPlanStatus = "approved"
	AiAgentActionPlanStatusExecuting           AiAgentActionPlanStatus = "executing"
	AiAgentActionPlanStatusCompleted           AiAgentActionPlanStatus = "completed"
	AiAgentActionPlanStatusFailed              AiAgentActionPlanStatus = "failed"
	AiAgentActionPlanStatusCancelled           AiAgentActionPlanStatus = "cancelled"
	AiAgentActionPlanStatusExpired             AiAgentActionPlanStatus = "expired"
)

var allowedAiAgentActionPlanStatusEnumValues = []AiAgentActionPlanStatus{
	AiAgentActionPlanStatusPendingConfirmation,
	AiAgentActionPlanStatusApproved,
	AiAgentActionPlanStatusExecuting,
	AiAgentActionPlanStatusCompleted,
	AiAgentActionPlanStatusFailed,
	AiAgentActionPlanStatusCancelled,
	AiAgentActionPlanStatusExpired,
}

// GetAllowedValues returns the list of possible values.
func (v *AiAgentActionPlanStatus) GetAllowedValues() []AiAgentActionPlanStatus {
	return allowedAiAgentActionPlanStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiAgentActionPlanStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiAgentActionPlanStatus(value)
	return nil
}

// NewAiAgentActionPlanStatusFromValue returns a pointer to a valid AiAgentActionPlanStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiAgentActionPlanStatusFromValue(v string) (*AiAgentActionPlanStatus, error) {
	ev := AiAgentActionPlanStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiAgentActionPlanStatus: valid values are %v", v, allowedAiAgentActionPlanStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiAgentActionPlanStatus) IsValid() bool {
	for _, existing := range allowedAiAgentActionPlanStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiAgentActionPlanStatus value.
func (v AiAgentActionPlanStatus) Ptr() *AiAgentActionPlanStatus {
	return &v
}
