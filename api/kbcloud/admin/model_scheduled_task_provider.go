// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ScheduledTaskProvider Scheduler provider
type ScheduledTaskProvider string

// List of ScheduledTaskProvider.
const (
	ScheduledTaskProviderTaskflow ScheduledTaskProvider = "taskflow"
)

var allowedScheduledTaskProviderEnumValues = []ScheduledTaskProvider{
	ScheduledTaskProviderTaskflow,
}

// GetAllowedValues returns the list of possible values.
func (v *ScheduledTaskProvider) GetAllowedValues() []ScheduledTaskProvider {
	return allowedScheduledTaskProviderEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScheduledTaskProvider) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScheduledTaskProvider(value)
	return nil
}

// NewScheduledTaskProviderFromValue returns a pointer to a valid ScheduledTaskProvider
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScheduledTaskProviderFromValue(v string) (*ScheduledTaskProvider, error) {
	ev := ScheduledTaskProvider(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScheduledTaskProvider: valid values are %v", v, allowedScheduledTaskProviderEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScheduledTaskProvider) IsValid() bool {
	for _, existing := range allowedScheduledTaskProviderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScheduledTaskProvider value.
func (v ScheduledTaskProvider) Ptr() *ScheduledTaskProvider {
	return &v
}
