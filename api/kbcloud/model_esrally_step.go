// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EsrallyStep Step of esrally
type EsrallyStep string

// List of EsrallyStep.
const (
	EsrallyStepPrepare EsrallyStep = "prepare"
	EsrallyStepRun     EsrallyStep = "run"
	EsrallyStepCleanup EsrallyStep = "cleanup"
	EsrallyStepAll     EsrallyStep = "all"
)

var allowedEsrallyStepEnumValues = []EsrallyStep{
	EsrallyStepPrepare,
	EsrallyStepRun,
	EsrallyStepCleanup,
	EsrallyStepAll,
}

// GetAllowedValues returns the list of possible values.
func (v *EsrallyStep) GetAllowedValues() []EsrallyStep {
	return allowedEsrallyStepEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EsrallyStep) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EsrallyStep(value)
	return nil
}

// NewEsrallyStepFromValue returns a pointer to a valid EsrallyStep
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEsrallyStepFromValue(v string) (*EsrallyStep, error) {
	ev := EsrallyStep(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EsrallyStep: valid values are %v", v, allowedEsrallyStepEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EsrallyStep) IsValid() bool {
	for _, existing := range allowedEsrallyStepEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EsrallyStep value.
func (v EsrallyStep) Ptr() *EsrallyStep {
	return &v
}
