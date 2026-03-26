// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EnvironmentCredentialType Type of the environment credential
type EnvironmentCredentialType string

// List of EnvironmentCredentialType.
const (
	EnvironmentCredentialTypeAliyunAccessKey EnvironmentCredentialType = "aliyunAccessKey"
)

var allowedEnvironmentCredentialTypeEnumValues = []EnvironmentCredentialType{
	EnvironmentCredentialTypeAliyunAccessKey,
}

// GetAllowedValues returns the list of possible values.
func (v *EnvironmentCredentialType) GetAllowedValues() []EnvironmentCredentialType {
	return allowedEnvironmentCredentialTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EnvironmentCredentialType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EnvironmentCredentialType(value)
	return nil
}

// NewEnvironmentCredentialTypeFromValue returns a pointer to a valid EnvironmentCredentialType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEnvironmentCredentialTypeFromValue(v string) (*EnvironmentCredentialType, error) {
	ev := EnvironmentCredentialType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EnvironmentCredentialType: valid values are %v", v, allowedEnvironmentCredentialTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EnvironmentCredentialType) IsValid() bool {
	for _, existing := range allowedEnvironmentCredentialTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnvironmentCredentialType value.
func (v EnvironmentCredentialType) Ptr() *EnvironmentCredentialType {
	return &v
}
