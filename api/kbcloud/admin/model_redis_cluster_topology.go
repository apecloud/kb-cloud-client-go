// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type RedisClusterTopology struct {
	Ready  *bool              `json:"ready,omitempty"`
	Reason *string            `json:"reason,omitempty"`
	Nodes  []RedisClusterNode `json:"nodes,omitempty"`
	Slots  []RedisClusterSlot `json:"slots,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRedisClusterTopology instantiates a new RedisClusterTopology object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRedisClusterTopology() *RedisClusterTopology {
	this := RedisClusterTopology{}
	return &this
}

// NewRedisClusterTopologyWithDefaults instantiates a new RedisClusterTopology object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRedisClusterTopologyWithDefaults() *RedisClusterTopology {
	this := RedisClusterTopology{}
	return &this
}

// GetReady returns the Ready field value if set, zero value otherwise.
func (o *RedisClusterTopology) GetReady() bool {
	if o == nil || o.Ready == nil {
		var ret bool
		return ret
	}
	return *o.Ready
}

// GetReadyOk returns a tuple with the Ready field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisClusterTopology) GetReadyOk() (*bool, bool) {
	if o == nil || o.Ready == nil {
		return nil, false
	}
	return o.Ready, true
}

// HasReady returns a boolean if a field has been set.
func (o *RedisClusterTopology) HasReady() bool {
	return o != nil && o.Ready != nil
}

// SetReady gets a reference to the given bool and assigns it to the Ready field.
func (o *RedisClusterTopology) SetReady(v bool) {
	o.Ready = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *RedisClusterTopology) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisClusterTopology) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *RedisClusterTopology) HasReason() bool {
	return o != nil && o.Reason != nil
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *RedisClusterTopology) SetReason(v string) {
	o.Reason = &v
}

// GetNodes returns the Nodes field value if set, zero value otherwise.
func (o *RedisClusterTopology) GetNodes() []RedisClusterNode {
	if o == nil || o.Nodes == nil {
		var ret []RedisClusterNode
		return ret
	}
	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisClusterTopology) GetNodesOk() (*[]RedisClusterNode, bool) {
	if o == nil || o.Nodes == nil {
		return nil, false
	}
	return &o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *RedisClusterTopology) HasNodes() bool {
	return o != nil && o.Nodes != nil
}

// SetNodes gets a reference to the given []RedisClusterNode and assigns it to the Nodes field.
func (o *RedisClusterTopology) SetNodes(v []RedisClusterNode) {
	o.Nodes = v
}

// GetSlots returns the Slots field value if set, zero value otherwise.
func (o *RedisClusterTopology) GetSlots() []RedisClusterSlot {
	if o == nil || o.Slots == nil {
		var ret []RedisClusterSlot
		return ret
	}
	return o.Slots
}

// GetSlotsOk returns a tuple with the Slots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisClusterTopology) GetSlotsOk() (*[]RedisClusterSlot, bool) {
	if o == nil || o.Slots == nil {
		return nil, false
	}
	return &o.Slots, true
}

// HasSlots returns a boolean if a field has been set.
func (o *RedisClusterTopology) HasSlots() bool {
	return o != nil && o.Slots != nil
}

// SetSlots gets a reference to the given []RedisClusterSlot and assigns it to the Slots field.
func (o *RedisClusterTopology) SetSlots(v []RedisClusterSlot) {
	o.Slots = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RedisClusterTopology) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Ready != nil {
		toSerialize["ready"] = o.Ready
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.Nodes != nil {
		toSerialize["nodes"] = o.Nodes
	}
	if o.Slots != nil {
		toSerialize["slots"] = o.Slots
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RedisClusterTopology) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Ready  *bool              `json:"ready,omitempty"`
		Reason *string            `json:"reason,omitempty"`
		Nodes  []RedisClusterNode `json:"nodes,omitempty"`
		Slots  []RedisClusterSlot `json:"slots,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"ready", "reason", "nodes", "slots"})
	} else {
		return err
	}
	o.Ready = all.Ready
	o.Reason = all.Reason
	o.Nodes = all.Nodes
	o.Slots = all.Slots

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
