// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// AggregateMetaDataType The type of the aggregate meta data.
type AggregateMetaDataType string

// List of AggregateMetaDataType.
const (
	AggregateMetaDataTypeOrganization AggregateMetaDataType = "organization"
	AggregateMetaDataTypeCluster      AggregateMetaDataType = "cluster"
	AggregateMetaDataTypeUser         AggregateMetaDataType = "user"
	AggregateMetaDataTypeEnvironment  AggregateMetaDataType = "environment"
)

var allowedAggregateMetaDataTypeEnumValues = []AggregateMetaDataType{
	AggregateMetaDataTypeOrganization,
	AggregateMetaDataTypeCluster,
	AggregateMetaDataTypeUser,
	AggregateMetaDataTypeEnvironment,
}

// GetAllowedValues returns the list of possible values.
func (v *AggregateMetaDataType) GetAllowedValues() []AggregateMetaDataType {
	return allowedAggregateMetaDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AggregateMetaDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AggregateMetaDataType(value)
	return nil
}

// NewAggregateMetaDataTypeFromValue returns a pointer to a valid AggregateMetaDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAggregateMetaDataTypeFromValue(v string) (*AggregateMetaDataType, error) {
	ev := AggregateMetaDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AggregateMetaDataType: valid values are %v", v, allowedAggregateMetaDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AggregateMetaDataType) IsValid() bool {
	for _, existing := range allowedAggregateMetaDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AggregateMetaDataType value.
func (v AggregateMetaDataType) Ptr() *AggregateMetaDataType {
	return &v
}
