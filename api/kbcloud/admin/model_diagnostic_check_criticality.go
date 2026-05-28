// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticCheckCriticality Whether this check is on the critical diagnostic path for health score calculation.
type DiagnosticCheckCriticality string

// List of DiagnosticCheckCriticality.
const (
	DiagnosticCheckCriticalityCriticalPath DiagnosticCheckCriticality = "critical_path"
	DiagnosticCheckCriticalitySupporting   DiagnosticCheckCriticality = "supporting"
)

var allowedDiagnosticCheckCriticalityEnumValues = []DiagnosticCheckCriticality{
	DiagnosticCheckCriticalityCriticalPath,
	DiagnosticCheckCriticalitySupporting,
}

// GetAllowedValues returns the list of possible values.
func (v *DiagnosticCheckCriticality) GetAllowedValues() []DiagnosticCheckCriticality {
	return allowedDiagnosticCheckCriticalityEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DiagnosticCheckCriticality) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DiagnosticCheckCriticality(value)
	return nil
}

// NewDiagnosticCheckCriticalityFromValue returns a pointer to a valid DiagnosticCheckCriticality
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDiagnosticCheckCriticalityFromValue(v string) (*DiagnosticCheckCriticality, error) {
	ev := DiagnosticCheckCriticality(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DiagnosticCheckCriticality: valid values are %v", v, allowedDiagnosticCheckCriticalityEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DiagnosticCheckCriticality) IsValid() bool {
	for _, existing := range allowedDiagnosticCheckCriticalityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DiagnosticCheckCriticality value.
func (v DiagnosticCheckCriticality) Ptr() *DiagnosticCheckCriticality {
	return &v
}
