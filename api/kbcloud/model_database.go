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
	// Specify account name who can manage this database
	AccountName *string `json:"accountName,omitempty"`
	// Specify charsetName of database.
	Charset common.NullableString `json:"charset,omitempty"`
	// Specify collateName of database.
	Collate common.NullableString `json:"collate,omitempty"`
	// Specify ctypeName of database.
	Ctype common.NullableString `json:"ctype,omitempty"`
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

// GetCharset returns the Charset field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Database) GetCharset() string {
	if o == nil || o.Charset.Get() == nil {
		var ret string
		return ret
	}
	return *o.Charset.Get()
}

// GetCharsetOk returns a tuple with the Charset field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Database) GetCharsetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Charset.Get(), o.Charset.IsSet()
}

// HasCharset returns a boolean if a field has been set.
func (o *Database) HasCharset() bool {
	return o != nil && o.Charset.IsSet()
}

// SetCharset gets a reference to the given common.NullableString and assigns it to the Charset field.
func (o *Database) SetCharset(v string) {
	o.Charset.Set(&v)
}

// SetCharsetNil sets the value for Charset to be an explicit nil.
func (o *Database) SetCharsetNil() {
	o.Charset.Set(nil)
}

// UnsetCharset ensures that no value is present for Charset, not even an explicit nil.
func (o *Database) UnsetCharset() {
	o.Charset.Unset()
}

// GetCollate returns the Collate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Database) GetCollate() string {
	if o == nil || o.Collate.Get() == nil {
		var ret string
		return ret
	}
	return *o.Collate.Get()
}

// GetCollateOk returns a tuple with the Collate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Database) GetCollateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Collate.Get(), o.Collate.IsSet()
}

// HasCollate returns a boolean if a field has been set.
func (o *Database) HasCollate() bool {
	return o != nil && o.Collate.IsSet()
}

// SetCollate gets a reference to the given common.NullableString and assigns it to the Collate field.
func (o *Database) SetCollate(v string) {
	o.Collate.Set(&v)
}

// SetCollateNil sets the value for Collate to be an explicit nil.
func (o *Database) SetCollateNil() {
	o.Collate.Set(nil)
}

// UnsetCollate ensures that no value is present for Collate, not even an explicit nil.
func (o *Database) UnsetCollate() {
	o.Collate.Unset()
}

// GetCtype returns the Ctype field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Database) GetCtype() string {
	if o == nil || o.Ctype.Get() == nil {
		var ret string
		return ret
	}
	return *o.Ctype.Get()
}

// GetCtypeOk returns a tuple with the Ctype field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Database) GetCtypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ctype.Get(), o.Ctype.IsSet()
}

// HasCtype returns a boolean if a field has been set.
func (o *Database) HasCtype() bool {
	return o != nil && o.Ctype.IsSet()
}

// SetCtype gets a reference to the given common.NullableString and assigns it to the Ctype field.
func (o *Database) SetCtype(v string) {
	o.Ctype.Set(&v)
}

// SetCtypeNil sets the value for Ctype to be an explicit nil.
func (o *Database) SetCtypeNil() {
	o.Ctype.Set(nil)
}

// UnsetCtype ensures that no value is present for Ctype, not even an explicit nil.
func (o *Database) UnsetCtype() {
	o.Ctype.Unset()
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
	if o.Charset.IsSet() {
		toSerialize["charset"] = o.Charset.Get()
	}
	if o.Collate.IsSet() {
		toSerialize["collate"] = o.Collate.Get()
	}
	if o.Ctype.IsSet() {
		toSerialize["ctype"] = o.Ctype.Get()
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
		Component   *string               `json:"component,omitempty"`
		Name        *string               `json:"name"`
		AccountName *string               `json:"accountName,omitempty"`
		Charset     common.NullableString `json:"charset,omitempty"`
		Collate     common.NullableString `json:"collate,omitempty"`
		Ctype       common.NullableString `json:"ctype,omitempty"`
		Description *string               `json:"description,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"component", "name", "accountName", "charset", "collate", "ctype", "description"})
	} else {
		return err
	}
	o.Component = all.Component
	o.Name = *all.Name
	o.AccountName = all.AccountName
	o.Charset = all.Charset
	o.Collate = all.Collate
	o.Ctype = all.Ctype
	o.Description = all.Description

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
