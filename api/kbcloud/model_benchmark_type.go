// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// BenchmarkType Type of benchmark
type BenchmarkType string

// List of BenchmarkType.
const (
	BENCHMARKTYPE_PGBENCH  BenchmarkType = "pgbench"
	BENCHMARKTYPE_SYSBENCH BenchmarkType = "sysbench"
	BENCHMARKTYPE_TPCC     BenchmarkType = "tpcc"
	BENCHMARKTYPE_TPCH     BenchmarkType = "tpch"
	BENCHMARKTYPE_YCSB     BenchmarkType = "ycsb"
)

var allowedBenchmarkTypeEnumValues = []BenchmarkType{
	BENCHMARKTYPE_PGBENCH,
	BENCHMARKTYPE_SYSBENCH,
	BENCHMARKTYPE_TPCC,
	BENCHMARKTYPE_TPCH,
	BENCHMARKTYPE_YCSB,
}

// GetAllowedValues returns the list of possible values.
func (v *BenchmarkType) GetAllowedValues() []BenchmarkType {
	return allowedBenchmarkTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *BenchmarkType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = BenchmarkType(value)
	return nil
}

// NewBenchmarkTypeFromValue returns a pointer to a valid BenchmarkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewBenchmarkTypeFromValue(v string) (*BenchmarkType, error) {
	ev := BenchmarkType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for BenchmarkType: valid values are %v", v, allowedBenchmarkTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v BenchmarkType) IsValid() bool {
	for _, existing := range allowedBenchmarkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BenchmarkType value.
func (v BenchmarkType) Ptr() *BenchmarkType {
	return &v
}
