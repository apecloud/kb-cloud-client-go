// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type TopicOffset struct {
	// 主题名称
	Topic *string `json:"topic,omitempty"`
	// 分区号
	Partition *int32 `json:"partition,omitempty"`
	// 消费者组当前消费位置
	ConsumerOffset common.NullableInt64 `json:"consumerOffset,omitempty"`
	// 分区最早消息的偏移量
	BeginningOffset common.NullableInt64 `json:"beginningOffset,omitempty"`
	// 分区最新消息的下一个偏移量
	EndOffset common.NullableInt64 `json:"endOffset,omitempty"`
	// 消费者组ID
	GroupId *string `json:"groupId,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTopicOffset instantiates a new TopicOffset object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTopicOffset() *TopicOffset {
	this := TopicOffset{}
	return &this
}

// NewTopicOffsetWithDefaults instantiates a new TopicOffset object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTopicOffsetWithDefaults() *TopicOffset {
	this := TopicOffset{}
	return &this
}

// GetTopic returns the Topic field value if set, zero value otherwise.
func (o *TopicOffset) GetTopic() string {
	if o == nil || o.Topic == nil {
		var ret string
		return ret
	}
	return *o.Topic
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopicOffset) GetTopicOk() (*string, bool) {
	if o == nil || o.Topic == nil {
		return nil, false
	}
	return o.Topic, true
}

// HasTopic returns a boolean if a field has been set.
func (o *TopicOffset) HasTopic() bool {
	return o != nil && o.Topic != nil
}

// SetTopic gets a reference to the given string and assigns it to the Topic field.
func (o *TopicOffset) SetTopic(v string) {
	o.Topic = &v
}

// GetPartition returns the Partition field value if set, zero value otherwise.
func (o *TopicOffset) GetPartition() int32 {
	if o == nil || o.Partition == nil {
		var ret int32
		return ret
	}
	return *o.Partition
}

// GetPartitionOk returns a tuple with the Partition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopicOffset) GetPartitionOk() (*int32, bool) {
	if o == nil || o.Partition == nil {
		return nil, false
	}
	return o.Partition, true
}

// HasPartition returns a boolean if a field has been set.
func (o *TopicOffset) HasPartition() bool {
	return o != nil && o.Partition != nil
}

// SetPartition gets a reference to the given int32 and assigns it to the Partition field.
func (o *TopicOffset) SetPartition(v int32) {
	o.Partition = &v
}

// GetConsumerOffset returns the ConsumerOffset field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TopicOffset) GetConsumerOffset() int64 {
	if o == nil || o.ConsumerOffset.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ConsumerOffset.Get()
}

// GetConsumerOffsetOk returns a tuple with the ConsumerOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TopicOffset) GetConsumerOffsetOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConsumerOffset.Get(), o.ConsumerOffset.IsSet()
}

// HasConsumerOffset returns a boolean if a field has been set.
func (o *TopicOffset) HasConsumerOffset() bool {
	return o != nil && o.ConsumerOffset.IsSet()
}

// SetConsumerOffset gets a reference to the given common.NullableInt64 and assigns it to the ConsumerOffset field.
func (o *TopicOffset) SetConsumerOffset(v int64) {
	o.ConsumerOffset.Set(&v)
}

// SetConsumerOffsetNil sets the value for ConsumerOffset to be an explicit nil.
func (o *TopicOffset) SetConsumerOffsetNil() {
	o.ConsumerOffset.Set(nil)
}

// UnsetConsumerOffset ensures that no value is present for ConsumerOffset, not even an explicit nil.
func (o *TopicOffset) UnsetConsumerOffset() {
	o.ConsumerOffset.Unset()
}

// GetBeginningOffset returns the BeginningOffset field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TopicOffset) GetBeginningOffset() int64 {
	if o == nil || o.BeginningOffset.Get() == nil {
		var ret int64
		return ret
	}
	return *o.BeginningOffset.Get()
}

// GetBeginningOffsetOk returns a tuple with the BeginningOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TopicOffset) GetBeginningOffsetOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.BeginningOffset.Get(), o.BeginningOffset.IsSet()
}

// HasBeginningOffset returns a boolean if a field has been set.
func (o *TopicOffset) HasBeginningOffset() bool {
	return o != nil && o.BeginningOffset.IsSet()
}

// SetBeginningOffset gets a reference to the given common.NullableInt64 and assigns it to the BeginningOffset field.
func (o *TopicOffset) SetBeginningOffset(v int64) {
	o.BeginningOffset.Set(&v)
}

// SetBeginningOffsetNil sets the value for BeginningOffset to be an explicit nil.
func (o *TopicOffset) SetBeginningOffsetNil() {
	o.BeginningOffset.Set(nil)
}

// UnsetBeginningOffset ensures that no value is present for BeginningOffset, not even an explicit nil.
func (o *TopicOffset) UnsetBeginningOffset() {
	o.BeginningOffset.Unset()
}

// GetEndOffset returns the EndOffset field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TopicOffset) GetEndOffset() int64 {
	if o == nil || o.EndOffset.Get() == nil {
		var ret int64
		return ret
	}
	return *o.EndOffset.Get()
}

// GetEndOffsetOk returns a tuple with the EndOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TopicOffset) GetEndOffsetOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndOffset.Get(), o.EndOffset.IsSet()
}

// HasEndOffset returns a boolean if a field has been set.
func (o *TopicOffset) HasEndOffset() bool {
	return o != nil && o.EndOffset.IsSet()
}

// SetEndOffset gets a reference to the given common.NullableInt64 and assigns it to the EndOffset field.
func (o *TopicOffset) SetEndOffset(v int64) {
	o.EndOffset.Set(&v)
}

// SetEndOffsetNil sets the value for EndOffset to be an explicit nil.
func (o *TopicOffset) SetEndOffsetNil() {
	o.EndOffset.Set(nil)
}

// UnsetEndOffset ensures that no value is present for EndOffset, not even an explicit nil.
func (o *TopicOffset) UnsetEndOffset() {
	o.EndOffset.Unset()
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *TopicOffset) GetGroupId() string {
	if o == nil || o.GroupId == nil {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopicOffset) GetGroupIdOk() (*string, bool) {
	if o == nil || o.GroupId == nil {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *TopicOffset) HasGroupId() bool {
	return o != nil && o.GroupId != nil
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *TopicOffset) SetGroupId(v string) {
	o.GroupId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TopicOffset) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Topic != nil {
		toSerialize["topic"] = o.Topic
	}
	if o.Partition != nil {
		toSerialize["partition"] = o.Partition
	}
	if o.ConsumerOffset.IsSet() {
		toSerialize["consumerOffset"] = o.ConsumerOffset.Get()
	}
	if o.BeginningOffset.IsSet() {
		toSerialize["beginningOffset"] = o.BeginningOffset.Get()
	}
	if o.EndOffset.IsSet() {
		toSerialize["endOffset"] = o.EndOffset.Get()
	}
	if o.GroupId != nil {
		toSerialize["groupId"] = o.GroupId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TopicOffset) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Topic           *string              `json:"topic,omitempty"`
		Partition       *int32               `json:"partition,omitempty"`
		ConsumerOffset  common.NullableInt64 `json:"consumerOffset,omitempty"`
		BeginningOffset common.NullableInt64 `json:"beginningOffset,omitempty"`
		EndOffset       common.NullableInt64 `json:"endOffset,omitempty"`
		GroupId         *string              `json:"groupId,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"topic", "partition", "consumerOffset", "beginningOffset", "endOffset", "groupId"})
	} else {
		return err
	}
	o.Topic = all.Topic
	o.Partition = all.Partition
	o.ConsumerOffset = all.ConsumerOffset
	o.BeginningOffset = all.BeginningOffset
	o.EndOffset = all.EndOffset
	o.GroupId = all.GroupId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
