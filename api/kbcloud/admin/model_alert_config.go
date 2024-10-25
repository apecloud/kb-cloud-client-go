// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// NODESCRIPTION AlertConfig
type AlertConfig struct {
	// Time zone offset in seconds, e.g. offset of UTC+08:00 is +8 * 60 * 60
	TimeZoneOffset int32 `json:"timeZoneOffset"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAlertConfig instantiates a new AlertConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAlertConfig(timeZoneOffset int32) *AlertConfig {
	this := AlertConfig{}
	this.TimeZoneOffset = timeZoneOffset
	return &this
}

// NewAlertConfigWithDefaults instantiates a new AlertConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAlertConfigWithDefaults() *AlertConfig {
	this := AlertConfig{}
	return &this
}

// GetTimeZoneOffset returns the TimeZoneOffset field value.
func (o *AlertConfig) GetTimeZoneOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.TimeZoneOffset
}

// GetTimeZoneOffsetOk returns a tuple with the TimeZoneOffset field value
// and a boolean to check if the value has been set.
func (o *AlertConfig) GetTimeZoneOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeZoneOffset, true
}

// SetTimeZoneOffset sets field value.
func (o *AlertConfig) SetTimeZoneOffset(v int32) {
	o.TimeZoneOffset = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AlertConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["timeZoneOffset"] = o.TimeZoneOffset

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AlertConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TimeZoneOffset *int32 `json:"timeZoneOffset"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.TimeZoneOffset == nil {
		return fmt.Errorf("required field timeZoneOffset missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"timeZoneOffset"})
	} else {
		return err
	}
	o.TimeZoneOffset = *all.TimeZoneOffset

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
