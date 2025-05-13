// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// Database Cluster database information
type Database struct {
	// component type
	Component *string `json:"component,omitempty"`
	// Specify the name of database, which must be unique.
	Name string `json:"name"`
	// Description of the database.
	Description *string `json:"description,omitempty"`
	// Specify the options of database.
	Options map[string]interface{} `json:"options,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDatabase instantiates a new Database object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatabase(name string) *Database {
	this := Database{}
	this.Name = name
	return &this
}

// NewDatabaseWithDefaults instantiates a new Database object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatabaseWithDefaults() *Database {
	this := Database{}
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *Database) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Database) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *Database) HasComponent() bool {
	return o != nil && o.Component != nil
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *Database) SetComponent(v string) {
	o.Component = &v
}

// GetName returns the Name field value.
func (o *Database) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Database) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *Database) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Database) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Database) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Database) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Database) SetDescription(v string) {
	o.Description = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *Database) GetOptions() map[string]interface{} {
	if o == nil || o.Options == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Database) GetOptionsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return &o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *Database) HasOptions() bool {
	return o != nil && o.Options != nil
}

// SetOptions gets a reference to the given map[string]interface{} and assigns it to the Options field.
func (o *Database) SetOptions(v map[string]interface{}) {
	o.Options = v
}

// MarshalJSON serializes the struct using spec logic.
func (o Database) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	toSerialize["name"] = o.Name
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Database) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Component   *string                `json:"component,omitempty"`
		Name        *string                `json:"name"`
		Description *string                `json:"description,omitempty"`
		Options     map[string]interface{} `json:"options,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"component", "name", "description", "options"})
	} else {
		return err
	}
	o.Component = all.Component
	o.Name = *all.Name
	o.Description = all.Description
	o.Options = all.Options

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
