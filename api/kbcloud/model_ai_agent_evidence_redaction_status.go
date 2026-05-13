// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentEvidenceRedactionStatus string

// List of AiAgentEvidenceRedactionStatus.
const (
	AiAgentEvidenceRedactionStatusNotRequired AiAgentEvidenceRedactionStatus = "not_required"
	AiAgentEvidenceRedactionStatusRedacted    AiAgentEvidenceRedactionStatus = "redacted"
	AiAgentEvidenceRedactionStatusFailed      AiAgentEvidenceRedactionStatus = "failed"
)

var allowedAiAgentEvidenceRedactionStatusEnumValues = []AiAgentEvidenceRedactionStatus{
	AiAgentEvidenceRedactionStatusNotRequired,
	AiAgentEvidenceRedactionStatusRedacted,
	AiAgentEvidenceRedactionStatusFailed,
}

// GetAllowedValues returns the list of possible values.
func (v *AiAgentEvidenceRedactionStatus) GetAllowedValues() []AiAgentEvidenceRedactionStatus {
	return allowedAiAgentEvidenceRedactionStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiAgentEvidenceRedactionStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiAgentEvidenceRedactionStatus(value)
	return nil
}

// NewAiAgentEvidenceRedactionStatusFromValue returns a pointer to a valid AiAgentEvidenceRedactionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiAgentEvidenceRedactionStatusFromValue(v string) (*AiAgentEvidenceRedactionStatus, error) {
	ev := AiAgentEvidenceRedactionStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiAgentEvidenceRedactionStatus: valid values are %v", v, allowedAiAgentEvidenceRedactionStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiAgentEvidenceRedactionStatus) IsValid() bool {
	for _, existing := range allowedAiAgentEvidenceRedactionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiAgentEvidenceRedactionStatus value.
func (v AiAgentEvidenceRedactionStatus) Ptr() *AiAgentEvidenceRedactionStatus {
	return &v
}
