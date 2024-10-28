// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// BackupRetentionPolicy backup retention policy when cluster is deleted
type BackupRetentionPolicy string

// List of BackupRetentionPolicy.
const (
	BACKUPRETENTIONPOLICY_ALL     BackupRetentionPolicy = "All"
	BACKUPRETENTIONPOLICY_LASTONE BackupRetentionPolicy = "LastOne"
	BACKUPRETENTIONPOLICY_WIPEOUT BackupRetentionPolicy = "WipeOut"
)

var allowedBackupRetentionPolicyEnumValues = []BackupRetentionPolicy{
	BACKUPRETENTIONPOLICY_ALL,
	BACKUPRETENTIONPOLICY_LASTONE,
	BACKUPRETENTIONPOLICY_WIPEOUT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *BackupRetentionPolicy) GetAllowedValues() []BackupRetentionPolicy {
	return allowedBackupRetentionPolicyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *BackupRetentionPolicy) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = BackupRetentionPolicy(value)
	return nil
}

// NewBackupRetentionPolicyFromValue returns a pointer to a valid BackupRetentionPolicy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewBackupRetentionPolicyFromValue(v string) (*BackupRetentionPolicy, error) {
	ev := BackupRetentionPolicy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for BackupRetentionPolicy: valid values are %v", v, allowedBackupRetentionPolicyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v BackupRetentionPolicy) IsValid() bool {
	for _, existing := range allowedBackupRetentionPolicyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BackupRetentionPolicy value.
func (v BackupRetentionPolicy) Ptr() *BackupRetentionPolicy {
	return &v
}
