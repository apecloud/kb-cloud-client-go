// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// SeriesItem A data point in a series for meta data, including a count and a timestamp.
// NODESCRIPTION SeriesItem
//
// Deprecated: This model is deprecated.
type SeriesItem struct {
	// The value of the series data point.
	Value float64 `json:"value"`
	// The timestamp representing the time of the data point, unit is seconds.
	Timestamp int64 `json:"timestamp"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSeriesItem instantiates a new SeriesItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSeriesItem(value float64, timestamp int64) *SeriesItem {
	this := SeriesItem{}
	this.Value = value
	this.Timestamp = timestamp
	return &this
}

// NewSeriesItemWithDefaults instantiates a new SeriesItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSeriesItemWithDefaults() *SeriesItem {
	this := SeriesItem{}
	return &this
}

// GetValue returns the Value field value.
func (o *SeriesItem) GetValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *SeriesItem) GetValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *SeriesItem) SetValue(v float64) {
	o.Value = v
}

// GetTimestamp returns the Timestamp field value.
func (o *SeriesItem) GetTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *SeriesItem) GetTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value.
func (o *SeriesItem) SetTimestamp(v int64) {
	o.Timestamp = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SeriesItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["value"] = o.Value
	toSerialize["timestamp"] = o.Timestamp

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SeriesItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Value     *float64 `json:"value"`
		Timestamp *int64   `json:"timestamp"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	if all.Timestamp == nil {
		return fmt.Errorf("required field timestamp missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"value", "timestamp"})
	} else {
		return err
	}
	o.Value = *all.Value
	o.Timestamp = *all.Timestamp

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
