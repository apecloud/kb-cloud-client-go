// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type ReplicationObjectTree struct {
	Nodes []ReplicationObjectTreeNode `json:"nodes,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewReplicationObjectTree instantiates a new ReplicationObjectTree object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReplicationObjectTree() *ReplicationObjectTree {
	this := ReplicationObjectTree{}
	return &this
}

// NewReplicationObjectTreeWithDefaults instantiates a new ReplicationObjectTree object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReplicationObjectTreeWithDefaults() *ReplicationObjectTree {
	this := ReplicationObjectTree{}
	return &this
}

// GetNodes returns the Nodes field value if set, zero value otherwise.
func (o *ReplicationObjectTree) GetNodes() []ReplicationObjectTreeNode {
	if o == nil || o.Nodes == nil {
		var ret []ReplicationObjectTreeNode
		return ret
	}
	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationObjectTree) GetNodesOk() (*[]ReplicationObjectTreeNode, bool) {
	if o == nil || o.Nodes == nil {
		return nil, false
	}
	return &o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *ReplicationObjectTree) HasNodes() bool {
	return o != nil && o.Nodes != nil
}

// SetNodes gets a reference to the given []ReplicationObjectTreeNode and assigns it to the Nodes field.
func (o *ReplicationObjectTree) SetNodes(v []ReplicationObjectTreeNode) {
	o.Nodes = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ReplicationObjectTree) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Nodes != nil {
		toSerialize["nodes"] = o.Nodes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ReplicationObjectTree) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Nodes []ReplicationObjectTreeNode `json:"nodes,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"nodes"})
	} else {
		return err
	}
	o.Nodes = all.Nodes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
