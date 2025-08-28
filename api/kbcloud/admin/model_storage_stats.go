// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// StorageStats StorageStats holds the resource stats of the volume, such as provisioned capacity, etc.
type StorageStats struct {
	// The total requested size of PVCs that are bound and currently used by pods, unit is GiB
	Requests float64 `json:"requests"`
	// Usage is the actual storage usage from PVCs that are used by pods, unit is GiB
	Usage float64 `json:"usage"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStorageStats instantiates a new StorageStats object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStorageStats(requests float64, usage float64) *StorageStats {
	this := StorageStats{}
	this.Requests = requests
	this.Usage = usage
	return &this
}

// NewStorageStatsWithDefaults instantiates a new StorageStats object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStorageStatsWithDefaults() *StorageStats {
	this := StorageStats{}
	return &this
}

// GetRequests returns the Requests field value.
func (o *StorageStats) GetRequests() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value
// and a boolean to check if the value has been set.
func (o *StorageStats) GetRequestsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requests, true
}

// SetRequests sets field value.
func (o *StorageStats) SetRequests(v float64) {
	o.Requests = v
}

// GetUsage returns the Usage field value.
func (o *StorageStats) GetUsage() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value
// and a boolean to check if the value has been set.
func (o *StorageStats) GetUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Usage, true
}

// SetUsage sets field value.
func (o *StorageStats) SetUsage(v float64) {
	o.Usage = v
}

// MarshalJSON serializes the struct using spec logic.
func (o StorageStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["requests"] = o.Requests
	toSerialize["usage"] = o.Usage

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StorageStats) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Requests *float64 `json:"requests"`
		Usage    *float64 `json:"usage"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Requests == nil {
		return fmt.Errorf("required field requests missing")
	}
	if all.Usage == nil {
		return fmt.Errorf("required field usage missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"requests", "usage"})
	} else {
		return err
	}
	o.Requests = *all.Requests
	o.Usage = *all.Usage

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
