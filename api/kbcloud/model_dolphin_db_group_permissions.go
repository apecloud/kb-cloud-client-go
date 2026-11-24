// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DolphinDBGroupPermissions struct {
	Group       string                     `json:"group"`
	Permissions []DolphinDBPermissionState `json:"permissions"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDolphinDBGroupPermissions instantiates a new DolphinDBGroupPermissions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDolphinDBGroupPermissions(group string, permissions []DolphinDBPermissionState) *DolphinDBGroupPermissions {
	this := DolphinDBGroupPermissions{}
	this.Group = group
	this.Permissions = permissions
	return &this
}

// NewDolphinDBGroupPermissionsWithDefaults instantiates a new DolphinDBGroupPermissions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDolphinDBGroupPermissionsWithDefaults() *DolphinDBGroupPermissions {
	this := DolphinDBGroupPermissions{}
	return &this
}

// GetGroup returns the Group field value.
func (o *DolphinDBGroupPermissions) GetGroup() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *DolphinDBGroupPermissions) GetGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value.
func (o *DolphinDBGroupPermissions) SetGroup(v string) {
	o.Group = v
}

// GetPermissions returns the Permissions field value.
func (o *DolphinDBGroupPermissions) GetPermissions() []DolphinDBPermissionState {
	if o == nil {
		var ret []DolphinDBPermissionState
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *DolphinDBGroupPermissions) GetPermissionsOk() (*[]DolphinDBPermissionState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Permissions, true
}

// SetPermissions sets field value.
func (o *DolphinDBGroupPermissions) SetPermissions(v []DolphinDBPermissionState) {
	o.Permissions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DolphinDBGroupPermissions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["group"] = o.Group
	toSerialize["permissions"] = o.Permissions

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DolphinDBGroupPermissions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Group       *string                     `json:"group"`
		Permissions *[]DolphinDBPermissionState `json:"permissions"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Group == nil {
		return fmt.Errorf("required field group missing")
	}
	if all.Permissions == nil {
		return fmt.Errorf("required field permissions missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"group", "permissions"})
	} else {
		return err
	}
	o.Group = *all.Group
	o.Permissions = *all.Permissions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
