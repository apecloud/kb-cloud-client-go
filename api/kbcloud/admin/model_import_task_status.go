// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ImportTaskStatus Import task status
type ImportTaskStatus string

// List of ImportTaskStatus.
const (
	ImportTaskStatusPending      ImportTaskStatus = "Pending"
	ImportTaskStatusInitializing ImportTaskStatus = "Initializing"
	ImportTaskStatusFinished     ImportTaskStatus = "Finished"
	ImportTaskStatusFailed       ImportTaskStatus = "Failed"
	ImportTaskStatusDeleting     ImportTaskStatus = "Deleting"
)

var allowedImportTaskStatusEnumValues = []ImportTaskStatus{
	ImportTaskStatusPending,
	ImportTaskStatusInitializing,
	ImportTaskStatusFinished,
	ImportTaskStatusFailed,
	ImportTaskStatusDeleting,
}

// GetAllowedValues returns the list of possible values.
func (v *ImportTaskStatus) GetAllowedValues() []ImportTaskStatus {
	return allowedImportTaskStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ImportTaskStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ImportTaskStatus(value)
	return nil
}

// NewImportTaskStatusFromValue returns a pointer to a valid ImportTaskStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewImportTaskStatusFromValue(v string) (*ImportTaskStatus, error) {
	ev := ImportTaskStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ImportTaskStatus: valid values are %v", v, allowedImportTaskStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ImportTaskStatus) IsValid() bool {
	for _, existing := range allowedImportTaskStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ImportTaskStatus value.
func (v ImportTaskStatus) Ptr() *ImportTaskStatus {
	return &v
}
