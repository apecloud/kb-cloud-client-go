// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EnvironmentPricing the information of environment pricing
type EnvironmentPricing struct {
	// The environment name
	EnvName common.NullableString `json:"envName"`
	// The price of CPU, the unit is 'Core'
	CpuPrice common.NullableString `json:"cpuPrice"`
	// The price of Memory, the unit is 'GB'
	MemoryPrice common.NullableString `json:"memoryPrice"`
	// The price of Storage, the unit is 'GB'
	StoragePrice common.NullableString `json:"storagePrice"`
	// The price of Backup, the unit is 'GB'
	BackupPrice common.NullableString `json:"backupPrice"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironmentPricing instantiates a new EnvironmentPricing object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironmentPricing(envName common.NullableString, cpuPrice common.NullableString, memoryPrice common.NullableString, storagePrice common.NullableString, backupPrice common.NullableString) *EnvironmentPricing {
	this := EnvironmentPricing{}
	this.EnvName = envName
	this.CpuPrice = cpuPrice
	this.MemoryPrice = memoryPrice
	this.StoragePrice = storagePrice
	this.BackupPrice = backupPrice
	return &this
}

// NewEnvironmentPricingWithDefaults instantiates a new EnvironmentPricing object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironmentPricingWithDefaults() *EnvironmentPricing {
	this := EnvironmentPricing{}
	return &this
}

// GetEnvName returns the EnvName field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *EnvironmentPricing) GetEnvName() string {
	if o == nil || o.EnvName.Get() == nil {
		var ret string
		return ret
	}
	return *o.EnvName.Get()
}

// GetEnvNameOk returns a tuple with the EnvName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EnvironmentPricing) GetEnvNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnvName.Get(), o.EnvName.IsSet()
}

// SetEnvName sets field value.
func (o *EnvironmentPricing) SetEnvName(v string) {
	o.EnvName.Set(&v)
}

// GetCpuPrice returns the CpuPrice field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *EnvironmentPricing) GetCpuPrice() string {
	if o == nil || o.CpuPrice.Get() == nil {
		var ret string
		return ret
	}
	return *o.CpuPrice.Get()
}

// GetCpuPriceOk returns a tuple with the CpuPrice field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EnvironmentPricing) GetCpuPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CpuPrice.Get(), o.CpuPrice.IsSet()
}

// SetCpuPrice sets field value.
func (o *EnvironmentPricing) SetCpuPrice(v string) {
	o.CpuPrice.Set(&v)
}

// GetMemoryPrice returns the MemoryPrice field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *EnvironmentPricing) GetMemoryPrice() string {
	if o == nil || o.MemoryPrice.Get() == nil {
		var ret string
		return ret
	}
	return *o.MemoryPrice.Get()
}

// GetMemoryPriceOk returns a tuple with the MemoryPrice field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EnvironmentPricing) GetMemoryPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MemoryPrice.Get(), o.MemoryPrice.IsSet()
}

// SetMemoryPrice sets field value.
func (o *EnvironmentPricing) SetMemoryPrice(v string) {
	o.MemoryPrice.Set(&v)
}

// GetStoragePrice returns the StoragePrice field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *EnvironmentPricing) GetStoragePrice() string {
	if o == nil || o.StoragePrice.Get() == nil {
		var ret string
		return ret
	}
	return *o.StoragePrice.Get()
}

// GetStoragePriceOk returns a tuple with the StoragePrice field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EnvironmentPricing) GetStoragePriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StoragePrice.Get(), o.StoragePrice.IsSet()
}

// SetStoragePrice sets field value.
func (o *EnvironmentPricing) SetStoragePrice(v string) {
	o.StoragePrice.Set(&v)
}

// GetBackupPrice returns the BackupPrice field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *EnvironmentPricing) GetBackupPrice() string {
	if o == nil || o.BackupPrice.Get() == nil {
		var ret string
		return ret
	}
	return *o.BackupPrice.Get()
}

// GetBackupPriceOk returns a tuple with the BackupPrice field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EnvironmentPricing) GetBackupPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BackupPrice.Get(), o.BackupPrice.IsSet()
}

// SetBackupPrice sets field value.
func (o *EnvironmentPricing) SetBackupPrice(v string) {
	o.BackupPrice.Set(&v)
}

// MarshalJSON serializes the struct using spec logic.
func (o EnvironmentPricing) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["envName"] = o.EnvName.Get()
	toSerialize["cpuPrice"] = o.CpuPrice.Get()
	toSerialize["memoryPrice"] = o.MemoryPrice.Get()
	toSerialize["storagePrice"] = o.StoragePrice.Get()
	toSerialize["backupPrice"] = o.BackupPrice.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EnvironmentPricing) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EnvName      common.NullableString `json:"envName"`
		CpuPrice     common.NullableString `json:"cpuPrice"`
		MemoryPrice  common.NullableString `json:"memoryPrice"`
		StoragePrice common.NullableString `json:"storagePrice"`
		BackupPrice  common.NullableString `json:"backupPrice"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if !all.EnvName.IsSet() {
		return fmt.Errorf("required field envName missing")
	}
	if !all.CpuPrice.IsSet() {
		return fmt.Errorf("required field cpuPrice missing")
	}
	if !all.MemoryPrice.IsSet() {
		return fmt.Errorf("required field memoryPrice missing")
	}
	if !all.StoragePrice.IsSet() {
		return fmt.Errorf("required field storagePrice missing")
	}
	if !all.BackupPrice.IsSet() {
		return fmt.Errorf("required field backupPrice missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"envName", "cpuPrice", "memoryPrice", "storagePrice", "backupPrice"})
	} else {
		return err
	}
	o.EnvName = all.EnvName
	o.CpuPrice = all.CpuPrice
	o.MemoryPrice = all.MemoryPrice
	o.StoragePrice = all.StoragePrice
	o.BackupPrice = all.BackupPrice

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
