// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentFinalStatus string

// List of AiAgentFinalStatus.
const (
	AiAgentFinalStatusCompleted            AiAgentFinalStatus = "completed"
	AiAgentFinalStatusInsufficientEvidence AiAgentFinalStatus = "insufficient_evidence"
	AiAgentFinalStatusBlockedByApproval    AiAgentFinalStatus = "blocked_by_approval"
	AiAgentFinalStatusFailed               AiAgentFinalStatus = "failed"
	AiAgentFinalStatusCancelled            AiAgentFinalStatus = "cancelled"
)

var allowedAiAgentFinalStatusEnumValues = []AiAgentFinalStatus{
	AiAgentFinalStatusCompleted,
	AiAgentFinalStatusInsufficientEvidence,
	AiAgentFinalStatusBlockedByApproval,
	AiAgentFinalStatusFailed,
	AiAgentFinalStatusCancelled,
}

// GetAllowedValues returns the list of possible values.
func (v *AiAgentFinalStatus) GetAllowedValues() []AiAgentFinalStatus {
	return allowedAiAgentFinalStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiAgentFinalStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiAgentFinalStatus(value)
	return nil
}

// NewAiAgentFinalStatusFromValue returns a pointer to a valid AiAgentFinalStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiAgentFinalStatusFromValue(v string) (*AiAgentFinalStatus, error) {
	ev := AiAgentFinalStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiAgentFinalStatus: valid values are %v", v, allowedAiAgentFinalStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiAgentFinalStatus) IsValid() bool {
	for _, existing := range allowedAiAgentFinalStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiAgentFinalStatus value.
func (v AiAgentFinalStatus) Ptr() *AiAgentFinalStatus {
	return &v
}
