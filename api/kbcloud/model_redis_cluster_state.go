// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type RedisClusterState struct {
	Ready *bool `json:"ready,omitempty"`
	// reason code when cluster data management is not ready, for example requires_redis_cluster_mirror
	Reason      *string            `json:"reason,omitempty"`
	NetworkMode *string            `json:"networkMode,omitempty"`
	MirrorReady *bool              `json:"mirrorReady,omitempty"`
	Nodes       []RedisClusterNode `json:"nodes,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRedisClusterState instantiates a new RedisClusterState object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRedisClusterState() *RedisClusterState {
	this := RedisClusterState{}
	return &this
}

// NewRedisClusterStateWithDefaults instantiates a new RedisClusterState object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRedisClusterStateWithDefaults() *RedisClusterState {
	this := RedisClusterState{}
	return &this
}

// GetReady returns the Ready field value if set, zero value otherwise.
func (o *RedisClusterState) GetReady() bool {
	if o == nil || o.Ready == nil {
		var ret bool
		return ret
	}
	return *o.Ready
}

// GetReadyOk returns a tuple with the Ready field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisClusterState) GetReadyOk() (*bool, bool) {
	if o == nil || o.Ready == nil {
		return nil, false
	}
	return o.Ready, true
}

// HasReady returns a boolean if a field has been set.
func (o *RedisClusterState) HasReady() bool {
	return o != nil && o.Ready != nil
}

// SetReady gets a reference to the given bool and assigns it to the Ready field.
func (o *RedisClusterState) SetReady(v bool) {
	o.Ready = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *RedisClusterState) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisClusterState) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *RedisClusterState) HasReason() bool {
	return o != nil && o.Reason != nil
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *RedisClusterState) SetReason(v string) {
	o.Reason = &v
}

// GetNetworkMode returns the NetworkMode field value if set, zero value otherwise.
func (o *RedisClusterState) GetNetworkMode() string {
	if o == nil || o.NetworkMode == nil {
		var ret string
		return ret
	}
	return *o.NetworkMode
}

// GetNetworkModeOk returns a tuple with the NetworkMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisClusterState) GetNetworkModeOk() (*string, bool) {
	if o == nil || o.NetworkMode == nil {
		return nil, false
	}
	return o.NetworkMode, true
}

// HasNetworkMode returns a boolean if a field has been set.
func (o *RedisClusterState) HasNetworkMode() bool {
	return o != nil && o.NetworkMode != nil
}

// SetNetworkMode gets a reference to the given string and assigns it to the NetworkMode field.
func (o *RedisClusterState) SetNetworkMode(v string) {
	o.NetworkMode = &v
}

// GetMirrorReady returns the MirrorReady field value if set, zero value otherwise.
func (o *RedisClusterState) GetMirrorReady() bool {
	if o == nil || o.MirrorReady == nil {
		var ret bool
		return ret
	}
	return *o.MirrorReady
}

// GetMirrorReadyOk returns a tuple with the MirrorReady field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisClusterState) GetMirrorReadyOk() (*bool, bool) {
	if o == nil || o.MirrorReady == nil {
		return nil, false
	}
	return o.MirrorReady, true
}

// HasMirrorReady returns a boolean if a field has been set.
func (o *RedisClusterState) HasMirrorReady() bool {
	return o != nil && o.MirrorReady != nil
}

// SetMirrorReady gets a reference to the given bool and assigns it to the MirrorReady field.
func (o *RedisClusterState) SetMirrorReady(v bool) {
	o.MirrorReady = &v
}

// GetNodes returns the Nodes field value if set, zero value otherwise.
func (o *RedisClusterState) GetNodes() []RedisClusterNode {
	if o == nil || o.Nodes == nil {
		var ret []RedisClusterNode
		return ret
	}
	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisClusterState) GetNodesOk() (*[]RedisClusterNode, bool) {
	if o == nil || o.Nodes == nil {
		return nil, false
	}
	return &o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *RedisClusterState) HasNodes() bool {
	return o != nil && o.Nodes != nil
}

// SetNodes gets a reference to the given []RedisClusterNode and assigns it to the Nodes field.
func (o *RedisClusterState) SetNodes(v []RedisClusterNode) {
	o.Nodes = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RedisClusterState) MarshalJSON() ([]byte, error) {
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
	if o.NetworkMode != nil {
		toSerialize["networkMode"] = o.NetworkMode
	}
	if o.MirrorReady != nil {
		toSerialize["mirrorReady"] = o.MirrorReady
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
func (o *RedisClusterState) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Ready       *bool              `json:"ready,omitempty"`
		Reason      *string            `json:"reason,omitempty"`
		NetworkMode *string            `json:"networkMode,omitempty"`
		MirrorReady *bool              `json:"mirrorReady,omitempty"`
		Nodes       []RedisClusterNode `json:"nodes,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"ready", "reason", "networkMode", "mirrorReady", "nodes"})
	} else {
		return err
	}
	o.Ready = all.Ready
	o.Reason = all.Reason
	o.NetworkMode = all.NetworkMode
	o.MirrorReady = all.MirrorReady
	o.Nodes = all.Nodes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
