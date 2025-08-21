// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// AggregateTaskType Specifies the type of the aggregate meta data.
type AggregateTaskType string

// List of AggregateTaskType.
const (
	AggregateTaskTypeOrganization AggregateTaskType = "organization"
	AggregateTaskTypeEnvironment  AggregateTaskType = "environment"
	AggregateTaskTypeAll          AggregateTaskType = "all"
)

var allowedAggregateTaskTypeEnumValues = []AggregateTaskType{
	AggregateTaskTypeOrganization,
	AggregateTaskTypeEnvironment,
	AggregateTaskTypeAll,
}

// GetAllowedValues returns the list of possible values.
func (v *AggregateTaskType) GetAllowedValues() []AggregateTaskType {
	return allowedAggregateTaskTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AggregateTaskType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AggregateTaskType(value)
	return nil
}

// NewAggregateTaskTypeFromValue returns a pointer to a valid AggregateTaskType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAggregateTaskTypeFromValue(v string) (*AggregateTaskType, error) {
	ev := AggregateTaskType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AggregateTaskType: valid values are %v", v, allowedAggregateTaskTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AggregateTaskType) IsValid() bool {
	for _, existing := range allowedAggregateTaskTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AggregateTaskType value.
func (v AggregateTaskType) Ptr() *AggregateTaskType {
	return &v
}
