// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentTurnAction struct {
	ActionId       string           `json:"actionId"`
	ConversationId string           `json:"conversationId"`
	TurnId         string           `json:"turnId"`
	MessageId      *string          `json:"messageId,omitempty"`
	Type           AiAgentEventType `json:"type"`
	Status         *string          `json:"status,omitempty"`
	Title          string           `json:"title"`
	Summary        *string          `json:"summary,omitempty"`
	ToolName       *string          `json:"toolName,omitempty"`
	Target         *string          `json:"target,omitempty"`
	ConfirmationId *string          `json:"confirmationId,omitempty"`
	HasDetail      bool             `json:"hasDetail"`
	// Sanitized action detail. Warning states are returned as `warning.value`
	// with stable values such as `output_truncated`, `final_analysis_missing`,
	// `no_tool_evidence`, `model_no_tool_call`, or `tooling_unavailable`.
	// `model_no_tool_call` and `tooling_unavailable` are turn-completed
	// warnings and are not emitted on tool/check detail cards.
	// The legacy `warningReason` string may be present for compatibility.
	//
	Detail    map[string]interface{} `json:"detail,omitempty"`
	CreatedAt time.Time              `json:"createdAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentTurnAction instantiates a new AiAgentTurnAction object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentTurnAction(actionId string, conversationId string, turnId string, typeVar AiAgentEventType, title string, hasDetail bool, createdAt time.Time) *AiAgentTurnAction {
	this := AiAgentTurnAction{}
	this.ActionId = actionId
	this.ConversationId = conversationId
	this.TurnId = turnId
	this.Type = typeVar
	this.Title = title
	this.HasDetail = hasDetail
	this.CreatedAt = createdAt
	return &this
}

// NewAiAgentTurnActionWithDefaults instantiates a new AiAgentTurnAction object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentTurnActionWithDefaults() *AiAgentTurnAction {
	this := AiAgentTurnAction{}
	return &this
}

// GetActionId returns the ActionId field value.
func (o *AiAgentTurnAction) GetActionId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ActionId
}

// GetActionIdOk returns a tuple with the ActionId field value
// and a boolean to check if the value has been set.
func (o *AiAgentTurnAction) GetActionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionId, true
}

// SetActionId sets field value.
func (o *AiAgentTurnAction) SetActionId(v string) {
	o.ActionId = v
}

// GetConversationId returns the ConversationId field value.
func (o *AiAgentTurnAction) GetConversationId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ConversationId
}

// GetConversationIdOk returns a tuple with the ConversationId field value
// and a boolean to check if the value has been set.
func (o *AiAgentTurnAction) GetConversationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConversationId, true
}

// SetConversationId sets field value.
func (o *AiAgentTurnAction) SetConversationId(v string) {
	o.ConversationId = v
}

// GetTurnId returns the TurnId field value.
func (o *AiAgentTurnAction) GetTurnId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TurnId
}

// GetTurnIdOk returns a tuple with the TurnId field value
// and a boolean to check if the value has been set.
func (o *AiAgentTurnAction) GetTurnIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TurnId, true
}

// SetTurnId sets field value.
func (o *AiAgentTurnAction) SetTurnId(v string) {
	o.TurnId = v
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *AiAgentTurnAction) GetMessageId() string {
	if o == nil || o.MessageId == nil {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentTurnAction) GetMessageIdOk() (*string, bool) {
	if o == nil || o.MessageId == nil {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *AiAgentTurnAction) HasMessageId() bool {
	return o != nil && o.MessageId != nil
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *AiAgentTurnAction) SetMessageId(v string) {
	o.MessageId = &v
}

// GetType returns the Type field value.
func (o *AiAgentTurnAction) GetType() AiAgentEventType {
	if o == nil {
		var ret AiAgentEventType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AiAgentTurnAction) GetTypeOk() (*AiAgentEventType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AiAgentTurnAction) SetType(v AiAgentEventType) {
	o.Type = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AiAgentTurnAction) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentTurnAction) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AiAgentTurnAction) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AiAgentTurnAction) SetStatus(v string) {
	o.Status = &v
}

// GetTitle returns the Title field value.
func (o *AiAgentTurnAction) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *AiAgentTurnAction) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *AiAgentTurnAction) SetTitle(v string) {
	o.Title = v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *AiAgentTurnAction) GetSummary() string {
	if o == nil || o.Summary == nil {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentTurnAction) GetSummaryOk() (*string, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *AiAgentTurnAction) HasSummary() bool {
	return o != nil && o.Summary != nil
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *AiAgentTurnAction) SetSummary(v string) {
	o.Summary = &v
}

// GetToolName returns the ToolName field value if set, zero value otherwise.
func (o *AiAgentTurnAction) GetToolName() string {
	if o == nil || o.ToolName == nil {
		var ret string
		return ret
	}
	return *o.ToolName
}

// GetToolNameOk returns a tuple with the ToolName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentTurnAction) GetToolNameOk() (*string, bool) {
	if o == nil || o.ToolName == nil {
		return nil, false
	}
	return o.ToolName, true
}

// HasToolName returns a boolean if a field has been set.
func (o *AiAgentTurnAction) HasToolName() bool {
	return o != nil && o.ToolName != nil
}

// SetToolName gets a reference to the given string and assigns it to the ToolName field.
func (o *AiAgentTurnAction) SetToolName(v string) {
	o.ToolName = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *AiAgentTurnAction) GetTarget() string {
	if o == nil || o.Target == nil {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentTurnAction) GetTargetOk() (*string, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *AiAgentTurnAction) HasTarget() bool {
	return o != nil && o.Target != nil
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *AiAgentTurnAction) SetTarget(v string) {
	o.Target = &v
}

// GetConfirmationId returns the ConfirmationId field value if set, zero value otherwise.
func (o *AiAgentTurnAction) GetConfirmationId() string {
	if o == nil || o.ConfirmationId == nil {
		var ret string
		return ret
	}
	return *o.ConfirmationId
}

// GetConfirmationIdOk returns a tuple with the ConfirmationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentTurnAction) GetConfirmationIdOk() (*string, bool) {
	if o == nil || o.ConfirmationId == nil {
		return nil, false
	}
	return o.ConfirmationId, true
}

// HasConfirmationId returns a boolean if a field has been set.
func (o *AiAgentTurnAction) HasConfirmationId() bool {
	return o != nil && o.ConfirmationId != nil
}

// SetConfirmationId gets a reference to the given string and assigns it to the ConfirmationId field.
func (o *AiAgentTurnAction) SetConfirmationId(v string) {
	o.ConfirmationId = &v
}

// GetHasDetail returns the HasDetail field value.
func (o *AiAgentTurnAction) GetHasDetail() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.HasDetail
}

// GetHasDetailOk returns a tuple with the HasDetail field value
// and a boolean to check if the value has been set.
func (o *AiAgentTurnAction) GetHasDetailOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasDetail, true
}

// SetHasDetail sets field value.
func (o *AiAgentTurnAction) SetHasDetail(v bool) {
	o.HasDetail = v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *AiAgentTurnAction) GetDetail() map[string]interface{} {
	if o == nil || o.Detail == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentTurnAction) GetDetailOk() (*map[string]interface{}, bool) {
	if o == nil || o.Detail == nil {
		return nil, false
	}
	return &o.Detail, true
}

// HasDetailSet returns a boolean if a field has been set.
func (o *AiAgentTurnAction) HasDetailSet() bool {
	return o != nil && o.Detail != nil
}

// SetDetail gets a reference to the given map[string]interface{} and assigns it to the Detail field.
func (o *AiAgentTurnAction) SetDetail(v map[string]interface{}) {
	o.Detail = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *AiAgentTurnAction) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AiAgentTurnAction) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *AiAgentTurnAction) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentTurnAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["actionId"] = o.ActionId
	toSerialize["conversationId"] = o.ConversationId
	toSerialize["turnId"] = o.TurnId
	if o.MessageId != nil {
		toSerialize["messageId"] = o.MessageId
	}
	toSerialize["type"] = o.Type
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	toSerialize["title"] = o.Title
	if o.Summary != nil {
		toSerialize["summary"] = o.Summary
	}
	if o.ToolName != nil {
		toSerialize["toolName"] = o.ToolName
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	if o.ConfirmationId != nil {
		toSerialize["confirmationId"] = o.ConfirmationId
	}
	toSerialize["hasDetail"] = o.HasDetail
	if o.Detail != nil {
		toSerialize["detail"] = o.Detail
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
func (o *AiAgentTurnAction) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ActionId       *string                `json:"actionId"`
		ConversationId *string                `json:"conversationId"`
		TurnId         *string                `json:"turnId"`
		MessageId      *string                `json:"messageId,omitempty"`
		Type           *AiAgentEventType      `json:"type"`
		Status         *string                `json:"status,omitempty"`
		Title          *string                `json:"title"`
		Summary        *string                `json:"summary,omitempty"`
		ToolName       *string                `json:"toolName,omitempty"`
		Target         *string                `json:"target,omitempty"`
		ConfirmationId *string                `json:"confirmationId,omitempty"`
		HasDetail      *bool                  `json:"hasDetail"`
		Detail         map[string]interface{} `json:"detail,omitempty"`
		CreatedAt      *time.Time             `json:"createdAt"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.ActionId == nil {
		return fmt.Errorf("required field actionId missing")
	}
	if all.ConversationId == nil {
		return fmt.Errorf("required field conversationId missing")
	}
	if all.TurnId == nil {
		return fmt.Errorf("required field turnId missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	if all.HasDetail == nil {
		return fmt.Errorf("required field hasDetail missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field createdAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"actionId", "conversationId", "turnId", "messageId", "type", "status", "title", "summary", "toolName", "target", "confirmationId", "hasDetail", "detail", "createdAt"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ActionId = *all.ActionId
	o.ConversationId = *all.ConversationId
	o.TurnId = *all.TurnId
	o.MessageId = all.MessageId
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.Status = all.Status
	o.Title = *all.Title
	o.Summary = all.Summary
	o.ToolName = all.ToolName
	o.Target = all.Target
	o.ConfirmationId = all.ConfirmationId
	o.HasDetail = *all.HasDetail
	o.Detail = all.Detail
	o.CreatedAt = *all.CreatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
