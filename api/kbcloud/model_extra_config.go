// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ExtraConfig struct {
	Name  *string            `json:"name,omitempty"`
	Title *InternationalDesc `json:"title,omitempty"`
	// the type of config. The difference between 'text' and 'string' is that 'text' refers to a longer string, which may need to be displayed using a textarea on the frontend rather than an input field.
	Type              *ConfigType           `json:"type,omitempty"`
	DefaultValue      common.NullableString `json:"defaultValue,omitempty"`
	IsRequired        *bool                 `json:"isRequired,omitempty"`
	ConnectionCfgType *ConnectionCfgType    `json:"connectionCfgType,omitempty"`
	// Whether to display this config only when the endpoint type is custom. If true, it will not be displayed when the endpoint type is kubeblocks.
	IsDisplayOnlyCustomEndpoint common.NullableBool `json:"isDisplayOnlyCustomEndpoint,omitempty"`
	// The display condition of this config. If empty, it will always be displayed; otherwise, it will be displayed only when the config corresponding to displayWith (of boolean type) is true.
	DisplayWith common.NullableString `json:"displayWith,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExtraConfig instantiates a new ExtraConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExtraConfig() *ExtraConfig {
	this := ExtraConfig{}
	return &this
}

// NewExtraConfigWithDefaults instantiates a new ExtraConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExtraConfigWithDefaults() *ExtraConfig {
	this := ExtraConfig{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExtraConfig) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtraConfig) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExtraConfig) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExtraConfig) SetName(v string) {
	o.Name = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ExtraConfig) GetTitle() InternationalDesc {
	if o == nil || o.Title == nil {
		var ret InternationalDesc
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtraConfig) GetTitleOk() (*InternationalDesc, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ExtraConfig) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given InternationalDesc and assigns it to the Title field.
func (o *ExtraConfig) SetTitle(v InternationalDesc) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ExtraConfig) GetType() ConfigType {
	if o == nil || o.Type == nil {
		var ret ConfigType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtraConfig) GetTypeOk() (*ConfigType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ExtraConfig) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given ConfigType and assigns it to the Type field.
func (o *ExtraConfig) SetType(v ConfigType) {
	o.Type = &v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExtraConfig) GetDefaultValue() string {
	if o == nil || o.DefaultValue.Get() == nil {
		var ret string
		return ret
	}
	return *o.DefaultValue.Get()
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExtraConfig) GetDefaultValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultValue.Get(), o.DefaultValue.IsSet()
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *ExtraConfig) HasDefaultValue() bool {
	return o != nil && o.DefaultValue.IsSet()
}

// SetDefaultValue gets a reference to the given common.NullableString and assigns it to the DefaultValue field.
func (o *ExtraConfig) SetDefaultValue(v string) {
	o.DefaultValue.Set(&v)
}

// SetDefaultValueNil sets the value for DefaultValue to be an explicit nil.
func (o *ExtraConfig) SetDefaultValueNil() {
	o.DefaultValue.Set(nil)
}

// UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil.
func (o *ExtraConfig) UnsetDefaultValue() {
	o.DefaultValue.Unset()
}

// GetIsRequired returns the IsRequired field value if set, zero value otherwise.
func (o *ExtraConfig) GetIsRequired() bool {
	if o == nil || o.IsRequired == nil {
		var ret bool
		return ret
	}
	return *o.IsRequired
}

// GetIsRequiredOk returns a tuple with the IsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtraConfig) GetIsRequiredOk() (*bool, bool) {
	if o == nil || o.IsRequired == nil {
		return nil, false
	}
	return o.IsRequired, true
}

// HasIsRequired returns a boolean if a field has been set.
func (o *ExtraConfig) HasIsRequired() bool {
	return o != nil && o.IsRequired != nil
}

// SetIsRequired gets a reference to the given bool and assigns it to the IsRequired field.
func (o *ExtraConfig) SetIsRequired(v bool) {
	o.IsRequired = &v
}

// GetConnectionCfgType returns the ConnectionCfgType field value if set, zero value otherwise.
func (o *ExtraConfig) GetConnectionCfgType() ConnectionCfgType {
	if o == nil || o.ConnectionCfgType == nil {
		var ret ConnectionCfgType
		return ret
	}
	return *o.ConnectionCfgType
}

// GetConnectionCfgTypeOk returns a tuple with the ConnectionCfgType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtraConfig) GetConnectionCfgTypeOk() (*ConnectionCfgType, bool) {
	if o == nil || o.ConnectionCfgType == nil {
		return nil, false
	}
	return o.ConnectionCfgType, true
}

// HasConnectionCfgType returns a boolean if a field has been set.
func (o *ExtraConfig) HasConnectionCfgType() bool {
	return o != nil && o.ConnectionCfgType != nil
}

// SetConnectionCfgType gets a reference to the given ConnectionCfgType and assigns it to the ConnectionCfgType field.
func (o *ExtraConfig) SetConnectionCfgType(v ConnectionCfgType) {
	o.ConnectionCfgType = &v
}

// GetIsDisplayOnlyCustomEndpoint returns the IsDisplayOnlyCustomEndpoint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExtraConfig) GetIsDisplayOnlyCustomEndpoint() bool {
	if o == nil || o.IsDisplayOnlyCustomEndpoint.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsDisplayOnlyCustomEndpoint.Get()
}

// GetIsDisplayOnlyCustomEndpointOk returns a tuple with the IsDisplayOnlyCustomEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExtraConfig) GetIsDisplayOnlyCustomEndpointOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsDisplayOnlyCustomEndpoint.Get(), o.IsDisplayOnlyCustomEndpoint.IsSet()
}

// HasIsDisplayOnlyCustomEndpoint returns a boolean if a field has been set.
func (o *ExtraConfig) HasIsDisplayOnlyCustomEndpoint() bool {
	return o != nil && o.IsDisplayOnlyCustomEndpoint.IsSet()
}

// SetIsDisplayOnlyCustomEndpoint gets a reference to the given common.NullableBool and assigns it to the IsDisplayOnlyCustomEndpoint field.
func (o *ExtraConfig) SetIsDisplayOnlyCustomEndpoint(v bool) {
	o.IsDisplayOnlyCustomEndpoint.Set(&v)
}

// SetIsDisplayOnlyCustomEndpointNil sets the value for IsDisplayOnlyCustomEndpoint to be an explicit nil.
func (o *ExtraConfig) SetIsDisplayOnlyCustomEndpointNil() {
	o.IsDisplayOnlyCustomEndpoint.Set(nil)
}

// UnsetIsDisplayOnlyCustomEndpoint ensures that no value is present for IsDisplayOnlyCustomEndpoint, not even an explicit nil.
func (o *ExtraConfig) UnsetIsDisplayOnlyCustomEndpoint() {
	o.IsDisplayOnlyCustomEndpoint.Unset()
}

// GetDisplayWith returns the DisplayWith field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExtraConfig) GetDisplayWith() string {
	if o == nil || o.DisplayWith.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayWith.Get()
}

// GetDisplayWithOk returns a tuple with the DisplayWith field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExtraConfig) GetDisplayWithOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayWith.Get(), o.DisplayWith.IsSet()
}

// HasDisplayWith returns a boolean if a field has been set.
func (o *ExtraConfig) HasDisplayWith() bool {
	return o != nil && o.DisplayWith.IsSet()
}

// SetDisplayWith gets a reference to the given common.NullableString and assigns it to the DisplayWith field.
func (o *ExtraConfig) SetDisplayWith(v string) {
	o.DisplayWith.Set(&v)
}

// SetDisplayWithNil sets the value for DisplayWith to be an explicit nil.
func (o *ExtraConfig) SetDisplayWithNil() {
	o.DisplayWith.Set(nil)
}

// UnsetDisplayWith ensures that no value is present for DisplayWith, not even an explicit nil.
func (o *ExtraConfig) UnsetDisplayWith() {
	o.DisplayWith.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ExtraConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.DefaultValue.IsSet() {
		toSerialize["defaultValue"] = o.DefaultValue.Get()
	}
	if o.IsRequired != nil {
		toSerialize["isRequired"] = o.IsRequired
	}
	if o.ConnectionCfgType != nil {
		toSerialize["connectionCfgType"] = o.ConnectionCfgType
	}
	if o.IsDisplayOnlyCustomEndpoint.IsSet() {
		toSerialize["isDisplayOnlyCustomEndpoint"] = o.IsDisplayOnlyCustomEndpoint.Get()
	}
	if o.DisplayWith.IsSet() {
		toSerialize["displayWith"] = o.DisplayWith.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExtraConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name                        *string               `json:"name,omitempty"`
		Title                       *InternationalDesc    `json:"title,omitempty"`
		Type                        *ConfigType           `json:"type,omitempty"`
		DefaultValue                common.NullableString `json:"defaultValue,omitempty"`
		IsRequired                  *bool                 `json:"isRequired,omitempty"`
		ConnectionCfgType           *ConnectionCfgType    `json:"connectionCfgType,omitempty"`
		IsDisplayOnlyCustomEndpoint common.NullableBool   `json:"isDisplayOnlyCustomEndpoint,omitempty"`
		DisplayWith                 common.NullableString `json:"displayWith,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "title", "type", "defaultValue", "isRequired", "connectionCfgType", "isDisplayOnlyCustomEndpoint", "displayWith"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = all.Name
	if all.Title != nil && all.Title.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Title = all.Title
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}
	o.DefaultValue = all.DefaultValue
	o.IsRequired = all.IsRequired
	if all.ConnectionCfgType != nil && !all.ConnectionCfgType.IsValid() {
		hasInvalidField = true
	} else {
		o.ConnectionCfgType = all.ConnectionCfgType
	}
	o.IsDisplayOnlyCustomEndpoint = all.IsDisplayOnlyCustomEndpoint
	o.DisplayWith = all.DisplayWith

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
