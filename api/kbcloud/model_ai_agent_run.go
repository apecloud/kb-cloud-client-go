// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentRun struct {
	RunId          string           `json:"runId"`
	ConversationId string           `json:"conversationId"`
	Status         AiAgentRunStatus `json:"status"`
	// Stable status code plus default user-facing copy. Frontend uses code for logic/tests and may override display text for i18n.
	StatusDisplay *AiAgentStatusDisplay `json:"statusDisplay,omitempty"`
	FinalStatus   *AiAgentFinalStatus   `json:"finalStatus,omitempty"`
	// Organization-scoped single-cluster Agent execution scope.
	Scope           *AiAgentResourceScope   `json:"scope,omitempty"`
	PlaybookId      *string                 `json:"playbookId,omitempty"`
	CurrentStage    *AiAgentStage           `json:"currentStage,omitempty"`
	DiagnosisBranch *AiAgentDiagnosisBranch `json:"diagnosisBranch,omitempty"`
	ActionBranch    *AiAgentActionBranch    `json:"actionBranch,omitempty"`
	RetryOfRunId    *string                 `json:"retryOfRunId,omitempty"`
	ForkedFromRunId *string                 `json:"forkedFromRunId,omitempty"`
	EventsUrl       *string                 `json:"eventsUrl,omitempty"`
	Metadata        *AiAgentRunMetadata     `json:"metadata,omitempty"`
	CreatedAt       time.Time               `json:"createdAt"`
	UpdatedAt       *time.Time              `json:"updatedAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentRun instantiates a new AiAgentRun object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentRun(runId string, conversationId string, status AiAgentRunStatus, createdAt time.Time) *AiAgentRun {
	this := AiAgentRun{}
	this.RunId = runId
	this.ConversationId = conversationId
	this.Status = status
	this.CreatedAt = createdAt
	return &this
}

// NewAiAgentRunWithDefaults instantiates a new AiAgentRun object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentRunWithDefaults() *AiAgentRun {
	this := AiAgentRun{}
	return &this
}

// GetRunId returns the RunId field value.
func (o *AiAgentRun) GetRunId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RunId
}

// GetRunIdOk returns a tuple with the RunId field value
// and a boolean to check if the value has been set.
func (o *AiAgentRun) GetRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunId, true
}

// SetRunId sets field value.
func (o *AiAgentRun) SetRunId(v string) {
	o.RunId = v
}

// GetConversationId returns the ConversationId field value.
func (o *AiAgentRun) GetConversationId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ConversationId
}

// GetConversationIdOk returns a tuple with the ConversationId field value
// and a boolean to check if the value has been set.
func (o *AiAgentRun) GetConversationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConversationId, true
}

// SetConversationId sets field value.
func (o *AiAgentRun) SetConversationId(v string) {
	o.ConversationId = v
}

// GetStatus returns the Status field value.
func (o *AiAgentRun) GetStatus() AiAgentRunStatus {
	if o == nil {
		var ret AiAgentRunStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AiAgentRun) GetStatusOk() (*AiAgentRunStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *AiAgentRun) SetStatus(v AiAgentRunStatus) {
	o.Status = v
}

// GetStatusDisplay returns the StatusDisplay field value if set, zero value otherwise.
func (o *AiAgentRun) GetStatusDisplay() AiAgentStatusDisplay {
	if o == nil || o.StatusDisplay == nil {
		var ret AiAgentStatusDisplay
		return ret
	}
	return *o.StatusDisplay
}

// GetStatusDisplayOk returns a tuple with the StatusDisplay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentRun) GetStatusDisplayOk() (*AiAgentStatusDisplay, bool) {
	if o == nil || o.StatusDisplay == nil {
		return nil, false
	}
	return o.StatusDisplay, true
}

// HasStatusDisplay returns a boolean if a field has been set.
func (o *AiAgentRun) HasStatusDisplay() bool {
	return o != nil && o.StatusDisplay != nil
}

// SetStatusDisplay gets a reference to the given AiAgentStatusDisplay and assigns it to the StatusDisplay field.
func (o *AiAgentRun) SetStatusDisplay(v AiAgentStatusDisplay) {
	o.StatusDisplay = &v
}

// GetFinalStatus returns the FinalStatus field value if set, zero value otherwise.
func (o *AiAgentRun) GetFinalStatus() AiAgentFinalStatus {
	if o == nil || o.FinalStatus == nil {
		var ret AiAgentFinalStatus
		return ret
	}
	return *o.FinalStatus
}

// GetFinalStatusOk returns a tuple with the FinalStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentRun) GetFinalStatusOk() (*AiAgentFinalStatus, bool) {
	if o == nil || o.FinalStatus == nil {
		return nil, false
	}
	return o.FinalStatus, true
}

// HasFinalStatus returns a boolean if a field has been set.
func (o *AiAgentRun) HasFinalStatus() bool {
	return o != nil && o.FinalStatus != nil
}

// SetFinalStatus gets a reference to the given AiAgentFinalStatus and assigns it to the FinalStatus field.
func (o *AiAgentRun) SetFinalStatus(v AiAgentFinalStatus) {
	o.FinalStatus = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *AiAgentRun) GetScope() AiAgentResourceScope {
	if o == nil || o.Scope == nil {
		var ret AiAgentResourceScope
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentRun) GetScopeOk() (*AiAgentResourceScope, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *AiAgentRun) HasScope() bool {
	return o != nil && o.Scope != nil
}

// SetScope gets a reference to the given AiAgentResourceScope and assigns it to the Scope field.
func (o *AiAgentRun) SetScope(v AiAgentResourceScope) {
	o.Scope = &v
}

// GetPlaybookId returns the PlaybookId field value if set, zero value otherwise.
func (o *AiAgentRun) GetPlaybookId() string {
	if o == nil || o.PlaybookId == nil {
		var ret string
		return ret
	}
	return *o.PlaybookId
}

// GetPlaybookIdOk returns a tuple with the PlaybookId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentRun) GetPlaybookIdOk() (*string, bool) {
	if o == nil || o.PlaybookId == nil {
		return nil, false
	}
	return o.PlaybookId, true
}

// HasPlaybookId returns a boolean if a field has been set.
func (o *AiAgentRun) HasPlaybookId() bool {
	return o != nil && o.PlaybookId != nil
}

// SetPlaybookId gets a reference to the given string and assigns it to the PlaybookId field.
func (o *AiAgentRun) SetPlaybookId(v string) {
	o.PlaybookId = &v
}

// GetCurrentStage returns the CurrentStage field value if set, zero value otherwise.
func (o *AiAgentRun) GetCurrentStage() AiAgentStage {
	if o == nil || o.CurrentStage == nil {
		var ret AiAgentStage
		return ret
	}
	return *o.CurrentStage
}

// GetCurrentStageOk returns a tuple with the CurrentStage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentRun) GetCurrentStageOk() (*AiAgentStage, bool) {
	if o == nil || o.CurrentStage == nil {
		return nil, false
	}
	return o.CurrentStage, true
}

// HasCurrentStage returns a boolean if a field has been set.
func (o *AiAgentRun) HasCurrentStage() bool {
	return o != nil && o.CurrentStage != nil
}

// SetCurrentStage gets a reference to the given AiAgentStage and assigns it to the CurrentStage field.
func (o *AiAgentRun) SetCurrentStage(v AiAgentStage) {
	o.CurrentStage = &v
}

// GetDiagnosisBranch returns the DiagnosisBranch field value if set, zero value otherwise.
func (o *AiAgentRun) GetDiagnosisBranch() AiAgentDiagnosisBranch {
	if o == nil || o.DiagnosisBranch == nil {
		var ret AiAgentDiagnosisBranch
		return ret
	}
	return *o.DiagnosisBranch
}

// GetDiagnosisBranchOk returns a tuple with the DiagnosisBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentRun) GetDiagnosisBranchOk() (*AiAgentDiagnosisBranch, bool) {
	if o == nil || o.DiagnosisBranch == nil {
		return nil, false
	}
	return o.DiagnosisBranch, true
}

// HasDiagnosisBranch returns a boolean if a field has been set.
func (o *AiAgentRun) HasDiagnosisBranch() bool {
	return o != nil && o.DiagnosisBranch != nil
}

// SetDiagnosisBranch gets a reference to the given AiAgentDiagnosisBranch and assigns it to the DiagnosisBranch field.
func (o *AiAgentRun) SetDiagnosisBranch(v AiAgentDiagnosisBranch) {
	o.DiagnosisBranch = &v
}

// GetActionBranch returns the ActionBranch field value if set, zero value otherwise.
func (o *AiAgentRun) GetActionBranch() AiAgentActionBranch {
	if o == nil || o.ActionBranch == nil {
		var ret AiAgentActionBranch
		return ret
	}
	return *o.ActionBranch
}

// GetActionBranchOk returns a tuple with the ActionBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentRun) GetActionBranchOk() (*AiAgentActionBranch, bool) {
	if o == nil || o.ActionBranch == nil {
		return nil, false
	}
	return o.ActionBranch, true
}

// HasActionBranch returns a boolean if a field has been set.
func (o *AiAgentRun) HasActionBranch() bool {
	return o != nil && o.ActionBranch != nil
}

// SetActionBranch gets a reference to the given AiAgentActionBranch and assigns it to the ActionBranch field.
func (o *AiAgentRun) SetActionBranch(v AiAgentActionBranch) {
	o.ActionBranch = &v
}

// GetRetryOfRunId returns the RetryOfRunId field value if set, zero value otherwise.
func (o *AiAgentRun) GetRetryOfRunId() string {
	if o == nil || o.RetryOfRunId == nil {
		var ret string
		return ret
	}
	return *o.RetryOfRunId
}

// GetRetryOfRunIdOk returns a tuple with the RetryOfRunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentRun) GetRetryOfRunIdOk() (*string, bool) {
	if o == nil || o.RetryOfRunId == nil {
		return nil, false
	}
	return o.RetryOfRunId, true
}

// HasRetryOfRunId returns a boolean if a field has been set.
func (o *AiAgentRun) HasRetryOfRunId() bool {
	return o != nil && o.RetryOfRunId != nil
}

// SetRetryOfRunId gets a reference to the given string and assigns it to the RetryOfRunId field.
func (o *AiAgentRun) SetRetryOfRunId(v string) {
	o.RetryOfRunId = &v
}

// GetForkedFromRunId returns the ForkedFromRunId field value if set, zero value otherwise.
func (o *AiAgentRun) GetForkedFromRunId() string {
	if o == nil || o.ForkedFromRunId == nil {
		var ret string
		return ret
	}
	return *o.ForkedFromRunId
}

// GetForkedFromRunIdOk returns a tuple with the ForkedFromRunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentRun) GetForkedFromRunIdOk() (*string, bool) {
	if o == nil || o.ForkedFromRunId == nil {
		return nil, false
	}
	return o.ForkedFromRunId, true
}

// HasForkedFromRunId returns a boolean if a field has been set.
func (o *AiAgentRun) HasForkedFromRunId() bool {
	return o != nil && o.ForkedFromRunId != nil
}

// SetForkedFromRunId gets a reference to the given string and assigns it to the ForkedFromRunId field.
func (o *AiAgentRun) SetForkedFromRunId(v string) {
	o.ForkedFromRunId = &v
}

// GetEventsUrl returns the EventsUrl field value if set, zero value otherwise.
func (o *AiAgentRun) GetEventsUrl() string {
	if o == nil || o.EventsUrl == nil {
		var ret string
		return ret
	}
	return *o.EventsUrl
}

// GetEventsUrlOk returns a tuple with the EventsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentRun) GetEventsUrlOk() (*string, bool) {
	if o == nil || o.EventsUrl == nil {
		return nil, false
	}
	return o.EventsUrl, true
}

// HasEventsUrl returns a boolean if a field has been set.
func (o *AiAgentRun) HasEventsUrl() bool {
	return o != nil && o.EventsUrl != nil
}

// SetEventsUrl gets a reference to the given string and assigns it to the EventsUrl field.
func (o *AiAgentRun) SetEventsUrl(v string) {
	o.EventsUrl = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *AiAgentRun) GetMetadata() AiAgentRunMetadata {
	if o == nil || o.Metadata == nil {
		var ret AiAgentRunMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentRun) GetMetadataOk() (*AiAgentRunMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *AiAgentRun) HasMetadata() bool {
	return o != nil && o.Metadata != nil
}

// SetMetadata gets a reference to the given AiAgentRunMetadata and assigns it to the Metadata field.
func (o *AiAgentRun) SetMetadata(v AiAgentRunMetadata) {
	o.Metadata = &v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *AiAgentRun) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AiAgentRun) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *AiAgentRun) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AiAgentRun) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentRun) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AiAgentRun) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *AiAgentRun) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentRun) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["runId"] = o.RunId
	toSerialize["conversationId"] = o.ConversationId
	toSerialize["status"] = o.Status
	if o.StatusDisplay != nil {
		toSerialize["statusDisplay"] = o.StatusDisplay
	}
	if o.FinalStatus != nil {
		toSerialize["finalStatus"] = o.FinalStatus
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.PlaybookId != nil {
		toSerialize["playbookId"] = o.PlaybookId
	}
	if o.CurrentStage != nil {
		toSerialize["currentStage"] = o.CurrentStage
	}
	if o.DiagnosisBranch != nil {
		toSerialize["diagnosisBranch"] = o.DiagnosisBranch
	}
	if o.ActionBranch != nil {
		toSerialize["actionBranch"] = o.ActionBranch
	}
	if o.RetryOfRunId != nil {
		toSerialize["retryOfRunId"] = o.RetryOfRunId
	}
	if o.ForkedFromRunId != nil {
		toSerialize["forkedFromRunId"] = o.ForkedFromRunId
	}
	if o.EventsUrl != nil {
		toSerialize["eventsUrl"] = o.EventsUrl
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
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
func (o *AiAgentRun) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		RunId           *string                 `json:"runId"`
		ConversationId  *string                 `json:"conversationId"`
		Status          *AiAgentRunStatus       `json:"status"`
		StatusDisplay   *AiAgentStatusDisplay   `json:"statusDisplay,omitempty"`
		FinalStatus     *AiAgentFinalStatus     `json:"finalStatus,omitempty"`
		Scope           *AiAgentResourceScope   `json:"scope,omitempty"`
		PlaybookId      *string                 `json:"playbookId,omitempty"`
		CurrentStage    *AiAgentStage           `json:"currentStage,omitempty"`
		DiagnosisBranch *AiAgentDiagnosisBranch `json:"diagnosisBranch,omitempty"`
		ActionBranch    *AiAgentActionBranch    `json:"actionBranch,omitempty"`
		RetryOfRunId    *string                 `json:"retryOfRunId,omitempty"`
		ForkedFromRunId *string                 `json:"forkedFromRunId,omitempty"`
		EventsUrl       *string                 `json:"eventsUrl,omitempty"`
		Metadata        *AiAgentRunMetadata     `json:"metadata,omitempty"`
		CreatedAt       *time.Time              `json:"createdAt"`
		UpdatedAt       *time.Time              `json:"updatedAt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.RunId == nil {
		return fmt.Errorf("required field runId missing")
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
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"runId", "conversationId", "status", "statusDisplay", "finalStatus", "scope", "playbookId", "currentStage", "diagnosisBranch", "actionBranch", "retryOfRunId", "forkedFromRunId", "eventsUrl", "metadata", "createdAt", "updatedAt"})
	} else {
		return err
	}

	hasInvalidField := false
	o.RunId = *all.RunId
	o.ConversationId = *all.ConversationId
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	if all.StatusDisplay != nil && all.StatusDisplay.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.StatusDisplay = all.StatusDisplay
	if all.FinalStatus != nil && !all.FinalStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.FinalStatus = all.FinalStatus
	}
	if all.Scope != nil && all.Scope.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Scope = all.Scope
	o.PlaybookId = all.PlaybookId
	if all.CurrentStage != nil && all.CurrentStage.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CurrentStage = all.CurrentStage
	if all.DiagnosisBranch != nil && all.DiagnosisBranch.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DiagnosisBranch = all.DiagnosisBranch
	if all.ActionBranch != nil && all.ActionBranch.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ActionBranch = all.ActionBranch
	o.RetryOfRunId = all.RetryOfRunId
	o.ForkedFromRunId = all.ForkedFromRunId
	o.EventsUrl = all.EventsUrl
	if all.Metadata != nil && all.Metadata.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Metadata = all.Metadata
	o.CreatedAt = *all.CreatedAt
	o.UpdatedAt = all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
