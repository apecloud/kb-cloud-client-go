// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentContextState string

// List of AiAgentContextState.
const (
	AiAgentContextStateGlobal       AiAgentContextState = "global"
	AiAgentContextStateClusterBound AiAgentContextState = "cluster_bound"
)

var allowedAiAgentContextStateEnumValues = []AiAgentContextState{
	AiAgentContextStateGlobal,
	AiAgentContextStateClusterBound,
}

// GetAllowedValues returns the list of possible values.
func (v *AiAgentContextState) GetAllowedValues() []AiAgentContextState {
	return allowedAiAgentContextStateEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiAgentContextState) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiAgentContextState(value)
	return nil
}

// NewAiAgentContextStateFromValue returns a pointer to a valid AiAgentContextState
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiAgentContextStateFromValue(v string) (*AiAgentContextState, error) {
	ev := AiAgentContextState(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiAgentContextState: valid values are %v", v, allowedAiAgentContextStateEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiAgentContextState) IsValid() bool {
	for _, existing := range allowedAiAgentContextStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiAgentContextState value.
func (v AiAgentContextState) Ptr() *AiAgentContextState {
	return &v
}
