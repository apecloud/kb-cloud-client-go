// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentScopeValidationStatus string

// List of AiAgentScopeValidationStatus.
const (
	AiAgentScopeValidationStatusAccepted AiAgentScopeValidationStatus = "accepted"
	AiAgentScopeValidationStatusRejected AiAgentScopeValidationStatus = "rejected"
)

var allowedAiAgentScopeValidationStatusEnumValues = []AiAgentScopeValidationStatus{
	AiAgentScopeValidationStatusAccepted,
	AiAgentScopeValidationStatusRejected,
}

// GetAllowedValues returns the list of possible values.
func (v *AiAgentScopeValidationStatus) GetAllowedValues() []AiAgentScopeValidationStatus {
	return allowedAiAgentScopeValidationStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiAgentScopeValidationStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiAgentScopeValidationStatus(value)
	return nil
}

// NewAiAgentScopeValidationStatusFromValue returns a pointer to a valid AiAgentScopeValidationStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiAgentScopeValidationStatusFromValue(v string) (*AiAgentScopeValidationStatus, error) {
	ev := AiAgentScopeValidationStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiAgentScopeValidationStatus: valid values are %v", v, allowedAiAgentScopeValidationStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiAgentScopeValidationStatus) IsValid() bool {
	for _, existing := range allowedAiAgentScopeValidationStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiAgentScopeValidationStatus value.
func (v AiAgentScopeValidationStatus) Ptr() *AiAgentScopeValidationStatus {
	return &v
}
