// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ElasticsearchRecoverySize struct {
	TotalBytes     common.NullableInt64  `json:"totalBytes,omitempty"`
	ReusedBytes    common.NullableInt64  `json:"reusedBytes,omitempty"`
	RecoveredBytes common.NullableInt64  `json:"recoveredBytes,omitempty"`
	Percent        common.NullableString `json:"percent,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewElasticsearchRecoverySize instantiates a new ElasticsearchRecoverySize object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewElasticsearchRecoverySize() *ElasticsearchRecoverySize {
	this := ElasticsearchRecoverySize{}
	return &this
}

// NewElasticsearchRecoverySizeWithDefaults instantiates a new ElasticsearchRecoverySize object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewElasticsearchRecoverySizeWithDefaults() *ElasticsearchRecoverySize {
	this := ElasticsearchRecoverySize{}
	return &this
}

// GetTotalBytes returns the TotalBytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElasticsearchRecoverySize) GetTotalBytes() int64 {
	if o == nil || o.TotalBytes.Get() == nil {
		var ret int64
		return ret
	}
	return *o.TotalBytes.Get()
}

// GetTotalBytesOk returns a tuple with the TotalBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchRecoverySize) GetTotalBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalBytes.Get(), o.TotalBytes.IsSet()
}

// HasTotalBytes returns a boolean if a field has been set.
func (o *ElasticsearchRecoverySize) HasTotalBytes() bool {
	return o != nil && o.TotalBytes.IsSet()
}

// SetTotalBytes gets a reference to the given common.NullableInt64 and assigns it to the TotalBytes field.
func (o *ElasticsearchRecoverySize) SetTotalBytes(v int64) {
	o.TotalBytes.Set(&v)
}

// SetTotalBytesNil sets the value for TotalBytes to be an explicit nil.
func (o *ElasticsearchRecoverySize) SetTotalBytesNil() {
	o.TotalBytes.Set(nil)
}

// UnsetTotalBytes ensures that no value is present for TotalBytes, not even an explicit nil.
func (o *ElasticsearchRecoverySize) UnsetTotalBytes() {
	o.TotalBytes.Unset()
}

// GetReusedBytes returns the ReusedBytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElasticsearchRecoverySize) GetReusedBytes() int64 {
	if o == nil || o.ReusedBytes.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ReusedBytes.Get()
}

// GetReusedBytesOk returns a tuple with the ReusedBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchRecoverySize) GetReusedBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReusedBytes.Get(), o.ReusedBytes.IsSet()
}

// HasReusedBytes returns a boolean if a field has been set.
func (o *ElasticsearchRecoverySize) HasReusedBytes() bool {
	return o != nil && o.ReusedBytes.IsSet()
}

// SetReusedBytes gets a reference to the given common.NullableInt64 and assigns it to the ReusedBytes field.
func (o *ElasticsearchRecoverySize) SetReusedBytes(v int64) {
	o.ReusedBytes.Set(&v)
}

// SetReusedBytesNil sets the value for ReusedBytes to be an explicit nil.
func (o *ElasticsearchRecoverySize) SetReusedBytesNil() {
	o.ReusedBytes.Set(nil)
}

// UnsetReusedBytes ensures that no value is present for ReusedBytes, not even an explicit nil.
func (o *ElasticsearchRecoverySize) UnsetReusedBytes() {
	o.ReusedBytes.Unset()
}

// GetRecoveredBytes returns the RecoveredBytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElasticsearchRecoverySize) GetRecoveredBytes() int64 {
	if o == nil || o.RecoveredBytes.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RecoveredBytes.Get()
}

// GetRecoveredBytesOk returns a tuple with the RecoveredBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchRecoverySize) GetRecoveredBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecoveredBytes.Get(), o.RecoveredBytes.IsSet()
}

// HasRecoveredBytes returns a boolean if a field has been set.
func (o *ElasticsearchRecoverySize) HasRecoveredBytes() bool {
	return o != nil && o.RecoveredBytes.IsSet()
}

// SetRecoveredBytes gets a reference to the given common.NullableInt64 and assigns it to the RecoveredBytes field.
func (o *ElasticsearchRecoverySize) SetRecoveredBytes(v int64) {
	o.RecoveredBytes.Set(&v)
}

// SetRecoveredBytesNil sets the value for RecoveredBytes to be an explicit nil.
func (o *ElasticsearchRecoverySize) SetRecoveredBytesNil() {
	o.RecoveredBytes.Set(nil)
}

// UnsetRecoveredBytes ensures that no value is present for RecoveredBytes, not even an explicit nil.
func (o *ElasticsearchRecoverySize) UnsetRecoveredBytes() {
	o.RecoveredBytes.Unset()
}

// GetPercent returns the Percent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElasticsearchRecoverySize) GetPercent() string {
	if o == nil || o.Percent.Get() == nil {
		var ret string
		return ret
	}
	return *o.Percent.Get()
}

// GetPercentOk returns a tuple with the Percent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchRecoverySize) GetPercentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Percent.Get(), o.Percent.IsSet()
}

// HasPercent returns a boolean if a field has been set.
func (o *ElasticsearchRecoverySize) HasPercent() bool {
	return o != nil && o.Percent.IsSet()
}

// SetPercent gets a reference to the given common.NullableString and assigns it to the Percent field.
func (o *ElasticsearchRecoverySize) SetPercent(v string) {
	o.Percent.Set(&v)
}

// SetPercentNil sets the value for Percent to be an explicit nil.
func (o *ElasticsearchRecoverySize) SetPercentNil() {
	o.Percent.Set(nil)
}

// UnsetPercent ensures that no value is present for Percent, not even an explicit nil.
func (o *ElasticsearchRecoverySize) UnsetPercent() {
	o.Percent.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ElasticsearchRecoverySize) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.TotalBytes.IsSet() {
		toSerialize["totalBytes"] = o.TotalBytes.Get()
	}
	if o.ReusedBytes.IsSet() {
		toSerialize["reusedBytes"] = o.ReusedBytes.Get()
	}
	if o.RecoveredBytes.IsSet() {
		toSerialize["recoveredBytes"] = o.RecoveredBytes.Get()
	}
	if o.Percent.IsSet() {
		toSerialize["percent"] = o.Percent.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ElasticsearchRecoverySize) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TotalBytes     common.NullableInt64  `json:"totalBytes,omitempty"`
		ReusedBytes    common.NullableInt64  `json:"reusedBytes,omitempty"`
		RecoveredBytes common.NullableInt64  `json:"recoveredBytes,omitempty"`
		Percent        common.NullableString `json:"percent,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"totalBytes", "reusedBytes", "recoveredBytes", "percent"})
	} else {
		return err
	}
	o.TotalBytes = all.TotalBytes
	o.ReusedBytes = all.ReusedBytes
	o.RecoveredBytes = all.RecoveredBytes
	o.Percent = all.Percent

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
