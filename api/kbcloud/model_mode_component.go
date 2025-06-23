// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ModeComponent struct {
	Component    string          `json:"component"`
	Shards       *IntegerOption  `json:"shards,omitempty"`
	Replicas     *IntegerOption  `json:"replicas,omitempty"`
	Cpu          *FloatOption    `json:"cpu,omitempty"`
	Memory       *FloatOption    `json:"memory,omitempty"`
	HideEnpoints bool            `json:"hideEnpoints"`
	HideOnCreate bool            `json:"hideOnCreate"`
	Storages     []StorageOption `json:"storages,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewModeComponent instantiates a new ModeComponent object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewModeComponent(component string, hideEnpoints bool, hideOnCreate bool) *ModeComponent {
	this := ModeComponent{}
	this.Component = component
	this.HideEnpoints = hideEnpoints
	this.HideOnCreate = hideOnCreate
	return &this
}

// NewModeComponentWithDefaults instantiates a new ModeComponent object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewModeComponentWithDefaults() *ModeComponent {
	this := ModeComponent{}
	return &this
}

// GetComponent returns the Component field value.
func (o *ModeComponent) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *ModeComponent) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value.
func (o *ModeComponent) SetComponent(v string) {
	o.Component = v
}

// GetShards returns the Shards field value if set, zero value otherwise.
func (o *ModeComponent) GetShards() IntegerOption {
	if o == nil || o.Shards == nil {
		var ret IntegerOption
		return ret
	}
	return *o.Shards
}

// GetShardsOk returns a tuple with the Shards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeComponent) GetShardsOk() (*IntegerOption, bool) {
	if o == nil || o.Shards == nil {
		return nil, false
	}
	return o.Shards, true
}

// HasShards returns a boolean if a field has been set.
func (o *ModeComponent) HasShards() bool {
	return o != nil && o.Shards != nil
}

// SetShards gets a reference to the given IntegerOption and assigns it to the Shards field.
func (o *ModeComponent) SetShards(v IntegerOption) {
	o.Shards = &v
}

// GetReplicas returns the Replicas field value if set, zero value otherwise.
func (o *ModeComponent) GetReplicas() IntegerOption {
	if o == nil || o.Replicas == nil {
		var ret IntegerOption
		return ret
	}
	return *o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeComponent) GetReplicasOk() (*IntegerOption, bool) {
	if o == nil || o.Replicas == nil {
		return nil, false
	}
	return o.Replicas, true
}

// HasReplicas returns a boolean if a field has been set.
func (o *ModeComponent) HasReplicas() bool {
	return o != nil && o.Replicas != nil
}

// SetReplicas gets a reference to the given IntegerOption and assigns it to the Replicas field.
func (o *ModeComponent) SetReplicas(v IntegerOption) {
	o.Replicas = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *ModeComponent) GetCpu() FloatOption {
	if o == nil || o.Cpu == nil {
		var ret FloatOption
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeComponent) GetCpuOk() (*FloatOption, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *ModeComponent) HasCpu() bool {
	return o != nil && o.Cpu != nil
}

// SetCpu gets a reference to the given FloatOption and assigns it to the Cpu field.
func (o *ModeComponent) SetCpu(v FloatOption) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *ModeComponent) GetMemory() FloatOption {
	if o == nil || o.Memory == nil {
		var ret FloatOption
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeComponent) GetMemoryOk() (*FloatOption, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *ModeComponent) HasMemory() bool {
	return o != nil && o.Memory != nil
}

// SetMemory gets a reference to the given FloatOption and assigns it to the Memory field.
func (o *ModeComponent) SetMemory(v FloatOption) {
	o.Memory = &v
}

// GetHideEnpoints returns the HideEnpoints field value.
func (o *ModeComponent) GetHideEnpoints() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.HideEnpoints
}

// GetHideEnpointsOk returns a tuple with the HideEnpoints field value
// and a boolean to check if the value has been set.
func (o *ModeComponent) GetHideEnpointsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HideEnpoints, true
}

// SetHideEnpoints sets field value.
func (o *ModeComponent) SetHideEnpoints(v bool) {
	o.HideEnpoints = v
}

// GetHideOnCreate returns the HideOnCreate field value.
func (o *ModeComponent) GetHideOnCreate() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.HideOnCreate
}

// GetHideOnCreateOk returns a tuple with the HideOnCreate field value
// and a boolean to check if the value has been set.
func (o *ModeComponent) GetHideOnCreateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HideOnCreate, true
}

// SetHideOnCreate sets field value.
func (o *ModeComponent) SetHideOnCreate(v bool) {
	o.HideOnCreate = v
}

// GetStorages returns the Storages field value if set, zero value otherwise.
func (o *ModeComponent) GetStorages() []StorageOption {
	if o == nil || o.Storages == nil {
		var ret []StorageOption
		return ret
	}
	return o.Storages
}

// GetStoragesOk returns a tuple with the Storages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeComponent) GetStoragesOk() (*[]StorageOption, bool) {
	if o == nil || o.Storages == nil {
		return nil, false
	}
	return &o.Storages, true
}

// HasStorages returns a boolean if a field has been set.
func (o *ModeComponent) HasStorages() bool {
	return o != nil && o.Storages != nil
}

// SetStorages gets a reference to the given []StorageOption and assigns it to the Storages field.
func (o *ModeComponent) SetStorages(v []StorageOption) {
	o.Storages = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ModeComponent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["component"] = o.Component
	if o.Shards != nil {
		toSerialize["shards"] = o.Shards
	}
	if o.Replicas != nil {
		toSerialize["replicas"] = o.Replicas
	}
	if o.Cpu != nil {
		toSerialize["cpu"] = o.Cpu
	}
	if o.Memory != nil {
		toSerialize["memory"] = o.Memory
	}
	toSerialize["hideEnpoints"] = o.HideEnpoints
	toSerialize["hideOnCreate"] = o.HideOnCreate
	if o.Storages != nil {
		toSerialize["storages"] = o.Storages
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ModeComponent) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Component    *string         `json:"component"`
		Shards       *IntegerOption  `json:"shards,omitempty"`
		Replicas     *IntegerOption  `json:"replicas,omitempty"`
		Cpu          *FloatOption    `json:"cpu,omitempty"`
		Memory       *FloatOption    `json:"memory,omitempty"`
		HideEnpoints *bool           `json:"hideEnpoints"`
		HideOnCreate *bool           `json:"hideOnCreate"`
		Storages     []StorageOption `json:"storages,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Component == nil {
		return fmt.Errorf("required field component missing")
	}
	if all.HideEnpoints == nil {
		return fmt.Errorf("required field hideEnpoints missing")
	}
	if all.HideOnCreate == nil {
		return fmt.Errorf("required field hideOnCreate missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"component", "shards", "replicas", "cpu", "memory", "hideEnpoints", "hideOnCreate", "storages"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Component = *all.Component
	if all.Shards != nil && all.Shards.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Shards = all.Shards
	if all.Replicas != nil && all.Replicas.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Replicas = all.Replicas
	if all.Cpu != nil && all.Cpu.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Cpu = all.Cpu
	if all.Memory != nil && all.Memory.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Memory = all.Memory
	o.HideEnpoints = *all.HideEnpoints
	o.HideOnCreate = *all.HideOnCreate
	o.Storages = all.Storages

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
