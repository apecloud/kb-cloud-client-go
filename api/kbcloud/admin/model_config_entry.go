// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type ConfigEntry struct {
	Name      *string `json:"name,omitempty"`
	Value     *string `json:"value,omitempty"`
	Default   *bool   `json:"default,omitempty"`
	ReadOnly  *bool   `json:"readOnly,omitempty"`
	Sensitive *bool   `json:"sensitive,omitempty"`
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

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *ConfigEntry) GetDefault() bool {
	if o == nil || o.Default == nil {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigEntry) GetDefaultOk() (*bool, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *ConfigEntry) HasDefault() bool {
	return o != nil && o.Default != nil
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *ConfigEntry) SetDefault(v bool) {
	o.Default = &v
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
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}
	if o.ReadOnly != nil {
		toSerialize["readOnly"] = o.ReadOnly
	}
	if o.Sensitive != nil {
		toSerialize["sensitive"] = o.Sensitive
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ConfigEntry) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name      *string `json:"name,omitempty"`
		Value     *string `json:"value,omitempty"`
		Default   *bool   `json:"default,omitempty"`
		ReadOnly  *bool   `json:"readOnly,omitempty"`
		Sensitive *bool   `json:"sensitive,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "value", "default", "readOnly", "sensitive"})
	} else {
		return err
	}
	o.Name = all.Name
	o.Value = all.Value
	o.Default = all.Default
	o.ReadOnly = all.ReadOnly
	o.Sensitive = all.Sensitive

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
