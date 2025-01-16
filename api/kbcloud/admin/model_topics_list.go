// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type TopicsList struct {
	Topics []Topic `json:"topics"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTopicsList instantiates a new TopicsList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTopicsList(topics []Topic) *TopicsList {
	this := TopicsList{}
	this.Topics = topics
	return &this
}

// NewTopicsListWithDefaults instantiates a new TopicsList object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTopicsListWithDefaults() *TopicsList {
	this := TopicsList{}
	return &this
}

// GetTopics returns the Topics field value.
func (o *TopicsList) GetTopics() []Topic {
	if o == nil {
		var ret []Topic
		return ret
	}
	return o.Topics
}

// GetTopicsOk returns a tuple with the Topics field value
// and a boolean to check if the value has been set.
func (o *TopicsList) GetTopicsOk() (*[]Topic, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Topics, true
}

// SetTopics sets field value.
func (o *TopicsList) SetTopics(v []Topic) {
	o.Topics = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TopicsList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["topics"] = o.Topics

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TopicsList) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Topics *[]Topic `json:"topics"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Topics == nil {
		return fmt.Errorf("required field topics missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"topics"})
	} else {
		return err
	}
	o.Topics = *all.Topics

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
