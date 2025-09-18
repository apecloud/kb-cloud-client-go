// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// GlobalPricing the global information of pricing
type GlobalPricing struct {
	// Enable global pricing
	Enabled *bool `json:"enabled,omitempty"`
	// The currency uint for display in bill query
	CurrencyUnit common.NullableString `json:"currencyUnit,omitempty"`
	// Scheduled production time for daily bills, the format is 'HH:mm:ss', such as '01:00:00' meas every day at 1:00 AM, the timezone is the configured timezone
	BillScheduleTime common.NullableString `json:"billScheduleTime,omitempty"`
	// The last billing time in RFC3339 format (e.g., '2025-09-01T00:00:00Z' or '2025-09-01T00:00:00+08:00').
	// This timestamp represents when the end time of the latest bill.
	//
	LastBillingTime common.NullableTime `json:"lastBillingTime,omitempty"`
	// The price of CPU, the unit is 'Core'
	CpuPrice common.NullableString `json:"cpuPrice,omitempty"`
	// The price of Memory, the unit is 'GB'
	MemoryPrice common.NullableString `json:"memoryPrice,omitempty"`
	// The price of Storage, the unit is 'GB'
	StoragePrice common.NullableString `json:"storagePrice,omitempty"`
	// The timezone used for billing calculations and time alignment, specified as an IANA timezone identifier.
	//
	// This timezone determines:
	// - When daily billing periods start and end (aligned to midnight in this timezone)
	// - How input timestamps are converted and normalized for billing calculations
	// - The timezone context for scheduled billing operations
	//
	// Examples of valid timezone values:
	// - "UTC" - Coordinated Universal Time
	// - "Asia/Shanghai" - China Standard Time (CST, UTC+8)
	// - "America/New_York" - Eastern Time (EST/EDT, UTC-5/-4)
	// - "Europe/London" - Greenwich Mean Time/British Summer Time (GMT/BST, UTC+0/+1)
	//
	// **The system defaults to UTC and not support to change for now.**
	//
	Timezone common.NullableString `json:"timezone,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGlobalPricing instantiates a new GlobalPricing object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGlobalPricing() *GlobalPricing {
	this := GlobalPricing{}
	return &this
}

// NewGlobalPricingWithDefaults instantiates a new GlobalPricing object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGlobalPricingWithDefaults() *GlobalPricing {
	this := GlobalPricing{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GlobalPricing) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalPricing) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GlobalPricing) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GlobalPricing) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetCurrencyUnit returns the CurrencyUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GlobalPricing) GetCurrencyUnit() string {
	if o == nil || o.CurrencyUnit.Get() == nil {
		var ret string
		return ret
	}
	return *o.CurrencyUnit.Get()
}

// GetCurrencyUnitOk returns a tuple with the CurrencyUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *GlobalPricing) GetCurrencyUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrencyUnit.Get(), o.CurrencyUnit.IsSet()
}

// HasCurrencyUnit returns a boolean if a field has been set.
func (o *GlobalPricing) HasCurrencyUnit() bool {
	return o != nil && o.CurrencyUnit.IsSet()
}

// SetCurrencyUnit gets a reference to the given common.NullableString and assigns it to the CurrencyUnit field.
func (o *GlobalPricing) SetCurrencyUnit(v string) {
	o.CurrencyUnit.Set(&v)
}

// SetCurrencyUnitNil sets the value for CurrencyUnit to be an explicit nil.
func (o *GlobalPricing) SetCurrencyUnitNil() {
	o.CurrencyUnit.Set(nil)
}

// UnsetCurrencyUnit ensures that no value is present for CurrencyUnit, not even an explicit nil.
func (o *GlobalPricing) UnsetCurrencyUnit() {
	o.CurrencyUnit.Unset()
}

// GetBillScheduleTime returns the BillScheduleTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GlobalPricing) GetBillScheduleTime() string {
	if o == nil || o.BillScheduleTime.Get() == nil {
		var ret string
		return ret
	}
	return *o.BillScheduleTime.Get()
}

// GetBillScheduleTimeOk returns a tuple with the BillScheduleTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *GlobalPricing) GetBillScheduleTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillScheduleTime.Get(), o.BillScheduleTime.IsSet()
}

// HasBillScheduleTime returns a boolean if a field has been set.
func (o *GlobalPricing) HasBillScheduleTime() bool {
	return o != nil && o.BillScheduleTime.IsSet()
}

// SetBillScheduleTime gets a reference to the given common.NullableString and assigns it to the BillScheduleTime field.
func (o *GlobalPricing) SetBillScheduleTime(v string) {
	o.BillScheduleTime.Set(&v)
}

// SetBillScheduleTimeNil sets the value for BillScheduleTime to be an explicit nil.
func (o *GlobalPricing) SetBillScheduleTimeNil() {
	o.BillScheduleTime.Set(nil)
}

// UnsetBillScheduleTime ensures that no value is present for BillScheduleTime, not even an explicit nil.
func (o *GlobalPricing) UnsetBillScheduleTime() {
	o.BillScheduleTime.Unset()
}

// GetLastBillingTime returns the LastBillingTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GlobalPricing) GetLastBillingTime() time.Time {
	if o == nil || o.LastBillingTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastBillingTime.Get()
}

// GetLastBillingTimeOk returns a tuple with the LastBillingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *GlobalPricing) GetLastBillingTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastBillingTime.Get(), o.LastBillingTime.IsSet()
}

// HasLastBillingTime returns a boolean if a field has been set.
func (o *GlobalPricing) HasLastBillingTime() bool {
	return o != nil && o.LastBillingTime.IsSet()
}

// SetLastBillingTime gets a reference to the given common.NullableTime and assigns it to the LastBillingTime field.
func (o *GlobalPricing) SetLastBillingTime(v time.Time) {
	o.LastBillingTime.Set(&v)
}

// SetLastBillingTimeNil sets the value for LastBillingTime to be an explicit nil.
func (o *GlobalPricing) SetLastBillingTimeNil() {
	o.LastBillingTime.Set(nil)
}

// UnsetLastBillingTime ensures that no value is present for LastBillingTime, not even an explicit nil.
func (o *GlobalPricing) UnsetLastBillingTime() {
	o.LastBillingTime.Unset()
}

// GetCpuPrice returns the CpuPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GlobalPricing) GetCpuPrice() string {
	if o == nil || o.CpuPrice.Get() == nil {
		var ret string
		return ret
	}
	return *o.CpuPrice.Get()
}

// GetCpuPriceOk returns a tuple with the CpuPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *GlobalPricing) GetCpuPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CpuPrice.Get(), o.CpuPrice.IsSet()
}

// HasCpuPrice returns a boolean if a field has been set.
func (o *GlobalPricing) HasCpuPrice() bool {
	return o != nil && o.CpuPrice.IsSet()
}

// SetCpuPrice gets a reference to the given common.NullableString and assigns it to the CpuPrice field.
func (o *GlobalPricing) SetCpuPrice(v string) {
	o.CpuPrice.Set(&v)
}

// SetCpuPriceNil sets the value for CpuPrice to be an explicit nil.
func (o *GlobalPricing) SetCpuPriceNil() {
	o.CpuPrice.Set(nil)
}

// UnsetCpuPrice ensures that no value is present for CpuPrice, not even an explicit nil.
func (o *GlobalPricing) UnsetCpuPrice() {
	o.CpuPrice.Unset()
}

// GetMemoryPrice returns the MemoryPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GlobalPricing) GetMemoryPrice() string {
	if o == nil || o.MemoryPrice.Get() == nil {
		var ret string
		return ret
	}
	return *o.MemoryPrice.Get()
}

// GetMemoryPriceOk returns a tuple with the MemoryPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *GlobalPricing) GetMemoryPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MemoryPrice.Get(), o.MemoryPrice.IsSet()
}

// HasMemoryPrice returns a boolean if a field has been set.
func (o *GlobalPricing) HasMemoryPrice() bool {
	return o != nil && o.MemoryPrice.IsSet()
}

// SetMemoryPrice gets a reference to the given common.NullableString and assigns it to the MemoryPrice field.
func (o *GlobalPricing) SetMemoryPrice(v string) {
	o.MemoryPrice.Set(&v)
}

// SetMemoryPriceNil sets the value for MemoryPrice to be an explicit nil.
func (o *GlobalPricing) SetMemoryPriceNil() {
	o.MemoryPrice.Set(nil)
}

// UnsetMemoryPrice ensures that no value is present for MemoryPrice, not even an explicit nil.
func (o *GlobalPricing) UnsetMemoryPrice() {
	o.MemoryPrice.Unset()
}

// GetStoragePrice returns the StoragePrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GlobalPricing) GetStoragePrice() string {
	if o == nil || o.StoragePrice.Get() == nil {
		var ret string
		return ret
	}
	return *o.StoragePrice.Get()
}

// GetStoragePriceOk returns a tuple with the StoragePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *GlobalPricing) GetStoragePriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StoragePrice.Get(), o.StoragePrice.IsSet()
}

// HasStoragePrice returns a boolean if a field has been set.
func (o *GlobalPricing) HasStoragePrice() bool {
	return o != nil && o.StoragePrice.IsSet()
}

// SetStoragePrice gets a reference to the given common.NullableString and assigns it to the StoragePrice field.
func (o *GlobalPricing) SetStoragePrice(v string) {
	o.StoragePrice.Set(&v)
}

// SetStoragePriceNil sets the value for StoragePrice to be an explicit nil.
func (o *GlobalPricing) SetStoragePriceNil() {
	o.StoragePrice.Set(nil)
}

// UnsetStoragePrice ensures that no value is present for StoragePrice, not even an explicit nil.
func (o *GlobalPricing) UnsetStoragePrice() {
	o.StoragePrice.Unset()
}

// GetTimezone returns the Timezone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GlobalPricing) GetTimezone() string {
	if o == nil || o.Timezone.Get() == nil {
		var ret string
		return ret
	}
	return *o.Timezone.Get()
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *GlobalPricing) GetTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timezone.Get(), o.Timezone.IsSet()
}

// HasTimezone returns a boolean if a field has been set.
func (o *GlobalPricing) HasTimezone() bool {
	return o != nil && o.Timezone.IsSet()
}

// SetTimezone gets a reference to the given common.NullableString and assigns it to the Timezone field.
func (o *GlobalPricing) SetTimezone(v string) {
	o.Timezone.Set(&v)
}

// SetTimezoneNil sets the value for Timezone to be an explicit nil.
func (o *GlobalPricing) SetTimezoneNil() {
	o.Timezone.Set(nil)
}

// UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil.
func (o *GlobalPricing) UnsetTimezone() {
	o.Timezone.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o GlobalPricing) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.CurrencyUnit.IsSet() {
		toSerialize["currencyUnit"] = o.CurrencyUnit.Get()
	}
	if o.BillScheduleTime.IsSet() {
		toSerialize["billScheduleTime"] = o.BillScheduleTime.Get()
	}
	if o.LastBillingTime.IsSet() {
		toSerialize["lastBillingTime"] = o.LastBillingTime.Get()
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
	if o.Timezone.IsSet() {
		toSerialize["timezone"] = o.Timezone.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GlobalPricing) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled          *bool                 `json:"enabled,omitempty"`
		CurrencyUnit     common.NullableString `json:"currencyUnit,omitempty"`
		BillScheduleTime common.NullableString `json:"billScheduleTime,omitempty"`
		LastBillingTime  common.NullableTime   `json:"lastBillingTime,omitempty"`
		CpuPrice         common.NullableString `json:"cpuPrice,omitempty"`
		MemoryPrice      common.NullableString `json:"memoryPrice,omitempty"`
		StoragePrice     common.NullableString `json:"storagePrice,omitempty"`
		Timezone         common.NullableString `json:"timezone,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"enabled", "currencyUnit", "billScheduleTime", "lastBillingTime", "cpuPrice", "memoryPrice", "storagePrice", "timezone"})
	} else {
		return err
	}
	o.Enabled = all.Enabled
	o.CurrencyUnit = all.CurrencyUnit
	o.BillScheduleTime = all.BillScheduleTime
	o.LastBillingTime = all.LastBillingTime
	o.CpuPrice = all.CpuPrice
	o.MemoryPrice = all.MemoryPrice
	o.StoragePrice = all.StoragePrice
	o.Timezone = all.Timezone

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
