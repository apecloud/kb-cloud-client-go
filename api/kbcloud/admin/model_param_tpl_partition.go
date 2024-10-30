// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package admin

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)



type ParamTplPartition string

// List of ParamTplPartition.
const (
	ParamTplPartitionDefault ParamTplPartition = "default"
	ParamTplPartitionCustom ParamTplPartition = "custom"
)

var allowedParamTplPartitionEnumValues = []ParamTplPartition{
	ParamTplPartitionDefault,
	ParamTplPartitionCustom,
}

// GetAllowedValues returns the list of possible values.
func (v *ParamTplPartition) GetAllowedValues() []ParamTplPartition {
	return allowedParamTplPartitionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ParamTplPartition) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ParamTplPartition(value)
	return nil
}

// NewParamTplPartitionFromValue returns a pointer to a valid ParamTplPartition
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewParamTplPartitionFromValue(v string) (*ParamTplPartition, error) {
	ev := ParamTplPartition(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ParamTplPartition: valid values are %v", v, allowedParamTplPartitionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ParamTplPartition) IsValid() bool {
	for _, existing := range allowedParamTplPartitionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ParamTplPartition value.
func (v ParamTplPartition) Ptr() *ParamTplPartition {
	return &v
}
