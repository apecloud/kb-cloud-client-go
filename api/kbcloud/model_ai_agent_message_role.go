// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentMessageRole string

// List of AiAgentMessageRole.
const (
	AiAgentMessageRoleUser      AiAgentMessageRole = "user"
	AiAgentMessageRoleAssistant AiAgentMessageRole = "assistant"
	AiAgentMessageRoleSystem    AiAgentMessageRole = "system"
)

var allowedAiAgentMessageRoleEnumValues = []AiAgentMessageRole{
	AiAgentMessageRoleUser,
	AiAgentMessageRoleAssistant,
	AiAgentMessageRoleSystem,
}

// GetAllowedValues returns the list of possible values.
func (v *AiAgentMessageRole) GetAllowedValues() []AiAgentMessageRole {
	return allowedAiAgentMessageRoleEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiAgentMessageRole) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiAgentMessageRole(value)
	return nil
}

// NewAiAgentMessageRoleFromValue returns a pointer to a valid AiAgentMessageRole
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiAgentMessageRoleFromValue(v string) (*AiAgentMessageRole, error) {
	ev := AiAgentMessageRole(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiAgentMessageRole: valid values are %v", v, allowedAiAgentMessageRoleEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiAgentMessageRole) IsValid() bool {
	for _, existing := range allowedAiAgentMessageRoleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiAgentMessageRole value.
func (v AiAgentMessageRole) Ptr() *AiAgentMessageRole {
	return &v
}
