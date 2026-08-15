// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// KoordinatorReservationResources Resource requests supported by KBE for Koordinator Reservations. Values are Kubernetes quantities and responses use Kubernetes canonical quantity strings.
type KoordinatorReservationResources struct {
	Cpu    *string `json:"cpu,omitempty"`
	Memory *string `json:"memory,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewKoordinatorReservationResources instantiates a new KoordinatorReservationResources object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewKoordinatorReservationResources() *KoordinatorReservationResources {
	this := KoordinatorReservationResources{}
	return &this
}

// NewKoordinatorReservationResourcesWithDefaults instantiates a new KoordinatorReservationResources object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewKoordinatorReservationResourcesWithDefaults() *KoordinatorReservationResources {
	this := KoordinatorReservationResources{}
	return &this
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *KoordinatorReservationResources) GetCpu() string {
	if o == nil || o.Cpu == nil {
		var ret string
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KoordinatorReservationResources) GetCpuOk() (*string, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *KoordinatorReservationResources) HasCpu() bool {
	return o != nil && o.Cpu != nil
}

// SetCpu gets a reference to the given string and assigns it to the Cpu field.
func (o *KoordinatorReservationResources) SetCpu(v string) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *KoordinatorReservationResources) GetMemory() string {
	if o == nil || o.Memory == nil {
		var ret string
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KoordinatorReservationResources) GetMemoryOk() (*string, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *KoordinatorReservationResources) HasMemory() bool {
	return o != nil && o.Memory != nil
}

// SetMemory gets a reference to the given string and assigns it to the Memory field.
func (o *KoordinatorReservationResources) SetMemory(v string) {
	o.Memory = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o KoordinatorReservationResources) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Cpu != nil {
		toSerialize["cpu"] = o.Cpu
	}
	if o.Memory != nil {
		toSerialize["memory"] = o.Memory
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *KoordinatorReservationResources) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Cpu    *string `json:"cpu,omitempty"`
		Memory *string `json:"memory,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	o.Cpu = all.Cpu
	o.Memory = all.Memory

	return nil
}
