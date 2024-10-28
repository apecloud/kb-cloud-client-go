// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// BenchmarkStatus the status of benchmark
type BenchmarkStatus string

// List of BenchmarkStatus.
const (
	BENCHMARKSTATUS_PENDING   BenchmarkStatus = "Pending"
	BENCHMARKSTATUS_RUNNING   BenchmarkStatus = "Running"
	BENCHMARKSTATUS_COMPLETED BenchmarkStatus = "Completed"
	BENCHMARKSTATUS_FAILED    BenchmarkStatus = "Failed"
)

var allowedBenchmarkStatusEnumValues = []BenchmarkStatus{
	BENCHMARKSTATUS_PENDING,
	BENCHMARKSTATUS_RUNNING,
	BENCHMARKSTATUS_COMPLETED,
	BENCHMARKSTATUS_FAILED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *BenchmarkStatus) GetAllowedValues() []BenchmarkStatus {
	return allowedBenchmarkStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *BenchmarkStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = BenchmarkStatus(value)
	return nil
}

// NewBenchmarkStatusFromValue returns a pointer to a valid BenchmarkStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewBenchmarkStatusFromValue(v string) (*BenchmarkStatus, error) {
	ev := BenchmarkStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for BenchmarkStatus: valid values are %v", v, allowedBenchmarkStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v BenchmarkStatus) IsValid() bool {
	for _, existing := range allowedBenchmarkStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BenchmarkStatus value.
func (v BenchmarkStatus) Ptr() *BenchmarkStatus {
	return &v
}
