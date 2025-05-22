// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// OrgMemberUpdate Org Member update
type OrgMemberUpdate struct {
	// The role of the User in the Org. Required
	Role string `json:"role"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgMemberUpdate instantiates a new OrgMemberUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgMemberUpdate(role string) *OrgMemberUpdate {
	this := OrgMemberUpdate{}
	this.Role = role
	return &this
}

// NewOrgMemberUpdateWithDefaults instantiates a new OrgMemberUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgMemberUpdateWithDefaults() *OrgMemberUpdate {
	this := OrgMemberUpdate{}
	return &this
}

// GetRole returns the Role field value.
func (o *OrgMemberUpdate) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *OrgMemberUpdate) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value.
func (o *OrgMemberUpdate) SetRole(v string) {
	o.Role = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgMemberUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["role"] = o.Role

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgMemberUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Role *string `json:"role"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Role == nil {
		return fmt.Errorf("required field role missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"role"})
	} else {
		return err
	}
	o.Role = *all.Role

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
