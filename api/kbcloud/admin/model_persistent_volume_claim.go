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
	Capacity string `json:"capacity"`
	// access mode of the PVC
	AccessMode string `json:"accessMode"`
	// the persistentvolume of the PVC
	VolumeName string `json:"volumeName"`
	// the node of the PVC
	Node string `json:"node"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPersistentVolumeClaim instantiates a new PersistentVolumeClaim object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPersistentVolumeClaim(nameSpace string, name string, phase string, capacity string, accessMode string, volumeName string, node string) *PersistentVolumeClaim {
	this := PersistentVolumeClaim{}
	this.NameSpace = nameSpace
	this.Name = name
	this.Phase = phase
	this.Capacity = capacity
	this.AccessMode = accessMode
	this.VolumeName = volumeName
	this.Node = node
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

// GetCapacity returns the Capacity field value.
func (o *PersistentVolumeClaim) GetCapacity() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value
// and a boolean to check if the value has been set.
func (o *PersistentVolumeClaim) GetCapacityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Capacity, true
}

// SetCapacity sets field value.
func (o *PersistentVolumeClaim) SetCapacity(v string) {
	o.Capacity = v
}

// GetAccessMode returns the AccessMode field value.
func (o *PersistentVolumeClaim) GetAccessMode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccessMode
}

// GetAccessModeOk returns a tuple with the AccessMode field value
// and a boolean to check if the value has been set.
func (o *PersistentVolumeClaim) GetAccessModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessMode, true
}

// SetAccessMode sets field value.
func (o *PersistentVolumeClaim) SetAccessMode(v string) {
	o.AccessMode = v
}

// GetVolumeName returns the VolumeName field value.
func (o *PersistentVolumeClaim) GetVolumeName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.VolumeName
}

// GetVolumeNameOk returns a tuple with the VolumeName field value
// and a boolean to check if the value has been set.
func (o *PersistentVolumeClaim) GetVolumeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VolumeName, true
}

// SetVolumeName sets field value.
func (o *PersistentVolumeClaim) SetVolumeName(v string) {
	o.VolumeName = v
}

// GetNode returns the Node field value.
func (o *PersistentVolumeClaim) GetNode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Node
}

// GetNodeOk returns a tuple with the Node field value
// and a boolean to check if the value has been set.
func (o *PersistentVolumeClaim) GetNodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Node, true
}

// SetNode sets field value.
func (o *PersistentVolumeClaim) SetNode(v string) {
	o.Node = v
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
	toSerialize["capacity"] = o.Capacity
	toSerialize["accessMode"] = o.AccessMode
	toSerialize["volumeName"] = o.VolumeName
	toSerialize["node"] = o.Node

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PersistentVolumeClaim) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NameSpace  *string `json:"nameSpace"`
		Name       *string `json:"name"`
		Phase      *string `json:"phase"`
		Capacity   *string `json:"capacity"`
		AccessMode *string `json:"accessMode"`
		VolumeName *string `json:"volumeName"`
		Node       *string `json:"node"`
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
	if all.Capacity == nil {
		return fmt.Errorf("required field capacity missing")
	}
	if all.AccessMode == nil {
		return fmt.Errorf("required field accessMode missing")
	}
	if all.VolumeName == nil {
		return fmt.Errorf("required field volumeName missing")
	}
	if all.Node == nil {
		return fmt.Errorf("required field node missing")
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
	o.Capacity = *all.Capacity
	o.AccessMode = *all.AccessMode
	o.VolumeName = *all.VolumeName
	o.Node = *all.Node

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
