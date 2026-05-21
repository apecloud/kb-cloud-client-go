// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentToolApprovalStatus string

// List of AiAgentToolApprovalStatus.
const (
	AiAgentToolApprovalStatusPendingUserConfirmation AiAgentToolApprovalStatus = "pending_user_confirmation"
	AiAgentToolApprovalStatusPolicyAutoApproved      AiAgentToolApprovalStatus = "policy_auto_approved"
	AiAgentToolApprovalStatusBlocked                 AiAgentToolApprovalStatus = "blocked"
	AiAgentToolApprovalStatusUserApproved            AiAgentToolApprovalStatus = "user_approved"
	AiAgentToolApprovalStatusUserRejected            AiAgentToolApprovalStatus = "user_rejected"
	AiAgentToolApprovalStatusUserCancelled           AiAgentToolApprovalStatus = "user_cancelled"
	AiAgentToolApprovalStatusApprovalExpired         AiAgentToolApprovalStatus = "approval_expired"
	AiAgentToolApprovalStatusExecuting               AiAgentToolApprovalStatus = "executing"
	AiAgentToolApprovalStatusCompleted               AiAgentToolApprovalStatus = "completed"
	AiAgentToolApprovalStatusFailed                  AiAgentToolApprovalStatus = "failed"
	AiAgentToolApprovalStatusPartial                 AiAgentToolApprovalStatus = "partial"
)

var allowedAiAgentToolApprovalStatusEnumValues = []AiAgentToolApprovalStatus{
	AiAgentToolApprovalStatusPendingUserConfirmation,
	AiAgentToolApprovalStatusPolicyAutoApproved,
	AiAgentToolApprovalStatusBlocked,
	AiAgentToolApprovalStatusUserApproved,
	AiAgentToolApprovalStatusUserRejected,
	AiAgentToolApprovalStatusUserCancelled,
	AiAgentToolApprovalStatusApprovalExpired,
	AiAgentToolApprovalStatusExecuting,
	AiAgentToolApprovalStatusCompleted,
	AiAgentToolApprovalStatusFailed,
	AiAgentToolApprovalStatusPartial,
}

// GetAllowedValues returns the list of possible values.
func (v *AiAgentToolApprovalStatus) GetAllowedValues() []AiAgentToolApprovalStatus {
	return allowedAiAgentToolApprovalStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiAgentToolApprovalStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiAgentToolApprovalStatus(value)
	return nil
}

// NewAiAgentToolApprovalStatusFromValue returns a pointer to a valid AiAgentToolApprovalStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiAgentToolApprovalStatusFromValue(v string) (*AiAgentToolApprovalStatus, error) {
	ev := AiAgentToolApprovalStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiAgentToolApprovalStatus: valid values are %v", v, allowedAiAgentToolApprovalStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiAgentToolApprovalStatus) IsValid() bool {
	for _, existing := range allowedAiAgentToolApprovalStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiAgentToolApprovalStatus value.
func (v AiAgentToolApprovalStatus) Ptr() *AiAgentToolApprovalStatus {
	return &v
}
