// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DmFile struct {
	// the id of this file
	Id *string `json:"id,omitempty"`
	// the name of this file
	Name string `json:"name"`
	// whether this file auto extend
	AutoExtend *bool `json:"autoExtend,omitempty"`
	// the size already allocated of this file, unit MB
	AllocatedSizeMb *string `json:"allocatedSizeMB,omitempty"`
	// the current used size of this file, unit MB
	UsedSizeMb *string `json:"usedSizeMB,omitempty"`
	// the current used ratio of this file
	UsedRatio *string `json:"usedRatio,omitempty"`
	// the extend step of this file, unit MB
	ExtendStepMb *string `json:"extendStepMB,omitempty"`
	// the max size of this file, unit GB
	MaxSizeGb *string `json:"maxSizeGB,omitempty"`
	// the used ratio in the max size of this file
	UsedRatioInMax *string `json:"usedRatioInMax,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmFile instantiates a new DmFile object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmFile(name string) *DmFile {
	this := DmFile{}
	this.Name = name
	return &this
}

// NewDmFileWithDefaults instantiates a new DmFile object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmFileWithDefaults() *DmFile {
	this := DmFile{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DmFile) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmFile) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DmFile) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DmFile) SetId(v string) {
	o.Id = &v
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

// GetAutoExtend returns the AutoExtend field value if set, zero value otherwise.
func (o *DmFile) GetAutoExtend() bool {
	if o == nil || o.AutoExtend == nil {
		var ret bool
		return ret
	}
	return *o.AutoExtend
}

// GetAutoExtendOk returns a tuple with the AutoExtend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmFile) GetAutoExtendOk() (*bool, bool) {
	if o == nil || o.AutoExtend == nil {
		return nil, false
	}
	return o.AutoExtend, true
}

// HasAutoExtend returns a boolean if a field has been set.
func (o *DmFile) HasAutoExtend() bool {
	return o != nil && o.AutoExtend != nil
}

// SetAutoExtend gets a reference to the given bool and assigns it to the AutoExtend field.
func (o *DmFile) SetAutoExtend(v bool) {
	o.AutoExtend = &v
}

// GetAllocatedSizeMb returns the AllocatedSizeMb field value if set, zero value otherwise.
func (o *DmFile) GetAllocatedSizeMb() string {
	if o == nil || o.AllocatedSizeMb == nil {
		var ret string
		return ret
	}
	return *o.AllocatedSizeMb
}

// GetAllocatedSizeMbOk returns a tuple with the AllocatedSizeMb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmFile) GetAllocatedSizeMbOk() (*string, bool) {
	if o == nil || o.AllocatedSizeMb == nil {
		return nil, false
	}
	return o.AllocatedSizeMb, true
}

// HasAllocatedSizeMb returns a boolean if a field has been set.
func (o *DmFile) HasAllocatedSizeMb() bool {
	return o != nil && o.AllocatedSizeMb != nil
}

// SetAllocatedSizeMb gets a reference to the given string and assigns it to the AllocatedSizeMb field.
func (o *DmFile) SetAllocatedSizeMb(v string) {
	o.AllocatedSizeMb = &v
}

// GetUsedSizeMb returns the UsedSizeMb field value if set, zero value otherwise.
func (o *DmFile) GetUsedSizeMb() string {
	if o == nil || o.UsedSizeMb == nil {
		var ret string
		return ret
	}
	return *o.UsedSizeMb
}

// GetUsedSizeMbOk returns a tuple with the UsedSizeMb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmFile) GetUsedSizeMbOk() (*string, bool) {
	if o == nil || o.UsedSizeMb == nil {
		return nil, false
	}
	return o.UsedSizeMb, true
}

// HasUsedSizeMb returns a boolean if a field has been set.
func (o *DmFile) HasUsedSizeMb() bool {
	return o != nil && o.UsedSizeMb != nil
}

// SetUsedSizeMb gets a reference to the given string and assigns it to the UsedSizeMb field.
func (o *DmFile) SetUsedSizeMb(v string) {
	o.UsedSizeMb = &v
}

// GetUsedRatio returns the UsedRatio field value if set, zero value otherwise.
func (o *DmFile) GetUsedRatio() string {
	if o == nil || o.UsedRatio == nil {
		var ret string
		return ret
	}
	return *o.UsedRatio
}

// GetUsedRatioOk returns a tuple with the UsedRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmFile) GetUsedRatioOk() (*string, bool) {
	if o == nil || o.UsedRatio == nil {
		return nil, false
	}
	return o.UsedRatio, true
}

// HasUsedRatio returns a boolean if a field has been set.
func (o *DmFile) HasUsedRatio() bool {
	return o != nil && o.UsedRatio != nil
}

// SetUsedRatio gets a reference to the given string and assigns it to the UsedRatio field.
func (o *DmFile) SetUsedRatio(v string) {
	o.UsedRatio = &v
}

// GetExtendStepMb returns the ExtendStepMb field value if set, zero value otherwise.
func (o *DmFile) GetExtendStepMb() string {
	if o == nil || o.ExtendStepMb == nil {
		var ret string
		return ret
	}
	return *o.ExtendStepMb
}

// GetExtendStepMbOk returns a tuple with the ExtendStepMb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmFile) GetExtendStepMbOk() (*string, bool) {
	if o == nil || o.ExtendStepMb == nil {
		return nil, false
	}
	return o.ExtendStepMb, true
}

// HasExtendStepMb returns a boolean if a field has been set.
func (o *DmFile) HasExtendStepMb() bool {
	return o != nil && o.ExtendStepMb != nil
}

// SetExtendStepMb gets a reference to the given string and assigns it to the ExtendStepMb field.
func (o *DmFile) SetExtendStepMb(v string) {
	o.ExtendStepMb = &v
}

// GetMaxSizeGb returns the MaxSizeGb field value if set, zero value otherwise.
func (o *DmFile) GetMaxSizeGb() string {
	if o == nil || o.MaxSizeGb == nil {
		var ret string
		return ret
	}
	return *o.MaxSizeGb
}

// GetMaxSizeGbOk returns a tuple with the MaxSizeGb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmFile) GetMaxSizeGbOk() (*string, bool) {
	if o == nil || o.MaxSizeGb == nil {
		return nil, false
	}
	return o.MaxSizeGb, true
}

// HasMaxSizeGb returns a boolean if a field has been set.
func (o *DmFile) HasMaxSizeGb() bool {
	return o != nil && o.MaxSizeGb != nil
}

// SetMaxSizeGb gets a reference to the given string and assigns it to the MaxSizeGb field.
func (o *DmFile) SetMaxSizeGb(v string) {
	o.MaxSizeGb = &v
}

// GetUsedRatioInMax returns the UsedRatioInMax field value if set, zero value otherwise.
func (o *DmFile) GetUsedRatioInMax() string {
	if o == nil || o.UsedRatioInMax == nil {
		var ret string
		return ret
	}
	return *o.UsedRatioInMax
}

// GetUsedRatioInMaxOk returns a tuple with the UsedRatioInMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmFile) GetUsedRatioInMaxOk() (*string, bool) {
	if o == nil || o.UsedRatioInMax == nil {
		return nil, false
	}
	return o.UsedRatioInMax, true
}

// HasUsedRatioInMax returns a boolean if a field has been set.
func (o *DmFile) HasUsedRatioInMax() bool {
	return o != nil && o.UsedRatioInMax != nil
}

// SetUsedRatioInMax gets a reference to the given string and assigns it to the UsedRatioInMax field.
func (o *DmFile) SetUsedRatioInMax(v string) {
	o.UsedRatioInMax = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DmFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if o.AutoExtend != nil {
		toSerialize["autoExtend"] = o.AutoExtend
	}
	if o.AllocatedSizeMb != nil {
		toSerialize["allocatedSizeMB"] = o.AllocatedSizeMb
	}
	if o.UsedSizeMb != nil {
		toSerialize["usedSizeMB"] = o.UsedSizeMb
	}
	if o.UsedRatio != nil {
		toSerialize["usedRatio"] = o.UsedRatio
	}
	if o.ExtendStepMb != nil {
		toSerialize["extendStepMB"] = o.ExtendStepMb
	}
	if o.MaxSizeGb != nil {
		toSerialize["maxSizeGB"] = o.MaxSizeGb
	}
	if o.UsedRatioInMax != nil {
		toSerialize["usedRatioInMax"] = o.UsedRatioInMax
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmFile) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id              *string `json:"id,omitempty"`
		Name            *string `json:"name"`
		AutoExtend      *bool   `json:"autoExtend,omitempty"`
		AllocatedSizeMb *string `json:"allocatedSizeMB,omitempty"`
		UsedSizeMb      *string `json:"usedSizeMB,omitempty"`
		UsedRatio       *string `json:"usedRatio,omitempty"`
		ExtendStepMb    *string `json:"extendStepMB,omitempty"`
		MaxSizeGb       *string `json:"maxSizeGB,omitempty"`
		UsedRatioInMax  *string `json:"usedRatioInMax,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "name", "autoExtend", "allocatedSizeMB", "usedSizeMB", "usedRatio", "extendStepMB", "maxSizeGB", "usedRatioInMax"})
	} else {
		return err
	}
	o.Id = all.Id
	o.Name = *all.Name
	o.AutoExtend = all.AutoExtend
	o.AllocatedSizeMb = all.AllocatedSizeMb
	o.UsedSizeMb = all.UsedSizeMb
	o.UsedRatio = all.UsedRatio
	o.ExtendStepMb = all.ExtendStepMb
	o.MaxSizeGb = all.MaxSizeGb
	o.UsedRatioInMax = all.UsedRatioInMax

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
