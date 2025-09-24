// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ConfigEntry struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
	// Whether the configuration is the default value
	IsDefault *bool `json:"isDefault,omitempty"`
	ReadOnly  *bool `json:"readOnly,omitempty"`
	// Whether the configuration is sensitive
	Sensitive *bool `json:"sensitive,omitempty"`
	// Configuration description
	Description *string `json:"description,omitempty"`
	// Valid values for the configuration
	ValidValues []string `json:"validValues,omitempty"`
	// Configuration type
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewConfigEntry instantiates a new ConfigEntry object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewConfigEntry() *ConfigEntry {
	this := ConfigEntry{}
	return &this
}

// NewConfigEntryWithDefaults instantiates a new ConfigEntry object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewConfigEntryWithDefaults() *ConfigEntry {
	this := ConfigEntry{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConfigEntry) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigEntry) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConfigEntry) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConfigEntry) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ConfigEntry) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigEntry) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ConfigEntry) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ConfigEntry) SetValue(v string) {
	o.Value = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *ConfigEntry) GetIsDefault() bool {
	if o == nil || o.IsDefault == nil {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigEntry) GetIsDefaultOk() (*bool, bool) {
	if o == nil || o.IsDefault == nil {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *ConfigEntry) HasIsDefault() bool {
	return o != nil && o.IsDefault != nil
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *ConfigEntry) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *ConfigEntry) GetReadOnly() bool {
	if o == nil || o.ReadOnly == nil {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigEntry) GetReadOnlyOk() (*bool, bool) {
	if o == nil || o.ReadOnly == nil {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *ConfigEntry) HasReadOnly() bool {
	return o != nil && o.ReadOnly != nil
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *ConfigEntry) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

// GetSensitive returns the Sensitive field value if set, zero value otherwise.
func (o *ConfigEntry) GetSensitive() bool {
	if o == nil || o.Sensitive == nil {
		var ret bool
		return ret
	}
	return *o.Sensitive
}

// GetSensitiveOk returns a tuple with the Sensitive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigEntry) GetSensitiveOk() (*bool, bool) {
	if o == nil || o.Sensitive == nil {
		return nil, false
	}
	return o.Sensitive, true
}

// HasSensitive returns a boolean if a field has been set.
func (o *ConfigEntry) HasSensitive() bool {
	return o != nil && o.Sensitive != nil
}

// SetSensitive gets a reference to the given bool and assigns it to the Sensitive field.
func (o *ConfigEntry) SetSensitive(v bool) {
	o.Sensitive = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConfigEntry) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigEntry) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConfigEntry) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ConfigEntry) SetDescription(v string) {
	o.Description = &v
}

// GetValidValues returns the ValidValues field value if set, zero value otherwise.
func (o *ConfigEntry) GetValidValues() []string {
	if o == nil || o.ValidValues == nil {
		var ret []string
		return ret
	}
	return o.ValidValues
}

// GetValidValuesOk returns a tuple with the ValidValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigEntry) GetValidValuesOk() (*[]string, bool) {
	if o == nil || o.ValidValues == nil {
		return nil, false
	}
	return &o.ValidValues, true
}

// HasValidValues returns a boolean if a field has been set.
func (o *ConfigEntry) HasValidValues() bool {
	return o != nil && o.ValidValues != nil
}

// SetValidValues gets a reference to the given []string and assigns it to the ValidValues field.
func (o *ConfigEntry) SetValidValues(v []string) {
	o.ValidValues = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ConfigEntry) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigEntry) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ConfigEntry) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ConfigEntry) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ConfigEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.IsDefault != nil {
		toSerialize["isDefault"] = o.IsDefault
	}
	if o.ReadOnly != nil {
		toSerialize["readOnly"] = o.ReadOnly
	}
	if o.Sensitive != nil {
		toSerialize["sensitive"] = o.Sensitive
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ValidValues != nil {
		toSerialize["validValues"] = o.ValidValues
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ConfigEntry) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name        *string  `json:"name,omitempty"`
		Value       *string  `json:"value,omitempty"`
		IsDefault   *bool    `json:"isDefault,omitempty"`
		ReadOnly    *bool    `json:"readOnly,omitempty"`
		Sensitive   *bool    `json:"sensitive,omitempty"`
		Description *string  `json:"description,omitempty"`
		ValidValues []string `json:"validValues,omitempty"`
		Type        *string  `json:"type,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "value", "isDefault", "readOnly", "sensitive", "description", "validValues", "type"})
	} else {
		return err
	}
	o.Name = all.Name
	o.Value = all.Value
	o.IsDefault = all.IsDefault
	o.ReadOnly = all.ReadOnly
	o.Sensitive = all.Sensitive
	o.Description = all.Description
	o.ValidValues = all.ValidValues
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
