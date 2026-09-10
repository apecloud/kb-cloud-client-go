// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ElasticsearchRecoveryTranslog struct {
	Recovered common.NullableInt64  `json:"recovered,omitempty"`
	Total     common.NullableInt64  `json:"total,omitempty"`
	Percent   common.NullableString `json:"percent,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewElasticsearchRecoveryTranslog instantiates a new ElasticsearchRecoveryTranslog object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewElasticsearchRecoveryTranslog() *ElasticsearchRecoveryTranslog {
	this := ElasticsearchRecoveryTranslog{}
	return &this
}

// NewElasticsearchRecoveryTranslogWithDefaults instantiates a new ElasticsearchRecoveryTranslog object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewElasticsearchRecoveryTranslogWithDefaults() *ElasticsearchRecoveryTranslog {
	this := ElasticsearchRecoveryTranslog{}
	return &this
}

// GetRecovered returns the Recovered field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElasticsearchRecoveryTranslog) GetRecovered() int64 {
	if o == nil || o.Recovered.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Recovered.Get()
}

// GetRecoveredOk returns a tuple with the Recovered field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchRecoveryTranslog) GetRecoveredOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Recovered.Get(), o.Recovered.IsSet()
}

// HasRecovered returns a boolean if a field has been set.
func (o *ElasticsearchRecoveryTranslog) HasRecovered() bool {
	return o != nil && o.Recovered.IsSet()
}

// SetRecovered gets a reference to the given common.NullableInt64 and assigns it to the Recovered field.
func (o *ElasticsearchRecoveryTranslog) SetRecovered(v int64) {
	o.Recovered.Set(&v)
}

// SetRecoveredNil sets the value for Recovered to be an explicit nil.
func (o *ElasticsearchRecoveryTranslog) SetRecoveredNil() {
	o.Recovered.Set(nil)
}

// UnsetRecovered ensures that no value is present for Recovered, not even an explicit nil.
func (o *ElasticsearchRecoveryTranslog) UnsetRecovered() {
	o.Recovered.Unset()
}

// GetTotal returns the Total field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElasticsearchRecoveryTranslog) GetTotal() int64 {
	if o == nil || o.Total.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Total.Get()
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchRecoveryTranslog) GetTotalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Total.Get(), o.Total.IsSet()
}

// HasTotal returns a boolean if a field has been set.
func (o *ElasticsearchRecoveryTranslog) HasTotal() bool {
	return o != nil && o.Total.IsSet()
}

// SetTotal gets a reference to the given common.NullableInt64 and assigns it to the Total field.
func (o *ElasticsearchRecoveryTranslog) SetTotal(v int64) {
	o.Total.Set(&v)
}

// SetTotalNil sets the value for Total to be an explicit nil.
func (o *ElasticsearchRecoveryTranslog) SetTotalNil() {
	o.Total.Set(nil)
}

// UnsetTotal ensures that no value is present for Total, not even an explicit nil.
func (o *ElasticsearchRecoveryTranslog) UnsetTotal() {
	o.Total.Unset()
}

// GetPercent returns the Percent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElasticsearchRecoveryTranslog) GetPercent() string {
	if o == nil || o.Percent.Get() == nil {
		var ret string
		return ret
	}
	return *o.Percent.Get()
}

// GetPercentOk returns a tuple with the Percent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchRecoveryTranslog) GetPercentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Percent.Get(), o.Percent.IsSet()
}

// HasPercent returns a boolean if a field has been set.
func (o *ElasticsearchRecoveryTranslog) HasPercent() bool {
	return o != nil && o.Percent.IsSet()
}

// SetPercent gets a reference to the given common.NullableString and assigns it to the Percent field.
func (o *ElasticsearchRecoveryTranslog) SetPercent(v string) {
	o.Percent.Set(&v)
}

// SetPercentNil sets the value for Percent to be an explicit nil.
func (o *ElasticsearchRecoveryTranslog) SetPercentNil() {
	o.Percent.Set(nil)
}

// UnsetPercent ensures that no value is present for Percent, not even an explicit nil.
func (o *ElasticsearchRecoveryTranslog) UnsetPercent() {
	o.Percent.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ElasticsearchRecoveryTranslog) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Recovered.IsSet() {
		toSerialize["recovered"] = o.Recovered.Get()
	}
	if o.Total.IsSet() {
		toSerialize["total"] = o.Total.Get()
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
func (o *ElasticsearchRecoveryTranslog) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Recovered common.NullableInt64  `json:"recovered,omitempty"`
		Total     common.NullableInt64  `json:"total,omitempty"`
		Percent   common.NullableString `json:"percent,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"recovered", "total", "percent"})
	} else {
		return err
	}
	o.Recovered = all.Recovered
	o.Total = all.Total
	o.Percent = all.Percent

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
