// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// EnvironmentPricing the information of environment pricing
type EnvironmentPricing struct {
	// The environment name
	EnvName common.NullableString `json:"envName,omitempty"`
	// The price of CPU, the unit is 'Core'
	CpuPrice common.NullableString `json:"cpuPrice,omitempty"`
	// The price of Memory, the unit is 'GB'
	MemoryPrice common.NullableString `json:"memoryPrice,omitempty"`
	// The price of Storage, the unit is 'GB'
	StoragePrice common.NullableString `json:"storagePrice,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironmentPricing instantiates a new EnvironmentPricing object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironmentPricing() *EnvironmentPricing {
	this := EnvironmentPricing{}
	return &this
}

// NewEnvironmentPricingWithDefaults instantiates a new EnvironmentPricing object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironmentPricingWithDefaults() *EnvironmentPricing {
	this := EnvironmentPricing{}
	return &this
}

// GetEnvName returns the EnvName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentPricing) GetEnvName() string {
	if o == nil || o.EnvName.Get() == nil {
		var ret string
		return ret
	}
	return *o.EnvName.Get()
}

// GetEnvNameOk returns a tuple with the EnvName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EnvironmentPricing) GetEnvNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnvName.Get(), o.EnvName.IsSet()
}

// HasEnvName returns a boolean if a field has been set.
func (o *EnvironmentPricing) HasEnvName() bool {
	return o != nil && o.EnvName.IsSet()
}

// SetEnvName gets a reference to the given common.NullableString and assigns it to the EnvName field.
func (o *EnvironmentPricing) SetEnvName(v string) {
	o.EnvName.Set(&v)
}

// SetEnvNameNil sets the value for EnvName to be an explicit nil.
func (o *EnvironmentPricing) SetEnvNameNil() {
	o.EnvName.Set(nil)
}

// UnsetEnvName ensures that no value is present for EnvName, not even an explicit nil.
func (o *EnvironmentPricing) UnsetEnvName() {
	o.EnvName.Unset()
}

// GetCpuPrice returns the CpuPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentPricing) GetCpuPrice() string {
	if o == nil || o.CpuPrice.Get() == nil {
		var ret string
		return ret
	}
	return *o.CpuPrice.Get()
}

// GetCpuPriceOk returns a tuple with the CpuPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EnvironmentPricing) GetCpuPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CpuPrice.Get(), o.CpuPrice.IsSet()
}

// HasCpuPrice returns a boolean if a field has been set.
func (o *EnvironmentPricing) HasCpuPrice() bool {
	return o != nil && o.CpuPrice.IsSet()
}

// SetCpuPrice gets a reference to the given common.NullableString and assigns it to the CpuPrice field.
func (o *EnvironmentPricing) SetCpuPrice(v string) {
	o.CpuPrice.Set(&v)
}

// SetCpuPriceNil sets the value for CpuPrice to be an explicit nil.
func (o *EnvironmentPricing) SetCpuPriceNil() {
	o.CpuPrice.Set(nil)
}

// UnsetCpuPrice ensures that no value is present for CpuPrice, not even an explicit nil.
func (o *EnvironmentPricing) UnsetCpuPrice() {
	o.CpuPrice.Unset()
}

// GetMemoryPrice returns the MemoryPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentPricing) GetMemoryPrice() string {
	if o == nil || o.MemoryPrice.Get() == nil {
		var ret string
		return ret
	}
	return *o.MemoryPrice.Get()
}

// GetMemoryPriceOk returns a tuple with the MemoryPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EnvironmentPricing) GetMemoryPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MemoryPrice.Get(), o.MemoryPrice.IsSet()
}

// HasMemoryPrice returns a boolean if a field has been set.
func (o *EnvironmentPricing) HasMemoryPrice() bool {
	return o != nil && o.MemoryPrice.IsSet()
}

// SetMemoryPrice gets a reference to the given common.NullableString and assigns it to the MemoryPrice field.
func (o *EnvironmentPricing) SetMemoryPrice(v string) {
	o.MemoryPrice.Set(&v)
}

// SetMemoryPriceNil sets the value for MemoryPrice to be an explicit nil.
func (o *EnvironmentPricing) SetMemoryPriceNil() {
	o.MemoryPrice.Set(nil)
}

// UnsetMemoryPrice ensures that no value is present for MemoryPrice, not even an explicit nil.
func (o *EnvironmentPricing) UnsetMemoryPrice() {
	o.MemoryPrice.Unset()
}

// GetStoragePrice returns the StoragePrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentPricing) GetStoragePrice() string {
	if o == nil || o.StoragePrice.Get() == nil {
		var ret string
		return ret
	}
	return *o.StoragePrice.Get()
}

// GetStoragePriceOk returns a tuple with the StoragePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EnvironmentPricing) GetStoragePriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StoragePrice.Get(), o.StoragePrice.IsSet()
}

// HasStoragePrice returns a boolean if a field has been set.
func (o *EnvironmentPricing) HasStoragePrice() bool {
	return o != nil && o.StoragePrice.IsSet()
}

// SetStoragePrice gets a reference to the given common.NullableString and assigns it to the StoragePrice field.
func (o *EnvironmentPricing) SetStoragePrice(v string) {
	o.StoragePrice.Set(&v)
}

// SetStoragePriceNil sets the value for StoragePrice to be an explicit nil.
func (o *EnvironmentPricing) SetStoragePriceNil() {
	o.StoragePrice.Set(nil)
}

// UnsetStoragePrice ensures that no value is present for StoragePrice, not even an explicit nil.
func (o *EnvironmentPricing) UnsetStoragePrice() {
	o.StoragePrice.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o EnvironmentPricing) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.EnvName.IsSet() {
		toSerialize["envName"] = o.EnvName.Get()
	}
	if o.CpuPrice.IsSet() {
		toSerialize["cpuPrice"] = o.CpuPrice.Get()
	}
	if o.MemoryPrice.IsSet() {
		toSerialize["memoryPrice"] = o.MemoryPrice.Get()
	}
	if o.StoragePrice.IsSet() {
		toSerialize["storagePrice"] = o.StoragePrice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EnvironmentPricing) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EnvName      common.NullableString `json:"envName,omitempty"`
		CpuPrice     common.NullableString `json:"cpuPrice,omitempty"`
		MemoryPrice  common.NullableString `json:"memoryPrice,omitempty"`
		StoragePrice common.NullableString `json:"storagePrice,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"envName", "cpuPrice", "memoryPrice", "storagePrice"})
	} else {
		return err
	}
	o.EnvName = all.EnvName
	o.CpuPrice = all.CpuPrice
	o.MemoryPrice = all.MemoryPrice
	o.StoragePrice = all.StoragePrice

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
