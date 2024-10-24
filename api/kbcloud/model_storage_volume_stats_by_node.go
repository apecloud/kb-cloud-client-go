// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// StorageVolumeStatsByNode storageVolumeStatsByNode provides detailed information about storage volumes of nodes.
// NODESCRIPTION StorageVolumeStatsByNode
//
// Deprecated: This model is deprecated.
type StorageVolumeStatsByNode struct {
	// the name of the node
	NodeName *string `json:"nodeName,omitempty"`
	// the status of the node
	NodeStatus *string `json:"nodeStatus,omitempty"`
	// the number of PVCs on the node
	VolumeNum *string `json:"VolumeNum,omitempty"`
	// the sum of capacity of PVC on the node
	Capacity *string `json:"capacity,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStorageVolumeStatsByNode instantiates a new StorageVolumeStatsByNode object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStorageVolumeStatsByNode() *StorageVolumeStatsByNode {
	this := StorageVolumeStatsByNode{}
	return &this
}

// NewStorageVolumeStatsByNodeWithDefaults instantiates a new StorageVolumeStatsByNode object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStorageVolumeStatsByNodeWithDefaults() *StorageVolumeStatsByNode {
	this := StorageVolumeStatsByNode{}
	return &this
}

// GetNodeName returns the NodeName field value if set, zero value otherwise.
func (o *StorageVolumeStatsByNode) GetNodeName() string {
	if o == nil || o.NodeName == nil {
		var ret string
		return ret
	}
	return *o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVolumeStatsByNode) GetNodeNameOk() (*string, bool) {
	if o == nil || o.NodeName == nil {
		return nil, false
	}
	return o.NodeName, true
}

// HasNodeName returns a boolean if a field has been set.
func (o *StorageVolumeStatsByNode) HasNodeName() bool {
	return o != nil && o.NodeName != nil
}

// SetNodeName gets a reference to the given string and assigns it to the NodeName field.
func (o *StorageVolumeStatsByNode) SetNodeName(v string) {
	o.NodeName = &v
}

// GetNodeStatus returns the NodeStatus field value if set, zero value otherwise.
func (o *StorageVolumeStatsByNode) GetNodeStatus() string {
	if o == nil || o.NodeStatus == nil {
		var ret string
		return ret
	}
	return *o.NodeStatus
}

// GetNodeStatusOk returns a tuple with the NodeStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVolumeStatsByNode) GetNodeStatusOk() (*string, bool) {
	if o == nil || o.NodeStatus == nil {
		return nil, false
	}
	return o.NodeStatus, true
}

// HasNodeStatus returns a boolean if a field has been set.
func (o *StorageVolumeStatsByNode) HasNodeStatus() bool {
	return o != nil && o.NodeStatus != nil
}

// SetNodeStatus gets a reference to the given string and assigns it to the NodeStatus field.
func (o *StorageVolumeStatsByNode) SetNodeStatus(v string) {
	o.NodeStatus = &v
}

// GetVolumeNum returns the VolumeNum field value if set, zero value otherwise.
func (o *StorageVolumeStatsByNode) GetVolumeNum() string {
	if o == nil || o.VolumeNum == nil {
		var ret string
		return ret
	}
	return *o.VolumeNum
}

// GetVolumeNumOk returns a tuple with the VolumeNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVolumeStatsByNode) GetVolumeNumOk() (*string, bool) {
	if o == nil || o.VolumeNum == nil {
		return nil, false
	}
	return o.VolumeNum, true
}

// HasVolumeNum returns a boolean if a field has been set.
func (o *StorageVolumeStatsByNode) HasVolumeNum() bool {
	return o != nil && o.VolumeNum != nil
}

// SetVolumeNum gets a reference to the given string and assigns it to the VolumeNum field.
func (o *StorageVolumeStatsByNode) SetVolumeNum(v string) {
	o.VolumeNum = &v
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *StorageVolumeStatsByNode) GetCapacity() string {
	if o == nil || o.Capacity == nil {
		var ret string
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVolumeStatsByNode) GetCapacityOk() (*string, bool) {
	if o == nil || o.Capacity == nil {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *StorageVolumeStatsByNode) HasCapacity() bool {
	return o != nil && o.Capacity != nil
}

// SetCapacity gets a reference to the given string and assigns it to the Capacity field.
func (o *StorageVolumeStatsByNode) SetCapacity(v string) {
	o.Capacity = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o StorageVolumeStatsByNode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.NodeName != nil {
		toSerialize["nodeName"] = o.NodeName
	}
	if o.NodeStatus != nil {
		toSerialize["nodeStatus"] = o.NodeStatus
	}
	if o.VolumeNum != nil {
		toSerialize["VolumeNum"] = o.VolumeNum
	}
	if o.Capacity != nil {
		toSerialize["capacity"] = o.Capacity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StorageVolumeStatsByNode) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NodeName   *string `json:"nodeName,omitempty"`
		NodeStatus *string `json:"nodeStatus,omitempty"`
		VolumeNum  *string `json:"VolumeNum,omitempty"`
		Capacity   *string `json:"capacity,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"nodeName", "nodeStatus", "VolumeNum", "capacity"})
	} else {
		return err
	}
	o.NodeName = all.NodeName
	o.NodeStatus = all.NodeStatus
	o.VolumeNum = all.VolumeNum
	o.Capacity = all.Capacity

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
