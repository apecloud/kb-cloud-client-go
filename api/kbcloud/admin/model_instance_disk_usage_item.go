// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// InstanceDiskUsageItem pvc disk usage with pvc name.
type InstanceDiskUsageItem struct {
	// pvc name
	PvcName string `json:"pvcName"`
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
func NewInstanceDiskUsageItem(pvcName string, diskUsage string) *InstanceDiskUsageItem {
	this := InstanceDiskUsageItem{}
	this.PvcName = pvcName
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

// GetPvcName returns the PvcName field value.
func (o *InstanceDiskUsageItem) GetPvcName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PvcName
}

// GetPvcNameOk returns a tuple with the PvcName field value
// and a boolean to check if the value has been set.
func (o *InstanceDiskUsageItem) GetPvcNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PvcName, true
}

// SetPvcName sets field value.
func (o *InstanceDiskUsageItem) SetPvcName(v string) {
	o.PvcName = v
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
	toSerialize["pvcName"] = o.PvcName
	toSerialize["diskUsage"] = o.DiskUsage

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InstanceDiskUsageItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		PvcName   *string `json:"pvcName"`
		DiskUsage *string `json:"diskUsage"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.PvcName == nil {
		return fmt.Errorf("required field pvcName missing")
	}
	if all.DiskUsage == nil {
		return fmt.Errorf("required field diskUsage missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"pvcName", "diskUsage"})
	} else {
		return err
	}
	o.PvcName = *all.PvcName
	o.DiskUsage = *all.DiskUsage

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
