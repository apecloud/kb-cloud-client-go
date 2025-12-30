// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type RbmqMessageStats struct {
	// Messages published in
	PublishIn *int32 `json:"publish_in,omitempty"`
	// Rate of messages published in per second
	PublishInRate *float64 `json:"publish_in_rate,omitempty"`
	// Messages published out
	PublishOut *int32 `json:"publish_out,omitempty"`
	// Rate of messages published out per second
	PublishOutRate *float64 `json:"publish_out_rate,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRbmqMessageStats instantiates a new RbmqMessageStats object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRbmqMessageStats() *RbmqMessageStats {
	this := RbmqMessageStats{}
	return &this
}

// NewRbmqMessageStatsWithDefaults instantiates a new RbmqMessageStats object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRbmqMessageStatsWithDefaults() *RbmqMessageStats {
	this := RbmqMessageStats{}
	return &this
}

// GetPublishIn returns the PublishIn field value if set, zero value otherwise.
func (o *RbmqMessageStats) GetPublishIn() int32 {
	if o == nil || o.PublishIn == nil {
		var ret int32
		return ret
	}
	return *o.PublishIn
}

// GetPublishInOk returns a tuple with the PublishIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbmqMessageStats) GetPublishInOk() (*int32, bool) {
	if o == nil || o.PublishIn == nil {
		return nil, false
	}
	return o.PublishIn, true
}

// HasPublishIn returns a boolean if a field has been set.
func (o *RbmqMessageStats) HasPublishIn() bool {
	return o != nil && o.PublishIn != nil
}

// SetPublishIn gets a reference to the given int32 and assigns it to the PublishIn field.
func (o *RbmqMessageStats) SetPublishIn(v int32) {
	o.PublishIn = &v
}

// GetPublishInRate returns the PublishInRate field value if set, zero value otherwise.
func (o *RbmqMessageStats) GetPublishInRate() float64 {
	if o == nil || o.PublishInRate == nil {
		var ret float64
		return ret
	}
	return *o.PublishInRate
}

// GetPublishInRateOk returns a tuple with the PublishInRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbmqMessageStats) GetPublishInRateOk() (*float64, bool) {
	if o == nil || o.PublishInRate == nil {
		return nil, false
	}
	return o.PublishInRate, true
}

// HasPublishInRate returns a boolean if a field has been set.
func (o *RbmqMessageStats) HasPublishInRate() bool {
	return o != nil && o.PublishInRate != nil
}

// SetPublishInRate gets a reference to the given float64 and assigns it to the PublishInRate field.
func (o *RbmqMessageStats) SetPublishInRate(v float64) {
	o.PublishInRate = &v
}

// GetPublishOut returns the PublishOut field value if set, zero value otherwise.
func (o *RbmqMessageStats) GetPublishOut() int32 {
	if o == nil || o.PublishOut == nil {
		var ret int32
		return ret
	}
	return *o.PublishOut
}

// GetPublishOutOk returns a tuple with the PublishOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbmqMessageStats) GetPublishOutOk() (*int32, bool) {
	if o == nil || o.PublishOut == nil {
		return nil, false
	}
	return o.PublishOut, true
}

// HasPublishOut returns a boolean if a field has been set.
func (o *RbmqMessageStats) HasPublishOut() bool {
	return o != nil && o.PublishOut != nil
}

// SetPublishOut gets a reference to the given int32 and assigns it to the PublishOut field.
func (o *RbmqMessageStats) SetPublishOut(v int32) {
	o.PublishOut = &v
}

// GetPublishOutRate returns the PublishOutRate field value if set, zero value otherwise.
func (o *RbmqMessageStats) GetPublishOutRate() float64 {
	if o == nil || o.PublishOutRate == nil {
		var ret float64
		return ret
	}
	return *o.PublishOutRate
}

// GetPublishOutRateOk returns a tuple with the PublishOutRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbmqMessageStats) GetPublishOutRateOk() (*float64, bool) {
	if o == nil || o.PublishOutRate == nil {
		return nil, false
	}
	return o.PublishOutRate, true
}

// HasPublishOutRate returns a boolean if a field has been set.
func (o *RbmqMessageStats) HasPublishOutRate() bool {
	return o != nil && o.PublishOutRate != nil
}

// SetPublishOutRate gets a reference to the given float64 and assigns it to the PublishOutRate field.
func (o *RbmqMessageStats) SetPublishOutRate(v float64) {
	o.PublishOutRate = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RbmqMessageStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.PublishIn != nil {
		toSerialize["publish_in"] = o.PublishIn
	}
	if o.PublishInRate != nil {
		toSerialize["publish_in_rate"] = o.PublishInRate
	}
	if o.PublishOut != nil {
		toSerialize["publish_out"] = o.PublishOut
	}
	if o.PublishOutRate != nil {
		toSerialize["publish_out_rate"] = o.PublishOutRate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RbmqMessageStats) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		PublishIn      *int32   `json:"publish_in,omitempty"`
		PublishInRate  *float64 `json:"publish_in_rate,omitempty"`
		PublishOut     *int32   `json:"publish_out,omitempty"`
		PublishOutRate *float64 `json:"publish_out_rate,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"publish_in", "publish_in_rate", "publish_out", "publish_out_rate"})
	} else {
		return err
	}
	o.PublishIn = all.PublishIn
	o.PublishInRate = all.PublishInRate
	o.PublishOut = all.PublishOut
	o.PublishOutRate = all.PublishOutRate

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
