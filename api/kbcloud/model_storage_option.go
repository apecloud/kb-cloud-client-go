// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type StorageOption struct {
	Title   LocalizedDescription `json:"title"`
	Name    string               `json:"name"`
	Min     int32                `json:"min"`
	Max     int32                `json:"max"`
	Default int32                `json:"default"`
	Step    int32                `json:"step"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStorageOption instantiates a new StorageOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStorageOption(title LocalizedDescription, name string, min int32, max int32, defaultVar int32, step int32) *StorageOption {
	this := StorageOption{}
	this.Title = title
	this.Name = name
	this.Min = min
	this.Max = max
	this.Default = defaultVar
	this.Step = step
	return &this
}

// NewStorageOptionWithDefaults instantiates a new StorageOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStorageOptionWithDefaults() *StorageOption {
	this := StorageOption{}
	return &this
}

// GetTitle returns the Title field value.
func (o *StorageOption) GetTitle() LocalizedDescription {
	if o == nil {
		var ret LocalizedDescription
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *StorageOption) GetTitleOk() (*LocalizedDescription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *StorageOption) SetTitle(v LocalizedDescription) {
	o.Title = v
}

// GetName returns the Name field value.
func (o *StorageOption) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StorageOption) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *StorageOption) SetName(v string) {
	o.Name = v
}

// GetMin returns the Min field value.
func (o *StorageOption) GetMin() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Min
}

// GetMinOk returns a tuple with the Min field value
// and a boolean to check if the value has been set.
func (o *StorageOption) GetMinOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Min, true
}

// SetMin sets field value.
func (o *StorageOption) SetMin(v int32) {
	o.Min = v
}

// GetMax returns the Max field value.
func (o *StorageOption) GetMax() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Max
}

// GetMaxOk returns a tuple with the Max field value
// and a boolean to check if the value has been set.
func (o *StorageOption) GetMaxOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Max, true
}

// SetMax sets field value.
func (o *StorageOption) SetMax(v int32) {
	o.Max = v
}

// GetDefault returns the Default field value.
func (o *StorageOption) GetDefault() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value
// and a boolean to check if the value has been set.
func (o *StorageOption) GetDefaultOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Default, true
}

// SetDefault sets field value.
func (o *StorageOption) SetDefault(v int32) {
	o.Default = v
}

// GetStep returns the Step field value.
func (o *StorageOption) GetStep() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Step
}

// GetStepOk returns a tuple with the Step field value
// and a boolean to check if the value has been set.
func (o *StorageOption) GetStepOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Step, true
}

// SetStep sets field value.
func (o *StorageOption) SetStep(v int32) {
	o.Step = v
}

// MarshalJSON serializes the struct using spec logic.
func (o StorageOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["title"] = o.Title
	toSerialize["name"] = o.Name
	toSerialize["min"] = o.Min
	toSerialize["max"] = o.Max
	toSerialize["default"] = o.Default
	toSerialize["step"] = o.Step

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StorageOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Title   *LocalizedDescription `json:"title"`
		Name    *string               `json:"name"`
		Min     *int32                `json:"min"`
		Max     *int32                `json:"max"`
		Default *int32                `json:"default"`
		Step    *int32                `json:"step"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Min == nil {
		return fmt.Errorf("required field min missing")
	}
	if all.Max == nil {
		return fmt.Errorf("required field max missing")
	}
	if all.Default == nil {
		return fmt.Errorf("required field default missing")
	}
	if all.Step == nil {
		return fmt.Errorf("required field step missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"title", "name", "min", "max", "default", "step"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Title.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Title = *all.Title
	o.Name = *all.Name
	o.Min = *all.Min
	o.Max = *all.Max
	o.Default = *all.Default
	o.Step = *all.Step

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
