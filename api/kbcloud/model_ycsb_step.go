// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// YcsbStep Step of benchmark
type YcsbStep string

// List of YcsbStep.
const (
	YCSBSTEP_PREPARE YcsbStep = "prepare"
	YCSBSTEP_RUN     YcsbStep = "run"
	YCSBSTEP_CLEANUP YcsbStep = "cleanup"
	YCSBSTEP_ALL     YcsbStep = "all"
)

var allowedYcsbStepEnumValues = []YcsbStep{
	YCSBSTEP_PREPARE,
	YCSBSTEP_RUN,
	YCSBSTEP_CLEANUP,
	YCSBSTEP_ALL,
}

// GetAllowedValues returns the list of possible values.
func (v *YcsbStep) GetAllowedValues() []YcsbStep {
	return allowedYcsbStepEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *YcsbStep) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = YcsbStep(value)
	return nil
}

// NewYcsbStepFromValue returns a pointer to a valid YcsbStep
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewYcsbStepFromValue(v string) (*YcsbStep, error) {
	ev := YcsbStep(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for YcsbStep: valid values are %v", v, allowedYcsbStepEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v YcsbStep) IsValid() bool {
	for _, existing := range allowedYcsbStepEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to YcsbStep value.
func (v YcsbStep) Ptr() *YcsbStep {
	return &v
}
