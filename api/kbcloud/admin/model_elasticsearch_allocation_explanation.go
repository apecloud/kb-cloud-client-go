// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ElasticsearchAllocationExplanation struct {
	Index                        common.NullableString                 `json:"index,omitempty"`
	Shard                        common.NullableInt64                  `json:"shard,omitempty"`
	Primary                      common.NullableBool                   `json:"primary,omitempty"`
	CurrentState                 common.NullableString                 `json:"currentState,omitempty"`
	CurrentNode                  *ElasticsearchShardNode               `json:"currentNode,omitempty"`
	UnassignedInfo               *ElasticsearchUnassignedInfo          `json:"unassignedInfo,omitempty"`
	CanAllocate                  common.NullableString                 `json:"canAllocate,omitempty"`
	AllocateExplanation          common.NullableString                 `json:"allocateExplanation,omitempty"`
	CanRemainOnCurrentNode       common.NullableString                 `json:"canRemainOnCurrentNode,omitempty"`
	CanRemainDecisions           []ElasticsearchAllocationDecider      `json:"canRemainDecisions,omitempty"`
	CanRebalanceCluster          common.NullableString                 `json:"canRebalanceCluster,omitempty"`
	CanRebalanceClusterDecisions []ElasticsearchAllocationDecider      `json:"canRebalanceClusterDecisions,omitempty"`
	CanRebalanceToOtherNode      common.NullableString                 `json:"canRebalanceToOtherNode,omitempty"`
	RebalanceExplanation         common.NullableString                 `json:"rebalanceExplanation,omitempty"`
	NodeAllocationDecisions      []ElasticsearchNodeAllocationDecision `json:"nodeAllocationDecisions,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewElasticsearchAllocationExplanation instantiates a new ElasticsearchAllocationExplanation object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewElasticsearchAllocationExplanation() *ElasticsearchAllocationExplanation {
	this := ElasticsearchAllocationExplanation{}
	return &this
}

// NewElasticsearchAllocationExplanationWithDefaults instantiates a new ElasticsearchAllocationExplanation object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewElasticsearchAllocationExplanationWithDefaults() *ElasticsearchAllocationExplanation {
	this := ElasticsearchAllocationExplanation{}
	return &this
}

// GetIndex returns the Index field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElasticsearchAllocationExplanation) GetIndex() string {
	if o == nil || o.Index.Get() == nil {
		var ret string
		return ret
	}
	return *o.Index.Get()
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchAllocationExplanation) GetIndexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Index.Get(), o.Index.IsSet()
}

// HasIndex returns a boolean if a field has been set.
func (o *ElasticsearchAllocationExplanation) HasIndex() bool {
	return o != nil && o.Index.IsSet()
}

// SetIndex gets a reference to the given common.NullableString and assigns it to the Index field.
func (o *ElasticsearchAllocationExplanation) SetIndex(v string) {
	o.Index.Set(&v)
}

// SetIndexNil sets the value for Index to be an explicit nil.
func (o *ElasticsearchAllocationExplanation) SetIndexNil() {
	o.Index.Set(nil)
}

// UnsetIndex ensures that no value is present for Index, not even an explicit nil.
func (o *ElasticsearchAllocationExplanation) UnsetIndex() {
	o.Index.Unset()
}

// GetShard returns the Shard field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElasticsearchAllocationExplanation) GetShard() int64 {
	if o == nil || o.Shard.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Shard.Get()
}

// GetShardOk returns a tuple with the Shard field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchAllocationExplanation) GetShardOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Shard.Get(), o.Shard.IsSet()
}

// HasShard returns a boolean if a field has been set.
func (o *ElasticsearchAllocationExplanation) HasShard() bool {
	return o != nil && o.Shard.IsSet()
}

// SetShard gets a reference to the given common.NullableInt64 and assigns it to the Shard field.
func (o *ElasticsearchAllocationExplanation) SetShard(v int64) {
	o.Shard.Set(&v)
}

// SetShardNil sets the value for Shard to be an explicit nil.
func (o *ElasticsearchAllocationExplanation) SetShardNil() {
	o.Shard.Set(nil)
}

// UnsetShard ensures that no value is present for Shard, not even an explicit nil.
func (o *ElasticsearchAllocationExplanation) UnsetShard() {
	o.Shard.Unset()
}

// GetPrimary returns the Primary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElasticsearchAllocationExplanation) GetPrimary() bool {
	if o == nil || o.Primary.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Primary.Get()
}

// GetPrimaryOk returns a tuple with the Primary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchAllocationExplanation) GetPrimaryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Primary.Get(), o.Primary.IsSet()
}

// HasPrimary returns a boolean if a field has been set.
func (o *ElasticsearchAllocationExplanation) HasPrimary() bool {
	return o != nil && o.Primary.IsSet()
}

// SetPrimary gets a reference to the given common.NullableBool and assigns it to the Primary field.
func (o *ElasticsearchAllocationExplanation) SetPrimary(v bool) {
	o.Primary.Set(&v)
}

// SetPrimaryNil sets the value for Primary to be an explicit nil.
func (o *ElasticsearchAllocationExplanation) SetPrimaryNil() {
	o.Primary.Set(nil)
}

// UnsetPrimary ensures that no value is present for Primary, not even an explicit nil.
func (o *ElasticsearchAllocationExplanation) UnsetPrimary() {
	o.Primary.Unset()
}

// GetCurrentState returns the CurrentState field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElasticsearchAllocationExplanation) GetCurrentState() string {
	if o == nil || o.CurrentState.Get() == nil {
		var ret string
		return ret
	}
	return *o.CurrentState.Get()
}

// GetCurrentStateOk returns a tuple with the CurrentState field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchAllocationExplanation) GetCurrentStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrentState.Get(), o.CurrentState.IsSet()
}

// HasCurrentState returns a boolean if a field has been set.
func (o *ElasticsearchAllocationExplanation) HasCurrentState() bool {
	return o != nil && o.CurrentState.IsSet()
}

// SetCurrentState gets a reference to the given common.NullableString and assigns it to the CurrentState field.
func (o *ElasticsearchAllocationExplanation) SetCurrentState(v string) {
	o.CurrentState.Set(&v)
}

// SetCurrentStateNil sets the value for CurrentState to be an explicit nil.
func (o *ElasticsearchAllocationExplanation) SetCurrentStateNil() {
	o.CurrentState.Set(nil)
}

// UnsetCurrentState ensures that no value is present for CurrentState, not even an explicit nil.
func (o *ElasticsearchAllocationExplanation) UnsetCurrentState() {
	o.CurrentState.Unset()
}

// GetCurrentNode returns the CurrentNode field value if set, zero value otherwise.
func (o *ElasticsearchAllocationExplanation) GetCurrentNode() ElasticsearchShardNode {
	if o == nil || o.CurrentNode == nil {
		var ret ElasticsearchShardNode
		return ret
	}
	return *o.CurrentNode
}

// GetCurrentNodeOk returns a tuple with the CurrentNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticsearchAllocationExplanation) GetCurrentNodeOk() (*ElasticsearchShardNode, bool) {
	if o == nil || o.CurrentNode == nil {
		return nil, false
	}
	return o.CurrentNode, true
}

// HasCurrentNode returns a boolean if a field has been set.
func (o *ElasticsearchAllocationExplanation) HasCurrentNode() bool {
	return o != nil && o.CurrentNode != nil
}

// SetCurrentNode gets a reference to the given ElasticsearchShardNode and assigns it to the CurrentNode field.
func (o *ElasticsearchAllocationExplanation) SetCurrentNode(v ElasticsearchShardNode) {
	o.CurrentNode = &v
}

// GetUnassignedInfo returns the UnassignedInfo field value if set, zero value otherwise.
func (o *ElasticsearchAllocationExplanation) GetUnassignedInfo() ElasticsearchUnassignedInfo {
	if o == nil || o.UnassignedInfo == nil {
		var ret ElasticsearchUnassignedInfo
		return ret
	}
	return *o.UnassignedInfo
}

// GetUnassignedInfoOk returns a tuple with the UnassignedInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticsearchAllocationExplanation) GetUnassignedInfoOk() (*ElasticsearchUnassignedInfo, bool) {
	if o == nil || o.UnassignedInfo == nil {
		return nil, false
	}
	return o.UnassignedInfo, true
}

// HasUnassignedInfo returns a boolean if a field has been set.
func (o *ElasticsearchAllocationExplanation) HasUnassignedInfo() bool {
	return o != nil && o.UnassignedInfo != nil
}

// SetUnassignedInfo gets a reference to the given ElasticsearchUnassignedInfo and assigns it to the UnassignedInfo field.
func (o *ElasticsearchAllocationExplanation) SetUnassignedInfo(v ElasticsearchUnassignedInfo) {
	o.UnassignedInfo = &v
}

// GetCanAllocate returns the CanAllocate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElasticsearchAllocationExplanation) GetCanAllocate() string {
	if o == nil || o.CanAllocate.Get() == nil {
		var ret string
		return ret
	}
	return *o.CanAllocate.Get()
}

// GetCanAllocateOk returns a tuple with the CanAllocate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchAllocationExplanation) GetCanAllocateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CanAllocate.Get(), o.CanAllocate.IsSet()
}

// HasCanAllocate returns a boolean if a field has been set.
func (o *ElasticsearchAllocationExplanation) HasCanAllocate() bool {
	return o != nil && o.CanAllocate.IsSet()
}

// SetCanAllocate gets a reference to the given common.NullableString and assigns it to the CanAllocate field.
func (o *ElasticsearchAllocationExplanation) SetCanAllocate(v string) {
	o.CanAllocate.Set(&v)
}

// SetCanAllocateNil sets the value for CanAllocate to be an explicit nil.
func (o *ElasticsearchAllocationExplanation) SetCanAllocateNil() {
	o.CanAllocate.Set(nil)
}

// UnsetCanAllocate ensures that no value is present for CanAllocate, not even an explicit nil.
func (o *ElasticsearchAllocationExplanation) UnsetCanAllocate() {
	o.CanAllocate.Unset()
}

// GetAllocateExplanation returns the AllocateExplanation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElasticsearchAllocationExplanation) GetAllocateExplanation() string {
	if o == nil || o.AllocateExplanation.Get() == nil {
		var ret string
		return ret
	}
	return *o.AllocateExplanation.Get()
}

// GetAllocateExplanationOk returns a tuple with the AllocateExplanation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchAllocationExplanation) GetAllocateExplanationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllocateExplanation.Get(), o.AllocateExplanation.IsSet()
}

// HasAllocateExplanation returns a boolean if a field has been set.
func (o *ElasticsearchAllocationExplanation) HasAllocateExplanation() bool {
	return o != nil && o.AllocateExplanation.IsSet()
}

// SetAllocateExplanation gets a reference to the given common.NullableString and assigns it to the AllocateExplanation field.
func (o *ElasticsearchAllocationExplanation) SetAllocateExplanation(v string) {
	o.AllocateExplanation.Set(&v)
}

// SetAllocateExplanationNil sets the value for AllocateExplanation to be an explicit nil.
func (o *ElasticsearchAllocationExplanation) SetAllocateExplanationNil() {
	o.AllocateExplanation.Set(nil)
}

// UnsetAllocateExplanation ensures that no value is present for AllocateExplanation, not even an explicit nil.
func (o *ElasticsearchAllocationExplanation) UnsetAllocateExplanation() {
	o.AllocateExplanation.Unset()
}

// GetCanRemainOnCurrentNode returns the CanRemainOnCurrentNode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElasticsearchAllocationExplanation) GetCanRemainOnCurrentNode() string {
	if o == nil || o.CanRemainOnCurrentNode.Get() == nil {
		var ret string
		return ret
	}
	return *o.CanRemainOnCurrentNode.Get()
}

// GetCanRemainOnCurrentNodeOk returns a tuple with the CanRemainOnCurrentNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchAllocationExplanation) GetCanRemainOnCurrentNodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CanRemainOnCurrentNode.Get(), o.CanRemainOnCurrentNode.IsSet()
}

// HasCanRemainOnCurrentNode returns a boolean if a field has been set.
func (o *ElasticsearchAllocationExplanation) HasCanRemainOnCurrentNode() bool {
	return o != nil && o.CanRemainOnCurrentNode.IsSet()
}

// SetCanRemainOnCurrentNode gets a reference to the given common.NullableString and assigns it to the CanRemainOnCurrentNode field.
func (o *ElasticsearchAllocationExplanation) SetCanRemainOnCurrentNode(v string) {
	o.CanRemainOnCurrentNode.Set(&v)
}

// SetCanRemainOnCurrentNodeNil sets the value for CanRemainOnCurrentNode to be an explicit nil.
func (o *ElasticsearchAllocationExplanation) SetCanRemainOnCurrentNodeNil() {
	o.CanRemainOnCurrentNode.Set(nil)
}

// UnsetCanRemainOnCurrentNode ensures that no value is present for CanRemainOnCurrentNode, not even an explicit nil.
func (o *ElasticsearchAllocationExplanation) UnsetCanRemainOnCurrentNode() {
	o.CanRemainOnCurrentNode.Unset()
}

// GetCanRemainDecisions returns the CanRemainDecisions field value if set, zero value otherwise.
func (o *ElasticsearchAllocationExplanation) GetCanRemainDecisions() []ElasticsearchAllocationDecider {
	if o == nil || o.CanRemainDecisions == nil {
		var ret []ElasticsearchAllocationDecider
		return ret
	}
	return o.CanRemainDecisions
}

// GetCanRemainDecisionsOk returns a tuple with the CanRemainDecisions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticsearchAllocationExplanation) GetCanRemainDecisionsOk() (*[]ElasticsearchAllocationDecider, bool) {
	if o == nil || o.CanRemainDecisions == nil {
		return nil, false
	}
	return &o.CanRemainDecisions, true
}

// HasCanRemainDecisions returns a boolean if a field has been set.
func (o *ElasticsearchAllocationExplanation) HasCanRemainDecisions() bool {
	return o != nil && o.CanRemainDecisions != nil
}

// SetCanRemainDecisions gets a reference to the given []ElasticsearchAllocationDecider and assigns it to the CanRemainDecisions field.
func (o *ElasticsearchAllocationExplanation) SetCanRemainDecisions(v []ElasticsearchAllocationDecider) {
	o.CanRemainDecisions = v
}

// GetCanRebalanceCluster returns the CanRebalanceCluster field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElasticsearchAllocationExplanation) GetCanRebalanceCluster() string {
	if o == nil || o.CanRebalanceCluster.Get() == nil {
		var ret string
		return ret
	}
	return *o.CanRebalanceCluster.Get()
}

// GetCanRebalanceClusterOk returns a tuple with the CanRebalanceCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchAllocationExplanation) GetCanRebalanceClusterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CanRebalanceCluster.Get(), o.CanRebalanceCluster.IsSet()
}

// HasCanRebalanceCluster returns a boolean if a field has been set.
func (o *ElasticsearchAllocationExplanation) HasCanRebalanceCluster() bool {
	return o != nil && o.CanRebalanceCluster.IsSet()
}

// SetCanRebalanceCluster gets a reference to the given common.NullableString and assigns it to the CanRebalanceCluster field.
func (o *ElasticsearchAllocationExplanation) SetCanRebalanceCluster(v string) {
	o.CanRebalanceCluster.Set(&v)
}

// SetCanRebalanceClusterNil sets the value for CanRebalanceCluster to be an explicit nil.
func (o *ElasticsearchAllocationExplanation) SetCanRebalanceClusterNil() {
	o.CanRebalanceCluster.Set(nil)
}

// UnsetCanRebalanceCluster ensures that no value is present for CanRebalanceCluster, not even an explicit nil.
func (o *ElasticsearchAllocationExplanation) UnsetCanRebalanceCluster() {
	o.CanRebalanceCluster.Unset()
}

// GetCanRebalanceClusterDecisions returns the CanRebalanceClusterDecisions field value if set, zero value otherwise.
func (o *ElasticsearchAllocationExplanation) GetCanRebalanceClusterDecisions() []ElasticsearchAllocationDecider {
	if o == nil || o.CanRebalanceClusterDecisions == nil {
		var ret []ElasticsearchAllocationDecider
		return ret
	}
	return o.CanRebalanceClusterDecisions
}

// GetCanRebalanceClusterDecisionsOk returns a tuple with the CanRebalanceClusterDecisions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticsearchAllocationExplanation) GetCanRebalanceClusterDecisionsOk() (*[]ElasticsearchAllocationDecider, bool) {
	if o == nil || o.CanRebalanceClusterDecisions == nil {
		return nil, false
	}
	return &o.CanRebalanceClusterDecisions, true
}

// HasCanRebalanceClusterDecisions returns a boolean if a field has been set.
func (o *ElasticsearchAllocationExplanation) HasCanRebalanceClusterDecisions() bool {
	return o != nil && o.CanRebalanceClusterDecisions != nil
}

// SetCanRebalanceClusterDecisions gets a reference to the given []ElasticsearchAllocationDecider and assigns it to the CanRebalanceClusterDecisions field.
func (o *ElasticsearchAllocationExplanation) SetCanRebalanceClusterDecisions(v []ElasticsearchAllocationDecider) {
	o.CanRebalanceClusterDecisions = v
}

// GetCanRebalanceToOtherNode returns the CanRebalanceToOtherNode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElasticsearchAllocationExplanation) GetCanRebalanceToOtherNode() string {
	if o == nil || o.CanRebalanceToOtherNode.Get() == nil {
		var ret string
		return ret
	}
	return *o.CanRebalanceToOtherNode.Get()
}

// GetCanRebalanceToOtherNodeOk returns a tuple with the CanRebalanceToOtherNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchAllocationExplanation) GetCanRebalanceToOtherNodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CanRebalanceToOtherNode.Get(), o.CanRebalanceToOtherNode.IsSet()
}

// HasCanRebalanceToOtherNode returns a boolean if a field has been set.
func (o *ElasticsearchAllocationExplanation) HasCanRebalanceToOtherNode() bool {
	return o != nil && o.CanRebalanceToOtherNode.IsSet()
}

// SetCanRebalanceToOtherNode gets a reference to the given common.NullableString and assigns it to the CanRebalanceToOtherNode field.
func (o *ElasticsearchAllocationExplanation) SetCanRebalanceToOtherNode(v string) {
	o.CanRebalanceToOtherNode.Set(&v)
}

// SetCanRebalanceToOtherNodeNil sets the value for CanRebalanceToOtherNode to be an explicit nil.
func (o *ElasticsearchAllocationExplanation) SetCanRebalanceToOtherNodeNil() {
	o.CanRebalanceToOtherNode.Set(nil)
}

// UnsetCanRebalanceToOtherNode ensures that no value is present for CanRebalanceToOtherNode, not even an explicit nil.
func (o *ElasticsearchAllocationExplanation) UnsetCanRebalanceToOtherNode() {
	o.CanRebalanceToOtherNode.Unset()
}

// GetRebalanceExplanation returns the RebalanceExplanation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElasticsearchAllocationExplanation) GetRebalanceExplanation() string {
	if o == nil || o.RebalanceExplanation.Get() == nil {
		var ret string
		return ret
	}
	return *o.RebalanceExplanation.Get()
}

// GetRebalanceExplanationOk returns a tuple with the RebalanceExplanation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchAllocationExplanation) GetRebalanceExplanationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RebalanceExplanation.Get(), o.RebalanceExplanation.IsSet()
}

// HasRebalanceExplanation returns a boolean if a field has been set.
func (o *ElasticsearchAllocationExplanation) HasRebalanceExplanation() bool {
	return o != nil && o.RebalanceExplanation.IsSet()
}

// SetRebalanceExplanation gets a reference to the given common.NullableString and assigns it to the RebalanceExplanation field.
func (o *ElasticsearchAllocationExplanation) SetRebalanceExplanation(v string) {
	o.RebalanceExplanation.Set(&v)
}

// SetRebalanceExplanationNil sets the value for RebalanceExplanation to be an explicit nil.
func (o *ElasticsearchAllocationExplanation) SetRebalanceExplanationNil() {
	o.RebalanceExplanation.Set(nil)
}

// UnsetRebalanceExplanation ensures that no value is present for RebalanceExplanation, not even an explicit nil.
func (o *ElasticsearchAllocationExplanation) UnsetRebalanceExplanation() {
	o.RebalanceExplanation.Unset()
}

// GetNodeAllocationDecisions returns the NodeAllocationDecisions field value if set, zero value otherwise.
func (o *ElasticsearchAllocationExplanation) GetNodeAllocationDecisions() []ElasticsearchNodeAllocationDecision {
	if o == nil || o.NodeAllocationDecisions == nil {
		var ret []ElasticsearchNodeAllocationDecision
		return ret
	}
	return o.NodeAllocationDecisions
}

// GetNodeAllocationDecisionsOk returns a tuple with the NodeAllocationDecisions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticsearchAllocationExplanation) GetNodeAllocationDecisionsOk() (*[]ElasticsearchNodeAllocationDecision, bool) {
	if o == nil || o.NodeAllocationDecisions == nil {
		return nil, false
	}
	return &o.NodeAllocationDecisions, true
}

// HasNodeAllocationDecisions returns a boolean if a field has been set.
func (o *ElasticsearchAllocationExplanation) HasNodeAllocationDecisions() bool {
	return o != nil && o.NodeAllocationDecisions != nil
}

// SetNodeAllocationDecisions gets a reference to the given []ElasticsearchNodeAllocationDecision and assigns it to the NodeAllocationDecisions field.
func (o *ElasticsearchAllocationExplanation) SetNodeAllocationDecisions(v []ElasticsearchNodeAllocationDecision) {
	o.NodeAllocationDecisions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ElasticsearchAllocationExplanation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Index.IsSet() {
		toSerialize["index"] = o.Index.Get()
	}
	if o.Shard.IsSet() {
		toSerialize["shard"] = o.Shard.Get()
	}
	if o.Primary.IsSet() {
		toSerialize["primary"] = o.Primary.Get()
	}
	if o.CurrentState.IsSet() {
		toSerialize["currentState"] = o.CurrentState.Get()
	}
	if o.CurrentNode != nil {
		toSerialize["currentNode"] = o.CurrentNode
	}
	if o.UnassignedInfo != nil {
		toSerialize["unassignedInfo"] = o.UnassignedInfo
	}
	if o.CanAllocate.IsSet() {
		toSerialize["canAllocate"] = o.CanAllocate.Get()
	}
	if o.AllocateExplanation.IsSet() {
		toSerialize["allocateExplanation"] = o.AllocateExplanation.Get()
	}
	if o.CanRemainOnCurrentNode.IsSet() {
		toSerialize["canRemainOnCurrentNode"] = o.CanRemainOnCurrentNode.Get()
	}
	if o.CanRemainDecisions != nil {
		toSerialize["canRemainDecisions"] = o.CanRemainDecisions
	}
	if o.CanRebalanceCluster.IsSet() {
		toSerialize["canRebalanceCluster"] = o.CanRebalanceCluster.Get()
	}
	if o.CanRebalanceClusterDecisions != nil {
		toSerialize["canRebalanceClusterDecisions"] = o.CanRebalanceClusterDecisions
	}
	if o.CanRebalanceToOtherNode.IsSet() {
		toSerialize["canRebalanceToOtherNode"] = o.CanRebalanceToOtherNode.Get()
	}
	if o.RebalanceExplanation.IsSet() {
		toSerialize["rebalanceExplanation"] = o.RebalanceExplanation.Get()
	}
	if o.NodeAllocationDecisions != nil {
		toSerialize["nodeAllocationDecisions"] = o.NodeAllocationDecisions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ElasticsearchAllocationExplanation) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Index                        common.NullableString                 `json:"index,omitempty"`
		Shard                        common.NullableInt64                  `json:"shard,omitempty"`
		Primary                      common.NullableBool                   `json:"primary,omitempty"`
		CurrentState                 common.NullableString                 `json:"currentState,omitempty"`
		CurrentNode                  *ElasticsearchShardNode               `json:"currentNode,omitempty"`
		UnassignedInfo               *ElasticsearchUnassignedInfo          `json:"unassignedInfo,omitempty"`
		CanAllocate                  common.NullableString                 `json:"canAllocate,omitempty"`
		AllocateExplanation          common.NullableString                 `json:"allocateExplanation,omitempty"`
		CanRemainOnCurrentNode       common.NullableString                 `json:"canRemainOnCurrentNode,omitempty"`
		CanRemainDecisions           []ElasticsearchAllocationDecider      `json:"canRemainDecisions,omitempty"`
		CanRebalanceCluster          common.NullableString                 `json:"canRebalanceCluster,omitempty"`
		CanRebalanceClusterDecisions []ElasticsearchAllocationDecider      `json:"canRebalanceClusterDecisions,omitempty"`
		CanRebalanceToOtherNode      common.NullableString                 `json:"canRebalanceToOtherNode,omitempty"`
		RebalanceExplanation         common.NullableString                 `json:"rebalanceExplanation,omitempty"`
		NodeAllocationDecisions      []ElasticsearchNodeAllocationDecision `json:"nodeAllocationDecisions,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"index", "shard", "primary", "currentState", "currentNode", "unassignedInfo", "canAllocate", "allocateExplanation", "canRemainOnCurrentNode", "canRemainDecisions", "canRebalanceCluster", "canRebalanceClusterDecisions", "canRebalanceToOtherNode", "rebalanceExplanation", "nodeAllocationDecisions"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Index = all.Index
	o.Shard = all.Shard
	o.Primary = all.Primary
	o.CurrentState = all.CurrentState
	if all.CurrentNode != nil && all.CurrentNode.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CurrentNode = all.CurrentNode
	if all.UnassignedInfo != nil && all.UnassignedInfo.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.UnassignedInfo = all.UnassignedInfo
	o.CanAllocate = all.CanAllocate
	o.AllocateExplanation = all.AllocateExplanation
	o.CanRemainOnCurrentNode = all.CanRemainOnCurrentNode
	o.CanRemainDecisions = all.CanRemainDecisions
	o.CanRebalanceCluster = all.CanRebalanceCluster
	o.CanRebalanceClusterDecisions = all.CanRebalanceClusterDecisions
	o.CanRebalanceToOtherNode = all.CanRebalanceToOtherNode
	o.RebalanceExplanation = all.RebalanceExplanation
	o.NodeAllocationDecisions = all.NodeAllocationDecisions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
