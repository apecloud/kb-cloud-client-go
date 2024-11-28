// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
)

// DisasterRecoveryStatus the status of promote event, support values: [Succeed, Failed, Running, Unknown]
type DisasterRecoveryStatus string

// List of DisasterRecoveryStatus.
const (
	DisasterRecoveryStatusSucceed DisasterRecoveryStatus = "Succeed"
	DisasterRecoveryStatusFailed  DisasterRecoveryStatus = "Failed"
	DisasterRecoveryStatusRunning DisasterRecoveryStatus = "Running"
	DisasterRecoveryStatusUnknown DisasterRecoveryStatus = "Unknown"
)

var allowedDisasterRecoveryStatusEnumValues = []DisasterRecoveryStatus{
	DisasterRecoveryStatusSucceed,
	DisasterRecoveryStatusFailed,
	DisasterRecoveryStatusRunning,
	DisasterRecoveryStatusUnknown,
}

// GetAllowedValues returns the list of possible values.
func (v *DisasterRecoveryStatus) GetAllowedValues() []DisasterRecoveryStatus {
	return allowedDisasterRecoveryStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DisasterRecoveryStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DisasterRecoveryStatus(value)
	return nil
}

// NewDisasterRecoveryStatusFromValue returns a pointer to a valid DisasterRecoveryStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDisasterRecoveryStatusFromValue(v string) (*DisasterRecoveryStatus, error) {
	ev := DisasterRecoveryStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DisasterRecoveryStatus: valid values are %v", v, allowedDisasterRecoveryStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DisasterRecoveryStatus) IsValid() bool {
	for _, existing := range allowedDisasterRecoveryStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DisasterRecoveryStatus value.
func (v DisasterRecoveryStatus) Ptr() *DisasterRecoveryStatus {
	return &v
}

// NullableDisasterRecoveryStatus handles when a null is used for DisasterRecoveryStatus.
type NullableDisasterRecoveryStatus struct {
	value *DisasterRecoveryStatus
	isSet bool
}

// Get returns the associated value.
func (v NullableDisasterRecoveryStatus) Get() *DisasterRecoveryStatus {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableDisasterRecoveryStatus) Set(val *DisasterRecoveryStatus) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableDisasterRecoveryStatus) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableDisasterRecoveryStatus) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDisasterRecoveryStatus initializes the struct as if Set has been called.
func NewNullableDisasterRecoveryStatus(val *DisasterRecoveryStatus) *NullableDisasterRecoveryStatus {
	return &NullableDisasterRecoveryStatus{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableDisasterRecoveryStatus) MarshalJSON() ([]byte, error) {
	return common.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableDisasterRecoveryStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return common.Unmarshal(src, &v.value)
}
