// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// InstanceTemplateEnv One fillable env var on an instance template, for frontend rendering.
type InstanceTemplateEnv struct {
	// Env var name written to the instance template.
	Name        string                `json:"name"`
	Label       *LocalizedDescription `json:"label,omitempty"`
	Description *LocalizedDescription `json:"description,omitempty"`
	// If true, create requests must set this env on every template.
	Required *bool `json:"required,omitempty"`
	// Default value shown by the frontend.
	Default *string `json:"default,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInstanceTemplateEnv instantiates a new InstanceTemplateEnv object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInstanceTemplateEnv(name string) *InstanceTemplateEnv {
	this := InstanceTemplateEnv{}
	this.Name = name
	return &this
}

// NewInstanceTemplateEnvWithDefaults instantiates a new InstanceTemplateEnv object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInstanceTemplateEnvWithDefaults() *InstanceTemplateEnv {
	this := InstanceTemplateEnv{}
	return &this
}

// GetName returns the Name field value.
func (o *InstanceTemplateEnv) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InstanceTemplateEnv) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *InstanceTemplateEnv) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *InstanceTemplateEnv) GetLabel() LocalizedDescription {
	if o == nil || o.Label == nil {
		var ret LocalizedDescription
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTemplateEnv) GetLabelOk() (*LocalizedDescription, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *InstanceTemplateEnv) HasLabel() bool {
	return o != nil && o.Label != nil
}

// SetLabel gets a reference to the given LocalizedDescription and assigns it to the Label field.
func (o *InstanceTemplateEnv) SetLabel(v LocalizedDescription) {
	o.Label = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InstanceTemplateEnv) GetDescription() LocalizedDescription {
	if o == nil || o.Description == nil {
		var ret LocalizedDescription
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTemplateEnv) GetDescriptionOk() (*LocalizedDescription, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InstanceTemplateEnv) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given LocalizedDescription and assigns it to the Description field.
func (o *InstanceTemplateEnv) SetDescription(v LocalizedDescription) {
	o.Description = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *InstanceTemplateEnv) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTemplateEnv) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *InstanceTemplateEnv) HasRequired() bool {
	return o != nil && o.Required != nil
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *InstanceTemplateEnv) SetRequired(v bool) {
	o.Required = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *InstanceTemplateEnv) GetDefault() string {
	if o == nil || o.Default == nil {
		var ret string
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTemplateEnv) GetDefaultOk() (*string, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *InstanceTemplateEnv) HasDefault() bool {
	return o != nil && o.Default != nil
}

// SetDefault gets a reference to the given string and assigns it to the Default field.
func (o *InstanceTemplateEnv) SetDefault(v string) {
	o.Default = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o InstanceTemplateEnv) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InstanceTemplateEnv) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name        *string               `json:"name"`
		Label       *LocalizedDescription `json:"label,omitempty"`
		Description *LocalizedDescription `json:"description,omitempty"`
		Required    *bool                 `json:"required,omitempty"`
		Default     *string               `json:"default,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "label", "description", "required", "default"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = *all.Name
	if all.Label != nil && all.Label.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Label = all.Label
	if all.Description != nil && all.Description.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Description = all.Description
	o.Required = all.Required
	o.Default = all.Default

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
