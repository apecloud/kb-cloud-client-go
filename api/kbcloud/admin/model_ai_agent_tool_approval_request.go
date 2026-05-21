// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentToolApprovalRequest struct {
	ApprovalRequestId string `json:"approvalRequestId"`
	RunId             string `json:"runId"`
	ToolCallId        string `json:"toolCallId"`
	AgentToolId       string `json:"agentToolId"`
	DisplayName       string `json:"displayName"`
	// Organization-scoped single-cluster Agent execution scope.
	Scope            *AiAgentResourceScope     `json:"scope,omitempty"`
	ScopeLabel       *string                   `json:"scopeLabel,omitempty"`
	Reason           *string                   `json:"reason,omitempty"`
	ExpectedEvidence []string                  `json:"expectedEvidence,omitempty"`
	RiskLevel        *string                   `json:"riskLevel,omitempty"`
	ExpectedImpact   []string                  `json:"expectedImpact,omitempty"`
	ChangesResources *bool                     `json:"changesResources,omitempty"`
	Budget           map[string]interface{}    `json:"budget,omitempty"`
	GuardrailReason  *string                   `json:"guardrailReason,omitempty"`
	ApprovalStatus   AiAgentToolApprovalStatus `json:"approvalStatus"`
	// Stable status code plus default user-facing copy. Frontend uses code for logic/tests and may override display text for i18n.
	StatusDisplay *AiAgentStatusDisplay `json:"statusDisplay,omitempty"`
	AuditRef      string                `json:"auditRef"`
	// Hash of the Runtime-validated tool arguments and scope used for idempotency and tamper checks. Raw tool args are not exposed.
	ArgsHash         string     `json:"argsHash"`
	ApprovalVersion  int64      `json:"approvalVersion"`
	ExpiresAt        time.Time  `json:"expiresAt"`
	DecidedAt        *time.Time `json:"decidedAt,omitempty"`
	Feedback         *string    `json:"feedback,omitempty"`
	SafeEvidenceRefs []string   `json:"safeEvidenceRefs,omitempty"`
	CreatedAt        time.Time  `json:"createdAt"`
	UpdatedAt        *time.Time `json:"updatedAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentToolApprovalRequest instantiates a new AiAgentToolApprovalRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentToolApprovalRequest(approvalRequestId string, runId string, toolCallId string, agentToolId string, displayName string, approvalStatus AiAgentToolApprovalStatus, auditRef string, argsHash string, approvalVersion int64, expiresAt time.Time, createdAt time.Time) *AiAgentToolApprovalRequest {
	this := AiAgentToolApprovalRequest{}
	this.ApprovalRequestId = approvalRequestId
	this.RunId = runId
	this.ToolCallId = toolCallId
	this.AgentToolId = agentToolId
	this.DisplayName = displayName
	this.ApprovalStatus = approvalStatus
	this.AuditRef = auditRef
	this.ArgsHash = argsHash
	this.ApprovalVersion = approvalVersion
	this.ExpiresAt = expiresAt
	this.CreatedAt = createdAt
	return &this
}

// NewAiAgentToolApprovalRequestWithDefaults instantiates a new AiAgentToolApprovalRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentToolApprovalRequestWithDefaults() *AiAgentToolApprovalRequest {
	this := AiAgentToolApprovalRequest{}
	return &this
}

// GetApprovalRequestId returns the ApprovalRequestId field value.
func (o *AiAgentToolApprovalRequest) GetApprovalRequestId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ApprovalRequestId
}

// GetApprovalRequestIdOk returns a tuple with the ApprovalRequestId field value
// and a boolean to check if the value has been set.
func (o *AiAgentToolApprovalRequest) GetApprovalRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApprovalRequestId, true
}

// SetApprovalRequestId sets field value.
func (o *AiAgentToolApprovalRequest) SetApprovalRequestId(v string) {
	o.ApprovalRequestId = v
}

// GetRunId returns the RunId field value.
func (o *AiAgentToolApprovalRequest) GetRunId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RunId
}

// GetRunIdOk returns a tuple with the RunId field value
// and a boolean to check if the value has been set.
func (o *AiAgentToolApprovalRequest) GetRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunId, true
}

// SetRunId sets field value.
func (o *AiAgentToolApprovalRequest) SetRunId(v string) {
	o.RunId = v
}

// GetToolCallId returns the ToolCallId field value.
func (o *AiAgentToolApprovalRequest) GetToolCallId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ToolCallId
}

// GetToolCallIdOk returns a tuple with the ToolCallId field value
// and a boolean to check if the value has been set.
func (o *AiAgentToolApprovalRequest) GetToolCallIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToolCallId, true
}

// SetToolCallId sets field value.
func (o *AiAgentToolApprovalRequest) SetToolCallId(v string) {
	o.ToolCallId = v
}

// GetAgentToolId returns the AgentToolId field value.
func (o *AiAgentToolApprovalRequest) GetAgentToolId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AgentToolId
}

// GetAgentToolIdOk returns a tuple with the AgentToolId field value
// and a boolean to check if the value has been set.
func (o *AiAgentToolApprovalRequest) GetAgentToolIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentToolId, true
}

// SetAgentToolId sets field value.
func (o *AiAgentToolApprovalRequest) SetAgentToolId(v string) {
	o.AgentToolId = v
}

// GetDisplayName returns the DisplayName field value.
func (o *AiAgentToolApprovalRequest) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *AiAgentToolApprovalRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value.
func (o *AiAgentToolApprovalRequest) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *AiAgentToolApprovalRequest) GetScope() AiAgentResourceScope {
	if o == nil || o.Scope == nil {
		var ret AiAgentResourceScope
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentToolApprovalRequest) GetScopeOk() (*AiAgentResourceScope, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *AiAgentToolApprovalRequest) HasScope() bool {
	return o != nil && o.Scope != nil
}

// SetScope gets a reference to the given AiAgentResourceScope and assigns it to the Scope field.
func (o *AiAgentToolApprovalRequest) SetScope(v AiAgentResourceScope) {
	o.Scope = &v
}

// GetScopeLabel returns the ScopeLabel field value if set, zero value otherwise.
func (o *AiAgentToolApprovalRequest) GetScopeLabel() string {
	if o == nil || o.ScopeLabel == nil {
		var ret string
		return ret
	}
	return *o.ScopeLabel
}

// GetScopeLabelOk returns a tuple with the ScopeLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentToolApprovalRequest) GetScopeLabelOk() (*string, bool) {
	if o == nil || o.ScopeLabel == nil {
		return nil, false
	}
	return o.ScopeLabel, true
}

// HasScopeLabel returns a boolean if a field has been set.
func (o *AiAgentToolApprovalRequest) HasScopeLabel() bool {
	return o != nil && o.ScopeLabel != nil
}

// SetScopeLabel gets a reference to the given string and assigns it to the ScopeLabel field.
func (o *AiAgentToolApprovalRequest) SetScopeLabel(v string) {
	o.ScopeLabel = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *AiAgentToolApprovalRequest) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentToolApprovalRequest) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *AiAgentToolApprovalRequest) HasReason() bool {
	return o != nil && o.Reason != nil
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *AiAgentToolApprovalRequest) SetReason(v string) {
	o.Reason = &v
}

// GetExpectedEvidence returns the ExpectedEvidence field value if set, zero value otherwise.
func (o *AiAgentToolApprovalRequest) GetExpectedEvidence() []string {
	if o == nil || o.ExpectedEvidence == nil {
		var ret []string
		return ret
	}
	return o.ExpectedEvidence
}

// GetExpectedEvidenceOk returns a tuple with the ExpectedEvidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentToolApprovalRequest) GetExpectedEvidenceOk() (*[]string, bool) {
	if o == nil || o.ExpectedEvidence == nil {
		return nil, false
	}
	return &o.ExpectedEvidence, true
}

// HasExpectedEvidence returns a boolean if a field has been set.
func (o *AiAgentToolApprovalRequest) HasExpectedEvidence() bool {
	return o != nil && o.ExpectedEvidence != nil
}

// SetExpectedEvidence gets a reference to the given []string and assigns it to the ExpectedEvidence field.
func (o *AiAgentToolApprovalRequest) SetExpectedEvidence(v []string) {
	o.ExpectedEvidence = v
}

// GetRiskLevel returns the RiskLevel field value if set, zero value otherwise.
func (o *AiAgentToolApprovalRequest) GetRiskLevel() string {
	if o == nil || o.RiskLevel == nil {
		var ret string
		return ret
	}
	return *o.RiskLevel
}

// GetRiskLevelOk returns a tuple with the RiskLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentToolApprovalRequest) GetRiskLevelOk() (*string, bool) {
	if o == nil || o.RiskLevel == nil {
		return nil, false
	}
	return o.RiskLevel, true
}

// HasRiskLevel returns a boolean if a field has been set.
func (o *AiAgentToolApprovalRequest) HasRiskLevel() bool {
	return o != nil && o.RiskLevel != nil
}

// SetRiskLevel gets a reference to the given string and assigns it to the RiskLevel field.
func (o *AiAgentToolApprovalRequest) SetRiskLevel(v string) {
	o.RiskLevel = &v
}

// GetExpectedImpact returns the ExpectedImpact field value if set, zero value otherwise.
func (o *AiAgentToolApprovalRequest) GetExpectedImpact() []string {
	if o == nil || o.ExpectedImpact == nil {
		var ret []string
		return ret
	}
	return o.ExpectedImpact
}

// GetExpectedImpactOk returns a tuple with the ExpectedImpact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentToolApprovalRequest) GetExpectedImpactOk() (*[]string, bool) {
	if o == nil || o.ExpectedImpact == nil {
		return nil, false
	}
	return &o.ExpectedImpact, true
}

// HasExpectedImpact returns a boolean if a field has been set.
func (o *AiAgentToolApprovalRequest) HasExpectedImpact() bool {
	return o != nil && o.ExpectedImpact != nil
}

// SetExpectedImpact gets a reference to the given []string and assigns it to the ExpectedImpact field.
func (o *AiAgentToolApprovalRequest) SetExpectedImpact(v []string) {
	o.ExpectedImpact = v
}

// GetChangesResources returns the ChangesResources field value if set, zero value otherwise.
func (o *AiAgentToolApprovalRequest) GetChangesResources() bool {
	if o == nil || o.ChangesResources == nil {
		var ret bool
		return ret
	}
	return *o.ChangesResources
}

// GetChangesResourcesOk returns a tuple with the ChangesResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentToolApprovalRequest) GetChangesResourcesOk() (*bool, bool) {
	if o == nil || o.ChangesResources == nil {
		return nil, false
	}
	return o.ChangesResources, true
}

// HasChangesResources returns a boolean if a field has been set.
func (o *AiAgentToolApprovalRequest) HasChangesResources() bool {
	return o != nil && o.ChangesResources != nil
}

// SetChangesResources gets a reference to the given bool and assigns it to the ChangesResources field.
func (o *AiAgentToolApprovalRequest) SetChangesResources(v bool) {
	o.ChangesResources = &v
}

// GetBudget returns the Budget field value if set, zero value otherwise.
func (o *AiAgentToolApprovalRequest) GetBudget() map[string]interface{} {
	if o == nil || o.Budget == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Budget
}

// GetBudgetOk returns a tuple with the Budget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentToolApprovalRequest) GetBudgetOk() (*map[string]interface{}, bool) {
	if o == nil || o.Budget == nil {
		return nil, false
	}
	return &o.Budget, true
}

// HasBudget returns a boolean if a field has been set.
func (o *AiAgentToolApprovalRequest) HasBudget() bool {
	return o != nil && o.Budget != nil
}

// SetBudget gets a reference to the given map[string]interface{} and assigns it to the Budget field.
func (o *AiAgentToolApprovalRequest) SetBudget(v map[string]interface{}) {
	o.Budget = v
}

// GetGuardrailReason returns the GuardrailReason field value if set, zero value otherwise.
func (o *AiAgentToolApprovalRequest) GetGuardrailReason() string {
	if o == nil || o.GuardrailReason == nil {
		var ret string
		return ret
	}
	return *o.GuardrailReason
}

// GetGuardrailReasonOk returns a tuple with the GuardrailReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentToolApprovalRequest) GetGuardrailReasonOk() (*string, bool) {
	if o == nil || o.GuardrailReason == nil {
		return nil, false
	}
	return o.GuardrailReason, true
}

// HasGuardrailReason returns a boolean if a field has been set.
func (o *AiAgentToolApprovalRequest) HasGuardrailReason() bool {
	return o != nil && o.GuardrailReason != nil
}

// SetGuardrailReason gets a reference to the given string and assigns it to the GuardrailReason field.
func (o *AiAgentToolApprovalRequest) SetGuardrailReason(v string) {
	o.GuardrailReason = &v
}

// GetApprovalStatus returns the ApprovalStatus field value.
func (o *AiAgentToolApprovalRequest) GetApprovalStatus() AiAgentToolApprovalStatus {
	if o == nil {
		var ret AiAgentToolApprovalStatus
		return ret
	}
	return o.ApprovalStatus
}

// GetApprovalStatusOk returns a tuple with the ApprovalStatus field value
// and a boolean to check if the value has been set.
func (o *AiAgentToolApprovalRequest) GetApprovalStatusOk() (*AiAgentToolApprovalStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApprovalStatus, true
}

// SetApprovalStatus sets field value.
func (o *AiAgentToolApprovalRequest) SetApprovalStatus(v AiAgentToolApprovalStatus) {
	o.ApprovalStatus = v
}

// GetStatusDisplay returns the StatusDisplay field value if set, zero value otherwise.
func (o *AiAgentToolApprovalRequest) GetStatusDisplay() AiAgentStatusDisplay {
	if o == nil || o.StatusDisplay == nil {
		var ret AiAgentStatusDisplay
		return ret
	}
	return *o.StatusDisplay
}

// GetStatusDisplayOk returns a tuple with the StatusDisplay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentToolApprovalRequest) GetStatusDisplayOk() (*AiAgentStatusDisplay, bool) {
	if o == nil || o.StatusDisplay == nil {
		return nil, false
	}
	return o.StatusDisplay, true
}

// HasStatusDisplay returns a boolean if a field has been set.
func (o *AiAgentToolApprovalRequest) HasStatusDisplay() bool {
	return o != nil && o.StatusDisplay != nil
}

// SetStatusDisplay gets a reference to the given AiAgentStatusDisplay and assigns it to the StatusDisplay field.
func (o *AiAgentToolApprovalRequest) SetStatusDisplay(v AiAgentStatusDisplay) {
	o.StatusDisplay = &v
}

// GetAuditRef returns the AuditRef field value.
func (o *AiAgentToolApprovalRequest) GetAuditRef() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AuditRef
}

// GetAuditRefOk returns a tuple with the AuditRef field value
// and a boolean to check if the value has been set.
func (o *AiAgentToolApprovalRequest) GetAuditRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuditRef, true
}

// SetAuditRef sets field value.
func (o *AiAgentToolApprovalRequest) SetAuditRef(v string) {
	o.AuditRef = v
}

// GetArgsHash returns the ArgsHash field value.
func (o *AiAgentToolApprovalRequest) GetArgsHash() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ArgsHash
}

// GetArgsHashOk returns a tuple with the ArgsHash field value
// and a boolean to check if the value has been set.
func (o *AiAgentToolApprovalRequest) GetArgsHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ArgsHash, true
}

// SetArgsHash sets field value.
func (o *AiAgentToolApprovalRequest) SetArgsHash(v string) {
	o.ArgsHash = v
}

// GetApprovalVersion returns the ApprovalVersion field value.
func (o *AiAgentToolApprovalRequest) GetApprovalVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ApprovalVersion
}

// GetApprovalVersionOk returns a tuple with the ApprovalVersion field value
// and a boolean to check if the value has been set.
func (o *AiAgentToolApprovalRequest) GetApprovalVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApprovalVersion, true
}

// SetApprovalVersion sets field value.
func (o *AiAgentToolApprovalRequest) SetApprovalVersion(v int64) {
	o.ApprovalVersion = v
}

// GetExpiresAt returns the ExpiresAt field value.
func (o *AiAgentToolApprovalRequest) GetExpiresAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *AiAgentToolApprovalRequest) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value.
func (o *AiAgentToolApprovalRequest) SetExpiresAt(v time.Time) {
	o.ExpiresAt = v
}

// GetDecidedAt returns the DecidedAt field value if set, zero value otherwise.
func (o *AiAgentToolApprovalRequest) GetDecidedAt() time.Time {
	if o == nil || o.DecidedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.DecidedAt
}

// GetDecidedAtOk returns a tuple with the DecidedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentToolApprovalRequest) GetDecidedAtOk() (*time.Time, bool) {
	if o == nil || o.DecidedAt == nil {
		return nil, false
	}
	return o.DecidedAt, true
}

// HasDecidedAt returns a boolean if a field has been set.
func (o *AiAgentToolApprovalRequest) HasDecidedAt() bool {
	return o != nil && o.DecidedAt != nil
}

// SetDecidedAt gets a reference to the given time.Time and assigns it to the DecidedAt field.
func (o *AiAgentToolApprovalRequest) SetDecidedAt(v time.Time) {
	o.DecidedAt = &v
}

// GetFeedback returns the Feedback field value if set, zero value otherwise.
func (o *AiAgentToolApprovalRequest) GetFeedback() string {
	if o == nil || o.Feedback == nil {
		var ret string
		return ret
	}
	return *o.Feedback
}

// GetFeedbackOk returns a tuple with the Feedback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentToolApprovalRequest) GetFeedbackOk() (*string, bool) {
	if o == nil || o.Feedback == nil {
		return nil, false
	}
	return o.Feedback, true
}

// HasFeedback returns a boolean if a field has been set.
func (o *AiAgentToolApprovalRequest) HasFeedback() bool {
	return o != nil && o.Feedback != nil
}

// SetFeedback gets a reference to the given string and assigns it to the Feedback field.
func (o *AiAgentToolApprovalRequest) SetFeedback(v string) {
	o.Feedback = &v
}

// GetSafeEvidenceRefs returns the SafeEvidenceRefs field value if set, zero value otherwise.
func (o *AiAgentToolApprovalRequest) GetSafeEvidenceRefs() []string {
	if o == nil || o.SafeEvidenceRefs == nil {
		var ret []string
		return ret
	}
	return o.SafeEvidenceRefs
}

// GetSafeEvidenceRefsOk returns a tuple with the SafeEvidenceRefs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentToolApprovalRequest) GetSafeEvidenceRefsOk() (*[]string, bool) {
	if o == nil || o.SafeEvidenceRefs == nil {
		return nil, false
	}
	return &o.SafeEvidenceRefs, true
}

// HasSafeEvidenceRefs returns a boolean if a field has been set.
func (o *AiAgentToolApprovalRequest) HasSafeEvidenceRefs() bool {
	return o != nil && o.SafeEvidenceRefs != nil
}

// SetSafeEvidenceRefs gets a reference to the given []string and assigns it to the SafeEvidenceRefs field.
func (o *AiAgentToolApprovalRequest) SetSafeEvidenceRefs(v []string) {
	o.SafeEvidenceRefs = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *AiAgentToolApprovalRequest) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AiAgentToolApprovalRequest) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *AiAgentToolApprovalRequest) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AiAgentToolApprovalRequest) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentToolApprovalRequest) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AiAgentToolApprovalRequest) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *AiAgentToolApprovalRequest) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentToolApprovalRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["approvalRequestId"] = o.ApprovalRequestId
	toSerialize["runId"] = o.RunId
	toSerialize["toolCallId"] = o.ToolCallId
	toSerialize["agentToolId"] = o.AgentToolId
	toSerialize["displayName"] = o.DisplayName
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.ScopeLabel != nil {
		toSerialize["scopeLabel"] = o.ScopeLabel
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.ExpectedEvidence != nil {
		toSerialize["expectedEvidence"] = o.ExpectedEvidence
	}
	if o.RiskLevel != nil {
		toSerialize["riskLevel"] = o.RiskLevel
	}
	if o.ExpectedImpact != nil {
		toSerialize["expectedImpact"] = o.ExpectedImpact
	}
	if o.ChangesResources != nil {
		toSerialize["changesResources"] = o.ChangesResources
	}
	if o.Budget != nil {
		toSerialize["budget"] = o.Budget
	}
	if o.GuardrailReason != nil {
		toSerialize["guardrailReason"] = o.GuardrailReason
	}
	toSerialize["approvalStatus"] = o.ApprovalStatus
	if o.StatusDisplay != nil {
		toSerialize["statusDisplay"] = o.StatusDisplay
	}
	toSerialize["auditRef"] = o.AuditRef
	toSerialize["argsHash"] = o.ArgsHash
	toSerialize["approvalVersion"] = o.ApprovalVersion
	if o.ExpiresAt.Nanosecond() == 0 {
		toSerialize["expiresAt"] = o.ExpiresAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["expiresAt"] = o.ExpiresAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.DecidedAt != nil {
		if o.DecidedAt.Nanosecond() == 0 {
			toSerialize["decidedAt"] = o.DecidedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["decidedAt"] = o.DecidedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Feedback != nil {
		toSerialize["feedback"] = o.Feedback
	}
	if o.SafeEvidenceRefs != nil {
		toSerialize["safeEvidenceRefs"] = o.SafeEvidenceRefs
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
func (o *AiAgentToolApprovalRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApprovalRequestId *string                    `json:"approvalRequestId"`
		RunId             *string                    `json:"runId"`
		ToolCallId        *string                    `json:"toolCallId"`
		AgentToolId       *string                    `json:"agentToolId"`
		DisplayName       *string                    `json:"displayName"`
		Scope             *AiAgentResourceScope      `json:"scope,omitempty"`
		ScopeLabel        *string                    `json:"scopeLabel,omitempty"`
		Reason            *string                    `json:"reason,omitempty"`
		ExpectedEvidence  []string                   `json:"expectedEvidence,omitempty"`
		RiskLevel         *string                    `json:"riskLevel,omitempty"`
		ExpectedImpact    []string                   `json:"expectedImpact,omitempty"`
		ChangesResources  *bool                      `json:"changesResources,omitempty"`
		Budget            map[string]interface{}     `json:"budget,omitempty"`
		GuardrailReason   *string                    `json:"guardrailReason,omitempty"`
		ApprovalStatus    *AiAgentToolApprovalStatus `json:"approvalStatus"`
		StatusDisplay     *AiAgentStatusDisplay      `json:"statusDisplay,omitempty"`
		AuditRef          *string                    `json:"auditRef"`
		ArgsHash          *string                    `json:"argsHash"`
		ApprovalVersion   *int64                     `json:"approvalVersion"`
		ExpiresAt         *time.Time                 `json:"expiresAt"`
		DecidedAt         *time.Time                 `json:"decidedAt,omitempty"`
		Feedback          *string                    `json:"feedback,omitempty"`
		SafeEvidenceRefs  []string                   `json:"safeEvidenceRefs,omitempty"`
		CreatedAt         *time.Time                 `json:"createdAt"`
		UpdatedAt         *time.Time                 `json:"updatedAt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.ApprovalRequestId == nil {
		return fmt.Errorf("required field approvalRequestId missing")
	}
	if all.RunId == nil {
		return fmt.Errorf("required field runId missing")
	}
	if all.ToolCallId == nil {
		return fmt.Errorf("required field toolCallId missing")
	}
	if all.AgentToolId == nil {
		return fmt.Errorf("required field agentToolId missing")
	}
	if all.DisplayName == nil {
		return fmt.Errorf("required field displayName missing")
	}
	if all.ApprovalStatus == nil {
		return fmt.Errorf("required field approvalStatus missing")
	}
	if all.AuditRef == nil {
		return fmt.Errorf("required field auditRef missing")
	}
	if all.ArgsHash == nil {
		return fmt.Errorf("required field argsHash missing")
	}
	if all.ApprovalVersion == nil {
		return fmt.Errorf("required field approvalVersion missing")
	}
	if all.ExpiresAt == nil {
		return fmt.Errorf("required field expiresAt missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field createdAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"approvalRequestId", "runId", "toolCallId", "agentToolId", "displayName", "scope", "scopeLabel", "reason", "expectedEvidence", "riskLevel", "expectedImpact", "changesResources", "budget", "guardrailReason", "approvalStatus", "statusDisplay", "auditRef", "argsHash", "approvalVersion", "expiresAt", "decidedAt", "feedback", "safeEvidenceRefs", "createdAt", "updatedAt"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ApprovalRequestId = *all.ApprovalRequestId
	o.RunId = *all.RunId
	o.ToolCallId = *all.ToolCallId
	o.AgentToolId = *all.AgentToolId
	o.DisplayName = *all.DisplayName
	if all.Scope != nil && all.Scope.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Scope = all.Scope
	o.ScopeLabel = all.ScopeLabel
	o.Reason = all.Reason
	o.ExpectedEvidence = all.ExpectedEvidence
	o.RiskLevel = all.RiskLevel
	o.ExpectedImpact = all.ExpectedImpact
	o.ChangesResources = all.ChangesResources
	o.Budget = all.Budget
	o.GuardrailReason = all.GuardrailReason
	if !all.ApprovalStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.ApprovalStatus = *all.ApprovalStatus
	}
	if all.StatusDisplay != nil && all.StatusDisplay.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.StatusDisplay = all.StatusDisplay
	o.AuditRef = *all.AuditRef
	o.ArgsHash = *all.ArgsHash
	o.ApprovalVersion = *all.ApprovalVersion
	o.ExpiresAt = *all.ExpiresAt
	o.DecidedAt = all.DecidedAt
	o.Feedback = all.Feedback
	o.SafeEvidenceRefs = all.SafeEvidenceRefs
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
