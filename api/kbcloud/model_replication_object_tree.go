// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ReplicationObjectTree struct {
	NodeType          *string                     `json:"nodeType,omitempty"`
	NodeValues        []ReplicationObjectNode     `json:"nodeValues,omitempty"`
	ChildrenNodeTypes common.NullableList[string] `json:"childrenNodeTypes,omitempty"`
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

// GetNodeType returns the NodeType field value if set, zero value otherwise.
func (o *ReplicationObjectTree) GetNodeType() string {
	if o == nil || o.NodeType == nil {
		var ret string
		return ret
	}
	return *o.NodeType
}

// GetNodeTypeOk returns a tuple with the NodeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationObjectTree) GetNodeTypeOk() (*string, bool) {
	if o == nil || o.NodeType == nil {
		return nil, false
	}
	return o.NodeType, true
}

// HasNodeType returns a boolean if a field has been set.
func (o *ReplicationObjectTree) HasNodeType() bool {
	return o != nil && o.NodeType != nil
}

// SetNodeType gets a reference to the given string and assigns it to the NodeType field.
func (o *ReplicationObjectTree) SetNodeType(v string) {
	o.NodeType = &v
}

// GetNodeValues returns the NodeValues field value if set, zero value otherwise.
func (o *ReplicationObjectTree) GetNodeValues() []ReplicationObjectNode {
	if o == nil || o.NodeValues == nil {
		var ret []ReplicationObjectNode
		return ret
	}
	return o.NodeValues
}

// GetNodeValuesOk returns a tuple with the NodeValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationObjectTree) GetNodeValuesOk() (*[]ReplicationObjectNode, bool) {
	if o == nil || o.NodeValues == nil {
		return nil, false
	}
	return &o.NodeValues, true
}

// HasNodeValues returns a boolean if a field has been set.
func (o *ReplicationObjectTree) HasNodeValues() bool {
	return o != nil && o.NodeValues != nil
}

// SetNodeValues gets a reference to the given []ReplicationObjectNode and assigns it to the NodeValues field.
func (o *ReplicationObjectTree) SetNodeValues(v []ReplicationObjectNode) {
	o.NodeValues = v
}

// GetChildrenNodeTypes returns the ChildrenNodeTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReplicationObjectTree) GetChildrenNodeTypes() []string {
	if o == nil || o.ChildrenNodeTypes.Get() == nil {
		var ret []string
		return ret
	}
	return *o.ChildrenNodeTypes.Get()
}

// GetChildrenNodeTypesOk returns a tuple with the ChildrenNodeTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReplicationObjectTree) GetChildrenNodeTypesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChildrenNodeTypes.Get(), o.ChildrenNodeTypes.IsSet()
}

// HasChildrenNodeTypes returns a boolean if a field has been set.
func (o *ReplicationObjectTree) HasChildrenNodeTypes() bool {
	return o != nil && o.ChildrenNodeTypes.IsSet()
}

// SetChildrenNodeTypes gets a reference to the given common.NullableList[string] and assigns it to the ChildrenNodeTypes field.
func (o *ReplicationObjectTree) SetChildrenNodeTypes(v []string) {
	o.ChildrenNodeTypes.Set(&v)
}

// SetChildrenNodeTypesNil sets the value for ChildrenNodeTypes to be an explicit nil.
func (o *ReplicationObjectTree) SetChildrenNodeTypesNil() {
	o.ChildrenNodeTypes.Set(nil)
}

// UnsetChildrenNodeTypes ensures that no value is present for ChildrenNodeTypes, not even an explicit nil.
func (o *ReplicationObjectTree) UnsetChildrenNodeTypes() {
	o.ChildrenNodeTypes.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ReplicationObjectTree) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.NodeType != nil {
		toSerialize["nodeType"] = o.NodeType
	}
	if o.NodeValues != nil {
		toSerialize["nodeValues"] = o.NodeValues
	}
	if o.ChildrenNodeTypes.IsSet() {
		toSerialize["childrenNodeTypes"] = o.ChildrenNodeTypes.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ReplicationObjectTree) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NodeType          *string                     `json:"nodeType,omitempty"`
		NodeValues        []ReplicationObjectNode     `json:"nodeValues,omitempty"`
		ChildrenNodeTypes common.NullableList[string] `json:"childrenNodeTypes,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"nodeType", "nodeValues", "childrenNodeTypes"})
	} else {
		return err
	}
	o.NodeType = all.NodeType
	o.NodeValues = all.NodeValues
	o.ChildrenNodeTypes = all.ChildrenNodeTypes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
