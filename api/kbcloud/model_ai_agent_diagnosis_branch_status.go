// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentDiagnosisBranchStatus string

// List of AiAgentDiagnosisBranchStatus.
const (
	AiAgentDiagnosisBranchStatusBuilding             AiAgentDiagnosisBranchStatus = "building"
	AiAgentDiagnosisBranchStatusReady                AiAgentDiagnosisBranchStatus = "ready"
	AiAgentDiagnosisBranchStatusFailed               AiAgentDiagnosisBranchStatus = "failed"
	AiAgentDiagnosisBranchStatusInsufficientEvidence AiAgentDiagnosisBranchStatus = "insufficient_evidence"
	AiAgentDiagnosisBranchStatusPermissionBlocked    AiAgentDiagnosisBranchStatus = "permission_blocked"
)

var allowedAiAgentDiagnosisBranchStatusEnumValues = []AiAgentDiagnosisBranchStatus{
	AiAgentDiagnosisBranchStatusBuilding,
	AiAgentDiagnosisBranchStatusReady,
	AiAgentDiagnosisBranchStatusFailed,
	AiAgentDiagnosisBranchStatusInsufficientEvidence,
	AiAgentDiagnosisBranchStatusPermissionBlocked,
}

// GetAllowedValues returns the list of possible values.
func (v *AiAgentDiagnosisBranchStatus) GetAllowedValues() []AiAgentDiagnosisBranchStatus {
	return allowedAiAgentDiagnosisBranchStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiAgentDiagnosisBranchStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiAgentDiagnosisBranchStatus(value)
	return nil
}

// NewAiAgentDiagnosisBranchStatusFromValue returns a pointer to a valid AiAgentDiagnosisBranchStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiAgentDiagnosisBranchStatusFromValue(v string) (*AiAgentDiagnosisBranchStatus, error) {
	ev := AiAgentDiagnosisBranchStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiAgentDiagnosisBranchStatus: valid values are %v", v, allowedAiAgentDiagnosisBranchStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiAgentDiagnosisBranchStatus) IsValid() bool {
	for _, existing := range allowedAiAgentDiagnosisBranchStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiAgentDiagnosisBranchStatus value.
func (v AiAgentDiagnosisBranchStatus) Ptr() *AiAgentDiagnosisBranchStatus {
	return &v
}
