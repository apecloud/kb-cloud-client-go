// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EnvironmentCredentialUsage The intended usage of the environment credential.
type EnvironmentCredentialUsage string

// List of EnvironmentCredentialUsage.
const (
	EnvironmentCredentialUsageDns EnvironmentCredentialUsage = "dns"
)

var allowedEnvironmentCredentialUsageEnumValues = []EnvironmentCredentialUsage{
	EnvironmentCredentialUsageDns,
}

// GetAllowedValues returns the list of possible values.
func (v *EnvironmentCredentialUsage) GetAllowedValues() []EnvironmentCredentialUsage {
	return allowedEnvironmentCredentialUsageEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EnvironmentCredentialUsage) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EnvironmentCredentialUsage(value)
	return nil
}

// NewEnvironmentCredentialUsageFromValue returns a pointer to a valid EnvironmentCredentialUsage
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEnvironmentCredentialUsageFromValue(v string) (*EnvironmentCredentialUsage, error) {
	ev := EnvironmentCredentialUsage(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EnvironmentCredentialUsage: valid values are %v", v, allowedEnvironmentCredentialUsageEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EnvironmentCredentialUsage) IsValid() bool {
	for _, existing := range allowedEnvironmentCredentialUsageEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnvironmentCredentialUsage value.
func (v EnvironmentCredentialUsage) Ptr() *EnvironmentCredentialUsage {
	return &v
}
