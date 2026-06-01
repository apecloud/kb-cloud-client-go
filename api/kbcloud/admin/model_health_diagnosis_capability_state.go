// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type HealthDiagnosisCapabilityState string

// List of HealthDiagnosisCapabilityState.
const (
	HealthDiagnosisCapabilityStateSupported          HealthDiagnosisCapabilityState = "supported"
	HealthDiagnosisCapabilityStateUnsupportedEngine  HealthDiagnosisCapabilityState = "unsupported_engine"
	HealthDiagnosisCapabilityStateUnsupportedVersion HealthDiagnosisCapabilityState = "unsupported_version"
	HealthDiagnosisCapabilityStateDependencyDisabled HealthDiagnosisCapabilityState = "dependency_disabled"
	HealthDiagnosisCapabilityStatePermissionDenied   HealthDiagnosisCapabilityState = "permission_denied"
	HealthDiagnosisCapabilityStateCollectionFailed   HealthDiagnosisCapabilityState = "collection_failed"
)

var allowedHealthDiagnosisCapabilityStateEnumValues = []HealthDiagnosisCapabilityState{
	HealthDiagnosisCapabilityStateSupported,
	HealthDiagnosisCapabilityStateUnsupportedEngine,
	HealthDiagnosisCapabilityStateUnsupportedVersion,
	HealthDiagnosisCapabilityStateDependencyDisabled,
	HealthDiagnosisCapabilityStatePermissionDenied,
	HealthDiagnosisCapabilityStateCollectionFailed,
}

// GetAllowedValues returns the list of possible values.
func (v *HealthDiagnosisCapabilityState) GetAllowedValues() []HealthDiagnosisCapabilityState {
	return allowedHealthDiagnosisCapabilityStateEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HealthDiagnosisCapabilityState) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HealthDiagnosisCapabilityState(value)
	return nil
}

// NewHealthDiagnosisCapabilityStateFromValue returns a pointer to a valid HealthDiagnosisCapabilityState
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHealthDiagnosisCapabilityStateFromValue(v string) (*HealthDiagnosisCapabilityState, error) {
	ev := HealthDiagnosisCapabilityState(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HealthDiagnosisCapabilityState: valid values are %v", v, allowedHealthDiagnosisCapabilityStateEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HealthDiagnosisCapabilityState) IsValid() bool {
	for _, existing := range allowedHealthDiagnosisCapabilityStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HealthDiagnosisCapabilityState value.
func (v HealthDiagnosisCapabilityState) Ptr() *HealthDiagnosisCapabilityState {
	return &v
}
