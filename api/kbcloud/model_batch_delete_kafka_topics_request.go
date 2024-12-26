// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

type BatchDeleteKafkaTopicsRequest struct {
	TopicNames []string `json:"topicNames,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBatchDeleteKafkaTopicsRequest instantiates a new BatchDeleteKafkaTopicsRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBatchDeleteKafkaTopicsRequest() *BatchDeleteKafkaTopicsRequest {
	this := BatchDeleteKafkaTopicsRequest{}
	return &this
}

// NewBatchDeleteKafkaTopicsRequestWithDefaults instantiates a new BatchDeleteKafkaTopicsRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBatchDeleteKafkaTopicsRequestWithDefaults() *BatchDeleteKafkaTopicsRequest {
	this := BatchDeleteKafkaTopicsRequest{}
	return &this
}

// GetTopicNames returns the TopicNames field value if set, zero value otherwise.
func (o *BatchDeleteKafkaTopicsRequest) GetTopicNames() []string {
	if o == nil || o.TopicNames == nil {
		var ret []string
		return ret
	}
	return o.TopicNames
}

// GetTopicNamesOk returns a tuple with the TopicNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchDeleteKafkaTopicsRequest) GetTopicNamesOk() (*[]string, bool) {
	if o == nil || o.TopicNames == nil {
		return nil, false
	}
	return &o.TopicNames, true
}

// HasTopicNames returns a boolean if a field has been set.
func (o *BatchDeleteKafkaTopicsRequest) HasTopicNames() bool {
	return o != nil && o.TopicNames != nil
}

// SetTopicNames gets a reference to the given []string and assigns it to the TopicNames field.
func (o *BatchDeleteKafkaTopicsRequest) SetTopicNames(v []string) {
	o.TopicNames = v
}

// MarshalJSON serializes the struct using spec logic.
func (o BatchDeleteKafkaTopicsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.TopicNames != nil {
		toSerialize["topicNames"] = o.TopicNames
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BatchDeleteKafkaTopicsRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TopicNames []string `json:"topicNames,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"topicNames"})
	} else {
		return err
	}
	o.TopicNames = all.TopicNames

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
