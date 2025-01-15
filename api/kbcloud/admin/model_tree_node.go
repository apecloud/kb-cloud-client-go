// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type TreeNode struct {
	// A string representing the level of the node.
	Level *string `json:"level,omitempty"`
	// A string representing the name of the node.
	Name *string `json:"name,omitempty"`
	// A string representing the uid of the resource.
	Uid *string `json:"uid,omitempty"`
	// A mapping where each key is a string representing the event name, and the value is a MapNode object.
	Events []Event `json:"events,omitempty"`
	// An array of child nodes.
	Children []TreeNode `json:"children,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTreeNode instantiates a new TreeNode object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTreeNode() *TreeNode {
	this := TreeNode{}
	return &this
}

// NewTreeNodeWithDefaults instantiates a new TreeNode object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTreeNodeWithDefaults() *TreeNode {
	this := TreeNode{}
	return &this
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *TreeNode) GetLevel() string {
	if o == nil || o.Level == nil {
		var ret string
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TreeNode) GetLevelOk() (*string, bool) {
	if o == nil || o.Level == nil {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *TreeNode) HasLevel() bool {
	return o != nil && o.Level != nil
}

// SetLevel gets a reference to the given string and assigns it to the Level field.
func (o *TreeNode) SetLevel(v string) {
	o.Level = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TreeNode) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TreeNode) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TreeNode) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TreeNode) SetName(v string) {
	o.Name = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *TreeNode) GetUid() string {
	if o == nil || o.Uid == nil {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TreeNode) GetUidOk() (*string, bool) {
	if o == nil || o.Uid == nil {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *TreeNode) HasUid() bool {
	return o != nil && o.Uid != nil
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *TreeNode) SetUid(v string) {
	o.Uid = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *TreeNode) GetEvents() []Event {
	if o == nil || o.Events == nil {
		var ret []Event
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TreeNode) GetEventsOk() (*[]Event, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return &o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *TreeNode) HasEvents() bool {
	return o != nil && o.Events != nil
}

// SetEvents gets a reference to the given []Event and assigns it to the Events field.
func (o *TreeNode) SetEvents(v []Event) {
	o.Events = v
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *TreeNode) GetChildren() []TreeNode {
	if o == nil || o.Children == nil {
		var ret []TreeNode
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TreeNode) GetChildrenOk() (*[]TreeNode, bool) {
	if o == nil || o.Children == nil {
		return nil, false
	}
	return &o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *TreeNode) HasChildren() bool {
	return o != nil && o.Children != nil
}

// SetChildren gets a reference to the given []TreeNode and assigns it to the Children field.
func (o *TreeNode) SetChildren(v []TreeNode) {
	o.Children = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TreeNode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Level != nil {
		toSerialize["level"] = o.Level
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Uid != nil {
		toSerialize["uid"] = o.Uid
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
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
func (o *TreeNode) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Level    *string    `json:"level,omitempty"`
		Name     *string    `json:"name,omitempty"`
		Uid      *string    `json:"uid,omitempty"`
		Events   []Event    `json:"events,omitempty"`
		Children []TreeNode `json:"children,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"level", "name", "uid", "events", "children"})
	} else {
		return err
	}
	o.Level = all.Level
	o.Name = all.Name
	o.Uid = all.Uid
	o.Events = all.Events
	o.Children = all.Children

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
