// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// AdminPermission Administrator Permission information
type AdminPermission struct {
	// The name of the permission
	Name string `json:"name"`
	// The group of the permission
	Group *string `json:"group,omitempty"`
	// The description of the permission
	Description *string `json:"description,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAdminPermission instantiates a new AdminPermission object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAdminPermission(name string) *AdminPermission {
	this := AdminPermission{}
	this.Name = name
	return &this
}

// NewAdminPermissionWithDefaults instantiates a new AdminPermission object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAdminPermissionWithDefaults() *AdminPermission {
	this := AdminPermission{}
	return &this
}

// GetName returns the Name field value.
func (o *AdminPermission) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AdminPermission) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *AdminPermission) SetName(v string) {
	o.Name = v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *AdminPermission) GetGroup() string {
	if o == nil || o.Group == nil {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminPermission) GetGroupOk() (*string, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *AdminPermission) HasGroup() bool {
	return o != nil && o.Group != nil
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *AdminPermission) SetGroup(v string) {
	o.Group = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AdminPermission) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminPermission) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AdminPermission) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AdminPermission) SetDescription(v string) {
	o.Description = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AdminPermission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	if o.Group != nil {
		toSerialize["group"] = o.Group
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AdminPermission) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name        *string `json:"name"`
		Group       *string `json:"group,omitempty"`
		Description *string `json:"description,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "group", "description"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Group = all.Group
	o.Description = all.Description

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
