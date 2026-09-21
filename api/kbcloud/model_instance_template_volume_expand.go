// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// InstanceTemplateVolumeExpand Volume expansion assignment for one instance template.
type InstanceTemplateVolumeExpand struct {
	// Instance template name declared on the engine option.
	Name    string                                    `json:"name"`
	Volumes []InstanceTemplateVolumeExpandVolumesItem `json:"volumes"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInstanceTemplateVolumeExpand instantiates a new InstanceTemplateVolumeExpand object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInstanceTemplateVolumeExpand(name string, volumes []InstanceTemplateVolumeExpandVolumesItem) *InstanceTemplateVolumeExpand {
	this := InstanceTemplateVolumeExpand{}
	this.Name = name
	this.Volumes = volumes
	return &this
}

// NewInstanceTemplateVolumeExpandWithDefaults instantiates a new InstanceTemplateVolumeExpand object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInstanceTemplateVolumeExpandWithDefaults() *InstanceTemplateVolumeExpand {
	this := InstanceTemplateVolumeExpand{}
	return &this
}

// GetName returns the Name field value.
func (o *InstanceTemplateVolumeExpand) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InstanceTemplateVolumeExpand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *InstanceTemplateVolumeExpand) SetName(v string) {
	o.Name = v
}

// GetVolumes returns the Volumes field value.
func (o *InstanceTemplateVolumeExpand) GetVolumes() []InstanceTemplateVolumeExpandVolumesItem {
	if o == nil {
		var ret []InstanceTemplateVolumeExpandVolumesItem
		return ret
	}
	return o.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value
// and a boolean to check if the value has been set.
func (o *InstanceTemplateVolumeExpand) GetVolumesOk() (*[]InstanceTemplateVolumeExpandVolumesItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Volumes, true
}

// SetVolumes sets field value.
func (o *InstanceTemplateVolumeExpand) SetVolumes(v []InstanceTemplateVolumeExpandVolumesItem) {
	o.Volumes = v
}

// MarshalJSON serializes the struct using spec logic.
func (o InstanceTemplateVolumeExpand) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["volumes"] = o.Volumes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InstanceTemplateVolumeExpand) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name    *string                                    `json:"name"`
		Volumes *[]InstanceTemplateVolumeExpandVolumesItem `json:"volumes"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Volumes == nil {
		return fmt.Errorf("required field volumes missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "volumes"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Volumes = *all.Volumes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
