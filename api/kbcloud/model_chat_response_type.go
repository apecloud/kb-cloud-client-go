// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ChatResponseType 消息类型
type ChatResponseType string

// List of ChatResponseType.
const (
	ChatResponseTypeMessage            ChatResponseType = "Message"
	ChatResponseTypeGraph              ChatResponseType = "Graph"
	ChatResponseTypeFinish             ChatResponseType = "Finish"
	ChatResponseTypeDataInterpretation ChatResponseType = "DataInterpretation"
	ChatResponseTypeError              ChatResponseType = "Error"
	ChatResponseTypeThink              ChatResponseType = "Think"
)

var allowedChatResponseTypeEnumValues = []ChatResponseType{
	ChatResponseTypeMessage,
	ChatResponseTypeGraph,
	ChatResponseTypeFinish,
	ChatResponseTypeDataInterpretation,
	ChatResponseTypeError,
	ChatResponseTypeThink,
}

// GetAllowedValues returns the list of possible values.
func (v *ChatResponseType) GetAllowedValues() []ChatResponseType {
	return allowedChatResponseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ChatResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ChatResponseType(value)
	return nil
}

// NewChatResponseTypeFromValue returns a pointer to a valid ChatResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewChatResponseTypeFromValue(v string) (*ChatResponseType, error) {
	ev := ChatResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ChatResponseType: valid values are %v", v, allowedChatResponseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ChatResponseType) IsValid() bool {
	for _, existing := range allowedChatResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ChatResponseType value.
func (v ChatResponseType) Ptr() *ChatResponseType {
	return &v
}
