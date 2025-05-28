// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// SenderType Type of the sender (user, assistant, or system)
type SenderType string

// List of SenderType.
const (
	SenderTypeUser      SenderType = "user"
	SenderTypeAssistant SenderType = "assistant"
	SenderTypeSystem    SenderType = "system"
)

var allowedSenderTypeEnumValues = []SenderType{
	SenderTypeUser,
	SenderTypeAssistant,
	SenderTypeSystem,
}

// GetAllowedValues returns the list of possible values.
func (v *SenderType) GetAllowedValues() []SenderType {
	return allowedSenderTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SenderType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SenderType(value)
	return nil
}

// NewSenderTypeFromValue returns a pointer to a valid SenderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSenderTypeFromValue(v string) (*SenderType, error) {
	ev := SenderType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SenderType: valid values are %v", v, allowedSenderTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SenderType) IsValid() bool {
	for _, existing := range allowedSenderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SenderType value.
func (v SenderType) Ptr() *SenderType {
	return &v
}
