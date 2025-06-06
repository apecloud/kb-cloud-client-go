// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ParameterProp struct {
	// The name of the parameter
	Name string `json:"name"`
	// The description of the parameter
	Description *string `json:"description,omitempty"`
	// The type of the parameter value
	Type string `json:"type"`
	// The value of the parameter, if parameter is not set in tpl, it's value equal to cue default value.
	Value interface{} `json:"value,omitempty"`
	// The value of the parameter, if parameter is not set in tpl, it's value equal to cue default value.
	Default interface{} `json:"default,omitempty"`
	// Whether the parameter requires a restart to take effect
	NeedRestart bool `json:"needRestart"`
	// Whether the parameter is an immutable parameter, immutable parameters cannot be modified
	Immutable bool `json:"immutable"`
	// The maximum value of the parameter
	Maximum common.NullableFloat64 `json:"maximum,omitempty"`
	// The minimum value of the parameter
	Minimum common.NullableFloat64 `json:"minimum,omitempty"`
	// The value options of the parameter
	Enum []interface{} `json:"enum,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewParameterProp instantiates a new ParameterProp object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewParameterProp(name string, typeVar string, needRestart bool, immutable bool) *ParameterProp {
	this := ParameterProp{}
	this.Name = name
	this.Type = typeVar
	this.NeedRestart = needRestart
	this.Immutable = immutable
	return &this
}

// NewParameterPropWithDefaults instantiates a new ParameterProp object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewParameterPropWithDefaults() *ParameterProp {
	this := ParameterProp{}
	return &this
}

// GetName returns the Name field value.
func (o *ParameterProp) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ParameterProp) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ParameterProp) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ParameterProp) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterProp) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ParameterProp) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ParameterProp) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value.
func (o *ParameterProp) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ParameterProp) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ParameterProp) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ParameterProp) GetValue() interface{} {
	if o == nil || o.Value == nil {
		var ret interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterProp) GetValueOk() (*interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ParameterProp) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given interface{} and assigns it to the Value field.
func (o *ParameterProp) SetValue(v interface{}) {
	o.Value = v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *ParameterProp) GetDefault() interface{} {
	if o == nil || o.Default == nil {
		var ret interface{}
		return ret
	}
	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterProp) GetDefaultOk() (*interface{}, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return &o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *ParameterProp) HasDefault() bool {
	return o != nil && o.Default != nil
}

// SetDefault gets a reference to the given interface{} and assigns it to the Default field.
func (o *ParameterProp) SetDefault(v interface{}) {
	o.Default = v
}

// GetNeedRestart returns the NeedRestart field value.
func (o *ParameterProp) GetNeedRestart() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.NeedRestart
}

// GetNeedRestartOk returns a tuple with the NeedRestart field value
// and a boolean to check if the value has been set.
func (o *ParameterProp) GetNeedRestartOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NeedRestart, true
}

// SetNeedRestart sets field value.
func (o *ParameterProp) SetNeedRestart(v bool) {
	o.NeedRestart = v
}

// GetImmutable returns the Immutable field value.
func (o *ParameterProp) GetImmutable() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Immutable
}

// GetImmutableOk returns a tuple with the Immutable field value
// and a boolean to check if the value has been set.
func (o *ParameterProp) GetImmutableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Immutable, true
}

// SetImmutable sets field value.
func (o *ParameterProp) SetImmutable(v bool) {
	o.Immutable = v
}

// GetMaximum returns the Maximum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParameterProp) GetMaximum() float64 {
	if o == nil || o.Maximum.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Maximum.Get()
}

// GetMaximumOk returns a tuple with the Maximum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ParameterProp) GetMaximumOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Maximum.Get(), o.Maximum.IsSet()
}

// HasMaximum returns a boolean if a field has been set.
func (o *ParameterProp) HasMaximum() bool {
	return o != nil && o.Maximum.IsSet()
}

// SetMaximum gets a reference to the given common.NullableFloat64 and assigns it to the Maximum field.
func (o *ParameterProp) SetMaximum(v float64) {
	o.Maximum.Set(&v)
}

// SetMaximumNil sets the value for Maximum to be an explicit nil.
func (o *ParameterProp) SetMaximumNil() {
	o.Maximum.Set(nil)
}

// UnsetMaximum ensures that no value is present for Maximum, not even an explicit nil.
func (o *ParameterProp) UnsetMaximum() {
	o.Maximum.Unset()
}

// GetMinimum returns the Minimum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParameterProp) GetMinimum() float64 {
	if o == nil || o.Minimum.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Minimum.Get()
}

// GetMinimumOk returns a tuple with the Minimum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ParameterProp) GetMinimumOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Minimum.Get(), o.Minimum.IsSet()
}

// HasMinimum returns a boolean if a field has been set.
func (o *ParameterProp) HasMinimum() bool {
	return o != nil && o.Minimum.IsSet()
}

// SetMinimum gets a reference to the given common.NullableFloat64 and assigns it to the Minimum field.
func (o *ParameterProp) SetMinimum(v float64) {
	o.Minimum.Set(&v)
}

// SetMinimumNil sets the value for Minimum to be an explicit nil.
func (o *ParameterProp) SetMinimumNil() {
	o.Minimum.Set(nil)
}

// UnsetMinimum ensures that no value is present for Minimum, not even an explicit nil.
func (o *ParameterProp) UnsetMinimum() {
	o.Minimum.Unset()
}

// GetEnum returns the Enum field value if set, zero value otherwise.
func (o *ParameterProp) GetEnum() []interface{} {
	if o == nil || o.Enum == nil {
		var ret []interface{}
		return ret
	}
	return o.Enum
}

// GetEnumOk returns a tuple with the Enum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterProp) GetEnumOk() (*[]interface{}, bool) {
	if o == nil || o.Enum == nil {
		return nil, false
	}
	return &o.Enum, true
}

// HasEnum returns a boolean if a field has been set.
func (o *ParameterProp) HasEnum() bool {
	return o != nil && o.Enum != nil
}

// SetEnum gets a reference to the given []interface{} and assigns it to the Enum field.
func (o *ParameterProp) SetEnum(v []interface{}) {
	o.Enum = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ParameterProp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	toSerialize["type"] = o.Type
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}
	toSerialize["needRestart"] = o.NeedRestart
	toSerialize["immutable"] = o.Immutable
	if o.Maximum.IsSet() {
		toSerialize["maximum"] = o.Maximum.Get()
	}
	if o.Minimum.IsSet() {
		toSerialize["minimum"] = o.Minimum.Get()
	}
	if o.Enum != nil {
		toSerialize["enum"] = o.Enum
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ParameterProp) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name        *string                `json:"name"`
		Description *string                `json:"description,omitempty"`
		Type        *string                `json:"type"`
		Value       interface{}            `json:"value,omitempty"`
		Default     interface{}            `json:"default,omitempty"`
		NeedRestart *bool                  `json:"needRestart"`
		Immutable   *bool                  `json:"immutable"`
		Maximum     common.NullableFloat64 `json:"maximum,omitempty"`
		Minimum     common.NullableFloat64 `json:"minimum,omitempty"`
		Enum        []interface{}          `json:"enum,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.NeedRestart == nil {
		return fmt.Errorf("required field needRestart missing")
	}
	if all.Immutable == nil {
		return fmt.Errorf("required field immutable missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "description", "type", "value", "default", "needRestart", "immutable", "maximum", "minimum", "enum"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Description = all.Description
	o.Type = *all.Type
	o.Value = all.Value
	o.Default = all.Default
	o.NeedRestart = *all.NeedRestart
	o.Immutable = *all.Immutable
	o.Maximum = all.Maximum
	o.Minimum = all.Minimum
	o.Enum = all.Enum

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
