// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentConversation struct {
	ConversationId   string                    `json:"conversationId"`
	EntryClusterName *string                   `json:"entryClusterName,omitempty"`
	Title            *string                   `json:"title,omitempty"`
	Status           AiAgentConversationStatus `json:"status"`
	RuntimeStatus    *AiAgentRuntimeStatusCode `json:"runtimeStatus,omitempty"`
	Model            *string                   `json:"model,omitempty"`
	CreatedAt        time.Time                 `json:"createdAt"`
	UpdatedAt        time.Time                 `json:"updatedAt"`
	LastMessage      *AiAgentMessage           `json:"lastMessage,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentConversation instantiates a new AiAgentConversation object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentConversation(conversationId string, status AiAgentConversationStatus, createdAt time.Time, updatedAt time.Time) *AiAgentConversation {
	this := AiAgentConversation{}
	this.ConversationId = conversationId
	this.Status = status
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
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

// GetEntryClusterName returns the EntryClusterName field value if set, zero value otherwise.
func (o *AiAgentConversation) GetEntryClusterName() string {
	if o == nil || o.EntryClusterName == nil {
		var ret string
		return ret
	}
	return *o.EntryClusterName
}

// GetEntryClusterNameOk returns a tuple with the EntryClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentConversation) GetEntryClusterNameOk() (*string, bool) {
	if o == nil || o.EntryClusterName == nil {
		return nil, false
	}
	return o.EntryClusterName, true
}

// HasEntryClusterName returns a boolean if a field has been set.
func (o *AiAgentConversation) HasEntryClusterName() bool {
	return o != nil && o.EntryClusterName != nil
}

// SetEntryClusterName gets a reference to the given string and assigns it to the EntryClusterName field.
func (o *AiAgentConversation) SetEntryClusterName(v string) {
	o.EntryClusterName = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AiAgentConversation) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentConversation) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AiAgentConversation) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AiAgentConversation) SetTitle(v string) {
	o.Title = &v
}

// GetStatus returns the Status field value.
func (o *AiAgentConversation) GetStatus() AiAgentConversationStatus {
	if o == nil {
		var ret AiAgentConversationStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AiAgentConversation) GetStatusOk() (*AiAgentConversationStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *AiAgentConversation) SetStatus(v AiAgentConversationStatus) {
	o.Status = v
}

// GetRuntimeStatus returns the RuntimeStatus field value if set, zero value otherwise.
func (o *AiAgentConversation) GetRuntimeStatus() AiAgentRuntimeStatusCode {
	if o == nil || o.RuntimeStatus == nil {
		var ret AiAgentRuntimeStatusCode
		return ret
	}
	return *o.RuntimeStatus
}

// GetRuntimeStatusOk returns a tuple with the RuntimeStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentConversation) GetRuntimeStatusOk() (*AiAgentRuntimeStatusCode, bool) {
	if o == nil || o.RuntimeStatus == nil {
		return nil, false
	}
	return o.RuntimeStatus, true
}

// HasRuntimeStatus returns a boolean if a field has been set.
func (o *AiAgentConversation) HasRuntimeStatus() bool {
	return o != nil && o.RuntimeStatus != nil
}

// SetRuntimeStatus gets a reference to the given AiAgentRuntimeStatusCode and assigns it to the RuntimeStatus field.
func (o *AiAgentConversation) SetRuntimeStatus(v AiAgentRuntimeStatusCode) {
	o.RuntimeStatus = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *AiAgentConversation) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentConversation) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *AiAgentConversation) HasModel() bool {
	return o != nil && o.Model != nil
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *AiAgentConversation) SetModel(v string) {
	o.Model = &v
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

// GetUpdatedAt returns the UpdatedAt field value.
func (o *AiAgentConversation) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *AiAgentConversation) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *AiAgentConversation) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetLastMessage returns the LastMessage field value if set, zero value otherwise.
func (o *AiAgentConversation) GetLastMessage() AiAgentMessage {
	if o == nil || o.LastMessage == nil {
		var ret AiAgentMessage
		return ret
	}
	return *o.LastMessage
}

// GetLastMessageOk returns a tuple with the LastMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentConversation) GetLastMessageOk() (*AiAgentMessage, bool) {
	if o == nil || o.LastMessage == nil {
		return nil, false
	}
	return o.LastMessage, true
}

// HasLastMessage returns a boolean if a field has been set.
func (o *AiAgentConversation) HasLastMessage() bool {
	return o != nil && o.LastMessage != nil
}

// SetLastMessage gets a reference to the given AiAgentMessage and assigns it to the LastMessage field.
func (o *AiAgentConversation) SetLastMessage(v AiAgentMessage) {
	o.LastMessage = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentConversation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["conversationId"] = o.ConversationId
	if o.EntryClusterName != nil {
		toSerialize["entryClusterName"] = o.EntryClusterName
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	toSerialize["status"] = o.Status
	if o.RuntimeStatus != nil {
		toSerialize["runtimeStatus"] = o.RuntimeStatus
	}
	if o.Model != nil {
		toSerialize["model"] = o.Model
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.UpdatedAt.Nanosecond() == 0 {
		toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.LastMessage != nil {
		toSerialize["lastMessage"] = o.LastMessage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentConversation) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConversationId   *string                    `json:"conversationId"`
		EntryClusterName *string                    `json:"entryClusterName,omitempty"`
		Title            *string                    `json:"title,omitempty"`
		Status           *AiAgentConversationStatus `json:"status"`
		RuntimeStatus    *AiAgentRuntimeStatusCode  `json:"runtimeStatus,omitempty"`
		Model            *string                    `json:"model,omitempty"`
		CreatedAt        *time.Time                 `json:"createdAt"`
		UpdatedAt        *time.Time                 `json:"updatedAt"`
		LastMessage      *AiAgentMessage            `json:"lastMessage,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.ConversationId == nil {
		return fmt.Errorf("required field conversationId missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field createdAt missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updatedAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"conversationId", "entryClusterName", "title", "status", "runtimeStatus", "model", "createdAt", "updatedAt", "lastMessage"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ConversationId = *all.ConversationId
	o.EntryClusterName = all.EntryClusterName
	o.Title = all.Title
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	if all.RuntimeStatus != nil && !all.RuntimeStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.RuntimeStatus = all.RuntimeStatus
	}
	o.Model = all.Model
	o.CreatedAt = *all.CreatedAt
	o.UpdatedAt = *all.UpdatedAt
	if all.LastMessage != nil && all.LastMessage.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LastMessage = all.LastMessage

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
