// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DataReplication_opsType string

// List of DataReplication_opsType.
const (
	DataReplication_opsTypePause   DataReplication_opsType = "pause"
	DataReplication_opsTypeResume  DataReplication_opsType = "resume"
	DataReplication_opsTypeRestart DataReplication_opsType = "restart"
)

var allowedDataReplication_opsTypeEnumValues = []DataReplication_opsType{
	DataReplication_opsTypePause,
	DataReplication_opsTypeResume,
	DataReplication_opsTypeRestart,
}

// GetAllowedValues returns the list of possible values.
func (v *DataReplication_opsType) GetAllowedValues() []DataReplication_opsType {
	return allowedDataReplication_opsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DataReplication_opsType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DataReplication_opsType(value)
	return nil
}

// NewDataReplication_opsTypeFromValue returns a pointer to a valid DataReplication_opsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDataReplication_opsTypeFromValue(v string) (*DataReplication_opsType, error) {
	ev := DataReplication_opsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DataReplication_opsType: valid values are %v", v, allowedDataReplication_opsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DataReplication_opsType) IsValid() bool {
	for _, existing := range allowedDataReplication_opsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataReplication_opsType value.
func (v DataReplication_opsType) Ptr() *DataReplication_opsType {
	return &v
}
