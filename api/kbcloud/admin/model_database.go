// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
)

// Database Cluster database information
type Database struct {
	// component type
	Component *string `json:"component,omitempty"`
	// Specify the name of database, which must be unique.
	Name string `json:"name"`
	// Specify account name who can manage this database
	AccountName *string `json:"accountName,omitempty"`
	// Specify charsetName of database.
	Charset *string `json:"charset,omitempty"`
	// Description of the database.
	Description *string `json:"description,omitempty"`
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

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *Database) GetAccountName() string {
	if o == nil || o.AccountName == nil {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Database) GetAccountNameOk() (*string, bool) {
	if o == nil || o.AccountName == nil {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *Database) HasAccountName() bool {
	return o != nil && o.AccountName != nil
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *Database) SetAccountName(v string) {
	o.AccountName = &v
}

// GetCharset returns the Charset field value if set, zero value otherwise.
func (o *Database) GetCharset() string {
	if o == nil || o.Charset == nil {
		var ret string
		return ret
	}
	return *o.Charset
}

// GetCharsetOk returns a tuple with the Charset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Database) GetCharsetOk() (*string, bool) {
	if o == nil || o.Charset == nil {
		return nil, false
	}
	return o.Charset, true
}

// HasCharset returns a boolean if a field has been set.
func (o *Database) HasCharset() bool {
	return o != nil && o.Charset != nil
}

// SetCharset gets a reference to the given string and assigns it to the Charset field.
func (o *Database) SetCharset(v string) {
	o.Charset = &v
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
	if o.AccountName != nil {
		toSerialize["accountName"] = o.AccountName
	}
	if o.Charset != nil {
		toSerialize["charset"] = o.Charset
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
func (o *Database) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Component   *string `json:"component,omitempty"`
		Name        *string `json:"name"`
		AccountName *string `json:"accountName,omitempty"`
		Charset     *string `json:"charset,omitempty"`
		Description *string `json:"description,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"component", "name", "accountName", "charset", "description"})
	} else {
		return err
	}
	o.Component = all.Component
	o.Name = *all.Name
	o.AccountName = all.AccountName
	o.Charset = all.Charset
	o.Description = all.Description

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
