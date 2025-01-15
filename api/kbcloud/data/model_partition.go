// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package data

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type Partition struct {
	Id       int32        `json:"id"`
	Leader   BrokerNode   `json:"leader"`
	Replicas []BrokerNode `json:"replicas,omitempty"`
	// In-Sync Replicas
	Isr             []BrokerNode `json:"isr,omitempty"`
	BeginningOffset *int64       `json:"beginningOffset,omitempty"`
	EndOffset       *int64       `json:"endOffset,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPartition instantiates a new Partition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPartition(id int32, leader BrokerNode) *Partition {
	this := Partition{}
	this.Id = id
	this.Leader = leader
	return &this
}

// NewPartitionWithDefaults instantiates a new Partition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPartitionWithDefaults() *Partition {
	this := Partition{}
	return &this
}

// GetId returns the Id field value.
func (o *Partition) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Partition) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *Partition) SetId(v int32) {
	o.Id = v
}

// GetLeader returns the Leader field value.
func (o *Partition) GetLeader() BrokerNode {
	if o == nil {
		var ret BrokerNode
		return ret
	}
	return o.Leader
}

// GetLeaderOk returns a tuple with the Leader field value
// and a boolean to check if the value has been set.
func (o *Partition) GetLeaderOk() (*BrokerNode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Leader, true
}

// SetLeader sets field value.
func (o *Partition) SetLeader(v BrokerNode) {
	o.Leader = v
}

// GetReplicas returns the Replicas field value if set, zero value otherwise.
func (o *Partition) GetReplicas() []BrokerNode {
	if o == nil || o.Replicas == nil {
		var ret []BrokerNode
		return ret
	}
	return o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Partition) GetReplicasOk() (*[]BrokerNode, bool) {
	if o == nil || o.Replicas == nil {
		return nil, false
	}
	return &o.Replicas, true
}

// HasReplicas returns a boolean if a field has been set.
func (o *Partition) HasReplicas() bool {
	return o != nil && o.Replicas != nil
}

// SetReplicas gets a reference to the given []BrokerNode and assigns it to the Replicas field.
func (o *Partition) SetReplicas(v []BrokerNode) {
	o.Replicas = v
}

// GetIsr returns the Isr field value if set, zero value otherwise.
func (o *Partition) GetIsr() []BrokerNode {
	if o == nil || o.Isr == nil {
		var ret []BrokerNode
		return ret
	}
	return o.Isr
}

// GetIsrOk returns a tuple with the Isr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Partition) GetIsrOk() (*[]BrokerNode, bool) {
	if o == nil || o.Isr == nil {
		return nil, false
	}
	return &o.Isr, true
}

// HasIsr returns a boolean if a field has been set.
func (o *Partition) HasIsr() bool {
	return o != nil && o.Isr != nil
}

// SetIsr gets a reference to the given []BrokerNode and assigns it to the Isr field.
func (o *Partition) SetIsr(v []BrokerNode) {
	o.Isr = v
}

// GetBeginningOffset returns the BeginningOffset field value if set, zero value otherwise.
func (o *Partition) GetBeginningOffset() int64 {
	if o == nil || o.BeginningOffset == nil {
		var ret int64
		return ret
	}
	return *o.BeginningOffset
}

// GetBeginningOffsetOk returns a tuple with the BeginningOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Partition) GetBeginningOffsetOk() (*int64, bool) {
	if o == nil || o.BeginningOffset == nil {
		return nil, false
	}
	return o.BeginningOffset, true
}

// HasBeginningOffset returns a boolean if a field has been set.
func (o *Partition) HasBeginningOffset() bool {
	return o != nil && o.BeginningOffset != nil
}

// SetBeginningOffset gets a reference to the given int64 and assigns it to the BeginningOffset field.
func (o *Partition) SetBeginningOffset(v int64) {
	o.BeginningOffset = &v
}

// GetEndOffset returns the EndOffset field value if set, zero value otherwise.
func (o *Partition) GetEndOffset() int64 {
	if o == nil || o.EndOffset == nil {
		var ret int64
		return ret
	}
	return *o.EndOffset
}

// GetEndOffsetOk returns a tuple with the EndOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Partition) GetEndOffsetOk() (*int64, bool) {
	if o == nil || o.EndOffset == nil {
		return nil, false
	}
	return o.EndOffset, true
}

// HasEndOffset returns a boolean if a field has been set.
func (o *Partition) HasEndOffset() bool {
	return o != nil && o.EndOffset != nil
}

// SetEndOffset gets a reference to the given int64 and assigns it to the EndOffset field.
func (o *Partition) SetEndOffset(v int64) {
	o.EndOffset = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o Partition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["leader"] = o.Leader
	if o.Replicas != nil {
		toSerialize["replicas"] = o.Replicas
	}
	if o.Isr != nil {
		toSerialize["isr"] = o.Isr
	}
	if o.BeginningOffset != nil {
		toSerialize["beginningOffset"] = o.BeginningOffset
	}
	if o.EndOffset != nil {
		toSerialize["endOffset"] = o.EndOffset
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Partition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id              *int32       `json:"id"`
		Leader          *BrokerNode  `json:"leader"`
		Replicas        []BrokerNode `json:"replicas,omitempty"`
		Isr             []BrokerNode `json:"isr,omitempty"`
		BeginningOffset *int64       `json:"beginningOffset,omitempty"`
		EndOffset       *int64       `json:"endOffset,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Leader == nil {
		return fmt.Errorf("required field leader missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "leader", "replicas", "isr", "beginningOffset", "endOffset"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = *all.Id
	if all.Leader.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Leader = *all.Leader
	o.Replicas = all.Replicas
	o.Isr = all.Isr
	o.BeginningOffset = all.BeginningOffset
	o.EndOffset = all.EndOffset

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
