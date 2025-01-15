// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// OpsVScale OpsVScale is the payload to vertically scale a KubeBlocks cluster
type OpsVScale struct {
	// component type
	Component string `json:"component"`
	// number of cpu
	Cpu *string `json:"cpu,omitempty"`
	// memory size
	Memory *string `json:"memory,omitempty"`
	// class code of the cluster
	ClassCode *string `json:"classCode,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOpsVScale instantiates a new OpsVScale object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpsVScale(component string) *OpsVScale {
	this := OpsVScale{}
	this.Component = component
	return &this
}

// NewOpsVScaleWithDefaults instantiates a new OpsVScale object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpsVScaleWithDefaults() *OpsVScale {
	this := OpsVScale{}
	return &this
}

// GetComponent returns the Component field value.
func (o *OpsVScale) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *OpsVScale) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value.
func (o *OpsVScale) SetComponent(v string) {
	o.Component = v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *OpsVScale) GetCpu() string {
	if o == nil || o.Cpu == nil {
		var ret string
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsVScale) GetCpuOk() (*string, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *OpsVScale) HasCpu() bool {
	return o != nil && o.Cpu != nil
}

// SetCpu gets a reference to the given string and assigns it to the Cpu field.
func (o *OpsVScale) SetCpu(v string) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *OpsVScale) GetMemory() string {
	if o == nil || o.Memory == nil {
		var ret string
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsVScale) GetMemoryOk() (*string, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *OpsVScale) HasMemory() bool {
	return o != nil && o.Memory != nil
}

// SetMemory gets a reference to the given string and assigns it to the Memory field.
func (o *OpsVScale) SetMemory(v string) {
	o.Memory = &v
}

// GetClassCode returns the ClassCode field value if set, zero value otherwise.
func (o *OpsVScale) GetClassCode() string {
	if o == nil || o.ClassCode == nil {
		var ret string
		return ret
	}
	return *o.ClassCode
}

// GetClassCodeOk returns a tuple with the ClassCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsVScale) GetClassCodeOk() (*string, bool) {
	if o == nil || o.ClassCode == nil {
		return nil, false
	}
	return o.ClassCode, true
}

// HasClassCode returns a boolean if a field has been set.
func (o *OpsVScale) HasClassCode() bool {
	return o != nil && o.ClassCode != nil
}

// SetClassCode gets a reference to the given string and assigns it to the ClassCode field.
func (o *OpsVScale) SetClassCode(v string) {
	o.ClassCode = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpsVScale) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["component"] = o.Component
	if o.Cpu != nil {
		toSerialize["cpu"] = o.Cpu
	}
	if o.Memory != nil {
		toSerialize["memory"] = o.Memory
	}
	if o.ClassCode != nil {
		toSerialize["classCode"] = o.ClassCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpsVScale) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Component *string `json:"component"`
		Cpu       *string `json:"cpu,omitempty"`
		Memory    *string `json:"memory,omitempty"`
		ClassCode *string `json:"classCode,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Component == nil {
		return fmt.Errorf("required field component missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"component", "cpu", "memory", "classCode"})
	} else {
		return err
	}
	o.Component = *all.Component
	o.Cpu = all.Cpu
	o.Memory = all.Memory
	o.ClassCode = all.ClassCode

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
