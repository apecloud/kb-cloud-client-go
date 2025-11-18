// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type StandardDefinition struct {
	Name        *string             `json:"name,omitempty"`
	Title       *InternationalDesc  `json:"title,omitempty"`
	Description *InternationalDesc  `json:"description,omitempty"`
	IsDefault   common.NullableBool `json:"isDefault,omitempty"`
	Cpu         *StandardResource   `json:"cpu,omitempty"`
	Memory      *StandardResource   `json:"memory,omitempty"`
	// module parameter template, the key is the module name, the value is the parameter template
	ParameterTemplate map[string][]DataChannelParameter `json:"parameterTemplate,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStandardDefinition instantiates a new StandardDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStandardDefinition() *StandardDefinition {
	this := StandardDefinition{}
	return &this
}

// NewStandardDefinitionWithDefaults instantiates a new StandardDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStandardDefinitionWithDefaults() *StandardDefinition {
	this := StandardDefinition{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StandardDefinition) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardDefinition) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StandardDefinition) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StandardDefinition) SetName(v string) {
	o.Name = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *StandardDefinition) GetTitle() InternationalDesc {
	if o == nil || o.Title == nil {
		var ret InternationalDesc
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardDefinition) GetTitleOk() (*InternationalDesc, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *StandardDefinition) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given InternationalDesc and assigns it to the Title field.
func (o *StandardDefinition) SetTitle(v InternationalDesc) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *StandardDefinition) GetDescription() InternationalDesc {
	if o == nil || o.Description == nil {
		var ret InternationalDesc
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardDefinition) GetDescriptionOk() (*InternationalDesc, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *StandardDefinition) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given InternationalDesc and assigns it to the Description field.
func (o *StandardDefinition) SetDescription(v InternationalDesc) {
	o.Description = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandardDefinition) GetIsDefault() bool {
	if o == nil || o.IsDefault.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsDefault.Get()
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *StandardDefinition) GetIsDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsDefault.Get(), o.IsDefault.IsSet()
}

// HasIsDefault returns a boolean if a field has been set.
func (o *StandardDefinition) HasIsDefault() bool {
	return o != nil && o.IsDefault.IsSet()
}

// SetIsDefault gets a reference to the given common.NullableBool and assigns it to the IsDefault field.
func (o *StandardDefinition) SetIsDefault(v bool) {
	o.IsDefault.Set(&v)
}

// SetIsDefaultNil sets the value for IsDefault to be an explicit nil.
func (o *StandardDefinition) SetIsDefaultNil() {
	o.IsDefault.Set(nil)
}

// UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil.
func (o *StandardDefinition) UnsetIsDefault() {
	o.IsDefault.Unset()
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *StandardDefinition) GetCpu() StandardResource {
	if o == nil || o.Cpu == nil {
		var ret StandardResource
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardDefinition) GetCpuOk() (*StandardResource, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *StandardDefinition) HasCpu() bool {
	return o != nil && o.Cpu != nil
}

// SetCpu gets a reference to the given StandardResource and assigns it to the Cpu field.
func (o *StandardDefinition) SetCpu(v StandardResource) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *StandardDefinition) GetMemory() StandardResource {
	if o == nil || o.Memory == nil {
		var ret StandardResource
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardDefinition) GetMemoryOk() (*StandardResource, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *StandardDefinition) HasMemory() bool {
	return o != nil && o.Memory != nil
}

// SetMemory gets a reference to the given StandardResource and assigns it to the Memory field.
func (o *StandardDefinition) SetMemory(v StandardResource) {
	o.Memory = &v
}

// GetParameterTemplate returns the ParameterTemplate field value if set, zero value otherwise.
func (o *StandardDefinition) GetParameterTemplate() map[string][]DataChannelParameter {
	if o == nil || o.ParameterTemplate == nil {
		var ret map[string][]DataChannelParameter
		return ret
	}
	return o.ParameterTemplate
}

// GetParameterTemplateOk returns a tuple with the ParameterTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardDefinition) GetParameterTemplateOk() (*map[string][]DataChannelParameter, bool) {
	if o == nil || o.ParameterTemplate == nil {
		return nil, false
	}
	return &o.ParameterTemplate, true
}

// HasParameterTemplate returns a boolean if a field has been set.
func (o *StandardDefinition) HasParameterTemplate() bool {
	return o != nil && o.ParameterTemplate != nil
}

// SetParameterTemplate gets a reference to the given map[string][]DataChannelParameter and assigns it to the ParameterTemplate field.
func (o *StandardDefinition) SetParameterTemplate(v map[string][]DataChannelParameter) {
	o.ParameterTemplate = v
}

// MarshalJSON serializes the struct using spec logic.
func (o StandardDefinition) MarshalJSON() ([]byte, error) {
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
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.IsDefault.IsSet() {
		toSerialize["isDefault"] = o.IsDefault.Get()
	}
	if o.Cpu != nil {
		toSerialize["cpu"] = o.Cpu
	}
	if o.Memory != nil {
		toSerialize["memory"] = o.Memory
	}
	if o.ParameterTemplate != nil {
		toSerialize["parameterTemplate"] = o.ParameterTemplate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StandardDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name              *string                           `json:"name,omitempty"`
		Title             *InternationalDesc                `json:"title,omitempty"`
		Description       *InternationalDesc                `json:"description,omitempty"`
		IsDefault         common.NullableBool               `json:"isDefault,omitempty"`
		Cpu               *StandardResource                 `json:"cpu,omitempty"`
		Memory            *StandardResource                 `json:"memory,omitempty"`
		ParameterTemplate map[string][]DataChannelParameter `json:"parameterTemplate,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "title", "description", "isDefault", "cpu", "memory", "parameterTemplate"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = all.Name
	if all.Title != nil && all.Title.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Title = all.Title
	if all.Description != nil && all.Description.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Description = all.Description
	o.IsDefault = all.IsDefault
	if all.Cpu != nil && all.Cpu.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Cpu = all.Cpu
	if all.Memory != nil && all.Memory.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Memory = all.Memory
	o.ParameterTemplate = all.ParameterTemplate

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
