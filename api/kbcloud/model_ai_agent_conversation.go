// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentConversation struct {
	ConversationId string `json:"conversationId"`
	Title          string `json:"title"`
	// Organization-scoped single-cluster Agent execution scope.
	Scope               *AiAgentResourceScope `json:"scope,omitempty"`
	ConversationSummary *string               `json:"conversationSummary,omitempty"`
	LastCompactedAt     *time.Time            `json:"lastCompactedAt,omitempty"`
	SummaryVersion      *string               `json:"summaryVersion,omitempty"`
	Messages            []AiAgentMessage      `json:"messages,omitempty"`
	Runs                []AiAgentRun          `json:"runs,omitempty"`
	CreatedAt           time.Time             `json:"createdAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentConversation instantiates a new AiAgentConversation object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentConversation(conversationId string, title string, createdAt time.Time) *AiAgentConversation {
	this := AiAgentConversation{}
	this.ConversationId = conversationId
	this.Title = title
	this.CreatedAt = createdAt
	return &this
}

// NewAiAgentConversationWithDefaults instantiates a new AiAgentConversation object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentConversationWithDefaults() *AiAgentConversation {
	this := AiAgentConversation{}
	return &this
}

// GetConversationId returns the ConversationId field value.
func (o *AiAgentConversation) GetConversationId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ConversationId
}

// GetConversationIdOk returns a tuple with the ConversationId field value
// and a boolean to check if the value has been set.
func (o *AiAgentConversation) GetConversationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConversationId, true
}

// SetConversationId sets field value.
func (o *AiAgentConversation) SetConversationId(v string) {
	o.ConversationId = v
}

// GetTitle returns the Title field value.
func (o *AiAgentConversation) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *AiAgentConversation) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *AiAgentConversation) SetTitle(v string) {
	o.Title = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *AiAgentConversation) GetScope() AiAgentResourceScope {
	if o == nil || o.Scope == nil {
		var ret AiAgentResourceScope
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentConversation) GetScopeOk() (*AiAgentResourceScope, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *AiAgentConversation) HasScope() bool {
	return o != nil && o.Scope != nil
}

// SetScope gets a reference to the given AiAgentResourceScope and assigns it to the Scope field.
func (o *AiAgentConversation) SetScope(v AiAgentResourceScope) {
	o.Scope = &v
}

// GetConversationSummary returns the ConversationSummary field value if set, zero value otherwise.
func (o *AiAgentConversation) GetConversationSummary() string {
	if o == nil || o.ConversationSummary == nil {
		var ret string
		return ret
	}
	return *o.ConversationSummary
}

// GetConversationSummaryOk returns a tuple with the ConversationSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentConversation) GetConversationSummaryOk() (*string, bool) {
	if o == nil || o.ConversationSummary == nil {
		return nil, false
	}
	return o.ConversationSummary, true
}

// HasConversationSummary returns a boolean if a field has been set.
func (o *AiAgentConversation) HasConversationSummary() bool {
	return o != nil && o.ConversationSummary != nil
}

// SetConversationSummary gets a reference to the given string and assigns it to the ConversationSummary field.
func (o *AiAgentConversation) SetConversationSummary(v string) {
	o.ConversationSummary = &v
}

// GetLastCompactedAt returns the LastCompactedAt field value if set, zero value otherwise.
func (o *AiAgentConversation) GetLastCompactedAt() time.Time {
	if o == nil || o.LastCompactedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.LastCompactedAt
}

// GetLastCompactedAtOk returns a tuple with the LastCompactedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentConversation) GetLastCompactedAtOk() (*time.Time, bool) {
	if o == nil || o.LastCompactedAt == nil {
		return nil, false
	}
	return o.LastCompactedAt, true
}

// HasLastCompactedAt returns a boolean if a field has been set.
func (o *AiAgentConversation) HasLastCompactedAt() bool {
	return o != nil && o.LastCompactedAt != nil
}

// SetLastCompactedAt gets a reference to the given time.Time and assigns it to the LastCompactedAt field.
func (o *AiAgentConversation) SetLastCompactedAt(v time.Time) {
	o.LastCompactedAt = &v
}

// GetSummaryVersion returns the SummaryVersion field value if set, zero value otherwise.
func (o *AiAgentConversation) GetSummaryVersion() string {
	if o == nil || o.SummaryVersion == nil {
		var ret string
		return ret
	}
	return *o.SummaryVersion
}

// GetSummaryVersionOk returns a tuple with the SummaryVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentConversation) GetSummaryVersionOk() (*string, bool) {
	if o == nil || o.SummaryVersion == nil {
		return nil, false
	}
	return o.SummaryVersion, true
}

// HasSummaryVersion returns a boolean if a field has been set.
func (o *AiAgentConversation) HasSummaryVersion() bool {
	return o != nil && o.SummaryVersion != nil
}

// SetSummaryVersion gets a reference to the given string and assigns it to the SummaryVersion field.
func (o *AiAgentConversation) SetSummaryVersion(v string) {
	o.SummaryVersion = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *AiAgentConversation) GetMessages() []AiAgentMessage {
	if o == nil || o.Messages == nil {
		var ret []AiAgentMessage
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentConversation) GetMessagesOk() (*[]AiAgentMessage, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return &o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *AiAgentConversation) HasMessages() bool {
	return o != nil && o.Messages != nil
}

// SetMessages gets a reference to the given []AiAgentMessage and assigns it to the Messages field.
func (o *AiAgentConversation) SetMessages(v []AiAgentMessage) {
	o.Messages = v
}

// GetRuns returns the Runs field value if set, zero value otherwise.
func (o *AiAgentConversation) GetRuns() []AiAgentRun {
	if o == nil || o.Runs == nil {
		var ret []AiAgentRun
		return ret
	}
	return o.Runs
}

// GetRunsOk returns a tuple with the Runs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentConversation) GetRunsOk() (*[]AiAgentRun, bool) {
	if o == nil || o.Runs == nil {
		return nil, false
	}
	return &o.Runs, true
}

// HasRuns returns a boolean if a field has been set.
func (o *AiAgentConversation) HasRuns() bool {
	return o != nil && o.Runs != nil
}

// SetRuns gets a reference to the given []AiAgentRun and assigns it to the Runs field.
func (o *AiAgentConversation) SetRuns(v []AiAgentRun) {
	o.Runs = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *AiAgentConversation) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AiAgentConversation) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *AiAgentConversation) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentConversation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["conversationId"] = o.ConversationId
	toSerialize["title"] = o.Title
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.ConversationSummary != nil {
		toSerialize["conversationSummary"] = o.ConversationSummary
	}
	if o.LastCompactedAt != nil {
		if o.LastCompactedAt.Nanosecond() == 0 {
			toSerialize["lastCompactedAt"] = o.LastCompactedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["lastCompactedAt"] = o.LastCompactedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.SummaryVersion != nil {
		toSerialize["summaryVersion"] = o.SummaryVersion
	}
	if o.Messages != nil {
		toSerialize["messages"] = o.Messages
	}
	if o.Runs != nil {
		toSerialize["runs"] = o.Runs
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
func (o *AiAgentConversation) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConversationId      *string               `json:"conversationId"`
		Title               *string               `json:"title"`
		Scope               *AiAgentResourceScope `json:"scope,omitempty"`
		ConversationSummary *string               `json:"conversationSummary,omitempty"`
		LastCompactedAt     *time.Time            `json:"lastCompactedAt,omitempty"`
		SummaryVersion      *string               `json:"summaryVersion,omitempty"`
		Messages            []AiAgentMessage      `json:"messages,omitempty"`
		Runs                []AiAgentRun          `json:"runs,omitempty"`
		CreatedAt           *time.Time            `json:"createdAt"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.ConversationId == nil {
		return fmt.Errorf("required field conversationId missing")
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field createdAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"conversationId", "title", "scope", "conversationSummary", "lastCompactedAt", "summaryVersion", "messages", "runs", "createdAt"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ConversationId = *all.ConversationId
	o.Title = *all.Title
	if all.Scope != nil && all.Scope.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Scope = all.Scope
	o.ConversationSummary = all.ConversationSummary
	o.LastCompactedAt = all.LastCompactedAt
	o.SummaryVersion = all.SummaryVersion
	o.Messages = all.Messages
	o.Runs = all.Runs
	o.CreatedAt = *all.CreatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
