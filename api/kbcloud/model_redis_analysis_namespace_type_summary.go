// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type RedisAnalysisNamespaceTypeSummary struct {
	Type        *string  `json:"type,omitempty"`
	Keys        *int64   `json:"keys,omitempty"`
	Memory      *int64   `json:"memory,omitempty"`
	KeyRatio    *float64 `json:"keyRatio,omitempty"`
	MemoryRatio *float64 `json:"memoryRatio,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRedisAnalysisNamespaceTypeSummary instantiates a new RedisAnalysisNamespaceTypeSummary object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRedisAnalysisNamespaceTypeSummary() *RedisAnalysisNamespaceTypeSummary {
	this := RedisAnalysisNamespaceTypeSummary{}
	return &this
}

// NewRedisAnalysisNamespaceTypeSummaryWithDefaults instantiates a new RedisAnalysisNamespaceTypeSummary object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRedisAnalysisNamespaceTypeSummaryWithDefaults() *RedisAnalysisNamespaceTypeSummary {
	this := RedisAnalysisNamespaceTypeSummary{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RedisAnalysisNamespaceTypeSummary) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisNamespaceTypeSummary) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RedisAnalysisNamespaceTypeSummary) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RedisAnalysisNamespaceTypeSummary) SetType(v string) {
	o.Type = &v
}

// GetKeys returns the Keys field value if set, zero value otherwise.
func (o *RedisAnalysisNamespaceTypeSummary) GetKeys() int64 {
	if o == nil || o.Keys == nil {
		var ret int64
		return ret
	}
	return *o.Keys
}

// GetKeysOk returns a tuple with the Keys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisNamespaceTypeSummary) GetKeysOk() (*int64, bool) {
	if o == nil || o.Keys == nil {
		return nil, false
	}
	return o.Keys, true
}

// HasKeys returns a boolean if a field has been set.
func (o *RedisAnalysisNamespaceTypeSummary) HasKeys() bool {
	return o != nil && o.Keys != nil
}

// SetKeys gets a reference to the given int64 and assigns it to the Keys field.
func (o *RedisAnalysisNamespaceTypeSummary) SetKeys(v int64) {
	o.Keys = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *RedisAnalysisNamespaceTypeSummary) GetMemory() int64 {
	if o == nil || o.Memory == nil {
		var ret int64
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisNamespaceTypeSummary) GetMemoryOk() (*int64, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *RedisAnalysisNamespaceTypeSummary) HasMemory() bool {
	return o != nil && o.Memory != nil
}

// SetMemory gets a reference to the given int64 and assigns it to the Memory field.
func (o *RedisAnalysisNamespaceTypeSummary) SetMemory(v int64) {
	o.Memory = &v
}

// GetKeyRatio returns the KeyRatio field value if set, zero value otherwise.
func (o *RedisAnalysisNamespaceTypeSummary) GetKeyRatio() float64 {
	if o == nil || o.KeyRatio == nil {
		var ret float64
		return ret
	}
	return *o.KeyRatio
}

// GetKeyRatioOk returns a tuple with the KeyRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisNamespaceTypeSummary) GetKeyRatioOk() (*float64, bool) {
	if o == nil || o.KeyRatio == nil {
		return nil, false
	}
	return o.KeyRatio, true
}

// HasKeyRatio returns a boolean if a field has been set.
func (o *RedisAnalysisNamespaceTypeSummary) HasKeyRatio() bool {
	return o != nil && o.KeyRatio != nil
}

// SetKeyRatio gets a reference to the given float64 and assigns it to the KeyRatio field.
func (o *RedisAnalysisNamespaceTypeSummary) SetKeyRatio(v float64) {
	o.KeyRatio = &v
}

// GetMemoryRatio returns the MemoryRatio field value if set, zero value otherwise.
func (o *RedisAnalysisNamespaceTypeSummary) GetMemoryRatio() float64 {
	if o == nil || o.MemoryRatio == nil {
		var ret float64
		return ret
	}
	return *o.MemoryRatio
}

// GetMemoryRatioOk returns a tuple with the MemoryRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisNamespaceTypeSummary) GetMemoryRatioOk() (*float64, bool) {
	if o == nil || o.MemoryRatio == nil {
		return nil, false
	}
	return o.MemoryRatio, true
}

// HasMemoryRatio returns a boolean if a field has been set.
func (o *RedisAnalysisNamespaceTypeSummary) HasMemoryRatio() bool {
	return o != nil && o.MemoryRatio != nil
}

// SetMemoryRatio gets a reference to the given float64 and assigns it to the MemoryRatio field.
func (o *RedisAnalysisNamespaceTypeSummary) SetMemoryRatio(v float64) {
	o.MemoryRatio = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RedisAnalysisNamespaceTypeSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Keys != nil {
		toSerialize["keys"] = o.Keys
	}
	if o.Memory != nil {
		toSerialize["memory"] = o.Memory
	}
	if o.KeyRatio != nil {
		toSerialize["keyRatio"] = o.KeyRatio
	}
	if o.MemoryRatio != nil {
		toSerialize["memoryRatio"] = o.MemoryRatio
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RedisAnalysisNamespaceTypeSummary) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type        *string  `json:"type,omitempty"`
		Keys        *int64   `json:"keys,omitempty"`
		Memory      *int64   `json:"memory,omitempty"`
		KeyRatio    *float64 `json:"keyRatio,omitempty"`
		MemoryRatio *float64 `json:"memoryRatio,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"type", "keys", "memory", "keyRatio", "memoryRatio"})
	} else {
		return err
	}
	o.Type = all.Type
	o.Keys = all.Keys
	o.Memory = all.Memory
	o.KeyRatio = all.KeyRatio
	o.MemoryRatio = all.MemoryRatio

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
