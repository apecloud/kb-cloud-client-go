// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ModeCompatibleKubeblocksVersion specify the compatible kubeblocks version. If empty, it means all versions are supported.
type ModeCompatibleKubeblocksVersion string

// List of ModeCompatibleKubeblocksVersion.
const (
	ModeCompatibleKubeblocksVersion09 ModeCompatibleKubeblocksVersion = "0.9"
	ModeCompatibleKubeblocksVersion10 ModeCompatibleKubeblocksVersion = "1.0"
)

var allowedModeCompatibleKubeblocksVersionEnumValues = []ModeCompatibleKubeblocksVersion{
	ModeCompatibleKubeblocksVersion09,
	ModeCompatibleKubeblocksVersion10,
}

// GetAllowedValues returns the list of possible values.
func (v *ModeCompatibleKubeblocksVersion) GetAllowedValues() []ModeCompatibleKubeblocksVersion {
	return allowedModeCompatibleKubeblocksVersionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ModeCompatibleKubeblocksVersion) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ModeCompatibleKubeblocksVersion(value)
	return nil
}

// NewModeCompatibleKubeblocksVersionFromValue returns a pointer to a valid ModeCompatibleKubeblocksVersion
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewModeCompatibleKubeblocksVersionFromValue(v string) (*ModeCompatibleKubeblocksVersion, error) {
	ev := ModeCompatibleKubeblocksVersion(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ModeCompatibleKubeblocksVersion: valid values are %v", v, allowedModeCompatibleKubeblocksVersionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ModeCompatibleKubeblocksVersion) IsValid() bool {
	for _, existing := range allowedModeCompatibleKubeblocksVersionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ModeCompatibleKubeblocksVersion value.
func (v ModeCompatibleKubeblocksVersion) Ptr() *ModeCompatibleKubeblocksVersion {
	return &v
}
