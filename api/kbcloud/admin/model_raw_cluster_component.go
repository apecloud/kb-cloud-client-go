// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// RawClusterComponent Actual resources of one component in a discovered cluster.
type RawClusterComponent struct {
	// Component name, such as vm-storage or vlselect.
	Component *string `json:"component,omitempty"`
	// the replicas of this component
	Replicas *int32 `json:"replicas,omitempty"`
	// Storage size, the unit is Gi
	Storage *float64 `json:"storage,omitempty"`
	// CPU cores.
	Cpu *float64 `json:"cpu,omitempty"`
	// Memory, the unit is Gi.
	Memory *float64 `json:"memory,omitempty"`
	// Class code recorded for this component.
	ClassCode *string `json:"classCode,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRawClusterComponent instantiates a new RawClusterComponent object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRawClusterComponent() *RawClusterComponent {
	this := RawClusterComponent{}
	return &this
}

// NewRawClusterComponentWithDefaults instantiates a new RawClusterComponent object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRawClusterComponentWithDefaults() *RawClusterComponent {
	this := RawClusterComponent{}
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *RawClusterComponent) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RawClusterComponent) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *RawClusterComponent) HasComponent() bool {
	return o != nil && o.Component != nil
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *RawClusterComponent) SetComponent(v string) {
	o.Component = &v
}

// GetReplicas returns the Replicas field value if set, zero value otherwise.
func (o *RawClusterComponent) GetReplicas() int32 {
	if o == nil || o.Replicas == nil {
		var ret int32
		return ret
	}
	return *o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RawClusterComponent) GetReplicasOk() (*int32, bool) {
	if o == nil || o.Replicas == nil {
		return nil, false
	}
	return o.Replicas, true
}

// HasReplicas returns a boolean if a field has been set.
func (o *RawClusterComponent) HasReplicas() bool {
	return o != nil && o.Replicas != nil
}

// SetReplicas gets a reference to the given int32 and assigns it to the Replicas field.
func (o *RawClusterComponent) SetReplicas(v int32) {
	o.Replicas = &v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *RawClusterComponent) GetStorage() float64 {
	if o == nil || o.Storage == nil {
		var ret float64
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RawClusterComponent) GetStorageOk() (*float64, bool) {
	if o == nil || o.Storage == nil {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *RawClusterComponent) HasStorage() bool {
	return o != nil && o.Storage != nil
}

// SetStorage gets a reference to the given float64 and assigns it to the Storage field.
func (o *RawClusterComponent) SetStorage(v float64) {
	o.Storage = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *RawClusterComponent) GetCpu() float64 {
	if o == nil || o.Cpu == nil {
		var ret float64
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RawClusterComponent) GetCpuOk() (*float64, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *RawClusterComponent) HasCpu() bool {
	return o != nil && o.Cpu != nil
}

// SetCpu gets a reference to the given float64 and assigns it to the Cpu field.
func (o *RawClusterComponent) SetCpu(v float64) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *RawClusterComponent) GetMemory() float64 {
	if o == nil || o.Memory == nil {
		var ret float64
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RawClusterComponent) GetMemoryOk() (*float64, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *RawClusterComponent) HasMemory() bool {
	return o != nil && o.Memory != nil
}

// SetMemory gets a reference to the given float64 and assigns it to the Memory field.
func (o *RawClusterComponent) SetMemory(v float64) {
	o.Memory = &v
}

// GetClassCode returns the ClassCode field value if set, zero value otherwise.
func (o *RawClusterComponent) GetClassCode() string {
	if o == nil || o.ClassCode == nil {
		var ret string
		return ret
	}
	return *o.ClassCode
}

// GetClassCodeOk returns a tuple with the ClassCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RawClusterComponent) GetClassCodeOk() (*string, bool) {
	if o == nil || o.ClassCode == nil {
		return nil, false
	}
	return o.ClassCode, true
}

// HasClassCode returns a boolean if a field has been set.
func (o *RawClusterComponent) HasClassCode() bool {
	return o != nil && o.ClassCode != nil
}

// SetClassCode gets a reference to the given string and assigns it to the ClassCode field.
func (o *RawClusterComponent) SetClassCode(v string) {
	o.ClassCode = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RawClusterComponent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	if o.Replicas != nil {
		toSerialize["replicas"] = o.Replicas
	}
	if o.Storage != nil {
		toSerialize["storage"] = o.Storage
	}
	if o.Cpu != nil {
		toSerialize["cpu"] = o.Cpu
	}
	if o.Memory != nil {
		toSerialize["memory"] = o.Memory
	}
	if o.ClassCode != nil {
		toSerialize["classCode"] = o.ClassCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RawClusterComponent) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Component *string  `json:"component,omitempty"`
		Replicas  *int32   `json:"replicas,omitempty"`
		Storage   *float64 `json:"storage,omitempty"`
		Cpu       *float64 `json:"cpu,omitempty"`
		Memory    *float64 `json:"memory,omitempty"`
		ClassCode *string  `json:"classCode,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"component", "replicas", "storage", "cpu", "memory", "classCode"})
	} else {
		return err
	}
	o.Component = all.Component
	o.Replicas = all.Replicas
	o.Storage = all.Storage
	o.Cpu = all.Cpu
	o.Memory = all.Memory
	o.ClassCode = all.ClassCode

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
