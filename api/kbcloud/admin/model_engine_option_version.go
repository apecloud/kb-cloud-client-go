// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type EngineOptionVersion string

// List of EngineOptionVersion.
const (
	ENGINEOPTIONVERSION_CURRENT  EngineOptionVersion = "current"
	ENGINEOPTIONVERSION_ORIGINAL EngineOptionVersion = "original"
)

var allowedEngineOptionVersionEnumValues = []EngineOptionVersion{
	ENGINEOPTIONVERSION_CURRENT,
	ENGINEOPTIONVERSION_ORIGINAL,
}

// GetAllowedValues returns the list of possible values.
func (v *EngineOptionVersion) GetAllowedValues() []EngineOptionVersion {
	return allowedEngineOptionVersionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EngineOptionVersion) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EngineOptionVersion(value)
	return nil
}

// NewEngineOptionVersionFromValue returns a pointer to a valid EngineOptionVersion
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEngineOptionVersionFromValue(v string) (*EngineOptionVersion, error) {
	ev := EngineOptionVersion(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EngineOptionVersion: valid values are %v", v, allowedEngineOptionVersionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EngineOptionVersion) IsValid() bool {
	for _, existing := range allowedEngineOptionVersionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EngineOptionVersion value.
func (v EngineOptionVersion) Ptr() *EngineOptionVersion {
	return &v
}
