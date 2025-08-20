// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DmFile struct {
	// the name of this file
	Name string `json:"name"`
	// whether this file auto extend
	AutoExtend bool `json:"autoExtend"`
	// the size of this file
	Size *string `json:"size,omitempty"`
	// the extend step of this file
	ExtendStep *string `json:"extendStep,omitempty"`
	// the max size of this file
	MaxSize *string `json:"maxSize,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmFile instantiates a new DmFile object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmFile(name string, autoExtend bool) *DmFile {
	this := DmFile{}
	this.Name = name
	this.AutoExtend = autoExtend
	return &this
}

// NewDmFileWithDefaults instantiates a new DmFile object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmFileWithDefaults() *DmFile {
	this := DmFile{}
	return &this
}

// GetName returns the Name field value.
func (o *DmFile) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DmFile) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *DmFile) SetName(v string) {
	o.Name = v
}

// GetAutoExtend returns the AutoExtend field value.
func (o *DmFile) GetAutoExtend() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.AutoExtend
}

// GetAutoExtendOk returns a tuple with the AutoExtend field value
// and a boolean to check if the value has been set.
func (o *DmFile) GetAutoExtendOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoExtend, true
}

// SetAutoExtend sets field value.
func (o *DmFile) SetAutoExtend(v bool) {
	o.AutoExtend = v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *DmFile) GetSize() string {
	if o == nil || o.Size == nil {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmFile) GetSizeOk() (*string, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *DmFile) HasSize() bool {
	return o != nil && o.Size != nil
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *DmFile) SetSize(v string) {
	o.Size = &v
}

// GetExtendStep returns the ExtendStep field value if set, zero value otherwise.
func (o *DmFile) GetExtendStep() string {
	if o == nil || o.ExtendStep == nil {
		var ret string
		return ret
	}
	return *o.ExtendStep
}

// GetExtendStepOk returns a tuple with the ExtendStep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmFile) GetExtendStepOk() (*string, bool) {
	if o == nil || o.ExtendStep == nil {
		return nil, false
	}
	return o.ExtendStep, true
}

// HasExtendStep returns a boolean if a field has been set.
func (o *DmFile) HasExtendStep() bool {
	return o != nil && o.ExtendStep != nil
}

// SetExtendStep gets a reference to the given string and assigns it to the ExtendStep field.
func (o *DmFile) SetExtendStep(v string) {
	o.ExtendStep = &v
}

// GetMaxSize returns the MaxSize field value if set, zero value otherwise.
func (o *DmFile) GetMaxSize() string {
	if o == nil || o.MaxSize == nil {
		var ret string
		return ret
	}
	return *o.MaxSize
}

// GetMaxSizeOk returns a tuple with the MaxSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmFile) GetMaxSizeOk() (*string, bool) {
	if o == nil || o.MaxSize == nil {
		return nil, false
	}
	return o.MaxSize, true
}

// HasMaxSize returns a boolean if a field has been set.
func (o *DmFile) HasMaxSize() bool {
	return o != nil && o.MaxSize != nil
}

// SetMaxSize gets a reference to the given string and assigns it to the MaxSize field.
func (o *DmFile) SetMaxSize(v string) {
	o.MaxSize = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DmFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["autoExtend"] = o.AutoExtend
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.ExtendStep != nil {
		toSerialize["extendStep"] = o.ExtendStep
	}
	if o.MaxSize != nil {
		toSerialize["maxSize"] = o.MaxSize
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmFile) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name       *string `json:"name"`
		AutoExtend *bool   `json:"autoExtend"`
		Size       *string `json:"size,omitempty"`
		ExtendStep *string `json:"extendStep,omitempty"`
		MaxSize    *string `json:"maxSize,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.AutoExtend == nil {
		return fmt.Errorf("required field autoExtend missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "autoExtend", "size", "extendStep", "maxSize"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.AutoExtend = *all.AutoExtend
	o.Size = all.Size
	o.ExtendStep = all.ExtendStep
	o.MaxSize = all.MaxSize

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
