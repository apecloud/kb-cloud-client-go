// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticCheckRiskDomain Whether this check describes database health or KBE control-plane risk.
type DiagnosticCheckRiskDomain string

// List of DiagnosticCheckRiskDomain.
const (
	DiagnosticCheckRiskDomainDatabase     DiagnosticCheckRiskDomain = "database"
	DiagnosticCheckRiskDomainControlPlane DiagnosticCheckRiskDomain = "control_plane"
)

var allowedDiagnosticCheckRiskDomainEnumValues = []DiagnosticCheckRiskDomain{
	DiagnosticCheckRiskDomainDatabase,
	DiagnosticCheckRiskDomainControlPlane,
}

// GetAllowedValues returns the list of possible values.
func (v *DiagnosticCheckRiskDomain) GetAllowedValues() []DiagnosticCheckRiskDomain {
	return allowedDiagnosticCheckRiskDomainEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DiagnosticCheckRiskDomain) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DiagnosticCheckRiskDomain(value)
	return nil
}

// NewDiagnosticCheckRiskDomainFromValue returns a pointer to a valid DiagnosticCheckRiskDomain
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDiagnosticCheckRiskDomainFromValue(v string) (*DiagnosticCheckRiskDomain, error) {
	ev := DiagnosticCheckRiskDomain(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DiagnosticCheckRiskDomain: valid values are %v", v, allowedDiagnosticCheckRiskDomainEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DiagnosticCheckRiskDomain) IsValid() bool {
	for _, existing := range allowedDiagnosticCheckRiskDomainEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DiagnosticCheckRiskDomain value.
func (v DiagnosticCheckRiskDomain) Ptr() *DiagnosticCheckRiskDomain {
	return &v
}
