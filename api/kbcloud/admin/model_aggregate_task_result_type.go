// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// AggregateTaskResultType Specifies the type of the aggregate task result.
type AggregateTaskResultType string

// List of AggregateTaskResultType.
const (
	AggregateTaskResultTypeGreen  AggregateTaskResultType = "green"
	AggregateTaskResultTypeYellow AggregateTaskResultType = "yellow"
	AggregateTaskResultTypeRed    AggregateTaskResultType = "red"
)

var allowedAggregateTaskResultTypeEnumValues = []AggregateTaskResultType{
	AggregateTaskResultTypeGreen,
	AggregateTaskResultTypeYellow,
	AggregateTaskResultTypeRed,
}

// GetAllowedValues returns the list of possible values.
func (v *AggregateTaskResultType) GetAllowedValues() []AggregateTaskResultType {
	return allowedAggregateTaskResultTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AggregateTaskResultType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AggregateTaskResultType(value)
	return nil
}

// NewAggregateTaskResultTypeFromValue returns a pointer to a valid AggregateTaskResultType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAggregateTaskResultTypeFromValue(v string) (*AggregateTaskResultType, error) {
	ev := AggregateTaskResultType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AggregateTaskResultType: valid values are %v", v, allowedAggregateTaskResultTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AggregateTaskResultType) IsValid() bool {
	for _, existing := range allowedAggregateTaskResultTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AggregateTaskResultType value.
func (v AggregateTaskResultType) Ptr() *AggregateTaskResultType {
	return &v
}
