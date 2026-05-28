// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentMessage struct {
	MessageId      string               `json:"messageId"`
	ConversationId string               `json:"conversationId"`
	TurnId         *string              `json:"turnId,omitempty"`
	Role           AiAgentMessageRole   `json:"role"`
	Contexts       []AiAgentScope       `json:"contexts,omitempty"`
	Parts          []AiAgentMessagePart `json:"parts"`
	CreatedAt      time.Time            `json:"createdAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentMessage instantiates a new AiAgentMessage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentMessage(messageId string, conversationId string, role AiAgentMessageRole, parts []AiAgentMessagePart, createdAt time.Time) *AiAgentMessage {
	this := AiAgentMessage{}
	this.MessageId = messageId
	this.ConversationId = conversationId
	this.Role = role
	this.Parts = parts
	this.CreatedAt = createdAt
	return &this
}

// NewAiAgentMessageWithDefaults instantiates a new AiAgentMessage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentMessageWithDefaults() *AiAgentMessage {
	this := AiAgentMessage{}
	return &this
}

// GetMessageId returns the MessageId field value.
func (o *AiAgentMessage) GetMessageId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value
// and a boolean to check if the value has been set.
func (o *AiAgentMessage) GetMessageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageId, true
}

// SetMessageId sets field value.
func (o *AiAgentMessage) SetMessageId(v string) {
	o.MessageId = v
}

// GetConversationId returns the ConversationId field value.
func (o *AiAgentMessage) GetConversationId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ConversationId
}

// GetConversationIdOk returns a tuple with the ConversationId field value
// and a boolean to check if the value has been set.
func (o *AiAgentMessage) GetConversationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConversationId, true
}

// SetConversationId sets field value.
func (o *AiAgentMessage) SetConversationId(v string) {
	o.ConversationId = v
}

// GetTurnId returns the TurnId field value if set, zero value otherwise.
func (o *AiAgentMessage) GetTurnId() string {
	if o == nil || o.TurnId == nil {
		var ret string
		return ret
	}
	return *o.TurnId
}

// GetTurnIdOk returns a tuple with the TurnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentMessage) GetTurnIdOk() (*string, bool) {
	if o == nil || o.TurnId == nil {
		return nil, false
	}
	return o.TurnId, true
}

// HasTurnId returns a boolean if a field has been set.
func (o *AiAgentMessage) HasTurnId() bool {
	return o != nil && o.TurnId != nil
}

// SetTurnId gets a reference to the given string and assigns it to the TurnId field.
func (o *AiAgentMessage) SetTurnId(v string) {
	o.TurnId = &v
}

// GetRole returns the Role field value.
func (o *AiAgentMessage) GetRole() AiAgentMessageRole {
	if o == nil {
		var ret AiAgentMessageRole
		return ret
	}
	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *AiAgentMessage) GetRoleOk() (*AiAgentMessageRole, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value.
func (o *AiAgentMessage) SetRole(v AiAgentMessageRole) {
	o.Role = v
}

// GetContexts returns the Contexts field value if set, zero value otherwise.
func (o *AiAgentMessage) GetContexts() []AiAgentScope {
	if o == nil || o.Contexts == nil {
		var ret []AiAgentScope
		return ret
	}
	return o.Contexts
}

// GetContextsOk returns a tuple with the Contexts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentMessage) GetContextsOk() (*[]AiAgentScope, bool) {
	if o == nil || o.Contexts == nil {
		return nil, false
	}
	return &o.Contexts, true
}

// HasContexts returns a boolean if a field has been set.
func (o *AiAgentMessage) HasContexts() bool {
	return o != nil && o.Contexts != nil
}

// SetContexts gets a reference to the given []AiAgentScope and assigns it to the Contexts field.
func (o *AiAgentMessage) SetContexts(v []AiAgentScope) {
	o.Contexts = v
}

// GetParts returns the Parts field value.
func (o *AiAgentMessage) GetParts() []AiAgentMessagePart {
	if o == nil {
		var ret []AiAgentMessagePart
		return ret
	}
	return o.Parts
}

// GetPartsOk returns a tuple with the Parts field value
// and a boolean to check if the value has been set.
func (o *AiAgentMessage) GetPartsOk() (*[]AiAgentMessagePart, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parts, true
}

// SetParts sets field value.
func (o *AiAgentMessage) SetParts(v []AiAgentMessagePart) {
	o.Parts = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *AiAgentMessage) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AiAgentMessage) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *AiAgentMessage) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["messageId"] = o.MessageId
	toSerialize["conversationId"] = o.ConversationId
	if o.TurnId != nil {
		toSerialize["turnId"] = o.TurnId
	}
	toSerialize["role"] = o.Role
	if o.Contexts != nil {
		toSerialize["contexts"] = o.Contexts
	}
	toSerialize["parts"] = o.Parts
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
func (o *AiAgentMessage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MessageId      *string               `json:"messageId"`
		ConversationId *string               `json:"conversationId"`
		TurnId         *string               `json:"turnId,omitempty"`
		Role           *AiAgentMessageRole   `json:"role"`
		Contexts       []AiAgentScope        `json:"contexts,omitempty"`
		Parts          *[]AiAgentMessagePart `json:"parts"`
		CreatedAt      *time.Time            `json:"createdAt"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.MessageId == nil {
		return fmt.Errorf("required field messageId missing")
	}
	if all.ConversationId == nil {
		return fmt.Errorf("required field conversationId missing")
	}
	if all.Role == nil {
		return fmt.Errorf("required field role missing")
	}
	if all.Parts == nil {
		return fmt.Errorf("required field parts missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field createdAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"messageId", "conversationId", "turnId", "role", "contexts", "parts", "createdAt"})
	} else {
		return err
	}

	hasInvalidField := false
	o.MessageId = *all.MessageId
	o.ConversationId = *all.ConversationId
	o.TurnId = all.TurnId
	if !all.Role.IsValid() {
		hasInvalidField = true
	} else {
		o.Role = *all.Role
	}
	o.Contexts = all.Contexts
	o.Parts = *all.Parts
	o.CreatedAt = *all.CreatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
