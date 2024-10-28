// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ParameterSpec Cluster parameter specification

type ParameterSpec struct {
	// The name of the parameter
	Name string `json:"name"`
	// The description of the parameter
	Description string `json:"description"`
	// The type of the parameter value
	Type string `json:"type"`
	// The default value of the parameter
	Default map[string]interface{} `json:"default"`
	// Whether the parameter requires a restart to take effect
	NeedRestart bool `json:"needRestart"`
	// Whether the parameter is an immutable parameter, immutable parameters cannot be modified
	Immutable bool `json:"immutable"`
	// The maximum value of the parameter
	Maximum float64 `json:"maximum"`
	// The minimum value of the parameter
	Minimum float64 `json:"minimum"`
	// The value options of the parameter
	Enum []map[string]interface{} `json:"enum"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewParameterSpec instantiates a new ParameterSpec object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewParameterSpec(name string, description string, typeVar string, defaultVar map[string]interface{}, needRestart bool, immutable bool, maximum float64, minimum float64, enum []map[string]interface{}) *ParameterSpec {
	this := ParameterSpec{}
	this.Name = name
	this.Description = description
	this.Type = typeVar
	this.Default = defaultVar
	this.NeedRestart = needRestart
	this.Immutable = immutable
	this.Maximum = maximum
	this.Minimum = minimum
	this.Enum = enum
	return &this
}

// NewParameterSpecWithDefaults instantiates a new ParameterSpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewParameterSpecWithDefaults() *ParameterSpec {
	this := ParameterSpec{}
	return &this
}

// GetName returns the Name field value.
func (o *ParameterSpec) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ParameterSpec) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ParameterSpec) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value.
func (o *ParameterSpec) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ParameterSpec) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *ParameterSpec) SetDescription(v string) {
	o.Description = v
}

// GetType returns the Type field value.
func (o *ParameterSpec) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ParameterSpec) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ParameterSpec) SetType(v string) {
	o.Type = v
}

// GetDefault returns the Default field value.
func (o *ParameterSpec) GetDefault() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value
// and a boolean to check if the value has been set.
func (o *ParameterSpec) GetDefaultOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Default, true
}

// SetDefault sets field value.
func (o *ParameterSpec) SetDefault(v map[string]interface{}) {
	o.Default = v
}

// GetNeedRestart returns the NeedRestart field value.
func (o *ParameterSpec) GetNeedRestart() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.NeedRestart
}

// GetNeedRestartOk returns a tuple with the NeedRestart field value
// and a boolean to check if the value has been set.
func (o *ParameterSpec) GetNeedRestartOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NeedRestart, true
}

// SetNeedRestart sets field value.
func (o *ParameterSpec) SetNeedRestart(v bool) {
	o.NeedRestart = v
}

// GetImmutable returns the Immutable field value.
func (o *ParameterSpec) GetImmutable() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Immutable
}

// GetImmutableOk returns a tuple with the Immutable field value
// and a boolean to check if the value has been set.
func (o *ParameterSpec) GetImmutableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Immutable, true
}

// SetImmutable sets field value.
func (o *ParameterSpec) SetImmutable(v bool) {
	o.Immutable = v
}

// GetMaximum returns the Maximum field value.
func (o *ParameterSpec) GetMaximum() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Maximum
}

// GetMaximumOk returns a tuple with the Maximum field value
// and a boolean to check if the value has been set.
func (o *ParameterSpec) GetMaximumOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Maximum, true
}

// SetMaximum sets field value.
func (o *ParameterSpec) SetMaximum(v float64) {
	o.Maximum = v
}

// GetMinimum returns the Minimum field value.
func (o *ParameterSpec) GetMinimum() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Minimum
}

// GetMinimumOk returns a tuple with the Minimum field value
// and a boolean to check if the value has been set.
func (o *ParameterSpec) GetMinimumOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Minimum, true
}

// SetMinimum sets field value.
func (o *ParameterSpec) SetMinimum(v float64) {
	o.Minimum = v
}

// GetEnum returns the Enum field value.
func (o *ParameterSpec) GetEnum() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Enum
}

// GetEnumOk returns a tuple with the Enum field value
// and a boolean to check if the value has been set.
func (o *ParameterSpec) GetEnumOk() (*[]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enum, true
}

// SetEnum sets field value.
func (o *ParameterSpec) SetEnum(v []map[string]interface{}) {
	o.Enum = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ParameterSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["type"] = o.Type
	toSerialize["default"] = o.Default
	toSerialize["needRestart"] = o.NeedRestart
	toSerialize["immutable"] = o.Immutable
	toSerialize["maximum"] = o.Maximum
	toSerialize["minimum"] = o.Minimum
	toSerialize["enum"] = o.Enum

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ParameterSpec) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name        *string                   `json:"name"`
		Description *string                   `json:"description"`
		Type        *string                   `json:"type"`
		Default     *map[string]interface{}   `json:"default"`
		NeedRestart *bool                     `json:"needRestart"`
		Immutable   *bool                     `json:"immutable"`
		Maximum     *float64                  `json:"maximum"`
		Minimum     *float64                  `json:"minimum"`
		Enum        *[]map[string]interface{} `json:"enum"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Default == nil {
		return fmt.Errorf("required field default missing")
	}
	if all.NeedRestart == nil {
		return fmt.Errorf("required field needRestart missing")
	}
	if all.Immutable == nil {
		return fmt.Errorf("required field immutable missing")
	}
	if all.Maximum == nil {
		return fmt.Errorf("required field maximum missing")
	}
	if all.Minimum == nil {
		return fmt.Errorf("required field minimum missing")
	}
	if all.Enum == nil {
		return fmt.Errorf("required field enum missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "description", "type", "default", "needRestart", "immutable", "maximum", "minimum", "enum"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Description = *all.Description
	o.Type = *all.Type
	o.Default = *all.Default
	o.NeedRestart = *all.NeedRestart
	o.Immutable = *all.Immutable
	o.Maximum = *all.Maximum
	o.Minimum = *all.Minimum
	o.Enum = *all.Enum

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
