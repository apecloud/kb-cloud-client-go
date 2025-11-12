// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type Class struct {
	Engine     string   `json:"engine"`
	Code       *string  `json:"code,omitempty"`
	CodeShort  *string  `json:"codeShort,omitempty"`
	Mode       string   `json:"mode"`
	Cpu        *float64 `json:"cpu,omitempty"`
	CpuRequest *float64 `json:"cpuRequest,omitempty"`
	CpuLimit   *float64 `json:"cpuLimit,omitempty"`
	// memory, unit is Gi
	Memory *float64 `json:"memory,omitempty"`
	// memory request, unit is Gi
	MemoryRequest *float64 `json:"memoryRequest,omitempty"`
	// memory limit, unit is Gi
	MemoryLimit *float64 `json:"memoryLimit,omitempty"`
	Component   *string  `json:"component,omitempty"`
	ClassType   string   `json:"classType"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClass instantiates a new Class object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClass(engine string, mode string, classType string) *Class {
	this := Class{}
	this.Engine = engine
	this.Mode = mode
	this.ClassType = classType
	return &this
}

// NewClassWithDefaults instantiates a new Class object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClassWithDefaults() *Class {
	this := Class{}
	return &this
}

// GetEngine returns the Engine field value.
func (o *Class) GetEngine() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Engine
}

// GetEngineOk returns a tuple with the Engine field value
// and a boolean to check if the value has been set.
func (o *Class) GetEngineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Engine, true
}

// SetEngine sets field value.
func (o *Class) SetEngine(v string) {
	o.Engine = v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Class) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Class) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Class) HasCode() bool {
	return o != nil && o.Code != nil
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *Class) SetCode(v string) {
	o.Code = &v
}

// GetCodeShort returns the CodeShort field value if set, zero value otherwise.
func (o *Class) GetCodeShort() string {
	if o == nil || o.CodeShort == nil {
		var ret string
		return ret
	}
	return *o.CodeShort
}

// GetCodeShortOk returns a tuple with the CodeShort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Class) GetCodeShortOk() (*string, bool) {
	if o == nil || o.CodeShort == nil {
		return nil, false
	}
	return o.CodeShort, true
}

// HasCodeShort returns a boolean if a field has been set.
func (o *Class) HasCodeShort() bool {
	return o != nil && o.CodeShort != nil
}

// SetCodeShort gets a reference to the given string and assigns it to the CodeShort field.
func (o *Class) SetCodeShort(v string) {
	o.CodeShort = &v
}

// GetMode returns the Mode field value.
func (o *Class) GetMode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *Class) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *Class) SetMode(v string) {
	o.Mode = v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *Class) GetCpu() float64 {
	if o == nil || o.Cpu == nil {
		var ret float64
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Class) GetCpuOk() (*float64, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *Class) HasCpu() bool {
	return o != nil && o.Cpu != nil
}

// SetCpu gets a reference to the given float64 and assigns it to the Cpu field.
func (o *Class) SetCpu(v float64) {
	o.Cpu = &v
}

// GetCpuRequest returns the CpuRequest field value if set, zero value otherwise.
func (o *Class) GetCpuRequest() float64 {
	if o == nil || o.CpuRequest == nil {
		var ret float64
		return ret
	}
	return *o.CpuRequest
}

// GetCpuRequestOk returns a tuple with the CpuRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Class) GetCpuRequestOk() (*float64, bool) {
	if o == nil || o.CpuRequest == nil {
		return nil, false
	}
	return o.CpuRequest, true
}

// HasCpuRequest returns a boolean if a field has been set.
func (o *Class) HasCpuRequest() bool {
	return o != nil && o.CpuRequest != nil
}

// SetCpuRequest gets a reference to the given float64 and assigns it to the CpuRequest field.
func (o *Class) SetCpuRequest(v float64) {
	o.CpuRequest = &v
}

// GetCpuLimit returns the CpuLimit field value if set, zero value otherwise.
func (o *Class) GetCpuLimit() float64 {
	if o == nil || o.CpuLimit == nil {
		var ret float64
		return ret
	}
	return *o.CpuLimit
}

// GetCpuLimitOk returns a tuple with the CpuLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Class) GetCpuLimitOk() (*float64, bool) {
	if o == nil || o.CpuLimit == nil {
		return nil, false
	}
	return o.CpuLimit, true
}

// HasCpuLimit returns a boolean if a field has been set.
func (o *Class) HasCpuLimit() bool {
	return o != nil && o.CpuLimit != nil
}

// SetCpuLimit gets a reference to the given float64 and assigns it to the CpuLimit field.
func (o *Class) SetCpuLimit(v float64) {
	o.CpuLimit = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *Class) GetMemory() float64 {
	if o == nil || o.Memory == nil {
		var ret float64
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Class) GetMemoryOk() (*float64, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *Class) HasMemory() bool {
	return o != nil && o.Memory != nil
}

// SetMemory gets a reference to the given float64 and assigns it to the Memory field.
func (o *Class) SetMemory(v float64) {
	o.Memory = &v
}

// GetMemoryRequest returns the MemoryRequest field value if set, zero value otherwise.
func (o *Class) GetMemoryRequest() float64 {
	if o == nil || o.MemoryRequest == nil {
		var ret float64
		return ret
	}
	return *o.MemoryRequest
}

// GetMemoryRequestOk returns a tuple with the MemoryRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Class) GetMemoryRequestOk() (*float64, bool) {
	if o == nil || o.MemoryRequest == nil {
		return nil, false
	}
	return o.MemoryRequest, true
}

// HasMemoryRequest returns a boolean if a field has been set.
func (o *Class) HasMemoryRequest() bool {
	return o != nil && o.MemoryRequest != nil
}

// SetMemoryRequest gets a reference to the given float64 and assigns it to the MemoryRequest field.
func (o *Class) SetMemoryRequest(v float64) {
	o.MemoryRequest = &v
}

// GetMemoryLimit returns the MemoryLimit field value if set, zero value otherwise.
func (o *Class) GetMemoryLimit() float64 {
	if o == nil || o.MemoryLimit == nil {
		var ret float64
		return ret
	}
	return *o.MemoryLimit
}

// GetMemoryLimitOk returns a tuple with the MemoryLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Class) GetMemoryLimitOk() (*float64, bool) {
	if o == nil || o.MemoryLimit == nil {
		return nil, false
	}
	return o.MemoryLimit, true
}

// HasMemoryLimit returns a boolean if a field has been set.
func (o *Class) HasMemoryLimit() bool {
	return o != nil && o.MemoryLimit != nil
}

// SetMemoryLimit gets a reference to the given float64 and assigns it to the MemoryLimit field.
func (o *Class) SetMemoryLimit(v float64) {
	o.MemoryLimit = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *Class) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Class) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *Class) HasComponent() bool {
	return o != nil && o.Component != nil
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *Class) SetComponent(v string) {
	o.Component = &v
}

// GetClassType returns the ClassType field value.
func (o *Class) GetClassType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClassType
}

// GetClassTypeOk returns a tuple with the ClassType field value
// and a boolean to check if the value has been set.
func (o *Class) GetClassTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassType, true
}

// SetClassType sets field value.
func (o *Class) SetClassType(v string) {
	o.ClassType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o Class) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["engine"] = o.Engine
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.CodeShort != nil {
		toSerialize["codeShort"] = o.CodeShort
	}
	toSerialize["mode"] = o.Mode
	if o.Cpu != nil {
		toSerialize["cpu"] = o.Cpu
	}
	if o.CpuRequest != nil {
		toSerialize["cpuRequest"] = o.CpuRequest
	}
	if o.CpuLimit != nil {
		toSerialize["cpuLimit"] = o.CpuLimit
	}
	if o.Memory != nil {
		toSerialize["memory"] = o.Memory
	}
	if o.MemoryRequest != nil {
		toSerialize["memoryRequest"] = o.MemoryRequest
	}
	if o.MemoryLimit != nil {
		toSerialize["memoryLimit"] = o.MemoryLimit
	}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	toSerialize["classType"] = o.ClassType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Class) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Engine        *string  `json:"engine"`
		Code          *string  `json:"code,omitempty"`
		CodeShort     *string  `json:"codeShort,omitempty"`
		Mode          *string  `json:"mode"`
		Cpu           *float64 `json:"cpu,omitempty"`
		CpuRequest    *float64 `json:"cpuRequest,omitempty"`
		CpuLimit      *float64 `json:"cpuLimit,omitempty"`
		Memory        *float64 `json:"memory,omitempty"`
		MemoryRequest *float64 `json:"memoryRequest,omitempty"`
		MemoryLimit   *float64 `json:"memoryLimit,omitempty"`
		Component     *string  `json:"component,omitempty"`
		ClassType     *string  `json:"classType"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Engine == nil {
		return fmt.Errorf("required field engine missing")
	}
	if all.Mode == nil {
		return fmt.Errorf("required field mode missing")
	}
	if all.ClassType == nil {
		return fmt.Errorf("required field classType missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"engine", "code", "codeShort", "mode", "cpu", "cpuRequest", "cpuLimit", "memory", "memoryRequest", "memoryLimit", "component", "classType"})
	} else {
		return err
	}
	o.Engine = *all.Engine
	o.Code = all.Code
	o.CodeShort = all.CodeShort
	o.Mode = *all.Mode
	o.Cpu = all.Cpu
	o.CpuRequest = all.CpuRequest
	o.CpuLimit = all.CpuLimit
	o.Memory = all.Memory
	o.MemoryRequest = all.MemoryRequest
	o.MemoryLimit = all.MemoryLimit
	o.Component = all.Component
	o.ClassType = *all.ClassType

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
