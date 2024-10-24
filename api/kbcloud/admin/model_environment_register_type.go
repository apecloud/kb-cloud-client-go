// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EnvironmentRegisterType Configuration for privisoining environment resources
type EnvironmentRegisterType string

// List of EnvironmentRegisterType.
const (
	ENVIRONMENTREGISTERTYPE_KUBECONFIG    EnvironmentRegisterType = "kubeconfig"
	ENVIRONMENTREGISTERTYPE_IAC           EnvironmentRegisterType = "iac"
	ENVIRONMENTREGISTERTYPE_CONNECT_AGENT EnvironmentRegisterType = "connect-agent"
)

var allowedEnvironmentRegisterTypeEnumValues = []EnvironmentRegisterType{
	ENVIRONMENTREGISTERTYPE_KUBECONFIG,
	ENVIRONMENTREGISTERTYPE_IAC,
	ENVIRONMENTREGISTERTYPE_CONNECT_AGENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EnvironmentRegisterType) GetAllowedValues() []EnvironmentRegisterType {
	return allowedEnvironmentRegisterTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EnvironmentRegisterType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EnvironmentRegisterType(value)
	return nil
}

// NewEnvironmentRegisterTypeFromValue returns a pointer to a valid EnvironmentRegisterType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEnvironmentRegisterTypeFromValue(v string) (*EnvironmentRegisterType, error) {
	ev := EnvironmentRegisterType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EnvironmentRegisterType: valid values are %v", v, allowedEnvironmentRegisterTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EnvironmentRegisterType) IsValid() bool {
	for _, existing := range allowedEnvironmentRegisterTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnvironmentRegisterType value.
func (v EnvironmentRegisterType) Ptr() *EnvironmentRegisterType {
	return &v
}
