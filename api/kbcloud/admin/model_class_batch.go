// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type ClassBatch struct {
	Engine           *string  `json:"engine,omitempty"`
	Mode             *string  `json:"mode,omitempty"`
	Component        *string  `json:"component,omitempty"`
	InstanceType     *string  `json:"instanceType,omitempty"`
	CpuOverCommit    *float64 `json:"cpuOverCommit,omitempty"`
	MemoryOverCommit *float64 `json:"memoryOverCommit,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClassBatch instantiates a new ClassBatch object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClassBatch() *ClassBatch {
	this := ClassBatch{}
	return &this
}

// NewClassBatchWithDefaults instantiates a new ClassBatch object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClassBatchWithDefaults() *ClassBatch {
	this := ClassBatch{}
	return &this
}

// GetEngine returns the Engine field value if set, zero value otherwise.
func (o *ClassBatch) GetEngine() string {
	if o == nil || o.Engine == nil {
		var ret string
		return ret
	}
	return *o.Engine
}

// GetEngineOk returns a tuple with the Engine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassBatch) GetEngineOk() (*string, bool) {
	if o == nil || o.Engine == nil {
		return nil, false
	}
	return o.Engine, true
}

// HasEngine returns a boolean if a field has been set.
func (o *ClassBatch) HasEngine() bool {
	return o != nil && o.Engine != nil
}

// SetEngine gets a reference to the given string and assigns it to the Engine field.
func (o *ClassBatch) SetEngine(v string) {
	o.Engine = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *ClassBatch) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassBatch) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *ClassBatch) HasMode() bool {
	return o != nil && o.Mode != nil
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *ClassBatch) SetMode(v string) {
	o.Mode = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *ClassBatch) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassBatch) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *ClassBatch) HasComponent() bool {
	return o != nil && o.Component != nil
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *ClassBatch) SetComponent(v string) {
	o.Component = &v
}

// GetInstanceType returns the InstanceType field value if set, zero value otherwise.
func (o *ClassBatch) GetInstanceType() string {
	if o == nil || o.InstanceType == nil {
		var ret string
		return ret
	}
	return *o.InstanceType
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassBatch) GetInstanceTypeOk() (*string, bool) {
	if o == nil || o.InstanceType == nil {
		return nil, false
	}
	return o.InstanceType, true
}

// HasInstanceType returns a boolean if a field has been set.
func (o *ClassBatch) HasInstanceType() bool {
	return o != nil && o.InstanceType != nil
}

// SetInstanceType gets a reference to the given string and assigns it to the InstanceType field.
func (o *ClassBatch) SetInstanceType(v string) {
	o.InstanceType = &v
}

// GetCpuOverCommit returns the CpuOverCommit field value if set, zero value otherwise.
func (o *ClassBatch) GetCpuOverCommit() float64 {
	if o == nil || o.CpuOverCommit == nil {
		var ret float64
		return ret
	}
	return *o.CpuOverCommit
}

// GetCpuOverCommitOk returns a tuple with the CpuOverCommit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassBatch) GetCpuOverCommitOk() (*float64, bool) {
	if o == nil || o.CpuOverCommit == nil {
		return nil, false
	}
	return o.CpuOverCommit, true
}

// HasCpuOverCommit returns a boolean if a field has been set.
func (o *ClassBatch) HasCpuOverCommit() bool {
	return o != nil && o.CpuOverCommit != nil
}

// SetCpuOverCommit gets a reference to the given float64 and assigns it to the CpuOverCommit field.
func (o *ClassBatch) SetCpuOverCommit(v float64) {
	o.CpuOverCommit = &v
}

// GetMemoryOverCommit returns the MemoryOverCommit field value if set, zero value otherwise.
func (o *ClassBatch) GetMemoryOverCommit() float64 {
	if o == nil || o.MemoryOverCommit == nil {
		var ret float64
		return ret
	}
	return *o.MemoryOverCommit
}

// GetMemoryOverCommitOk returns a tuple with the MemoryOverCommit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassBatch) GetMemoryOverCommitOk() (*float64, bool) {
	if o == nil || o.MemoryOverCommit == nil {
		return nil, false
	}
	return o.MemoryOverCommit, true
}

// HasMemoryOverCommit returns a boolean if a field has been set.
func (o *ClassBatch) HasMemoryOverCommit() bool {
	return o != nil && o.MemoryOverCommit != nil
}

// SetMemoryOverCommit gets a reference to the given float64 and assigns it to the MemoryOverCommit field.
func (o *ClassBatch) SetMemoryOverCommit(v float64) {
	o.MemoryOverCommit = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClassBatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Engine != nil {
		toSerialize["engine"] = o.Engine
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	if o.InstanceType != nil {
		toSerialize["instanceType"] = o.InstanceType
	}
	if o.CpuOverCommit != nil {
		toSerialize["cpuOverCommit"] = o.CpuOverCommit
	}
	if o.MemoryOverCommit != nil {
		toSerialize["memoryOverCommit"] = o.MemoryOverCommit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClassBatch) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Engine           *string  `json:"engine,omitempty"`
		Mode             *string  `json:"mode,omitempty"`
		Component        *string  `json:"component,omitempty"`
		InstanceType     *string  `json:"instanceType,omitempty"`
		CpuOverCommit    *float64 `json:"cpuOverCommit,omitempty"`
		MemoryOverCommit *float64 `json:"memoryOverCommit,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"engine", "mode", "component", "instanceType", "cpuOverCommit", "memoryOverCommit"})
	} else {
		return err
	}
	o.Engine = all.Engine
	o.Mode = all.Mode
	o.Component = all.Component
	o.InstanceType = all.InstanceType
	o.CpuOverCommit = all.CpuOverCommit
	o.MemoryOverCommit = all.MemoryOverCommit

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
