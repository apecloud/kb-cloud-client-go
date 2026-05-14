// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentActionPlanApprovalStatus string

// List of AiAgentActionPlanApprovalStatus.
const (
	AiAgentActionPlanApprovalStatusApproved  AiAgentActionPlanApprovalStatus = "approved"
	AiAgentActionPlanApprovalStatusRejected  AiAgentActionPlanApprovalStatus = "rejected"
	AiAgentActionPlanApprovalStatusExpired   AiAgentActionPlanApprovalStatus = "expired"
	AiAgentActionPlanApprovalStatusCancelled AiAgentActionPlanApprovalStatus = "cancelled"
)

var allowedAiAgentActionPlanApprovalStatusEnumValues = []AiAgentActionPlanApprovalStatus{
	AiAgentActionPlanApprovalStatusApproved,
	AiAgentActionPlanApprovalStatusRejected,
	AiAgentActionPlanApprovalStatusExpired,
	AiAgentActionPlanApprovalStatusCancelled,
}

// GetAllowedValues returns the list of possible values.
func (v *AiAgentActionPlanApprovalStatus) GetAllowedValues() []AiAgentActionPlanApprovalStatus {
	return allowedAiAgentActionPlanApprovalStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiAgentActionPlanApprovalStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiAgentActionPlanApprovalStatus(value)
	return nil
}

// NewAiAgentActionPlanApprovalStatusFromValue returns a pointer to a valid AiAgentActionPlanApprovalStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiAgentActionPlanApprovalStatusFromValue(v string) (*AiAgentActionPlanApprovalStatus, error) {
	ev := AiAgentActionPlanApprovalStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiAgentActionPlanApprovalStatus: valid values are %v", v, allowedAiAgentActionPlanApprovalStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiAgentActionPlanApprovalStatus) IsValid() bool {
	for _, existing := range allowedAiAgentActionPlanApprovalStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiAgentActionPlanApprovalStatus value.
func (v AiAgentActionPlanApprovalStatus) Ptr() *AiAgentActionPlanApprovalStatus {
	return &v
}
