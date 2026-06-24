// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type RedisPubSubPublishRequest struct {
	Channel string `json:"channel"`
	Message string `json:"message"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRedisPubSubPublishRequest instantiates a new RedisPubSubPublishRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRedisPubSubPublishRequest(channel string, message string) *RedisPubSubPublishRequest {
	this := RedisPubSubPublishRequest{}
	this.Channel = channel
	this.Message = message
	return &this
}

// NewRedisPubSubPublishRequestWithDefaults instantiates a new RedisPubSubPublishRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRedisPubSubPublishRequestWithDefaults() *RedisPubSubPublishRequest {
	this := RedisPubSubPublishRequest{}
	return &this
}

// GetChannel returns the Channel field value.
func (o *RedisPubSubPublishRequest) GetChannel() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Channel
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
func (o *RedisPubSubPublishRequest) GetChannelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Channel, true
}

// SetChannel sets field value.
func (o *RedisPubSubPublishRequest) SetChannel(v string) {
	o.Channel = v
}

// GetMessage returns the Message field value.
func (o *RedisPubSubPublishRequest) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *RedisPubSubPublishRequest) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value.
func (o *RedisPubSubPublishRequest) SetMessage(v string) {
	o.Message = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RedisPubSubPublishRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["channel"] = o.Channel
	toSerialize["message"] = o.Message

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RedisPubSubPublishRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Channel *string `json:"channel"`
		Message *string `json:"message"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Channel == nil {
		return fmt.Errorf("required field channel missing")
	}
	if all.Message == nil {
		return fmt.Errorf("required field message missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"channel", "message"})
	} else {
		return err
	}
	o.Channel = *all.Channel
	o.Message = *all.Message

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
