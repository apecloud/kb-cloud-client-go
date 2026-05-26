// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentUIVisibility string

// List of AiAgentUIVisibility.
const (
	AiAgentUIVisibilityHidden  AiAgentUIVisibility = "hidden"
	AiAgentUIVisibilityInline  AiAgentUIVisibility = "inline"
	AiAgentUIVisibilityWorkLog AiAgentUIVisibility = "work_log"
	AiAgentUIVisibilityDebug   AiAgentUIVisibility = "debug"
)

var allowedAiAgentUIVisibilityEnumValues = []AiAgentUIVisibility{
	AiAgentUIVisibilityHidden,
	AiAgentUIVisibilityInline,
	AiAgentUIVisibilityWorkLog,
	AiAgentUIVisibilityDebug,
}

// GetAllowedValues returns the list of possible values.
func (v *AiAgentUIVisibility) GetAllowedValues() []AiAgentUIVisibility {
	return allowedAiAgentUIVisibilityEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiAgentUIVisibility) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiAgentUIVisibility(value)
	return nil
}

// NewAiAgentUIVisibilityFromValue returns a pointer to a valid AiAgentUIVisibility
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiAgentUIVisibilityFromValue(v string) (*AiAgentUIVisibility, error) {
	ev := AiAgentUIVisibility(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiAgentUIVisibility: valid values are %v", v, allowedAiAgentUIVisibilityEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiAgentUIVisibility) IsValid() bool {
	for _, existing := range allowedAiAgentUIVisibilityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiAgentUIVisibility value.
func (v AiAgentUIVisibility) Ptr() *AiAgentUIVisibility {
	return &v
}
