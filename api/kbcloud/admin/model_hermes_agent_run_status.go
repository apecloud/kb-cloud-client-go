// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type HermesAgentRunStatus string

// List of HermesAgentRunStatus.
const (
	HermesAgentRunStatusStarted            HermesAgentRunStatus = "started"
	HermesAgentRunStatusRunning            HermesAgentRunStatus = "running"
	HermesAgentRunStatusWaitingForApproval HermesAgentRunStatus = "waiting_for_approval"
	HermesAgentRunStatusCompleted          HermesAgentRunStatus = "completed"
	HermesAgentRunStatusFailed             HermesAgentRunStatus = "failed"
	HermesAgentRunStatusCancelled          HermesAgentRunStatus = "cancelled"
)

var allowedHermesAgentRunStatusEnumValues = []HermesAgentRunStatus{
	HermesAgentRunStatusStarted,
	HermesAgentRunStatusRunning,
	HermesAgentRunStatusWaitingForApproval,
	HermesAgentRunStatusCompleted,
	HermesAgentRunStatusFailed,
	HermesAgentRunStatusCancelled,
}

// GetAllowedValues returns the list of possible values.
func (v *HermesAgentRunStatus) GetAllowedValues() []HermesAgentRunStatus {
	return allowedHermesAgentRunStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HermesAgentRunStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HermesAgentRunStatus(value)
	return nil
}

// NewHermesAgentRunStatusFromValue returns a pointer to a valid HermesAgentRunStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHermesAgentRunStatusFromValue(v string) (*HermesAgentRunStatus, error) {
	ev := HermesAgentRunStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HermesAgentRunStatus: valid values are %v", v, allowedHermesAgentRunStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HermesAgentRunStatus) IsValid() bool {
	for _, existing := range allowedHermesAgentRunStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HermesAgentRunStatus value.
func (v HermesAgentRunStatus) Ptr() *HermesAgentRunStatus {
	return &v
}
