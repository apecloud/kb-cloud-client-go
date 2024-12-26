// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
)

type OpsStatus string

// List of OpsStatus.
const (
	OpsStatusPending    OpsStatus = "Pending"
	OpsStatusCreating   OpsStatus = "Creating"
	OpsStatusRunning    OpsStatus = "Running"
	OpsStatusCancelling OpsStatus = "Cancelling"
	OpsStatusSucceed    OpsStatus = "Succeed"
	OpsStatusCancelled  OpsStatus = "Cancelled"
	OpsStatusFailed     OpsStatus = "Failed"
	OpsStatusAborted    OpsStatus = "Aborted"
)

var allowedOpsStatusEnumValues = []OpsStatus{
	OpsStatusPending,
	OpsStatusCreating,
	OpsStatusRunning,
	OpsStatusCancelling,
	OpsStatusSucceed,
	OpsStatusCancelled,
	OpsStatusFailed,
	OpsStatusAborted,
}

// GetAllowedValues returns the list of possible values.
func (v *OpsStatus) GetAllowedValues() []OpsStatus {
	return allowedOpsStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OpsStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OpsStatus(value)
	return nil
}

// NewOpsStatusFromValue returns a pointer to a valid OpsStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOpsStatusFromValue(v string) (*OpsStatus, error) {
	ev := OpsStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OpsStatus: valid values are %v", v, allowedOpsStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OpsStatus) IsValid() bool {
	for _, existing := range allowedOpsStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OpsStatus value.
func (v OpsStatus) Ptr() *OpsStatus {
	return &v
}
