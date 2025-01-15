// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// InstanceDiskUsageItem disk usage with volume name.
type InstanceDiskUsageItem struct {
	// volume name
	VolumeName string `json:"volumeName"`
	// disk usage with unit Gi
	DiskUsage string `json:"diskUsage"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInstanceDiskUsageItem instantiates a new InstanceDiskUsageItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInstanceDiskUsageItem(volumeName string, diskUsage string) *InstanceDiskUsageItem {
	this := InstanceDiskUsageItem{}
	this.VolumeName = volumeName
	this.DiskUsage = diskUsage
	return &this
}

// NewInstanceDiskUsageItemWithDefaults instantiates a new InstanceDiskUsageItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInstanceDiskUsageItemWithDefaults() *InstanceDiskUsageItem {
	this := InstanceDiskUsageItem{}
	return &this
}

// GetVolumeName returns the VolumeName field value.
func (o *InstanceDiskUsageItem) GetVolumeName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.VolumeName
}

// GetVolumeNameOk returns a tuple with the VolumeName field value
// and a boolean to check if the value has been set.
func (o *InstanceDiskUsageItem) GetVolumeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VolumeName, true
}

// SetVolumeName sets field value.
func (o *InstanceDiskUsageItem) SetVolumeName(v string) {
	o.VolumeName = v
}

// GetDiskUsage returns the DiskUsage field value.
func (o *InstanceDiskUsageItem) GetDiskUsage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DiskUsage
}

// GetDiskUsageOk returns a tuple with the DiskUsage field value
// and a boolean to check if the value has been set.
func (o *InstanceDiskUsageItem) GetDiskUsageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiskUsage, true
}

// SetDiskUsage sets field value.
func (o *InstanceDiskUsageItem) SetDiskUsage(v string) {
	o.DiskUsage = v
}

// MarshalJSON serializes the struct using spec logic.
func (o InstanceDiskUsageItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["volumeName"] = o.VolumeName
	toSerialize["diskUsage"] = o.DiskUsage

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InstanceDiskUsageItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		VolumeName *string `json:"volumeName"`
		DiskUsage  *string `json:"diskUsage"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.VolumeName == nil {
		return fmt.Errorf("required field volumeName missing")
	}
	if all.DiskUsage == nil {
		return fmt.Errorf("required field diskUsage missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"volumeName", "diskUsage"})
	} else {
		return err
	}
	o.VolumeName = *all.VolumeName
	o.DiskUsage = *all.DiskUsage

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
