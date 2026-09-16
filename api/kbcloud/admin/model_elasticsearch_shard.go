// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ElasticsearchShard struct {
	Index            common.NullableString `json:"index,omitempty"`
	Shard            common.NullableInt64  `json:"shard,omitempty"`
	Primary          common.NullableBool   `json:"primary,omitempty"`
	State            common.NullableString `json:"state,omitempty"`
	Documents        common.NullableInt64  `json:"documents,omitempty"`
	StoreBytes       common.NullableInt64  `json:"storeBytes,omitempty"`
	NodeId           common.NullableString `json:"nodeId,omitempty"`
	NodeName         common.NullableString `json:"nodeName,omitempty"`
	UnassignedReason common.NullableString `json:"unassignedReason,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewElasticsearchShard instantiates a new ElasticsearchShard object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewElasticsearchShard() *ElasticsearchShard {
	this := ElasticsearchShard{}
	return &this
}

// NewElasticsearchShardWithDefaults instantiates a new ElasticsearchShard object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewElasticsearchShardWithDefaults() *ElasticsearchShard {
	this := ElasticsearchShard{}
	return &this
}

// GetIndex returns the Index field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElasticsearchShard) GetIndex() string {
	if o == nil || o.Index.Get() == nil {
		var ret string
		return ret
	}
	return *o.Index.Get()
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchShard) GetIndexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Index.Get(), o.Index.IsSet()
}

// HasIndex returns a boolean if a field has been set.
func (o *ElasticsearchShard) HasIndex() bool {
	return o != nil && o.Index.IsSet()
}

// SetIndex gets a reference to the given common.NullableString and assigns it to the Index field.
func (o *ElasticsearchShard) SetIndex(v string) {
	o.Index.Set(&v)
}

// SetIndexNil sets the value for Index to be an explicit nil.
func (o *ElasticsearchShard) SetIndexNil() {
	o.Index.Set(nil)
}

// UnsetIndex ensures that no value is present for Index, not even an explicit nil.
func (o *ElasticsearchShard) UnsetIndex() {
	o.Index.Unset()
}

// GetShard returns the Shard field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElasticsearchShard) GetShard() int64 {
	if o == nil || o.Shard.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Shard.Get()
}

// GetShardOk returns a tuple with the Shard field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchShard) GetShardOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Shard.Get(), o.Shard.IsSet()
}

// HasShard returns a boolean if a field has been set.
func (o *ElasticsearchShard) HasShard() bool {
	return o != nil && o.Shard.IsSet()
}

// SetShard gets a reference to the given common.NullableInt64 and assigns it to the Shard field.
func (o *ElasticsearchShard) SetShard(v int64) {
	o.Shard.Set(&v)
}

// SetShardNil sets the value for Shard to be an explicit nil.
func (o *ElasticsearchShard) SetShardNil() {
	o.Shard.Set(nil)
}

// UnsetShard ensures that no value is present for Shard, not even an explicit nil.
func (o *ElasticsearchShard) UnsetShard() {
	o.Shard.Unset()
}

// GetPrimary returns the Primary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElasticsearchShard) GetPrimary() bool {
	if o == nil || o.Primary.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Primary.Get()
}

// GetPrimaryOk returns a tuple with the Primary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchShard) GetPrimaryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Primary.Get(), o.Primary.IsSet()
}

// HasPrimary returns a boolean if a field has been set.
func (o *ElasticsearchShard) HasPrimary() bool {
	return o != nil && o.Primary.IsSet()
}

// SetPrimary gets a reference to the given common.NullableBool and assigns it to the Primary field.
func (o *ElasticsearchShard) SetPrimary(v bool) {
	o.Primary.Set(&v)
}

// SetPrimaryNil sets the value for Primary to be an explicit nil.
func (o *ElasticsearchShard) SetPrimaryNil() {
	o.Primary.Set(nil)
}

// UnsetPrimary ensures that no value is present for Primary, not even an explicit nil.
func (o *ElasticsearchShard) UnsetPrimary() {
	o.Primary.Unset()
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElasticsearchShard) GetState() string {
	if o == nil || o.State.Get() == nil {
		var ret string
		return ret
	}
	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchShard) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// HasState returns a boolean if a field has been set.
func (o *ElasticsearchShard) HasState() bool {
	return o != nil && o.State.IsSet()
}

// SetState gets a reference to the given common.NullableString and assigns it to the State field.
func (o *ElasticsearchShard) SetState(v string) {
	o.State.Set(&v)
}

// SetStateNil sets the value for State to be an explicit nil.
func (o *ElasticsearchShard) SetStateNil() {
	o.State.Set(nil)
}

// UnsetState ensures that no value is present for State, not even an explicit nil.
func (o *ElasticsearchShard) UnsetState() {
	o.State.Unset()
}

// GetDocuments returns the Documents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElasticsearchShard) GetDocuments() int64 {
	if o == nil || o.Documents.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Documents.Get()
}

// GetDocumentsOk returns a tuple with the Documents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchShard) GetDocumentsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Documents.Get(), o.Documents.IsSet()
}

// HasDocuments returns a boolean if a field has been set.
func (o *ElasticsearchShard) HasDocuments() bool {
	return o != nil && o.Documents.IsSet()
}

// SetDocuments gets a reference to the given common.NullableInt64 and assigns it to the Documents field.
func (o *ElasticsearchShard) SetDocuments(v int64) {
	o.Documents.Set(&v)
}

// SetDocumentsNil sets the value for Documents to be an explicit nil.
func (o *ElasticsearchShard) SetDocumentsNil() {
	o.Documents.Set(nil)
}

// UnsetDocuments ensures that no value is present for Documents, not even an explicit nil.
func (o *ElasticsearchShard) UnsetDocuments() {
	o.Documents.Unset()
}

// GetStoreBytes returns the StoreBytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElasticsearchShard) GetStoreBytes() int64 {
	if o == nil || o.StoreBytes.Get() == nil {
		var ret int64
		return ret
	}
	return *o.StoreBytes.Get()
}

// GetStoreBytesOk returns a tuple with the StoreBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchShard) GetStoreBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.StoreBytes.Get(), o.StoreBytes.IsSet()
}

// HasStoreBytes returns a boolean if a field has been set.
func (o *ElasticsearchShard) HasStoreBytes() bool {
	return o != nil && o.StoreBytes.IsSet()
}

// SetStoreBytes gets a reference to the given common.NullableInt64 and assigns it to the StoreBytes field.
func (o *ElasticsearchShard) SetStoreBytes(v int64) {
	o.StoreBytes.Set(&v)
}

// SetStoreBytesNil sets the value for StoreBytes to be an explicit nil.
func (o *ElasticsearchShard) SetStoreBytesNil() {
	o.StoreBytes.Set(nil)
}

// UnsetStoreBytes ensures that no value is present for StoreBytes, not even an explicit nil.
func (o *ElasticsearchShard) UnsetStoreBytes() {
	o.StoreBytes.Unset()
}

// GetNodeId returns the NodeId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElasticsearchShard) GetNodeId() string {
	if o == nil || o.NodeId.Get() == nil {
		var ret string
		return ret
	}
	return *o.NodeId.Get()
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchShard) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NodeId.Get(), o.NodeId.IsSet()
}

// HasNodeId returns a boolean if a field has been set.
func (o *ElasticsearchShard) HasNodeId() bool {
	return o != nil && o.NodeId.IsSet()
}

// SetNodeId gets a reference to the given common.NullableString and assigns it to the NodeId field.
func (o *ElasticsearchShard) SetNodeId(v string) {
	o.NodeId.Set(&v)
}

// SetNodeIdNil sets the value for NodeId to be an explicit nil.
func (o *ElasticsearchShard) SetNodeIdNil() {
	o.NodeId.Set(nil)
}

// UnsetNodeId ensures that no value is present for NodeId, not even an explicit nil.
func (o *ElasticsearchShard) UnsetNodeId() {
	o.NodeId.Unset()
}

// GetNodeName returns the NodeName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElasticsearchShard) GetNodeName() string {
	if o == nil || o.NodeName.Get() == nil {
		var ret string
		return ret
	}
	return *o.NodeName.Get()
}

// GetNodeNameOk returns a tuple with the NodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchShard) GetNodeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NodeName.Get(), o.NodeName.IsSet()
}

// HasNodeName returns a boolean if a field has been set.
func (o *ElasticsearchShard) HasNodeName() bool {
	return o != nil && o.NodeName.IsSet()
}

// SetNodeName gets a reference to the given common.NullableString and assigns it to the NodeName field.
func (o *ElasticsearchShard) SetNodeName(v string) {
	o.NodeName.Set(&v)
}

// SetNodeNameNil sets the value for NodeName to be an explicit nil.
func (o *ElasticsearchShard) SetNodeNameNil() {
	o.NodeName.Set(nil)
}

// UnsetNodeName ensures that no value is present for NodeName, not even an explicit nil.
func (o *ElasticsearchShard) UnsetNodeName() {
	o.NodeName.Unset()
}

// GetUnassignedReason returns the UnassignedReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElasticsearchShard) GetUnassignedReason() string {
	if o == nil || o.UnassignedReason.Get() == nil {
		var ret string
		return ret
	}
	return *o.UnassignedReason.Get()
}

// GetUnassignedReasonOk returns a tuple with the UnassignedReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchShard) GetUnassignedReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UnassignedReason.Get(), o.UnassignedReason.IsSet()
}

// HasUnassignedReason returns a boolean if a field has been set.
func (o *ElasticsearchShard) HasUnassignedReason() bool {
	return o != nil && o.UnassignedReason.IsSet()
}

// SetUnassignedReason gets a reference to the given common.NullableString and assigns it to the UnassignedReason field.
func (o *ElasticsearchShard) SetUnassignedReason(v string) {
	o.UnassignedReason.Set(&v)
}

// SetUnassignedReasonNil sets the value for UnassignedReason to be an explicit nil.
func (o *ElasticsearchShard) SetUnassignedReasonNil() {
	o.UnassignedReason.Set(nil)
}

// UnsetUnassignedReason ensures that no value is present for UnassignedReason, not even an explicit nil.
func (o *ElasticsearchShard) UnsetUnassignedReason() {
	o.UnassignedReason.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ElasticsearchShard) MarshalJSON() ([]byte, error) {
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
	if o.State.IsSet() {
		toSerialize["state"] = o.State.Get()
	}
	if o.Documents.IsSet() {
		toSerialize["documents"] = o.Documents.Get()
	}
	if o.StoreBytes.IsSet() {
		toSerialize["storeBytes"] = o.StoreBytes.Get()
	}
	if o.NodeId.IsSet() {
		toSerialize["nodeId"] = o.NodeId.Get()
	}
	if o.NodeName.IsSet() {
		toSerialize["nodeName"] = o.NodeName.Get()
	}
	if o.UnassignedReason.IsSet() {
		toSerialize["unassignedReason"] = o.UnassignedReason.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ElasticsearchShard) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Index            common.NullableString `json:"index,omitempty"`
		Shard            common.NullableInt64  `json:"shard,omitempty"`
		Primary          common.NullableBool   `json:"primary,omitempty"`
		State            common.NullableString `json:"state,omitempty"`
		Documents        common.NullableInt64  `json:"documents,omitempty"`
		StoreBytes       common.NullableInt64  `json:"storeBytes,omitempty"`
		NodeId           common.NullableString `json:"nodeId,omitempty"`
		NodeName         common.NullableString `json:"nodeName,omitempty"`
		UnassignedReason common.NullableString `json:"unassignedReason,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"index", "shard", "primary", "state", "documents", "storeBytes", "nodeId", "nodeName", "unassignedReason"})
	} else {
		return err
	}
	o.Index = all.Index
	o.Shard = all.Shard
	o.Primary = all.Primary
	o.State = all.State
	o.Documents = all.Documents
	o.StoreBytes = all.StoreBytes
	o.NodeId = all.NodeId
	o.NodeName = all.NodeName
	o.UnassignedReason = all.UnassignedReason

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
