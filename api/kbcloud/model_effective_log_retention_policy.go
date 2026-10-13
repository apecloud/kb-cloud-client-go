// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type EffectiveLogRetentionPolicy struct {
	LogType       *string `json:"logType,omitempty"`
	RetentionDays *int64  `json:"retentionDays,omitempty"`
	// Policy source: cluster, environment or environmentDefault.
	Source *string `json:"source,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEffectiveLogRetentionPolicy instantiates a new EffectiveLogRetentionPolicy object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEffectiveLogRetentionPolicy() *EffectiveLogRetentionPolicy {
	this := EffectiveLogRetentionPolicy{}
	return &this
}

// NewEffectiveLogRetentionPolicyWithDefaults instantiates a new EffectiveLogRetentionPolicy object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEffectiveLogRetentionPolicyWithDefaults() *EffectiveLogRetentionPolicy {
	this := EffectiveLogRetentionPolicy{}
	return &this
}

// GetLogType returns the LogType field value if set, zero value otherwise.
func (o *EffectiveLogRetentionPolicy) GetLogType() string {
	if o == nil || o.LogType == nil {
		var ret string
		return ret
	}
	return *o.LogType
}

// GetLogTypeOk returns a tuple with the LogType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EffectiveLogRetentionPolicy) GetLogTypeOk() (*string, bool) {
	if o == nil || o.LogType == nil {
		return nil, false
	}
	return o.LogType, true
}

// HasLogType returns a boolean if a field has been set.
func (o *EffectiveLogRetentionPolicy) HasLogType() bool {
	return o != nil && o.LogType != nil
}

// SetLogType gets a reference to the given string and assigns it to the LogType field.
func (o *EffectiveLogRetentionPolicy) SetLogType(v string) {
	o.LogType = &v
}

// GetRetentionDays returns the RetentionDays field value if set, zero value otherwise.
func (o *EffectiveLogRetentionPolicy) GetRetentionDays() int64 {
	if o == nil || o.RetentionDays == nil {
		var ret int64
		return ret
	}
	return *o.RetentionDays
}

// GetRetentionDaysOk returns a tuple with the RetentionDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EffectiveLogRetentionPolicy) GetRetentionDaysOk() (*int64, bool) {
	if o == nil || o.RetentionDays == nil {
		return nil, false
	}
	return o.RetentionDays, true
}

// HasRetentionDays returns a boolean if a field has been set.
func (o *EffectiveLogRetentionPolicy) HasRetentionDays() bool {
	return o != nil && o.RetentionDays != nil
}

// SetRetentionDays gets a reference to the given int64 and assigns it to the RetentionDays field.
func (o *EffectiveLogRetentionPolicy) SetRetentionDays(v int64) {
	o.RetentionDays = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *EffectiveLogRetentionPolicy) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EffectiveLogRetentionPolicy) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *EffectiveLogRetentionPolicy) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *EffectiveLogRetentionPolicy) SetSource(v string) {
	o.Source = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EffectiveLogRetentionPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.LogType != nil {
		toSerialize["logType"] = o.LogType
	}
	if o.RetentionDays != nil {
		toSerialize["retentionDays"] = o.RetentionDays
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EffectiveLogRetentionPolicy) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		LogType       *string `json:"logType,omitempty"`
		RetentionDays *int64  `json:"retentionDays,omitempty"`
		Source        *string `json:"source,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"logType", "retentionDays", "source"})
	} else {
		return err
	}
	o.LogType = all.LogType
	o.RetentionDays = all.RetentionDays
	o.Source = all.Source

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
