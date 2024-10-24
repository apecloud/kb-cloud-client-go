// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ModeComponent struct {
	// NODESCRIPTION Component
	Component string `json:"component"`
	// NODESCRIPTION Shards
	Shards *IntegerOption `json:"shards,omitempty"`
	// NODESCRIPTION Replicas
	Replicas IntegerOption `json:"replicas"`
	// NODESCRIPTION Cpu
	Cpu FloatOption `json:"cpu"`
	// NODESCRIPTION Memory
	Memory FloatOption `json:"memory"`
	// NODESCRIPTION HideEnpoints
	HideEnpoints bool `json:"hideEnpoints"`
	// NODESCRIPTION HideOnCreate
	HideOnCreate bool `json:"hideOnCreate"`
	// NODESCRIPTION Storages
	Storages []StorageOption `json:"storages"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewModeComponent instantiates a new ModeComponent object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewModeComponent(component string, replicas IntegerOption, cpu FloatOption, memory FloatOption, hideEnpoints bool, hideOnCreate bool, storages []StorageOption) *ModeComponent {
	this := ModeComponent{}
	this.Component = component
	this.Replicas = replicas
	this.Cpu = cpu
	this.Memory = memory
	this.HideEnpoints = hideEnpoints
	this.HideOnCreate = hideOnCreate
	this.Storages = storages
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

// GetReplicas returns the Replicas field value.
func (o *ModeComponent) GetReplicas() IntegerOption {
	if o == nil {
		var ret IntegerOption
		return ret
	}
	return o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value
// and a boolean to check if the value has been set.
func (o *ModeComponent) GetReplicasOk() (*IntegerOption, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Replicas, true
}

// SetReplicas sets field value.
func (o *ModeComponent) SetReplicas(v IntegerOption) {
	o.Replicas = v
}

// GetCpu returns the Cpu field value.
func (o *ModeComponent) GetCpu() FloatOption {
	if o == nil {
		var ret FloatOption
		return ret
	}
	return o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value
// and a boolean to check if the value has been set.
func (o *ModeComponent) GetCpuOk() (*FloatOption, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cpu, true
}

// SetCpu sets field value.
func (o *ModeComponent) SetCpu(v FloatOption) {
	o.Cpu = v
}

// GetMemory returns the Memory field value.
func (o *ModeComponent) GetMemory() FloatOption {
	if o == nil {
		var ret FloatOption
		return ret
	}
	return o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value
// and a boolean to check if the value has been set.
func (o *ModeComponent) GetMemoryOk() (*FloatOption, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Memory, true
}

// SetMemory sets field value.
func (o *ModeComponent) SetMemory(v FloatOption) {
	o.Memory = v
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

// GetStorages returns the Storages field value.
func (o *ModeComponent) GetStorages() []StorageOption {
	if o == nil {
		var ret []StorageOption
		return ret
	}
	return o.Storages
}

// GetStoragesOk returns a tuple with the Storages field value
// and a boolean to check if the value has been set.
func (o *ModeComponent) GetStoragesOk() (*[]StorageOption, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Storages, true
}

// SetStorages sets field value.
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
	toSerialize["replicas"] = o.Replicas
	toSerialize["cpu"] = o.Cpu
	toSerialize["memory"] = o.Memory
	toSerialize["hideEnpoints"] = o.HideEnpoints
	toSerialize["hideOnCreate"] = o.HideOnCreate
	toSerialize["storages"] = o.Storages

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ModeComponent) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Component    *string          `json:"component"`
		Shards       *IntegerOption   `json:"shards,omitempty"`
		Replicas     *IntegerOption   `json:"replicas"`
		Cpu          *FloatOption     `json:"cpu"`
		Memory       *FloatOption     `json:"memory"`
		HideEnpoints *bool            `json:"hideEnpoints"`
		HideOnCreate *bool            `json:"hideOnCreate"`
		Storages     *[]StorageOption `json:"storages"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Component == nil {
		return fmt.Errorf("required field component missing")
	}
	if all.Replicas == nil {
		return fmt.Errorf("required field replicas missing")
	}
	if all.Cpu == nil {
		return fmt.Errorf("required field cpu missing")
	}
	if all.Memory == nil {
		return fmt.Errorf("required field memory missing")
	}
	if all.HideEnpoints == nil {
		return fmt.Errorf("required field hideEnpoints missing")
	}
	if all.HideOnCreate == nil {
		return fmt.Errorf("required field hideOnCreate missing")
	}
	if all.Storages == nil {
		return fmt.Errorf("required field storages missing")
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
	if all.Replicas.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Replicas = *all.Replicas
	if all.Cpu.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Cpu = *all.Cpu
	if all.Memory.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Memory = *all.Memory
	o.HideEnpoints = *all.HideEnpoints
	o.HideOnCreate = *all.HideOnCreate
	o.Storages = *all.Storages

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
