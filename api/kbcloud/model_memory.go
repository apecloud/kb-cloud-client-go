// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type Memory struct {
	// NODESCRIPTION MemCapacity
	MemCapacity *string `json:"mem_capacity,omitempty"`
	// NODESCRIPTION MemoryLimit
	MemoryLimit *string `json:"memory_limit,omitempty"`
	// NODESCRIPTION MemAssigned
	MemAssigned *string `json:"mem_assigned,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMemory instantiates a new Memory object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMemory() *Memory {
	this := Memory{}
	return &this
}

// NewMemoryWithDefaults instantiates a new Memory object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMemoryWithDefaults() *Memory {
	this := Memory{}
	return &this
}

// GetMemCapacity returns the MemCapacity field value if set, zero value otherwise.
func (o *Memory) GetMemCapacity() string {
	if o == nil || o.MemCapacity == nil {
		var ret string
		return ret
	}
	return *o.MemCapacity
}

// GetMemCapacityOk returns a tuple with the MemCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Memory) GetMemCapacityOk() (*string, bool) {
	if o == nil || o.MemCapacity == nil {
		return nil, false
	}
	return o.MemCapacity, true
}

// HasMemCapacity returns a boolean if a field has been set.
func (o *Memory) HasMemCapacity() bool {
	return o != nil && o.MemCapacity != nil
}

// SetMemCapacity gets a reference to the given string and assigns it to the MemCapacity field.
func (o *Memory) SetMemCapacity(v string) {
	o.MemCapacity = &v
}

// GetMemoryLimit returns the MemoryLimit field value if set, zero value otherwise.
func (o *Memory) GetMemoryLimit() string {
	if o == nil || o.MemoryLimit == nil {
		var ret string
		return ret
	}
	return *o.MemoryLimit
}

// GetMemoryLimitOk returns a tuple with the MemoryLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Memory) GetMemoryLimitOk() (*string, bool) {
	if o == nil || o.MemoryLimit == nil {
		return nil, false
	}
	return o.MemoryLimit, true
}

// HasMemoryLimit returns a boolean if a field has been set.
func (o *Memory) HasMemoryLimit() bool {
	return o != nil && o.MemoryLimit != nil
}

// SetMemoryLimit gets a reference to the given string and assigns it to the MemoryLimit field.
func (o *Memory) SetMemoryLimit(v string) {
	o.MemoryLimit = &v
}

// GetMemAssigned returns the MemAssigned field value if set, zero value otherwise.
func (o *Memory) GetMemAssigned() string {
	if o == nil || o.MemAssigned == nil {
		var ret string
		return ret
	}
	return *o.MemAssigned
}

// GetMemAssignedOk returns a tuple with the MemAssigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Memory) GetMemAssignedOk() (*string, bool) {
	if o == nil || o.MemAssigned == nil {
		return nil, false
	}
	return o.MemAssigned, true
}

// HasMemAssigned returns a boolean if a field has been set.
func (o *Memory) HasMemAssigned() bool {
	return o != nil && o.MemAssigned != nil
}

// SetMemAssigned gets a reference to the given string and assigns it to the MemAssigned field.
func (o *Memory) SetMemAssigned(v string) {
	o.MemAssigned = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o Memory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.MemCapacity != nil {
		toSerialize["mem_capacity"] = o.MemCapacity
	}
	if o.MemoryLimit != nil {
		toSerialize["memory_limit"] = o.MemoryLimit
	}
	if o.MemAssigned != nil {
		toSerialize["mem_assigned"] = o.MemAssigned
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Memory) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MemCapacity *string `json:"mem_capacity,omitempty"`
		MemoryLimit *string `json:"memory_limit,omitempty"`
		MemAssigned *string `json:"mem_assigned,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"mem_capacity", "memory_limit", "mem_assigned"})
	} else {
		return err
	}
	o.MemCapacity = all.MemCapacity
	o.MemoryLimit = all.MemoryLimit
	o.MemAssigned = all.MemAssigned

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
