// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// OrgResourceQuota org resource quota
type OrgResourceQuota struct {
	// Maximum available vCPU. if set to 0, no limit
	Cpu float64 `json:"cpu"`
	// Maximum available memory in GB. if set to 0, no limit
	Memory float64 `json:"memory"`
	// Maximum available storage in GB. if set to 0, no limit
	Storage float64 `json:"storage"`
	// Number of the clusters. key is engine type, values is the maximum number of engine
	Clusters map[string]int32 `json:"clusters"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgResourceQuota instantiates a new OrgResourceQuota object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgResourceQuota(cpu float64, memory float64, storage float64, clusters map[string]int32) *OrgResourceQuota {
	this := OrgResourceQuota{}
	this.Cpu = cpu
	this.Memory = memory
	this.Storage = storage
	this.Clusters = clusters
	return &this
}

// NewOrgResourceQuotaWithDefaults instantiates a new OrgResourceQuota object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgResourceQuotaWithDefaults() *OrgResourceQuota {
	this := OrgResourceQuota{}
	return &this
}

// GetCpu returns the Cpu field value.
func (o *OrgResourceQuota) GetCpu() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value
// and a boolean to check if the value has been set.
func (o *OrgResourceQuota) GetCpuOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cpu, true
}

// SetCpu sets field value.
func (o *OrgResourceQuota) SetCpu(v float64) {
	o.Cpu = v
}

// GetMemory returns the Memory field value.
func (o *OrgResourceQuota) GetMemory() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value
// and a boolean to check if the value has been set.
func (o *OrgResourceQuota) GetMemoryOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Memory, true
}

// SetMemory sets field value.
func (o *OrgResourceQuota) SetMemory(v float64) {
	o.Memory = v
}

// GetStorage returns the Storage field value.
func (o *OrgResourceQuota) GetStorage() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value
// and a boolean to check if the value has been set.
func (o *OrgResourceQuota) GetStorageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Storage, true
}

// SetStorage sets field value.
func (o *OrgResourceQuota) SetStorage(v float64) {
	o.Storage = v
}

// GetClusters returns the Clusters field value.
func (o *OrgResourceQuota) GetClusters() map[string]int32 {
	if o == nil {
		var ret map[string]int32
		return ret
	}
	return o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value
// and a boolean to check if the value has been set.
func (o *OrgResourceQuota) GetClustersOk() (*map[string]int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Clusters, true
}

// SetClusters sets field value.
func (o *OrgResourceQuota) SetClusters(v map[string]int32) {
	o.Clusters = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgResourceQuota) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["cpu"] = o.Cpu
	toSerialize["memory"] = o.Memory
	toSerialize["storage"] = o.Storage
	toSerialize["clusters"] = o.Clusters

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgResourceQuota) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Cpu      *float64          `json:"cpu"`
		Memory   *float64          `json:"memory"`
		Storage  *float64          `json:"storage"`
		Clusters *map[string]int32 `json:"clusters"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Cpu == nil {
		return fmt.Errorf("required field cpu missing")
	}
	if all.Memory == nil {
		return fmt.Errorf("required field memory missing")
	}
	if all.Storage == nil {
		return fmt.Errorf("required field storage missing")
	}
	if all.Clusters == nil {
		return fmt.Errorf("required field clusters missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"cpu", "memory", "storage", "clusters"})
	} else {
		return err
	}
	o.Cpu = *all.Cpu
	o.Memory = *all.Memory
	o.Storage = *all.Storage
	o.Clusters = *all.Clusters

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
