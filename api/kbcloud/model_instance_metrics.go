// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// InstanceMetrics instance metrics
type InstanceMetrics struct {
	// the name of the instance
	InstanceName string `json:"instanceName"`
	// cpu with uint cores.
	CpuUsage common.NullableString `json:"cpuUsage,omitempty"`
	// memory with uint Gi.
	MemoryUsage common.NullableString `json:"memoryUsage,omitempty"`
	// disk usage items
	DiskUsageItems []InstanceDiskUsageItem `json:"diskUsageItems,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInstanceMetrics instantiates a new InstanceMetrics object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInstanceMetrics(instanceName string) *InstanceMetrics {
	this := InstanceMetrics{}
	this.InstanceName = instanceName
	return &this
}

// NewInstanceMetricsWithDefaults instantiates a new InstanceMetrics object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInstanceMetricsWithDefaults() *InstanceMetrics {
	this := InstanceMetrics{}
	return &this
}

// GetInstanceName returns the InstanceName field value.
func (o *InstanceMetrics) GetInstanceName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.InstanceName
}

// GetInstanceNameOk returns a tuple with the InstanceName field value
// and a boolean to check if the value has been set.
func (o *InstanceMetrics) GetInstanceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceName, true
}

// SetInstanceName sets field value.
func (o *InstanceMetrics) SetInstanceName(v string) {
	o.InstanceName = v
}

// GetCpuUsage returns the CpuUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceMetrics) GetCpuUsage() string {
	if o == nil || o.CpuUsage.Get() == nil {
		var ret string
		return ret
	}
	return *o.CpuUsage.Get()
}

// GetCpuUsageOk returns a tuple with the CpuUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *InstanceMetrics) GetCpuUsageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CpuUsage.Get(), o.CpuUsage.IsSet()
}

// HasCpuUsage returns a boolean if a field has been set.
func (o *InstanceMetrics) HasCpuUsage() bool {
	return o != nil && o.CpuUsage.IsSet()
}

// SetCpuUsage gets a reference to the given common.NullableString and assigns it to the CpuUsage field.
func (o *InstanceMetrics) SetCpuUsage(v string) {
	o.CpuUsage.Set(&v)
}

// SetCpuUsageNil sets the value for CpuUsage to be an explicit nil.
func (o *InstanceMetrics) SetCpuUsageNil() {
	o.CpuUsage.Set(nil)
}

// UnsetCpuUsage ensures that no value is present for CpuUsage, not even an explicit nil.
func (o *InstanceMetrics) UnsetCpuUsage() {
	o.CpuUsage.Unset()
}

// GetMemoryUsage returns the MemoryUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceMetrics) GetMemoryUsage() string {
	if o == nil || o.MemoryUsage.Get() == nil {
		var ret string
		return ret
	}
	return *o.MemoryUsage.Get()
}

// GetMemoryUsageOk returns a tuple with the MemoryUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *InstanceMetrics) GetMemoryUsageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MemoryUsage.Get(), o.MemoryUsage.IsSet()
}

// HasMemoryUsage returns a boolean if a field has been set.
func (o *InstanceMetrics) HasMemoryUsage() bool {
	return o != nil && o.MemoryUsage.IsSet()
}

// SetMemoryUsage gets a reference to the given common.NullableString and assigns it to the MemoryUsage field.
func (o *InstanceMetrics) SetMemoryUsage(v string) {
	o.MemoryUsage.Set(&v)
}

// SetMemoryUsageNil sets the value for MemoryUsage to be an explicit nil.
func (o *InstanceMetrics) SetMemoryUsageNil() {
	o.MemoryUsage.Set(nil)
}

// UnsetMemoryUsage ensures that no value is present for MemoryUsage, not even an explicit nil.
func (o *InstanceMetrics) UnsetMemoryUsage() {
	o.MemoryUsage.Unset()
}

// GetDiskUsageItems returns the DiskUsageItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceMetrics) GetDiskUsageItems() []InstanceDiskUsageItem {
	if o == nil {
		var ret []InstanceDiskUsageItem
		return ret
	}
	return o.DiskUsageItems
}

// GetDiskUsageItemsOk returns a tuple with the DiskUsageItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *InstanceMetrics) GetDiskUsageItemsOk() (*[]InstanceDiskUsageItem, bool) {
	if o == nil || o.DiskUsageItems == nil {
		return nil, false
	}
	return &o.DiskUsageItems, true
}

// HasDiskUsageItems returns a boolean if a field has been set.
func (o *InstanceMetrics) HasDiskUsageItems() bool {
	return o != nil && o.DiskUsageItems != nil
}

// SetDiskUsageItems gets a reference to the given []InstanceDiskUsageItem and assigns it to the DiskUsageItems field.
func (o *InstanceMetrics) SetDiskUsageItems(v []InstanceDiskUsageItem) {
	o.DiskUsageItems = v
}

// MarshalJSON serializes the struct using spec logic.
func (o InstanceMetrics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["instanceName"] = o.InstanceName
	if o.CpuUsage.IsSet() {
		toSerialize["cpuUsage"] = o.CpuUsage.Get()
	}
	if o.MemoryUsage.IsSet() {
		toSerialize["memoryUsage"] = o.MemoryUsage.Get()
	}
	if o.DiskUsageItems != nil {
		toSerialize["diskUsageItems"] = o.DiskUsageItems
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InstanceMetrics) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		InstanceName   *string                 `json:"instanceName"`
		CpuUsage       common.NullableString   `json:"cpuUsage,omitempty"`
		MemoryUsage    common.NullableString   `json:"memoryUsage,omitempty"`
		DiskUsageItems []InstanceDiskUsageItem `json:"diskUsageItems,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.InstanceName == nil {
		return fmt.Errorf("required field instanceName missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"instanceName", "cpuUsage", "memoryUsage", "diskUsageItems"})
	} else {
		return err
	}
	o.InstanceName = *all.InstanceName
	o.CpuUsage = all.CpuUsage
	o.MemoryUsage = all.MemoryUsage
	o.DiskUsageItems = all.DiskUsageItems

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
