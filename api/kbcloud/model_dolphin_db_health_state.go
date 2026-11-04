// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DolphinDBHealthState string

// List of DolphinDBHealthState.
const (
	DolphinDBHealthStatePassed        DolphinDBHealthState = "passed"
	DolphinDBHealthStateFailed        DolphinDBHealthState = "failed"
	DolphinDBHealthStateUnknown       DolphinDBHealthState = "unknown"
	DolphinDBHealthStateNotApplicable DolphinDBHealthState = "notApplicable"
)

var allowedDolphinDBHealthStateEnumValues = []DolphinDBHealthState{
	DolphinDBHealthStatePassed,
	DolphinDBHealthStateFailed,
	DolphinDBHealthStateUnknown,
	DolphinDBHealthStateNotApplicable,
}

// GetAllowedValues returns the list of possible values.
func (v *DolphinDBHealthState) GetAllowedValues() []DolphinDBHealthState {
	return allowedDolphinDBHealthStateEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DolphinDBHealthState) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DolphinDBHealthState(value)
	return nil
}

// NewDolphinDBHealthStateFromValue returns a pointer to a valid DolphinDBHealthState
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDolphinDBHealthStateFromValue(v string) (*DolphinDBHealthState, error) {
	ev := DolphinDBHealthState(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DolphinDBHealthState: valid values are %v", v, allowedDolphinDBHealthStateEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DolphinDBHealthState) IsValid() bool {
	for _, existing := range allowedDolphinDBHealthStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DolphinDBHealthState value.
func (v DolphinDBHealthState) Ptr() *DolphinDBHealthState {
	return &v
}
