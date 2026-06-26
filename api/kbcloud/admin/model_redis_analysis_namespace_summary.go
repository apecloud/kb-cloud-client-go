// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type RedisAnalysisNamespaceSummary struct {
	Nsp         *string                             `json:"nsp,omitempty"`
	Keys        *int64                              `json:"keys,omitempty"`
	Memory      *int64                              `json:"memory,omitempty"`
	KeyRatio    *float64                            `json:"keyRatio,omitempty"`
	MemoryRatio *float64                            `json:"memoryRatio,omitempty"`
	Types       []RedisAnalysisNamespaceTypeSummary `json:"types,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRedisAnalysisNamespaceSummary instantiates a new RedisAnalysisNamespaceSummary object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRedisAnalysisNamespaceSummary() *RedisAnalysisNamespaceSummary {
	this := RedisAnalysisNamespaceSummary{}
	return &this
}

// NewRedisAnalysisNamespaceSummaryWithDefaults instantiates a new RedisAnalysisNamespaceSummary object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRedisAnalysisNamespaceSummaryWithDefaults() *RedisAnalysisNamespaceSummary {
	this := RedisAnalysisNamespaceSummary{}
	return &this
}

// GetNsp returns the Nsp field value if set, zero value otherwise.
func (o *RedisAnalysisNamespaceSummary) GetNsp() string {
	if o == nil || o.Nsp == nil {
		var ret string
		return ret
	}
	return *o.Nsp
}

// GetNspOk returns a tuple with the Nsp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisNamespaceSummary) GetNspOk() (*string, bool) {
	if o == nil || o.Nsp == nil {
		return nil, false
	}
	return o.Nsp, true
}

// HasNsp returns a boolean if a field has been set.
func (o *RedisAnalysisNamespaceSummary) HasNsp() bool {
	return o != nil && o.Nsp != nil
}

// SetNsp gets a reference to the given string and assigns it to the Nsp field.
func (o *RedisAnalysisNamespaceSummary) SetNsp(v string) {
	o.Nsp = &v
}

// GetKeys returns the Keys field value if set, zero value otherwise.
func (o *RedisAnalysisNamespaceSummary) GetKeys() int64 {
	if o == nil || o.Keys == nil {
		var ret int64
		return ret
	}
	return *o.Keys
}

// GetKeysOk returns a tuple with the Keys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisNamespaceSummary) GetKeysOk() (*int64, bool) {
	if o == nil || o.Keys == nil {
		return nil, false
	}
	return o.Keys, true
}

// HasKeys returns a boolean if a field has been set.
func (o *RedisAnalysisNamespaceSummary) HasKeys() bool {
	return o != nil && o.Keys != nil
}

// SetKeys gets a reference to the given int64 and assigns it to the Keys field.
func (o *RedisAnalysisNamespaceSummary) SetKeys(v int64) {
	o.Keys = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *RedisAnalysisNamespaceSummary) GetMemory() int64 {
	if o == nil || o.Memory == nil {
		var ret int64
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisNamespaceSummary) GetMemoryOk() (*int64, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *RedisAnalysisNamespaceSummary) HasMemory() bool {
	return o != nil && o.Memory != nil
}

// SetMemory gets a reference to the given int64 and assigns it to the Memory field.
func (o *RedisAnalysisNamespaceSummary) SetMemory(v int64) {
	o.Memory = &v
}

// GetKeyRatio returns the KeyRatio field value if set, zero value otherwise.
func (o *RedisAnalysisNamespaceSummary) GetKeyRatio() float64 {
	if o == nil || o.KeyRatio == nil {
		var ret float64
		return ret
	}
	return *o.KeyRatio
}

// GetKeyRatioOk returns a tuple with the KeyRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisNamespaceSummary) GetKeyRatioOk() (*float64, bool) {
	if o == nil || o.KeyRatio == nil {
		return nil, false
	}
	return o.KeyRatio, true
}

// HasKeyRatio returns a boolean if a field has been set.
func (o *RedisAnalysisNamespaceSummary) HasKeyRatio() bool {
	return o != nil && o.KeyRatio != nil
}

// SetKeyRatio gets a reference to the given float64 and assigns it to the KeyRatio field.
func (o *RedisAnalysisNamespaceSummary) SetKeyRatio(v float64) {
	o.KeyRatio = &v
}

// GetMemoryRatio returns the MemoryRatio field value if set, zero value otherwise.
func (o *RedisAnalysisNamespaceSummary) GetMemoryRatio() float64 {
	if o == nil || o.MemoryRatio == nil {
		var ret float64
		return ret
	}
	return *o.MemoryRatio
}

// GetMemoryRatioOk returns a tuple with the MemoryRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisNamespaceSummary) GetMemoryRatioOk() (*float64, bool) {
	if o == nil || o.MemoryRatio == nil {
		return nil, false
	}
	return o.MemoryRatio, true
}

// HasMemoryRatio returns a boolean if a field has been set.
func (o *RedisAnalysisNamespaceSummary) HasMemoryRatio() bool {
	return o != nil && o.MemoryRatio != nil
}

// SetMemoryRatio gets a reference to the given float64 and assigns it to the MemoryRatio field.
func (o *RedisAnalysisNamespaceSummary) SetMemoryRatio(v float64) {
	o.MemoryRatio = &v
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *RedisAnalysisNamespaceSummary) GetTypes() []RedisAnalysisNamespaceTypeSummary {
	if o == nil || o.Types == nil {
		var ret []RedisAnalysisNamespaceTypeSummary
		return ret
	}
	return o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisNamespaceSummary) GetTypesOk() (*[]RedisAnalysisNamespaceTypeSummary, bool) {
	if o == nil || o.Types == nil {
		return nil, false
	}
	return &o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *RedisAnalysisNamespaceSummary) HasTypes() bool {
	return o != nil && o.Types != nil
}

// SetTypes gets a reference to the given []RedisAnalysisNamespaceTypeSummary and assigns it to the Types field.
func (o *RedisAnalysisNamespaceSummary) SetTypes(v []RedisAnalysisNamespaceTypeSummary) {
	o.Types = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RedisAnalysisNamespaceSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Nsp != nil {
		toSerialize["nsp"] = o.Nsp
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
	if o.Types != nil {
		toSerialize["types"] = o.Types
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RedisAnalysisNamespaceSummary) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Nsp         *string                             `json:"nsp,omitempty"`
		Keys        *int64                              `json:"keys,omitempty"`
		Memory      *int64                              `json:"memory,omitempty"`
		KeyRatio    *float64                            `json:"keyRatio,omitempty"`
		MemoryRatio *float64                            `json:"memoryRatio,omitempty"`
		Types       []RedisAnalysisNamespaceTypeSummary `json:"types,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"nsp", "keys", "memory", "keyRatio", "memoryRatio", "types"})
	} else {
		return err
	}
	o.Nsp = all.Nsp
	o.Keys = all.Keys
	o.Memory = all.Memory
	o.KeyRatio = all.KeyRatio
	o.MemoryRatio = all.MemoryRatio
	o.Types = all.Types

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
