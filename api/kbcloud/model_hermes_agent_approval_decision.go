// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type HermesAgentApprovalDecision string

// List of HermesAgentApprovalDecision.
const (
	HermesAgentApprovalDecisionApprove HermesAgentApprovalDecision = "approve"
	HermesAgentApprovalDecisionReject  HermesAgentApprovalDecision = "reject"
)

var allowedHermesAgentApprovalDecisionEnumValues = []HermesAgentApprovalDecision{
	HermesAgentApprovalDecisionApprove,
	HermesAgentApprovalDecisionReject,
}

// GetAllowedValues returns the list of possible values.
func (v *HermesAgentApprovalDecision) GetAllowedValues() []HermesAgentApprovalDecision {
	return allowedHermesAgentApprovalDecisionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HermesAgentApprovalDecision) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HermesAgentApprovalDecision(value)
	return nil
}

// NewHermesAgentApprovalDecisionFromValue returns a pointer to a valid HermesAgentApprovalDecision
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHermesAgentApprovalDecisionFromValue(v string) (*HermesAgentApprovalDecision, error) {
	ev := HermesAgentApprovalDecision(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HermesAgentApprovalDecision: valid values are %v", v, allowedHermesAgentApprovalDecisionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HermesAgentApprovalDecision) IsValid() bool {
	for _, existing := range allowedHermesAgentApprovalDecisionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HermesAgentApprovalDecision value.
func (v HermesAgentApprovalDecision) Ptr() *HermesAgentApprovalDecision {
	return &v
}
