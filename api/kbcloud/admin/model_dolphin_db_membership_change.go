// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DolphinDBMembershipChange struct {
	// add or remove; incremental, preserves other memberships.
	Action string   `json:"action"`
	Users  []string `json:"users"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDolphinDBMembershipChange instantiates a new DolphinDBMembershipChange object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDolphinDBMembershipChange(action string, users []string) *DolphinDBMembershipChange {
	this := DolphinDBMembershipChange{}
	this.Action = action
	this.Users = users
	return &this
}

// NewDolphinDBMembershipChangeWithDefaults instantiates a new DolphinDBMembershipChange object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDolphinDBMembershipChangeWithDefaults() *DolphinDBMembershipChange {
	this := DolphinDBMembershipChange{}
	return &this
}

// GetAction returns the Action field value.
func (o *DolphinDBMembershipChange) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *DolphinDBMembershipChange) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value.
func (o *DolphinDBMembershipChange) SetAction(v string) {
	o.Action = v
}

// GetUsers returns the Users field value.
func (o *DolphinDBMembershipChange) GetUsers() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *DolphinDBMembershipChange) GetUsersOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Users, true
}

// SetUsers sets field value.
func (o *DolphinDBMembershipChange) SetUsers(v []string) {
	o.Users = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DolphinDBMembershipChange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["action"] = o.Action
	toSerialize["users"] = o.Users

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DolphinDBMembershipChange) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Action *string   `json:"action"`
		Users  *[]string `json:"users"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Action == nil {
		return fmt.Errorf("required field action missing")
	}
	if all.Users == nil {
		return fmt.Errorf("required field users missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"action", "users"})
	} else {
		return err
	}
	o.Action = *all.Action
	o.Users = *all.Users

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
