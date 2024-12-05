// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EnvironmentStorageType the storage type
type EnvironmentStorageType string

// List of EnvironmentStorageType.
const (
	EnvironmentStorageTypeRemote EnvironmentStorageType = "remote"
	EnvironmentStorageTypeLocal  EnvironmentStorageType = "local"
)

var allowedEnvironmentStorageTypeEnumValues = []EnvironmentStorageType{
	EnvironmentStorageTypeRemote,
	EnvironmentStorageTypeLocal,
}

// GetAllowedValues returns the list of possible values.
func (v *EnvironmentStorageType) GetAllowedValues() []EnvironmentStorageType {
	return allowedEnvironmentStorageTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EnvironmentStorageType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EnvironmentStorageType(value)
	return nil
}

// NewEnvironmentStorageTypeFromValue returns a pointer to a valid EnvironmentStorageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEnvironmentStorageTypeFromValue(v string) (*EnvironmentStorageType, error) {
	ev := EnvironmentStorageType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EnvironmentStorageType: valid values are %v", v, allowedEnvironmentStorageTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EnvironmentStorageType) IsValid() bool {
	for _, existing := range allowedEnvironmentStorageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnvironmentStorageType value.
func (v EnvironmentStorageType) Ptr() *EnvironmentStorageType {
	return &v
}
