// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// BenchmarkStep Step of benchmark
type BenchmarkStep string

// List of BenchmarkStep.
const (
	BenchmarkStepPrepare BenchmarkStep = "prepare"
	BenchmarkStepRun     BenchmarkStep = "run"
	BenchmarkStepCleanup BenchmarkStep = "cleanup"
	BenchmarkStepAll     BenchmarkStep = "all"
)

var allowedBenchmarkStepEnumValues = []BenchmarkStep{
	BenchmarkStepPrepare,
	BenchmarkStepRun,
	BenchmarkStepCleanup,
	BenchmarkStepAll,
}

// GetAllowedValues returns the list of possible values.
func (v *BenchmarkStep) GetAllowedValues() []BenchmarkStep {
	return allowedBenchmarkStepEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *BenchmarkStep) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = BenchmarkStep(value)
	return nil
}

// NewBenchmarkStepFromValue returns a pointer to a valid BenchmarkStep
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewBenchmarkStepFromValue(v string) (*BenchmarkStep, error) {
	ev := BenchmarkStep(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for BenchmarkStep: valid values are %v", v, allowedBenchmarkStepEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v BenchmarkStep) IsValid() bool {
	for _, existing := range allowedBenchmarkStepEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BenchmarkStep value.
func (v BenchmarkStep) Ptr() *BenchmarkStep {
	return &v
}
