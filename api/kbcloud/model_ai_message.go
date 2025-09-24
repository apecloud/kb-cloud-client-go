// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"time"

	"github.com/google/uuid"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiMessage struct {
	// Message ID
	Id *uuid.UUID `json:"id,omitempty"`
	// ID of the conversation this message belongs to
	ConversationId *uuid.UUID `json:"conversationId,omitempty"`
	// Type of the sender (user, assistant, or system)
	SenderType *SenderType `json:"senderType,omitempty"`
	// Content of the message
	Content *string `json:"content,omitempty"`
	// Whether the message is favorited by user
	IsFavorited *bool `json:"isFavorited,omitempty"`
	// Optional metadata for AI messages (e.g., tool calls, results)
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Timestamp of when the message was created
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Timestamp of when the message was last updated
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiMessage instantiates a new AiMessage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiMessage() *AiMessage {
	this := AiMessage{}
	var isFavorited bool = false
	this.IsFavorited = &isFavorited
	return &this
}

// NewAiMessageWithDefaults instantiates a new AiMessage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiMessageWithDefaults() *AiMessage {
	this := AiMessage{}
	var isFavorited bool = false
	this.IsFavorited = &isFavorited
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AiMessage) GetId() uuid.UUID {
	if o == nil || o.Id == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiMessage) GetIdOk() (*uuid.UUID, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AiMessage) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given uuid.UUID and assigns it to the Id field.
func (o *AiMessage) SetId(v uuid.UUID) {
	o.Id = &v
}

// GetConversationId returns the ConversationId field value if set, zero value otherwise.
func (o *AiMessage) GetConversationId() uuid.UUID {
	if o == nil || o.ConversationId == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.ConversationId
}

// GetConversationIdOk returns a tuple with the ConversationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiMessage) GetConversationIdOk() (*uuid.UUID, bool) {
	if o == nil || o.ConversationId == nil {
		return nil, false
	}
	return o.ConversationId, true
}

// HasConversationId returns a boolean if a field has been set.
func (o *AiMessage) HasConversationId() bool {
	return o != nil && o.ConversationId != nil
}

// SetConversationId gets a reference to the given uuid.UUID and assigns it to the ConversationId field.
func (o *AiMessage) SetConversationId(v uuid.UUID) {
	o.ConversationId = &v
}

// GetSenderType returns the SenderType field value if set, zero value otherwise.
func (o *AiMessage) GetSenderType() SenderType {
	if o == nil || o.SenderType == nil {
		var ret SenderType
		return ret
	}
	return *o.SenderType
}

// GetSenderTypeOk returns a tuple with the SenderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiMessage) GetSenderTypeOk() (*SenderType, bool) {
	if o == nil || o.SenderType == nil {
		return nil, false
	}
	return o.SenderType, true
}

// HasSenderType returns a boolean if a field has been set.
func (o *AiMessage) HasSenderType() bool {
	return o != nil && o.SenderType != nil
}

// SetSenderType gets a reference to the given SenderType and assigns it to the SenderType field.
func (o *AiMessage) SetSenderType(v SenderType) {
	o.SenderType = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *AiMessage) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiMessage) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *AiMessage) HasContent() bool {
	return o != nil && o.Content != nil
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *AiMessage) SetContent(v string) {
	o.Content = &v
}

// GetIsFavorited returns the IsFavorited field value if set, zero value otherwise.
func (o *AiMessage) GetIsFavorited() bool {
	if o == nil || o.IsFavorited == nil {
		var ret bool
		return ret
	}
	return *o.IsFavorited
}

// GetIsFavoritedOk returns a tuple with the IsFavorited field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiMessage) GetIsFavoritedOk() (*bool, bool) {
	if o == nil || o.IsFavorited == nil {
		return nil, false
	}
	return o.IsFavorited, true
}

// HasIsFavorited returns a boolean if a field has been set.
func (o *AiMessage) HasIsFavorited() bool {
	return o != nil && o.IsFavorited != nil
}

// SetIsFavorited gets a reference to the given bool and assigns it to the IsFavorited field.
func (o *AiMessage) SetIsFavorited(v bool) {
	o.IsFavorited = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *AiMessage) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiMessage) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *AiMessage) HasMetadata() bool {
	return o != nil && o.Metadata != nil
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *AiMessage) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AiMessage) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiMessage) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AiMessage) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AiMessage) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AiMessage) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiMessage) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AiMessage) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *AiMessage) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiMessage) MarshalJSON() ([]byte, error) {
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
	if o.IsFavorited != nil {
		toSerialize["isFavorited"] = o.IsFavorited
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
	if o.UpdatedAt != nil {
		if o.UpdatedAt.Nanosecond() == 0 {
			toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiMessage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id             *uuid.UUID             `json:"id,omitempty"`
		ConversationId *uuid.UUID             `json:"conversationId,omitempty"`
		SenderType     *SenderType            `json:"senderType,omitempty"`
		Content        *string                `json:"content,omitempty"`
		IsFavorited    *bool                  `json:"isFavorited,omitempty"`
		Metadata       map[string]interface{} `json:"metadata,omitempty"`
		CreatedAt      *time.Time             `json:"createdAt,omitempty"`
		UpdatedAt      *time.Time             `json:"updatedAt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "conversationId", "senderType", "content", "isFavorited", "metadata", "createdAt", "updatedAt"})
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
	o.IsFavorited = all.IsFavorited
	o.Metadata = all.Metadata
	o.CreatedAt = all.CreatedAt
	o.UpdatedAt = all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
