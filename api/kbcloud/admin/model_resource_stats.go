// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ResourceStats ResourceStats holds the requests, limits, and available stats for a resource.
type ResourceStats struct {
	// The amount of CPU or Memory resources that are available on the node. Unit is GiB for memory and Cores for CPU.
	Allocatable float64 `json:"allocatable"`
	// The maximum number of CPU or Memory resources pods are allowed to use on the node.  Unit is GiB for memory and Cores for CPU
	Limits float64 `json:"limits"`
	// The number of CPU or Memory resources requested by pods running on the node. Unit is GiB for memory and Cores for CPU.
	Requests float64 `json:"requests"`
	// The amount of CPU or Memory resources that are already used on the node. Unit is GiB for memory and Cores for CPU.
	Usage float64 `json:"usage"`
	// The total amount of physical resources available on the node. Unit is GiB for memory and Cores for CPU.
	Capacity *float64 `json:"capacity,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewResourceStats instantiates a new ResourceStats object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewResourceStats(allocatable float64, limits float64, requests float64, usage float64) *ResourceStats {
	this := ResourceStats{}
	this.Allocatable = allocatable
	this.Limits = limits
	this.Requests = requests
	this.Usage = usage
	return &this
}

// NewResourceStatsWithDefaults instantiates a new ResourceStats object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewResourceStatsWithDefaults() *ResourceStats {
	this := ResourceStats{}
	return &this
}

// GetAllocatable returns the Allocatable field value.
func (o *ResourceStats) GetAllocatable() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Allocatable
}

// GetAllocatableOk returns a tuple with the Allocatable field value
// and a boolean to check if the value has been set.
func (o *ResourceStats) GetAllocatableOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Allocatable, true
}

// SetAllocatable sets field value.
func (o *ResourceStats) SetAllocatable(v float64) {
	o.Allocatable = v
}

// GetLimits returns the Limits field value.
func (o *ResourceStats) GetLimits() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Limits
}

// GetLimitsOk returns a tuple with the Limits field value
// and a boolean to check if the value has been set.
func (o *ResourceStats) GetLimitsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limits, true
}

// SetLimits sets field value.
func (o *ResourceStats) SetLimits(v float64) {
	o.Limits = v
}

// GetRequests returns the Requests field value.
func (o *ResourceStats) GetRequests() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value
// and a boolean to check if the value has been set.
func (o *ResourceStats) GetRequestsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requests, true
}

// SetRequests sets field value.
func (o *ResourceStats) SetRequests(v float64) {
	o.Requests = v
}

// GetUsage returns the Usage field value.
func (o *ResourceStats) GetUsage() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value
// and a boolean to check if the value has been set.
func (o *ResourceStats) GetUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Usage, true
}

// SetUsage sets field value.
func (o *ResourceStats) SetUsage(v float64) {
	o.Usage = v
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
	toSerialize["allocatable"] = o.Allocatable
	toSerialize["limits"] = o.Limits
	toSerialize["requests"] = o.Requests
	toSerialize["usage"] = o.Usage
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
		Allocatable *float64 `json:"allocatable"`
		Limits      *float64 `json:"limits"`
		Requests    *float64 `json:"requests"`
		Usage       *float64 `json:"usage"`
		Capacity    *float64 `json:"capacity,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Allocatable == nil {
		return fmt.Errorf("required field allocatable missing")
	}
	if all.Limits == nil {
		return fmt.Errorf("required field limits missing")
	}
	if all.Requests == nil {
		return fmt.Errorf("required field requests missing")
	}
	if all.Usage == nil {
		return fmt.Errorf("required field usage missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"allocatable", "limits", "requests", "usage", "capacity"})
	} else {
		return err
	}
	o.Allocatable = *all.Allocatable
	o.Limits = *all.Limits
	o.Requests = *all.Requests
	o.Usage = *all.Usage
	o.Capacity = all.Capacity

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
