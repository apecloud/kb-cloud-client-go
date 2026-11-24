// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DolphinDBUserCreate struct {
	Name string `json:"name"`
	// Initial password. Never returned in responses.
	Password string   `json:"password"`
	IsAdmin  *bool    `json:"isAdmin,omitempty"`
	Groups   []string `json:"groups,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDolphinDBUserCreate instantiates a new DolphinDBUserCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDolphinDBUserCreate(name string, password string) *DolphinDBUserCreate {
	this := DolphinDBUserCreate{}
	this.Name = name
	this.Password = password
	var isAdmin bool = false
	this.IsAdmin = &isAdmin
	return &this
}

// NewDolphinDBUserCreateWithDefaults instantiates a new DolphinDBUserCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDolphinDBUserCreateWithDefaults() *DolphinDBUserCreate {
	this := DolphinDBUserCreate{}
	var isAdmin bool = false
	this.IsAdmin = &isAdmin
	return &this
}

// GetName returns the Name field value.
func (o *DolphinDBUserCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DolphinDBUserCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *DolphinDBUserCreate) SetName(v string) {
	o.Name = v
}

// GetPassword returns the Password field value.
func (o *DolphinDBUserCreate) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *DolphinDBUserCreate) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value.
func (o *DolphinDBUserCreate) SetPassword(v string) {
	o.Password = v
}

// GetIsAdmin returns the IsAdmin field value if set, zero value otherwise.
func (o *DolphinDBUserCreate) GetIsAdmin() bool {
	if o == nil || o.IsAdmin == nil {
		var ret bool
		return ret
	}
	return *o.IsAdmin
}

// GetIsAdminOk returns a tuple with the IsAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBUserCreate) GetIsAdminOk() (*bool, bool) {
	if o == nil || o.IsAdmin == nil {
		return nil, false
	}
	return o.IsAdmin, true
}

// HasIsAdmin returns a boolean if a field has been set.
func (o *DolphinDBUserCreate) HasIsAdmin() bool {
	return o != nil && o.IsAdmin != nil
}

// SetIsAdmin gets a reference to the given bool and assigns it to the IsAdmin field.
func (o *DolphinDBUserCreate) SetIsAdmin(v bool) {
	o.IsAdmin = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *DolphinDBUserCreate) GetGroups() []string {
	if o == nil || o.Groups == nil {
		var ret []string
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBUserCreate) GetGroupsOk() (*[]string, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return &o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *DolphinDBUserCreate) HasGroups() bool {
	return o != nil && o.Groups != nil
}

// SetGroups gets a reference to the given []string and assigns it to the Groups field.
func (o *DolphinDBUserCreate) SetGroups(v []string) {
	o.Groups = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DolphinDBUserCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["password"] = o.Password
	if o.IsAdmin != nil {
		toSerialize["isAdmin"] = o.IsAdmin
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DolphinDBUserCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name     *string  `json:"name"`
		Password *string  `json:"password"`
		IsAdmin  *bool    `json:"isAdmin,omitempty"`
		Groups   []string `json:"groups,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Password == nil {
		return fmt.Errorf("required field password missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "password", "isAdmin", "groups"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Password = *all.Password
	o.IsAdmin = all.IsAdmin
	o.Groups = all.Groups

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
