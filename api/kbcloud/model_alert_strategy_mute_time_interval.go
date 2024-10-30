// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package kbcloud

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)


 
type AlertStrategyMuteTimeInterval struct {
	Weekdays []int32 `json:"weekdays,omitempty"`
	Times *AlertStrategyMuteTimeIntervalTimes `json:"times,omitempty"`
	// Only mute once for `onceMinutes` minutes
	OnceMinutes *int32 `json:"onceMinutes,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewAlertStrategyMuteTimeInterval instantiates a new AlertStrategyMuteTimeInterval object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAlertStrategyMuteTimeInterval() *AlertStrategyMuteTimeInterval {
	this := AlertStrategyMuteTimeInterval{}
	return &this
}

// NewAlertStrategyMuteTimeIntervalWithDefaults instantiates a new AlertStrategyMuteTimeInterval object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAlertStrategyMuteTimeIntervalWithDefaults() *AlertStrategyMuteTimeInterval {
	this := AlertStrategyMuteTimeInterval{}
	return &this
}
// GetWeekdays returns the Weekdays field value if set, zero value otherwise.
func (o *AlertStrategyMuteTimeInterval) GetWeekdays() []int32 {
	if o == nil || o.Weekdays == nil {
		var ret []int32
		return ret
	}
	return o.Weekdays
}

// GetWeekdaysOk returns a tuple with the Weekdays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertStrategyMuteTimeInterval) GetWeekdaysOk() (*[]int32, bool) {
	if o == nil || o.Weekdays == nil {
		return nil, false
	}
	return &o.Weekdays, true
}

// HasWeekdays returns a boolean if a field has been set.
func (o *AlertStrategyMuteTimeInterval) HasWeekdays() bool {
	return o != nil && o.Weekdays != nil
}

// SetWeekdays gets a reference to the given []int32 and assigns it to the Weekdays field.
func (o *AlertStrategyMuteTimeInterval) SetWeekdays(v []int32) {
	o.Weekdays = v
}


// GetTimes returns the Times field value if set, zero value otherwise.
func (o *AlertStrategyMuteTimeInterval) GetTimes() AlertStrategyMuteTimeIntervalTimes {
	if o == nil || o.Times == nil {
		var ret AlertStrategyMuteTimeIntervalTimes
		return ret
	}
	return *o.Times
}

// GetTimesOk returns a tuple with the Times field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertStrategyMuteTimeInterval) GetTimesOk() (*AlertStrategyMuteTimeIntervalTimes, bool) {
	if o == nil || o.Times == nil {
		return nil, false
	}
	return o.Times, true
}

// HasTimes returns a boolean if a field has been set.
func (o *AlertStrategyMuteTimeInterval) HasTimes() bool {
	return o != nil && o.Times != nil
}

// SetTimes gets a reference to the given AlertStrategyMuteTimeIntervalTimes and assigns it to the Times field.
func (o *AlertStrategyMuteTimeInterval) SetTimes(v AlertStrategyMuteTimeIntervalTimes) {
	o.Times = &v
}


// GetOnceMinutes returns the OnceMinutes field value if set, zero value otherwise.
func (o *AlertStrategyMuteTimeInterval) GetOnceMinutes() int32 {
	if o == nil || o.OnceMinutes == nil {
		var ret int32
		return ret
	}
	return *o.OnceMinutes
}

// GetOnceMinutesOk returns a tuple with the OnceMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertStrategyMuteTimeInterval) GetOnceMinutesOk() (*int32, bool) {
	if o == nil || o.OnceMinutes == nil {
		return nil, false
	}
	return o.OnceMinutes, true
}

// HasOnceMinutes returns a boolean if a field has been set.
func (o *AlertStrategyMuteTimeInterval) HasOnceMinutes() bool {
	return o != nil && o.OnceMinutes != nil
}

// SetOnceMinutes gets a reference to the given int32 and assigns it to the OnceMinutes field.
func (o *AlertStrategyMuteTimeInterval) SetOnceMinutes(v int32) {
	o.OnceMinutes = &v
}



// MarshalJSON serializes the struct using spec logic.
func (o AlertStrategyMuteTimeInterval) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Weekdays != nil {
		toSerialize["weekdays"] = o.Weekdays
	}
	if o.Times != nil {
		toSerialize["times"] = o.Times
	}
	if o.OnceMinutes != nil {
		toSerialize["onceMinutes"] = o.OnceMinutes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AlertStrategyMuteTimeInterval) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Weekdays []int32 `json:"weekdays,omitempty"`
		Times *AlertStrategyMuteTimeIntervalTimes `json:"times,omitempty"`
		OnceMinutes *int32 `json:"onceMinutes,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{ "weekdays", "times", "onceMinutes",  })
	} else {
		return err
	}

	hasInvalidField := false
	o.Weekdays = all.Weekdays
	if  all.Times != nil && all.Times.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Times = all.Times
	o.OnceMinutes = all.OnceMinutes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
