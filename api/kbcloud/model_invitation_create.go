// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// InvitationCreate Invitation create payload
type InvitationCreate struct {
	// The ID of the user as invitee, if email is not provided, the user ID is required
	UserId string `json:"userId"`
	// The name of the organization, for output only
	OrgName string `json:"orgName"`
	// The name of the role
	RoleName string `json:"roleName"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInvitationCreate instantiates a new InvitationCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInvitationCreate(userId string, orgName string, roleName string) *InvitationCreate {
	this := InvitationCreate{}
	this.UserId = userId
	this.OrgName = orgName
	this.RoleName = roleName
	return &this
}

// NewInvitationCreateWithDefaults instantiates a new InvitationCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInvitationCreateWithDefaults() *InvitationCreate {
	this := InvitationCreate{}
	return &this
}

// GetUserId returns the UserId field value.
func (o *InvitationCreate) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *InvitationCreate) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value.
func (o *InvitationCreate) SetUserId(v string) {
	o.UserId = v
}

// GetOrgName returns the OrgName field value.
func (o *InvitationCreate) GetOrgName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value
// and a boolean to check if the value has been set.
func (o *InvitationCreate) GetOrgNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgName, true
}

// SetOrgName sets field value.
func (o *InvitationCreate) SetOrgName(v string) {
	o.OrgName = v
}

// GetRoleName returns the RoleName field value.
func (o *InvitationCreate) GetRoleName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value
// and a boolean to check if the value has been set.
func (o *InvitationCreate) GetRoleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleName, true
}

// SetRoleName sets field value.
func (o *InvitationCreate) SetRoleName(v string) {
	o.RoleName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o InvitationCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["userId"] = o.UserId
	toSerialize["orgName"] = o.OrgName
	toSerialize["roleName"] = o.RoleName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InvitationCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		UserId   *string `json:"userId"`
		OrgName  *string `json:"orgName"`
		RoleName *string `json:"roleName"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.UserId == nil {
		return fmt.Errorf("required field userId missing")
	}
	if all.OrgName == nil {
		return fmt.Errorf("required field orgName missing")
	}
	if all.RoleName == nil {
		return fmt.Errorf("required field roleName missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"userId", "orgName", "roleName"})
	} else {
		return err
	}
	o.UserId = *all.UserId
	o.OrgName = *all.OrgName
	o.RoleName = *all.RoleName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
