// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package kbcloud

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)


 
type CPU struct {
	CpuCapacity *string `json:"cpu_capacity,omitempty"`
	CpuCapacityMax *string `json:"cpu_capacity_max,omitempty"`
	CpuAssigned *string `json:"cpu_assigned,omitempty"`
	CpuAssignedMax *string `json:"cpu_assigned_max,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewCPU instantiates a new CPU object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCPU() *CPU {
	this := CPU{}
	return &this
}

// NewCPUWithDefaults instantiates a new CPU object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCPUWithDefaults() *CPU {
	this := CPU{}
	return &this
}
// GetCpuCapacity returns the CpuCapacity field value if set, zero value otherwise.
func (o *CPU) GetCpuCapacity() string {
	if o == nil || o.CpuCapacity == nil {
		var ret string
		return ret
	}
	return *o.CpuCapacity
}

// GetCpuCapacityOk returns a tuple with the CpuCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CPU) GetCpuCapacityOk() (*string, bool) {
	if o == nil || o.CpuCapacity == nil {
		return nil, false
	}
	return o.CpuCapacity, true
}

// HasCpuCapacity returns a boolean if a field has been set.
func (o *CPU) HasCpuCapacity() bool {
	return o != nil && o.CpuCapacity != nil
}

// SetCpuCapacity gets a reference to the given string and assigns it to the CpuCapacity field.
func (o *CPU) SetCpuCapacity(v string) {
	o.CpuCapacity = &v
}


// GetCpuCapacityMax returns the CpuCapacityMax field value if set, zero value otherwise.
func (o *CPU) GetCpuCapacityMax() string {
	if o == nil || o.CpuCapacityMax == nil {
		var ret string
		return ret
	}
	return *o.CpuCapacityMax
}

// GetCpuCapacityMaxOk returns a tuple with the CpuCapacityMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CPU) GetCpuCapacityMaxOk() (*string, bool) {
	if o == nil || o.CpuCapacityMax == nil {
		return nil, false
	}
	return o.CpuCapacityMax, true
}

// HasCpuCapacityMax returns a boolean if a field has been set.
func (o *CPU) HasCpuCapacityMax() bool {
	return o != nil && o.CpuCapacityMax != nil
}

// SetCpuCapacityMax gets a reference to the given string and assigns it to the CpuCapacityMax field.
func (o *CPU) SetCpuCapacityMax(v string) {
	o.CpuCapacityMax = &v
}


// GetCpuAssigned returns the CpuAssigned field value if set, zero value otherwise.
func (o *CPU) GetCpuAssigned() string {
	if o == nil || o.CpuAssigned == nil {
		var ret string
		return ret
	}
	return *o.CpuAssigned
}

// GetCpuAssignedOk returns a tuple with the CpuAssigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CPU) GetCpuAssignedOk() (*string, bool) {
	if o == nil || o.CpuAssigned == nil {
		return nil, false
	}
	return o.CpuAssigned, true
}

// HasCpuAssigned returns a boolean if a field has been set.
func (o *CPU) HasCpuAssigned() bool {
	return o != nil && o.CpuAssigned != nil
}

// SetCpuAssigned gets a reference to the given string and assigns it to the CpuAssigned field.
func (o *CPU) SetCpuAssigned(v string) {
	o.CpuAssigned = &v
}


// GetCpuAssignedMax returns the CpuAssignedMax field value if set, zero value otherwise.
func (o *CPU) GetCpuAssignedMax() string {
	if o == nil || o.CpuAssignedMax == nil {
		var ret string
		return ret
	}
	return *o.CpuAssignedMax
}

// GetCpuAssignedMaxOk returns a tuple with the CpuAssignedMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CPU) GetCpuAssignedMaxOk() (*string, bool) {
	if o == nil || o.CpuAssignedMax == nil {
		return nil, false
	}
	return o.CpuAssignedMax, true
}

// HasCpuAssignedMax returns a boolean if a field has been set.
func (o *CPU) HasCpuAssignedMax() bool {
	return o != nil && o.CpuAssignedMax != nil
}

// SetCpuAssignedMax gets a reference to the given string and assigns it to the CpuAssignedMax field.
func (o *CPU) SetCpuAssignedMax(v string) {
	o.CpuAssignedMax = &v
}



// MarshalJSON serializes the struct using spec logic.
func (o CPU) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.CpuCapacity != nil {
		toSerialize["cpu_capacity"] = o.CpuCapacity
	}
	if o.CpuCapacityMax != nil {
		toSerialize["cpu_capacity_max"] = o.CpuCapacityMax
	}
	if o.CpuAssigned != nil {
		toSerialize["cpu_assigned"] = o.CpuAssigned
	}
	if o.CpuAssignedMax != nil {
		toSerialize["cpu_assigned_max"] = o.CpuAssignedMax
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CPU) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CpuCapacity *string `json:"cpu_capacity,omitempty"`
		CpuCapacityMax *string `json:"cpu_capacity_max,omitempty"`
		CpuAssigned *string `json:"cpu_assigned,omitempty"`
		CpuAssignedMax *string `json:"cpu_assigned_max,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{ "cpu_capacity", "cpu_capacity_max", "cpu_assigned", "cpu_assigned_max",  })
	} else {
		return err
	}
	o.CpuCapacity = all.CpuCapacity
	o.CpuCapacityMax = all.CpuCapacityMax
	o.CpuAssigned = all.CpuAssigned
	o.CpuAssignedMax = all.CpuAssignedMax

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
