// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// BackupRepoPVReclaimPolicy Specify the reclaim policy for PVs created by this backup repo
type BackupRepoPVReclaimPolicy string

// List of BackupRepoPVReclaimPolicy.
const (
	BACKUPREPOPVRECLAIMPOLICY_RETAIN BackupRepoPVReclaimPolicy = "Retain"
	BACKUPREPOPVRECLAIMPOLICY_DELETE BackupRepoPVReclaimPolicy = "Delete"
)

var allowedBackupRepoPVReclaimPolicyEnumValues = []BackupRepoPVReclaimPolicy{
	BACKUPREPOPVRECLAIMPOLICY_RETAIN,
	BACKUPREPOPVRECLAIMPOLICY_DELETE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *BackupRepoPVReclaimPolicy) GetAllowedValues() []BackupRepoPVReclaimPolicy {
	return allowedBackupRepoPVReclaimPolicyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *BackupRepoPVReclaimPolicy) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = BackupRepoPVReclaimPolicy(value)
	return nil
}

// NewBackupRepoPVReclaimPolicyFromValue returns a pointer to a valid BackupRepoPVReclaimPolicy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewBackupRepoPVReclaimPolicyFromValue(v string) (*BackupRepoPVReclaimPolicy, error) {
	ev := BackupRepoPVReclaimPolicy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for BackupRepoPVReclaimPolicy: valid values are %v", v, allowedBackupRepoPVReclaimPolicyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v BackupRepoPVReclaimPolicy) IsValid() bool {
	for _, existing := range allowedBackupRepoPVReclaimPolicyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BackupRepoPVReclaimPolicy value.
func (v BackupRepoPVReclaimPolicy) Ptr() *BackupRepoPVReclaimPolicy {
	return &v
}
