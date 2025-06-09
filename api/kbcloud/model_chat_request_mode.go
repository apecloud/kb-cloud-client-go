// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ChatRequestMode Chat interaction mode
type ChatRequestMode string

// List of ChatRequestMode.
const (
	ChatRequestModeInteractive     ChatRequestMode = "interactive"
	ChatRequestModeAutomaticReport ChatRequestMode = "automatic_report"
)

var allowedChatRequestModeEnumValues = []ChatRequestMode{
	ChatRequestModeInteractive,
	ChatRequestModeAutomaticReport,
}

// GetAllowedValues returns the list of possible values.
func (v *ChatRequestMode) GetAllowedValues() []ChatRequestMode {
	return allowedChatRequestModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ChatRequestMode) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ChatRequestMode(value)
	return nil
}

// NewChatRequestModeFromValue returns a pointer to a valid ChatRequestMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewChatRequestModeFromValue(v string) (*ChatRequestMode, error) {
	ev := ChatRequestMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ChatRequestMode: valid values are %v", v, allowedChatRequestModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ChatRequestMode) IsValid() bool {
	for _, existing := range allowedChatRequestModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ChatRequestMode value.
func (v ChatRequestMode) Ptr() *ChatRequestMode {
	return &v
}
