// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package data

import "github.com/apecloud/kb-cloud-client-go/api/common"

type TopicMessage struct {
	// 消息所属的主题
	Topic *string `json:"topic,omitempty"`
	// 消息所在的分区号
	Partition *int32 `json:"partition,omitempty"`
	// 消息在分区中的偏移量
	Offset *int64 `json:"offset,omitempty"`
	// 消息的时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`
	// 消息的key
	Key *string `json:"key,omitempty"`
	// 消息的内容
	Value *string `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTopicMessage instantiates a new TopicMessage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTopicMessage() *TopicMessage {
	this := TopicMessage{}
	return &this
}

// NewTopicMessageWithDefaults instantiates a new TopicMessage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTopicMessageWithDefaults() *TopicMessage {
	this := TopicMessage{}
	return &this
}

// GetTopic returns the Topic field value if set, zero value otherwise.
func (o *TopicMessage) GetTopic() string {
	if o == nil || o.Topic == nil {
		var ret string
		return ret
	}
	return *o.Topic
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopicMessage) GetTopicOk() (*string, bool) {
	if o == nil || o.Topic == nil {
		return nil, false
	}
	return o.Topic, true
}

// HasTopic returns a boolean if a field has been set.
func (o *TopicMessage) HasTopic() bool {
	return o != nil && o.Topic != nil
}

// SetTopic gets a reference to the given string and assigns it to the Topic field.
func (o *TopicMessage) SetTopic(v string) {
	o.Topic = &v
}

// GetPartition returns the Partition field value if set, zero value otherwise.
func (o *TopicMessage) GetPartition() int32 {
	if o == nil || o.Partition == nil {
		var ret int32
		return ret
	}
	return *o.Partition
}

// GetPartitionOk returns a tuple with the Partition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopicMessage) GetPartitionOk() (*int32, bool) {
	if o == nil || o.Partition == nil {
		return nil, false
	}
	return o.Partition, true
}

// HasPartition returns a boolean if a field has been set.
func (o *TopicMessage) HasPartition() bool {
	return o != nil && o.Partition != nil
}

// SetPartition gets a reference to the given int32 and assigns it to the Partition field.
func (o *TopicMessage) SetPartition(v int32) {
	o.Partition = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *TopicMessage) GetOffset() int64 {
	if o == nil || o.Offset == nil {
		var ret int64
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopicMessage) GetOffsetOk() (*int64, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *TopicMessage) HasOffset() bool {
	return o != nil && o.Offset != nil
}

// SetOffset gets a reference to the given int64 and assigns it to the Offset field.
func (o *TopicMessage) SetOffset(v int64) {
	o.Offset = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *TopicMessage) GetTimestamp() int64 {
	if o == nil || o.Timestamp == nil {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopicMessage) GetTimestampOk() (*int64, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *TopicMessage) HasTimestamp() bool {
	return o != nil && o.Timestamp != nil
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *TopicMessage) SetTimestamp(v int64) {
	o.Timestamp = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *TopicMessage) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopicMessage) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *TopicMessage) HasKey() bool {
	return o != nil && o.Key != nil
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *TopicMessage) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TopicMessage) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopicMessage) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TopicMessage) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *TopicMessage) SetValue(v string) {
	o.Value = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TopicMessage) MarshalJSON() ([]byte, error) {
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
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TopicMessage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Topic     *string `json:"topic,omitempty"`
		Partition *int32  `json:"partition,omitempty"`
		Offset    *int64  `json:"offset,omitempty"`
		Timestamp *int64  `json:"timestamp,omitempty"`
		Key       *string `json:"key,omitempty"`
		Value     *string `json:"value,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"topic", "partition", "offset", "timestamp", "key", "value"})
	} else {
		return err
	}
	o.Topic = all.Topic
	o.Partition = all.Partition
	o.Offset = all.Offset
	o.Timestamp = all.Timestamp
	o.Key = all.Key
	o.Value = all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
