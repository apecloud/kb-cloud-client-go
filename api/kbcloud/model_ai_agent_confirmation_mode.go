// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentConfirmationMode string

// List of AiAgentConfirmationMode.
const (
	AiAgentConfirmationModeButton   AiAgentConfirmationMode = "button"
	AiAgentConfirmationModeTypeText AiAgentConfirmationMode = "type_text"
)

var allowedAiAgentConfirmationModeEnumValues = []AiAgentConfirmationMode{
	AiAgentConfirmationModeButton,
	AiAgentConfirmationModeTypeText,
}

// GetAllowedValues returns the list of possible values.
func (v *AiAgentConfirmationMode) GetAllowedValues() []AiAgentConfirmationMode {
	return allowedAiAgentConfirmationModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiAgentConfirmationMode) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiAgentConfirmationMode(value)
	return nil
}

// NewAiAgentConfirmationModeFromValue returns a pointer to a valid AiAgentConfirmationMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiAgentConfirmationModeFromValue(v string) (*AiAgentConfirmationMode, error) {
	ev := AiAgentConfirmationMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiAgentConfirmationMode: valid values are %v", v, allowedAiAgentConfirmationModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiAgentConfirmationMode) IsValid() bool {
	for _, existing := range allowedAiAgentConfirmationModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiAgentConfirmationMode value.
func (v AiAgentConfirmationMode) Ptr() *AiAgentConfirmationMode {
	return &v
}
