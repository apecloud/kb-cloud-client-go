// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package admin

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)



// BackupRepoAccessMethod the access method for backup repo
type BackupRepoAccessMethod string

// List of BackupRepoAccessMethod.
const (
	BackupRepoAccessMethodMount BackupRepoAccessMethod = "Mount"
	BackupRepoAccessMethodTool BackupRepoAccessMethod = "Tool"
)

var allowedBackupRepoAccessMethodEnumValues = []BackupRepoAccessMethod{
	BackupRepoAccessMethodMount,
	BackupRepoAccessMethodTool,
}

// GetAllowedValues returns the list of possible values.
func (v *BackupRepoAccessMethod) GetAllowedValues() []BackupRepoAccessMethod {
	return allowedBackupRepoAccessMethodEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *BackupRepoAccessMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = BackupRepoAccessMethod(value)
	return nil
}

// NewBackupRepoAccessMethodFromValue returns a pointer to a valid BackupRepoAccessMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewBackupRepoAccessMethodFromValue(v string) (*BackupRepoAccessMethod, error) {
	ev := BackupRepoAccessMethod(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for BackupRepoAccessMethod: valid values are %v", v, allowedBackupRepoAccessMethodEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v BackupRepoAccessMethod) IsValid() bool {
	for _, existing := range allowedBackupRepoAccessMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BackupRepoAccessMethod value.
func (v BackupRepoAccessMethod) Ptr() *BackupRepoAccessMethod {
	return &v
}
