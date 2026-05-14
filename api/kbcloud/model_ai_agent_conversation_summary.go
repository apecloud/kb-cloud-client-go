// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentConversationSummary struct {
	ConversationId string `json:"conversationId"`
	Title          string `json:"title"`
	// Organization-scoped single-cluster Agent execution scope.
	Scope       *AiAgentResourceScope     `json:"scope,omitempty"`
	Status      AiAgentConversationStatus `json:"status"`
	LastSummary *string                   `json:"lastSummary,omitempty"`
	LastRunId   *string                   `json:"lastRunId,omitempty"`
	UpdatedAt   time.Time                 `json:"updatedAt"`
	CreatedAt   time.Time                 `json:"createdAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentConversationSummary instantiates a new AiAgentConversationSummary object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentConversationSummary(conversationId string, title string, status AiAgentConversationStatus, updatedAt time.Time, createdAt time.Time) *AiAgentConversationSummary {
	this := AiAgentConversationSummary{}
	this.ConversationId = conversationId
	this.Title = title
	this.Status = status
	this.UpdatedAt = updatedAt
	this.CreatedAt = createdAt
	return &this
}

// NewAiAgentConversationSummaryWithDefaults instantiates a new AiAgentConversationSummary object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentConversationSummaryWithDefaults() *AiAgentConversationSummary {
	this := AiAgentConversationSummary{}
	return &this
}

// GetConversationId returns the ConversationId field value.
func (o *AiAgentConversationSummary) GetConversationId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ConversationId
}

// GetConversationIdOk returns a tuple with the ConversationId field value
// and a boolean to check if the value has been set.
func (o *AiAgentConversationSummary) GetConversationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConversationId, true
}

// SetConversationId sets field value.
func (o *AiAgentConversationSummary) SetConversationId(v string) {
	o.ConversationId = v
}

// GetTitle returns the Title field value.
func (o *AiAgentConversationSummary) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *AiAgentConversationSummary) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *AiAgentConversationSummary) SetTitle(v string) {
	o.Title = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *AiAgentConversationSummary) GetScope() AiAgentResourceScope {
	if o == nil || o.Scope == nil {
		var ret AiAgentResourceScope
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentConversationSummary) GetScopeOk() (*AiAgentResourceScope, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *AiAgentConversationSummary) HasScope() bool {
	return o != nil && o.Scope != nil
}

// SetScope gets a reference to the given AiAgentResourceScope and assigns it to the Scope field.
func (o *AiAgentConversationSummary) SetScope(v AiAgentResourceScope) {
	o.Scope = &v
}

// GetStatus returns the Status field value.
func (o *AiAgentConversationSummary) GetStatus() AiAgentConversationStatus {
	if o == nil {
		var ret AiAgentConversationStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AiAgentConversationSummary) GetStatusOk() (*AiAgentConversationStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *AiAgentConversationSummary) SetStatus(v AiAgentConversationStatus) {
	o.Status = v
}

// GetLastSummary returns the LastSummary field value if set, zero value otherwise.
func (o *AiAgentConversationSummary) GetLastSummary() string {
	if o == nil || o.LastSummary == nil {
		var ret string
		return ret
	}
	return *o.LastSummary
}

// GetLastSummaryOk returns a tuple with the LastSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentConversationSummary) GetLastSummaryOk() (*string, bool) {
	if o == nil || o.LastSummary == nil {
		return nil, false
	}
	return o.LastSummary, true
}

// HasLastSummary returns a boolean if a field has been set.
func (o *AiAgentConversationSummary) HasLastSummary() bool {
	return o != nil && o.LastSummary != nil
}

// SetLastSummary gets a reference to the given string and assigns it to the LastSummary field.
func (o *AiAgentConversationSummary) SetLastSummary(v string) {
	o.LastSummary = &v
}

// GetLastRunId returns the LastRunId field value if set, zero value otherwise.
func (o *AiAgentConversationSummary) GetLastRunId() string {
	if o == nil || o.LastRunId == nil {
		var ret string
		return ret
	}
	return *o.LastRunId
}

// GetLastRunIdOk returns a tuple with the LastRunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentConversationSummary) GetLastRunIdOk() (*string, bool) {
	if o == nil || o.LastRunId == nil {
		return nil, false
	}
	return o.LastRunId, true
}

// HasLastRunId returns a boolean if a field has been set.
func (o *AiAgentConversationSummary) HasLastRunId() bool {
	return o != nil && o.LastRunId != nil
}

// SetLastRunId gets a reference to the given string and assigns it to the LastRunId field.
func (o *AiAgentConversationSummary) SetLastRunId(v string) {
	o.LastRunId = &v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *AiAgentConversationSummary) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *AiAgentConversationSummary) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *AiAgentConversationSummary) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *AiAgentConversationSummary) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AiAgentConversationSummary) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *AiAgentConversationSummary) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentConversationSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["conversationId"] = o.ConversationId
	toSerialize["title"] = o.Title
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	toSerialize["status"] = o.Status
	if o.LastSummary != nil {
		toSerialize["lastSummary"] = o.LastSummary
	}
	if o.LastRunId != nil {
		toSerialize["lastRunId"] = o.LastRunId
	}
	if o.UpdatedAt.Nanosecond() == 0 {
		toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
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
func (o *AiAgentConversationSummary) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConversationId *string                    `json:"conversationId"`
		Title          *string                    `json:"title"`
		Scope          *AiAgentResourceScope      `json:"scope,omitempty"`
		Status         *AiAgentConversationStatus `json:"status"`
		LastSummary    *string                    `json:"lastSummary,omitempty"`
		LastRunId      *string                    `json:"lastRunId,omitempty"`
		UpdatedAt      *time.Time                 `json:"updatedAt"`
		CreatedAt      *time.Time                 `json:"createdAt"`
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
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updatedAt missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field createdAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"conversationId", "title", "scope", "status", "lastSummary", "lastRunId", "updatedAt", "createdAt"})
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
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.LastSummary = all.LastSummary
	o.LastRunId = all.LastRunId
	o.UpdatedAt = *all.UpdatedAt
	o.CreatedAt = *all.CreatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
