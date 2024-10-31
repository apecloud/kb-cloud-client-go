// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// PgbenchStep Step of pgbench
type PgbenchStep string

// List of PgbenchStep.
const (
	PgbenchStepPrepare PgbenchStep = "prepare"
	PgbenchStepRun     PgbenchStep = "run"
	PgbenchStepCleanup PgbenchStep = "cleanup"
	PgbenchStepAll     PgbenchStep = "all"
)

var allowedPgbenchStepEnumValues = []PgbenchStep{
	PgbenchStepPrepare,
	PgbenchStepRun,
	PgbenchStepCleanup,
	PgbenchStepAll,
}

// GetAllowedValues returns the list of possible values.
func (v *PgbenchStep) GetAllowedValues() []PgbenchStep {
	return allowedPgbenchStepEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PgbenchStep) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PgbenchStep(value)
	return nil
}

// NewPgbenchStepFromValue returns a pointer to a valid PgbenchStep
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPgbenchStepFromValue(v string) (*PgbenchStep, error) {
	ev := PgbenchStep(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PgbenchStep: valid values are %v", v, allowedPgbenchStepEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PgbenchStep) IsValid() bool {
	for _, existing := range allowedPgbenchStepEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PgbenchStep value.
func (v PgbenchStep) Ptr() *PgbenchStep {
	return &v
}
