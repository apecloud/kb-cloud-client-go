// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package data

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)


 
type LiveMessage struct {
	// 消息ID
	Id *string `json:"id,omitempty"`
	Message *TopicMessage `json:"message,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewLiveMessage instantiates a new LiveMessage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLiveMessage() *LiveMessage {
	this := LiveMessage{}
	return &this
}

// NewLiveMessageWithDefaults instantiates a new LiveMessage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLiveMessageWithDefaults() *LiveMessage {
	this := LiveMessage{}
	return &this
}
// GetId returns the Id field value if set, zero value otherwise.
func (o *LiveMessage) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveMessage) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LiveMessage) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LiveMessage) SetId(v string) {
	o.Id = &v
}


// GetMessage returns the Message field value if set, zero value otherwise.
func (o *LiveMessage) GetMessage() TopicMessage {
	if o == nil || o.Message == nil {
		var ret TopicMessage
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveMessage) GetMessageOk() (*TopicMessage, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *LiveMessage) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given TopicMessage and assigns it to the Message field.
func (o *LiveMessage) SetMessage(v TopicMessage) {
	o.Message = &v
}



// MarshalJSON serializes the struct using spec logic.
func (o LiveMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LiveMessage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id *string `json:"id,omitempty"`
		Message *TopicMessage `json:"message,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{ "id", "message",  })
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = all.Id
	if  all.Message != nil && all.Message.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Message = all.Message

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
