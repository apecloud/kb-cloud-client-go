// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type HermesAgentModelConfigStatus string

// List of HermesAgentModelConfigStatus.
const (
	HermesAgentModelConfigStatusReady   HermesAgentModelConfigStatus = "ready"
	HermesAgentModelConfigStatusMissing HermesAgentModelConfigStatus = "missing"
)

var allowedHermesAgentModelConfigStatusEnumValues = []HermesAgentModelConfigStatus{
	HermesAgentModelConfigStatusReady,
	HermesAgentModelConfigStatusMissing,
}

// GetAllowedValues returns the list of possible values.
func (v *HermesAgentModelConfigStatus) GetAllowedValues() []HermesAgentModelConfigStatus {
	return allowedHermesAgentModelConfigStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HermesAgentModelConfigStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HermesAgentModelConfigStatus(value)
	return nil
}

// NewHermesAgentModelConfigStatusFromValue returns a pointer to a valid HermesAgentModelConfigStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHermesAgentModelConfigStatusFromValue(v string) (*HermesAgentModelConfigStatus, error) {
	ev := HermesAgentModelConfigStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HermesAgentModelConfigStatus: valid values are %v", v, allowedHermesAgentModelConfigStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HermesAgentModelConfigStatus) IsValid() bool {
	for _, existing := range allowedHermesAgentModelConfigStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HermesAgentModelConfigStatus value.
func (v HermesAgentModelConfigStatus) Ptr() *HermesAgentModelConfigStatus {
	return &v
}
