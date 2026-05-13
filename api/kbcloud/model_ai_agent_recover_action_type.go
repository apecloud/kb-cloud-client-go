// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentRecoverActionType string

// List of AiAgentRecoverActionType.
const (
	AiAgentRecoverActionTypeRetry                 AiAgentRecoverActionType = "retry"
	AiAgentRecoverActionTypeNarrowScope           AiAgentRecoverActionType = "narrow_scope"
	AiAgentRecoverActionTypeRequestPermission     AiAgentRecoverActionType = "request_permission"
	AiAgentRecoverActionTypeAskUser               AiAgentRecoverActionType = "ask_user"
	AiAgentRecoverActionTypeViewReport            AiAgentRecoverActionType = "view_report"
	AiAgentRecoverActionTypeRegenerateReport      AiAgentRecoverActionType = "regenerate_report"
	AiAgentRecoverActionTypeContinueInvestigation AiAgentRecoverActionType = "continue_investigation"
	AiAgentRecoverActionTypeNon                   AiAgentRecoverActionType = "none"
)

var allowedAiAgentRecoverActionTypeEnumValues = []AiAgentRecoverActionType{
	AiAgentRecoverActionTypeRetry,
	AiAgentRecoverActionTypeNarrowScope,
	AiAgentRecoverActionTypeRequestPermission,
	AiAgentRecoverActionTypeAskUser,
	AiAgentRecoverActionTypeViewReport,
	AiAgentRecoverActionTypeRegenerateReport,
	AiAgentRecoverActionTypeContinueInvestigation,
	AiAgentRecoverActionTypeNon,
}

// GetAllowedValues returns the list of possible values.
func (v *AiAgentRecoverActionType) GetAllowedValues() []AiAgentRecoverActionType {
	return allowedAiAgentRecoverActionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiAgentRecoverActionType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiAgentRecoverActionType(value)
	return nil
}

// NewAiAgentRecoverActionTypeFromValue returns a pointer to a valid AiAgentRecoverActionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiAgentRecoverActionTypeFromValue(v string) (*AiAgentRecoverActionType, error) {
	ev := AiAgentRecoverActionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiAgentRecoverActionType: valid values are %v", v, allowedAiAgentRecoverActionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiAgentRecoverActionType) IsValid() bool {
	for _, existing := range allowedAiAgentRecoverActionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiAgentRecoverActionType value.
func (v AiAgentRecoverActionType) Ptr() *AiAgentRecoverActionType {
	return &v
}
