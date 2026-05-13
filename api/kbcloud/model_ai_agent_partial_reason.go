// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentPartialReason string

// List of AiAgentPartialReason.
const (
	AiAgentPartialReasonSomeRequirementsMissing    AiAgentPartialReason = "some_requirements_missing"
	AiAgentPartialReasonSomeSourcesFailed          AiAgentPartialReason = "some_sources_failed"
	AiAgentPartialReasonTruncatedByLimit           AiAgentPartialReason = "truncated_by_limit"
	AiAgentPartialReasonPermissionPartiallyBlocked AiAgentPartialReason = "permission_partially_blocked"
	AiAgentPartialReasonTimeoutPartial             AiAgentPartialReason = "timeout_partial"
)

var allowedAiAgentPartialReasonEnumValues = []AiAgentPartialReason{
	AiAgentPartialReasonSomeRequirementsMissing,
	AiAgentPartialReasonSomeSourcesFailed,
	AiAgentPartialReasonTruncatedByLimit,
	AiAgentPartialReasonPermissionPartiallyBlocked,
	AiAgentPartialReasonTimeoutPartial,
}

// GetAllowedValues returns the list of possible values.
func (v *AiAgentPartialReason) GetAllowedValues() []AiAgentPartialReason {
	return allowedAiAgentPartialReasonEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiAgentPartialReason) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiAgentPartialReason(value)
	return nil
}

// NewAiAgentPartialReasonFromValue returns a pointer to a valid AiAgentPartialReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiAgentPartialReasonFromValue(v string) (*AiAgentPartialReason, error) {
	ev := AiAgentPartialReason(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiAgentPartialReason: valid values are %v", v, allowedAiAgentPartialReasonEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiAgentPartialReason) IsValid() bool {
	for _, existing := range allowedAiAgentPartialReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiAgentPartialReason value.
func (v AiAgentPartialReason) Ptr() *AiAgentPartialReason {
	return &v
}
