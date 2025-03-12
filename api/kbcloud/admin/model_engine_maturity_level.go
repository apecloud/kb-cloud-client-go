// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EngineMaturityLevel engine maturity level
type EngineMaturityLevel string

// List of EngineMaturityLevel.
const (
	EngineMaturityLevelStable       EngineMaturityLevel = "stable"
	EngineMaturityLevelBeta         EngineMaturityLevel = "beta"
	EngineMaturityLevelExperimental EngineMaturityLevel = "experimental"
	EngineMaturityLevelDeprecated   EngineMaturityLevel = "deprecated"
)

var allowedEngineMaturityLevelEnumValues = []EngineMaturityLevel{
	EngineMaturityLevelStable,
	EngineMaturityLevelBeta,
	EngineMaturityLevelExperimental,
	EngineMaturityLevelDeprecated,
}

// GetAllowedValues returns the list of possible values.
func (v *EngineMaturityLevel) GetAllowedValues() []EngineMaturityLevel {
	return allowedEngineMaturityLevelEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EngineMaturityLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EngineMaturityLevel(value)
	return nil
}

// NewEngineMaturityLevelFromValue returns a pointer to a valid EngineMaturityLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEngineMaturityLevelFromValue(v string) (*EngineMaturityLevel, error) {
	ev := EngineMaturityLevel(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EngineMaturityLevel: valid values are %v", v, allowedEngineMaturityLevelEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EngineMaturityLevel) IsValid() bool {
	for _, existing := range allowedEngineMaturityLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EngineMaturityLevel value.
func (v EngineMaturityLevel) Ptr() *EngineMaturityLevel {
	return &v
}
