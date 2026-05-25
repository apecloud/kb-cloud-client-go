// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentEvent struct {
	EventId        string                 `json:"eventId"`
	ConversationId string                 `json:"conversationId"`
	TurnId         *string                `json:"turnId,omitempty"`
	Type           AiAgentEventType       `json:"type"`
	Status         *string                `json:"status,omitempty"`
	Message        *string                `json:"message,omitempty"`
	Payload        map[string]interface{} `json:"payload,omitempty"`
	CreatedAt      time.Time              `json:"createdAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentEvent instantiates a new AiAgentEvent object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentEvent(eventId string, conversationId string, typeVar AiAgentEventType, createdAt time.Time) *AiAgentEvent {
	this := AiAgentEvent{}
	this.EventId = eventId
	this.ConversationId = conversationId
	this.Type = typeVar
	this.CreatedAt = createdAt
	return &this
}

// NewAiAgentEventWithDefaults instantiates a new AiAgentEvent object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentEventWithDefaults() *AiAgentEvent {
	this := AiAgentEvent{}
	return &this
}

// GetEventId returns the EventId field value.
func (o *AiAgentEvent) GetEventId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value
// and a boolean to check if the value has been set.
func (o *AiAgentEvent) GetEventIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventId, true
}

// SetEventId sets field value.
func (o *AiAgentEvent) SetEventId(v string) {
	o.EventId = v
}

// GetConversationId returns the ConversationId field value.
func (o *AiAgentEvent) GetConversationId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ConversationId
}

// GetConversationIdOk returns a tuple with the ConversationId field value
// and a boolean to check if the value has been set.
func (o *AiAgentEvent) GetConversationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConversationId, true
}

// SetConversationId sets field value.
func (o *AiAgentEvent) SetConversationId(v string) {
	o.ConversationId = v
}

// GetTurnId returns the TurnId field value if set, zero value otherwise.
func (o *AiAgentEvent) GetTurnId() string {
	if o == nil || o.TurnId == nil {
		var ret string
		return ret
	}
	return *o.TurnId
}

// GetTurnIdOk returns a tuple with the TurnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEvent) GetTurnIdOk() (*string, bool) {
	if o == nil || o.TurnId == nil {
		return nil, false
	}
	return o.TurnId, true
}

// HasTurnId returns a boolean if a field has been set.
func (o *AiAgentEvent) HasTurnId() bool {
	return o != nil && o.TurnId != nil
}

// SetTurnId gets a reference to the given string and assigns it to the TurnId field.
func (o *AiAgentEvent) SetTurnId(v string) {
	o.TurnId = &v
}

// GetType returns the Type field value.
func (o *AiAgentEvent) GetType() AiAgentEventType {
	if o == nil {
		var ret AiAgentEventType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AiAgentEvent) GetTypeOk() (*AiAgentEventType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AiAgentEvent) SetType(v AiAgentEventType) {
	o.Type = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AiAgentEvent) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEvent) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AiAgentEvent) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AiAgentEvent) SetStatus(v string) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *AiAgentEvent) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEvent) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *AiAgentEvent) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *AiAgentEvent) SetMessage(v string) {
	o.Message = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *AiAgentEvent) GetPayload() map[string]interface{} {
	if o == nil || o.Payload == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEvent) GetPayloadOk() (*map[string]interface{}, bool) {
	if o == nil || o.Payload == nil {
		return nil, false
	}
	return &o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *AiAgentEvent) HasPayload() bool {
	return o != nil && o.Payload != nil
}

// SetPayload gets a reference to the given map[string]interface{} and assigns it to the Payload field.
func (o *AiAgentEvent) SetPayload(v map[string]interface{}) {
	o.Payload = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *AiAgentEvent) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AiAgentEvent) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *AiAgentEvent) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["eventId"] = o.EventId
	toSerialize["conversationId"] = o.ConversationId
	if o.TurnId != nil {
		toSerialize["turnId"] = o.TurnId
	}
	toSerialize["type"] = o.Type
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Payload != nil {
		toSerialize["payload"] = o.Payload
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentEvent) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EventId        *string                `json:"eventId"`
		ConversationId *string                `json:"conversationId"`
		TurnId         *string                `json:"turnId,omitempty"`
		Type           *AiAgentEventType      `json:"type"`
		Status         *string                `json:"status,omitempty"`
		Message        *string                `json:"message,omitempty"`
		Payload        map[string]interface{} `json:"payload,omitempty"`
		CreatedAt      *time.Time             `json:"createdAt"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.EventId == nil {
		return fmt.Errorf("required field eventId missing")
	}
	if all.ConversationId == nil {
		return fmt.Errorf("required field conversationId missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field createdAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"eventId", "conversationId", "turnId", "type", "status", "message", "payload", "createdAt"})
	} else {
		return err
	}

	hasInvalidField := false
	o.EventId = *all.EventId
	o.ConversationId = *all.ConversationId
	o.TurnId = all.TurnId
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.Status = all.Status
	o.Message = all.Message
	o.Payload = all.Payload
	o.CreatedAt = *all.CreatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
