// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentCostClass string

// List of AiAgentCostClass.
const (
	AiAgentCostClassLow    AiAgentCostClass = "low"
	AiAgentCostClassMedium AiAgentCostClass = "medium"
	AiAgentCostClassHigh   AiAgentCostClass = "high"
)

var allowedAiAgentCostClassEnumValues = []AiAgentCostClass{
	AiAgentCostClassLow,
	AiAgentCostClassMedium,
	AiAgentCostClassHigh,
}

// GetAllowedValues returns the list of possible values.
func (v *AiAgentCostClass) GetAllowedValues() []AiAgentCostClass {
	return allowedAiAgentCostClassEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiAgentCostClass) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiAgentCostClass(value)
	return nil
}

// NewAiAgentCostClassFromValue returns a pointer to a valid AiAgentCostClass
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiAgentCostClassFromValue(v string) (*AiAgentCostClass, error) {
	ev := AiAgentCostClass(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiAgentCostClass: valid values are %v", v, allowedAiAgentCostClassEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiAgentCostClass) IsValid() bool {
	for _, existing := range allowedAiAgentCostClassEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiAgentCostClass value.
func (v AiAgentCostClass) Ptr() *AiAgentCostClass {
	return &v
}
