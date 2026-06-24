// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type RedisClusterSlot struct {
	Start          *int64   `json:"start,omitempty"`
	End            *int64   `json:"end,omitempty"`
	PrimaryNodeId  *string  `json:"primaryNodeId,omitempty"`
	ReplicaNodeIds []string `json:"replicaNodeIds,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRedisClusterSlot instantiates a new RedisClusterSlot object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRedisClusterSlot() *RedisClusterSlot {
	this := RedisClusterSlot{}
	return &this
}

// NewRedisClusterSlotWithDefaults instantiates a new RedisClusterSlot object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRedisClusterSlotWithDefaults() *RedisClusterSlot {
	this := RedisClusterSlot{}
	return &this
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *RedisClusterSlot) GetStart() int64 {
	if o == nil || o.Start == nil {
		var ret int64
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisClusterSlot) GetStartOk() (*int64, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *RedisClusterSlot) HasStart() bool {
	return o != nil && o.Start != nil
}

// SetStart gets a reference to the given int64 and assigns it to the Start field.
func (o *RedisClusterSlot) SetStart(v int64) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *RedisClusterSlot) GetEnd() int64 {
	if o == nil || o.End == nil {
		var ret int64
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisClusterSlot) GetEndOk() (*int64, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *RedisClusterSlot) HasEnd() bool {
	return o != nil && o.End != nil
}

// SetEnd gets a reference to the given int64 and assigns it to the End field.
func (o *RedisClusterSlot) SetEnd(v int64) {
	o.End = &v
}

// GetPrimaryNodeId returns the PrimaryNodeId field value if set, zero value otherwise.
func (o *RedisClusterSlot) GetPrimaryNodeId() string {
	if o == nil || o.PrimaryNodeId == nil {
		var ret string
		return ret
	}
	return *o.PrimaryNodeId
}

// GetPrimaryNodeIdOk returns a tuple with the PrimaryNodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisClusterSlot) GetPrimaryNodeIdOk() (*string, bool) {
	if o == nil || o.PrimaryNodeId == nil {
		return nil, false
	}
	return o.PrimaryNodeId, true
}

// HasPrimaryNodeId returns a boolean if a field has been set.
func (o *RedisClusterSlot) HasPrimaryNodeId() bool {
	return o != nil && o.PrimaryNodeId != nil
}

// SetPrimaryNodeId gets a reference to the given string and assigns it to the PrimaryNodeId field.
func (o *RedisClusterSlot) SetPrimaryNodeId(v string) {
	o.PrimaryNodeId = &v
}

// GetReplicaNodeIds returns the ReplicaNodeIds field value if set, zero value otherwise.
func (o *RedisClusterSlot) GetReplicaNodeIds() []string {
	if o == nil || o.ReplicaNodeIds == nil {
		var ret []string
		return ret
	}
	return o.ReplicaNodeIds
}

// GetReplicaNodeIdsOk returns a tuple with the ReplicaNodeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisClusterSlot) GetReplicaNodeIdsOk() (*[]string, bool) {
	if o == nil || o.ReplicaNodeIds == nil {
		return nil, false
	}
	return &o.ReplicaNodeIds, true
}

// HasReplicaNodeIds returns a boolean if a field has been set.
func (o *RedisClusterSlot) HasReplicaNodeIds() bool {
	return o != nil && o.ReplicaNodeIds != nil
}

// SetReplicaNodeIds gets a reference to the given []string and assigns it to the ReplicaNodeIds field.
func (o *RedisClusterSlot) SetReplicaNodeIds(v []string) {
	o.ReplicaNodeIds = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RedisClusterSlot) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Start != nil {
		toSerialize["start"] = o.Start
	}
	if o.End != nil {
		toSerialize["end"] = o.End
	}
	if o.PrimaryNodeId != nil {
		toSerialize["primaryNodeId"] = o.PrimaryNodeId
	}
	if o.ReplicaNodeIds != nil {
		toSerialize["replicaNodeIds"] = o.ReplicaNodeIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RedisClusterSlot) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Start          *int64   `json:"start,omitempty"`
		End            *int64   `json:"end,omitempty"`
		PrimaryNodeId  *string  `json:"primaryNodeId,omitempty"`
		ReplicaNodeIds []string `json:"replicaNodeIds,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"start", "end", "primaryNodeId", "replicaNodeIds"})
	} else {
		return err
	}
	o.Start = all.Start
	o.End = all.End
	o.PrimaryNodeId = all.PrimaryNodeId
	o.ReplicaNodeIds = all.ReplicaNodeIds

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
