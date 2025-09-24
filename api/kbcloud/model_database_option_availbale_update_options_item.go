// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DatabaseOptionAvailbaleUpdateOptionsItem struct {
	Name        *string               `json:"name,omitempty"`
	Type        *string               `json:"type,omitempty"`
	Description *LocalizedDescription `json:"description,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDatabaseOptionAvailbaleUpdateOptionsItem instantiates a new DatabaseOptionAvailbaleUpdateOptionsItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatabaseOptionAvailbaleUpdateOptionsItem() *DatabaseOptionAvailbaleUpdateOptionsItem {
	this := DatabaseOptionAvailbaleUpdateOptionsItem{}
	return &this
}

// NewDatabaseOptionAvailbaleUpdateOptionsItemWithDefaults instantiates a new DatabaseOptionAvailbaleUpdateOptionsItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatabaseOptionAvailbaleUpdateOptionsItemWithDefaults() *DatabaseOptionAvailbaleUpdateOptionsItem {
	this := DatabaseOptionAvailbaleUpdateOptionsItem{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DatabaseOptionAvailbaleUpdateOptionsItem) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseOptionAvailbaleUpdateOptionsItem) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DatabaseOptionAvailbaleUpdateOptionsItem) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DatabaseOptionAvailbaleUpdateOptionsItem) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DatabaseOptionAvailbaleUpdateOptionsItem) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseOptionAvailbaleUpdateOptionsItem) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DatabaseOptionAvailbaleUpdateOptionsItem) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DatabaseOptionAvailbaleUpdateOptionsItem) SetType(v string) {
	o.Type = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DatabaseOptionAvailbaleUpdateOptionsItem) GetDescription() LocalizedDescription {
	if o == nil || o.Description == nil {
		var ret LocalizedDescription
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseOptionAvailbaleUpdateOptionsItem) GetDescriptionOk() (*LocalizedDescription, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DatabaseOptionAvailbaleUpdateOptionsItem) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given LocalizedDescription and assigns it to the Description field.
func (o *DatabaseOptionAvailbaleUpdateOptionsItem) SetDescription(v LocalizedDescription) {
	o.Description = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DatabaseOptionAvailbaleUpdateOptionsItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
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
func (o *DatabaseOptionAvailbaleUpdateOptionsItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name        *string               `json:"name,omitempty"`
		Type        *string               `json:"type,omitempty"`
		Description *LocalizedDescription `json:"description,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "type", "description"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = all.Name
	o.Type = all.Type
	if all.Description != nil && all.Description.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Description = all.Description

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
