// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentEventSource string

// List of AiAgentEventSource.
const (
	AiAgentEventSourceRuntime         AiAgentEventSource = "runtime"
	AiAgentEventSourceLlm             AiAgentEventSource = "llm"
	AiAgentEventSourceToolGateway     AiAgentEventSource = "tool_gateway"
	AiAgentEventSourceEvidenceBuilder AiAgentEventSource = "evidence_builder"
	AiAgentEventSourceResultBuilder   AiAgentEventSource = "result_builder"
	AiAgentEventSourceApprovalService AiAgentEventSource = "approval_service"
)

var allowedAiAgentEventSourceEnumValues = []AiAgentEventSource{
	AiAgentEventSourceRuntime,
	AiAgentEventSourceLlm,
	AiAgentEventSourceToolGateway,
	AiAgentEventSourceEvidenceBuilder,
	AiAgentEventSourceResultBuilder,
	AiAgentEventSourceApprovalService,
}

// GetAllowedValues returns the list of possible values.
func (v *AiAgentEventSource) GetAllowedValues() []AiAgentEventSource {
	return allowedAiAgentEventSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiAgentEventSource) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiAgentEventSource(value)
	return nil
}

// NewAiAgentEventSourceFromValue returns a pointer to a valid AiAgentEventSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiAgentEventSourceFromValue(v string) (*AiAgentEventSource, error) {
	ev := AiAgentEventSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiAgentEventSource: valid values are %v", v, allowedAiAgentEventSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiAgentEventSource) IsValid() bool {
	for _, existing := range allowedAiAgentEventSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiAgentEventSource value.
func (v AiAgentEventSource) Ptr() *AiAgentEventSource {
	return &v
}
