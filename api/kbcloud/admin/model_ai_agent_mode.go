// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentMode string

// List of AiAgentMode.
const (
	AiAgentModeGlobal           AiAgentMode = "global"
	AiAgentModeClusterDiagnosis AiAgentMode = "cluster_diagnosis"
)

var allowedAiAgentModeEnumValues = []AiAgentMode{
	AiAgentModeGlobal,
	AiAgentModeClusterDiagnosis,
}

// GetAllowedValues returns the list of possible values.
func (v *AiAgentMode) GetAllowedValues() []AiAgentMode {
	return allowedAiAgentModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiAgentMode) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiAgentMode(value)
	return nil
}

// NewAiAgentModeFromValue returns a pointer to a valid AiAgentMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiAgentModeFromValue(v string) (*AiAgentMode, error) {
	ev := AiAgentMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiAgentMode: valid values are %v", v, allowedAiAgentModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiAgentMode) IsValid() bool {
	for _, existing := range allowedAiAgentModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiAgentMode value.
func (v AiAgentMode) Ptr() *AiAgentMode {
	return &v
}
