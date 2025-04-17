// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type EngineOptionStatus string

// List of EngineOptionStatus.
const (
	EngineOptionStatusRelease     EngineOptionStatus = "Release"
	EngineOptionStatusUncommitted EngineOptionStatus = "Uncommitted"
)

var allowedEngineOptionStatusEnumValues = []EngineOptionStatus{
	EngineOptionStatusRelease,
	EngineOptionStatusUncommitted,
}

// GetAllowedValues returns the list of possible values.
func (v *EngineOptionStatus) GetAllowedValues() []EngineOptionStatus {
	return allowedEngineOptionStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EngineOptionStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EngineOptionStatus(value)
	return nil
}

// NewEngineOptionStatusFromValue returns a pointer to a valid EngineOptionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEngineOptionStatusFromValue(v string) (*EngineOptionStatus, error) {
	ev := EngineOptionStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EngineOptionStatus: valid values are %v", v, allowedEngineOptionStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EngineOptionStatus) IsValid() bool {
	for _, existing := range allowedEngineOptionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EngineOptionStatus value.
func (v EngineOptionStatus) Ptr() *EngineOptionStatus {
	return &v
}
