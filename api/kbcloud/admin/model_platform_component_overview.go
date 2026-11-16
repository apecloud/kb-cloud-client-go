// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// PlatformComponentOverview Central platform resources. CPU is measured in cores; memory and storage in GiB. Missing usage is -1. Storage includes only PVCs referenced by platform component Pods.
type PlatformComponentOverview struct {
	Nodes []PlatformComponentNode `json:"nodes"`
	// ResourceStats holds the requests, limits, and available stats for a resource.
	CpuStats ResourceStats `json:"cpuStats"`
	// ResourceStats holds the requests, limits, and available stats for a resource.
	MemoryStats ResourceStats `json:"memoryStats"`
	// StorageStats holds the resource stats of the volume, including bound PVC capacity, pod-used PVC capacity, and actual storage usage.
	StorageStats StorageStats              `json:"storageStats"`
	PodStats     PlatformComponentPodStats `json:"podStats"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPlatformComponentOverview instantiates a new PlatformComponentOverview object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPlatformComponentOverview(nodes []PlatformComponentNode, cpuStats ResourceStats, memoryStats ResourceStats, storageStats StorageStats, podStats PlatformComponentPodStats) *PlatformComponentOverview {
	this := PlatformComponentOverview{}
	this.Nodes = nodes
	this.CpuStats = cpuStats
	this.MemoryStats = memoryStats
	this.StorageStats = storageStats
	this.PodStats = podStats
	return &this
}

// NewPlatformComponentOverviewWithDefaults instantiates a new PlatformComponentOverview object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPlatformComponentOverviewWithDefaults() *PlatformComponentOverview {
	this := PlatformComponentOverview{}
	return &this
}

// GetNodes returns the Nodes field value.
func (o *PlatformComponentOverview) GetNodes() []PlatformComponentNode {
	if o == nil {
		var ret []PlatformComponentNode
		return ret
	}
	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value
// and a boolean to check if the value has been set.
func (o *PlatformComponentOverview) GetNodesOk() (*[]PlatformComponentNode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nodes, true
}

// SetNodes sets field value.
func (o *PlatformComponentOverview) SetNodes(v []PlatformComponentNode) {
	o.Nodes = v
}

// GetCpuStats returns the CpuStats field value.
func (o *PlatformComponentOverview) GetCpuStats() ResourceStats {
	if o == nil {
		var ret ResourceStats
		return ret
	}
	return o.CpuStats
}

// GetCpuStatsOk returns a tuple with the CpuStats field value
// and a boolean to check if the value has been set.
func (o *PlatformComponentOverview) GetCpuStatsOk() (*ResourceStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CpuStats, true
}

// SetCpuStats sets field value.
func (o *PlatformComponentOverview) SetCpuStats(v ResourceStats) {
	o.CpuStats = v
}

// GetMemoryStats returns the MemoryStats field value.
func (o *PlatformComponentOverview) GetMemoryStats() ResourceStats {
	if o == nil {
		var ret ResourceStats
		return ret
	}
	return o.MemoryStats
}

// GetMemoryStatsOk returns a tuple with the MemoryStats field value
// and a boolean to check if the value has been set.
func (o *PlatformComponentOverview) GetMemoryStatsOk() (*ResourceStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MemoryStats, true
}

// SetMemoryStats sets field value.
func (o *PlatformComponentOverview) SetMemoryStats(v ResourceStats) {
	o.MemoryStats = v
}

// GetStorageStats returns the StorageStats field value.
func (o *PlatformComponentOverview) GetStorageStats() StorageStats {
	if o == nil {
		var ret StorageStats
		return ret
	}
	return o.StorageStats
}

// GetStorageStatsOk returns a tuple with the StorageStats field value
// and a boolean to check if the value has been set.
func (o *PlatformComponentOverview) GetStorageStatsOk() (*StorageStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageStats, true
}

// SetStorageStats sets field value.
func (o *PlatformComponentOverview) SetStorageStats(v StorageStats) {
	o.StorageStats = v
}

// GetPodStats returns the PodStats field value.
func (o *PlatformComponentOverview) GetPodStats() PlatformComponentPodStats {
	if o == nil {
		var ret PlatformComponentPodStats
		return ret
	}
	return o.PodStats
}

// GetPodStatsOk returns a tuple with the PodStats field value
// and a boolean to check if the value has been set.
func (o *PlatformComponentOverview) GetPodStatsOk() (*PlatformComponentPodStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PodStats, true
}

// SetPodStats sets field value.
func (o *PlatformComponentOverview) SetPodStats(v PlatformComponentPodStats) {
	o.PodStats = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PlatformComponentOverview) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["nodes"] = o.Nodes
	toSerialize["cpuStats"] = o.CpuStats
	toSerialize["memoryStats"] = o.MemoryStats
	toSerialize["storageStats"] = o.StorageStats
	toSerialize["podStats"] = o.PodStats

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PlatformComponentOverview) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Nodes        *[]PlatformComponentNode   `json:"nodes"`
		CpuStats     *ResourceStats             `json:"cpuStats"`
		MemoryStats  *ResourceStats             `json:"memoryStats"`
		StorageStats *StorageStats              `json:"storageStats"`
		PodStats     *PlatformComponentPodStats `json:"podStats"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Nodes == nil {
		return fmt.Errorf("required field nodes missing")
	}
	if all.CpuStats == nil {
		return fmt.Errorf("required field cpuStats missing")
	}
	if all.MemoryStats == nil {
		return fmt.Errorf("required field memoryStats missing")
	}
	if all.StorageStats == nil {
		return fmt.Errorf("required field storageStats missing")
	}
	if all.PodStats == nil {
		return fmt.Errorf("required field podStats missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"nodes", "cpuStats", "memoryStats", "storageStats", "podStats"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Nodes = *all.Nodes
	if all.CpuStats.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CpuStats = *all.CpuStats
	if all.MemoryStats.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.MemoryStats = *all.MemoryStats
	if all.StorageStats.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.StorageStats = *all.StorageStats
	if all.PodStats.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.PodStats = *all.PodStats

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
