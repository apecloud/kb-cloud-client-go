// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticHealthBucket Health score bucket for the report first screen.
type DiagnosticHealthBucket string

// List of DiagnosticHealthBucket.
const (
	DiagnosticHealthBucketHealthy   DiagnosticHealthBucket = "healthy"
	DiagnosticHealthBucketAttention DiagnosticHealthBucket = "attention"
	DiagnosticHealthBucketAbnormal  DiagnosticHealthBucket = "abnormal"
	DiagnosticHealthBucketCritical  DiagnosticHealthBucket = "critical"
)

var allowedDiagnosticHealthBucketEnumValues = []DiagnosticHealthBucket{
	DiagnosticHealthBucketHealthy,
	DiagnosticHealthBucketAttention,
	DiagnosticHealthBucketAbnormal,
	DiagnosticHealthBucketCritical,
}

// GetAllowedValues returns the list of possible values.
func (v *DiagnosticHealthBucket) GetAllowedValues() []DiagnosticHealthBucket {
	return allowedDiagnosticHealthBucketEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DiagnosticHealthBucket) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DiagnosticHealthBucket(value)
	return nil
}

// NewDiagnosticHealthBucketFromValue returns a pointer to a valid DiagnosticHealthBucket
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDiagnosticHealthBucketFromValue(v string) (*DiagnosticHealthBucket, error) {
	ev := DiagnosticHealthBucket(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DiagnosticHealthBucket: valid values are %v", v, allowedDiagnosticHealthBucketEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DiagnosticHealthBucket) IsValid() bool {
	for _, existing := range allowedDiagnosticHealthBucketEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DiagnosticHealthBucket value.
func (v DiagnosticHealthBucket) Ptr() *DiagnosticHealthBucket {
	return &v
}
