// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentEvidenceDetailRedactionStatus string

// List of AiAgentEvidenceDetailRedactionStatus.
const (
	AiAgentEvidenceDetailRedactionStatusNotRequired AiAgentEvidenceDetailRedactionStatus = "not_required"
	AiAgentEvidenceDetailRedactionStatusRedacted    AiAgentEvidenceDetailRedactionStatus = "redacted"
	AiAgentEvidenceDetailRedactionStatusFailed      AiAgentEvidenceDetailRedactionStatus = "failed"
)

var allowedAiAgentEvidenceDetailRedactionStatusEnumValues = []AiAgentEvidenceDetailRedactionStatus{
	AiAgentEvidenceDetailRedactionStatusNotRequired,
	AiAgentEvidenceDetailRedactionStatusRedacted,
	AiAgentEvidenceDetailRedactionStatusFailed,
}

// GetAllowedValues returns the list of possible values.
func (v *AiAgentEvidenceDetailRedactionStatus) GetAllowedValues() []AiAgentEvidenceDetailRedactionStatus {
	return allowedAiAgentEvidenceDetailRedactionStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiAgentEvidenceDetailRedactionStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiAgentEvidenceDetailRedactionStatus(value)
	return nil
}

// NewAiAgentEvidenceDetailRedactionStatusFromValue returns a pointer to a valid AiAgentEvidenceDetailRedactionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiAgentEvidenceDetailRedactionStatusFromValue(v string) (*AiAgentEvidenceDetailRedactionStatus, error) {
	ev := AiAgentEvidenceDetailRedactionStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiAgentEvidenceDetailRedactionStatus: valid values are %v", v, allowedAiAgentEvidenceDetailRedactionStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiAgentEvidenceDetailRedactionStatus) IsValid() bool {
	for _, existing := range allowedAiAgentEvidenceDetailRedactionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiAgentEvidenceDetailRedactionStatus value.
func (v AiAgentEvidenceDetailRedactionStatus) Ptr() *AiAgentEvidenceDetailRedactionStatus {
	return &v
}
