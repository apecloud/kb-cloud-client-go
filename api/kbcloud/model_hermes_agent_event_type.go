// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type HermesAgentEventType string

// List of HermesAgentEventType.
const (
	HermesAgentEventTypeRunStarted         HermesAgentEventType = "run_started"
	HermesAgentEventTypeAssistantDelta     HermesAgentEventType = "assistant_delta"
	HermesAgentEventTypeReasoningAvailable HermesAgentEventType = "reasoning_available"
	HermesAgentEventTypeToolStarted        HermesAgentEventType = "tool_started"
	HermesAgentEventTypeToolCompleted      HermesAgentEventType = "tool_completed"
	HermesAgentEventTypeApprovalRequested  HermesAgentEventType = "approval_requested"
	HermesAgentEventTypeApprovalDecided    HermesAgentEventType = "approval_decided"
	HermesAgentEventTypeRunCompleted       HermesAgentEventType = "run_completed"
	HermesAgentEventTypeRunFailed          HermesAgentEventType = "run_failed"
	HermesAgentEventTypeRunCancelled       HermesAgentEventType = "run_cancelled"
	HermesAgentEventTypeHermesEvent        HermesAgentEventType = "hermes_event"
)

var allowedHermesAgentEventTypeEnumValues = []HermesAgentEventType{
	HermesAgentEventTypeRunStarted,
	HermesAgentEventTypeAssistantDelta,
	HermesAgentEventTypeReasoningAvailable,
	HermesAgentEventTypeToolStarted,
	HermesAgentEventTypeToolCompleted,
	HermesAgentEventTypeApprovalRequested,
	HermesAgentEventTypeApprovalDecided,
	HermesAgentEventTypeRunCompleted,
	HermesAgentEventTypeRunFailed,
	HermesAgentEventTypeRunCancelled,
	HermesAgentEventTypeHermesEvent,
}

// GetAllowedValues returns the list of possible values.
func (v *HermesAgentEventType) GetAllowedValues() []HermesAgentEventType {
	return allowedHermesAgentEventTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HermesAgentEventType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HermesAgentEventType(value)
	return nil
}

// NewHermesAgentEventTypeFromValue returns a pointer to a valid HermesAgentEventType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHermesAgentEventTypeFromValue(v string) (*HermesAgentEventType, error) {
	ev := HermesAgentEventType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HermesAgentEventType: valid values are %v", v, allowedHermesAgentEventTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HermesAgentEventType) IsValid() bool {
	for _, existing := range allowedHermesAgentEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HermesAgentEventType value.
func (v HermesAgentEventType) Ptr() *HermesAgentEventType {
	return &v
}
