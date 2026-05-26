// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentEventStatus string

// List of AiAgentEventStatus.
const (
	AiAgentEventStatusPending           AiAgentEventStatus = "pending"
	AiAgentEventStatusActive            AiAgentEventStatus = "active"
	AiAgentEventStatusCompleted         AiAgentEventStatus = "completed"
	AiAgentEventStatusFailed            AiAgentEventStatus = "failed"
	AiAgentEventStatusSkipped           AiAgentEventStatus = "skipped"
	AiAgentEventStatusPermissionBlocked AiAgentEventStatus = "permission_blocked"
	AiAgentEventStatusNotApplicable     AiAgentEventStatus = "not_applicable"
)

var allowedAiAgentEventStatusEnumValues = []AiAgentEventStatus{
	AiAgentEventStatusPending,
	AiAgentEventStatusActive,
	AiAgentEventStatusCompleted,
	AiAgentEventStatusFailed,
	AiAgentEventStatusSkipped,
	AiAgentEventStatusPermissionBlocked,
	AiAgentEventStatusNotApplicable,
}

// GetAllowedValues returns the list of possible values.
func (v *AiAgentEventStatus) GetAllowedValues() []AiAgentEventStatus {
	return allowedAiAgentEventStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiAgentEventStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiAgentEventStatus(value)
	return nil
}

// NewAiAgentEventStatusFromValue returns a pointer to a valid AiAgentEventStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiAgentEventStatusFromValue(v string) (*AiAgentEventStatus, error) {
	ev := AiAgentEventStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiAgentEventStatus: valid values are %v", v, allowedAiAgentEventStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiAgentEventStatus) IsValid() bool {
	for _, existing := range allowedAiAgentEventStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiAgentEventStatus value.
func (v AiAgentEventStatus) Ptr() *AiAgentEventStatus {
	return &v
}
