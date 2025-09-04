// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DataReplication_endpointType string

// List of DataReplication_endpointType.
const (
	DataReplication_endpointTypeCustom     DataReplication_endpointType = "custom"
	DataReplication_endpointTypeKubeblocks DataReplication_endpointType = "kubeblocks"
)

var allowedDataReplication_endpointTypeEnumValues = []DataReplication_endpointType{
	DataReplication_endpointTypeCustom,
	DataReplication_endpointTypeKubeblocks,
}

// GetAllowedValues returns the list of possible values.
func (v *DataReplication_endpointType) GetAllowedValues() []DataReplication_endpointType {
	return allowedDataReplication_endpointTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DataReplication_endpointType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DataReplication_endpointType(value)
	return nil
}

// NewDataReplication_endpointTypeFromValue returns a pointer to a valid DataReplication_endpointType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDataReplication_endpointTypeFromValue(v string) (*DataReplication_endpointType, error) {
	ev := DataReplication_endpointType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DataReplication_endpointType: valid values are %v", v, allowedDataReplication_endpointTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DataReplication_endpointType) IsValid() bool {
	for _, existing := range allowedDataReplication_endpointTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataReplication_endpointType value.
func (v DataReplication_endpointType) Ptr() *DataReplication_endpointType {
	return &v
}
