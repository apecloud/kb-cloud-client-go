// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type RedisPubSubMessage struct {
	Channel    *string    `json:"channel,omitempty"`
	Pattern    *string    `json:"pattern,omitempty"`
	Message    *string    `json:"message,omitempty"`
	ReceivedAt *time.Time `json:"receivedAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRedisPubSubMessage instantiates a new RedisPubSubMessage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRedisPubSubMessage() *RedisPubSubMessage {
	this := RedisPubSubMessage{}
	return &this
}

// NewRedisPubSubMessageWithDefaults instantiates a new RedisPubSubMessage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRedisPubSubMessageWithDefaults() *RedisPubSubMessage {
	this := RedisPubSubMessage{}
	return &this
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *RedisPubSubMessage) GetChannel() string {
	if o == nil || o.Channel == nil {
		var ret string
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisPubSubMessage) GetChannelOk() (*string, bool) {
	if o == nil || o.Channel == nil {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *RedisPubSubMessage) HasChannel() bool {
	return o != nil && o.Channel != nil
}

// SetChannel gets a reference to the given string and assigns it to the Channel field.
func (o *RedisPubSubMessage) SetChannel(v string) {
	o.Channel = &v
}

// GetPattern returns the Pattern field value if set, zero value otherwise.
func (o *RedisPubSubMessage) GetPattern() string {
	if o == nil || o.Pattern == nil {
		var ret string
		return ret
	}
	return *o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisPubSubMessage) GetPatternOk() (*string, bool) {
	if o == nil || o.Pattern == nil {
		return nil, false
	}
	return o.Pattern, true
}

// HasPattern returns a boolean if a field has been set.
func (o *RedisPubSubMessage) HasPattern() bool {
	return o != nil && o.Pattern != nil
}

// SetPattern gets a reference to the given string and assigns it to the Pattern field.
func (o *RedisPubSubMessage) SetPattern(v string) {
	o.Pattern = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *RedisPubSubMessage) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisPubSubMessage) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *RedisPubSubMessage) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *RedisPubSubMessage) SetMessage(v string) {
	o.Message = &v
}

// GetReceivedAt returns the ReceivedAt field value if set, zero value otherwise.
func (o *RedisPubSubMessage) GetReceivedAt() time.Time {
	if o == nil || o.ReceivedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ReceivedAt
}

// GetReceivedAtOk returns a tuple with the ReceivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisPubSubMessage) GetReceivedAtOk() (*time.Time, bool) {
	if o == nil || o.ReceivedAt == nil {
		return nil, false
	}
	return o.ReceivedAt, true
}

// HasReceivedAt returns a boolean if a field has been set.
func (o *RedisPubSubMessage) HasReceivedAt() bool {
	return o != nil && o.ReceivedAt != nil
}

// SetReceivedAt gets a reference to the given time.Time and assigns it to the ReceivedAt field.
func (o *RedisPubSubMessage) SetReceivedAt(v time.Time) {
	o.ReceivedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RedisPubSubMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Channel != nil {
		toSerialize["channel"] = o.Channel
	}
	if o.Pattern != nil {
		toSerialize["pattern"] = o.Pattern
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.ReceivedAt != nil {
		if o.ReceivedAt.Nanosecond() == 0 {
			toSerialize["receivedAt"] = o.ReceivedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["receivedAt"] = o.ReceivedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RedisPubSubMessage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Channel    *string    `json:"channel,omitempty"`
		Pattern    *string    `json:"pattern,omitempty"`
		Message    *string    `json:"message,omitempty"`
		ReceivedAt *time.Time `json:"receivedAt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"channel", "pattern", "message", "receivedAt"})
	} else {
		return err
	}
	o.Channel = all.Channel
	o.Pattern = all.Pattern
	o.Message = all.Message
	o.ReceivedAt = all.ReceivedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
