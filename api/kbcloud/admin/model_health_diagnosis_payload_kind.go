// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type HealthDiagnosisPayloadKind string

// List of HealthDiagnosisPayloadKind.
const (
	HealthDiagnosisPayloadKindPostgresqlRealtimeSessionsV1 HealthDiagnosisPayloadKind = "postgresql.realtimeSessions.v1"
)

var allowedHealthDiagnosisPayloadKindEnumValues = []HealthDiagnosisPayloadKind{
	HealthDiagnosisPayloadKindPostgresqlRealtimeSessionsV1,
}

// GetAllowedValues returns the list of possible values.
func (v *HealthDiagnosisPayloadKind) GetAllowedValues() []HealthDiagnosisPayloadKind {
	return allowedHealthDiagnosisPayloadKindEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HealthDiagnosisPayloadKind) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HealthDiagnosisPayloadKind(value)
	return nil
}

// NewHealthDiagnosisPayloadKindFromValue returns a pointer to a valid HealthDiagnosisPayloadKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHealthDiagnosisPayloadKindFromValue(v string) (*HealthDiagnosisPayloadKind, error) {
	ev := HealthDiagnosisPayloadKind(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HealthDiagnosisPayloadKind: valid values are %v", v, allowedHealthDiagnosisPayloadKindEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HealthDiagnosisPayloadKind) IsValid() bool {
	for _, existing := range allowedHealthDiagnosisPayloadKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HealthDiagnosisPayloadKind value.
func (v HealthDiagnosisPayloadKind) Ptr() *HealthDiagnosisPayloadKind {
	return &v
}
