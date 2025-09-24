// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ReplicationObjectNode struct {
	NodeName  *string `json:"nodeName,omitempty"`
	NodeValue *string `json:"nodeValue,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewReplicationObjectNode instantiates a new ReplicationObjectNode object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReplicationObjectNode() *ReplicationObjectNode {
	this := ReplicationObjectNode{}
	return &this
}

// NewReplicationObjectNodeWithDefaults instantiates a new ReplicationObjectNode object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReplicationObjectNodeWithDefaults() *ReplicationObjectNode {
	this := ReplicationObjectNode{}
	return &this
}

// GetNodeName returns the NodeName field value if set, zero value otherwise.
func (o *ReplicationObjectNode) GetNodeName() string {
	if o == nil || o.NodeName == nil {
		var ret string
		return ret
	}
	return *o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationObjectNode) GetNodeNameOk() (*string, bool) {
	if o == nil || o.NodeName == nil {
		return nil, false
	}
	return o.NodeName, true
}

// HasNodeName returns a boolean if a field has been set.
func (o *ReplicationObjectNode) HasNodeName() bool {
	return o != nil && o.NodeName != nil
}

// SetNodeName gets a reference to the given string and assigns it to the NodeName field.
func (o *ReplicationObjectNode) SetNodeName(v string) {
	o.NodeName = &v
}

// GetNodeValue returns the NodeValue field value if set, zero value otherwise.
func (o *ReplicationObjectNode) GetNodeValue() string {
	if o == nil || o.NodeValue == nil {
		var ret string
		return ret
	}
	return *o.NodeValue
}

// GetNodeValueOk returns a tuple with the NodeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationObjectNode) GetNodeValueOk() (*string, bool) {
	if o == nil || o.NodeValue == nil {
		return nil, false
	}
	return o.NodeValue, true
}

// HasNodeValue returns a boolean if a field has been set.
func (o *ReplicationObjectNode) HasNodeValue() bool {
	return o != nil && o.NodeValue != nil
}

// SetNodeValue gets a reference to the given string and assigns it to the NodeValue field.
func (o *ReplicationObjectNode) SetNodeValue(v string) {
	o.NodeValue = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ReplicationObjectNode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.NodeName != nil {
		toSerialize["nodeName"] = o.NodeName
	}
	if o.NodeValue != nil {
		toSerialize["nodeValue"] = o.NodeValue
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ReplicationObjectNode) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NodeName  *string `json:"nodeName,omitempty"`
		NodeValue *string `json:"nodeValue,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"nodeName", "nodeValue"})
	} else {
		return err
	}
	o.NodeName = all.NodeName
	o.NodeValue = all.NodeValue

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
