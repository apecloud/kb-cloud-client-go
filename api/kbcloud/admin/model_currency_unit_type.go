// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// CurrencyUnitType The currency uint for display in bill query
type CurrencyUnitType string

// List of CurrencyUnitType.
const (
	CurrencyUnitTypeRmb CurrencyUnitType = "RMB"
)

var allowedCurrencyUnitTypeEnumValues = []CurrencyUnitType{
	CurrencyUnitTypeRmb,
}

// GetAllowedValues returns the list of possible values.
func (v *CurrencyUnitType) GetAllowedValues() []CurrencyUnitType {
	return allowedCurrencyUnitTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CurrencyUnitType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CurrencyUnitType(value)
	return nil
}

// NewCurrencyUnitTypeFromValue returns a pointer to a valid CurrencyUnitType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCurrencyUnitTypeFromValue(v string) (*CurrencyUnitType, error) {
	ev := CurrencyUnitType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CurrencyUnitType: valid values are %v", v, allowedCurrencyUnitTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CurrencyUnitType) IsValid() bool {
	for _, existing := range allowedCurrencyUnitTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CurrencyUnitType value.
func (v CurrencyUnitType) Ptr() *CurrencyUnitType {
	return &v
}

// NullableCurrencyUnitType handles when a null is used for CurrencyUnitType.
type NullableCurrencyUnitType struct {
	value *CurrencyUnitType
	isSet bool
}

// Get returns the associated value.
func (v NullableCurrencyUnitType) Get() *CurrencyUnitType {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableCurrencyUnitType) Set(val *CurrencyUnitType) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableCurrencyUnitType) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableCurrencyUnitType) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableCurrencyUnitType initializes the struct as if Set has been called.
func NewNullableCurrencyUnitType(val *CurrencyUnitType) *NullableCurrencyUnitType {
	return &NullableCurrencyUnitType{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableCurrencyUnitType) MarshalJSON() ([]byte, error) {
	return common.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableCurrencyUnitType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return common.Unmarshal(src, &v.value)
}
