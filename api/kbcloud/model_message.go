// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
	"github.com/google/uuid"
)

type Message struct {
	// Message ID
	Id *uuid.UUID `json:"id,omitempty"`
	// ID of the conversation this message belongs to
	ConversationId *uuid.UUID `json:"conversationId,omitempty"`
	// Type of the sender (user, assistant, or system)
	SenderType *SenderType `json:"senderType,omitempty"`
	// Content of the message
	Content *string `json:"content,omitempty"`
	// Optional metadata for AI messages (e.g., tool calls, results)
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Timestamp of when the message was created
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMessage instantiates a new Message object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMessage() *Message {
	this := Message{}
	return &this
}

// NewMessageWithDefaults instantiates a new Message object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMessageWithDefaults() *Message {
	this := Message{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Message) GetId() uuid.UUID {
	if o == nil || o.Id == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetIdOk() (*uuid.UUID, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Message) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given uuid.UUID and assigns it to the Id field.
func (o *Message) SetId(v uuid.UUID) {
	o.Id = &v
}

// GetConversationId returns the ConversationId field value if set, zero value otherwise.
func (o *Message) GetConversationId() uuid.UUID {
	if o == nil || o.ConversationId == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.ConversationId
}

// GetConversationIdOk returns a tuple with the ConversationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetConversationIdOk() (*uuid.UUID, bool) {
	if o == nil || o.ConversationId == nil {
		return nil, false
	}
	return o.ConversationId, true
}

// HasConversationId returns a boolean if a field has been set.
func (o *Message) HasConversationId() bool {
	return o != nil && o.ConversationId != nil
}

// SetConversationId gets a reference to the given uuid.UUID and assigns it to the ConversationId field.
func (o *Message) SetConversationId(v uuid.UUID) {
	o.ConversationId = &v
}

// GetSenderType returns the SenderType field value if set, zero value otherwise.
func (o *Message) GetSenderType() SenderType {
	if o == nil || o.SenderType == nil {
		var ret SenderType
		return ret
	}
	return *o.SenderType
}

// GetSenderTypeOk returns a tuple with the SenderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetSenderTypeOk() (*SenderType, bool) {
	if o == nil || o.SenderType == nil {
		return nil, false
	}
	return o.SenderType, true
}

// HasSenderType returns a boolean if a field has been set.
func (o *Message) HasSenderType() bool {
	return o != nil && o.SenderType != nil
}

// SetSenderType gets a reference to the given SenderType and assigns it to the SenderType field.
func (o *Message) SetSenderType(v SenderType) {
	o.SenderType = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *Message) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *Message) HasContent() bool {
	return o != nil && o.Content != nil
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *Message) SetContent(v string) {
	o.Content = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Message) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Message) HasMetadata() bool {
	return o != nil && o.Metadata != nil
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *Message) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Message) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Message) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Message) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o Message) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ConversationId != nil {
		toSerialize["conversationId"] = o.ConversationId
	}
	if o.SenderType != nil {
		toSerialize["senderType"] = o.SenderType
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Message) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id             *uuid.UUID             `json:"id,omitempty"`
		ConversationId *uuid.UUID             `json:"conversationId,omitempty"`
		SenderType     *SenderType            `json:"senderType,omitempty"`
		Content        *string                `json:"content,omitempty"`
		Metadata       map[string]interface{} `json:"metadata,omitempty"`
		CreatedAt      *time.Time             `json:"createdAt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "conversationId", "senderType", "content", "metadata", "createdAt"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = all.Id
	o.ConversationId = all.ConversationId
	if all.SenderType != nil && !all.SenderType.IsValid() {
		hasInvalidField = true
	} else {
		o.SenderType = all.SenderType
	}
	o.Content = all.Content
	o.Metadata = all.Metadata
	o.CreatedAt = all.CreatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
