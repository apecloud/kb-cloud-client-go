// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type RedisKeyDetail struct {
	Metadata *RedisKeySummary `json:"metadata,omitempty"`
	// Discriminated Redis key value payload for string, hash, list, set, zset, stream, and json editors.
	Value      *RedisKeyValue `json:"value,omitempty"`
	NextCursor *string        `json:"nextCursor,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRedisKeyDetail instantiates a new RedisKeyDetail object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRedisKeyDetail() *RedisKeyDetail {
	this := RedisKeyDetail{}
	return &this
}

// NewRedisKeyDetailWithDefaults instantiates a new RedisKeyDetail object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRedisKeyDetailWithDefaults() *RedisKeyDetail {
	this := RedisKeyDetail{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *RedisKeyDetail) GetMetadata() RedisKeySummary {
	if o == nil || o.Metadata == nil {
		var ret RedisKeySummary
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyDetail) GetMetadataOk() (*RedisKeySummary, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *RedisKeyDetail) HasMetadata() bool {
	return o != nil && o.Metadata != nil
}

// SetMetadata gets a reference to the given RedisKeySummary and assigns it to the Metadata field.
func (o *RedisKeyDetail) SetMetadata(v RedisKeySummary) {
	o.Metadata = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *RedisKeyDetail) GetValue() RedisKeyValue {
	if o == nil || o.Value == nil {
		var ret RedisKeyValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyDetail) GetValueOk() (*RedisKeyValue, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RedisKeyDetail) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given RedisKeyValue and assigns it to the Value field.
func (o *RedisKeyDetail) SetValue(v RedisKeyValue) {
	o.Value = &v
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise.
func (o *RedisKeyDetail) GetNextCursor() string {
	if o == nil || o.NextCursor == nil {
		var ret string
		return ret
	}
	return *o.NextCursor
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyDetail) GetNextCursorOk() (*string, bool) {
	if o == nil || o.NextCursor == nil {
		return nil, false
	}
	return o.NextCursor, true
}

// HasNextCursor returns a boolean if a field has been set.
func (o *RedisKeyDetail) HasNextCursor() bool {
	return o != nil && o.NextCursor != nil
}

// SetNextCursor gets a reference to the given string and assigns it to the NextCursor field.
func (o *RedisKeyDetail) SetNextCursor(v string) {
	o.NextCursor = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RedisKeyDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.NextCursor != nil {
		toSerialize["nextCursor"] = o.NextCursor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RedisKeyDetail) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Metadata   *RedisKeySummary `json:"metadata,omitempty"`
		Value      *RedisKeyValue   `json:"value,omitempty"`
		NextCursor *string          `json:"nextCursor,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"metadata", "value", "nextCursor"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Metadata != nil && all.Metadata.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Metadata = all.Metadata
	if all.Value != nil && all.Value.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Value = all.Value
	o.NextCursor = all.NextCursor

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
