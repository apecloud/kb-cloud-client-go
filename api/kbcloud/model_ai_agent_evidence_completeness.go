// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentEvidenceCompleteness string

// List of AiAgentEvidenceCompleteness.
const (
	AiAgentEvidenceCompletenessComplete          AiAgentEvidenceCompleteness = "complete"
	AiAgentEvidenceCompletenessPartialDegraded   AiAgentEvidenceCompleteness = "partial_degraded"
	AiAgentEvidenceCompletenessInsufficient      AiAgentEvidenceCompleteness = "insufficient"
	AiAgentEvidenceCompletenessPermissionBlocked AiAgentEvidenceCompleteness = "permission_blocked"
)

var allowedAiAgentEvidenceCompletenessEnumValues = []AiAgentEvidenceCompleteness{
	AiAgentEvidenceCompletenessComplete,
	AiAgentEvidenceCompletenessPartialDegraded,
	AiAgentEvidenceCompletenessInsufficient,
	AiAgentEvidenceCompletenessPermissionBlocked,
}

// GetAllowedValues returns the list of possible values.
func (v *AiAgentEvidenceCompleteness) GetAllowedValues() []AiAgentEvidenceCompleteness {
	return allowedAiAgentEvidenceCompletenessEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiAgentEvidenceCompleteness) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiAgentEvidenceCompleteness(value)
	return nil
}

// NewAiAgentEvidenceCompletenessFromValue returns a pointer to a valid AiAgentEvidenceCompleteness
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiAgentEvidenceCompletenessFromValue(v string) (*AiAgentEvidenceCompleteness, error) {
	ev := AiAgentEvidenceCompleteness(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiAgentEvidenceCompleteness: valid values are %v", v, allowedAiAgentEvidenceCompletenessEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiAgentEvidenceCompleteness) IsValid() bool {
	for _, existing := range allowedAiAgentEvidenceCompletenessEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiAgentEvidenceCompleteness value.
func (v AiAgentEvidenceCompleteness) Ptr() *AiAgentEvidenceCompleteness {
	return &v
}
