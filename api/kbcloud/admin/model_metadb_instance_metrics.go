// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// Metadb_instanceMetrics instance metrics

type Metadb_instanceMetrics struct {
	// the name of the instance
	InstanceName string `json:"instanceName"`
	// cpu with uint cores.
	CpuUsage string `json:"cpuUsage"`
	// memory with uint Gi.
	MemoryUsage string `json:"memoryUsage"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMetadb_instanceMetrics instantiates a new Metadb_instanceMetrics object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMetadb_instanceMetrics(instanceName string, cpuUsage string, memoryUsage string) *Metadb_instanceMetrics {
	this := Metadb_instanceMetrics{}
	this.InstanceName = instanceName
	this.CpuUsage = cpuUsage
	this.MemoryUsage = memoryUsage
	return &this
}

// NewMetadb_instanceMetricsWithDefaults instantiates a new Metadb_instanceMetrics object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMetadb_instanceMetricsWithDefaults() *Metadb_instanceMetrics {
	this := Metadb_instanceMetrics{}
	return &this
}

// GetInstanceName returns the InstanceName field value.
func (o *Metadb_instanceMetrics) GetInstanceName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.InstanceName
}

// GetInstanceNameOk returns a tuple with the InstanceName field value
// and a boolean to check if the value has been set.
func (o *Metadb_instanceMetrics) GetInstanceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceName, true
}

// SetInstanceName sets field value.
func (o *Metadb_instanceMetrics) SetInstanceName(v string) {
	o.InstanceName = v
}

// GetCpuUsage returns the CpuUsage field value.
func (o *Metadb_instanceMetrics) GetCpuUsage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CpuUsage
}

// GetCpuUsageOk returns a tuple with the CpuUsage field value
// and a boolean to check if the value has been set.
func (o *Metadb_instanceMetrics) GetCpuUsageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CpuUsage, true
}

// SetCpuUsage sets field value.
func (o *Metadb_instanceMetrics) SetCpuUsage(v string) {
	o.CpuUsage = v
}

// GetMemoryUsage returns the MemoryUsage field value.
func (o *Metadb_instanceMetrics) GetMemoryUsage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MemoryUsage
}

// GetMemoryUsageOk returns a tuple with the MemoryUsage field value
// and a boolean to check if the value has been set.
func (o *Metadb_instanceMetrics) GetMemoryUsageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MemoryUsage, true
}

// SetMemoryUsage sets field value.
func (o *Metadb_instanceMetrics) SetMemoryUsage(v string) {
	o.MemoryUsage = v
}

// MarshalJSON serializes the struct using spec logic.
func (o Metadb_instanceMetrics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["instanceName"] = o.InstanceName
	toSerialize["cpuUsage"] = o.CpuUsage
	toSerialize["memoryUsage"] = o.MemoryUsage

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Metadb_instanceMetrics) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		InstanceName *string `json:"instanceName"`
		CpuUsage     *string `json:"cpuUsage"`
		MemoryUsage  *string `json:"memoryUsage"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.InstanceName == nil {
		return fmt.Errorf("required field instanceName missing")
	}
	if all.CpuUsage == nil {
		return fmt.Errorf("required field cpuUsage missing")
	}
	if all.MemoryUsage == nil {
		return fmt.Errorf("required field memoryUsage missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"instanceName", "cpuUsage", "memoryUsage"})
	} else {
		return err
	}
	o.InstanceName = *all.InstanceName
	o.CpuUsage = *all.CpuUsage
	o.MemoryUsage = *all.MemoryUsage

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
