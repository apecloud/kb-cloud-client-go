// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentEvidenceStatus string

// List of AiAgentEvidenceStatus.
const (
	AiAgentEvidenceStatusValid             AiAgentEvidenceStatus = "valid"
	AiAgentEvidenceStatusPartial           AiAgentEvidenceStatus = "partial"
	AiAgentEvidenceStatusMissing           AiAgentEvidenceStatus = "missing"
	AiAgentEvidenceStatusPermissionBlocked AiAgentEvidenceStatus = "permission_blocked"
	AiAgentEvidenceStatusFailed            AiAgentEvidenceStatus = "failed"
)

var allowedAiAgentEvidenceStatusEnumValues = []AiAgentEvidenceStatus{
	AiAgentEvidenceStatusValid,
	AiAgentEvidenceStatusPartial,
	AiAgentEvidenceStatusMissing,
	AiAgentEvidenceStatusPermissionBlocked,
	AiAgentEvidenceStatusFailed,
}

// GetAllowedValues returns the list of possible values.
func (v *AiAgentEvidenceStatus) GetAllowedValues() []AiAgentEvidenceStatus {
	return allowedAiAgentEvidenceStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiAgentEvidenceStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiAgentEvidenceStatus(value)
	return nil
}

// NewAiAgentEvidenceStatusFromValue returns a pointer to a valid AiAgentEvidenceStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiAgentEvidenceStatusFromValue(v string) (*AiAgentEvidenceStatus, error) {
	ev := AiAgentEvidenceStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiAgentEvidenceStatus: valid values are %v", v, allowedAiAgentEvidenceStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiAgentEvidenceStatus) IsValid() bool {
	for _, existing := range allowedAiAgentEvidenceStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiAgentEvidenceStatus value.
func (v AiAgentEvidenceStatus) Ptr() *AiAgentEvidenceStatus {
	return &v
}
