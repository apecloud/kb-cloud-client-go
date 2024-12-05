// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// BackupStatus The current status. Valid values are New, InProgress, Completed, Failed.
type BackupStatus string

// List of BackupStatus.
const (
	BackupStatusNew        BackupStatus = "New"
	BackupStatusInProgress BackupStatus = "InProgress"
	BackupStatusRunning    BackupStatus = "Running"
	BackupStatusCompleted  BackupStatus = "Completed"
	BackupStatusFailed     BackupStatus = "Failed"
	BackupStatusDeleting   BackupStatus = "Deleting"
)

var allowedBackupStatusEnumValues = []BackupStatus{
	BackupStatusNew,
	BackupStatusInProgress,
	BackupStatusRunning,
	BackupStatusCompleted,
	BackupStatusFailed,
	BackupStatusDeleting,
}

// GetAllowedValues returns the list of possible values.
func (v *BackupStatus) GetAllowedValues() []BackupStatus {
	return allowedBackupStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *BackupStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = BackupStatus(value)
	return nil
}

// NewBackupStatusFromValue returns a pointer to a valid BackupStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewBackupStatusFromValue(v string) (*BackupStatus, error) {
	ev := BackupStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for BackupStatus: valid values are %v", v, allowedBackupStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v BackupStatus) IsValid() bool {
	for _, existing := range allowedBackupStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BackupStatus value.
func (v BackupStatus) Ptr() *BackupStatus {
	return &v
}
