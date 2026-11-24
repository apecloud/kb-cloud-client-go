// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DolphinDBUser struct {
	Name    string   `json:"name"`
	IsAdmin bool     `json:"isAdmin"`
	Groups  []string `json:"groups"`
	// Built-in admin and management account are protected from mutation.
	ReadOnly bool `json:"readOnly"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDolphinDBUser instantiates a new DolphinDBUser object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDolphinDBUser(name string, isAdmin bool, groups []string, readOnly bool) *DolphinDBUser {
	this := DolphinDBUser{}
	this.Name = name
	this.IsAdmin = isAdmin
	this.Groups = groups
	this.ReadOnly = readOnly
	return &this
}

// NewDolphinDBUserWithDefaults instantiates a new DolphinDBUser object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDolphinDBUserWithDefaults() *DolphinDBUser {
	this := DolphinDBUser{}
	return &this
}

// GetName returns the Name field value.
func (o *DolphinDBUser) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DolphinDBUser) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *DolphinDBUser) SetName(v string) {
	o.Name = v
}

// GetIsAdmin returns the IsAdmin field value.
func (o *DolphinDBUser) GetIsAdmin() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsAdmin
}

// GetIsAdminOk returns a tuple with the IsAdmin field value
// and a boolean to check if the value has been set.
func (o *DolphinDBUser) GetIsAdminOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsAdmin, true
}

// SetIsAdmin sets field value.
func (o *DolphinDBUser) SetIsAdmin(v bool) {
	o.IsAdmin = v
}

// GetGroups returns the Groups field value.
func (o *DolphinDBUser) GetGroups() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value
// and a boolean to check if the value has been set.
func (o *DolphinDBUser) GetGroupsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Groups, true
}

// SetGroups sets field value.
func (o *DolphinDBUser) SetGroups(v []string) {
	o.Groups = v
}

// GetReadOnly returns the ReadOnly field value.
func (o *DolphinDBUser) GetReadOnly() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value
// and a boolean to check if the value has been set.
func (o *DolphinDBUser) GetReadOnlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReadOnly, true
}

// SetReadOnly sets field value.
func (o *DolphinDBUser) SetReadOnly(v bool) {
	o.ReadOnly = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DolphinDBUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["isAdmin"] = o.IsAdmin
	toSerialize["groups"] = o.Groups
	toSerialize["readOnly"] = o.ReadOnly

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DolphinDBUser) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name     *string   `json:"name"`
		IsAdmin  *bool     `json:"isAdmin"`
		Groups   *[]string `json:"groups"`
		ReadOnly *bool     `json:"readOnly"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.IsAdmin == nil {
		return fmt.Errorf("required field isAdmin missing")
	}
	if all.Groups == nil {
		return fmt.Errorf("required field groups missing")
	}
	if all.ReadOnly == nil {
		return fmt.Errorf("required field readOnly missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "isAdmin", "groups", "readOnly"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.IsAdmin = *all.IsAdmin
	o.Groups = *all.Groups
	o.ReadOnly = *all.ReadOnly

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
