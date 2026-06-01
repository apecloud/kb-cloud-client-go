// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type HealthDiagnosisCapabilityModule string

// List of HealthDiagnosisCapabilityModule.
const (
	HealthDiagnosisCapabilityModuleRealtimeSessions HealthDiagnosisCapabilityModule = "realtimeSessions"
)

var allowedHealthDiagnosisCapabilityModuleEnumValues = []HealthDiagnosisCapabilityModule{
	HealthDiagnosisCapabilityModuleRealtimeSessions,
}

// GetAllowedValues returns the list of possible values.
func (v *HealthDiagnosisCapabilityModule) GetAllowedValues() []HealthDiagnosisCapabilityModule {
	return allowedHealthDiagnosisCapabilityModuleEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HealthDiagnosisCapabilityModule) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HealthDiagnosisCapabilityModule(value)
	return nil
}

// NewHealthDiagnosisCapabilityModuleFromValue returns a pointer to a valid HealthDiagnosisCapabilityModule
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHealthDiagnosisCapabilityModuleFromValue(v string) (*HealthDiagnosisCapabilityModule, error) {
	ev := HealthDiagnosisCapabilityModule(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HealthDiagnosisCapabilityModule: valid values are %v", v, allowedHealthDiagnosisCapabilityModuleEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HealthDiagnosisCapabilityModule) IsValid() bool {
	for _, existing := range allowedHealthDiagnosisCapabilityModuleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HealthDiagnosisCapabilityModule value.
func (v HealthDiagnosisCapabilityModule) Ptr() *HealthDiagnosisCapabilityModule {
	return &v
}
