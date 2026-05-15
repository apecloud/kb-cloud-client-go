// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentMessageOutcome string

// List of AiAgentMessageOutcome.
const (
	AiAgentMessageOutcomeAcceptedPending AiAgentMessageOutcome = "accepted_pending"
	AiAgentMessageOutcomeStartedNewRun   AiAgentMessageOutcome = "started_new_run"
	AiAgentMessageOutcomeRejected        AiAgentMessageOutcome = "rejected"
)

var allowedAiAgentMessageOutcomeEnumValues = []AiAgentMessageOutcome{
	AiAgentMessageOutcomeAcceptedPending,
	AiAgentMessageOutcomeStartedNewRun,
	AiAgentMessageOutcomeRejected,
}

// GetAllowedValues returns the list of possible values.
func (v *AiAgentMessageOutcome) GetAllowedValues() []AiAgentMessageOutcome {
	return allowedAiAgentMessageOutcomeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiAgentMessageOutcome) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiAgentMessageOutcome(value)
	return nil
}

// NewAiAgentMessageOutcomeFromValue returns a pointer to a valid AiAgentMessageOutcome
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiAgentMessageOutcomeFromValue(v string) (*AiAgentMessageOutcome, error) {
	ev := AiAgentMessageOutcome(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiAgentMessageOutcome: valid values are %v", v, allowedAiAgentMessageOutcomeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiAgentMessageOutcome) IsValid() bool {
	for _, existing := range allowedAiAgentMessageOutcomeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiAgentMessageOutcome value.
func (v AiAgentMessageOutcome) Ptr() *AiAgentMessageOutcome {
	return &v
}
