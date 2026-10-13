// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EnvironmentLogRetentionConfig Environment log retention days.
type EnvironmentLogRetentionConfig struct {
	PodLogRetentionDays     int64 `json:"podLogRetentionDays"`
	RunningLogRetentionDays int64 `json:"runningLogRetentionDays"`
	ErrorLogRetentionDays   int64 `json:"errorLogRetentionDays"`
	SlowLogRetentionDays    int64 `json:"slowLogRetentionDays"`
	AuditLogRetentionDays   int64 `json:"auditLogRetentionDays"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironmentLogRetentionConfig instantiates a new EnvironmentLogRetentionConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironmentLogRetentionConfig(podLogRetentionDays int64, runningLogRetentionDays int64, errorLogRetentionDays int64, slowLogRetentionDays int64, auditLogRetentionDays int64) *EnvironmentLogRetentionConfig {
	this := EnvironmentLogRetentionConfig{}
	this.PodLogRetentionDays = podLogRetentionDays
	this.RunningLogRetentionDays = runningLogRetentionDays
	this.ErrorLogRetentionDays = errorLogRetentionDays
	this.SlowLogRetentionDays = slowLogRetentionDays
	this.AuditLogRetentionDays = auditLogRetentionDays
	return &this
}

// NewEnvironmentLogRetentionConfigWithDefaults instantiates a new EnvironmentLogRetentionConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironmentLogRetentionConfigWithDefaults() *EnvironmentLogRetentionConfig {
	this := EnvironmentLogRetentionConfig{}
	return &this
}

// GetPodLogRetentionDays returns the PodLogRetentionDays field value.
func (o *EnvironmentLogRetentionConfig) GetPodLogRetentionDays() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.PodLogRetentionDays
}

// GetPodLogRetentionDaysOk returns a tuple with the PodLogRetentionDays field value
// and a boolean to check if the value has been set.
func (o *EnvironmentLogRetentionConfig) GetPodLogRetentionDaysOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PodLogRetentionDays, true
}

// SetPodLogRetentionDays sets field value.
func (o *EnvironmentLogRetentionConfig) SetPodLogRetentionDays(v int64) {
	o.PodLogRetentionDays = v
}

// GetRunningLogRetentionDays returns the RunningLogRetentionDays field value.
func (o *EnvironmentLogRetentionConfig) GetRunningLogRetentionDays() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.RunningLogRetentionDays
}

// GetRunningLogRetentionDaysOk returns a tuple with the RunningLogRetentionDays field value
// and a boolean to check if the value has been set.
func (o *EnvironmentLogRetentionConfig) GetRunningLogRetentionDaysOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunningLogRetentionDays, true
}

// SetRunningLogRetentionDays sets field value.
func (o *EnvironmentLogRetentionConfig) SetRunningLogRetentionDays(v int64) {
	o.RunningLogRetentionDays = v
}

// GetErrorLogRetentionDays returns the ErrorLogRetentionDays field value.
func (o *EnvironmentLogRetentionConfig) GetErrorLogRetentionDays() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ErrorLogRetentionDays
}

// GetErrorLogRetentionDaysOk returns a tuple with the ErrorLogRetentionDays field value
// and a boolean to check if the value has been set.
func (o *EnvironmentLogRetentionConfig) GetErrorLogRetentionDaysOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorLogRetentionDays, true
}

// SetErrorLogRetentionDays sets field value.
func (o *EnvironmentLogRetentionConfig) SetErrorLogRetentionDays(v int64) {
	o.ErrorLogRetentionDays = v
}

// GetSlowLogRetentionDays returns the SlowLogRetentionDays field value.
func (o *EnvironmentLogRetentionConfig) GetSlowLogRetentionDays() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.SlowLogRetentionDays
}

// GetSlowLogRetentionDaysOk returns a tuple with the SlowLogRetentionDays field value
// and a boolean to check if the value has been set.
func (o *EnvironmentLogRetentionConfig) GetSlowLogRetentionDaysOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SlowLogRetentionDays, true
}

// SetSlowLogRetentionDays sets field value.
func (o *EnvironmentLogRetentionConfig) SetSlowLogRetentionDays(v int64) {
	o.SlowLogRetentionDays = v
}

// GetAuditLogRetentionDays returns the AuditLogRetentionDays field value.
func (o *EnvironmentLogRetentionConfig) GetAuditLogRetentionDays() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.AuditLogRetentionDays
}

// GetAuditLogRetentionDaysOk returns a tuple with the AuditLogRetentionDays field value
// and a boolean to check if the value has been set.
func (o *EnvironmentLogRetentionConfig) GetAuditLogRetentionDaysOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuditLogRetentionDays, true
}

// SetAuditLogRetentionDays sets field value.
func (o *EnvironmentLogRetentionConfig) SetAuditLogRetentionDays(v int64) {
	o.AuditLogRetentionDays = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EnvironmentLogRetentionConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["podLogRetentionDays"] = o.PodLogRetentionDays
	toSerialize["runningLogRetentionDays"] = o.RunningLogRetentionDays
	toSerialize["errorLogRetentionDays"] = o.ErrorLogRetentionDays
	toSerialize["slowLogRetentionDays"] = o.SlowLogRetentionDays
	toSerialize["auditLogRetentionDays"] = o.AuditLogRetentionDays

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EnvironmentLogRetentionConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		PodLogRetentionDays     *int64 `json:"podLogRetentionDays"`
		RunningLogRetentionDays *int64 `json:"runningLogRetentionDays"`
		ErrorLogRetentionDays   *int64 `json:"errorLogRetentionDays"`
		SlowLogRetentionDays    *int64 `json:"slowLogRetentionDays"`
		AuditLogRetentionDays   *int64 `json:"auditLogRetentionDays"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.PodLogRetentionDays == nil {
		return fmt.Errorf("required field podLogRetentionDays missing")
	}
	if all.RunningLogRetentionDays == nil {
		return fmt.Errorf("required field runningLogRetentionDays missing")
	}
	if all.ErrorLogRetentionDays == nil {
		return fmt.Errorf("required field errorLogRetentionDays missing")
	}
	if all.SlowLogRetentionDays == nil {
		return fmt.Errorf("required field slowLogRetentionDays missing")
	}
	if all.AuditLogRetentionDays == nil {
		return fmt.Errorf("required field auditLogRetentionDays missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"podLogRetentionDays", "runningLogRetentionDays", "errorLogRetentionDays", "slowLogRetentionDays", "auditLogRetentionDays"})
	} else {
		return err
	}
	o.PodLogRetentionDays = *all.PodLogRetentionDays
	o.RunningLogRetentionDays = *all.RunningLogRetentionDays
	o.ErrorLogRetentionDays = *all.ErrorLogRetentionDays
	o.SlowLogRetentionDays = *all.SlowLogRetentionDays
	o.AuditLogRetentionDays = *all.AuditLogRetentionDays

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
