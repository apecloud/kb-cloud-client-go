// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentToolResponseStatus string

// List of AiAgentToolResponseStatus.
const (
	AiAgentToolResponseStatusCompleted         AiAgentToolResponseStatus = "completed"
	AiAgentToolResponseStatusPartial           AiAgentToolResponseStatus = "partial"
	AiAgentToolResponseStatusFailed            AiAgentToolResponseStatus = "failed"
	AiAgentToolResponseStatusPermissionBlocked AiAgentToolResponseStatus = "permission_blocked"
)

var allowedAiAgentToolResponseStatusEnumValues = []AiAgentToolResponseStatus{
	AiAgentToolResponseStatusCompleted,
	AiAgentToolResponseStatusPartial,
	AiAgentToolResponseStatusFailed,
	AiAgentToolResponseStatusPermissionBlocked,
}

// GetAllowedValues returns the list of possible values.
func (v *AiAgentToolResponseStatus) GetAllowedValues() []AiAgentToolResponseStatus {
	return allowedAiAgentToolResponseStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiAgentToolResponseStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiAgentToolResponseStatus(value)
	return nil
}

// NewAiAgentToolResponseStatusFromValue returns a pointer to a valid AiAgentToolResponseStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiAgentToolResponseStatusFromValue(v string) (*AiAgentToolResponseStatus, error) {
	ev := AiAgentToolResponseStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiAgentToolResponseStatus: valid values are %v", v, allowedAiAgentToolResponseStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiAgentToolResponseStatus) IsValid() bool {
	for _, existing := range allowedAiAgentToolResponseStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiAgentToolResponseStatus value.
func (v AiAgentToolResponseStatus) Ptr() *AiAgentToolResponseStatus {
	return &v
}
