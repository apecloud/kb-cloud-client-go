// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ElasticsearchNodeAllocationDecision struct {
	NodeId       common.NullableString            `json:"nodeId,omitempty"`
	NodeName     common.NullableString            `json:"nodeName,omitempty"`
	NodeDecision common.NullableString            `json:"nodeDecision,omitempty"`
	Deciders     []ElasticsearchAllocationDecider `json:"deciders,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewElasticsearchNodeAllocationDecision instantiates a new ElasticsearchNodeAllocationDecision object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewElasticsearchNodeAllocationDecision() *ElasticsearchNodeAllocationDecision {
	this := ElasticsearchNodeAllocationDecision{}
	return &this
}

// NewElasticsearchNodeAllocationDecisionWithDefaults instantiates a new ElasticsearchNodeAllocationDecision object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewElasticsearchNodeAllocationDecisionWithDefaults() *ElasticsearchNodeAllocationDecision {
	this := ElasticsearchNodeAllocationDecision{}
	return &this
}

// GetNodeId returns the NodeId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElasticsearchNodeAllocationDecision) GetNodeId() string {
	if o == nil || o.NodeId.Get() == nil {
		var ret string
		return ret
	}
	return *o.NodeId.Get()
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchNodeAllocationDecision) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NodeId.Get(), o.NodeId.IsSet()
}

// HasNodeId returns a boolean if a field has been set.
func (o *ElasticsearchNodeAllocationDecision) HasNodeId() bool {
	return o != nil && o.NodeId.IsSet()
}

// SetNodeId gets a reference to the given common.NullableString and assigns it to the NodeId field.
func (o *ElasticsearchNodeAllocationDecision) SetNodeId(v string) {
	o.NodeId.Set(&v)
}

// SetNodeIdNil sets the value for NodeId to be an explicit nil.
func (o *ElasticsearchNodeAllocationDecision) SetNodeIdNil() {
	o.NodeId.Set(nil)
}

// UnsetNodeId ensures that no value is present for NodeId, not even an explicit nil.
func (o *ElasticsearchNodeAllocationDecision) UnsetNodeId() {
	o.NodeId.Unset()
}

// GetNodeName returns the NodeName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElasticsearchNodeAllocationDecision) GetNodeName() string {
	if o == nil || o.NodeName.Get() == nil {
		var ret string
		return ret
	}
	return *o.NodeName.Get()
}

// GetNodeNameOk returns a tuple with the NodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchNodeAllocationDecision) GetNodeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NodeName.Get(), o.NodeName.IsSet()
}

// HasNodeName returns a boolean if a field has been set.
func (o *ElasticsearchNodeAllocationDecision) HasNodeName() bool {
	return o != nil && o.NodeName.IsSet()
}

// SetNodeName gets a reference to the given common.NullableString and assigns it to the NodeName field.
func (o *ElasticsearchNodeAllocationDecision) SetNodeName(v string) {
	o.NodeName.Set(&v)
}

// SetNodeNameNil sets the value for NodeName to be an explicit nil.
func (o *ElasticsearchNodeAllocationDecision) SetNodeNameNil() {
	o.NodeName.Set(nil)
}

// UnsetNodeName ensures that no value is present for NodeName, not even an explicit nil.
func (o *ElasticsearchNodeAllocationDecision) UnsetNodeName() {
	o.NodeName.Unset()
}

// GetNodeDecision returns the NodeDecision field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElasticsearchNodeAllocationDecision) GetNodeDecision() string {
	if o == nil || o.NodeDecision.Get() == nil {
		var ret string
		return ret
	}
	return *o.NodeDecision.Get()
}

// GetNodeDecisionOk returns a tuple with the NodeDecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchNodeAllocationDecision) GetNodeDecisionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NodeDecision.Get(), o.NodeDecision.IsSet()
}

// HasNodeDecision returns a boolean if a field has been set.
func (o *ElasticsearchNodeAllocationDecision) HasNodeDecision() bool {
	return o != nil && o.NodeDecision.IsSet()
}

// SetNodeDecision gets a reference to the given common.NullableString and assigns it to the NodeDecision field.
func (o *ElasticsearchNodeAllocationDecision) SetNodeDecision(v string) {
	o.NodeDecision.Set(&v)
}

// SetNodeDecisionNil sets the value for NodeDecision to be an explicit nil.
func (o *ElasticsearchNodeAllocationDecision) SetNodeDecisionNil() {
	o.NodeDecision.Set(nil)
}

// UnsetNodeDecision ensures that no value is present for NodeDecision, not even an explicit nil.
func (o *ElasticsearchNodeAllocationDecision) UnsetNodeDecision() {
	o.NodeDecision.Unset()
}

// GetDeciders returns the Deciders field value if set, zero value otherwise.
func (o *ElasticsearchNodeAllocationDecision) GetDeciders() []ElasticsearchAllocationDecider {
	if o == nil || o.Deciders == nil {
		var ret []ElasticsearchAllocationDecider
		return ret
	}
	return o.Deciders
}

// GetDecidersOk returns a tuple with the Deciders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticsearchNodeAllocationDecision) GetDecidersOk() (*[]ElasticsearchAllocationDecider, bool) {
	if o == nil || o.Deciders == nil {
		return nil, false
	}
	return &o.Deciders, true
}

// HasDeciders returns a boolean if a field has been set.
func (o *ElasticsearchNodeAllocationDecision) HasDeciders() bool {
	return o != nil && o.Deciders != nil
}

// SetDeciders gets a reference to the given []ElasticsearchAllocationDecider and assigns it to the Deciders field.
func (o *ElasticsearchNodeAllocationDecision) SetDeciders(v []ElasticsearchAllocationDecider) {
	o.Deciders = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ElasticsearchNodeAllocationDecision) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.NodeId.IsSet() {
		toSerialize["nodeId"] = o.NodeId.Get()
	}
	if o.NodeName.IsSet() {
		toSerialize["nodeName"] = o.NodeName.Get()
	}
	if o.NodeDecision.IsSet() {
		toSerialize["nodeDecision"] = o.NodeDecision.Get()
	}
	if o.Deciders != nil {
		toSerialize["deciders"] = o.Deciders
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ElasticsearchNodeAllocationDecision) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NodeId       common.NullableString            `json:"nodeId,omitempty"`
		NodeName     common.NullableString            `json:"nodeName,omitempty"`
		NodeDecision common.NullableString            `json:"nodeDecision,omitempty"`
		Deciders     []ElasticsearchAllocationDecider `json:"deciders,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"nodeId", "nodeName", "nodeDecision", "deciders"})
	} else {
		return err
	}
	o.NodeId = all.NodeId
	o.NodeName = all.NodeName
	o.NodeDecision = all.NodeDecision
	o.Deciders = all.Deciders

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
