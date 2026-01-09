// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DataReplicationOpsType data replication ops type
type DataReplicationOpsType string

// List of DataReplicationOpsType.
const (
	DataReplicationOpsTypePause   DataReplicationOpsType = "pause"
	DataReplicationOpsTypeResume  DataReplicationOpsType = "resume"
	DataReplicationOpsTypeRestart DataReplicationOpsType = "restart"
)

var allowedDataReplicationOpsTypeEnumValues = []DataReplicationOpsType{
	DataReplicationOpsTypePause,
	DataReplicationOpsTypeResume,
	DataReplicationOpsTypeRestart,
}

// GetAllowedValues returns the list of possible values.
func (v *DataReplicationOpsType) GetAllowedValues() []DataReplicationOpsType {
	return allowedDataReplicationOpsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DataReplicationOpsType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DataReplicationOpsType(value)
	return nil
}

// NewDataReplicationOpsTypeFromValue returns a pointer to a valid DataReplicationOpsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDataReplicationOpsTypeFromValue(v string) (*DataReplicationOpsType, error) {
	ev := DataReplicationOpsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DataReplicationOpsType: valid values are %v", v, allowedDataReplicationOpsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DataReplicationOpsType) IsValid() bool {
	for _, existing := range allowedDataReplicationOpsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataReplicationOpsType value.
func (v DataReplicationOpsType) Ptr() *DataReplicationOpsType {
	return &v
}
