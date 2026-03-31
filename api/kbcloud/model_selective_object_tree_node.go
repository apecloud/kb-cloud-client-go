// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// SelectiveObjectTreeNode tree node for selective backup objects
type SelectiveObjectTreeNode struct {
	// the type of the node
	Type string `json:"type"`
	// the name of the node
	Name string `json:"name"`
	// whether all children nodes are selected
	SelectedAll *bool `json:"selectedAll,omitempty"`
	// children nodes
	Children []SelectiveObjectTreeNode `json:"children,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSelectiveObjectTreeNode instantiates a new SelectiveObjectTreeNode object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSelectiveObjectTreeNode(typeVar string, name string) *SelectiveObjectTreeNode {
	this := SelectiveObjectTreeNode{}
	this.Type = typeVar
	this.Name = name
	return &this
}

// NewSelectiveObjectTreeNodeWithDefaults instantiates a new SelectiveObjectTreeNode object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSelectiveObjectTreeNodeWithDefaults() *SelectiveObjectTreeNode {
	this := SelectiveObjectTreeNode{}
	return &this
}

// GetType returns the Type field value.
func (o *SelectiveObjectTreeNode) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SelectiveObjectTreeNode) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SelectiveObjectTreeNode) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value.
func (o *SelectiveObjectTreeNode) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SelectiveObjectTreeNode) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *SelectiveObjectTreeNode) SetName(v string) {
	o.Name = v
}

// GetSelectedAll returns the SelectedAll field value if set, zero value otherwise.
func (o *SelectiveObjectTreeNode) GetSelectedAll() bool {
	if o == nil || o.SelectedAll == nil {
		var ret bool
		return ret
	}
	return *o.SelectedAll
}

// GetSelectedAllOk returns a tuple with the SelectedAll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectiveObjectTreeNode) GetSelectedAllOk() (*bool, bool) {
	if o == nil || o.SelectedAll == nil {
		return nil, false
	}
	return o.SelectedAll, true
}

// HasSelectedAll returns a boolean if a field has been set.
func (o *SelectiveObjectTreeNode) HasSelectedAll() bool {
	return o != nil && o.SelectedAll != nil
}

// SetSelectedAll gets a reference to the given bool and assigns it to the SelectedAll field.
func (o *SelectiveObjectTreeNode) SetSelectedAll(v bool) {
	o.SelectedAll = &v
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *SelectiveObjectTreeNode) GetChildren() []SelectiveObjectTreeNode {
	if o == nil || o.Children == nil {
		var ret []SelectiveObjectTreeNode
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectiveObjectTreeNode) GetChildrenOk() (*[]SelectiveObjectTreeNode, bool) {
	if o == nil || o.Children == nil {
		return nil, false
	}
	return &o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *SelectiveObjectTreeNode) HasChildren() bool {
	return o != nil && o.Children != nil
}

// SetChildren gets a reference to the given []SelectiveObjectTreeNode and assigns it to the Children field.
func (o *SelectiveObjectTreeNode) SetChildren(v []SelectiveObjectTreeNode) {
	o.Children = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SelectiveObjectTreeNode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["type"] = o.Type
	toSerialize["name"] = o.Name
	if o.SelectedAll != nil {
		toSerialize["selectedAll"] = o.SelectedAll
	}
	if o.Children != nil {
		toSerialize["children"] = o.Children
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SelectiveObjectTreeNode) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type        *string                   `json:"type"`
		Name        *string                   `json:"name"`
		SelectedAll *bool                     `json:"selectedAll,omitempty"`
		Children    []SelectiveObjectTreeNode `json:"children,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"type", "name", "selectedAll", "children"})
	} else {
		return err
	}
	o.Type = *all.Type
	o.Name = *all.Name
	o.SelectedAll = all.SelectedAll
	o.Children = all.Children

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
