// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// UserAuthorization User authorization
type UserAuthorization struct {
	// The organization name
	OrgName string `json:"orgName"`
	// The role name in the organization
	RoleName string `json:"roleName"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUserAuthorization instantiates a new UserAuthorization object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUserAuthorization(orgName string, roleName string) *UserAuthorization {
	this := UserAuthorization{}
	this.OrgName = orgName
	this.RoleName = roleName
	return &this
}

// NewUserAuthorizationWithDefaults instantiates a new UserAuthorization object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUserAuthorizationWithDefaults() *UserAuthorization {
	this := UserAuthorization{}
	return &this
}

// GetOrgName returns the OrgName field value.
func (o *UserAuthorization) GetOrgName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value
// and a boolean to check if the value has been set.
func (o *UserAuthorization) GetOrgNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgName, true
}

// SetOrgName sets field value.
func (o *UserAuthorization) SetOrgName(v string) {
	o.OrgName = v
}

// GetRoleName returns the RoleName field value.
func (o *UserAuthorization) GetRoleName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value
// and a boolean to check if the value has been set.
func (o *UserAuthorization) GetRoleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleName, true
}

// SetRoleName sets field value.
func (o *UserAuthorization) SetRoleName(v string) {
	o.RoleName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UserAuthorization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["orgName"] = o.OrgName
	toSerialize["roleName"] = o.RoleName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UserAuthorization) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OrgName  *string `json:"orgName"`
		RoleName *string `json:"roleName"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.OrgName == nil {
		return fmt.Errorf("required field orgName missing")
	}
	if all.RoleName == nil {
		return fmt.Errorf("required field roleName missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"orgName", "roleName"})
	} else {
		return err
	}
	o.OrgName = *all.OrgName
	o.RoleName = *all.RoleName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
