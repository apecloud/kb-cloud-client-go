// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type OrgStatus string

// List of OrgStatus.
const (
	OrgStatusActive         OrgStatus = "Active"
	OrgStatusDeletionLocked OrgStatus = "DeletionLocked"
	OrgStatusPurging        OrgStatus = "Purging"
	OrgStatusPurgeFailed    OrgStatus = "PurgeFailed"
	OrgStatusDeleted        OrgStatus = "Deleted"
)

var allowedOrgStatusEnumValues = []OrgStatus{
	OrgStatusActive,
	OrgStatusDeletionLocked,
	OrgStatusPurging,
	OrgStatusPurgeFailed,
	OrgStatusDeleted,
}

// GetAllowedValues returns the list of possible values.
func (v *OrgStatus) GetAllowedValues() []OrgStatus {
	return allowedOrgStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OrgStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OrgStatus(value)
	return nil
}

// NewOrgStatusFromValue returns a pointer to a valid OrgStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOrgStatusFromValue(v string) (*OrgStatus, error) {
	ev := OrgStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OrgStatus: valid values are %v", v, allowedOrgStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OrgStatus) IsValid() bool {
	for _, existing := range allowedOrgStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrgStatus value.
func (v OrgStatus) Ptr() *OrgStatus {
	return &v
}
