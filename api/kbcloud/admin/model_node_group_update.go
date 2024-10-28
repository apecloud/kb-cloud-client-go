// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// NodeGroupUpdate NodeGroup patch info

type NodeGroupUpdate struct {
	// NodeGroup description
	Description common.NullableString `json:"description,omitempty"`
	// NODESCRIPTION Nodes
	Nodes []NodeOperation `json:"nodes,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNodeGroupUpdate instantiates a new NodeGroupUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNodeGroupUpdate() *NodeGroupUpdate {
	this := NodeGroupUpdate{}
	return &this
}

// NewNodeGroupUpdateWithDefaults instantiates a new NodeGroupUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNodeGroupUpdateWithDefaults() *NodeGroupUpdate {
	this := NodeGroupUpdate{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NodeGroupUpdate) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *NodeGroupUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *NodeGroupUpdate) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given common.NullableString and assigns it to the Description field.
func (o *NodeGroupUpdate) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *NodeGroupUpdate) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *NodeGroupUpdate) UnsetDescription() {
	o.Description.Unset()
}

// GetNodes returns the Nodes field value if set, zero value otherwise.
func (o *NodeGroupUpdate) GetNodes() []NodeOperation {
	if o == nil || o.Nodes == nil {
		var ret []NodeOperation
		return ret
	}
	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeGroupUpdate) GetNodesOk() (*[]NodeOperation, bool) {
	if o == nil || o.Nodes == nil {
		return nil, false
	}
	return &o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *NodeGroupUpdate) HasNodes() bool {
	return o != nil && o.Nodes != nil
}

// SetNodes gets a reference to the given []NodeOperation and assigns it to the Nodes field.
func (o *NodeGroupUpdate) SetNodes(v []NodeOperation) {
	o.Nodes = v
}

// MarshalJSON serializes the struct using spec logic.
func (o NodeGroupUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
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
func (o *NodeGroupUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description common.NullableString `json:"description,omitempty"`
		Nodes       []NodeOperation       `json:"nodes,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"description", "nodes"})
	} else {
		return err
	}
	o.Description = all.Description
	o.Nodes = all.Nodes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
