// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
)

// OpsVolumeExpand OpsVolumeExpand is the payload to expand volume for a KubeBlocks cluster
type OpsVolumeExpand struct {
	Component *string                      `json:"component,omitempty"`
	Volumes   []OpsVolumeExpandVolumesItem `json:"volumes"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOpsVolumeExpand instantiates a new OpsVolumeExpand object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpsVolumeExpand(volumes []OpsVolumeExpandVolumesItem) *OpsVolumeExpand {
	this := OpsVolumeExpand{}
	this.Volumes = volumes
	return &this
}

// NewOpsVolumeExpandWithDefaults instantiates a new OpsVolumeExpand object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpsVolumeExpandWithDefaults() *OpsVolumeExpand {
	this := OpsVolumeExpand{}
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *OpsVolumeExpand) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsVolumeExpand) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *OpsVolumeExpand) HasComponent() bool {
	return o != nil && o.Component != nil
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *OpsVolumeExpand) SetComponent(v string) {
	o.Component = &v
}

// GetVolumes returns the Volumes field value.
func (o *OpsVolumeExpand) GetVolumes() []OpsVolumeExpandVolumesItem {
	if o == nil {
		var ret []OpsVolumeExpandVolumesItem
		return ret
	}
	return o.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value
// and a boolean to check if the value has been set.
func (o *OpsVolumeExpand) GetVolumesOk() (*[]OpsVolumeExpandVolumesItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Volumes, true
}

// SetVolumes sets field value.
func (o *OpsVolumeExpand) SetVolumes(v []OpsVolumeExpandVolumesItem) {
	o.Volumes = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpsVolumeExpand) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	toSerialize["volumes"] = o.Volumes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpsVolumeExpand) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Component *string                       `json:"component,omitempty"`
		Volumes   *[]OpsVolumeExpandVolumesItem `json:"volumes"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Volumes == nil {
		return fmt.Errorf("required field volumes missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"component", "volumes"})
	} else {
		return err
	}
	o.Component = all.Component
	o.Volumes = *all.Volumes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
