// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type RedisClusterNode struct {
	NodeId    *string `json:"nodeId,omitempty"`
	Address   *string `json:"address,omitempty"`
	Role      *string `json:"role,omitempty"`
	ShardId   *string `json:"shardId,omitempty"`
	Reachable *bool   `json:"reachable,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRedisClusterNode instantiates a new RedisClusterNode object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRedisClusterNode() *RedisClusterNode {
	this := RedisClusterNode{}
	return &this
}

// NewRedisClusterNodeWithDefaults instantiates a new RedisClusterNode object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRedisClusterNodeWithDefaults() *RedisClusterNode {
	this := RedisClusterNode{}
	return &this
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *RedisClusterNode) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisClusterNode) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *RedisClusterNode) HasNodeId() bool {
	return o != nil && o.NodeId != nil
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *RedisClusterNode) SetNodeId(v string) {
	o.NodeId = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *RedisClusterNode) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisClusterNode) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *RedisClusterNode) HasAddress() bool {
	return o != nil && o.Address != nil
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *RedisClusterNode) SetAddress(v string) {
	o.Address = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *RedisClusterNode) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisClusterNode) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *RedisClusterNode) HasRole() bool {
	return o != nil && o.Role != nil
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *RedisClusterNode) SetRole(v string) {
	o.Role = &v
}

// GetShardId returns the ShardId field value if set, zero value otherwise.
func (o *RedisClusterNode) GetShardId() string {
	if o == nil || o.ShardId == nil {
		var ret string
		return ret
	}
	return *o.ShardId
}

// GetShardIdOk returns a tuple with the ShardId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisClusterNode) GetShardIdOk() (*string, bool) {
	if o == nil || o.ShardId == nil {
		return nil, false
	}
	return o.ShardId, true
}

// HasShardId returns a boolean if a field has been set.
func (o *RedisClusterNode) HasShardId() bool {
	return o != nil && o.ShardId != nil
}

// SetShardId gets a reference to the given string and assigns it to the ShardId field.
func (o *RedisClusterNode) SetShardId(v string) {
	o.ShardId = &v
}

// GetReachable returns the Reachable field value if set, zero value otherwise.
func (o *RedisClusterNode) GetReachable() bool {
	if o == nil || o.Reachable == nil {
		var ret bool
		return ret
	}
	return *o.Reachable
}

// GetReachableOk returns a tuple with the Reachable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisClusterNode) GetReachableOk() (*bool, bool) {
	if o == nil || o.Reachable == nil {
		return nil, false
	}
	return o.Reachable, true
}

// HasReachable returns a boolean if a field has been set.
func (o *RedisClusterNode) HasReachable() bool {
	return o != nil && o.Reachable != nil
}

// SetReachable gets a reference to the given bool and assigns it to the Reachable field.
func (o *RedisClusterNode) SetReachable(v bool) {
	o.Reachable = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RedisClusterNode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.ShardId != nil {
		toSerialize["shardId"] = o.ShardId
	}
	if o.Reachable != nil {
		toSerialize["reachable"] = o.Reachable
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RedisClusterNode) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NodeId    *string `json:"nodeId,omitempty"`
		Address   *string `json:"address,omitempty"`
		Role      *string `json:"role,omitempty"`
		ShardId   *string `json:"shardId,omitempty"`
		Reachable *bool   `json:"reachable,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"nodeId", "address", "role", "shardId", "reachable"})
	} else {
		return err
	}
	o.NodeId = all.NodeId
	o.Address = all.Address
	o.Role = all.Role
	o.ShardId = all.ShardId
	o.Reachable = all.Reachable

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
