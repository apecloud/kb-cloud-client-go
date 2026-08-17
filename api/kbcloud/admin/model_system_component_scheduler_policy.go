// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type SystemComponentSchedulerPolicy string

// List of SystemComponentSchedulerPolicy.
const (
	SystemComponentSchedulerPolicyDefault     SystemComponentSchedulerPolicy = "default"
	SystemComponentSchedulerPolicyKoordinator SystemComponentSchedulerPolicy = "koordinator"
)

var allowedSystemComponentSchedulerPolicyEnumValues = []SystemComponentSchedulerPolicy{
	SystemComponentSchedulerPolicyDefault,
	SystemComponentSchedulerPolicyKoordinator,
}

// GetAllowedValues returns the list of possible values.
func (v *SystemComponentSchedulerPolicy) GetAllowedValues() []SystemComponentSchedulerPolicy {
	return allowedSystemComponentSchedulerPolicyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SystemComponentSchedulerPolicy) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SystemComponentSchedulerPolicy(value)
	return nil
}

// NewSystemComponentSchedulerPolicyFromValue returns a pointer to a valid SystemComponentSchedulerPolicy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSystemComponentSchedulerPolicyFromValue(v string) (*SystemComponentSchedulerPolicy, error) {
	ev := SystemComponentSchedulerPolicy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SystemComponentSchedulerPolicy: valid values are %v", v, allowedSystemComponentSchedulerPolicyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SystemComponentSchedulerPolicy) IsValid() bool {
	for _, existing := range allowedSystemComponentSchedulerPolicyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SystemComponentSchedulerPolicy value.
func (v SystemComponentSchedulerPolicy) Ptr() *SystemComponentSchedulerPolicy {
	return &v
}
