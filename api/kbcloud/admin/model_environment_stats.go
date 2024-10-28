// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EnvironmentStats The cumulative statistics for all environments.

type EnvironmentStats struct {
	// The total number of environments.
	EnvTotal int32 `json:"envTotal"`
	// The cumulative CPU capacity of all environments, typically measured in CPU units.
	CpuTotal float64 `json:"cpuTotal"`
	// The cumulative memory capacity of all environments, typically measured in bytes or GiB.
	MemoryTotal float64 `json:"memoryTotal"`
	// The cumulative storage capacity of all environments, typically measured in bytes or GiB.
	StorageTotal float64 `json:"storageTotal"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironmentStats instantiates a new EnvironmentStats object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironmentStats(envTotal int32, cpuTotal float64, memoryTotal float64, storageTotal float64) *EnvironmentStats {
	this := EnvironmentStats{}
	this.EnvTotal = envTotal
	this.CpuTotal = cpuTotal
	this.MemoryTotal = memoryTotal
	this.StorageTotal = storageTotal
	return &this
}

// NewEnvironmentStatsWithDefaults instantiates a new EnvironmentStats object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironmentStatsWithDefaults() *EnvironmentStats {
	this := EnvironmentStats{}
	return &this
}

// GetEnvTotal returns the EnvTotal field value.
func (o *EnvironmentStats) GetEnvTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.EnvTotal
}

// GetEnvTotalOk returns a tuple with the EnvTotal field value
// and a boolean to check if the value has been set.
func (o *EnvironmentStats) GetEnvTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvTotal, true
}

// SetEnvTotal sets field value.
func (o *EnvironmentStats) SetEnvTotal(v int32) {
	o.EnvTotal = v
}

// GetCpuTotal returns the CpuTotal field value.
func (o *EnvironmentStats) GetCpuTotal() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.CpuTotal
}

// GetCpuTotalOk returns a tuple with the CpuTotal field value
// and a boolean to check if the value has been set.
func (o *EnvironmentStats) GetCpuTotalOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CpuTotal, true
}

// SetCpuTotal sets field value.
func (o *EnvironmentStats) SetCpuTotal(v float64) {
	o.CpuTotal = v
}

// GetMemoryTotal returns the MemoryTotal field value.
func (o *EnvironmentStats) GetMemoryTotal() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.MemoryTotal
}

// GetMemoryTotalOk returns a tuple with the MemoryTotal field value
// and a boolean to check if the value has been set.
func (o *EnvironmentStats) GetMemoryTotalOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MemoryTotal, true
}

// SetMemoryTotal sets field value.
func (o *EnvironmentStats) SetMemoryTotal(v float64) {
	o.MemoryTotal = v
}

// GetStorageTotal returns the StorageTotal field value.
func (o *EnvironmentStats) GetStorageTotal() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.StorageTotal
}

// GetStorageTotalOk returns a tuple with the StorageTotal field value
// and a boolean to check if the value has been set.
func (o *EnvironmentStats) GetStorageTotalOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageTotal, true
}

// SetStorageTotal sets field value.
func (o *EnvironmentStats) SetStorageTotal(v float64) {
	o.StorageTotal = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EnvironmentStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["envTotal"] = o.EnvTotal
	toSerialize["cpuTotal"] = o.CpuTotal
	toSerialize["memoryTotal"] = o.MemoryTotal
	toSerialize["storageTotal"] = o.StorageTotal

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EnvironmentStats) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EnvTotal     *int32   `json:"envTotal"`
		CpuTotal     *float64 `json:"cpuTotal"`
		MemoryTotal  *float64 `json:"memoryTotal"`
		StorageTotal *float64 `json:"storageTotal"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.EnvTotal == nil {
		return fmt.Errorf("required field envTotal missing")
	}
	if all.CpuTotal == nil {
		return fmt.Errorf("required field cpuTotal missing")
	}
	if all.MemoryTotal == nil {
		return fmt.Errorf("required field memoryTotal missing")
	}
	if all.StorageTotal == nil {
		return fmt.Errorf("required field storageTotal missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"envTotal", "cpuTotal", "memoryTotal", "storageTotal"})
	} else {
		return err
	}
	o.EnvTotal = *all.EnvTotal
	o.CpuTotal = *all.CpuTotal
	o.MemoryTotal = *all.MemoryTotal
	o.StorageTotal = *all.StorageTotal

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
