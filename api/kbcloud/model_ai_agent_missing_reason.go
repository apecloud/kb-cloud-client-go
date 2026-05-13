// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentMissingReason string

// List of AiAgentMissingReason.
const (
	AiAgentMissingReasonNotFound         AiAgentMissingReason = "not_found"
	AiAgentMissingReasonPermissionDenied AiAgentMissingReason = "permission_denied"
	AiAgentMissingReasonToolFailed       AiAgentMissingReason = "tool_failed"
	AiAgentMissingReasonTimeout          AiAgentMissingReason = "timeout"
	AiAgentMissingReasonNotConfigured    AiAgentMissingReason = "not_configured"
	AiAgentMissingReasonOutOfScope       AiAgentMissingReason = "out_of_scope"
	AiAgentMissingReasonBudgetExceeded   AiAgentMissingReason = "budget_exceeded"
)

var allowedAiAgentMissingReasonEnumValues = []AiAgentMissingReason{
	AiAgentMissingReasonNotFound,
	AiAgentMissingReasonPermissionDenied,
	AiAgentMissingReasonToolFailed,
	AiAgentMissingReasonTimeout,
	AiAgentMissingReasonNotConfigured,
	AiAgentMissingReasonOutOfScope,
	AiAgentMissingReasonBudgetExceeded,
}

// GetAllowedValues returns the list of possible values.
func (v *AiAgentMissingReason) GetAllowedValues() []AiAgentMissingReason {
	return allowedAiAgentMissingReasonEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiAgentMissingReason) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiAgentMissingReason(value)
	return nil
}

// NewAiAgentMissingReasonFromValue returns a pointer to a valid AiAgentMissingReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiAgentMissingReasonFromValue(v string) (*AiAgentMissingReason, error) {
	ev := AiAgentMissingReason(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiAgentMissingReason: valid values are %v", v, allowedAiAgentMissingReasonEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiAgentMissingReason) IsValid() bool {
	for _, existing := range allowedAiAgentMissingReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiAgentMissingReason value.
func (v AiAgentMissingReason) Ptr() *AiAgentMissingReason {
	return &v
}
