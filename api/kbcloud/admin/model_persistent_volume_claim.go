// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
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
	Capacity common.NullableFloat64 `json:"capacity,omitempty"`
	// the usage of the PVC
	Usage common.NullableFloat64 `json:"usage,omitempty"`
	// access mode of the PVC
	AccessMode common.NullableString `json:"accessMode,omitempty"`
	// the persistentvolume of the PVC
	VolumeName common.NullableString `json:"volumeName,omitempty"`
	// the node of the PVC
	Node common.NullableString `json:"node,omitempty"`
	// the org name of the PVC
	OrgName common.NullableString `json:"orgName,omitempty"`
	// the cluster name of the PVC
	ClusterName common.NullableString `json:"clusterName,omitempty"`
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
func (o *PersistentVolumeClaim) GetCapacity() float64 {
	if o == nil || o.Capacity.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Capacity.Get()
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PersistentVolumeClaim) GetCapacityOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Capacity.Get(), o.Capacity.IsSet()
}

// HasCapacity returns a boolean if a field has been set.
func (o *PersistentVolumeClaim) HasCapacity() bool {
	return o != nil && o.Capacity.IsSet()
}

// SetCapacity gets a reference to the given common.NullableFloat64 and assigns it to the Capacity field.
func (o *PersistentVolumeClaim) SetCapacity(v float64) {
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

// GetUsage returns the Usage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PersistentVolumeClaim) GetUsage() float64 {
	if o == nil || o.Usage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Usage.Get()
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PersistentVolumeClaim) GetUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Usage.Get(), o.Usage.IsSet()
}

// HasUsage returns a boolean if a field has been set.
func (o *PersistentVolumeClaim) HasUsage() bool {
	return o != nil && o.Usage.IsSet()
}

// SetUsage gets a reference to the given common.NullableFloat64 and assigns it to the Usage field.
func (o *PersistentVolumeClaim) SetUsage(v float64) {
	o.Usage.Set(&v)
}

// SetUsageNil sets the value for Usage to be an explicit nil.
func (o *PersistentVolumeClaim) SetUsageNil() {
	o.Usage.Set(nil)
}

// UnsetUsage ensures that no value is present for Usage, not even an explicit nil.
func (o *PersistentVolumeClaim) UnsetUsage() {
	o.Usage.Unset()
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

// GetOrgName returns the OrgName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PersistentVolumeClaim) GetOrgName() string {
	if o == nil || o.OrgName.Get() == nil {
		var ret string
		return ret
	}
	return *o.OrgName.Get()
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PersistentVolumeClaim) GetOrgNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrgName.Get(), o.OrgName.IsSet()
}

// HasOrgName returns a boolean if a field has been set.
func (o *PersistentVolumeClaim) HasOrgName() bool {
	return o != nil && o.OrgName.IsSet()
}

// SetOrgName gets a reference to the given common.NullableString and assigns it to the OrgName field.
func (o *PersistentVolumeClaim) SetOrgName(v string) {
	o.OrgName.Set(&v)
}

// SetOrgNameNil sets the value for OrgName to be an explicit nil.
func (o *PersistentVolumeClaim) SetOrgNameNil() {
	o.OrgName.Set(nil)
}

// UnsetOrgName ensures that no value is present for OrgName, not even an explicit nil.
func (o *PersistentVolumeClaim) UnsetOrgName() {
	o.OrgName.Unset()
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PersistentVolumeClaim) GetClusterName() string {
	if o == nil || o.ClusterName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClusterName.Get()
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PersistentVolumeClaim) GetClusterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterName.Get(), o.ClusterName.IsSet()
}

// HasClusterName returns a boolean if a field has been set.
func (o *PersistentVolumeClaim) HasClusterName() bool {
	return o != nil && o.ClusterName.IsSet()
}

// SetClusterName gets a reference to the given common.NullableString and assigns it to the ClusterName field.
func (o *PersistentVolumeClaim) SetClusterName(v string) {
	o.ClusterName.Set(&v)
}

// SetClusterNameNil sets the value for ClusterName to be an explicit nil.
func (o *PersistentVolumeClaim) SetClusterNameNil() {
	o.ClusterName.Set(nil)
}

// UnsetClusterName ensures that no value is present for ClusterName, not even an explicit nil.
func (o *PersistentVolumeClaim) UnsetClusterName() {
	o.ClusterName.Unset()
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
	if o.Usage.IsSet() {
		toSerialize["usage"] = o.Usage.Get()
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
	if o.OrgName.IsSet() {
		toSerialize["orgName"] = o.OrgName.Get()
	}
	if o.ClusterName.IsSet() {
		toSerialize["clusterName"] = o.ClusterName.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PersistentVolumeClaim) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NameSpace   *string                `json:"nameSpace"`
		Name        *string                `json:"name"`
		Phase       *string                `json:"phase"`
		Capacity    common.NullableFloat64 `json:"capacity,omitempty"`
		Usage       common.NullableFloat64 `json:"usage,omitempty"`
		AccessMode  common.NullableString  `json:"accessMode,omitempty"`
		VolumeName  common.NullableString  `json:"volumeName,omitempty"`
		Node        common.NullableString  `json:"node,omitempty"`
		OrgName     common.NullableString  `json:"orgName,omitempty"`
		ClusterName common.NullableString  `json:"clusterName,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
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
		common.DeleteKeys(additionalProperties, &[]string{"nameSpace", "name", "phase", "capacity", "usage", "accessMode", "volumeName", "node", "orgName", "clusterName"})
	} else {
		return err
	}
	o.NameSpace = *all.NameSpace
	o.Name = *all.Name
	o.Phase = *all.Phase
	o.Capacity = all.Capacity
	o.Usage = all.Usage
	o.AccessMode = all.AccessMode
	o.VolumeName = all.VolumeName
	o.Node = all.Node
	o.OrgName = all.OrgName
	o.ClusterName = all.ClusterName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
