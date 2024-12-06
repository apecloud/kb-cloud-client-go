// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
)

// DisasterRecoveryEventType the event type of disasterRecovery history, support values: [CreateInstance, DeleteInstance, Promote]
type DisasterRecoveryEventType string

// List of DisasterRecoveryEventType.
const (
	DisasterRecoveryEventTypeCreateInstance DisasterRecoveryEventType = "CreateInstance"
	DisasterRecoveryEventTypeDeleteInstance DisasterRecoveryEventType = "DeleteInstance"
	DisasterRecoveryEventTypePromote        DisasterRecoveryEventType = "Promote"
)

var allowedDisasterRecoveryEventTypeEnumValues = []DisasterRecoveryEventType{
	DisasterRecoveryEventTypeCreateInstance,
	DisasterRecoveryEventTypeDeleteInstance,
	DisasterRecoveryEventTypePromote,
}

// GetAllowedValues returns the list of possible values.
func (v *DisasterRecoveryEventType) GetAllowedValues() []DisasterRecoveryEventType {
	return allowedDisasterRecoveryEventTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DisasterRecoveryEventType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DisasterRecoveryEventType(value)
	return nil
}

// NewDisasterRecoveryEventTypeFromValue returns a pointer to a valid DisasterRecoveryEventType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDisasterRecoveryEventTypeFromValue(v string) (*DisasterRecoveryEventType, error) {
	ev := DisasterRecoveryEventType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DisasterRecoveryEventType: valid values are %v", v, allowedDisasterRecoveryEventTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DisasterRecoveryEventType) IsValid() bool {
	for _, existing := range allowedDisasterRecoveryEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DisasterRecoveryEventType value.
func (v DisasterRecoveryEventType) Ptr() *DisasterRecoveryEventType {
	return &v
}

// NullableDisasterRecoveryEventType handles when a null is used for DisasterRecoveryEventType.
type NullableDisasterRecoveryEventType struct {
	value *DisasterRecoveryEventType
	isSet bool
}

// Get returns the associated value.
func (v NullableDisasterRecoveryEventType) Get() *DisasterRecoveryEventType {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableDisasterRecoveryEventType) Set(val *DisasterRecoveryEventType) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableDisasterRecoveryEventType) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableDisasterRecoveryEventType) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDisasterRecoveryEventType initializes the struct as if Set has been called.
func NewNullableDisasterRecoveryEventType(val *DisasterRecoveryEventType) *NullableDisasterRecoveryEventType {
	return &NullableDisasterRecoveryEventType{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableDisasterRecoveryEventType) MarshalJSON() ([]byte, error) {
	return common.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableDisasterRecoveryEventType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return common.Unmarshal(src, &v.value)
}
