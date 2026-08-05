// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentProfile string

// List of AiAgentProfile.
const (
	AiAgentProfileCloudAiAgent AiAgentProfile = "cloud-ai-agent"
)

var allowedAiAgentProfileEnumValues = []AiAgentProfile{
	AiAgentProfileCloudAiAgent,
}

// GetAllowedValues returns the list of possible values.
func (v *AiAgentProfile) GetAllowedValues() []AiAgentProfile {
	return allowedAiAgentProfileEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiAgentProfile) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiAgentProfile(value)
	return nil
}

// NewAiAgentProfileFromValue returns a pointer to a valid AiAgentProfile
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiAgentProfileFromValue(v string) (*AiAgentProfile, error) {
	ev := AiAgentProfile(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiAgentProfile: valid values are %v", v, allowedAiAgentProfileEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiAgentProfile) IsValid() bool {
	for _, existing := range allowedAiAgentProfileEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiAgentProfile value.
func (v AiAgentProfile) Ptr() *AiAgentProfile {
	return &v
}
