// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type InfluxDBTokenStatus string

// List of InfluxDBTokenStatus.
const (
	InfluxDBTokenStatusActive   InfluxDBTokenStatus = "active"
	InfluxDBTokenStatusInactive InfluxDBTokenStatus = "inactive"
)

var allowedInfluxDBTokenStatusEnumValues = []InfluxDBTokenStatus{
	InfluxDBTokenStatusActive,
	InfluxDBTokenStatusInactive,
}

// GetAllowedValues returns the list of possible values.
func (v *InfluxDBTokenStatus) GetAllowedValues() []InfluxDBTokenStatus {
	return allowedInfluxDBTokenStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *InfluxDBTokenStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = InfluxDBTokenStatus(value)
	return nil
}

// NewInfluxDBTokenStatusFromValue returns a pointer to a valid InfluxDBTokenStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewInfluxDBTokenStatusFromValue(v string) (*InfluxDBTokenStatus, error) {
	ev := InfluxDBTokenStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for InfluxDBTokenStatus: valid values are %v", v, allowedInfluxDBTokenStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v InfluxDBTokenStatus) IsValid() bool {
	for _, existing := range allowedInfluxDBTokenStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InfluxDBTokenStatus value.
func (v InfluxDBTokenStatus) Ptr() *InfluxDBTokenStatus {
	return &v
}
