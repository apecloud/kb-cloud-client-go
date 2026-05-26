// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentActionPlan struct {
	ActionPlanId   string                  `json:"actionPlanId"`
	RunId          string                  `json:"runId"`
	ConversationId string                  `json:"conversationId"`
	Status         AiAgentActionPlanStatus `json:"status"`
	// Stable status code plus default user-facing copy. Frontend uses code for logic/tests and may override display text for i18n.
	StatusDisplay     *AiAgentStatusDisplay `json:"statusDisplay,omitempty"`
	OperationCategory string                `json:"operationCategory"`
	Title             string                `json:"title"`
	// Organization-scoped single-cluster Agent execution scope.
	Target                   *AiAgentResourceScope    `json:"target,omitempty"`
	ArgsSummary              map[string]interface{}   `json:"argsSummary,omitempty"`
	Impact                   []string                 `json:"impact,omitempty"`
	RiskLevel                *string                  `json:"riskLevel,omitempty"`
	RiskExplanation          *string                  `json:"riskExplanation,omitempty"`
	PermissionRequired       []string                 `json:"permissionRequired,omitempty"`
	Verification             []string                 `json:"verification,omitempty"`
	Rollback                 []string                 `json:"rollback,omitempty"`
	ConfirmLabel             *string                  `json:"confirmLabel,omitempty"`
	CancelLabel              *string                  `json:"cancelLabel,omitempty"`
	ConfirmationMode         *AiAgentConfirmationMode `json:"confirmationMode,omitempty"`
	ConfirmationRequiredText *string                  `json:"confirmationRequiredText,omitempty"`
	// Stable status code plus default user-facing copy. Frontend uses code for logic/tests and may override display text for i18n.
	DisabledReason     *AiAgentStatusDisplay `json:"disabledReason,omitempty"`
	DisabledReasonCode *string               `json:"disabledReasonCode,omitempty"`
	ExpiresAt          *time.Time            `json:"expiresAt,omitempty"`
	ExpiresInSeconds   *int64                `json:"expiresInSeconds,omitempty"`
	SourceEvidenceIds  []string              `json:"sourceEvidenceIds,omitempty"`
	CreatedAt          *time.Time            `json:"createdAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentActionPlan instantiates a new AiAgentActionPlan object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentActionPlan(actionPlanId string, runId string, conversationId string, status AiAgentActionPlanStatus, operationCategory string, title string) *AiAgentActionPlan {
	this := AiAgentActionPlan{}
	this.ActionPlanId = actionPlanId
	this.RunId = runId
	this.ConversationId = conversationId
	this.Status = status
	this.OperationCategory = operationCategory
	this.Title = title
	return &this
}

// NewAiAgentActionPlanWithDefaults instantiates a new AiAgentActionPlan object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentActionPlanWithDefaults() *AiAgentActionPlan {
	this := AiAgentActionPlan{}
	return &this
}

// GetActionPlanId returns the ActionPlanId field value.
func (o *AiAgentActionPlan) GetActionPlanId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ActionPlanId
}

// GetActionPlanIdOk returns a tuple with the ActionPlanId field value
// and a boolean to check if the value has been set.
func (o *AiAgentActionPlan) GetActionPlanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionPlanId, true
}

// SetActionPlanId sets field value.
func (o *AiAgentActionPlan) SetActionPlanId(v string) {
	o.ActionPlanId = v
}

// GetRunId returns the RunId field value.
func (o *AiAgentActionPlan) GetRunId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RunId
}

// GetRunIdOk returns a tuple with the RunId field value
// and a boolean to check if the value has been set.
func (o *AiAgentActionPlan) GetRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunId, true
}

// SetRunId sets field value.
func (o *AiAgentActionPlan) SetRunId(v string) {
	o.RunId = v
}

// GetConversationId returns the ConversationId field value.
func (o *AiAgentActionPlan) GetConversationId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ConversationId
}

// GetConversationIdOk returns a tuple with the ConversationId field value
// and a boolean to check if the value has been set.
func (o *AiAgentActionPlan) GetConversationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConversationId, true
}

// SetConversationId sets field value.
func (o *AiAgentActionPlan) SetConversationId(v string) {
	o.ConversationId = v
}

// GetStatus returns the Status field value.
func (o *AiAgentActionPlan) GetStatus() AiAgentActionPlanStatus {
	if o == nil {
		var ret AiAgentActionPlanStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AiAgentActionPlan) GetStatusOk() (*AiAgentActionPlanStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *AiAgentActionPlan) SetStatus(v AiAgentActionPlanStatus) {
	o.Status = v
}

// GetStatusDisplay returns the StatusDisplay field value if set, zero value otherwise.
func (o *AiAgentActionPlan) GetStatusDisplay() AiAgentStatusDisplay {
	if o == nil || o.StatusDisplay == nil {
		var ret AiAgentStatusDisplay
		return ret
	}
	return *o.StatusDisplay
}

// GetStatusDisplayOk returns a tuple with the StatusDisplay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentActionPlan) GetStatusDisplayOk() (*AiAgentStatusDisplay, bool) {
	if o == nil || o.StatusDisplay == nil {
		return nil, false
	}
	return o.StatusDisplay, true
}

// HasStatusDisplay returns a boolean if a field has been set.
func (o *AiAgentActionPlan) HasStatusDisplay() bool {
	return o != nil && o.StatusDisplay != nil
}

// SetStatusDisplay gets a reference to the given AiAgentStatusDisplay and assigns it to the StatusDisplay field.
func (o *AiAgentActionPlan) SetStatusDisplay(v AiAgentStatusDisplay) {
	o.StatusDisplay = &v
}

// GetOperationCategory returns the OperationCategory field value.
func (o *AiAgentActionPlan) GetOperationCategory() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OperationCategory
}

// GetOperationCategoryOk returns a tuple with the OperationCategory field value
// and a boolean to check if the value has been set.
func (o *AiAgentActionPlan) GetOperationCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationCategory, true
}

// SetOperationCategory sets field value.
func (o *AiAgentActionPlan) SetOperationCategory(v string) {
	o.OperationCategory = v
}

// GetTitle returns the Title field value.
func (o *AiAgentActionPlan) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *AiAgentActionPlan) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *AiAgentActionPlan) SetTitle(v string) {
	o.Title = v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *AiAgentActionPlan) GetTarget() AiAgentResourceScope {
	if o == nil || o.Target == nil {
		var ret AiAgentResourceScope
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentActionPlan) GetTargetOk() (*AiAgentResourceScope, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *AiAgentActionPlan) HasTarget() bool {
	return o != nil && o.Target != nil
}

// SetTarget gets a reference to the given AiAgentResourceScope and assigns it to the Target field.
func (o *AiAgentActionPlan) SetTarget(v AiAgentResourceScope) {
	o.Target = &v
}

// GetArgsSummary returns the ArgsSummary field value if set, zero value otherwise.
func (o *AiAgentActionPlan) GetArgsSummary() map[string]interface{} {
	if o == nil || o.ArgsSummary == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ArgsSummary
}

// GetArgsSummaryOk returns a tuple with the ArgsSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentActionPlan) GetArgsSummaryOk() (*map[string]interface{}, bool) {
	if o == nil || o.ArgsSummary == nil {
		return nil, false
	}
	return &o.ArgsSummary, true
}

// HasArgsSummary returns a boolean if a field has been set.
func (o *AiAgentActionPlan) HasArgsSummary() bool {
	return o != nil && o.ArgsSummary != nil
}

// SetArgsSummary gets a reference to the given map[string]interface{} and assigns it to the ArgsSummary field.
func (o *AiAgentActionPlan) SetArgsSummary(v map[string]interface{}) {
	o.ArgsSummary = v
}

// GetImpact returns the Impact field value if set, zero value otherwise.
func (o *AiAgentActionPlan) GetImpact() []string {
	if o == nil || o.Impact == nil {
		var ret []string
		return ret
	}
	return o.Impact
}

// GetImpactOk returns a tuple with the Impact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentActionPlan) GetImpactOk() (*[]string, bool) {
	if o == nil || o.Impact == nil {
		return nil, false
	}
	return &o.Impact, true
}

// HasImpact returns a boolean if a field has been set.
func (o *AiAgentActionPlan) HasImpact() bool {
	return o != nil && o.Impact != nil
}

// SetImpact gets a reference to the given []string and assigns it to the Impact field.
func (o *AiAgentActionPlan) SetImpact(v []string) {
	o.Impact = v
}

// GetRiskLevel returns the RiskLevel field value if set, zero value otherwise.
func (o *AiAgentActionPlan) GetRiskLevel() string {
	if o == nil || o.RiskLevel == nil {
		var ret string
		return ret
	}
	return *o.RiskLevel
}

// GetRiskLevelOk returns a tuple with the RiskLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentActionPlan) GetRiskLevelOk() (*string, bool) {
	if o == nil || o.RiskLevel == nil {
		return nil, false
	}
	return o.RiskLevel, true
}

// HasRiskLevel returns a boolean if a field has been set.
func (o *AiAgentActionPlan) HasRiskLevel() bool {
	return o != nil && o.RiskLevel != nil
}

// SetRiskLevel gets a reference to the given string and assigns it to the RiskLevel field.
func (o *AiAgentActionPlan) SetRiskLevel(v string) {
	o.RiskLevel = &v
}

// GetRiskExplanation returns the RiskExplanation field value if set, zero value otherwise.
func (o *AiAgentActionPlan) GetRiskExplanation() string {
	if o == nil || o.RiskExplanation == nil {
		var ret string
		return ret
	}
	return *o.RiskExplanation
}

// GetRiskExplanationOk returns a tuple with the RiskExplanation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentActionPlan) GetRiskExplanationOk() (*string, bool) {
	if o == nil || o.RiskExplanation == nil {
		return nil, false
	}
	return o.RiskExplanation, true
}

// HasRiskExplanation returns a boolean if a field has been set.
func (o *AiAgentActionPlan) HasRiskExplanation() bool {
	return o != nil && o.RiskExplanation != nil
}

// SetRiskExplanation gets a reference to the given string and assigns it to the RiskExplanation field.
func (o *AiAgentActionPlan) SetRiskExplanation(v string) {
	o.RiskExplanation = &v
}

// GetPermissionRequired returns the PermissionRequired field value if set, zero value otherwise.
func (o *AiAgentActionPlan) GetPermissionRequired() []string {
	if o == nil || o.PermissionRequired == nil {
		var ret []string
		return ret
	}
	return o.PermissionRequired
}

// GetPermissionRequiredOk returns a tuple with the PermissionRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentActionPlan) GetPermissionRequiredOk() (*[]string, bool) {
	if o == nil || o.PermissionRequired == nil {
		return nil, false
	}
	return &o.PermissionRequired, true
}

// HasPermissionRequired returns a boolean if a field has been set.
func (o *AiAgentActionPlan) HasPermissionRequired() bool {
	return o != nil && o.PermissionRequired != nil
}

// SetPermissionRequired gets a reference to the given []string and assigns it to the PermissionRequired field.
func (o *AiAgentActionPlan) SetPermissionRequired(v []string) {
	o.PermissionRequired = v
}

// GetVerification returns the Verification field value if set, zero value otherwise.
func (o *AiAgentActionPlan) GetVerification() []string {
	if o == nil || o.Verification == nil {
		var ret []string
		return ret
	}
	return o.Verification
}

// GetVerificationOk returns a tuple with the Verification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentActionPlan) GetVerificationOk() (*[]string, bool) {
	if o == nil || o.Verification == nil {
		return nil, false
	}
	return &o.Verification, true
}

// HasVerification returns a boolean if a field has been set.
func (o *AiAgentActionPlan) HasVerification() bool {
	return o != nil && o.Verification != nil
}

// SetVerification gets a reference to the given []string and assigns it to the Verification field.
func (o *AiAgentActionPlan) SetVerification(v []string) {
	o.Verification = v
}

// GetRollback returns the Rollback field value if set, zero value otherwise.
func (o *AiAgentActionPlan) GetRollback() []string {
	if o == nil || o.Rollback == nil {
		var ret []string
		return ret
	}
	return o.Rollback
}

// GetRollbackOk returns a tuple with the Rollback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentActionPlan) GetRollbackOk() (*[]string, bool) {
	if o == nil || o.Rollback == nil {
		return nil, false
	}
	return &o.Rollback, true
}

// HasRollback returns a boolean if a field has been set.
func (o *AiAgentActionPlan) HasRollback() bool {
	return o != nil && o.Rollback != nil
}

// SetRollback gets a reference to the given []string and assigns it to the Rollback field.
func (o *AiAgentActionPlan) SetRollback(v []string) {
	o.Rollback = v
}

// GetConfirmLabel returns the ConfirmLabel field value if set, zero value otherwise.
func (o *AiAgentActionPlan) GetConfirmLabel() string {
	if o == nil || o.ConfirmLabel == nil {
		var ret string
		return ret
	}
	return *o.ConfirmLabel
}

// GetConfirmLabelOk returns a tuple with the ConfirmLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentActionPlan) GetConfirmLabelOk() (*string, bool) {
	if o == nil || o.ConfirmLabel == nil {
		return nil, false
	}
	return o.ConfirmLabel, true
}

// HasConfirmLabel returns a boolean if a field has been set.
func (o *AiAgentActionPlan) HasConfirmLabel() bool {
	return o != nil && o.ConfirmLabel != nil
}

// SetConfirmLabel gets a reference to the given string and assigns it to the ConfirmLabel field.
func (o *AiAgentActionPlan) SetConfirmLabel(v string) {
	o.ConfirmLabel = &v
}

// GetCancelLabel returns the CancelLabel field value if set, zero value otherwise.
func (o *AiAgentActionPlan) GetCancelLabel() string {
	if o == nil || o.CancelLabel == nil {
		var ret string
		return ret
	}
	return *o.CancelLabel
}

// GetCancelLabelOk returns a tuple with the CancelLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentActionPlan) GetCancelLabelOk() (*string, bool) {
	if o == nil || o.CancelLabel == nil {
		return nil, false
	}
	return o.CancelLabel, true
}

// HasCancelLabel returns a boolean if a field has been set.
func (o *AiAgentActionPlan) HasCancelLabel() bool {
	return o != nil && o.CancelLabel != nil
}

// SetCancelLabel gets a reference to the given string and assigns it to the CancelLabel field.
func (o *AiAgentActionPlan) SetCancelLabel(v string) {
	o.CancelLabel = &v
}

// GetConfirmationMode returns the ConfirmationMode field value if set, zero value otherwise.
func (o *AiAgentActionPlan) GetConfirmationMode() AiAgentConfirmationMode {
	if o == nil || o.ConfirmationMode == nil {
		var ret AiAgentConfirmationMode
		return ret
	}
	return *o.ConfirmationMode
}

// GetConfirmationModeOk returns a tuple with the ConfirmationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentActionPlan) GetConfirmationModeOk() (*AiAgentConfirmationMode, bool) {
	if o == nil || o.ConfirmationMode == nil {
		return nil, false
	}
	return o.ConfirmationMode, true
}

// HasConfirmationMode returns a boolean if a field has been set.
func (o *AiAgentActionPlan) HasConfirmationMode() bool {
	return o != nil && o.ConfirmationMode != nil
}

// SetConfirmationMode gets a reference to the given AiAgentConfirmationMode and assigns it to the ConfirmationMode field.
func (o *AiAgentActionPlan) SetConfirmationMode(v AiAgentConfirmationMode) {
	o.ConfirmationMode = &v
}

// GetConfirmationRequiredText returns the ConfirmationRequiredText field value if set, zero value otherwise.
func (o *AiAgentActionPlan) GetConfirmationRequiredText() string {
	if o == nil || o.ConfirmationRequiredText == nil {
		var ret string
		return ret
	}
	return *o.ConfirmationRequiredText
}

// GetConfirmationRequiredTextOk returns a tuple with the ConfirmationRequiredText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentActionPlan) GetConfirmationRequiredTextOk() (*string, bool) {
	if o == nil || o.ConfirmationRequiredText == nil {
		return nil, false
	}
	return o.ConfirmationRequiredText, true
}

// HasConfirmationRequiredText returns a boolean if a field has been set.
func (o *AiAgentActionPlan) HasConfirmationRequiredText() bool {
	return o != nil && o.ConfirmationRequiredText != nil
}

// SetConfirmationRequiredText gets a reference to the given string and assigns it to the ConfirmationRequiredText field.
func (o *AiAgentActionPlan) SetConfirmationRequiredText(v string) {
	o.ConfirmationRequiredText = &v
}

// GetDisabledReason returns the DisabledReason field value if set, zero value otherwise.
func (o *AiAgentActionPlan) GetDisabledReason() AiAgentStatusDisplay {
	if o == nil || o.DisabledReason == nil {
		var ret AiAgentStatusDisplay
		return ret
	}
	return *o.DisabledReason
}

// GetDisabledReasonOk returns a tuple with the DisabledReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentActionPlan) GetDisabledReasonOk() (*AiAgentStatusDisplay, bool) {
	if o == nil || o.DisabledReason == nil {
		return nil, false
	}
	return o.DisabledReason, true
}

// HasDisabledReason returns a boolean if a field has been set.
func (o *AiAgentActionPlan) HasDisabledReason() bool {
	return o != nil && o.DisabledReason != nil
}

// SetDisabledReason gets a reference to the given AiAgentStatusDisplay and assigns it to the DisabledReason field.
func (o *AiAgentActionPlan) SetDisabledReason(v AiAgentStatusDisplay) {
	o.DisabledReason = &v
}

// GetDisabledReasonCode returns the DisabledReasonCode field value if set, zero value otherwise.
func (o *AiAgentActionPlan) GetDisabledReasonCode() string {
	if o == nil || o.DisabledReasonCode == nil {
		var ret string
		return ret
	}
	return *o.DisabledReasonCode
}

// GetDisabledReasonCodeOk returns a tuple with the DisabledReasonCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentActionPlan) GetDisabledReasonCodeOk() (*string, bool) {
	if o == nil || o.DisabledReasonCode == nil {
		return nil, false
	}
	return o.DisabledReasonCode, true
}

// HasDisabledReasonCode returns a boolean if a field has been set.
func (o *AiAgentActionPlan) HasDisabledReasonCode() bool {
	return o != nil && o.DisabledReasonCode != nil
}

// SetDisabledReasonCode gets a reference to the given string and assigns it to the DisabledReasonCode field.
func (o *AiAgentActionPlan) SetDisabledReasonCode(v string) {
	o.DisabledReasonCode = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *AiAgentActionPlan) GetExpiresAt() time.Time {
	if o == nil || o.ExpiresAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentActionPlan) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *AiAgentActionPlan) HasExpiresAt() bool {
	return o != nil && o.ExpiresAt != nil
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *AiAgentActionPlan) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetExpiresInSeconds returns the ExpiresInSeconds field value if set, zero value otherwise.
func (o *AiAgentActionPlan) GetExpiresInSeconds() int64 {
	if o == nil || o.ExpiresInSeconds == nil {
		var ret int64
		return ret
	}
	return *o.ExpiresInSeconds
}

// GetExpiresInSecondsOk returns a tuple with the ExpiresInSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentActionPlan) GetExpiresInSecondsOk() (*int64, bool) {
	if o == nil || o.ExpiresInSeconds == nil {
		return nil, false
	}
	return o.ExpiresInSeconds, true
}

// HasExpiresInSeconds returns a boolean if a field has been set.
func (o *AiAgentActionPlan) HasExpiresInSeconds() bool {
	return o != nil && o.ExpiresInSeconds != nil
}

// SetExpiresInSeconds gets a reference to the given int64 and assigns it to the ExpiresInSeconds field.
func (o *AiAgentActionPlan) SetExpiresInSeconds(v int64) {
	o.ExpiresInSeconds = &v
}

// GetSourceEvidenceIds returns the SourceEvidenceIds field value if set, zero value otherwise.
func (o *AiAgentActionPlan) GetSourceEvidenceIds() []string {
	if o == nil || o.SourceEvidenceIds == nil {
		var ret []string
		return ret
	}
	return o.SourceEvidenceIds
}

// GetSourceEvidenceIdsOk returns a tuple with the SourceEvidenceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentActionPlan) GetSourceEvidenceIdsOk() (*[]string, bool) {
	if o == nil || o.SourceEvidenceIds == nil {
		return nil, false
	}
	return &o.SourceEvidenceIds, true
}

// HasSourceEvidenceIds returns a boolean if a field has been set.
func (o *AiAgentActionPlan) HasSourceEvidenceIds() bool {
	return o != nil && o.SourceEvidenceIds != nil
}

// SetSourceEvidenceIds gets a reference to the given []string and assigns it to the SourceEvidenceIds field.
func (o *AiAgentActionPlan) SetSourceEvidenceIds(v []string) {
	o.SourceEvidenceIds = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AiAgentActionPlan) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentActionPlan) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AiAgentActionPlan) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AiAgentActionPlan) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentActionPlan) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["actionPlanId"] = o.ActionPlanId
	toSerialize["runId"] = o.RunId
	toSerialize["conversationId"] = o.ConversationId
	toSerialize["status"] = o.Status
	if o.StatusDisplay != nil {
		toSerialize["statusDisplay"] = o.StatusDisplay
	}
	toSerialize["operationCategory"] = o.OperationCategory
	toSerialize["title"] = o.Title
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	if o.ArgsSummary != nil {
		toSerialize["argsSummary"] = o.ArgsSummary
	}
	if o.Impact != nil {
		toSerialize["impact"] = o.Impact
	}
	if o.RiskLevel != nil {
		toSerialize["riskLevel"] = o.RiskLevel
	}
	if o.RiskExplanation != nil {
		toSerialize["riskExplanation"] = o.RiskExplanation
	}
	if o.PermissionRequired != nil {
		toSerialize["permissionRequired"] = o.PermissionRequired
	}
	if o.Verification != nil {
		toSerialize["verification"] = o.Verification
	}
	if o.Rollback != nil {
		toSerialize["rollback"] = o.Rollback
	}
	if o.ConfirmLabel != nil {
		toSerialize["confirmLabel"] = o.ConfirmLabel
	}
	if o.CancelLabel != nil {
		toSerialize["cancelLabel"] = o.CancelLabel
	}
	if o.ConfirmationMode != nil {
		toSerialize["confirmationMode"] = o.ConfirmationMode
	}
	if o.ConfirmationRequiredText != nil {
		toSerialize["confirmationRequiredText"] = o.ConfirmationRequiredText
	}
	if o.DisabledReason != nil {
		toSerialize["disabledReason"] = o.DisabledReason
	}
	if o.DisabledReasonCode != nil {
		toSerialize["disabledReasonCode"] = o.DisabledReasonCode
	}
	if o.ExpiresAt != nil {
		if o.ExpiresAt.Nanosecond() == 0 {
			toSerialize["expiresAt"] = o.ExpiresAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["expiresAt"] = o.ExpiresAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.ExpiresInSeconds != nil {
		toSerialize["expiresInSeconds"] = o.ExpiresInSeconds
	}
	if o.SourceEvidenceIds != nil {
		toSerialize["sourceEvidenceIds"] = o.SourceEvidenceIds
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentActionPlan) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ActionPlanId             *string                  `json:"actionPlanId"`
		RunId                    *string                  `json:"runId"`
		ConversationId           *string                  `json:"conversationId"`
		Status                   *AiAgentActionPlanStatus `json:"status"`
		StatusDisplay            *AiAgentStatusDisplay    `json:"statusDisplay,omitempty"`
		OperationCategory        *string                  `json:"operationCategory"`
		Title                    *string                  `json:"title"`
		Target                   *AiAgentResourceScope    `json:"target,omitempty"`
		ArgsSummary              map[string]interface{}   `json:"argsSummary,omitempty"`
		Impact                   []string                 `json:"impact,omitempty"`
		RiskLevel                *string                  `json:"riskLevel,omitempty"`
		RiskExplanation          *string                  `json:"riskExplanation,omitempty"`
		PermissionRequired       []string                 `json:"permissionRequired,omitempty"`
		Verification             []string                 `json:"verification,omitempty"`
		Rollback                 []string                 `json:"rollback,omitempty"`
		ConfirmLabel             *string                  `json:"confirmLabel,omitempty"`
		CancelLabel              *string                  `json:"cancelLabel,omitempty"`
		ConfirmationMode         *AiAgentConfirmationMode `json:"confirmationMode,omitempty"`
		ConfirmationRequiredText *string                  `json:"confirmationRequiredText,omitempty"`
		DisabledReason           *AiAgentStatusDisplay    `json:"disabledReason,omitempty"`
		DisabledReasonCode       *string                  `json:"disabledReasonCode,omitempty"`
		ExpiresAt                *time.Time               `json:"expiresAt,omitempty"`
		ExpiresInSeconds         *int64                   `json:"expiresInSeconds,omitempty"`
		SourceEvidenceIds        []string                 `json:"sourceEvidenceIds,omitempty"`
		CreatedAt                *time.Time               `json:"createdAt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.ActionPlanId == nil {
		return fmt.Errorf("required field actionPlanId missing")
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
	if all.OperationCategory == nil {
		return fmt.Errorf("required field operationCategory missing")
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"actionPlanId", "runId", "conversationId", "status", "statusDisplay", "operationCategory", "title", "target", "argsSummary", "impact", "riskLevel", "riskExplanation", "permissionRequired", "verification", "rollback", "confirmLabel", "cancelLabel", "confirmationMode", "confirmationRequiredText", "disabledReason", "disabledReasonCode", "expiresAt", "expiresInSeconds", "sourceEvidenceIds", "createdAt"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ActionPlanId = *all.ActionPlanId
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
	o.OperationCategory = *all.OperationCategory
	o.Title = *all.Title
	if all.Target != nil && all.Target.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Target = all.Target
	o.ArgsSummary = all.ArgsSummary
	o.Impact = all.Impact
	o.RiskLevel = all.RiskLevel
	o.RiskExplanation = all.RiskExplanation
	o.PermissionRequired = all.PermissionRequired
	o.Verification = all.Verification
	o.Rollback = all.Rollback
	o.ConfirmLabel = all.ConfirmLabel
	o.CancelLabel = all.CancelLabel
	if all.ConfirmationMode != nil && !all.ConfirmationMode.IsValid() {
		hasInvalidField = true
	} else {
		o.ConfirmationMode = all.ConfirmationMode
	}
	o.ConfirmationRequiredText = all.ConfirmationRequiredText
	if all.DisabledReason != nil && all.DisabledReason.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DisabledReason = all.DisabledReason
	o.DisabledReasonCode = all.DisabledReasonCode
	o.ExpiresAt = all.ExpiresAt
	o.ExpiresInSeconds = all.ExpiresInSeconds
	o.SourceEvidenceIds = all.SourceEvidenceIds
	o.CreatedAt = all.CreatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
