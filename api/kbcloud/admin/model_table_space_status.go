// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// TableSpaceStatus the status of this tablespace
type TableSpaceStatus string

// List of TableSpaceStatus.
const (
	TableSpaceStatusOnline  TableSpaceStatus = "ONLINE"
	TableSpaceStatusOffline TableSpaceStatus = "OFFLINE"
	TableSpaceStatusCorrupt TableSpaceStatus = "CORRUPT"
)

var allowedTableSpaceStatusEnumValues = []TableSpaceStatus{
	TableSpaceStatusOnline,
	TableSpaceStatusOffline,
	TableSpaceStatusCorrupt,
}

// GetAllowedValues returns the list of possible values.
func (v *TableSpaceStatus) GetAllowedValues() []TableSpaceStatus {
	return allowedTableSpaceStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TableSpaceStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TableSpaceStatus(value)
	return nil
}

// NewTableSpaceStatusFromValue returns a pointer to a valid TableSpaceStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTableSpaceStatusFromValue(v string) (*TableSpaceStatus, error) {
	ev := TableSpaceStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TableSpaceStatus: valid values are %v", v, allowedTableSpaceStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TableSpaceStatus) IsValid() bool {
	for _, existing := range allowedTableSpaceStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TableSpaceStatus value.
func (v TableSpaceStatus) Ptr() *TableSpaceStatus {
	return &v
}
