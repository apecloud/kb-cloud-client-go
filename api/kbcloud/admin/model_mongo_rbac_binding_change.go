// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type MongoRBACBindingChange struct {
	// grant or revoke. Incremental; all other native roles are retained.
	Action string             `json:"action"`
	Roles  []MongoRBACRoleRef `json:"roles"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMongoRBACBindingChange instantiates a new MongoRBACBindingChange object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMongoRBACBindingChange(action string, roles []MongoRBACRoleRef) *MongoRBACBindingChange {
	this := MongoRBACBindingChange{}
	this.Action = action
	this.Roles = roles
	return &this
}

// NewMongoRBACBindingChangeWithDefaults instantiates a new MongoRBACBindingChange object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMongoRBACBindingChangeWithDefaults() *MongoRBACBindingChange {
	this := MongoRBACBindingChange{}
	return &this
}

// GetAction returns the Action field value.
func (o *MongoRBACBindingChange) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *MongoRBACBindingChange) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value.
func (o *MongoRBACBindingChange) SetAction(v string) {
	o.Action = v
}

// GetRoles returns the Roles field value.
func (o *MongoRBACBindingChange) GetRoles() []MongoRBACRoleRef {
	if o == nil {
		var ret []MongoRBACRoleRef
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *MongoRBACBindingChange) GetRolesOk() (*[]MongoRBACRoleRef, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Roles, true
}

// SetRoles sets field value.
func (o *MongoRBACBindingChange) SetRoles(v []MongoRBACRoleRef) {
	o.Roles = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MongoRBACBindingChange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["action"] = o.Action
	toSerialize["roles"] = o.Roles

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MongoRBACBindingChange) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Action *string             `json:"action"`
		Roles  *[]MongoRBACRoleRef `json:"roles"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Action == nil {
		return fmt.Errorf("required field action missing")
	}
	if all.Roles == nil {
		return fmt.Errorf("required field roles missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"action", "roles"})
	} else {
		return err
	}
	o.Action = *all.Action
	o.Roles = *all.Roles

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
