// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentReport struct {
	ReportId       string `json:"reportId"`
	RunId          string `json:"runId"`
	ConversationId string `json:"conversationId"`
	// Stable status code plus default user-facing copy. Frontend uses code for logic/tests and may override display text for i18n.
	Status                          AiAgentStatusDisplay            `json:"status"`
	FinalStatus                     AiAgentFinalStatus              `json:"finalStatus"`
	EvidenceCompleteness            AiAgentEvidenceCompleteness     `json:"evidenceCompleteness"`
	Confidence                      *AiAgentConfidence              `json:"confidence,omitempty"`
	Title                           *string                         `json:"title,omitempty"`
	Summary                         *string                         `json:"summary,omitempty"`
	RootCause                       *AiAgentRootCause               `json:"rootCause,omitempty"`
	CheckedScope                    []AiAgentCheckedScopeItem       `json:"checkedScope,omitempty"`
	EvidenceSummary                 []AiAgentEvidenceSummary        `json:"evidenceSummary,omitempty"`
	Uncertainty                     []AiAgentUncertainty            `json:"uncertainty,omitempty"`
	BlockingRequirementIds          []string                        `json:"blockingRequirementIds,omitempty"`
	DegradedRequirementIds          []string                        `json:"degradedRequirementIds,omitempty"`
	PermissionBlockedRequirementIds []string                        `json:"permissionBlockedRequirementIds,omitempty"`
	MissingReasonByRequirement      map[string]AiAgentMissingReason `json:"missingReasonByRequirement,omitempty"`
	PartialReason                   *AiAgentPartialReason           `json:"partialReason,omitempty"`
	BudgetStatus                    *string                         `json:"budgetStatus,omitempty"`
	ScopeValidationStatus           *AiAgentScopeValidationStatus   `json:"scopeValidationStatus,omitempty"`
	ApprovalRequired                *bool                           `json:"approvalRequired,omitempty"`
	ApprovalStatus                  *AiAgentApprovalStatus          `json:"approvalStatus,omitempty"`
	Recommendations                 []AiAgentRecommendation         `json:"recommendations,omitempty"`
	Visualizations                  []AiAgentVisualization          `json:"visualizations,omitempty"`
	// References aiAgentMetricSeries.metricSeriesId values for batch hydration.
	MetricSeriesRefs []string  `json:"metricSeriesRefs,omitempty"`
	SourceEventIds   []string  `json:"sourceEventIds,omitempty"`
	CreatedAt        time.Time `json:"createdAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentReport instantiates a new AiAgentReport object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentReport(reportId string, runId string, conversationId string, status AiAgentStatusDisplay, finalStatus AiAgentFinalStatus, evidenceCompleteness AiAgentEvidenceCompleteness, createdAt time.Time) *AiAgentReport {
	this := AiAgentReport{}
	this.ReportId = reportId
	this.RunId = runId
	this.ConversationId = conversationId
	this.Status = status
	this.FinalStatus = finalStatus
	this.EvidenceCompleteness = evidenceCompleteness
	this.CreatedAt = createdAt
	return &this
}

// NewAiAgentReportWithDefaults instantiates a new AiAgentReport object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentReportWithDefaults() *AiAgentReport {
	this := AiAgentReport{}
	return &this
}

// GetReportId returns the ReportId field value.
func (o *AiAgentReport) GetReportId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ReportId
}

// GetReportIdOk returns a tuple with the ReportId field value
// and a boolean to check if the value has been set.
func (o *AiAgentReport) GetReportIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportId, true
}

// SetReportId sets field value.
func (o *AiAgentReport) SetReportId(v string) {
	o.ReportId = v
}

// GetRunId returns the RunId field value.
func (o *AiAgentReport) GetRunId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RunId
}

// GetRunIdOk returns a tuple with the RunId field value
// and a boolean to check if the value has been set.
func (o *AiAgentReport) GetRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunId, true
}

// SetRunId sets field value.
func (o *AiAgentReport) SetRunId(v string) {
	o.RunId = v
}

// GetConversationId returns the ConversationId field value.
func (o *AiAgentReport) GetConversationId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ConversationId
}

// GetConversationIdOk returns a tuple with the ConversationId field value
// and a boolean to check if the value has been set.
func (o *AiAgentReport) GetConversationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConversationId, true
}

// SetConversationId sets field value.
func (o *AiAgentReport) SetConversationId(v string) {
	o.ConversationId = v
}

// GetStatus returns the Status field value.
func (o *AiAgentReport) GetStatus() AiAgentStatusDisplay {
	if o == nil {
		var ret AiAgentStatusDisplay
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AiAgentReport) GetStatusOk() (*AiAgentStatusDisplay, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *AiAgentReport) SetStatus(v AiAgentStatusDisplay) {
	o.Status = v
}

// GetFinalStatus returns the FinalStatus field value.
func (o *AiAgentReport) GetFinalStatus() AiAgentFinalStatus {
	if o == nil {
		var ret AiAgentFinalStatus
		return ret
	}
	return o.FinalStatus
}

// GetFinalStatusOk returns a tuple with the FinalStatus field value
// and a boolean to check if the value has been set.
func (o *AiAgentReport) GetFinalStatusOk() (*AiAgentFinalStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FinalStatus, true
}

// SetFinalStatus sets field value.
func (o *AiAgentReport) SetFinalStatus(v AiAgentFinalStatus) {
	o.FinalStatus = v
}

// GetEvidenceCompleteness returns the EvidenceCompleteness field value.
func (o *AiAgentReport) GetEvidenceCompleteness() AiAgentEvidenceCompleteness {
	if o == nil {
		var ret AiAgentEvidenceCompleteness
		return ret
	}
	return o.EvidenceCompleteness
}

// GetEvidenceCompletenessOk returns a tuple with the EvidenceCompleteness field value
// and a boolean to check if the value has been set.
func (o *AiAgentReport) GetEvidenceCompletenessOk() (*AiAgentEvidenceCompleteness, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EvidenceCompleteness, true
}

// SetEvidenceCompleteness sets field value.
func (o *AiAgentReport) SetEvidenceCompleteness(v AiAgentEvidenceCompleteness) {
	o.EvidenceCompleteness = v
}

// GetConfidence returns the Confidence field value if set, zero value otherwise.
func (o *AiAgentReport) GetConfidence() AiAgentConfidence {
	if o == nil || o.Confidence == nil {
		var ret AiAgentConfidence
		return ret
	}
	return *o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentReport) GetConfidenceOk() (*AiAgentConfidence, bool) {
	if o == nil || o.Confidence == nil {
		return nil, false
	}
	return o.Confidence, true
}

// HasConfidence returns a boolean if a field has been set.
func (o *AiAgentReport) HasConfidence() bool {
	return o != nil && o.Confidence != nil
}

// SetConfidence gets a reference to the given AiAgentConfidence and assigns it to the Confidence field.
func (o *AiAgentReport) SetConfidence(v AiAgentConfidence) {
	o.Confidence = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AiAgentReport) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentReport) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AiAgentReport) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AiAgentReport) SetTitle(v string) {
	o.Title = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *AiAgentReport) GetSummary() string {
	if o == nil || o.Summary == nil {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentReport) GetSummaryOk() (*string, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *AiAgentReport) HasSummary() bool {
	return o != nil && o.Summary != nil
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *AiAgentReport) SetSummary(v string) {
	o.Summary = &v
}

// GetRootCause returns the RootCause field value if set, zero value otherwise.
func (o *AiAgentReport) GetRootCause() AiAgentRootCause {
	if o == nil || o.RootCause == nil {
		var ret AiAgentRootCause
		return ret
	}
	return *o.RootCause
}

// GetRootCauseOk returns a tuple with the RootCause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentReport) GetRootCauseOk() (*AiAgentRootCause, bool) {
	if o == nil || o.RootCause == nil {
		return nil, false
	}
	return o.RootCause, true
}

// HasRootCause returns a boolean if a field has been set.
func (o *AiAgentReport) HasRootCause() bool {
	return o != nil && o.RootCause != nil
}

// SetRootCause gets a reference to the given AiAgentRootCause and assigns it to the RootCause field.
func (o *AiAgentReport) SetRootCause(v AiAgentRootCause) {
	o.RootCause = &v
}

// GetCheckedScope returns the CheckedScope field value if set, zero value otherwise.
func (o *AiAgentReport) GetCheckedScope() []AiAgentCheckedScopeItem {
	if o == nil || o.CheckedScope == nil {
		var ret []AiAgentCheckedScopeItem
		return ret
	}
	return o.CheckedScope
}

// GetCheckedScopeOk returns a tuple with the CheckedScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentReport) GetCheckedScopeOk() (*[]AiAgentCheckedScopeItem, bool) {
	if o == nil || o.CheckedScope == nil {
		return nil, false
	}
	return &o.CheckedScope, true
}

// HasCheckedScope returns a boolean if a field has been set.
func (o *AiAgentReport) HasCheckedScope() bool {
	return o != nil && o.CheckedScope != nil
}

// SetCheckedScope gets a reference to the given []AiAgentCheckedScopeItem and assigns it to the CheckedScope field.
func (o *AiAgentReport) SetCheckedScope(v []AiAgentCheckedScopeItem) {
	o.CheckedScope = v
}

// GetEvidenceSummary returns the EvidenceSummary field value if set, zero value otherwise.
func (o *AiAgentReport) GetEvidenceSummary() []AiAgentEvidenceSummary {
	if o == nil || o.EvidenceSummary == nil {
		var ret []AiAgentEvidenceSummary
		return ret
	}
	return o.EvidenceSummary
}

// GetEvidenceSummaryOk returns a tuple with the EvidenceSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentReport) GetEvidenceSummaryOk() (*[]AiAgentEvidenceSummary, bool) {
	if o == nil || o.EvidenceSummary == nil {
		return nil, false
	}
	return &o.EvidenceSummary, true
}

// HasEvidenceSummary returns a boolean if a field has been set.
func (o *AiAgentReport) HasEvidenceSummary() bool {
	return o != nil && o.EvidenceSummary != nil
}

// SetEvidenceSummary gets a reference to the given []AiAgentEvidenceSummary and assigns it to the EvidenceSummary field.
func (o *AiAgentReport) SetEvidenceSummary(v []AiAgentEvidenceSummary) {
	o.EvidenceSummary = v
}

// GetUncertainty returns the Uncertainty field value if set, zero value otherwise.
func (o *AiAgentReport) GetUncertainty() []AiAgentUncertainty {
	if o == nil || o.Uncertainty == nil {
		var ret []AiAgentUncertainty
		return ret
	}
	return o.Uncertainty
}

// GetUncertaintyOk returns a tuple with the Uncertainty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentReport) GetUncertaintyOk() (*[]AiAgentUncertainty, bool) {
	if o == nil || o.Uncertainty == nil {
		return nil, false
	}
	return &o.Uncertainty, true
}

// HasUncertainty returns a boolean if a field has been set.
func (o *AiAgentReport) HasUncertainty() bool {
	return o != nil && o.Uncertainty != nil
}

// SetUncertainty gets a reference to the given []AiAgentUncertainty and assigns it to the Uncertainty field.
func (o *AiAgentReport) SetUncertainty(v []AiAgentUncertainty) {
	o.Uncertainty = v
}

// GetBlockingRequirementIds returns the BlockingRequirementIds field value if set, zero value otherwise.
func (o *AiAgentReport) GetBlockingRequirementIds() []string {
	if o == nil || o.BlockingRequirementIds == nil {
		var ret []string
		return ret
	}
	return o.BlockingRequirementIds
}

// GetBlockingRequirementIdsOk returns a tuple with the BlockingRequirementIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentReport) GetBlockingRequirementIdsOk() (*[]string, bool) {
	if o == nil || o.BlockingRequirementIds == nil {
		return nil, false
	}
	return &o.BlockingRequirementIds, true
}

// HasBlockingRequirementIds returns a boolean if a field has been set.
func (o *AiAgentReport) HasBlockingRequirementIds() bool {
	return o != nil && o.BlockingRequirementIds != nil
}

// SetBlockingRequirementIds gets a reference to the given []string and assigns it to the BlockingRequirementIds field.
func (o *AiAgentReport) SetBlockingRequirementIds(v []string) {
	o.BlockingRequirementIds = v
}

// GetDegradedRequirementIds returns the DegradedRequirementIds field value if set, zero value otherwise.
func (o *AiAgentReport) GetDegradedRequirementIds() []string {
	if o == nil || o.DegradedRequirementIds == nil {
		var ret []string
		return ret
	}
	return o.DegradedRequirementIds
}

// GetDegradedRequirementIdsOk returns a tuple with the DegradedRequirementIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentReport) GetDegradedRequirementIdsOk() (*[]string, bool) {
	if o == nil || o.DegradedRequirementIds == nil {
		return nil, false
	}
	return &o.DegradedRequirementIds, true
}

// HasDegradedRequirementIds returns a boolean if a field has been set.
func (o *AiAgentReport) HasDegradedRequirementIds() bool {
	return o != nil && o.DegradedRequirementIds != nil
}

// SetDegradedRequirementIds gets a reference to the given []string and assigns it to the DegradedRequirementIds field.
func (o *AiAgentReport) SetDegradedRequirementIds(v []string) {
	o.DegradedRequirementIds = v
}

// GetPermissionBlockedRequirementIds returns the PermissionBlockedRequirementIds field value if set, zero value otherwise.
func (o *AiAgentReport) GetPermissionBlockedRequirementIds() []string {
	if o == nil || o.PermissionBlockedRequirementIds == nil {
		var ret []string
		return ret
	}
	return o.PermissionBlockedRequirementIds
}

// GetPermissionBlockedRequirementIdsOk returns a tuple with the PermissionBlockedRequirementIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentReport) GetPermissionBlockedRequirementIdsOk() (*[]string, bool) {
	if o == nil || o.PermissionBlockedRequirementIds == nil {
		return nil, false
	}
	return &o.PermissionBlockedRequirementIds, true
}

// HasPermissionBlockedRequirementIds returns a boolean if a field has been set.
func (o *AiAgentReport) HasPermissionBlockedRequirementIds() bool {
	return o != nil && o.PermissionBlockedRequirementIds != nil
}

// SetPermissionBlockedRequirementIds gets a reference to the given []string and assigns it to the PermissionBlockedRequirementIds field.
func (o *AiAgentReport) SetPermissionBlockedRequirementIds(v []string) {
	o.PermissionBlockedRequirementIds = v
}

// GetMissingReasonByRequirement returns the MissingReasonByRequirement field value if set, zero value otherwise.
func (o *AiAgentReport) GetMissingReasonByRequirement() map[string]AiAgentMissingReason {
	if o == nil || o.MissingReasonByRequirement == nil {
		var ret map[string]AiAgentMissingReason
		return ret
	}
	return o.MissingReasonByRequirement
}

// GetMissingReasonByRequirementOk returns a tuple with the MissingReasonByRequirement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentReport) GetMissingReasonByRequirementOk() (*map[string]AiAgentMissingReason, bool) {
	if o == nil || o.MissingReasonByRequirement == nil {
		return nil, false
	}
	return &o.MissingReasonByRequirement, true
}

// HasMissingReasonByRequirement returns a boolean if a field has been set.
func (o *AiAgentReport) HasMissingReasonByRequirement() bool {
	return o != nil && o.MissingReasonByRequirement != nil
}

// SetMissingReasonByRequirement gets a reference to the given map[string]AiAgentMissingReason and assigns it to the MissingReasonByRequirement field.
func (o *AiAgentReport) SetMissingReasonByRequirement(v map[string]AiAgentMissingReason) {
	o.MissingReasonByRequirement = v
}

// GetPartialReason returns the PartialReason field value if set, zero value otherwise.
func (o *AiAgentReport) GetPartialReason() AiAgentPartialReason {
	if o == nil || o.PartialReason == nil {
		var ret AiAgentPartialReason
		return ret
	}
	return *o.PartialReason
}

// GetPartialReasonOk returns a tuple with the PartialReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentReport) GetPartialReasonOk() (*AiAgentPartialReason, bool) {
	if o == nil || o.PartialReason == nil {
		return nil, false
	}
	return o.PartialReason, true
}

// HasPartialReason returns a boolean if a field has been set.
func (o *AiAgentReport) HasPartialReason() bool {
	return o != nil && o.PartialReason != nil
}

// SetPartialReason gets a reference to the given AiAgentPartialReason and assigns it to the PartialReason field.
func (o *AiAgentReport) SetPartialReason(v AiAgentPartialReason) {
	o.PartialReason = &v
}

// GetBudgetStatus returns the BudgetStatus field value if set, zero value otherwise.
func (o *AiAgentReport) GetBudgetStatus() string {
	if o == nil || o.BudgetStatus == nil {
		var ret string
		return ret
	}
	return *o.BudgetStatus
}

// GetBudgetStatusOk returns a tuple with the BudgetStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentReport) GetBudgetStatusOk() (*string, bool) {
	if o == nil || o.BudgetStatus == nil {
		return nil, false
	}
	return o.BudgetStatus, true
}

// HasBudgetStatus returns a boolean if a field has been set.
func (o *AiAgentReport) HasBudgetStatus() bool {
	return o != nil && o.BudgetStatus != nil
}

// SetBudgetStatus gets a reference to the given string and assigns it to the BudgetStatus field.
func (o *AiAgentReport) SetBudgetStatus(v string) {
	o.BudgetStatus = &v
}

// GetScopeValidationStatus returns the ScopeValidationStatus field value if set, zero value otherwise.
func (o *AiAgentReport) GetScopeValidationStatus() AiAgentScopeValidationStatus {
	if o == nil || o.ScopeValidationStatus == nil {
		var ret AiAgentScopeValidationStatus
		return ret
	}
	return *o.ScopeValidationStatus
}

// GetScopeValidationStatusOk returns a tuple with the ScopeValidationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentReport) GetScopeValidationStatusOk() (*AiAgentScopeValidationStatus, bool) {
	if o == nil || o.ScopeValidationStatus == nil {
		return nil, false
	}
	return o.ScopeValidationStatus, true
}

// HasScopeValidationStatus returns a boolean if a field has been set.
func (o *AiAgentReport) HasScopeValidationStatus() bool {
	return o != nil && o.ScopeValidationStatus != nil
}

// SetScopeValidationStatus gets a reference to the given AiAgentScopeValidationStatus and assigns it to the ScopeValidationStatus field.
func (o *AiAgentReport) SetScopeValidationStatus(v AiAgentScopeValidationStatus) {
	o.ScopeValidationStatus = &v
}

// GetApprovalRequired returns the ApprovalRequired field value if set, zero value otherwise.
func (o *AiAgentReport) GetApprovalRequired() bool {
	if o == nil || o.ApprovalRequired == nil {
		var ret bool
		return ret
	}
	return *o.ApprovalRequired
}

// GetApprovalRequiredOk returns a tuple with the ApprovalRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentReport) GetApprovalRequiredOk() (*bool, bool) {
	if o == nil || o.ApprovalRequired == nil {
		return nil, false
	}
	return o.ApprovalRequired, true
}

// HasApprovalRequired returns a boolean if a field has been set.
func (o *AiAgentReport) HasApprovalRequired() bool {
	return o != nil && o.ApprovalRequired != nil
}

// SetApprovalRequired gets a reference to the given bool and assigns it to the ApprovalRequired field.
func (o *AiAgentReport) SetApprovalRequired(v bool) {
	o.ApprovalRequired = &v
}

// GetApprovalStatus returns the ApprovalStatus field value if set, zero value otherwise.
func (o *AiAgentReport) GetApprovalStatus() AiAgentApprovalStatus {
	if o == nil || o.ApprovalStatus == nil {
		var ret AiAgentApprovalStatus
		return ret
	}
	return *o.ApprovalStatus
}

// GetApprovalStatusOk returns a tuple with the ApprovalStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentReport) GetApprovalStatusOk() (*AiAgentApprovalStatus, bool) {
	if o == nil || o.ApprovalStatus == nil {
		return nil, false
	}
	return o.ApprovalStatus, true
}

// HasApprovalStatus returns a boolean if a field has been set.
func (o *AiAgentReport) HasApprovalStatus() bool {
	return o != nil && o.ApprovalStatus != nil
}

// SetApprovalStatus gets a reference to the given AiAgentApprovalStatus and assigns it to the ApprovalStatus field.
func (o *AiAgentReport) SetApprovalStatus(v AiAgentApprovalStatus) {
	o.ApprovalStatus = &v
}

// GetRecommendations returns the Recommendations field value if set, zero value otherwise.
func (o *AiAgentReport) GetRecommendations() []AiAgentRecommendation {
	if o == nil || o.Recommendations == nil {
		var ret []AiAgentRecommendation
		return ret
	}
	return o.Recommendations
}

// GetRecommendationsOk returns a tuple with the Recommendations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentReport) GetRecommendationsOk() (*[]AiAgentRecommendation, bool) {
	if o == nil || o.Recommendations == nil {
		return nil, false
	}
	return &o.Recommendations, true
}

// HasRecommendations returns a boolean if a field has been set.
func (o *AiAgentReport) HasRecommendations() bool {
	return o != nil && o.Recommendations != nil
}

// SetRecommendations gets a reference to the given []AiAgentRecommendation and assigns it to the Recommendations field.
func (o *AiAgentReport) SetRecommendations(v []AiAgentRecommendation) {
	o.Recommendations = v
}

// GetVisualizations returns the Visualizations field value if set, zero value otherwise.
func (o *AiAgentReport) GetVisualizations() []AiAgentVisualization {
	if o == nil || o.Visualizations == nil {
		var ret []AiAgentVisualization
		return ret
	}
	return o.Visualizations
}

// GetVisualizationsOk returns a tuple with the Visualizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentReport) GetVisualizationsOk() (*[]AiAgentVisualization, bool) {
	if o == nil || o.Visualizations == nil {
		return nil, false
	}
	return &o.Visualizations, true
}

// HasVisualizations returns a boolean if a field has been set.
func (o *AiAgentReport) HasVisualizations() bool {
	return o != nil && o.Visualizations != nil
}

// SetVisualizations gets a reference to the given []AiAgentVisualization and assigns it to the Visualizations field.
func (o *AiAgentReport) SetVisualizations(v []AiAgentVisualization) {
	o.Visualizations = v
}

// GetMetricSeriesRefs returns the MetricSeriesRefs field value if set, zero value otherwise.
func (o *AiAgentReport) GetMetricSeriesRefs() []string {
	if o == nil || o.MetricSeriesRefs == nil {
		var ret []string
		return ret
	}
	return o.MetricSeriesRefs
}

// GetMetricSeriesRefsOk returns a tuple with the MetricSeriesRefs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentReport) GetMetricSeriesRefsOk() (*[]string, bool) {
	if o == nil || o.MetricSeriesRefs == nil {
		return nil, false
	}
	return &o.MetricSeriesRefs, true
}

// HasMetricSeriesRefs returns a boolean if a field has been set.
func (o *AiAgentReport) HasMetricSeriesRefs() bool {
	return o != nil && o.MetricSeriesRefs != nil
}

// SetMetricSeriesRefs gets a reference to the given []string and assigns it to the MetricSeriesRefs field.
func (o *AiAgentReport) SetMetricSeriesRefs(v []string) {
	o.MetricSeriesRefs = v
}

// GetSourceEventIds returns the SourceEventIds field value if set, zero value otherwise.
func (o *AiAgentReport) GetSourceEventIds() []string {
	if o == nil || o.SourceEventIds == nil {
		var ret []string
		return ret
	}
	return o.SourceEventIds
}

// GetSourceEventIdsOk returns a tuple with the SourceEventIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentReport) GetSourceEventIdsOk() (*[]string, bool) {
	if o == nil || o.SourceEventIds == nil {
		return nil, false
	}
	return &o.SourceEventIds, true
}

// HasSourceEventIds returns a boolean if a field has been set.
func (o *AiAgentReport) HasSourceEventIds() bool {
	return o != nil && o.SourceEventIds != nil
}

// SetSourceEventIds gets a reference to the given []string and assigns it to the SourceEventIds field.
func (o *AiAgentReport) SetSourceEventIds(v []string) {
	o.SourceEventIds = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *AiAgentReport) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AiAgentReport) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *AiAgentReport) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentReport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["reportId"] = o.ReportId
	toSerialize["runId"] = o.RunId
	toSerialize["conversationId"] = o.ConversationId
	toSerialize["status"] = o.Status
	toSerialize["finalStatus"] = o.FinalStatus
	toSerialize["evidenceCompleteness"] = o.EvidenceCompleteness
	if o.Confidence != nil {
		toSerialize["confidence"] = o.Confidence
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Summary != nil {
		toSerialize["summary"] = o.Summary
	}
	if o.RootCause != nil {
		toSerialize["rootCause"] = o.RootCause
	}
	if o.CheckedScope != nil {
		toSerialize["checkedScope"] = o.CheckedScope
	}
	if o.EvidenceSummary != nil {
		toSerialize["evidenceSummary"] = o.EvidenceSummary
	}
	if o.Uncertainty != nil {
		toSerialize["uncertainty"] = o.Uncertainty
	}
	if o.BlockingRequirementIds != nil {
		toSerialize["blockingRequirementIds"] = o.BlockingRequirementIds
	}
	if o.DegradedRequirementIds != nil {
		toSerialize["degradedRequirementIds"] = o.DegradedRequirementIds
	}
	if o.PermissionBlockedRequirementIds != nil {
		toSerialize["permissionBlockedRequirementIds"] = o.PermissionBlockedRequirementIds
	}
	if o.MissingReasonByRequirement != nil {
		toSerialize["missingReasonByRequirement"] = o.MissingReasonByRequirement
	}
	if o.PartialReason != nil {
		toSerialize["partialReason"] = o.PartialReason
	}
	if o.BudgetStatus != nil {
		toSerialize["budgetStatus"] = o.BudgetStatus
	}
	if o.ScopeValidationStatus != nil {
		toSerialize["scopeValidationStatus"] = o.ScopeValidationStatus
	}
	if o.ApprovalRequired != nil {
		toSerialize["approvalRequired"] = o.ApprovalRequired
	}
	if o.ApprovalStatus != nil {
		toSerialize["approvalStatus"] = o.ApprovalStatus
	}
	if o.Recommendations != nil {
		toSerialize["recommendations"] = o.Recommendations
	}
	if o.Visualizations != nil {
		toSerialize["visualizations"] = o.Visualizations
	}
	if o.MetricSeriesRefs != nil {
		toSerialize["metricSeriesRefs"] = o.MetricSeriesRefs
	}
	if o.SourceEventIds != nil {
		toSerialize["sourceEventIds"] = o.SourceEventIds
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
func (o *AiAgentReport) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ReportId                        *string                         `json:"reportId"`
		RunId                           *string                         `json:"runId"`
		ConversationId                  *string                         `json:"conversationId"`
		Status                          *AiAgentStatusDisplay           `json:"status"`
		FinalStatus                     *AiAgentFinalStatus             `json:"finalStatus"`
		EvidenceCompleteness            *AiAgentEvidenceCompleteness    `json:"evidenceCompleteness"`
		Confidence                      *AiAgentConfidence              `json:"confidence,omitempty"`
		Title                           *string                         `json:"title,omitempty"`
		Summary                         *string                         `json:"summary,omitempty"`
		RootCause                       *AiAgentRootCause               `json:"rootCause,omitempty"`
		CheckedScope                    []AiAgentCheckedScopeItem       `json:"checkedScope,omitempty"`
		EvidenceSummary                 []AiAgentEvidenceSummary        `json:"evidenceSummary,omitempty"`
		Uncertainty                     []AiAgentUncertainty            `json:"uncertainty,omitempty"`
		BlockingRequirementIds          []string                        `json:"blockingRequirementIds,omitempty"`
		DegradedRequirementIds          []string                        `json:"degradedRequirementIds,omitempty"`
		PermissionBlockedRequirementIds []string                        `json:"permissionBlockedRequirementIds,omitempty"`
		MissingReasonByRequirement      map[string]AiAgentMissingReason `json:"missingReasonByRequirement,omitempty"`
		PartialReason                   *AiAgentPartialReason           `json:"partialReason,omitempty"`
		BudgetStatus                    *string                         `json:"budgetStatus,omitempty"`
		ScopeValidationStatus           *AiAgentScopeValidationStatus   `json:"scopeValidationStatus,omitempty"`
		ApprovalRequired                *bool                           `json:"approvalRequired,omitempty"`
		ApprovalStatus                  *AiAgentApprovalStatus          `json:"approvalStatus,omitempty"`
		Recommendations                 []AiAgentRecommendation         `json:"recommendations,omitempty"`
		Visualizations                  []AiAgentVisualization          `json:"visualizations,omitempty"`
		MetricSeriesRefs                []string                        `json:"metricSeriesRefs,omitempty"`
		SourceEventIds                  []string                        `json:"sourceEventIds,omitempty"`
		CreatedAt                       *time.Time                      `json:"createdAt"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.ReportId == nil {
		return fmt.Errorf("required field reportId missing")
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
	if all.FinalStatus == nil {
		return fmt.Errorf("required field finalStatus missing")
	}
	if all.EvidenceCompleteness == nil {
		return fmt.Errorf("required field evidenceCompleteness missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field createdAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"reportId", "runId", "conversationId", "status", "finalStatus", "evidenceCompleteness", "confidence", "title", "summary", "rootCause", "checkedScope", "evidenceSummary", "uncertainty", "blockingRequirementIds", "degradedRequirementIds", "permissionBlockedRequirementIds", "missingReasonByRequirement", "partialReason", "budgetStatus", "scopeValidationStatus", "approvalRequired", "approvalStatus", "recommendations", "visualizations", "metricSeriesRefs", "sourceEventIds", "createdAt"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ReportId = *all.ReportId
	o.RunId = *all.RunId
	o.ConversationId = *all.ConversationId
	if all.Status.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Status = *all.Status
	if !all.FinalStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.FinalStatus = *all.FinalStatus
	}
	if !all.EvidenceCompleteness.IsValid() {
		hasInvalidField = true
	} else {
		o.EvidenceCompleteness = *all.EvidenceCompleteness
	}
	if all.Confidence != nil && !all.Confidence.IsValid() {
		hasInvalidField = true
	} else {
		o.Confidence = all.Confidence
	}
	o.Title = all.Title
	o.Summary = all.Summary
	if all.RootCause != nil && all.RootCause.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.RootCause = all.RootCause
	o.CheckedScope = all.CheckedScope
	o.EvidenceSummary = all.EvidenceSummary
	o.Uncertainty = all.Uncertainty
	o.BlockingRequirementIds = all.BlockingRequirementIds
	o.DegradedRequirementIds = all.DegradedRequirementIds
	o.PermissionBlockedRequirementIds = all.PermissionBlockedRequirementIds
	o.MissingReasonByRequirement = all.MissingReasonByRequirement
	if all.PartialReason != nil && !all.PartialReason.IsValid() {
		hasInvalidField = true
	} else {
		o.PartialReason = all.PartialReason
	}
	o.BudgetStatus = all.BudgetStatus
	if all.ScopeValidationStatus != nil && !all.ScopeValidationStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.ScopeValidationStatus = all.ScopeValidationStatus
	}
	o.ApprovalRequired = all.ApprovalRequired
	if all.ApprovalStatus != nil && !all.ApprovalStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.ApprovalStatus = all.ApprovalStatus
	}
	o.Recommendations = all.Recommendations
	o.Visualizations = all.Visualizations
	o.MetricSeriesRefs = all.MetricSeriesRefs
	o.SourceEventIds = all.SourceEventIds
	o.CreatedAt = *all.CreatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
