// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type LogRetentionPolicy struct {
	// Environment default retention days.
	DefaultRetentionDays *int64 `json:"defaultRetentionDays,omitempty"`
	// Configured retention days by log type.
	RetentionDaysByType map[string]int64 `json:"retentionDaysByType"`
	// Effective retention days and sources.
	EffectivePolicies []EffectiveLogRetentionPolicy `json:"effectivePolicies,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLogRetentionPolicy instantiates a new LogRetentionPolicy object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogRetentionPolicy(retentionDaysByType map[string]int64) *LogRetentionPolicy {
	this := LogRetentionPolicy{}
	this.RetentionDaysByType = retentionDaysByType
	return &this
}

// NewLogRetentionPolicyWithDefaults instantiates a new LogRetentionPolicy object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogRetentionPolicyWithDefaults() *LogRetentionPolicy {
	this := LogRetentionPolicy{}
	return &this
}

// GetDefaultRetentionDays returns the DefaultRetentionDays field value if set, zero value otherwise.
func (o *LogRetentionPolicy) GetDefaultRetentionDays() int64 {
	if o == nil || o.DefaultRetentionDays == nil {
		var ret int64
		return ret
	}
	return *o.DefaultRetentionDays
}

// GetDefaultRetentionDaysOk returns a tuple with the DefaultRetentionDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogRetentionPolicy) GetDefaultRetentionDaysOk() (*int64, bool) {
	if o == nil || o.DefaultRetentionDays == nil {
		return nil, false
	}
	return o.DefaultRetentionDays, true
}

// HasDefaultRetentionDays returns a boolean if a field has been set.
func (o *LogRetentionPolicy) HasDefaultRetentionDays() bool {
	return o != nil && o.DefaultRetentionDays != nil
}

// SetDefaultRetentionDays gets a reference to the given int64 and assigns it to the DefaultRetentionDays field.
func (o *LogRetentionPolicy) SetDefaultRetentionDays(v int64) {
	o.DefaultRetentionDays = &v
}

// GetRetentionDaysByType returns the RetentionDaysByType field value.
func (o *LogRetentionPolicy) GetRetentionDaysByType() map[string]int64 {
	if o == nil {
		var ret map[string]int64
		return ret
	}
	return o.RetentionDaysByType
}

// GetRetentionDaysByTypeOk returns a tuple with the RetentionDaysByType field value
// and a boolean to check if the value has been set.
func (o *LogRetentionPolicy) GetRetentionDaysByTypeOk() (*map[string]int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RetentionDaysByType, true
}

// SetRetentionDaysByType sets field value.
func (o *LogRetentionPolicy) SetRetentionDaysByType(v map[string]int64) {
	o.RetentionDaysByType = v
}

// GetEffectivePolicies returns the EffectivePolicies field value if set, zero value otherwise.
func (o *LogRetentionPolicy) GetEffectivePolicies() []EffectiveLogRetentionPolicy {
	if o == nil || o.EffectivePolicies == nil {
		var ret []EffectiveLogRetentionPolicy
		return ret
	}
	return o.EffectivePolicies
}

// GetEffectivePoliciesOk returns a tuple with the EffectivePolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogRetentionPolicy) GetEffectivePoliciesOk() (*[]EffectiveLogRetentionPolicy, bool) {
	if o == nil || o.EffectivePolicies == nil {
		return nil, false
	}
	return &o.EffectivePolicies, true
}

// HasEffectivePolicies returns a boolean if a field has been set.
func (o *LogRetentionPolicy) HasEffectivePolicies() bool {
	return o != nil && o.EffectivePolicies != nil
}

// SetEffectivePolicies gets a reference to the given []EffectiveLogRetentionPolicy and assigns it to the EffectivePolicies field.
func (o *LogRetentionPolicy) SetEffectivePolicies(v []EffectiveLogRetentionPolicy) {
	o.EffectivePolicies = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogRetentionPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.DefaultRetentionDays != nil {
		toSerialize["defaultRetentionDays"] = o.DefaultRetentionDays
	}
	toSerialize["retentionDaysByType"] = o.RetentionDaysByType
	if o.EffectivePolicies != nil {
		toSerialize["effectivePolicies"] = o.EffectivePolicies
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogRetentionPolicy) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DefaultRetentionDays *int64                        `json:"defaultRetentionDays,omitempty"`
		RetentionDaysByType  *map[string]int64             `json:"retentionDaysByType"`
		EffectivePolicies    []EffectiveLogRetentionPolicy `json:"effectivePolicies,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.RetentionDaysByType == nil {
		return fmt.Errorf("required field retentionDaysByType missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"defaultRetentionDays", "retentionDaysByType", "effectivePolicies"})
	} else {
		return err
	}
	o.DefaultRetentionDays = all.DefaultRetentionDays
	o.RetentionDaysByType = *all.RetentionDaysByType
	o.EffectivePolicies = all.EffectivePolicies

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
