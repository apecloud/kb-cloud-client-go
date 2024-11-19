// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type AlertStrategyMuteTimeIntervalTimes struct {
	// Mute start time, e.g. '17:00', should be in UTC time.
	StartTime *string `json:"startTime,omitempty"`
	// Mute end time, e.g. '24:00', should be in UTC time.
	EndTime *string `json:"endTime,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAlertStrategyMuteTimeIntervalTimes instantiates a new AlertStrategyMuteTimeIntervalTimes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAlertStrategyMuteTimeIntervalTimes() *AlertStrategyMuteTimeIntervalTimes {
	this := AlertStrategyMuteTimeIntervalTimes{}
	return &this
}

// NewAlertStrategyMuteTimeIntervalTimesWithDefaults instantiates a new AlertStrategyMuteTimeIntervalTimes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAlertStrategyMuteTimeIntervalTimesWithDefaults() *AlertStrategyMuteTimeIntervalTimes {
	this := AlertStrategyMuteTimeIntervalTimes{}
	return &this
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *AlertStrategyMuteTimeIntervalTimes) GetStartTime() string {
	if o == nil || o.StartTime == nil {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertStrategyMuteTimeIntervalTimes) GetStartTimeOk() (*string, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *AlertStrategyMuteTimeIntervalTimes) HasStartTime() bool {
	return o != nil && o.StartTime != nil
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *AlertStrategyMuteTimeIntervalTimes) SetStartTime(v string) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *AlertStrategyMuteTimeIntervalTimes) GetEndTime() string {
	if o == nil || o.EndTime == nil {
		var ret string
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertStrategyMuteTimeIntervalTimes) GetEndTimeOk() (*string, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *AlertStrategyMuteTimeIntervalTimes) HasEndTime() bool {
	return o != nil && o.EndTime != nil
}

// SetEndTime gets a reference to the given string and assigns it to the EndTime field.
func (o *AlertStrategyMuteTimeIntervalTimes) SetEndTime(v string) {
	o.EndTime = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AlertStrategyMuteTimeIntervalTimes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.StartTime != nil {
		toSerialize["startTime"] = o.StartTime
	}
	if o.EndTime != nil {
		toSerialize["endTime"] = o.EndTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AlertStrategyMuteTimeIntervalTimes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		StartTime *string `json:"startTime,omitempty"`
		EndTime   *string `json:"endTime,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"startTime", "endTime"})
	} else {
		return err
	}
	o.StartTime = all.StartTime
	o.EndTime = all.EndTime

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
