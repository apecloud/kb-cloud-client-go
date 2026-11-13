// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DmsAuditStatus string

// List of DmsAuditStatus.
const (
	DmsAuditStatusSuccess DmsAuditStatus = "success"
	DmsAuditStatusFailed  DmsAuditStatus = "failed"
	DmsAuditStatusUnknown DmsAuditStatus = "unknown"
)

var allowedDmsAuditStatusEnumValues = []DmsAuditStatus{
	DmsAuditStatusSuccess,
	DmsAuditStatusFailed,
	DmsAuditStatusUnknown,
}

// GetAllowedValues returns the list of possible values.
func (v *DmsAuditStatus) GetAllowedValues() []DmsAuditStatus {
	return allowedDmsAuditStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DmsAuditStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DmsAuditStatus(value)
	return nil
}

// NewDmsAuditStatusFromValue returns a pointer to a valid DmsAuditStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDmsAuditStatusFromValue(v string) (*DmsAuditStatus, error) {
	ev := DmsAuditStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DmsAuditStatus: valid values are %v", v, allowedDmsAuditStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DmsAuditStatus) IsValid() bool {
	for _, existing := range allowedDmsAuditStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DmsAuditStatus value.
func (v DmsAuditStatus) Ptr() *DmsAuditStatus {
	return &v
}
