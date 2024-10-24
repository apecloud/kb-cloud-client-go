// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type Vscale struct {
	// the cpu resource of vertical scale
	Cpu string `json:"cpu"`
	// the memory resource of vertical scale
	Memory string `json:"memory"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewVscale instantiates a new Vscale object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewVscale(cpu string, memory string) *Vscale {
	this := Vscale{}
	this.Cpu = cpu
	this.Memory = memory
	return &this
}

// NewVscaleWithDefaults instantiates a new Vscale object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewVscaleWithDefaults() *Vscale {
	this := Vscale{}
	return &this
}

// GetCpu returns the Cpu field value.
func (o *Vscale) GetCpu() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value
// and a boolean to check if the value has been set.
func (o *Vscale) GetCpuOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cpu, true
}

// SetCpu sets field value.
func (o *Vscale) SetCpu(v string) {
	o.Cpu = v
}

// GetMemory returns the Memory field value.
func (o *Vscale) GetMemory() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value
// and a boolean to check if the value has been set.
func (o *Vscale) GetMemoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Memory, true
}

// SetMemory sets field value.
func (o *Vscale) SetMemory(v string) {
	o.Memory = v
}

// MarshalJSON serializes the struct using spec logic.
func (o Vscale) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["cpu"] = o.Cpu
	toSerialize["memory"] = o.Memory

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Vscale) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Cpu    *string `json:"cpu"`
		Memory *string `json:"memory"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Cpu == nil {
		return fmt.Errorf("required field cpu missing")
	}
	if all.Memory == nil {
		return fmt.Errorf("required field memory missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"cpu", "memory"})
	} else {
		return err
	}
	o.Cpu = *all.Cpu
	o.Memory = *all.Memory

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
