// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
)

// PersistentVolumeClaim persistentVolumeClaim provides detailed information about the PVC .
type PersistentVolumeClaim struct {
	// the namespace of the PVC
	NameSpace string `json:"nameSpace"`
	// the name of the PVC
	Name string `json:"name"`
	// the phase of the PVC
	Phase string `json:"phase"`
	// the capacity of the PVC
	Capacity common.NullableString `json:"capacity,omitempty"`
	// access mode of the PVC
	AccessMode common.NullableString `json:"accessMode,omitempty"`
	// the persistentvolume of the PVC
	VolumeName common.NullableString `json:"volumeName,omitempty"`
	// the node of the PVC
	Node common.NullableString `json:"node,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPersistentVolumeClaim instantiates a new PersistentVolumeClaim object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPersistentVolumeClaim(nameSpace string, name string, phase string) *PersistentVolumeClaim {
	this := PersistentVolumeClaim{}
	this.NameSpace = nameSpace
	this.Name = name
	this.Phase = phase
	return &this
}

// NewPersistentVolumeClaimWithDefaults instantiates a new PersistentVolumeClaim object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPersistentVolumeClaimWithDefaults() *PersistentVolumeClaim {
	this := PersistentVolumeClaim{}
	return &this
}

// GetNameSpace returns the NameSpace field value.
func (o *PersistentVolumeClaim) GetNameSpace() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.NameSpace
}

// GetNameSpaceOk returns a tuple with the NameSpace field value
// and a boolean to check if the value has been set.
func (o *PersistentVolumeClaim) GetNameSpaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NameSpace, true
}

// SetNameSpace sets field value.
func (o *PersistentVolumeClaim) SetNameSpace(v string) {
	o.NameSpace = v
}

// GetName returns the Name field value.
func (o *PersistentVolumeClaim) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PersistentVolumeClaim) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *PersistentVolumeClaim) SetName(v string) {
	o.Name = v
}

// GetPhase returns the Phase field value.
func (o *PersistentVolumeClaim) GetPhase() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value
// and a boolean to check if the value has been set.
func (o *PersistentVolumeClaim) GetPhaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Phase, true
}

// SetPhase sets field value.
func (o *PersistentVolumeClaim) SetPhase(v string) {
	o.Phase = v
}

// GetCapacity returns the Capacity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PersistentVolumeClaim) GetCapacity() string {
	if o == nil || o.Capacity.Get() == nil {
		var ret string
		return ret
	}
	return *o.Capacity.Get()
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PersistentVolumeClaim) GetCapacityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Capacity.Get(), o.Capacity.IsSet()
}

// HasCapacity returns a boolean if a field has been set.
func (o *PersistentVolumeClaim) HasCapacity() bool {
	return o != nil && o.Capacity.IsSet()
}

// SetCapacity gets a reference to the given common.NullableString and assigns it to the Capacity field.
func (o *PersistentVolumeClaim) SetCapacity(v string) {
	o.Capacity.Set(&v)
}

// SetCapacityNil sets the value for Capacity to be an explicit nil.
func (o *PersistentVolumeClaim) SetCapacityNil() {
	o.Capacity.Set(nil)
}

// UnsetCapacity ensures that no value is present for Capacity, not even an explicit nil.
func (o *PersistentVolumeClaim) UnsetCapacity() {
	o.Capacity.Unset()
}

// GetAccessMode returns the AccessMode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PersistentVolumeClaim) GetAccessMode() string {
	if o == nil || o.AccessMode.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccessMode.Get()
}

// GetAccessModeOk returns a tuple with the AccessMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PersistentVolumeClaim) GetAccessModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessMode.Get(), o.AccessMode.IsSet()
}

// HasAccessMode returns a boolean if a field has been set.
func (o *PersistentVolumeClaim) HasAccessMode() bool {
	return o != nil && o.AccessMode.IsSet()
}

// SetAccessMode gets a reference to the given common.NullableString and assigns it to the AccessMode field.
func (o *PersistentVolumeClaim) SetAccessMode(v string) {
	o.AccessMode.Set(&v)
}

// SetAccessModeNil sets the value for AccessMode to be an explicit nil.
func (o *PersistentVolumeClaim) SetAccessModeNil() {
	o.AccessMode.Set(nil)
}

// UnsetAccessMode ensures that no value is present for AccessMode, not even an explicit nil.
func (o *PersistentVolumeClaim) UnsetAccessMode() {
	o.AccessMode.Unset()
}

// GetVolumeName returns the VolumeName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PersistentVolumeClaim) GetVolumeName() string {
	if o == nil || o.VolumeName.Get() == nil {
		var ret string
		return ret
	}
	return *o.VolumeName.Get()
}

// GetVolumeNameOk returns a tuple with the VolumeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PersistentVolumeClaim) GetVolumeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VolumeName.Get(), o.VolumeName.IsSet()
}

// HasVolumeName returns a boolean if a field has been set.
func (o *PersistentVolumeClaim) HasVolumeName() bool {
	return o != nil && o.VolumeName.IsSet()
}

// SetVolumeName gets a reference to the given common.NullableString and assigns it to the VolumeName field.
func (o *PersistentVolumeClaim) SetVolumeName(v string) {
	o.VolumeName.Set(&v)
}

// SetVolumeNameNil sets the value for VolumeName to be an explicit nil.
func (o *PersistentVolumeClaim) SetVolumeNameNil() {
	o.VolumeName.Set(nil)
}

// UnsetVolumeName ensures that no value is present for VolumeName, not even an explicit nil.
func (o *PersistentVolumeClaim) UnsetVolumeName() {
	o.VolumeName.Unset()
}

// GetNode returns the Node field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PersistentVolumeClaim) GetNode() string {
	if o == nil || o.Node.Get() == nil {
		var ret string
		return ret
	}
	return *o.Node.Get()
}

// GetNodeOk returns a tuple with the Node field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PersistentVolumeClaim) GetNodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Node.Get(), o.Node.IsSet()
}

// HasNode returns a boolean if a field has been set.
func (o *PersistentVolumeClaim) HasNode() bool {
	return o != nil && o.Node.IsSet()
}

// SetNode gets a reference to the given common.NullableString and assigns it to the Node field.
func (o *PersistentVolumeClaim) SetNode(v string) {
	o.Node.Set(&v)
}

// SetNodeNil sets the value for Node to be an explicit nil.
func (o *PersistentVolumeClaim) SetNodeNil() {
	o.Node.Set(nil)
}

// UnsetNode ensures that no value is present for Node, not even an explicit nil.
func (o *PersistentVolumeClaim) UnsetNode() {
	o.Node.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o PersistentVolumeClaim) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["nameSpace"] = o.NameSpace
	toSerialize["name"] = o.Name
	toSerialize["phase"] = o.Phase
	if o.Capacity.IsSet() {
		toSerialize["capacity"] = o.Capacity.Get()
	}
	if o.AccessMode.IsSet() {
		toSerialize["accessMode"] = o.AccessMode.Get()
	}
	if o.VolumeName.IsSet() {
		toSerialize["volumeName"] = o.VolumeName.Get()
	}
	if o.Node.IsSet() {
		toSerialize["node"] = o.Node.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PersistentVolumeClaim) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NameSpace  *string               `json:"nameSpace"`
		Name       *string               `json:"name"`
		Phase      *string               `json:"phase"`
		Capacity   common.NullableString `json:"capacity,omitempty"`
		AccessMode common.NullableString `json:"accessMode,omitempty"`
		VolumeName common.NullableString `json:"volumeName,omitempty"`
		Node       common.NullableString `json:"node,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.NameSpace == nil {
		return fmt.Errorf("required field nameSpace missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Phase == nil {
		return fmt.Errorf("required field phase missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"nameSpace", "name", "phase", "capacity", "accessMode", "volumeName", "node"})
	} else {
		return err
	}
	o.NameSpace = *all.NameSpace
	o.Name = *all.Name
	o.Phase = *all.Phase
	o.Capacity = all.Capacity
	o.AccessMode = all.AccessMode
	o.VolumeName = all.VolumeName
	o.Node = all.Node

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
