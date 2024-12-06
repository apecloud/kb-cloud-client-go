// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
)

// StorageStats StorageStats holds the resource stats of the volume, such as provisioned capacity, etc.
type StorageStats struct {
	// ProvisionedCapacity is the actual size of the volumes that is bound to the PVC, unit is GiB
	ProvisionedCapacity float64 `json:"provisionedCapacity"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStorageStats instantiates a new StorageStats object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStorageStats(provisionedCapacity float64) *StorageStats {
	this := StorageStats{}
	this.ProvisionedCapacity = provisionedCapacity
	return &this
}

// NewStorageStatsWithDefaults instantiates a new StorageStats object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStorageStatsWithDefaults() *StorageStats {
	this := StorageStats{}
	return &this
}

// GetProvisionedCapacity returns the ProvisionedCapacity field value.
func (o *StorageStats) GetProvisionedCapacity() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.ProvisionedCapacity
}

// GetProvisionedCapacityOk returns a tuple with the ProvisionedCapacity field value
// and a boolean to check if the value has been set.
func (o *StorageStats) GetProvisionedCapacityOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProvisionedCapacity, true
}

// SetProvisionedCapacity sets field value.
func (o *StorageStats) SetProvisionedCapacity(v float64) {
	o.ProvisionedCapacity = v
}

// MarshalJSON serializes the struct using spec logic.
func (o StorageStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["provisionedCapacity"] = o.ProvisionedCapacity

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StorageStats) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ProvisionedCapacity *float64 `json:"provisionedCapacity"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ProvisionedCapacity == nil {
		return fmt.Errorf("required field provisionedCapacity missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"provisionedCapacity"})
	} else {
		return err
	}
	o.ProvisionedCapacity = *all.ProvisionedCapacity

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
