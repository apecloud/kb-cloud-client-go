// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EnvironmentArchitecture Architecture of the environment data plane nodes (arm64, amd64, or multiarch for multiple architectures)
type EnvironmentArchitecture string

// List of EnvironmentArchitecture.
const (
	EnvironmentArchitectureArm64     EnvironmentArchitecture = "arm64"
	EnvironmentArchitectureAmd64     EnvironmentArchitecture = "amd64"
	EnvironmentArchitectureMultiarch EnvironmentArchitecture = "multiarch"
)

var allowedEnvironmentArchitectureEnumValues = []EnvironmentArchitecture{
	EnvironmentArchitectureArm64,
	EnvironmentArchitectureAmd64,
	EnvironmentArchitectureMultiarch,
}

// GetAllowedValues returns the list of possible values.
func (v *EnvironmentArchitecture) GetAllowedValues() []EnvironmentArchitecture {
	return allowedEnvironmentArchitectureEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EnvironmentArchitecture) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EnvironmentArchitecture(value)
	return nil
}

// NewEnvironmentArchitectureFromValue returns a pointer to a valid EnvironmentArchitecture
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEnvironmentArchitectureFromValue(v string) (*EnvironmentArchitecture, error) {
	ev := EnvironmentArchitecture(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EnvironmentArchitecture: valid values are %v", v, allowedEnvironmentArchitectureEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EnvironmentArchitecture) IsValid() bool {
	for _, existing := range allowedEnvironmentArchitectureEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnvironmentArchitecture value.
func (v EnvironmentArchitecture) Ptr() *EnvironmentArchitecture {
	return &v
}
