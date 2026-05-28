// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentRuntimeStatusCode string

// List of AiAgentRuntimeStatusCode.
const (
	AiAgentRuntimeStatusCodeReady              AiAgentRuntimeStatusCode = "ready"
	AiAgentRuntimeStatusCodeMissingModelConfig AiAgentRuntimeStatusCode = "missing_model_config"
	AiAgentRuntimeStatusCodeUnreachable        AiAgentRuntimeStatusCode = "unreachable"
	AiAgentRuntimeStatusCodeUnauthorized       AiAgentRuntimeStatusCode = "unauthorized"
	AiAgentRuntimeStatusCodeNotConfigured      AiAgentRuntimeStatusCode = "not_configured"
)

var allowedAiAgentRuntimeStatusCodeEnumValues = []AiAgentRuntimeStatusCode{
	AiAgentRuntimeStatusCodeReady,
	AiAgentRuntimeStatusCodeMissingModelConfig,
	AiAgentRuntimeStatusCodeUnreachable,
	AiAgentRuntimeStatusCodeUnauthorized,
	AiAgentRuntimeStatusCodeNotConfigured,
}

// GetAllowedValues returns the list of possible values.
func (v *AiAgentRuntimeStatusCode) GetAllowedValues() []AiAgentRuntimeStatusCode {
	return allowedAiAgentRuntimeStatusCodeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiAgentRuntimeStatusCode) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiAgentRuntimeStatusCode(value)
	return nil
}

// NewAiAgentRuntimeStatusCodeFromValue returns a pointer to a valid AiAgentRuntimeStatusCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiAgentRuntimeStatusCodeFromValue(v string) (*AiAgentRuntimeStatusCode, error) {
	ev := AiAgentRuntimeStatusCode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiAgentRuntimeStatusCode: valid values are %v", v, allowedAiAgentRuntimeStatusCodeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiAgentRuntimeStatusCode) IsValid() bool {
	for _, existing := range allowedAiAgentRuntimeStatusCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiAgentRuntimeStatusCode value.
func (v AiAgentRuntimeStatusCode) Ptr() *AiAgentRuntimeStatusCode {
	return &v
}
