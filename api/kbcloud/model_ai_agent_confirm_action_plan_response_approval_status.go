// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentConfirmActionPlanResponseApprovalStatus string

// List of AiAgentConfirmActionPlanResponseApprovalStatus.
const (
	AiAgentConfirmActionPlanResponseApprovalStatusApproved  AiAgentConfirmActionPlanResponseApprovalStatus = "approved"
	AiAgentConfirmActionPlanResponseApprovalStatusRejected  AiAgentConfirmActionPlanResponseApprovalStatus = "rejected"
	AiAgentConfirmActionPlanResponseApprovalStatusExpired   AiAgentConfirmActionPlanResponseApprovalStatus = "expired"
	AiAgentConfirmActionPlanResponseApprovalStatusCancelled AiAgentConfirmActionPlanResponseApprovalStatus = "cancelled"
)

var allowedAiAgentConfirmActionPlanResponseApprovalStatusEnumValues = []AiAgentConfirmActionPlanResponseApprovalStatus{
	AiAgentConfirmActionPlanResponseApprovalStatusApproved,
	AiAgentConfirmActionPlanResponseApprovalStatusRejected,
	AiAgentConfirmActionPlanResponseApprovalStatusExpired,
	AiAgentConfirmActionPlanResponseApprovalStatusCancelled,
}

// GetAllowedValues returns the list of possible values.
func (v *AiAgentConfirmActionPlanResponseApprovalStatus) GetAllowedValues() []AiAgentConfirmActionPlanResponseApprovalStatus {
	return allowedAiAgentConfirmActionPlanResponseApprovalStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiAgentConfirmActionPlanResponseApprovalStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiAgentConfirmActionPlanResponseApprovalStatus(value)
	return nil
}

// NewAiAgentConfirmActionPlanResponseApprovalStatusFromValue returns a pointer to a valid AiAgentConfirmActionPlanResponseApprovalStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiAgentConfirmActionPlanResponseApprovalStatusFromValue(v string) (*AiAgentConfirmActionPlanResponseApprovalStatus, error) {
	ev := AiAgentConfirmActionPlanResponseApprovalStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiAgentConfirmActionPlanResponseApprovalStatus: valid values are %v", v, allowedAiAgentConfirmActionPlanResponseApprovalStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiAgentConfirmActionPlanResponseApprovalStatus) IsValid() bool {
	for _, existing := range allowedAiAgentConfirmActionPlanResponseApprovalStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiAgentConfirmActionPlanResponseApprovalStatus value.
func (v AiAgentConfirmActionPlanResponseApprovalStatus) Ptr() *AiAgentConfirmActionPlanResponseApprovalStatus {
	return &v
}
