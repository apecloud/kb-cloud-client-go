// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EngineSchedulingStrategy * `HardAntiAffinity` - Strictly enforced; pods will not be scheduled if constraints cannot be met.
// * `SoftAntiAffinity` - Best-effort; the scheduler prefers to satisfy constraints but may place pods together if necessary.
// * `Disabled` - No anti-affinity constraints applied.
type EngineSchedulingStrategy string

// List of EngineSchedulingStrategy.
const (
	EngineSchedulingStrategyHardAntiAffinity EngineSchedulingStrategy = "HardAntiAffinity"
	EngineSchedulingStrategySoftAntiAffinity EngineSchedulingStrategy = "SoftAntiAffinity"
	EngineSchedulingStrategyDisabled         EngineSchedulingStrategy = "Disabled"
)

var allowedEngineSchedulingStrategyEnumValues = []EngineSchedulingStrategy{
	EngineSchedulingStrategyHardAntiAffinity,
	EngineSchedulingStrategySoftAntiAffinity,
	EngineSchedulingStrategyDisabled,
}

// GetAllowedValues returns the list of possible values.
func (v *EngineSchedulingStrategy) GetAllowedValues() []EngineSchedulingStrategy {
	return allowedEngineSchedulingStrategyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EngineSchedulingStrategy) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EngineSchedulingStrategy(value)
	return nil
}

// NewEngineSchedulingStrategyFromValue returns a pointer to a valid EngineSchedulingStrategy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEngineSchedulingStrategyFromValue(v string) (*EngineSchedulingStrategy, error) {
	ev := EngineSchedulingStrategy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EngineSchedulingStrategy: valid values are %v", v, allowedEngineSchedulingStrategyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EngineSchedulingStrategy) IsValid() bool {
	for _, existing := range allowedEngineSchedulingStrategyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EngineSchedulingStrategy value.
func (v EngineSchedulingStrategy) Ptr() *EngineSchedulingStrategy {
	return &v
}

// NullableEngineSchedulingStrategy handles when a null is used for EngineSchedulingStrategy.
type NullableEngineSchedulingStrategy struct {
	value *EngineSchedulingStrategy
	isSet bool
}

// Get returns the associated value.
func (v NullableEngineSchedulingStrategy) Get() *EngineSchedulingStrategy {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableEngineSchedulingStrategy) Set(val *EngineSchedulingStrategy) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableEngineSchedulingStrategy) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableEngineSchedulingStrategy) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableEngineSchedulingStrategy initializes the struct as if Set has been called.
func NewNullableEngineSchedulingStrategy(val *EngineSchedulingStrategy) *NullableEngineSchedulingStrategy {
	return &NullableEngineSchedulingStrategy{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableEngineSchedulingStrategy) MarshalJSON() ([]byte, error) {
	return common.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableEngineSchedulingStrategy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return common.Unmarshal(src, &v.value)
}
