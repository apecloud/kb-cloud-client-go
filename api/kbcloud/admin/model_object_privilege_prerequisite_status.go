// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ObjectPrivilegePrerequisiteStatus string

// List of ObjectPrivilegePrerequisiteStatus.
const (
	ObjectPrivilegePrerequisiteStatusGranted ObjectPrivilegePrerequisiteStatus = "granted"
	ObjectPrivilegePrerequisiteStatusDenied  ObjectPrivilegePrerequisiteStatus = "denied"
	ObjectPrivilegePrerequisiteStatusUnknown ObjectPrivilegePrerequisiteStatus = "unknown"
)

var allowedObjectPrivilegePrerequisiteStatusEnumValues = []ObjectPrivilegePrerequisiteStatus{
	ObjectPrivilegePrerequisiteStatusGranted,
	ObjectPrivilegePrerequisiteStatusDenied,
	ObjectPrivilegePrerequisiteStatusUnknown,
}

// GetAllowedValues returns the list of possible values.
func (v *ObjectPrivilegePrerequisiteStatus) GetAllowedValues() []ObjectPrivilegePrerequisiteStatus {
	return allowedObjectPrivilegePrerequisiteStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObjectPrivilegePrerequisiteStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObjectPrivilegePrerequisiteStatus(value)
	return nil
}

// NewObjectPrivilegePrerequisiteStatusFromValue returns a pointer to a valid ObjectPrivilegePrerequisiteStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObjectPrivilegePrerequisiteStatusFromValue(v string) (*ObjectPrivilegePrerequisiteStatus, error) {
	ev := ObjectPrivilegePrerequisiteStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObjectPrivilegePrerequisiteStatus: valid values are %v", v, allowedObjectPrivilegePrerequisiteStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObjectPrivilegePrerequisiteStatus) IsValid() bool {
	for _, existing := range allowedObjectPrivilegePrerequisiteStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObjectPrivilegePrerequisiteStatus value.
func (v ObjectPrivilegePrerequisiteStatus) Ptr() *ObjectPrivilegePrerequisiteStatus {
	return &v
}
