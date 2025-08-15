// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ComponentItemCreate ComponentItem is the information of a component
type ComponentItemCreate struct {
	// component type, refer to componentDef and support NamePrefix
	Component string `json:"component"`
	// The number of components, if often used as shards number
	CompNum *int32 `json:"compNum,omitempty"`
	// The number of replicas, for standalone mode, the replicas is 1, for raftGroup mode, the default replicas is 3.
	Replicas  *int32  `json:"replicas,omitempty"`
	ClassCode *string `json:"classCode,omitempty"`
	// CPU cores.
	Cpu *float64 `json:"cpu,omitempty"`
	// Memory, the unit is Gi.
	Memory *float64 `json:"memory,omitempty"`
	// StorageClass name
	StorageClass *string               `json:"storageClass,omitempty"`
	Volumes      []ComponentVolumeItem `json:"volumes,omitempty"`
	// The name of the secret that contains the system account credentials
	SystemAccountSecretName *string `json:"systemAccountSecretName,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewComponentItemCreate instantiates a new ComponentItemCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewComponentItemCreate(component string) *ComponentItemCreate {
	this := ComponentItemCreate{}
	this.Component = component
	return &this
}

// NewComponentItemCreateWithDefaults instantiates a new ComponentItemCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewComponentItemCreateWithDefaults() *ComponentItemCreate {
	this := ComponentItemCreate{}
	return &this
}

// GetComponent returns the Component field value.
func (o *ComponentItemCreate) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *ComponentItemCreate) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value.
func (o *ComponentItemCreate) SetComponent(v string) {
	o.Component = v
}

// GetCompNum returns the CompNum field value if set, zero value otherwise.
func (o *ComponentItemCreate) GetCompNum() int32 {
	if o == nil || o.CompNum == nil {
		var ret int32
		return ret
	}
	return *o.CompNum
}

// GetCompNumOk returns a tuple with the CompNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentItemCreate) GetCompNumOk() (*int32, bool) {
	if o == nil || o.CompNum == nil {
		return nil, false
	}
	return o.CompNum, true
}

// HasCompNum returns a boolean if a field has been set.
func (o *ComponentItemCreate) HasCompNum() bool {
	return o != nil && o.CompNum != nil
}

// SetCompNum gets a reference to the given int32 and assigns it to the CompNum field.
func (o *ComponentItemCreate) SetCompNum(v int32) {
	o.CompNum = &v
}

// GetReplicas returns the Replicas field value if set, zero value otherwise.
func (o *ComponentItemCreate) GetReplicas() int32 {
	if o == nil || o.Replicas == nil {
		var ret int32
		return ret
	}
	return *o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentItemCreate) GetReplicasOk() (*int32, bool) {
	if o == nil || o.Replicas == nil {
		return nil, false
	}
	return o.Replicas, true
}

// HasReplicas returns a boolean if a field has been set.
func (o *ComponentItemCreate) HasReplicas() bool {
	return o != nil && o.Replicas != nil
}

// SetReplicas gets a reference to the given int32 and assigns it to the Replicas field.
func (o *ComponentItemCreate) SetReplicas(v int32) {
	o.Replicas = &v
}

// GetClassCode returns the ClassCode field value if set, zero value otherwise.
func (o *ComponentItemCreate) GetClassCode() string {
	if o == nil || o.ClassCode == nil {
		var ret string
		return ret
	}
	return *o.ClassCode
}

// GetClassCodeOk returns a tuple with the ClassCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentItemCreate) GetClassCodeOk() (*string, bool) {
	if o == nil || o.ClassCode == nil {
		return nil, false
	}
	return o.ClassCode, true
}

// HasClassCode returns a boolean if a field has been set.
func (o *ComponentItemCreate) HasClassCode() bool {
	return o != nil && o.ClassCode != nil
}

// SetClassCode gets a reference to the given string and assigns it to the ClassCode field.
func (o *ComponentItemCreate) SetClassCode(v string) {
	o.ClassCode = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *ComponentItemCreate) GetCpu() float64 {
	if o == nil || o.Cpu == nil {
		var ret float64
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentItemCreate) GetCpuOk() (*float64, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *ComponentItemCreate) HasCpu() bool {
	return o != nil && o.Cpu != nil
}

// SetCpu gets a reference to the given float64 and assigns it to the Cpu field.
func (o *ComponentItemCreate) SetCpu(v float64) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *ComponentItemCreate) GetMemory() float64 {
	if o == nil || o.Memory == nil {
		var ret float64
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentItemCreate) GetMemoryOk() (*float64, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *ComponentItemCreate) HasMemory() bool {
	return o != nil && o.Memory != nil
}

// SetMemory gets a reference to the given float64 and assigns it to the Memory field.
func (o *ComponentItemCreate) SetMemory(v float64) {
	o.Memory = &v
}

// GetStorageClass returns the StorageClass field value if set, zero value otherwise.
func (o *ComponentItemCreate) GetStorageClass() string {
	if o == nil || o.StorageClass == nil {
		var ret string
		return ret
	}
	return *o.StorageClass
}

// GetStorageClassOk returns a tuple with the StorageClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentItemCreate) GetStorageClassOk() (*string, bool) {
	if o == nil || o.StorageClass == nil {
		return nil, false
	}
	return o.StorageClass, true
}

// HasStorageClass returns a boolean if a field has been set.
func (o *ComponentItemCreate) HasStorageClass() bool {
	return o != nil && o.StorageClass != nil
}

// SetStorageClass gets a reference to the given string and assigns it to the StorageClass field.
func (o *ComponentItemCreate) SetStorageClass(v string) {
	o.StorageClass = &v
}

// GetVolumes returns the Volumes field value if set, zero value otherwise.
func (o *ComponentItemCreate) GetVolumes() []ComponentVolumeItem {
	if o == nil || o.Volumes == nil {
		var ret []ComponentVolumeItem
		return ret
	}
	return o.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentItemCreate) GetVolumesOk() (*[]ComponentVolumeItem, bool) {
	if o == nil || o.Volumes == nil {
		return nil, false
	}
	return &o.Volumes, true
}

// HasVolumes returns a boolean if a field has been set.
func (o *ComponentItemCreate) HasVolumes() bool {
	return o != nil && o.Volumes != nil
}

// SetVolumes gets a reference to the given []ComponentVolumeItem and assigns it to the Volumes field.
func (o *ComponentItemCreate) SetVolumes(v []ComponentVolumeItem) {
	o.Volumes = v
}

// GetSystemAccountSecretName returns the SystemAccountSecretName field value if set, zero value otherwise.
func (o *ComponentItemCreate) GetSystemAccountSecretName() string {
	if o == nil || o.SystemAccountSecretName == nil {
		var ret string
		return ret
	}
	return *o.SystemAccountSecretName
}

// GetSystemAccountSecretNameOk returns a tuple with the SystemAccountSecretName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentItemCreate) GetSystemAccountSecretNameOk() (*string, bool) {
	if o == nil || o.SystemAccountSecretName == nil {
		return nil, false
	}
	return o.SystemAccountSecretName, true
}

// HasSystemAccountSecretName returns a boolean if a field has been set.
func (o *ComponentItemCreate) HasSystemAccountSecretName() bool {
	return o != nil && o.SystemAccountSecretName != nil
}

// SetSystemAccountSecretName gets a reference to the given string and assigns it to the SystemAccountSecretName field.
func (o *ComponentItemCreate) SetSystemAccountSecretName(v string) {
	o.SystemAccountSecretName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ComponentItemCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["component"] = o.Component
	if o.CompNum != nil {
		toSerialize["compNum"] = o.CompNum
	}
	if o.Replicas != nil {
		toSerialize["replicas"] = o.Replicas
	}
	if o.ClassCode != nil {
		toSerialize["classCode"] = o.ClassCode
	}
	if o.Cpu != nil {
		toSerialize["cpu"] = o.Cpu
	}
	if o.Memory != nil {
		toSerialize["memory"] = o.Memory
	}
	if o.StorageClass != nil {
		toSerialize["storageClass"] = o.StorageClass
	}
	if o.Volumes != nil {
		toSerialize["volumes"] = o.Volumes
	}
	if o.SystemAccountSecretName != nil {
		toSerialize["systemAccountSecretName"] = o.SystemAccountSecretName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ComponentItemCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Component               *string               `json:"component"`
		CompNum                 *int32                `json:"compNum,omitempty"`
		Replicas                *int32                `json:"replicas,omitempty"`
		ClassCode               *string               `json:"classCode,omitempty"`
		Cpu                     *float64              `json:"cpu,omitempty"`
		Memory                  *float64              `json:"memory,omitempty"`
		StorageClass            *string               `json:"storageClass,omitempty"`
		Volumes                 []ComponentVolumeItem `json:"volumes,omitempty"`
		SystemAccountSecretName *string               `json:"systemAccountSecretName,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Component == nil {
		return fmt.Errorf("required field component missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"component", "compNum", "replicas", "classCode", "cpu", "memory", "storageClass", "volumes", "systemAccountSecretName"})
	} else {
		return err
	}
	o.Component = *all.Component
	o.CompNum = all.CompNum
	o.Replicas = all.Replicas
	o.ClassCode = all.ClassCode
	o.Cpu = all.Cpu
	o.Memory = all.Memory
	o.StorageClass = all.StorageClass
	o.Volumes = all.Volumes
	o.SystemAccountSecretName = all.SystemAccountSecretName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
