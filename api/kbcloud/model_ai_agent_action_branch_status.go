// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentActionBranchStatus string

// List of AiAgentActionBranchStatus.
const (
	AiAgentActionBranchStatusNon                 AiAgentActionBranchStatus = "none"
	AiAgentActionBranchStatusPendingConfirmation AiAgentActionBranchStatus = "pending_confirmation"
	AiAgentActionBranchStatusApproved            AiAgentActionBranchStatus = "approved"
	AiAgentActionBranchStatusExecuting           AiAgentActionBranchStatus = "executing"
	AiAgentActionBranchStatusCompleted           AiAgentActionBranchStatus = "completed"
	AiAgentActionBranchStatusFailed              AiAgentActionBranchStatus = "failed"
	AiAgentActionBranchStatusCancelled           AiAgentActionBranchStatus = "cancelled"
	AiAgentActionBranchStatusExpired             AiAgentActionBranchStatus = "expired"
)

var allowedAiAgentActionBranchStatusEnumValues = []AiAgentActionBranchStatus{
	AiAgentActionBranchStatusNon,
	AiAgentActionBranchStatusPendingConfirmation,
	AiAgentActionBranchStatusApproved,
	AiAgentActionBranchStatusExecuting,
	AiAgentActionBranchStatusCompleted,
	AiAgentActionBranchStatusFailed,
	AiAgentActionBranchStatusCancelled,
	AiAgentActionBranchStatusExpired,
}

// GetAllowedValues returns the list of possible values.
func (v *AiAgentActionBranchStatus) GetAllowedValues() []AiAgentActionBranchStatus {
	return allowedAiAgentActionBranchStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiAgentActionBranchStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiAgentActionBranchStatus(value)
	return nil
}

// NewAiAgentActionBranchStatusFromValue returns a pointer to a valid AiAgentActionBranchStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiAgentActionBranchStatusFromValue(v string) (*AiAgentActionBranchStatus, error) {
	ev := AiAgentActionBranchStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiAgentActionBranchStatus: valid values are %v", v, allowedAiAgentActionBranchStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiAgentActionBranchStatus) IsValid() bool {
	for _, existing := range allowedAiAgentActionBranchStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiAgentActionBranchStatus value.
func (v AiAgentActionBranchStatus) Ptr() *AiAgentActionBranchStatus {
	return &v
}
