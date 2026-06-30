// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type RedisKeyDetailRequest struct {
	DbIndex     *int64  `json:"dbIndex,omitempty"`
	Key         string  `json:"key"`
	ValueCursor *string `json:"valueCursor,omitempty"`
	Offset      *int64  `json:"offset,omitempty"`
	Limit       *int64  `json:"limit,omitempty"`
	JsonPath    *string `json:"jsonPath,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRedisKeyDetailRequest instantiates a new RedisKeyDetailRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRedisKeyDetailRequest(key string) *RedisKeyDetailRequest {
	this := RedisKeyDetailRequest{}
	this.Key = key
	return &this
}

// NewRedisKeyDetailRequestWithDefaults instantiates a new RedisKeyDetailRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRedisKeyDetailRequestWithDefaults() *RedisKeyDetailRequest {
	this := RedisKeyDetailRequest{}
	return &this
}

// GetDbIndex returns the DbIndex field value if set, zero value otherwise.
func (o *RedisKeyDetailRequest) GetDbIndex() int64 {
	if o == nil || o.DbIndex == nil {
		var ret int64
		return ret
	}
	return *o.DbIndex
}

// GetDbIndexOk returns a tuple with the DbIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyDetailRequest) GetDbIndexOk() (*int64, bool) {
	if o == nil || o.DbIndex == nil {
		return nil, false
	}
	return o.DbIndex, true
}

// HasDbIndex returns a boolean if a field has been set.
func (o *RedisKeyDetailRequest) HasDbIndex() bool {
	return o != nil && o.DbIndex != nil
}

// SetDbIndex gets a reference to the given int64 and assigns it to the DbIndex field.
func (o *RedisKeyDetailRequest) SetDbIndex(v int64) {
	o.DbIndex = &v
}

// GetKey returns the Key field value.
func (o *RedisKeyDetailRequest) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *RedisKeyDetailRequest) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value.
func (o *RedisKeyDetailRequest) SetKey(v string) {
	o.Key = v
}

// GetValueCursor returns the ValueCursor field value if set, zero value otherwise.
func (o *RedisKeyDetailRequest) GetValueCursor() string {
	if o == nil || o.ValueCursor == nil {
		var ret string
		return ret
	}
	return *o.ValueCursor
}

// GetValueCursorOk returns a tuple with the ValueCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyDetailRequest) GetValueCursorOk() (*string, bool) {
	if o == nil || o.ValueCursor == nil {
		return nil, false
	}
	return o.ValueCursor, true
}

// HasValueCursor returns a boolean if a field has been set.
func (o *RedisKeyDetailRequest) HasValueCursor() bool {
	return o != nil && o.ValueCursor != nil
}

// SetValueCursor gets a reference to the given string and assigns it to the ValueCursor field.
func (o *RedisKeyDetailRequest) SetValueCursor(v string) {
	o.ValueCursor = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *RedisKeyDetailRequest) GetOffset() int64 {
	if o == nil || o.Offset == nil {
		var ret int64
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyDetailRequest) GetOffsetOk() (*int64, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *RedisKeyDetailRequest) HasOffset() bool {
	return o != nil && o.Offset != nil
}

// SetOffset gets a reference to the given int64 and assigns it to the Offset field.
func (o *RedisKeyDetailRequest) SetOffset(v int64) {
	o.Offset = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *RedisKeyDetailRequest) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyDetailRequest) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *RedisKeyDetailRequest) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *RedisKeyDetailRequest) SetLimit(v int64) {
	o.Limit = &v
}

// GetJsonPath returns the JsonPath field value if set, zero value otherwise.
func (o *RedisKeyDetailRequest) GetJsonPath() string {
	if o == nil || o.JsonPath == nil {
		var ret string
		return ret
	}
	return *o.JsonPath
}

// GetJsonPathOk returns a tuple with the JsonPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyDetailRequest) GetJsonPathOk() (*string, bool) {
	if o == nil || o.JsonPath == nil {
		return nil, false
	}
	return o.JsonPath, true
}

// HasJsonPath returns a boolean if a field has been set.
func (o *RedisKeyDetailRequest) HasJsonPath() bool {
	return o != nil && o.JsonPath != nil
}

// SetJsonPath gets a reference to the given string and assigns it to the JsonPath field.
func (o *RedisKeyDetailRequest) SetJsonPath(v string) {
	o.JsonPath = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RedisKeyDetailRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.DbIndex != nil {
		toSerialize["dbIndex"] = o.DbIndex
	}
	toSerialize["key"] = o.Key
	if o.ValueCursor != nil {
		toSerialize["valueCursor"] = o.ValueCursor
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.JsonPath != nil {
		toSerialize["jsonPath"] = o.JsonPath
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RedisKeyDetailRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DbIndex     *int64  `json:"dbIndex,omitempty"`
		Key         *string `json:"key"`
		ValueCursor *string `json:"valueCursor,omitempty"`
		Offset      *int64  `json:"offset,omitempty"`
		Limit       *int64  `json:"limit,omitempty"`
		JsonPath    *string `json:"jsonPath,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Key == nil {
		return fmt.Errorf("required field key missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"dbIndex", "key", "valueCursor", "offset", "limit", "jsonPath"})
	} else {
		return err
	}
	o.DbIndex = all.DbIndex
	o.Key = *all.Key
	o.ValueCursor = all.ValueCursor
	o.Offset = all.Offset
	o.Limit = all.Limit
	o.JsonPath = all.JsonPath

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
