// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// SysbenchStep Step of sysbench
type SysbenchStep string

// List of SysbenchStep.
const (
	SysbenchStepPrepare SysbenchStep = "prepare"
	SysbenchStepRun     SysbenchStep = "run"
	SysbenchStepCleanup SysbenchStep = "cleanup"
	SysbenchStepAll     SysbenchStep = "all"
)

var allowedSysbenchStepEnumValues = []SysbenchStep{
	SysbenchStepPrepare,
	SysbenchStepRun,
	SysbenchStepCleanup,
	SysbenchStepAll,
}

// GetAllowedValues returns the list of possible values.
func (v *SysbenchStep) GetAllowedValues() []SysbenchStep {
	return allowedSysbenchStepEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SysbenchStep) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SysbenchStep(value)
	return nil
}

// NewSysbenchStepFromValue returns a pointer to a valid SysbenchStep
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSysbenchStepFromValue(v string) (*SysbenchStep, error) {
	ev := SysbenchStep(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SysbenchStep: valid values are %v", v, allowedSysbenchStepEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SysbenchStep) IsValid() bool {
	for _, existing := range allowedSysbenchStepEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SysbenchStep value.
func (v SysbenchStep) Ptr() *SysbenchStep {
	return &v
}
