// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type TopicMessageList struct {
	Items []TopicMessage

	// UnparsedObject contains the raw value of the array if there was an error when deserializing into the struct
	UnparsedObject []interface{} `json:"-"`
}

// NewTopicMessageList instantiates a new TopicMessageList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTopicMessageList() *TopicMessageList {
	this := TopicMessageList{}
	return &this
}

// NewTopicMessageListWithDefaults instantiates a new TopicMessageList object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTopicMessageListWithDefaults() *TopicMessageList {
	this := TopicMessageList{}
	return &this
}

// MarshalJSON serializes the struct using spec logic.
func (o TopicMessageList) MarshalJSON() ([]byte, error) {
	toSerialize := make([]interface{}, len(o.Items))
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TopicMessageList) UnmarshalJSON(bytes []byte) (err error) {
	if err = common.Unmarshal(bytes, &o.Items); err != nil {
		return err
	}

	if o.Items != nil && len(o.Items) > 0 {
		for _, v := range o.Items {
			if v.UnparsedObject != nil {
				return common.Unmarshal(bytes, &o.UnparsedObject)
			}
		}
	}

	return nil
}
