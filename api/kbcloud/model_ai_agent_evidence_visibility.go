// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentEvidenceVisibility string

// List of AiAgentEvidenceVisibility.
const (
	AiAgentEvidenceVisibilityUser     AiAgentEvidenceVisibility = "user"
	AiAgentEvidenceVisibilityInternal AiAgentEvidenceVisibility = "internal"
	AiAgentEvidenceVisibilityDebug    AiAgentEvidenceVisibility = "debug"
)

var allowedAiAgentEvidenceVisibilityEnumValues = []AiAgentEvidenceVisibility{
	AiAgentEvidenceVisibilityUser,
	AiAgentEvidenceVisibilityInternal,
	AiAgentEvidenceVisibilityDebug,
}

// GetAllowedValues returns the list of possible values.
func (v *AiAgentEvidenceVisibility) GetAllowedValues() []AiAgentEvidenceVisibility {
	return allowedAiAgentEvidenceVisibilityEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiAgentEvidenceVisibility) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiAgentEvidenceVisibility(value)
	return nil
}

// NewAiAgentEvidenceVisibilityFromValue returns a pointer to a valid AiAgentEvidenceVisibility
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiAgentEvidenceVisibilityFromValue(v string) (*AiAgentEvidenceVisibility, error) {
	ev := AiAgentEvidenceVisibility(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiAgentEvidenceVisibility: valid values are %v", v, allowedAiAgentEvidenceVisibilityEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiAgentEvidenceVisibility) IsValid() bool {
	for _, existing := range allowedAiAgentEvidenceVisibilityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiAgentEvidenceVisibility value.
func (v AiAgentEvidenceVisibility) Ptr() *AiAgentEvidenceVisibility {
	return &v
}
