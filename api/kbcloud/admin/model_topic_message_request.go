// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type TopicMessageRequest struct {
	// 指定消息将被发送到的Kafka分区
	Partition *int32 `json:"partition,omitempty"`
	// 消息的键（可选）
	Key *string `json:"key,omitempty"`
	// 消息的实际内容
	Value *string `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTopicMessageRequest instantiates a new TopicMessageRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTopicMessageRequest() *TopicMessageRequest {
	this := TopicMessageRequest{}
	return &this
}

// NewTopicMessageRequestWithDefaults instantiates a new TopicMessageRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTopicMessageRequestWithDefaults() *TopicMessageRequest {
	this := TopicMessageRequest{}
	return &this
}

// GetPartition returns the Partition field value if set, zero value otherwise.
func (o *TopicMessageRequest) GetPartition() int32 {
	if o == nil || o.Partition == nil {
		var ret int32
		return ret
	}
	return *o.Partition
}

// GetPartitionOk returns a tuple with the Partition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopicMessageRequest) GetPartitionOk() (*int32, bool) {
	if o == nil || o.Partition == nil {
		return nil, false
	}
	return o.Partition, true
}

// HasPartition returns a boolean if a field has been set.
func (o *TopicMessageRequest) HasPartition() bool {
	return o != nil && o.Partition != nil
}

// SetPartition gets a reference to the given int32 and assigns it to the Partition field.
func (o *TopicMessageRequest) SetPartition(v int32) {
	o.Partition = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *TopicMessageRequest) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopicMessageRequest) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *TopicMessageRequest) HasKey() bool {
	return o != nil && o.Key != nil
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *TopicMessageRequest) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TopicMessageRequest) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopicMessageRequest) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TopicMessageRequest) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *TopicMessageRequest) SetValue(v string) {
	o.Value = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TopicMessageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Partition != nil {
		toSerialize["partition"] = o.Partition
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
func (o *TopicMessageRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Partition *int32  `json:"partition,omitempty"`
		Key       *string `json:"key,omitempty"`
		Value     *string `json:"value,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"partition", "key", "value"})
	} else {
		return err
	}
	o.Partition = all.Partition
	o.Key = all.Key
	o.Value = all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
