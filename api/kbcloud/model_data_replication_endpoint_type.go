// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DataReplicationEndpointType string

// List of DataReplicationEndpointType.
const (
	DataReplicationEndpointTypeCustom     DataReplicationEndpointType = "custom"
	DataReplicationEndpointTypeKubeblocks DataReplicationEndpointType = "kubeblocks"
)

var allowedDataReplicationEndpointTypeEnumValues = []DataReplicationEndpointType{
	DataReplicationEndpointTypeCustom,
	DataReplicationEndpointTypeKubeblocks,
}

// GetAllowedValues returns the list of possible values.
func (v *DataReplicationEndpointType) GetAllowedValues() []DataReplicationEndpointType {
	return allowedDataReplicationEndpointTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DataReplicationEndpointType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DataReplicationEndpointType(value)
	return nil
}

// NewDataReplicationEndpointTypeFromValue returns a pointer to a valid DataReplicationEndpointType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDataReplicationEndpointTypeFromValue(v string) (*DataReplicationEndpointType, error) {
	ev := DataReplicationEndpointType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DataReplicationEndpointType: valid values are %v", v, allowedDataReplicationEndpointTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DataReplicationEndpointType) IsValid() bool {
	for _, existing := range allowedDataReplicationEndpointTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataReplicationEndpointType value.
func (v DataReplicationEndpointType) Ptr() *DataReplicationEndpointType {
	return &v
}
