// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// Target_role Role name to match. Use "*" to match all roles (or no role if the component has none).
type Target_role string

// List of Target_role.
const (
	Target_rolePrimary Target_role = "Primary"
	Target_role        Target_role = "*"
)

var allowedTarget_roleEnumValues = []Target_role{
	Target_rolePrimary,
	Target_role,
}

// GetAllowedValues returns the list of possible values.
func (v *Target_role) GetAllowedValues() []Target_role {
	return allowedTarget_roleEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *Target_role) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = Target_role(value)
	return nil
}

// NewTarget_roleFromValue returns a pointer to a valid Target_role
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTarget_roleFromValue(v string) (*Target_role, error) {
	ev := Target_role(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for Target_role: valid values are %v", v, allowedTarget_roleEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v Target_role) IsValid() bool {
	for _, existing := range allowedTarget_roleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Target_role value.
func (v Target_role) Ptr() *Target_role {
	return &v
}
