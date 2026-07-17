// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type RedisAnalysisKey struct {
	Name     *string `json:"name,omitempty"`
	Type     *string `json:"type,omitempty"`
	Nsp      *string `json:"nsp,omitempty"`
	Ttl      *int64  `json:"ttl,omitempty"`
	Memory   *int64  `json:"memory,omitempty"`
	Length   *int64  `json:"length,omitempty"`
	Encoding *string `json:"encoding,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRedisAnalysisKey instantiates a new RedisAnalysisKey object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRedisAnalysisKey() *RedisAnalysisKey {
	this := RedisAnalysisKey{}
	return &this
}

// NewRedisAnalysisKeyWithDefaults instantiates a new RedisAnalysisKey object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRedisAnalysisKeyWithDefaults() *RedisAnalysisKey {
	this := RedisAnalysisKey{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RedisAnalysisKey) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisKey) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RedisAnalysisKey) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RedisAnalysisKey) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RedisAnalysisKey) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisKey) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RedisAnalysisKey) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RedisAnalysisKey) SetType(v string) {
	o.Type = &v
}

// GetNsp returns the Nsp field value if set, zero value otherwise.
func (o *RedisAnalysisKey) GetNsp() string {
	if o == nil || o.Nsp == nil {
		var ret string
		return ret
	}
	return *o.Nsp
}

// GetNspOk returns a tuple with the Nsp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisKey) GetNspOk() (*string, bool) {
	if o == nil || o.Nsp == nil {
		return nil, false
	}
	return o.Nsp, true
}

// HasNsp returns a boolean if a field has been set.
func (o *RedisAnalysisKey) HasNsp() bool {
	return o != nil && o.Nsp != nil
}

// SetNsp gets a reference to the given string and assigns it to the Nsp field.
func (o *RedisAnalysisKey) SetNsp(v string) {
	o.Nsp = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *RedisAnalysisKey) GetTtl() int64 {
	if o == nil || o.Ttl == nil {
		var ret int64
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisKey) GetTtlOk() (*int64, bool) {
	if o == nil || o.Ttl == nil {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *RedisAnalysisKey) HasTtl() bool {
	return o != nil && o.Ttl != nil
}

// SetTtl gets a reference to the given int64 and assigns it to the Ttl field.
func (o *RedisAnalysisKey) SetTtl(v int64) {
	o.Ttl = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *RedisAnalysisKey) GetMemory() int64 {
	if o == nil || o.Memory == nil {
		var ret int64
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisKey) GetMemoryOk() (*int64, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *RedisAnalysisKey) HasMemory() bool {
	return o != nil && o.Memory != nil
}

// SetMemory gets a reference to the given int64 and assigns it to the Memory field.
func (o *RedisAnalysisKey) SetMemory(v int64) {
	o.Memory = &v
}

// GetLength returns the Length field value if set, zero value otherwise.
func (o *RedisAnalysisKey) GetLength() int64 {
	if o == nil || o.Length == nil {
		var ret int64
		return ret
	}
	return *o.Length
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisKey) GetLengthOk() (*int64, bool) {
	if o == nil || o.Length == nil {
		return nil, false
	}
	return o.Length, true
}

// HasLength returns a boolean if a field has been set.
func (o *RedisAnalysisKey) HasLength() bool {
	return o != nil && o.Length != nil
}

// SetLength gets a reference to the given int64 and assigns it to the Length field.
func (o *RedisAnalysisKey) SetLength(v int64) {
	o.Length = &v
}

// GetEncoding returns the Encoding field value if set, zero value otherwise.
func (o *RedisAnalysisKey) GetEncoding() string {
	if o == nil || o.Encoding == nil {
		var ret string
		return ret
	}
	return *o.Encoding
}

// GetEncodingOk returns a tuple with the Encoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisKey) GetEncodingOk() (*string, bool) {
	if o == nil || o.Encoding == nil {
		return nil, false
	}
	return o.Encoding, true
}

// HasEncoding returns a boolean if a field has been set.
func (o *RedisAnalysisKey) HasEncoding() bool {
	return o != nil && o.Encoding != nil
}

// SetEncoding gets a reference to the given string and assigns it to the Encoding field.
func (o *RedisAnalysisKey) SetEncoding(v string) {
	o.Encoding = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RedisAnalysisKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Nsp != nil {
		toSerialize["nsp"] = o.Nsp
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RedisAnalysisKey) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name     *string `json:"name,omitempty"`
		Type     *string `json:"type,omitempty"`
		Nsp      *string `json:"nsp,omitempty"`
		Ttl      *int64  `json:"ttl,omitempty"`
		Memory   *int64  `json:"memory,omitempty"`
		Length   *int64  `json:"length,omitempty"`
		Encoding *string `json:"encoding,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "type", "nsp", "ttl", "memory", "length", "encoding"})
	} else {
		return err
	}
	o.Name = all.Name
	o.Type = all.Type
	o.Nsp = all.Nsp
	o.Ttl = all.Ttl
	o.Memory = all.Memory
	o.Length = all.Length
	o.Encoding = all.Encoding

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
