// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type EnvironmentLogRetentionConfig struct {
	// Default retention days.
	DefaultRetentionDays int64 `json:"defaultRetentionDays"`
	// Retention days by log type. Omitted types use the default.
	RetentionDaysByType map[string]int64 `json:"retentionDaysByType"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironmentLogRetentionConfig instantiates a new EnvironmentLogRetentionConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironmentLogRetentionConfig(defaultRetentionDays int64, retentionDaysByType map[string]int64) *EnvironmentLogRetentionConfig {
	this := EnvironmentLogRetentionConfig{}
	this.DefaultRetentionDays = defaultRetentionDays
	this.RetentionDaysByType = retentionDaysByType
	return &this
}

// NewEnvironmentLogRetentionConfigWithDefaults instantiates a new EnvironmentLogRetentionConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironmentLogRetentionConfigWithDefaults() *EnvironmentLogRetentionConfig {
	this := EnvironmentLogRetentionConfig{}
	return &this
}

// GetDefaultRetentionDays returns the DefaultRetentionDays field value.
func (o *EnvironmentLogRetentionConfig) GetDefaultRetentionDays() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.DefaultRetentionDays
}

// GetDefaultRetentionDaysOk returns a tuple with the DefaultRetentionDays field value
// and a boolean to check if the value has been set.
func (o *EnvironmentLogRetentionConfig) GetDefaultRetentionDaysOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultRetentionDays, true
}

// SetDefaultRetentionDays sets field value.
func (o *EnvironmentLogRetentionConfig) SetDefaultRetentionDays(v int64) {
	o.DefaultRetentionDays = v
}

// GetRetentionDaysByType returns the RetentionDaysByType field value.
func (o *EnvironmentLogRetentionConfig) GetRetentionDaysByType() map[string]int64 {
	if o == nil {
		var ret map[string]int64
		return ret
	}
	return o.RetentionDaysByType
}

// GetRetentionDaysByTypeOk returns a tuple with the RetentionDaysByType field value
// and a boolean to check if the value has been set.
func (o *EnvironmentLogRetentionConfig) GetRetentionDaysByTypeOk() (*map[string]int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RetentionDaysByType, true
}

// SetRetentionDaysByType sets field value.
func (o *EnvironmentLogRetentionConfig) SetRetentionDaysByType(v map[string]int64) {
	o.RetentionDaysByType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EnvironmentLogRetentionConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["defaultRetentionDays"] = o.DefaultRetentionDays
	toSerialize["retentionDaysByType"] = o.RetentionDaysByType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EnvironmentLogRetentionConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DefaultRetentionDays *int64            `json:"defaultRetentionDays"`
		RetentionDaysByType  *map[string]int64 `json:"retentionDaysByType"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.DefaultRetentionDays == nil {
		return fmt.Errorf("required field defaultRetentionDays missing")
	}
	if all.RetentionDaysByType == nil {
		return fmt.Errorf("required field retentionDaysByType missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"defaultRetentionDays", "retentionDaysByType"})
	} else {
		return err
	}
	o.DefaultRetentionDays = *all.DefaultRetentionDays
	o.RetentionDaysByType = *all.RetentionDaysByType

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
