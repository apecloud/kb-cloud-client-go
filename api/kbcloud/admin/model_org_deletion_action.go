// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type OrgDeletionAction string

// List of OrgDeletionAction.
const (
	OrgDeletionActionDelete OrgDeletionAction = "delete"
	OrgDeletionActionPurge  OrgDeletionAction = "purge"
)

var allowedOrgDeletionActionEnumValues = []OrgDeletionAction{
	OrgDeletionActionDelete,
	OrgDeletionActionPurge,
}

// GetAllowedValues returns the list of possible values.
func (v *OrgDeletionAction) GetAllowedValues() []OrgDeletionAction {
	return allowedOrgDeletionActionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OrgDeletionAction) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OrgDeletionAction(value)
	return nil
}

// NewOrgDeletionActionFromValue returns a pointer to a valid OrgDeletionAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOrgDeletionActionFromValue(v string) (*OrgDeletionAction, error) {
	ev := OrgDeletionAction(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OrgDeletionAction: valid values are %v", v, allowedOrgDeletionActionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OrgDeletionAction) IsValid() bool {
	for _, existing := range allowedOrgDeletionActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrgDeletionAction value.
func (v OrgDeletionAction) Ptr() *OrgDeletionAction {
	return &v
}
