// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// StorageStats StorageStats holds the resource stats of the volume, including bound PVC capacity, pod-used PVC capacity, and actual storage usage.
type StorageStats struct {
	// BoundCapacity is the total requested size of all PVCs that are bound to volumes, unit is GiB
	BoundCapacity float64 `json:"boundCapacity"`
	// PodUsedCapacity is the total requested size of PVCs that are bound and currently used by pods, unit is GiB
	PodUsedCapacity float64 `json:"podUsedCapacity"`
	// PodUsedUsage is the actual storage usage from PVCs that are used by pods, unit is GiB
	PodUsedUsage float64 `json:"podUsedUsage"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStorageStats instantiates a new StorageStats object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStorageStats(boundCapacity float64, podUsedCapacity float64, podUsedUsage float64) *StorageStats {
	this := StorageStats{}
	this.BoundCapacity = boundCapacity
	this.PodUsedCapacity = podUsedCapacity
	this.PodUsedUsage = podUsedUsage
	return &this
}

// NewStorageStatsWithDefaults instantiates a new StorageStats object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStorageStatsWithDefaults() *StorageStats {
	this := StorageStats{}
	return &this
}

// GetBoundCapacity returns the BoundCapacity field value.
func (o *StorageStats) GetBoundCapacity() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.BoundCapacity
}

// GetBoundCapacityOk returns a tuple with the BoundCapacity field value
// and a boolean to check if the value has been set.
func (o *StorageStats) GetBoundCapacityOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BoundCapacity, true
}

// SetBoundCapacity sets field value.
func (o *StorageStats) SetBoundCapacity(v float64) {
	o.BoundCapacity = v
}

// GetPodUsedCapacity returns the PodUsedCapacity field value.
func (o *StorageStats) GetPodUsedCapacity() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.PodUsedCapacity
}

// GetPodUsedCapacityOk returns a tuple with the PodUsedCapacity field value
// and a boolean to check if the value has been set.
func (o *StorageStats) GetPodUsedCapacityOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PodUsedCapacity, true
}

// SetPodUsedCapacity sets field value.
func (o *StorageStats) SetPodUsedCapacity(v float64) {
	o.PodUsedCapacity = v
}

// GetPodUsedUsage returns the PodUsedUsage field value.
func (o *StorageStats) GetPodUsedUsage() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.PodUsedUsage
}

// GetPodUsedUsageOk returns a tuple with the PodUsedUsage field value
// and a boolean to check if the value has been set.
func (o *StorageStats) GetPodUsedUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PodUsedUsage, true
}

// SetPodUsedUsage sets field value.
func (o *StorageStats) SetPodUsedUsage(v float64) {
	o.PodUsedUsage = v
}

// MarshalJSON serializes the struct using spec logic.
func (o StorageStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["boundCapacity"] = o.BoundCapacity
	toSerialize["podUsedCapacity"] = o.PodUsedCapacity
	toSerialize["podUsedUsage"] = o.PodUsedUsage

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StorageStats) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BoundCapacity   *float64 `json:"boundCapacity"`
		PodUsedCapacity *float64 `json:"podUsedCapacity"`
		PodUsedUsage    *float64 `json:"podUsedUsage"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.BoundCapacity == nil {
		return fmt.Errorf("required field boundCapacity missing")
	}
	if all.PodUsedCapacity == nil {
		return fmt.Errorf("required field podUsedCapacity missing")
	}
	if all.PodUsedUsage == nil {
		return fmt.Errorf("required field podUsedUsage missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"boundCapacity", "podUsedCapacity", "podUsedUsage"})
	} else {
		return err
	}
	o.BoundCapacity = *all.BoundCapacity
	o.PodUsedCapacity = *all.PodUsedCapacity
	o.PodUsedUsage = *all.PodUsedUsage

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
