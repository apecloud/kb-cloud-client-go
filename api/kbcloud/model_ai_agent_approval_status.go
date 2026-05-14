// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentApprovalStatus string

// List of AiAgentApprovalStatus.
const (
	AiAgentApprovalStatusNotRequired AiAgentApprovalStatus = "not_required"
	AiAgentApprovalStatusPending     AiAgentApprovalStatus = "pending"
	AiAgentApprovalStatusApproved    AiAgentApprovalStatus = "approved"
	AiAgentApprovalStatusRejected    AiAgentApprovalStatus = "rejected"
	AiAgentApprovalStatusExpired     AiAgentApprovalStatus = "expired"
	AiAgentApprovalStatusCancelled   AiAgentApprovalStatus = "cancelled"
)

var allowedAiAgentApprovalStatusEnumValues = []AiAgentApprovalStatus{
	AiAgentApprovalStatusNotRequired,
	AiAgentApprovalStatusPending,
	AiAgentApprovalStatusApproved,
	AiAgentApprovalStatusRejected,
	AiAgentApprovalStatusExpired,
	AiAgentApprovalStatusCancelled,
}

// GetAllowedValues returns the list of possible values.
func (v *AiAgentApprovalStatus) GetAllowedValues() []AiAgentApprovalStatus {
	return allowedAiAgentApprovalStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiAgentApprovalStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiAgentApprovalStatus(value)
	return nil
}

// NewAiAgentApprovalStatusFromValue returns a pointer to a valid AiAgentApprovalStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiAgentApprovalStatusFromValue(v string) (*AiAgentApprovalStatus, error) {
	ev := AiAgentApprovalStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiAgentApprovalStatus: valid values are %v", v, allowedAiAgentApprovalStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiAgentApprovalStatus) IsValid() bool {
	for _, existing := range allowedAiAgentApprovalStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiAgentApprovalStatus value.
func (v AiAgentApprovalStatus) Ptr() *AiAgentApprovalStatus {
	return &v
}
