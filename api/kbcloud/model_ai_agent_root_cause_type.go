// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentRootCauseType string

// List of AiAgentRootCauseType.
const (
	AiAgentRootCauseTypeConfirmed AiAgentRootCauseType = "confirmed"
	AiAgentRootCauseTypeCandidate AiAgentRootCauseType = "candidate"
	AiAgentRootCauseTypeNon       AiAgentRootCauseType = "none"
)

var allowedAiAgentRootCauseTypeEnumValues = []AiAgentRootCauseType{
	AiAgentRootCauseTypeConfirmed,
	AiAgentRootCauseTypeCandidate,
	AiAgentRootCauseTypeNon,
}

// GetAllowedValues returns the list of possible values.
func (v *AiAgentRootCauseType) GetAllowedValues() []AiAgentRootCauseType {
	return allowedAiAgentRootCauseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiAgentRootCauseType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiAgentRootCauseType(value)
	return nil
}

// NewAiAgentRootCauseTypeFromValue returns a pointer to a valid AiAgentRootCauseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiAgentRootCauseTypeFromValue(v string) (*AiAgentRootCauseType, error) {
	ev := AiAgentRootCauseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiAgentRootCauseType: valid values are %v", v, allowedAiAgentRootCauseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiAgentRootCauseType) IsValid() bool {
	for _, existing := range allowedAiAgentRootCauseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiAgentRootCauseType value.
func (v AiAgentRootCauseType) Ptr() *AiAgentRootCauseType {
	return &v
}
