// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type HealthDiagnosisModuleState string

// List of HealthDiagnosisModuleState.
const (
	HealthDiagnosisModuleStateWithData           HealthDiagnosisModuleState = "with_data"
	HealthDiagnosisModuleStateNoData             HealthDiagnosisModuleState = "no_data"
	HealthDiagnosisModuleStateUnsupportedEngine  HealthDiagnosisModuleState = "unsupported_engine"
	HealthDiagnosisModuleStateUnsupportedVersion HealthDiagnosisModuleState = "unsupported_version"
	HealthDiagnosisModuleStateDependencyDisabled HealthDiagnosisModuleState = "dependency_disabled"
	HealthDiagnosisModuleStatePermissionDenied   HealthDiagnosisModuleState = "permission_denied"
	HealthDiagnosisModuleStateCollectionFailed   HealthDiagnosisModuleState = "collection_failed"
)

var allowedHealthDiagnosisModuleStateEnumValues = []HealthDiagnosisModuleState{
	HealthDiagnosisModuleStateWithData,
	HealthDiagnosisModuleStateNoData,
	HealthDiagnosisModuleStateUnsupportedEngine,
	HealthDiagnosisModuleStateUnsupportedVersion,
	HealthDiagnosisModuleStateDependencyDisabled,
	HealthDiagnosisModuleStatePermissionDenied,
	HealthDiagnosisModuleStateCollectionFailed,
}

// GetAllowedValues returns the list of possible values.
func (v *HealthDiagnosisModuleState) GetAllowedValues() []HealthDiagnosisModuleState {
	return allowedHealthDiagnosisModuleStateEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HealthDiagnosisModuleState) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HealthDiagnosisModuleState(value)
	return nil
}

// NewHealthDiagnosisModuleStateFromValue returns a pointer to a valid HealthDiagnosisModuleState
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHealthDiagnosisModuleStateFromValue(v string) (*HealthDiagnosisModuleState, error) {
	ev := HealthDiagnosisModuleState(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HealthDiagnosisModuleState: valid values are %v", v, allowedHealthDiagnosisModuleStateEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HealthDiagnosisModuleState) IsValid() bool {
	for _, existing := range allowedHealthDiagnosisModuleStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HealthDiagnosisModuleState value.
func (v HealthDiagnosisModuleState) Ptr() *HealthDiagnosisModuleState {
	return &v
}
