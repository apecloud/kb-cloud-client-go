// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// TolerationOperator The operator to use for matching the taint's key and value.
type TolerationOperator string

// List of TolerationOperator.
const (
	TolerationOperatorEqual  TolerationOperator = "Equal"
	TolerationOperatorExists TolerationOperator = "Exists"
)

var allowedTolerationOperatorEnumValues = []TolerationOperator{
	TolerationOperatorEqual,
	TolerationOperatorExists,
}

// GetAllowedValues returns the list of possible values.
func (v *TolerationOperator) GetAllowedValues() []TolerationOperator {
	return allowedTolerationOperatorEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TolerationOperator) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TolerationOperator(value)
	return nil
}

// NewTolerationOperatorFromValue returns a pointer to a valid TolerationOperator
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTolerationOperatorFromValue(v string) (*TolerationOperator, error) {
	ev := TolerationOperator(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TolerationOperator: valid values are %v", v, allowedTolerationOperatorEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TolerationOperator) IsValid() bool {
	for _, existing := range allowedTolerationOperatorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TolerationOperator value.
func (v TolerationOperator) Ptr() *TolerationOperator {
	return &v
}
