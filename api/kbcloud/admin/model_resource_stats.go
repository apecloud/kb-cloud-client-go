// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package admin

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)



// ResourceStats ResourceStats holds the requests, limits, and available stats for a resource. 
type ResourceStats struct {
	// The amount of CPU or Memory resources that are available on the node. Unit is GiB for memory and Cores for CPU.
	Allocatable *float64 `json:"allocatable,omitempty"`
	// The maximum number of CPU or Memory resources pods are allowed to use on the node.  Unit is GiB for memory and Cores for CPU
	Limits *float64 `json:"limits,omitempty"`
	// The number of CPU or Memory resources requested by pods running on the node. Unit is GiB for memory and Cores for CPU.
	Requests *float64 `json:"requests,omitempty"`
	// The amount of CPU or Memory resources that are already used on the node. Unit is GiB for memory and Cores for CPU.
	Usage *float64 `json:"usage,omitempty"`
	// The total amount of physical resources available on the node. Unit is GiB for memory and Cores for CPU.
	Capacity *float64 `json:"capacity,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewResourceStats instantiates a new ResourceStats object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewResourceStats() *ResourceStats {
	this := ResourceStats{}
	return &this
}

// NewResourceStatsWithDefaults instantiates a new ResourceStats object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewResourceStatsWithDefaults() *ResourceStats {
	this := ResourceStats{}
	return &this
}
// GetAllocatable returns the Allocatable field value if set, zero value otherwise.
func (o *ResourceStats) GetAllocatable() float64 {
	if o == nil || o.Allocatable == nil {
		var ret float64
		return ret
	}
	return *o.Allocatable
}

// GetAllocatableOk returns a tuple with the Allocatable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceStats) GetAllocatableOk() (*float64, bool) {
	if o == nil || o.Allocatable == nil {
		return nil, false
	}
	return o.Allocatable, true
}

// HasAllocatable returns a boolean if a field has been set.
func (o *ResourceStats) HasAllocatable() bool {
	return o != nil && o.Allocatable != nil
}

// SetAllocatable gets a reference to the given float64 and assigns it to the Allocatable field.
func (o *ResourceStats) SetAllocatable(v float64) {
	o.Allocatable = &v
}


// GetLimits returns the Limits field value if set, zero value otherwise.
func (o *ResourceStats) GetLimits() float64 {
	if o == nil || o.Limits == nil {
		var ret float64
		return ret
	}
	return *o.Limits
}

// GetLimitsOk returns a tuple with the Limits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceStats) GetLimitsOk() (*float64, bool) {
	if o == nil || o.Limits == nil {
		return nil, false
	}
	return o.Limits, true
}

// HasLimits returns a boolean if a field has been set.
func (o *ResourceStats) HasLimits() bool {
	return o != nil && o.Limits != nil
}

// SetLimits gets a reference to the given float64 and assigns it to the Limits field.
func (o *ResourceStats) SetLimits(v float64) {
	o.Limits = &v
}


// GetRequests returns the Requests field value if set, zero value otherwise.
func (o *ResourceStats) GetRequests() float64 {
	if o == nil || o.Requests == nil {
		var ret float64
		return ret
	}
	return *o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceStats) GetRequestsOk() (*float64, bool) {
	if o == nil || o.Requests == nil {
		return nil, false
	}
	return o.Requests, true
}

// HasRequests returns a boolean if a field has been set.
func (o *ResourceStats) HasRequests() bool {
	return o != nil && o.Requests != nil
}

// SetRequests gets a reference to the given float64 and assigns it to the Requests field.
func (o *ResourceStats) SetRequests(v float64) {
	o.Requests = &v
}


// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *ResourceStats) GetUsage() float64 {
	if o == nil || o.Usage == nil {
		var ret float64
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceStats) GetUsageOk() (*float64, bool) {
	if o == nil || o.Usage == nil {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *ResourceStats) HasUsage() bool {
	return o != nil && o.Usage != nil
}

// SetUsage gets a reference to the given float64 and assigns it to the Usage field.
func (o *ResourceStats) SetUsage(v float64) {
	o.Usage = &v
}


// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *ResourceStats) GetCapacity() float64 {
	if o == nil || o.Capacity == nil {
		var ret float64
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceStats) GetCapacityOk() (*float64, bool) {
	if o == nil || o.Capacity == nil {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *ResourceStats) HasCapacity() bool {
	return o != nil && o.Capacity != nil
}

// SetCapacity gets a reference to the given float64 and assigns it to the Capacity field.
func (o *ResourceStats) SetCapacity(v float64) {
	o.Capacity = &v
}



// MarshalJSON serializes the struct using spec logic.
func (o ResourceStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Allocatable != nil {
		toSerialize["allocatable"] = o.Allocatable
	}
	if o.Limits != nil {
		toSerialize["limits"] = o.Limits
	}
	if o.Requests != nil {
		toSerialize["requests"] = o.Requests
	}
	if o.Usage != nil {
		toSerialize["usage"] = o.Usage
	}
	if o.Capacity != nil {
		toSerialize["capacity"] = o.Capacity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ResourceStats) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Allocatable *float64 `json:"allocatable,omitempty"`
		Limits *float64 `json:"limits,omitempty"`
		Requests *float64 `json:"requests,omitempty"`
		Usage *float64 `json:"usage,omitempty"`
		Capacity *float64 `json:"capacity,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{ "allocatable", "limits", "requests", "usage", "capacity",  })
	} else {
		return err
	}
	o.Allocatable = all.Allocatable
	o.Limits = all.Limits
	o.Requests = all.Requests
	o.Usage = all.Usage
	o.Capacity = all.Capacity

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
