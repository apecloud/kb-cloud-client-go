// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// OrgMemberRole The default role of the User in the Org. Required
type OrgMemberRole string

// List of OrgMemberRole.
const (
	OrgMemberRoleOrgadmin  OrgMemberRole = "orgadmin"
	OrgMemberRoleOrgaudit  OrgMemberRole = "orgaudit"
	OrgMemberRoleOrgmember OrgMemberRole = "orgmember"
	OrgMemberRoleOrgviewer OrgMemberRole = "orgviewer"
)

var allowedOrgMemberRoleEnumValues = []OrgMemberRole{
	OrgMemberRoleOrgadmin,
	OrgMemberRoleOrgaudit,
	OrgMemberRoleOrgmember,
	OrgMemberRoleOrgviewer,
}

// GetAllowedValues returns the list of possible values.
func (v *OrgMemberRole) GetAllowedValues() []OrgMemberRole {
	return allowedOrgMemberRoleEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OrgMemberRole) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OrgMemberRole(value)
	return nil
}

// NewOrgMemberRoleFromValue returns a pointer to a valid OrgMemberRole
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOrgMemberRoleFromValue(v string) (*OrgMemberRole, error) {
	ev := OrgMemberRole(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OrgMemberRole: valid values are %v", v, allowedOrgMemberRoleEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OrgMemberRole) IsValid() bool {
	for _, existing := range allowedOrgMemberRoleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrgMemberRole value.
func (v OrgMemberRole) Ptr() *OrgMemberRole {
	return &v
}
