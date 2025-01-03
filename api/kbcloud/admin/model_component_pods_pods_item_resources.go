// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ComponentPodsPodsItemResources Resource usage (only for kb-cluster type)
type ComponentPodsPodsItemResources struct {
	// CPU usage
	Cpu *string `json:"cpu,omitempty"`
	// Memory usage
	Memory *string `json:"memory,omitempty"`
	// Storage usage
	Storage *string `json:"storage,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewComponentPodsPodsItemResources instantiates a new ComponentPodsPodsItemResources object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewComponentPodsPodsItemResources() *ComponentPodsPodsItemResources {
	this := ComponentPodsPodsItemResources{}
	return &this
}

// NewComponentPodsPodsItemResourcesWithDefaults instantiates a new ComponentPodsPodsItemResources object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewComponentPodsPodsItemResourcesWithDefaults() *ComponentPodsPodsItemResources {
	this := ComponentPodsPodsItemResources{}
	return &this
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *ComponentPodsPodsItemResources) GetCpu() string {
	if o == nil || o.Cpu == nil {
		var ret string
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentPodsPodsItemResources) GetCpuOk() (*string, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *ComponentPodsPodsItemResources) HasCpu() bool {
	return o != nil && o.Cpu != nil
}

// SetCpu gets a reference to the given string and assigns it to the Cpu field.
func (o *ComponentPodsPodsItemResources) SetCpu(v string) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *ComponentPodsPodsItemResources) GetMemory() string {
	if o == nil || o.Memory == nil {
		var ret string
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentPodsPodsItemResources) GetMemoryOk() (*string, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *ComponentPodsPodsItemResources) HasMemory() bool {
	return o != nil && o.Memory != nil
}

// SetMemory gets a reference to the given string and assigns it to the Memory field.
func (o *ComponentPodsPodsItemResources) SetMemory(v string) {
	o.Memory = &v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *ComponentPodsPodsItemResources) GetStorage() string {
	if o == nil || o.Storage == nil {
		var ret string
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentPodsPodsItemResources) GetStorageOk() (*string, bool) {
	if o == nil || o.Storage == nil {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *ComponentPodsPodsItemResources) HasStorage() bool {
	return o != nil && o.Storage != nil
}

// SetStorage gets a reference to the given string and assigns it to the Storage field.
func (o *ComponentPodsPodsItemResources) SetStorage(v string) {
	o.Storage = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ComponentPodsPodsItemResources) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Cpu != nil {
		toSerialize["cpu"] = o.Cpu
	}
	if o.Memory != nil {
		toSerialize["memory"] = o.Memory
	}
	if o.Storage != nil {
		toSerialize["storage"] = o.Storage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ComponentPodsPodsItemResources) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Cpu     *string `json:"cpu,omitempty"`
		Memory  *string `json:"memory,omitempty"`
		Storage *string `json:"storage,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"cpu", "memory", "storage"})
	} else {
		return err
	}
	o.Cpu = all.Cpu
	o.Memory = all.Memory
	o.Storage = all.Storage

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
