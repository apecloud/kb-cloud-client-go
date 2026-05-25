// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type HermesAgentRuntimeStatusCode string

// List of HermesAgentRuntimeStatusCode.
const (
	HermesAgentRuntimeStatusCodeReady              HermesAgentRuntimeStatusCode = "ready"
	HermesAgentRuntimeStatusCodeNotConfigured      HermesAgentRuntimeStatusCode = "not_configured"
	HermesAgentRuntimeStatusCodeMissingModelConfig HermesAgentRuntimeStatusCode = "missing_model_config"
	HermesAgentRuntimeStatusCodeUnreachable        HermesAgentRuntimeStatusCode = "unreachable"
	HermesAgentRuntimeStatusCodeUnauthorized       HermesAgentRuntimeStatusCode = "unauthorized"
)

var allowedHermesAgentRuntimeStatusCodeEnumValues = []HermesAgentRuntimeStatusCode{
	HermesAgentRuntimeStatusCodeReady,
	HermesAgentRuntimeStatusCodeNotConfigured,
	HermesAgentRuntimeStatusCodeMissingModelConfig,
	HermesAgentRuntimeStatusCodeUnreachable,
	HermesAgentRuntimeStatusCodeUnauthorized,
}

// GetAllowedValues returns the list of possible values.
func (v *HermesAgentRuntimeStatusCode) GetAllowedValues() []HermesAgentRuntimeStatusCode {
	return allowedHermesAgentRuntimeStatusCodeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HermesAgentRuntimeStatusCode) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HermesAgentRuntimeStatusCode(value)
	return nil
}

// NewHermesAgentRuntimeStatusCodeFromValue returns a pointer to a valid HermesAgentRuntimeStatusCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHermesAgentRuntimeStatusCodeFromValue(v string) (*HermesAgentRuntimeStatusCode, error) {
	ev := HermesAgentRuntimeStatusCode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HermesAgentRuntimeStatusCode: valid values are %v", v, allowedHermesAgentRuntimeStatusCodeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HermesAgentRuntimeStatusCode) IsValid() bool {
	for _, existing := range allowedHermesAgentRuntimeStatusCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HermesAgentRuntimeStatusCode value.
func (v HermesAgentRuntimeStatusCode) Ptr() *HermesAgentRuntimeStatusCode {
	return &v
}
