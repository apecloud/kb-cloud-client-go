// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ModuleDefinition struct {
	Name   *string                  `json:"name,omitempty"`
	Title  *InternationalDesc       `json:"title,omitempty"`
	Values []ModuleDefinitionValues `json:"values,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewModuleDefinition instantiates a new ModuleDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewModuleDefinition() *ModuleDefinition {
	this := ModuleDefinition{}
	return &this
}

// NewModuleDefinitionWithDefaults instantiates a new ModuleDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewModuleDefinitionWithDefaults() *ModuleDefinition {
	this := ModuleDefinition{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModuleDefinition) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleDefinition) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModuleDefinition) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModuleDefinition) SetName(v string) {
	o.Name = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ModuleDefinition) GetTitle() InternationalDesc {
	if o == nil || o.Title == nil {
		var ret InternationalDesc
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleDefinition) GetTitleOk() (*InternationalDesc, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ModuleDefinition) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given InternationalDesc and assigns it to the Title field.
func (o *ModuleDefinition) SetTitle(v InternationalDesc) {
	o.Title = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *ModuleDefinition) GetValues() []ModuleDefinitionValues {
	if o == nil || o.Values == nil {
		var ret []ModuleDefinitionValues
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleDefinition) GetValuesOk() (*[]ModuleDefinitionValues, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return &o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *ModuleDefinition) HasValues() bool {
	return o != nil && o.Values != nil
}

// SetValues gets a reference to the given []ModuleDefinitionValues and assigns it to the Values field.
func (o *ModuleDefinition) SetValues(v []ModuleDefinitionValues) {
	o.Values = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ModuleDefinition) MarshalJSON() ([]byte, error) {
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
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ModuleDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name   *string                  `json:"name,omitempty"`
		Title  *InternationalDesc       `json:"title,omitempty"`
		Values []ModuleDefinitionValues `json:"values,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "title", "values"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = all.Name
	if all.Title != nil && all.Title.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Title = all.Title
	o.Values = all.Values

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
