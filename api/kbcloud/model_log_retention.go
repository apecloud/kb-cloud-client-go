// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// LogRetention Log retention in days.
type LogRetention struct {
	Pod     int64 `json:"pod"`
	Running int64 `json:"running"`
	Error   int64 `json:"error"`
	Slow    int64 `json:"slow"`
	Audit   int64 `json:"audit"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLogRetention instantiates a new LogRetention object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogRetention(pod int64, running int64, error int64, slow int64, audit int64) *LogRetention {
	this := LogRetention{}
	this.Pod = pod
	this.Running = running
	this.Error = error
	this.Slow = slow
	this.Audit = audit
	return &this
}

// NewLogRetentionWithDefaults instantiates a new LogRetention object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogRetentionWithDefaults() *LogRetention {
	this := LogRetention{}
	return &this
}

// GetPod returns the Pod field value.
func (o *LogRetention) GetPod() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Pod
}

// GetPodOk returns a tuple with the Pod field value
// and a boolean to check if the value has been set.
func (o *LogRetention) GetPodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pod, true
}

// SetPod sets field value.
func (o *LogRetention) SetPod(v int64) {
	o.Pod = v
}

// GetRunning returns the Running field value.
func (o *LogRetention) GetRunning() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Running
}

// GetRunningOk returns a tuple with the Running field value
// and a boolean to check if the value has been set.
func (o *LogRetention) GetRunningOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Running, true
}

// SetRunning sets field value.
func (o *LogRetention) SetRunning(v int64) {
	o.Running = v
}

// GetError returns the Error field value.
func (o *LogRetention) GetError() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *LogRetention) GetErrorOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value.
func (o *LogRetention) SetError(v int64) {
	o.Error = v
}

// GetSlow returns the Slow field value.
func (o *LogRetention) GetSlow() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Slow
}

// GetSlowOk returns a tuple with the Slow field value
// and a boolean to check if the value has been set.
func (o *LogRetention) GetSlowOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slow, true
}

// SetSlow sets field value.
func (o *LogRetention) SetSlow(v int64) {
	o.Slow = v
}

// GetAudit returns the Audit field value.
func (o *LogRetention) GetAudit() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Audit
}

// GetAuditOk returns a tuple with the Audit field value
// and a boolean to check if the value has been set.
func (o *LogRetention) GetAuditOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Audit, true
}

// SetAudit sets field value.
func (o *LogRetention) SetAudit(v int64) {
	o.Audit = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogRetention) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["pod"] = o.Pod
	toSerialize["running"] = o.Running
	toSerialize["error"] = o.Error
	toSerialize["slow"] = o.Slow
	toSerialize["audit"] = o.Audit

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogRetention) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Pod     *int64 `json:"pod"`
		Running *int64 `json:"running"`
		Error   *int64 `json:"error"`
		Slow    *int64 `json:"slow"`
		Audit   *int64 `json:"audit"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Pod == nil {
		return fmt.Errorf("required field pod missing")
	}
	if all.Running == nil {
		return fmt.Errorf("required field running missing")
	}
	if all.Error == nil {
		return fmt.Errorf("required field error missing")
	}
	if all.Slow == nil {
		return fmt.Errorf("required field slow missing")
	}
	if all.Audit == nil {
		return fmt.Errorf("required field audit missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"pod", "running", "error", "slow", "audit"})
	} else {
		return err
	}
	o.Pod = *all.Pod
	o.Running = *all.Running
	o.Error = *all.Error
	o.Slow = *all.Slow
	o.Audit = *all.Audit

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
