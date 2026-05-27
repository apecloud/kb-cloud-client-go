// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DataCheckStatus string

// List of DataCheckStatus.
const (
	DataCheckStatusRunning  DataCheckStatus = "Running"
	DataCheckStatusFinished DataCheckStatus = "Finished"
	DataCheckStatusFailed   DataCheckStatus = "Failed"
)

var allowedDataCheckStatusEnumValues = []DataCheckStatus{
	DataCheckStatusRunning,
	DataCheckStatusFinished,
	DataCheckStatusFailed,
}

// GetAllowedValues returns the list of possible values.
func (v *DataCheckStatus) GetAllowedValues() []DataCheckStatus {
	return allowedDataCheckStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DataCheckStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DataCheckStatus(value)
	return nil
}

// NewDataCheckStatusFromValue returns a pointer to a valid DataCheckStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDataCheckStatusFromValue(v string) (*DataCheckStatus, error) {
	ev := DataCheckStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DataCheckStatus: valid values are %v", v, allowedDataCheckStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DataCheckStatus) IsValid() bool {
	for _, existing := range allowedDataCheckStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataCheckStatus value.
func (v DataCheckStatus) Ptr() *DataCheckStatus {
	return &v
}
