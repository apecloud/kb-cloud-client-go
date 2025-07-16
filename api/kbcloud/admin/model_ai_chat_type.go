// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// AIChatType Query type
type AIChatType string

// List of AIChatType.
const (
	AIChatTypeChatbi  AIChatType = "chatbi"
	AIChatTypeChatops AIChatType = "chatops"
)

var allowedAIChatTypeEnumValues = []AIChatType{
	AIChatTypeChatbi,
	AIChatTypeChatops,
}

// GetAllowedValues returns the list of possible values.
func (v *AIChatType) GetAllowedValues() []AIChatType {
	return allowedAIChatTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AIChatType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AIChatType(value)
	return nil
}

// NewAIChatTypeFromValue returns a pointer to a valid AIChatType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAIChatTypeFromValue(v string) (*AIChatType, error) {
	ev := AIChatType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AIChatType: valid values are %v", v, allowedAIChatTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AIChatType) IsValid() bool {
	for _, existing := range allowedAIChatTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AIChatType value.
func (v AIChatType) Ptr() *AIChatType {
	return &v
}
