// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ClusterStorageUsageHistoryInstanceRole Database replica role normalized from metrics labels.
type ClusterStorageUsageHistoryInstanceRole string

// List of ClusterStorageUsageHistoryInstanceRole.
const (
	ClusterStorageUsageHistoryInstanceRolePrimary ClusterStorageUsageHistoryInstanceRole = "primary"
	ClusterStorageUsageHistoryInstanceRoleStandby ClusterStorageUsageHistoryInstanceRole = "standby"
	ClusterStorageUsageHistoryInstanceRoleUnknown ClusterStorageUsageHistoryInstanceRole = "unknown"
)

var allowedClusterStorageUsageHistoryInstanceRoleEnumValues = []ClusterStorageUsageHistoryInstanceRole{
	ClusterStorageUsageHistoryInstanceRolePrimary,
	ClusterStorageUsageHistoryInstanceRoleStandby,
	ClusterStorageUsageHistoryInstanceRoleUnknown,
}

// GetAllowedValues returns the list of possible values.
func (v *ClusterStorageUsageHistoryInstanceRole) GetAllowedValues() []ClusterStorageUsageHistoryInstanceRole {
	return allowedClusterStorageUsageHistoryInstanceRoleEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ClusterStorageUsageHistoryInstanceRole) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ClusterStorageUsageHistoryInstanceRole(value)
	return nil
}

// NewClusterStorageUsageHistoryInstanceRoleFromValue returns a pointer to a valid ClusterStorageUsageHistoryInstanceRole
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewClusterStorageUsageHistoryInstanceRoleFromValue(v string) (*ClusterStorageUsageHistoryInstanceRole, error) {
	ev := ClusterStorageUsageHistoryInstanceRole(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ClusterStorageUsageHistoryInstanceRole: valid values are %v", v, allowedClusterStorageUsageHistoryInstanceRoleEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ClusterStorageUsageHistoryInstanceRole) IsValid() bool {
	for _, existing := range allowedClusterStorageUsageHistoryInstanceRoleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ClusterStorageUsageHistoryInstanceRole value.
func (v ClusterStorageUsageHistoryInstanceRole) Ptr() *ClusterStorageUsageHistoryInstanceRole {
	return &v
}
