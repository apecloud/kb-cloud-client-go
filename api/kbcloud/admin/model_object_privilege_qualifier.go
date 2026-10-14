// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ObjectPrivilegeQualifier struct {
	Name         string  `json:"name"`
	Required     *bool   `json:"required,omitempty"`
	DefaultValue *string `json:"defaultValue,omitempty"`
	Description  *string `json:"description,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObjectPrivilegeQualifier instantiates a new ObjectPrivilegeQualifier object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObjectPrivilegeQualifier(name string) *ObjectPrivilegeQualifier {
	this := ObjectPrivilegeQualifier{}
	this.Name = name
	return &this
}

// NewObjectPrivilegeQualifierWithDefaults instantiates a new ObjectPrivilegeQualifier object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObjectPrivilegeQualifierWithDefaults() *ObjectPrivilegeQualifier {
	this := ObjectPrivilegeQualifier{}
	return &this
}

// GetName returns the Name field value.
func (o *ObjectPrivilegeQualifier) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeQualifier) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ObjectPrivilegeQualifier) SetName(v string) {
	o.Name = v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *ObjectPrivilegeQualifier) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeQualifier) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *ObjectPrivilegeQualifier) HasRequired() bool {
	return o != nil && o.Required != nil
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *ObjectPrivilegeQualifier) SetRequired(v bool) {
	o.Required = &v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *ObjectPrivilegeQualifier) GetDefaultValue() string {
	if o == nil || o.DefaultValue == nil {
		var ret string
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeQualifier) GetDefaultValueOk() (*string, bool) {
	if o == nil || o.DefaultValue == nil {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *ObjectPrivilegeQualifier) HasDefaultValue() bool {
	return o != nil && o.DefaultValue != nil
}

// SetDefaultValue gets a reference to the given string and assigns it to the DefaultValue field.
func (o *ObjectPrivilegeQualifier) SetDefaultValue(v string) {
	o.DefaultValue = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ObjectPrivilegeQualifier) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeQualifier) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ObjectPrivilegeQualifier) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ObjectPrivilegeQualifier) SetDescription(v string) {
	o.Description = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObjectPrivilegeQualifier) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}
	if o.DefaultValue != nil {
		toSerialize["defaultValue"] = o.DefaultValue
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
func (o *ObjectPrivilegeQualifier) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name         *string `json:"name"`
		Required     *bool   `json:"required,omitempty"`
		DefaultValue *string `json:"defaultValue,omitempty"`
		Description  *string `json:"description,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "required", "defaultValue", "description"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Required = all.Required
	o.DefaultValue = all.DefaultValue
	o.Description = all.Description

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
