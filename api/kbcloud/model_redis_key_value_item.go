// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// RedisKeyValueItem Stable Redis collection item shape. Fields are populated according to Redis value type.
type RedisKeyValueItem struct {
	// Hash field name.
	Field *string `json:"field,omitempty"`
	// Hash/list/set/json item value.
	Value map[string]interface{} `json:"value,omitempty"`
	// Set or sorted-set member.
	Member *string `json:"member,omitempty"`
	// Sorted-set score.
	Score *float64 `json:"score,omitempty"`
	// Stream entry ID.
	Id *string `json:"id,omitempty"`
	// Stream entry field/value map.
	Fields map[string]string `json:"fields,omitempty"`
	// List item index when available.
	Index *int64 `json:"index,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRedisKeyValueItem instantiates a new RedisKeyValueItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRedisKeyValueItem() *RedisKeyValueItem {
	this := RedisKeyValueItem{}
	return &this
}

// NewRedisKeyValueItemWithDefaults instantiates a new RedisKeyValueItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRedisKeyValueItemWithDefaults() *RedisKeyValueItem {
	this := RedisKeyValueItem{}
	return &this
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *RedisKeyValueItem) GetField() string {
	if o == nil || o.Field == nil {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyValueItem) GetFieldOk() (*string, bool) {
	if o == nil || o.Field == nil {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *RedisKeyValueItem) HasField() bool {
	return o != nil && o.Field != nil
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *RedisKeyValueItem) SetField(v string) {
	o.Field = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *RedisKeyValueItem) GetValue() map[string]interface{} {
	if o == nil || o.Value == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyValueItem) GetValueOk() (*map[string]interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RedisKeyValueItem) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given map[string]interface{} and assigns it to the Value field.
func (o *RedisKeyValueItem) SetValue(v map[string]interface{}) {
	o.Value = v
}

// GetMember returns the Member field value if set, zero value otherwise.
func (o *RedisKeyValueItem) GetMember() string {
	if o == nil || o.Member == nil {
		var ret string
		return ret
	}
	return *o.Member
}

// GetMemberOk returns a tuple with the Member field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyValueItem) GetMemberOk() (*string, bool) {
	if o == nil || o.Member == nil {
		return nil, false
	}
	return o.Member, true
}

// HasMember returns a boolean if a field has been set.
func (o *RedisKeyValueItem) HasMember() bool {
	return o != nil && o.Member != nil
}

// SetMember gets a reference to the given string and assigns it to the Member field.
func (o *RedisKeyValueItem) SetMember(v string) {
	o.Member = &v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *RedisKeyValueItem) GetScore() float64 {
	if o == nil || o.Score == nil {
		var ret float64
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyValueItem) GetScoreOk() (*float64, bool) {
	if o == nil || o.Score == nil {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *RedisKeyValueItem) HasScore() bool {
	return o != nil && o.Score != nil
}

// SetScore gets a reference to the given float64 and assigns it to the Score field.
func (o *RedisKeyValueItem) SetScore(v float64) {
	o.Score = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RedisKeyValueItem) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyValueItem) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RedisKeyValueItem) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RedisKeyValueItem) SetId(v string) {
	o.Id = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *RedisKeyValueItem) GetFields() map[string]string {
	if o == nil || o.Fields == nil {
		var ret map[string]string
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyValueItem) GetFieldsOk() (*map[string]string, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return &o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *RedisKeyValueItem) HasFields() bool {
	return o != nil && o.Fields != nil
}

// SetFields gets a reference to the given map[string]string and assigns it to the Fields field.
func (o *RedisKeyValueItem) SetFields(v map[string]string) {
	o.Fields = v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *RedisKeyValueItem) GetIndex() int64 {
	if o == nil || o.Index == nil {
		var ret int64
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyValueItem) GetIndexOk() (*int64, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *RedisKeyValueItem) HasIndex() bool {
	return o != nil && o.Index != nil
}

// SetIndex gets a reference to the given int64 and assigns it to the Index field.
func (o *RedisKeyValueItem) SetIndex(v int64) {
	o.Index = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RedisKeyValueItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Field != nil {
		toSerialize["field"] = o.Field
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Member != nil {
		toSerialize["member"] = o.Member
	}
	if o.Score != nil {
		toSerialize["score"] = o.Score
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	if o.Index != nil {
		toSerialize["index"] = o.Index
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RedisKeyValueItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Field  *string                `json:"field,omitempty"`
		Value  map[string]interface{} `json:"value,omitempty"`
		Member *string                `json:"member,omitempty"`
		Score  *float64               `json:"score,omitempty"`
		Id     *string                `json:"id,omitempty"`
		Fields map[string]string      `json:"fields,omitempty"`
		Index  *int64                 `json:"index,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"field", "value", "member", "score", "id", "fields", "index"})
	} else {
		return err
	}
	o.Field = all.Field
	o.Value = all.Value
	o.Member = all.Member
	o.Score = all.Score
	o.Id = all.Id
	o.Fields = all.Fields
	o.Index = all.Index

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
