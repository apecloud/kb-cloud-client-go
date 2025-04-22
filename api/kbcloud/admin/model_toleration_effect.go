// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// TolerationEffect The effect of the taint that this toleration tolerates.
type TolerationEffect string

// List of TolerationEffect.
const (
	TolerationEffectNoSchedule       TolerationEffect = "NoSchedule"
	TolerationEffectPreferNoSchedule TolerationEffect = "PreferNoSchedule"
	TolerationEffectNoExecute        TolerationEffect = "NoExecute"
)

var allowedTolerationEffectEnumValues = []TolerationEffect{
	TolerationEffectNoSchedule,
	TolerationEffectPreferNoSchedule,
	TolerationEffectNoExecute,
}

// GetAllowedValues returns the list of possible values.
func (v *TolerationEffect) GetAllowedValues() []TolerationEffect {
	return allowedTolerationEffectEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TolerationEffect) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TolerationEffect(value)
	return nil
}

// NewTolerationEffectFromValue returns a pointer to a valid TolerationEffect
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTolerationEffectFromValue(v string) (*TolerationEffect, error) {
	ev := TolerationEffect(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TolerationEffect: valid values are %v", v, allowedTolerationEffectEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TolerationEffect) IsValid() bool {
	for _, existing := range allowedTolerationEffectEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TolerationEffect value.
func (v TolerationEffect) Ptr() *TolerationEffect {
	return &v
}
