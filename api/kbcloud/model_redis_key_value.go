// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// RedisKeyValue Discriminated Redis key value payload for string, hash, list, set, zset, stream, and json editors.
type RedisKeyValue struct {
	// Redis value type, such as string, hash, list, set, zset, stream, or json.
	Type *string `json:"type,omitempty"`
	// String or serialized scalar value for string/json previews.
	Text *string `json:"text,omitempty"`
	// Paged collection items for hash, list, set, zset, stream, and json collection previews.
	Items []RedisKeyValueItem `json:"items,omitempty"`
	// Cursor for the next value page. Empty or "0" means the value page is complete.
	Cursor *string `json:"cursor,omitempty"`
	// Offset used by offset-based value pagination.
	Offset *int64 `json:"offset,omitempty"`
	// Maximum number of value items returned.
	Limit *int64 `json:"limit,omitempty"`
	// Whether the returned value is truncated by size or item limit.
	Truncated *bool `json:"truncated,omitempty"`
	// JSON path used for the returned JSON value.
	JsonPath *string `json:"jsonPath,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRedisKeyValue instantiates a new RedisKeyValue object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRedisKeyValue() *RedisKeyValue {
	this := RedisKeyValue{}
	return &this
}

// NewRedisKeyValueWithDefaults instantiates a new RedisKeyValue object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRedisKeyValueWithDefaults() *RedisKeyValue {
	this := RedisKeyValue{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RedisKeyValue) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyValue) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RedisKeyValue) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RedisKeyValue) SetType(v string) {
	o.Type = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *RedisKeyValue) GetText() string {
	if o == nil || o.Text == nil {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyValue) GetTextOk() (*string, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *RedisKeyValue) HasText() bool {
	return o != nil && o.Text != nil
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *RedisKeyValue) SetText(v string) {
	o.Text = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *RedisKeyValue) GetItems() []RedisKeyValueItem {
	if o == nil || o.Items == nil {
		var ret []RedisKeyValueItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyValue) GetItemsOk() (*[]RedisKeyValueItem, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return &o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *RedisKeyValue) HasItems() bool {
	return o != nil && o.Items != nil
}

// SetItems gets a reference to the given []RedisKeyValueItem and assigns it to the Items field.
func (o *RedisKeyValue) SetItems(v []RedisKeyValueItem) {
	o.Items = v
}

// GetCursor returns the Cursor field value if set, zero value otherwise.
func (o *RedisKeyValue) GetCursor() string {
	if o == nil || o.Cursor == nil {
		var ret string
		return ret
	}
	return *o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyValue) GetCursorOk() (*string, bool) {
	if o == nil || o.Cursor == nil {
		return nil, false
	}
	return o.Cursor, true
}

// HasCursor returns a boolean if a field has been set.
func (o *RedisKeyValue) HasCursor() bool {
	return o != nil && o.Cursor != nil
}

// SetCursor gets a reference to the given string and assigns it to the Cursor field.
func (o *RedisKeyValue) SetCursor(v string) {
	o.Cursor = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *RedisKeyValue) GetOffset() int64 {
	if o == nil || o.Offset == nil {
		var ret int64
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyValue) GetOffsetOk() (*int64, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *RedisKeyValue) HasOffset() bool {
	return o != nil && o.Offset != nil
}

// SetOffset gets a reference to the given int64 and assigns it to the Offset field.
func (o *RedisKeyValue) SetOffset(v int64) {
	o.Offset = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *RedisKeyValue) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyValue) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *RedisKeyValue) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *RedisKeyValue) SetLimit(v int64) {
	o.Limit = &v
}

// GetTruncated returns the Truncated field value if set, zero value otherwise.
func (o *RedisKeyValue) GetTruncated() bool {
	if o == nil || o.Truncated == nil {
		var ret bool
		return ret
	}
	return *o.Truncated
}

// GetTruncatedOk returns a tuple with the Truncated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyValue) GetTruncatedOk() (*bool, bool) {
	if o == nil || o.Truncated == nil {
		return nil, false
	}
	return o.Truncated, true
}

// HasTruncated returns a boolean if a field has been set.
func (o *RedisKeyValue) HasTruncated() bool {
	return o != nil && o.Truncated != nil
}

// SetTruncated gets a reference to the given bool and assigns it to the Truncated field.
func (o *RedisKeyValue) SetTruncated(v bool) {
	o.Truncated = &v
}

// GetJsonPath returns the JsonPath field value if set, zero value otherwise.
func (o *RedisKeyValue) GetJsonPath() string {
	if o == nil || o.JsonPath == nil {
		var ret string
		return ret
	}
	return *o.JsonPath
}

// GetJsonPathOk returns a tuple with the JsonPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyValue) GetJsonPathOk() (*string, bool) {
	if o == nil || o.JsonPath == nil {
		return nil, false
	}
	return o.JsonPath, true
}

// HasJsonPath returns a boolean if a field has been set.
func (o *RedisKeyValue) HasJsonPath() bool {
	return o != nil && o.JsonPath != nil
}

// SetJsonPath gets a reference to the given string and assigns it to the JsonPath field.
func (o *RedisKeyValue) SetJsonPath(v string) {
	o.JsonPath = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RedisKeyValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Text != nil {
		toSerialize["text"] = o.Text
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.Cursor != nil {
		toSerialize["cursor"] = o.Cursor
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Truncated != nil {
		toSerialize["truncated"] = o.Truncated
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
func (o *RedisKeyValue) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type      *string             `json:"type,omitempty"`
		Text      *string             `json:"text,omitempty"`
		Items     []RedisKeyValueItem `json:"items,omitempty"`
		Cursor    *string             `json:"cursor,omitempty"`
		Offset    *int64              `json:"offset,omitempty"`
		Limit     *int64              `json:"limit,omitempty"`
		Truncated *bool               `json:"truncated,omitempty"`
		JsonPath  *string             `json:"jsonPath,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"type", "text", "items", "cursor", "offset", "limit", "truncated", "jsonPath"})
	} else {
		return err
	}
	o.Type = all.Type
	o.Text = all.Text
	o.Items = all.Items
	o.Cursor = all.Cursor
	o.Offset = all.Offset
	o.Limit = all.Limit
	o.Truncated = all.Truncated
	o.JsonPath = all.JsonPath

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
