// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// NodeResourceStats ResourceStats holds the requests, limits, and available stats for a resource.
type NodeResourceStats struct {
	// ResourceStats holds the requests, limits, and available stats for a resource.
	CpuStats ResourceStats `json:"cpuStats"`
	// ResourceStats holds the requests, limits, and available stats for a resource.
	MemoryStats ResourceStats `json:"memoryStats"`
	// Name of the node.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNodeResourceStats instantiates a new NodeResourceStats object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNodeResourceStats(cpuStats ResourceStats, memoryStats ResourceStats, name string) *NodeResourceStats {
	this := NodeResourceStats{}
	this.CpuStats = cpuStats
	this.MemoryStats = memoryStats
	this.Name = name
	return &this
}

// NewNodeResourceStatsWithDefaults instantiates a new NodeResourceStats object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNodeResourceStatsWithDefaults() *NodeResourceStats {
	this := NodeResourceStats{}
	return &this
}

// GetCpuStats returns the CpuStats field value.
func (o *NodeResourceStats) GetCpuStats() ResourceStats {
	if o == nil {
		var ret ResourceStats
		return ret
	}
	return o.CpuStats
}

// GetCpuStatsOk returns a tuple with the CpuStats field value
// and a boolean to check if the value has been set.
func (o *NodeResourceStats) GetCpuStatsOk() (*ResourceStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CpuStats, true
}

// SetCpuStats sets field value.
func (o *NodeResourceStats) SetCpuStats(v ResourceStats) {
	o.CpuStats = v
}

// GetMemoryStats returns the MemoryStats field value.
func (o *NodeResourceStats) GetMemoryStats() ResourceStats {
	if o == nil {
		var ret ResourceStats
		return ret
	}
	return o.MemoryStats
}

// GetMemoryStatsOk returns a tuple with the MemoryStats field value
// and a boolean to check if the value has been set.
func (o *NodeResourceStats) GetMemoryStatsOk() (*ResourceStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MemoryStats, true
}

// SetMemoryStats sets field value.
func (o *NodeResourceStats) SetMemoryStats(v ResourceStats) {
	o.MemoryStats = v
}

// GetName returns the Name field value.
func (o *NodeResourceStats) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NodeResourceStats) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *NodeResourceStats) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o NodeResourceStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["cpuStats"] = o.CpuStats
	toSerialize["memoryStats"] = o.MemoryStats
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NodeResourceStats) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CpuStats    *ResourceStats `json:"cpuStats"`
		MemoryStats *ResourceStats `json:"memoryStats"`
		Name        *string        `json:"name"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.CpuStats == nil {
		return fmt.Errorf("required field cpuStats missing")
	}
	if all.MemoryStats == nil {
		return fmt.Errorf("required field memoryStats missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"cpuStats", "memoryStats", "name"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.CpuStats.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CpuStats = *all.CpuStats
	if all.MemoryStats.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.MemoryStats = *all.MemoryStats
	o.Name = *all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
