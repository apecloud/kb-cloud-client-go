// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type ReplicationObjectTreeNode struct {
	NodeName  *string                     `json:"nodeName,omitempty"`
	NodeType  *string                     `json:"nodeType,omitempty"`
	NodeValue *string                     `json:"nodeValue,omitempty"`
	Children  []ReplicationObjectTreeNode `json:"children,omitempty"`
	IsLeaf    common.NullableBool         `json:"isLeaf,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewReplicationObjectTreeNode instantiates a new ReplicationObjectTreeNode object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReplicationObjectTreeNode() *ReplicationObjectTreeNode {
	this := ReplicationObjectTreeNode{}
	return &this
}

// NewReplicationObjectTreeNodeWithDefaults instantiates a new ReplicationObjectTreeNode object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReplicationObjectTreeNodeWithDefaults() *ReplicationObjectTreeNode {
	this := ReplicationObjectTreeNode{}
	return &this
}

// GetNodeName returns the NodeName field value if set, zero value otherwise.
func (o *ReplicationObjectTreeNode) GetNodeName() string {
	if o == nil || o.NodeName == nil {
		var ret string
		return ret
	}
	return *o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationObjectTreeNode) GetNodeNameOk() (*string, bool) {
	if o == nil || o.NodeName == nil {
		return nil, false
	}
	return o.NodeName, true
}

// HasNodeName returns a boolean if a field has been set.
func (o *ReplicationObjectTreeNode) HasNodeName() bool {
	return o != nil && o.NodeName != nil
}

// SetNodeName gets a reference to the given string and assigns it to the NodeName field.
func (o *ReplicationObjectTreeNode) SetNodeName(v string) {
	o.NodeName = &v
}

// GetNodeType returns the NodeType field value if set, zero value otherwise.
func (o *ReplicationObjectTreeNode) GetNodeType() string {
	if o == nil || o.NodeType == nil {
		var ret string
		return ret
	}
	return *o.NodeType
}

// GetNodeTypeOk returns a tuple with the NodeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationObjectTreeNode) GetNodeTypeOk() (*string, bool) {
	if o == nil || o.NodeType == nil {
		return nil, false
	}
	return o.NodeType, true
}

// HasNodeType returns a boolean if a field has been set.
func (o *ReplicationObjectTreeNode) HasNodeType() bool {
	return o != nil && o.NodeType != nil
}

// SetNodeType gets a reference to the given string and assigns it to the NodeType field.
func (o *ReplicationObjectTreeNode) SetNodeType(v string) {
	o.NodeType = &v
}

// GetNodeValue returns the NodeValue field value if set, zero value otherwise.
func (o *ReplicationObjectTreeNode) GetNodeValue() string {
	if o == nil || o.NodeValue == nil {
		var ret string
		return ret
	}
	return *o.NodeValue
}

// GetNodeValueOk returns a tuple with the NodeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationObjectTreeNode) GetNodeValueOk() (*string, bool) {
	if o == nil || o.NodeValue == nil {
		return nil, false
	}
	return o.NodeValue, true
}

// HasNodeValue returns a boolean if a field has been set.
func (o *ReplicationObjectTreeNode) HasNodeValue() bool {
	return o != nil && o.NodeValue != nil
}

// SetNodeValue gets a reference to the given string and assigns it to the NodeValue field.
func (o *ReplicationObjectTreeNode) SetNodeValue(v string) {
	o.NodeValue = &v
}

// GetChildren returns the Children field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReplicationObjectTreeNode) GetChildren() []ReplicationObjectTreeNode {
	if o == nil {
		var ret []ReplicationObjectTreeNode
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReplicationObjectTreeNode) GetChildrenOk() (*[]ReplicationObjectTreeNode, bool) {
	if o == nil || o.Children == nil {
		return nil, false
	}
	return &o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *ReplicationObjectTreeNode) HasChildren() bool {
	return o != nil && o.Children != nil
}

// SetChildren gets a reference to the given []ReplicationObjectTreeNode and assigns it to the Children field.
func (o *ReplicationObjectTreeNode) SetChildren(v []ReplicationObjectTreeNode) {
	o.Children = v
}

// GetIsLeaf returns the IsLeaf field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReplicationObjectTreeNode) GetIsLeaf() bool {
	if o == nil || o.IsLeaf.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsLeaf.Get()
}

// GetIsLeafOk returns a tuple with the IsLeaf field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReplicationObjectTreeNode) GetIsLeafOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsLeaf.Get(), o.IsLeaf.IsSet()
}

// HasIsLeaf returns a boolean if a field has been set.
func (o *ReplicationObjectTreeNode) HasIsLeaf() bool {
	return o != nil && o.IsLeaf.IsSet()
}

// SetIsLeaf gets a reference to the given common.NullableBool and assigns it to the IsLeaf field.
func (o *ReplicationObjectTreeNode) SetIsLeaf(v bool) {
	o.IsLeaf.Set(&v)
}

// SetIsLeafNil sets the value for IsLeaf to be an explicit nil.
func (o *ReplicationObjectTreeNode) SetIsLeafNil() {
	o.IsLeaf.Set(nil)
}

// UnsetIsLeaf ensures that no value is present for IsLeaf, not even an explicit nil.
func (o *ReplicationObjectTreeNode) UnsetIsLeaf() {
	o.IsLeaf.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ReplicationObjectTreeNode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.NodeName != nil {
		toSerialize["nodeName"] = o.NodeName
	}
	if o.NodeType != nil {
		toSerialize["nodeType"] = o.NodeType
	}
	if o.NodeValue != nil {
		toSerialize["nodeValue"] = o.NodeValue
	}
	if o.Children != nil {
		toSerialize["children"] = o.Children
	}
	if o.IsLeaf.IsSet() {
		toSerialize["isLeaf"] = o.IsLeaf.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ReplicationObjectTreeNode) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NodeName  *string                     `json:"nodeName,omitempty"`
		NodeType  *string                     `json:"nodeType,omitempty"`
		NodeValue *string                     `json:"nodeValue,omitempty"`
		Children  []ReplicationObjectTreeNode `json:"children,omitempty"`
		IsLeaf    common.NullableBool         `json:"isLeaf,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"nodeName", "nodeType", "nodeValue", "children", "isLeaf"})
	} else {
		return err
	}
	o.NodeName = all.NodeName
	o.NodeType = all.NodeType
	o.NodeValue = all.NodeValue
	o.Children = all.Children
	o.IsLeaf = all.IsLeaf

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
