// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// BackupType the type of backup
type BackupType string

// List of BackupType.
const (
	BACKUPTYPE_FULL         BackupType = "Full"
	BACKUPTYPE_INCREMENTAL  BackupType = "Incremental"
	BACKUPTYPE_DIFFERENTIAL BackupType = "Differential"
	BACKUPTYPE_CONTINUOUS   BackupType = "Continuous"
)

var allowedBackupTypeEnumValues = []BackupType{
	BACKUPTYPE_FULL,
	BACKUPTYPE_INCREMENTAL,
	BACKUPTYPE_DIFFERENTIAL,
	BACKUPTYPE_CONTINUOUS,
}

// GetAllowedValues returns the list of possible values.
func (v *BackupType) GetAllowedValues() []BackupType {
	return allowedBackupTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *BackupType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = BackupType(value)
	return nil
}

// NewBackupTypeFromValue returns a pointer to a valid BackupType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewBackupTypeFromValue(v string) (*BackupType, error) {
	ev := BackupType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for BackupType: valid values are %v", v, allowedBackupTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v BackupType) IsValid() bool {
	for _, existing := range allowedBackupTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BackupType value.
func (v BackupType) Ptr() *BackupType {
	return &v
}
