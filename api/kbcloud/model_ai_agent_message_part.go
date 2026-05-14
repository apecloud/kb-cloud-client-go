// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentMessagePart struct {
	PartId                string                  `json:"partId"`
	Type                  AiAgentMessagePartType  `json:"type"`
	Text                  *string                 `json:"text,omitempty"`
	ReportId              *string                 `json:"reportId,omitempty"`
	EvidenceId            *string                 `json:"evidenceId,omitempty"`
	ActionPlanId          *string                 `json:"actionPlanId,omitempty"`
	RunId                 *string                 `json:"runId,omitempty"`
	ErrorCode             *string                 `json:"errorCode,omitempty"`
	DisplayTitle          *string                 `json:"displayTitle,omitempty"`
	Question              *string                 `json:"question,omitempty"`
	RequiredContextFields []string                `json:"requiredContextFields,omitempty"`
	SuggestedReplies      []AiAgentSuggestedReply `json:"suggestedReplies,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentMessagePart instantiates a new AiAgentMessagePart object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentMessagePart(partId string, typeVar AiAgentMessagePartType) *AiAgentMessagePart {
	this := AiAgentMessagePart{}
	this.PartId = partId
	this.Type = typeVar
	return &this
}

// NewAiAgentMessagePartWithDefaults instantiates a new AiAgentMessagePart object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentMessagePartWithDefaults() *AiAgentMessagePart {
	this := AiAgentMessagePart{}
	return &this
}

// GetPartId returns the PartId field value.
func (o *AiAgentMessagePart) GetPartId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PartId
}

// GetPartIdOk returns a tuple with the PartId field value
// and a boolean to check if the value has been set.
func (o *AiAgentMessagePart) GetPartIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PartId, true
}

// SetPartId sets field value.
func (o *AiAgentMessagePart) SetPartId(v string) {
	o.PartId = v
}

// GetType returns the Type field value.
func (o *AiAgentMessagePart) GetType() AiAgentMessagePartType {
	if o == nil {
		var ret AiAgentMessagePartType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AiAgentMessagePart) GetTypeOk() (*AiAgentMessagePartType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AiAgentMessagePart) SetType(v AiAgentMessagePartType) {
	o.Type = v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *AiAgentMessagePart) GetText() string {
	if o == nil || o.Text == nil {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentMessagePart) GetTextOk() (*string, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *AiAgentMessagePart) HasText() bool {
	return o != nil && o.Text != nil
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *AiAgentMessagePart) SetText(v string) {
	o.Text = &v
}

// GetReportId returns the ReportId field value if set, zero value otherwise.
func (o *AiAgentMessagePart) GetReportId() string {
	if o == nil || o.ReportId == nil {
		var ret string
		return ret
	}
	return *o.ReportId
}

// GetReportIdOk returns a tuple with the ReportId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentMessagePart) GetReportIdOk() (*string, bool) {
	if o == nil || o.ReportId == nil {
		return nil, false
	}
	return o.ReportId, true
}

// HasReportId returns a boolean if a field has been set.
func (o *AiAgentMessagePart) HasReportId() bool {
	return o != nil && o.ReportId != nil
}

// SetReportId gets a reference to the given string and assigns it to the ReportId field.
func (o *AiAgentMessagePart) SetReportId(v string) {
	o.ReportId = &v
}

// GetEvidenceId returns the EvidenceId field value if set, zero value otherwise.
func (o *AiAgentMessagePart) GetEvidenceId() string {
	if o == nil || o.EvidenceId == nil {
		var ret string
		return ret
	}
	return *o.EvidenceId
}

// GetEvidenceIdOk returns a tuple with the EvidenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentMessagePart) GetEvidenceIdOk() (*string, bool) {
	if o == nil || o.EvidenceId == nil {
		return nil, false
	}
	return o.EvidenceId, true
}

// HasEvidenceId returns a boolean if a field has been set.
func (o *AiAgentMessagePart) HasEvidenceId() bool {
	return o != nil && o.EvidenceId != nil
}

// SetEvidenceId gets a reference to the given string and assigns it to the EvidenceId field.
func (o *AiAgentMessagePart) SetEvidenceId(v string) {
	o.EvidenceId = &v
}

// GetActionPlanId returns the ActionPlanId field value if set, zero value otherwise.
func (o *AiAgentMessagePart) GetActionPlanId() string {
	if o == nil || o.ActionPlanId == nil {
		var ret string
		return ret
	}
	return *o.ActionPlanId
}

// GetActionPlanIdOk returns a tuple with the ActionPlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentMessagePart) GetActionPlanIdOk() (*string, bool) {
	if o == nil || o.ActionPlanId == nil {
		return nil, false
	}
	return o.ActionPlanId, true
}

// HasActionPlanId returns a boolean if a field has been set.
func (o *AiAgentMessagePart) HasActionPlanId() bool {
	return o != nil && o.ActionPlanId != nil
}

// SetActionPlanId gets a reference to the given string and assigns it to the ActionPlanId field.
func (o *AiAgentMessagePart) SetActionPlanId(v string) {
	o.ActionPlanId = &v
}

// GetRunId returns the RunId field value if set, zero value otherwise.
func (o *AiAgentMessagePart) GetRunId() string {
	if o == nil || o.RunId == nil {
		var ret string
		return ret
	}
	return *o.RunId
}

// GetRunIdOk returns a tuple with the RunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentMessagePart) GetRunIdOk() (*string, bool) {
	if o == nil || o.RunId == nil {
		return nil, false
	}
	return o.RunId, true
}

// HasRunId returns a boolean if a field has been set.
func (o *AiAgentMessagePart) HasRunId() bool {
	return o != nil && o.RunId != nil
}

// SetRunId gets a reference to the given string and assigns it to the RunId field.
func (o *AiAgentMessagePart) SetRunId(v string) {
	o.RunId = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *AiAgentMessagePart) GetErrorCode() string {
	if o == nil || o.ErrorCode == nil {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentMessagePart) GetErrorCodeOk() (*string, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *AiAgentMessagePart) HasErrorCode() bool {
	return o != nil && o.ErrorCode != nil
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *AiAgentMessagePart) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetDisplayTitle returns the DisplayTitle field value if set, zero value otherwise.
func (o *AiAgentMessagePart) GetDisplayTitle() string {
	if o == nil || o.DisplayTitle == nil {
		var ret string
		return ret
	}
	return *o.DisplayTitle
}

// GetDisplayTitleOk returns a tuple with the DisplayTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentMessagePart) GetDisplayTitleOk() (*string, bool) {
	if o == nil || o.DisplayTitle == nil {
		return nil, false
	}
	return o.DisplayTitle, true
}

// HasDisplayTitle returns a boolean if a field has been set.
func (o *AiAgentMessagePart) HasDisplayTitle() bool {
	return o != nil && o.DisplayTitle != nil
}

// SetDisplayTitle gets a reference to the given string and assigns it to the DisplayTitle field.
func (o *AiAgentMessagePart) SetDisplayTitle(v string) {
	o.DisplayTitle = &v
}

// GetQuestion returns the Question field value if set, zero value otherwise.
func (o *AiAgentMessagePart) GetQuestion() string {
	if o == nil || o.Question == nil {
		var ret string
		return ret
	}
	return *o.Question
}

// GetQuestionOk returns a tuple with the Question field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentMessagePart) GetQuestionOk() (*string, bool) {
	if o == nil || o.Question == nil {
		return nil, false
	}
	return o.Question, true
}

// HasQuestion returns a boolean if a field has been set.
func (o *AiAgentMessagePart) HasQuestion() bool {
	return o != nil && o.Question != nil
}

// SetQuestion gets a reference to the given string and assigns it to the Question field.
func (o *AiAgentMessagePart) SetQuestion(v string) {
	o.Question = &v
}

// GetRequiredContextFields returns the RequiredContextFields field value if set, zero value otherwise.
func (o *AiAgentMessagePart) GetRequiredContextFields() []string {
	if o == nil || o.RequiredContextFields == nil {
		var ret []string
		return ret
	}
	return o.RequiredContextFields
}

// GetRequiredContextFieldsOk returns a tuple with the RequiredContextFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentMessagePart) GetRequiredContextFieldsOk() (*[]string, bool) {
	if o == nil || o.RequiredContextFields == nil {
		return nil, false
	}
	return &o.RequiredContextFields, true
}

// HasRequiredContextFields returns a boolean if a field has been set.
func (o *AiAgentMessagePart) HasRequiredContextFields() bool {
	return o != nil && o.RequiredContextFields != nil
}

// SetRequiredContextFields gets a reference to the given []string and assigns it to the RequiredContextFields field.
func (o *AiAgentMessagePart) SetRequiredContextFields(v []string) {
	o.RequiredContextFields = v
}

// GetSuggestedReplies returns the SuggestedReplies field value if set, zero value otherwise.
func (o *AiAgentMessagePart) GetSuggestedReplies() []AiAgentSuggestedReply {
	if o == nil || o.SuggestedReplies == nil {
		var ret []AiAgentSuggestedReply
		return ret
	}
	return o.SuggestedReplies
}

// GetSuggestedRepliesOk returns a tuple with the SuggestedReplies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentMessagePart) GetSuggestedRepliesOk() (*[]AiAgentSuggestedReply, bool) {
	if o == nil || o.SuggestedReplies == nil {
		return nil, false
	}
	return &o.SuggestedReplies, true
}

// HasSuggestedReplies returns a boolean if a field has been set.
func (o *AiAgentMessagePart) HasSuggestedReplies() bool {
	return o != nil && o.SuggestedReplies != nil
}

// SetSuggestedReplies gets a reference to the given []AiAgentSuggestedReply and assigns it to the SuggestedReplies field.
func (o *AiAgentMessagePart) SetSuggestedReplies(v []AiAgentSuggestedReply) {
	o.SuggestedReplies = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentMessagePart) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["partId"] = o.PartId
	toSerialize["type"] = o.Type
	if o.Text != nil {
		toSerialize["text"] = o.Text
	}
	if o.ReportId != nil {
		toSerialize["reportId"] = o.ReportId
	}
	if o.EvidenceId != nil {
		toSerialize["evidenceId"] = o.EvidenceId
	}
	if o.ActionPlanId != nil {
		toSerialize["actionPlanId"] = o.ActionPlanId
	}
	if o.RunId != nil {
		toSerialize["runId"] = o.RunId
	}
	if o.ErrorCode != nil {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if o.DisplayTitle != nil {
		toSerialize["displayTitle"] = o.DisplayTitle
	}
	if o.Question != nil {
		toSerialize["question"] = o.Question
	}
	if o.RequiredContextFields != nil {
		toSerialize["requiredContextFields"] = o.RequiredContextFields
	}
	if o.SuggestedReplies != nil {
		toSerialize["suggestedReplies"] = o.SuggestedReplies
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentMessagePart) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		PartId                *string                 `json:"partId"`
		Type                  *AiAgentMessagePartType `json:"type"`
		Text                  *string                 `json:"text,omitempty"`
		ReportId              *string                 `json:"reportId,omitempty"`
		EvidenceId            *string                 `json:"evidenceId,omitempty"`
		ActionPlanId          *string                 `json:"actionPlanId,omitempty"`
		RunId                 *string                 `json:"runId,omitempty"`
		ErrorCode             *string                 `json:"errorCode,omitempty"`
		DisplayTitle          *string                 `json:"displayTitle,omitempty"`
		Question              *string                 `json:"question,omitempty"`
		RequiredContextFields []string                `json:"requiredContextFields,omitempty"`
		SuggestedReplies      []AiAgentSuggestedReply `json:"suggestedReplies,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.PartId == nil {
		return fmt.Errorf("required field partId missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"partId", "type", "text", "reportId", "evidenceId", "actionPlanId", "runId", "errorCode", "displayTitle", "question", "requiredContextFields", "suggestedReplies"})
	} else {
		return err
	}

	hasInvalidField := false
	o.PartId = *all.PartId
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.Text = all.Text
	o.ReportId = all.ReportId
	o.EvidenceId = all.EvidenceId
	o.ActionPlanId = all.ActionPlanId
	o.RunId = all.RunId
	o.ErrorCode = all.ErrorCode
	o.DisplayTitle = all.DisplayTitle
	o.Question = all.Question
	o.RequiredContextFields = all.RequiredContextFields
	o.SuggestedReplies = all.SuggestedReplies

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
