// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type BackupRepoCheckStatus string

// List of BackupRepoCheckStatus.
const (
	BackupRepoCheckStatusPassed     BackupRepoCheckStatus = "Passed"
	BackupRepoCheckStatusFailed     BackupRepoCheckStatus = "Failed"
	BackupRepoCheckStatusNotChecked BackupRepoCheckStatus = "NotChecked"
)

var allowedBackupRepoCheckStatusEnumValues = []BackupRepoCheckStatus{
	BackupRepoCheckStatusPassed,
	BackupRepoCheckStatusFailed,
	BackupRepoCheckStatusNotChecked,
}

// GetAllowedValues returns the list of possible values.
func (v *BackupRepoCheckStatus) GetAllowedValues() []BackupRepoCheckStatus {
	return allowedBackupRepoCheckStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *BackupRepoCheckStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = BackupRepoCheckStatus(value)
	return nil
}

// NewBackupRepoCheckStatusFromValue returns a pointer to a valid BackupRepoCheckStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewBackupRepoCheckStatusFromValue(v string) (*BackupRepoCheckStatus, error) {
	ev := BackupRepoCheckStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for BackupRepoCheckStatus: valid values are %v", v, allowedBackupRepoCheckStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v BackupRepoCheckStatus) IsValid() bool {
	for _, existing := range allowedBackupRepoCheckStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BackupRepoCheckStatus value.
func (v BackupRepoCheckStatus) Ptr() *BackupRepoCheckStatus {
	return &v
}
