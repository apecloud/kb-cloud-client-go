// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
)

// LogOption the log type (error, running, etc.) corresponds to MonitorDefinition's logType
type LogOption struct {
	Component string `json:"component"`
	Error     bool   `json:"error"`
	Slow      bool   `json:"slow"`
	Audit     bool   `json:"audit"`
	Running   bool   `json:"running"`
	// serves the same purpose as the running log
	Pod *bool `json:"pod,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLogOption instantiates a new LogOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogOption(component string, error bool, slow bool, audit bool, running bool) *LogOption {
	this := LogOption{}
	this.Component = component
	this.Error = error
	this.Slow = slow
	this.Audit = audit
	this.Running = running
	var pod bool = false
	this.Pod = &pod
	return &this
}

// NewLogOptionWithDefaults instantiates a new LogOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogOptionWithDefaults() *LogOption {
	this := LogOption{}
	var pod bool = false
	this.Pod = &pod
	return &this
}

// GetComponent returns the Component field value.
func (o *LogOption) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *LogOption) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value.
func (o *LogOption) SetComponent(v string) {
	o.Component = v
}

// GetError returns the Error field value.
func (o *LogOption) GetError() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *LogOption) GetErrorOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value.
func (o *LogOption) SetError(v bool) {
	o.Error = v
}

// GetSlow returns the Slow field value.
func (o *LogOption) GetSlow() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Slow
}

// GetSlowOk returns a tuple with the Slow field value
// and a boolean to check if the value has been set.
func (o *LogOption) GetSlowOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slow, true
}

// SetSlow sets field value.
func (o *LogOption) SetSlow(v bool) {
	o.Slow = v
}

// GetAudit returns the Audit field value.
func (o *LogOption) GetAudit() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Audit
}

// GetAuditOk returns a tuple with the Audit field value
// and a boolean to check if the value has been set.
func (o *LogOption) GetAuditOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Audit, true
}

// SetAudit sets field value.
func (o *LogOption) SetAudit(v bool) {
	o.Audit = v
}

// GetRunning returns the Running field value.
func (o *LogOption) GetRunning() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Running
}

// GetRunningOk returns a tuple with the Running field value
// and a boolean to check if the value has been set.
func (o *LogOption) GetRunningOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Running, true
}

// SetRunning sets field value.
func (o *LogOption) SetRunning(v bool) {
	o.Running = v
}

// GetPod returns the Pod field value if set, zero value otherwise.
func (o *LogOption) GetPod() bool {
	if o == nil || o.Pod == nil {
		var ret bool
		return ret
	}
	return *o.Pod
}

// GetPodOk returns a tuple with the Pod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogOption) GetPodOk() (*bool, bool) {
	if o == nil || o.Pod == nil {
		return nil, false
	}
	return o.Pod, true
}

// HasPod returns a boolean if a field has been set.
func (o *LogOption) HasPod() bool {
	return o != nil && o.Pod != nil
}

// SetPod gets a reference to the given bool and assigns it to the Pod field.
func (o *LogOption) SetPod(v bool) {
	o.Pod = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["component"] = o.Component
	toSerialize["error"] = o.Error
	toSerialize["slow"] = o.Slow
	toSerialize["audit"] = o.Audit
	toSerialize["running"] = o.Running
	if o.Pod != nil {
		toSerialize["pod"] = o.Pod
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Component *string `json:"component"`
		Error     *bool   `json:"error"`
		Slow      *bool   `json:"slow"`
		Audit     *bool   `json:"audit"`
		Running   *bool   `json:"running"`
		Pod       *bool   `json:"pod,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Component == nil {
		return fmt.Errorf("required field component missing")
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
	if all.Running == nil {
		return fmt.Errorf("required field running missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"component", "error", "slow", "audit", "running", "pod"})
	} else {
		return err
	}
	o.Component = *all.Component
	o.Error = *all.Error
	o.Slow = *all.Slow
	o.Audit = *all.Audit
	o.Running = *all.Running
	o.Pod = all.Pod

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
