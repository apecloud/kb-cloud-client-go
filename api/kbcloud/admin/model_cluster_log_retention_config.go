// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ClusterLogRetentionConfig Cluster log retention overrides. Omitted fields inherit from the environment.
type ClusterLogRetentionConfig struct {
	PodLogRetentionDays     common.NullableInt64 `json:"podLogRetentionDays,omitempty"`
	RunningLogRetentionDays common.NullableInt64 `json:"runningLogRetentionDays,omitempty"`
	ErrorLogRetentionDays   common.NullableInt64 `json:"errorLogRetentionDays,omitempty"`
	SlowLogRetentionDays    common.NullableInt64 `json:"slowLogRetentionDays,omitempty"`
	AuditLogRetentionDays   common.NullableInt64 `json:"auditLogRetentionDays,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterLogRetentionConfig instantiates a new ClusterLogRetentionConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterLogRetentionConfig() *ClusterLogRetentionConfig {
	this := ClusterLogRetentionConfig{}
	return &this
}

// NewClusterLogRetentionConfigWithDefaults instantiates a new ClusterLogRetentionConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterLogRetentionConfigWithDefaults() *ClusterLogRetentionConfig {
	this := ClusterLogRetentionConfig{}
	return &this
}

// GetPodLogRetentionDays returns the PodLogRetentionDays field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterLogRetentionConfig) GetPodLogRetentionDays() int64 {
	if o == nil || o.PodLogRetentionDays.Get() == nil {
		var ret int64
		return ret
	}
	return *o.PodLogRetentionDays.Get()
}

// GetPodLogRetentionDaysOk returns a tuple with the PodLogRetentionDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ClusterLogRetentionConfig) GetPodLogRetentionDaysOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PodLogRetentionDays.Get(), o.PodLogRetentionDays.IsSet()
}

// HasPodLogRetentionDays returns a boolean if a field has been set.
func (o *ClusterLogRetentionConfig) HasPodLogRetentionDays() bool {
	return o != nil && o.PodLogRetentionDays.IsSet()
}

// SetPodLogRetentionDays gets a reference to the given common.NullableInt64 and assigns it to the PodLogRetentionDays field.
func (o *ClusterLogRetentionConfig) SetPodLogRetentionDays(v int64) {
	o.PodLogRetentionDays.Set(&v)
}

// SetPodLogRetentionDaysNil sets the value for PodLogRetentionDays to be an explicit nil.
func (o *ClusterLogRetentionConfig) SetPodLogRetentionDaysNil() {
	o.PodLogRetentionDays.Set(nil)
}

// UnsetPodLogRetentionDays ensures that no value is present for PodLogRetentionDays, not even an explicit nil.
func (o *ClusterLogRetentionConfig) UnsetPodLogRetentionDays() {
	o.PodLogRetentionDays.Unset()
}

// GetRunningLogRetentionDays returns the RunningLogRetentionDays field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterLogRetentionConfig) GetRunningLogRetentionDays() int64 {
	if o == nil || o.RunningLogRetentionDays.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RunningLogRetentionDays.Get()
}

// GetRunningLogRetentionDaysOk returns a tuple with the RunningLogRetentionDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ClusterLogRetentionConfig) GetRunningLogRetentionDaysOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RunningLogRetentionDays.Get(), o.RunningLogRetentionDays.IsSet()
}

// HasRunningLogRetentionDays returns a boolean if a field has been set.
func (o *ClusterLogRetentionConfig) HasRunningLogRetentionDays() bool {
	return o != nil && o.RunningLogRetentionDays.IsSet()
}

// SetRunningLogRetentionDays gets a reference to the given common.NullableInt64 and assigns it to the RunningLogRetentionDays field.
func (o *ClusterLogRetentionConfig) SetRunningLogRetentionDays(v int64) {
	o.RunningLogRetentionDays.Set(&v)
}

// SetRunningLogRetentionDaysNil sets the value for RunningLogRetentionDays to be an explicit nil.
func (o *ClusterLogRetentionConfig) SetRunningLogRetentionDaysNil() {
	o.RunningLogRetentionDays.Set(nil)
}

// UnsetRunningLogRetentionDays ensures that no value is present for RunningLogRetentionDays, not even an explicit nil.
func (o *ClusterLogRetentionConfig) UnsetRunningLogRetentionDays() {
	o.RunningLogRetentionDays.Unset()
}

// GetErrorLogRetentionDays returns the ErrorLogRetentionDays field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterLogRetentionConfig) GetErrorLogRetentionDays() int64 {
	if o == nil || o.ErrorLogRetentionDays.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ErrorLogRetentionDays.Get()
}

// GetErrorLogRetentionDaysOk returns a tuple with the ErrorLogRetentionDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ClusterLogRetentionConfig) GetErrorLogRetentionDaysOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorLogRetentionDays.Get(), o.ErrorLogRetentionDays.IsSet()
}

// HasErrorLogRetentionDays returns a boolean if a field has been set.
func (o *ClusterLogRetentionConfig) HasErrorLogRetentionDays() bool {
	return o != nil && o.ErrorLogRetentionDays.IsSet()
}

// SetErrorLogRetentionDays gets a reference to the given common.NullableInt64 and assigns it to the ErrorLogRetentionDays field.
func (o *ClusterLogRetentionConfig) SetErrorLogRetentionDays(v int64) {
	o.ErrorLogRetentionDays.Set(&v)
}

// SetErrorLogRetentionDaysNil sets the value for ErrorLogRetentionDays to be an explicit nil.
func (o *ClusterLogRetentionConfig) SetErrorLogRetentionDaysNil() {
	o.ErrorLogRetentionDays.Set(nil)
}

// UnsetErrorLogRetentionDays ensures that no value is present for ErrorLogRetentionDays, not even an explicit nil.
func (o *ClusterLogRetentionConfig) UnsetErrorLogRetentionDays() {
	o.ErrorLogRetentionDays.Unset()
}

// GetSlowLogRetentionDays returns the SlowLogRetentionDays field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterLogRetentionConfig) GetSlowLogRetentionDays() int64 {
	if o == nil || o.SlowLogRetentionDays.Get() == nil {
		var ret int64
		return ret
	}
	return *o.SlowLogRetentionDays.Get()
}

// GetSlowLogRetentionDaysOk returns a tuple with the SlowLogRetentionDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ClusterLogRetentionConfig) GetSlowLogRetentionDaysOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SlowLogRetentionDays.Get(), o.SlowLogRetentionDays.IsSet()
}

// HasSlowLogRetentionDays returns a boolean if a field has been set.
func (o *ClusterLogRetentionConfig) HasSlowLogRetentionDays() bool {
	return o != nil && o.SlowLogRetentionDays.IsSet()
}

// SetSlowLogRetentionDays gets a reference to the given common.NullableInt64 and assigns it to the SlowLogRetentionDays field.
func (o *ClusterLogRetentionConfig) SetSlowLogRetentionDays(v int64) {
	o.SlowLogRetentionDays.Set(&v)
}

// SetSlowLogRetentionDaysNil sets the value for SlowLogRetentionDays to be an explicit nil.
func (o *ClusterLogRetentionConfig) SetSlowLogRetentionDaysNil() {
	o.SlowLogRetentionDays.Set(nil)
}

// UnsetSlowLogRetentionDays ensures that no value is present for SlowLogRetentionDays, not even an explicit nil.
func (o *ClusterLogRetentionConfig) UnsetSlowLogRetentionDays() {
	o.SlowLogRetentionDays.Unset()
}

// GetAuditLogRetentionDays returns the AuditLogRetentionDays field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterLogRetentionConfig) GetAuditLogRetentionDays() int64 {
	if o == nil || o.AuditLogRetentionDays.Get() == nil {
		var ret int64
		return ret
	}
	return *o.AuditLogRetentionDays.Get()
}

// GetAuditLogRetentionDaysOk returns a tuple with the AuditLogRetentionDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ClusterLogRetentionConfig) GetAuditLogRetentionDaysOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuditLogRetentionDays.Get(), o.AuditLogRetentionDays.IsSet()
}

// HasAuditLogRetentionDays returns a boolean if a field has been set.
func (o *ClusterLogRetentionConfig) HasAuditLogRetentionDays() bool {
	return o != nil && o.AuditLogRetentionDays.IsSet()
}

// SetAuditLogRetentionDays gets a reference to the given common.NullableInt64 and assigns it to the AuditLogRetentionDays field.
func (o *ClusterLogRetentionConfig) SetAuditLogRetentionDays(v int64) {
	o.AuditLogRetentionDays.Set(&v)
}

// SetAuditLogRetentionDaysNil sets the value for AuditLogRetentionDays to be an explicit nil.
func (o *ClusterLogRetentionConfig) SetAuditLogRetentionDaysNil() {
	o.AuditLogRetentionDays.Set(nil)
}

// UnsetAuditLogRetentionDays ensures that no value is present for AuditLogRetentionDays, not even an explicit nil.
func (o *ClusterLogRetentionConfig) UnsetAuditLogRetentionDays() {
	o.AuditLogRetentionDays.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterLogRetentionConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.PodLogRetentionDays.IsSet() {
		toSerialize["podLogRetentionDays"] = o.PodLogRetentionDays.Get()
	}
	if o.RunningLogRetentionDays.IsSet() {
		toSerialize["runningLogRetentionDays"] = o.RunningLogRetentionDays.Get()
	}
	if o.ErrorLogRetentionDays.IsSet() {
		toSerialize["errorLogRetentionDays"] = o.ErrorLogRetentionDays.Get()
	}
	if o.SlowLogRetentionDays.IsSet() {
		toSerialize["slowLogRetentionDays"] = o.SlowLogRetentionDays.Get()
	}
	if o.AuditLogRetentionDays.IsSet() {
		toSerialize["auditLogRetentionDays"] = o.AuditLogRetentionDays.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterLogRetentionConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		PodLogRetentionDays     common.NullableInt64 `json:"podLogRetentionDays,omitempty"`
		RunningLogRetentionDays common.NullableInt64 `json:"runningLogRetentionDays,omitempty"`
		ErrorLogRetentionDays   common.NullableInt64 `json:"errorLogRetentionDays,omitempty"`
		SlowLogRetentionDays    common.NullableInt64 `json:"slowLogRetentionDays,omitempty"`
		AuditLogRetentionDays   common.NullableInt64 `json:"auditLogRetentionDays,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"podLogRetentionDays", "runningLogRetentionDays", "errorLogRetentionDays", "slowLogRetentionDays", "auditLogRetentionDays"})
	} else {
		return err
	}
	o.PodLogRetentionDays = all.PodLogRetentionDays
	o.RunningLogRetentionDays = all.RunningLogRetentionDays
	o.ErrorLogRetentionDays = all.ErrorLogRetentionDays
	o.SlowLogRetentionDays = all.SlowLogRetentionDays
	o.AuditLogRetentionDays = all.AuditLogRetentionDays

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
