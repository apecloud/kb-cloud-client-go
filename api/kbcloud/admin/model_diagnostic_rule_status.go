// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticRuleStatus Result status for one diagnostic rule.
type DiagnosticRuleStatus string

// List of DiagnosticRuleStatus.
const (
	DiagnosticRuleStatusPass          DiagnosticRuleStatus = "pass"
	DiagnosticRuleStatusWarning       DiagnosticRuleStatus = "warning"
	DiagnosticRuleStatusFailed        DiagnosticRuleStatus = "failed"
	DiagnosticRuleStatusUnknown       DiagnosticRuleStatus = "unknown"
	DiagnosticRuleStatusNotApplicable DiagnosticRuleStatus = "not_applicable"
)

var allowedDiagnosticRuleStatusEnumValues = []DiagnosticRuleStatus{
	DiagnosticRuleStatusPass,
	DiagnosticRuleStatusWarning,
	DiagnosticRuleStatusFailed,
	DiagnosticRuleStatusUnknown,
	DiagnosticRuleStatusNotApplicable,
}

// GetAllowedValues returns the list of possible values.
func (v *DiagnosticRuleStatus) GetAllowedValues() []DiagnosticRuleStatus {
	return allowedDiagnosticRuleStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DiagnosticRuleStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DiagnosticRuleStatus(value)
	return nil
}

// NewDiagnosticRuleStatusFromValue returns a pointer to a valid DiagnosticRuleStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDiagnosticRuleStatusFromValue(v string) (*DiagnosticRuleStatus, error) {
	ev := DiagnosticRuleStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DiagnosticRuleStatus: valid values are %v", v, allowedDiagnosticRuleStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DiagnosticRuleStatus) IsValid() bool {
	for _, existing := range allowedDiagnosticRuleStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DiagnosticRuleStatus value.
func (v DiagnosticRuleStatus) Ptr() *DiagnosticRuleStatus {
	return &v
}
