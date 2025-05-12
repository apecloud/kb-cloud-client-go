// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// OrgMemberAdd MemberAdd is the payload for adding organization member
type OrgMemberAdd struct {
	// The name of the role in the organization
	Role string `json:"role"`
	// The name of the user
	UserName string `json:"userName"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgMemberAdd instantiates a new OrgMemberAdd object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgMemberAdd(role string, userName string) *OrgMemberAdd {
	this := OrgMemberAdd{}
	this.Role = role
	this.UserName = userName
	return &this
}

// NewOrgMemberAddWithDefaults instantiates a new OrgMemberAdd object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgMemberAddWithDefaults() *OrgMemberAdd {
	this := OrgMemberAdd{}
	return &this
}

// GetRole returns the Role field value.
func (o *OrgMemberAdd) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *OrgMemberAdd) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value.
func (o *OrgMemberAdd) SetRole(v string) {
	o.Role = v
}

// GetUserName returns the UserName field value.
func (o *OrgMemberAdd) GetUserName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
func (o *OrgMemberAdd) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserName, true
}

// SetUserName sets field value.
func (o *OrgMemberAdd) SetUserName(v string) {
	o.UserName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgMemberAdd) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["role"] = o.Role
	toSerialize["userName"] = o.UserName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgMemberAdd) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Role     *string `json:"role"`
		UserName *string `json:"userName"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Role == nil {
		return fmt.Errorf("required field role missing")
	}
	if all.UserName == nil {
		return fmt.Errorf("required field userName missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"role", "userName"})
	} else {
		return err
	}
	o.Role = *all.Role
	o.UserName = *all.UserName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
