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
	MessageId string             `json:"messageId"`
	Role      AiAgentMessageRole `json:"role"`
	RunId     *string            `json:"runId,omitempty"`
	Summary   *string            `json:"summary,omitempty"`
	// Deprecated for assistant UI replay; use parts instead.
	Content   *string              `json:"content,omitempty"`
	Parts     []AiAgentMessagePart `json:"parts,omitempty"`
	CreatedAt time.Time            `json:"createdAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentMessage instantiates a new AiAgentMessage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentMessage(messageId string, role AiAgentMessageRole, createdAt time.Time) *AiAgentMessage {
	this := AiAgentMessage{}
	this.MessageId = messageId
	this.Role = role
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

// GetRunId returns the RunId field value if set, zero value otherwise.
func (o *AiAgentMessage) GetRunId() string {
	if o == nil || o.RunId == nil {
		var ret string
		return ret
	}
	return *o.RunId
}

// GetRunIdOk returns a tuple with the RunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentMessage) GetRunIdOk() (*string, bool) {
	if o == nil || o.RunId == nil {
		return nil, false
	}
	return o.RunId, true
}

// HasRunId returns a boolean if a field has been set.
func (o *AiAgentMessage) HasRunId() bool {
	return o != nil && o.RunId != nil
}

// SetRunId gets a reference to the given string and assigns it to the RunId field.
func (o *AiAgentMessage) SetRunId(v string) {
	o.RunId = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *AiAgentMessage) GetSummary() string {
	if o == nil || o.Summary == nil {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentMessage) GetSummaryOk() (*string, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *AiAgentMessage) HasSummary() bool {
	return o != nil && o.Summary != nil
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *AiAgentMessage) SetSummary(v string) {
	o.Summary = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *AiAgentMessage) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentMessage) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *AiAgentMessage) HasContent() bool {
	return o != nil && o.Content != nil
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *AiAgentMessage) SetContent(v string) {
	o.Content = &v
}

// GetParts returns the Parts field value if set, zero value otherwise.
func (o *AiAgentMessage) GetParts() []AiAgentMessagePart {
	if o == nil || o.Parts == nil {
		var ret []AiAgentMessagePart
		return ret
	}
	return o.Parts
}

// GetPartsOk returns a tuple with the Parts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentMessage) GetPartsOk() (*[]AiAgentMessagePart, bool) {
	if o == nil || o.Parts == nil {
		return nil, false
	}
	return &o.Parts, true
}

// HasParts returns a boolean if a field has been set.
func (o *AiAgentMessage) HasParts() bool {
	return o != nil && o.Parts != nil
}

// SetParts gets a reference to the given []AiAgentMessagePart and assigns it to the Parts field.
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
	toSerialize["role"] = o.Role
	if o.RunId != nil {
		toSerialize["runId"] = o.RunId
	}
	if o.Summary != nil {
		toSerialize["summary"] = o.Summary
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.Parts != nil {
		toSerialize["parts"] = o.Parts
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
func (o *AiAgentMessage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MessageId *string              `json:"messageId"`
		Role      *AiAgentMessageRole  `json:"role"`
		RunId     *string              `json:"runId,omitempty"`
		Summary   *string              `json:"summary,omitempty"`
		Content   *string              `json:"content,omitempty"`
		Parts     []AiAgentMessagePart `json:"parts,omitempty"`
		CreatedAt *time.Time           `json:"createdAt"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.MessageId == nil {
		return fmt.Errorf("required field messageId missing")
	}
	if all.Role == nil {
		return fmt.Errorf("required field role missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field createdAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"messageId", "role", "runId", "summary", "content", "parts", "createdAt"})
	} else {
		return err
	}

	hasInvalidField := false
	o.MessageId = *all.MessageId
	if !all.Role.IsValid() {
		hasInvalidField = true
	} else {
		o.Role = *all.Role
	}
	o.RunId = all.RunId
	o.Summary = all.Summary
	o.Content = all.Content
	o.Parts = all.Parts
	o.CreatedAt = *all.CreatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
