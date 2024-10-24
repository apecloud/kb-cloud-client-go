// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// OrgMemberAdd MemberAdd is the payload for adding organization member
// NODESCRIPTION OrgMemberAdd
//
// Deprecated: This model is deprecated.
type OrgMemberAdd struct {
	// The name of the role in the organization
	Role string `json:"role"`
	// The ID of the user
	UserId string `json:"userId"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgMemberAdd instantiates a new OrgMemberAdd object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgMemberAdd(role string, userId string) *OrgMemberAdd {
	this := OrgMemberAdd{}
	this.Role = role
	this.UserId = userId
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

// GetUserId returns the UserId field value.
func (o *OrgMemberAdd) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *OrgMemberAdd) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value.
func (o *OrgMemberAdd) SetUserId(v string) {
	o.UserId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgMemberAdd) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["role"] = o.Role
	toSerialize["userId"] = o.UserId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgMemberAdd) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Role   *string `json:"role"`
		UserId *string `json:"userId"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Role == nil {
		return fmt.Errorf("required field role missing")
	}
	if all.UserId == nil {
		return fmt.Errorf("required field userId missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"role", "userId"})
	} else {
		return err
	}
	o.Role = *all.Role
	o.UserId = *all.UserId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
