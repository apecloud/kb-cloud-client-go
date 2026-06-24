// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type RedisKeySummary struct {
	Key  *string `json:"key,omitempty"`
	Type *string `json:"type,omitempty"`
	// key TTL in seconds; -1 means no expiration and -2 means missing key
	Ttl *int64 `json:"ttl,omitempty"`
	// memory usage in bytes when available
	Memory   *int64  `json:"memory,omitempty"`
	Length   *int64  `json:"length,omitempty"`
	Encoding *string `json:"encoding,omitempty"`
	Slot     *int64  `json:"slot,omitempty"`
	NodeId   *string `json:"nodeId,omitempty"`
	// datasource ACL access result for this key
	AclAccess *string `json:"aclAccess,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRedisKeySummary instantiates a new RedisKeySummary object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRedisKeySummary() *RedisKeySummary {
	this := RedisKeySummary{}
	return &this
}

// NewRedisKeySummaryWithDefaults instantiates a new RedisKeySummary object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRedisKeySummaryWithDefaults() *RedisKeySummary {
	this := RedisKeySummary{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *RedisKeySummary) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeySummary) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *RedisKeySummary) HasKey() bool {
	return o != nil && o.Key != nil
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *RedisKeySummary) SetKey(v string) {
	o.Key = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RedisKeySummary) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeySummary) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RedisKeySummary) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RedisKeySummary) SetType(v string) {
	o.Type = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *RedisKeySummary) GetTtl() int64 {
	if o == nil || o.Ttl == nil {
		var ret int64
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeySummary) GetTtlOk() (*int64, bool) {
	if o == nil || o.Ttl == nil {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *RedisKeySummary) HasTtl() bool {
	return o != nil && o.Ttl != nil
}

// SetTtl gets a reference to the given int64 and assigns it to the Ttl field.
func (o *RedisKeySummary) SetTtl(v int64) {
	o.Ttl = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *RedisKeySummary) GetMemory() int64 {
	if o == nil || o.Memory == nil {
		var ret int64
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeySummary) GetMemoryOk() (*int64, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *RedisKeySummary) HasMemory() bool {
	return o != nil && o.Memory != nil
}

// SetMemory gets a reference to the given int64 and assigns it to the Memory field.
func (o *RedisKeySummary) SetMemory(v int64) {
	o.Memory = &v
}

// GetLength returns the Length field value if set, zero value otherwise.
func (o *RedisKeySummary) GetLength() int64 {
	if o == nil || o.Length == nil {
		var ret int64
		return ret
	}
	return *o.Length
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeySummary) GetLengthOk() (*int64, bool) {
	if o == nil || o.Length == nil {
		return nil, false
	}
	return o.Length, true
}

// HasLength returns a boolean if a field has been set.
func (o *RedisKeySummary) HasLength() bool {
	return o != nil && o.Length != nil
}

// SetLength gets a reference to the given int64 and assigns it to the Length field.
func (o *RedisKeySummary) SetLength(v int64) {
	o.Length = &v
}

// GetEncoding returns the Encoding field value if set, zero value otherwise.
func (o *RedisKeySummary) GetEncoding() string {
	if o == nil || o.Encoding == nil {
		var ret string
		return ret
	}
	return *o.Encoding
}

// GetEncodingOk returns a tuple with the Encoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeySummary) GetEncodingOk() (*string, bool) {
	if o == nil || o.Encoding == nil {
		return nil, false
	}
	return o.Encoding, true
}

// HasEncoding returns a boolean if a field has been set.
func (o *RedisKeySummary) HasEncoding() bool {
	return o != nil && o.Encoding != nil
}

// SetEncoding gets a reference to the given string and assigns it to the Encoding field.
func (o *RedisKeySummary) SetEncoding(v string) {
	o.Encoding = &v
}

// GetSlot returns the Slot field value if set, zero value otherwise.
func (o *RedisKeySummary) GetSlot() int64 {
	if o == nil || o.Slot == nil {
		var ret int64
		return ret
	}
	return *o.Slot
}

// GetSlotOk returns a tuple with the Slot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeySummary) GetSlotOk() (*int64, bool) {
	if o == nil || o.Slot == nil {
		return nil, false
	}
	return o.Slot, true
}

// HasSlot returns a boolean if a field has been set.
func (o *RedisKeySummary) HasSlot() bool {
	return o != nil && o.Slot != nil
}

// SetSlot gets a reference to the given int64 and assigns it to the Slot field.
func (o *RedisKeySummary) SetSlot(v int64) {
	o.Slot = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *RedisKeySummary) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeySummary) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *RedisKeySummary) HasNodeId() bool {
	return o != nil && o.NodeId != nil
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *RedisKeySummary) SetNodeId(v string) {
	o.NodeId = &v
}

// GetAclAccess returns the AclAccess field value if set, zero value otherwise.
func (o *RedisKeySummary) GetAclAccess() string {
	if o == nil || o.AclAccess == nil {
		var ret string
		return ret
	}
	return *o.AclAccess
}

// GetAclAccessOk returns a tuple with the AclAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeySummary) GetAclAccessOk() (*string, bool) {
	if o == nil || o.AclAccess == nil {
		return nil, false
	}
	return o.AclAccess, true
}

// HasAclAccess returns a boolean if a field has been set.
func (o *RedisKeySummary) HasAclAccess() bool {
	return o != nil && o.AclAccess != nil
}

// SetAclAccess gets a reference to the given string and assigns it to the AclAccess field.
func (o *RedisKeySummary) SetAclAccess(v string) {
	o.AclAccess = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RedisKeySummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Ttl != nil {
		toSerialize["ttl"] = o.Ttl
	}
	if o.Memory != nil {
		toSerialize["memory"] = o.Memory
	}
	if o.Length != nil {
		toSerialize["length"] = o.Length
	}
	if o.Encoding != nil {
		toSerialize["encoding"] = o.Encoding
	}
	if o.Slot != nil {
		toSerialize["slot"] = o.Slot
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.AclAccess != nil {
		toSerialize["aclAccess"] = o.AclAccess
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RedisKeySummary) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Key       *string `json:"key,omitempty"`
		Type      *string `json:"type,omitempty"`
		Ttl       *int64  `json:"ttl,omitempty"`
		Memory    *int64  `json:"memory,omitempty"`
		Length    *int64  `json:"length,omitempty"`
		Encoding  *string `json:"encoding,omitempty"`
		Slot      *int64  `json:"slot,omitempty"`
		NodeId    *string `json:"nodeId,omitempty"`
		AclAccess *string `json:"aclAccess,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"key", "type", "ttl", "memory", "length", "encoding", "slot", "nodeId", "aclAccess"})
	} else {
		return err
	}
	o.Key = all.Key
	o.Type = all.Type
	o.Ttl = all.Ttl
	o.Memory = all.Memory
	o.Length = all.Length
	o.Encoding = all.Encoding
	o.Slot = all.Slot
	o.NodeId = all.NodeId
	o.AclAccess = all.AclAccess

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
