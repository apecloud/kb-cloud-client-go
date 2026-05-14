// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// AiAgentEventPayload Frontend-safe union payload for typed Agent events. For eventType=assistant_delta, messageId, partId, deltaKind, sequence, delta, and isFinal are required by contract; sequence is monotonically increasing per partId from 1 and is used by clients to order and deduplicate streamed deltas. Debug or hidden projections must not be sent on the ordinary user SSE stream.
type AiAgentEventPayload struct {
	MessageId *string           `json:"messageId,omitempty"`
	PartId    *string           `json:"partId,omitempty"`
	DeltaKind *AiAgentDeltaKind `json:"deltaKind,omitempty"`
	// Monotonic sequence number per message part, starting at 1 for each partId.
	Sequence *int64  `json:"sequence,omitempty"`
	Delta    *string `json:"delta,omitempty"`
	IsFinal  *bool   `json:"isFinal,omitempty"`
	StageId  *string `json:"stageId,omitempty"`
	// Stable status code plus default user-facing copy. Frontend uses code for logic/tests and may override display text for i18n.
	State                 *AiAgentStatusDisplay        `json:"state,omitempty"`
	DisplayLabel          *string                      `json:"displayLabel,omitempty"`
	DisplayDescription    *string                      `json:"displayDescription,omitempty"`
	UiVisibility          *AiAgentUIVisibility         `json:"uiVisibility,omitempty"`
	ToolCallId            *string                      `json:"toolCallId,omitempty"`
	AgentToolId           *string                      `json:"agentToolId,omitempty"`
	RequirementIds        []string                     `json:"requirementIds,omitempty"`
	EvidenceIds           []string                     `json:"evidenceIds,omitempty"`
	MissingRequirements   []AiAgentMissingRequirement  `json:"missingRequirements,omitempty"`
	ToolResponseStatus    *AiAgentToolResponseStatus   `json:"toolResponseStatus,omitempty"`
	MissingReason         *AiAgentMissingReason        `json:"missingReason,omitempty"`
	PartialReason         *AiAgentPartialReason        `json:"partialReason,omitempty"`
	StatusReason          *string                      `json:"statusReason,omitempty"`
	PrimaryMissingReason  *string                      `json:"primaryMissingReason,omitempty"`
	EvidenceCompleteness  *AiAgentEvidenceCompleteness `json:"evidenceCompleteness,omitempty"`
	Summary               *string                      `json:"summary,omitempty"`
	ReportId              *string                      `json:"reportId,omitempty"`
	ReportUrl             *string                      `json:"reportUrl,omitempty"`
	ActionPlanId          *string                      `json:"actionPlanId,omitempty"`
	Title                 *string                      `json:"title,omitempty"`
	RiskLevel             *string                      `json:"riskLevel,omitempty"`
	ImpactSummary         *string                      `json:"impactSummary,omitempty"`
	Question              *string                      `json:"question,omitempty"`
	RequiredContextFields []string                     `json:"requiredContextFields,omitempty"`
	SuggestedReplies      []AiAgentSuggestedReply      `json:"suggestedReplies,omitempty"`
	RecoverActions        []AiAgentRecoverAction       `json:"recoverActions,omitempty"`
	CostClass             *AiAgentCostClass            `json:"costClass,omitempty"`
	BudgetStatus          *string                      `json:"budgetStatus,omitempty"`
	RemainingBudgetClass  *AiAgentCostClass            `json:"remainingBudgetClass,omitempty"`
	Reason                *string                      `json:"reason,omitempty"`
	PlanTitle             *string                      `json:"planTitle,omitempty"`
	PlanSteps             []string                     `json:"planSteps,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentEventPayload instantiates a new AiAgentEventPayload object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentEventPayload() *AiAgentEventPayload {
	this := AiAgentEventPayload{}
	return &this
}

// NewAiAgentEventPayloadWithDefaults instantiates a new AiAgentEventPayload object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentEventPayloadWithDefaults() *AiAgentEventPayload {
	this := AiAgentEventPayload{}
	return &this
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *AiAgentEventPayload) GetMessageId() string {
	if o == nil || o.MessageId == nil {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEventPayload) GetMessageIdOk() (*string, bool) {
	if o == nil || o.MessageId == nil {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *AiAgentEventPayload) HasMessageId() bool {
	return o != nil && o.MessageId != nil
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *AiAgentEventPayload) SetMessageId(v string) {
	o.MessageId = &v
}

// GetPartId returns the PartId field value if set, zero value otherwise.
func (o *AiAgentEventPayload) GetPartId() string {
	if o == nil || o.PartId == nil {
		var ret string
		return ret
	}
	return *o.PartId
}

// GetPartIdOk returns a tuple with the PartId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEventPayload) GetPartIdOk() (*string, bool) {
	if o == nil || o.PartId == nil {
		return nil, false
	}
	return o.PartId, true
}

// HasPartId returns a boolean if a field has been set.
func (o *AiAgentEventPayload) HasPartId() bool {
	return o != nil && o.PartId != nil
}

// SetPartId gets a reference to the given string and assigns it to the PartId field.
func (o *AiAgentEventPayload) SetPartId(v string) {
	o.PartId = &v
}

// GetDeltaKind returns the DeltaKind field value if set, zero value otherwise.
func (o *AiAgentEventPayload) GetDeltaKind() AiAgentDeltaKind {
	if o == nil || o.DeltaKind == nil {
		var ret AiAgentDeltaKind
		return ret
	}
	return *o.DeltaKind
}

// GetDeltaKindOk returns a tuple with the DeltaKind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEventPayload) GetDeltaKindOk() (*AiAgentDeltaKind, bool) {
	if o == nil || o.DeltaKind == nil {
		return nil, false
	}
	return o.DeltaKind, true
}

// HasDeltaKind returns a boolean if a field has been set.
func (o *AiAgentEventPayload) HasDeltaKind() bool {
	return o != nil && o.DeltaKind != nil
}

// SetDeltaKind gets a reference to the given AiAgentDeltaKind and assigns it to the DeltaKind field.
func (o *AiAgentEventPayload) SetDeltaKind(v AiAgentDeltaKind) {
	o.DeltaKind = &v
}

// GetSequence returns the Sequence field value if set, zero value otherwise.
func (o *AiAgentEventPayload) GetSequence() int64 {
	if o == nil || o.Sequence == nil {
		var ret int64
		return ret
	}
	return *o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEventPayload) GetSequenceOk() (*int64, bool) {
	if o == nil || o.Sequence == nil {
		return nil, false
	}
	return o.Sequence, true
}

// HasSequence returns a boolean if a field has been set.
func (o *AiAgentEventPayload) HasSequence() bool {
	return o != nil && o.Sequence != nil
}

// SetSequence gets a reference to the given int64 and assigns it to the Sequence field.
func (o *AiAgentEventPayload) SetSequence(v int64) {
	o.Sequence = &v
}

// GetDelta returns the Delta field value if set, zero value otherwise.
func (o *AiAgentEventPayload) GetDelta() string {
	if o == nil || o.Delta == nil {
		var ret string
		return ret
	}
	return *o.Delta
}

// GetDeltaOk returns a tuple with the Delta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEventPayload) GetDeltaOk() (*string, bool) {
	if o == nil || o.Delta == nil {
		return nil, false
	}
	return o.Delta, true
}

// HasDelta returns a boolean if a field has been set.
func (o *AiAgentEventPayload) HasDelta() bool {
	return o != nil && o.Delta != nil
}

// SetDelta gets a reference to the given string and assigns it to the Delta field.
func (o *AiAgentEventPayload) SetDelta(v string) {
	o.Delta = &v
}

// GetIsFinal returns the IsFinal field value if set, zero value otherwise.
func (o *AiAgentEventPayload) GetIsFinal() bool {
	if o == nil || o.IsFinal == nil {
		var ret bool
		return ret
	}
	return *o.IsFinal
}

// GetIsFinalOk returns a tuple with the IsFinal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEventPayload) GetIsFinalOk() (*bool, bool) {
	if o == nil || o.IsFinal == nil {
		return nil, false
	}
	return o.IsFinal, true
}

// HasIsFinal returns a boolean if a field has been set.
func (o *AiAgentEventPayload) HasIsFinal() bool {
	return o != nil && o.IsFinal != nil
}

// SetIsFinal gets a reference to the given bool and assigns it to the IsFinal field.
func (o *AiAgentEventPayload) SetIsFinal(v bool) {
	o.IsFinal = &v
}

// GetStageId returns the StageId field value if set, zero value otherwise.
func (o *AiAgentEventPayload) GetStageId() string {
	if o == nil || o.StageId == nil {
		var ret string
		return ret
	}
	return *o.StageId
}

// GetStageIdOk returns a tuple with the StageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEventPayload) GetStageIdOk() (*string, bool) {
	if o == nil || o.StageId == nil {
		return nil, false
	}
	return o.StageId, true
}

// HasStageId returns a boolean if a field has been set.
func (o *AiAgentEventPayload) HasStageId() bool {
	return o != nil && o.StageId != nil
}

// SetStageId gets a reference to the given string and assigns it to the StageId field.
func (o *AiAgentEventPayload) SetStageId(v string) {
	o.StageId = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *AiAgentEventPayload) GetState() AiAgentStatusDisplay {
	if o == nil || o.State == nil {
		var ret AiAgentStatusDisplay
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEventPayload) GetStateOk() (*AiAgentStatusDisplay, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *AiAgentEventPayload) HasState() bool {
	return o != nil && o.State != nil
}

// SetState gets a reference to the given AiAgentStatusDisplay and assigns it to the State field.
func (o *AiAgentEventPayload) SetState(v AiAgentStatusDisplay) {
	o.State = &v
}

// GetDisplayLabel returns the DisplayLabel field value if set, zero value otherwise.
func (o *AiAgentEventPayload) GetDisplayLabel() string {
	if o == nil || o.DisplayLabel == nil {
		var ret string
		return ret
	}
	return *o.DisplayLabel
}

// GetDisplayLabelOk returns a tuple with the DisplayLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEventPayload) GetDisplayLabelOk() (*string, bool) {
	if o == nil || o.DisplayLabel == nil {
		return nil, false
	}
	return o.DisplayLabel, true
}

// HasDisplayLabel returns a boolean if a field has been set.
func (o *AiAgentEventPayload) HasDisplayLabel() bool {
	return o != nil && o.DisplayLabel != nil
}

// SetDisplayLabel gets a reference to the given string and assigns it to the DisplayLabel field.
func (o *AiAgentEventPayload) SetDisplayLabel(v string) {
	o.DisplayLabel = &v
}

// GetDisplayDescription returns the DisplayDescription field value if set, zero value otherwise.
func (o *AiAgentEventPayload) GetDisplayDescription() string {
	if o == nil || o.DisplayDescription == nil {
		var ret string
		return ret
	}
	return *o.DisplayDescription
}

// GetDisplayDescriptionOk returns a tuple with the DisplayDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEventPayload) GetDisplayDescriptionOk() (*string, bool) {
	if o == nil || o.DisplayDescription == nil {
		return nil, false
	}
	return o.DisplayDescription, true
}

// HasDisplayDescription returns a boolean if a field has been set.
func (o *AiAgentEventPayload) HasDisplayDescription() bool {
	return o != nil && o.DisplayDescription != nil
}

// SetDisplayDescription gets a reference to the given string and assigns it to the DisplayDescription field.
func (o *AiAgentEventPayload) SetDisplayDescription(v string) {
	o.DisplayDescription = &v
}

// GetUiVisibility returns the UiVisibility field value if set, zero value otherwise.
func (o *AiAgentEventPayload) GetUiVisibility() AiAgentUIVisibility {
	if o == nil || o.UiVisibility == nil {
		var ret AiAgentUIVisibility
		return ret
	}
	return *o.UiVisibility
}

// GetUiVisibilityOk returns a tuple with the UiVisibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEventPayload) GetUiVisibilityOk() (*AiAgentUIVisibility, bool) {
	if o == nil || o.UiVisibility == nil {
		return nil, false
	}
	return o.UiVisibility, true
}

// HasUiVisibility returns a boolean if a field has been set.
func (o *AiAgentEventPayload) HasUiVisibility() bool {
	return o != nil && o.UiVisibility != nil
}

// SetUiVisibility gets a reference to the given AiAgentUIVisibility and assigns it to the UiVisibility field.
func (o *AiAgentEventPayload) SetUiVisibility(v AiAgentUIVisibility) {
	o.UiVisibility = &v
}

// GetToolCallId returns the ToolCallId field value if set, zero value otherwise.
func (o *AiAgentEventPayload) GetToolCallId() string {
	if o == nil || o.ToolCallId == nil {
		var ret string
		return ret
	}
	return *o.ToolCallId
}

// GetToolCallIdOk returns a tuple with the ToolCallId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEventPayload) GetToolCallIdOk() (*string, bool) {
	if o == nil || o.ToolCallId == nil {
		return nil, false
	}
	return o.ToolCallId, true
}

// HasToolCallId returns a boolean if a field has been set.
func (o *AiAgentEventPayload) HasToolCallId() bool {
	return o != nil && o.ToolCallId != nil
}

// SetToolCallId gets a reference to the given string and assigns it to the ToolCallId field.
func (o *AiAgentEventPayload) SetToolCallId(v string) {
	o.ToolCallId = &v
}

// GetAgentToolId returns the AgentToolId field value if set, zero value otherwise.
func (o *AiAgentEventPayload) GetAgentToolId() string {
	if o == nil || o.AgentToolId == nil {
		var ret string
		return ret
	}
	return *o.AgentToolId
}

// GetAgentToolIdOk returns a tuple with the AgentToolId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEventPayload) GetAgentToolIdOk() (*string, bool) {
	if o == nil || o.AgentToolId == nil {
		return nil, false
	}
	return o.AgentToolId, true
}

// HasAgentToolId returns a boolean if a field has been set.
func (o *AiAgentEventPayload) HasAgentToolId() bool {
	return o != nil && o.AgentToolId != nil
}

// SetAgentToolId gets a reference to the given string and assigns it to the AgentToolId field.
func (o *AiAgentEventPayload) SetAgentToolId(v string) {
	o.AgentToolId = &v
}

// GetRequirementIds returns the RequirementIds field value if set, zero value otherwise.
func (o *AiAgentEventPayload) GetRequirementIds() []string {
	if o == nil || o.RequirementIds == nil {
		var ret []string
		return ret
	}
	return o.RequirementIds
}

// GetRequirementIdsOk returns a tuple with the RequirementIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEventPayload) GetRequirementIdsOk() (*[]string, bool) {
	if o == nil || o.RequirementIds == nil {
		return nil, false
	}
	return &o.RequirementIds, true
}

// HasRequirementIds returns a boolean if a field has been set.
func (o *AiAgentEventPayload) HasRequirementIds() bool {
	return o != nil && o.RequirementIds != nil
}

// SetRequirementIds gets a reference to the given []string and assigns it to the RequirementIds field.
func (o *AiAgentEventPayload) SetRequirementIds(v []string) {
	o.RequirementIds = v
}

// GetEvidenceIds returns the EvidenceIds field value if set, zero value otherwise.
func (o *AiAgentEventPayload) GetEvidenceIds() []string {
	if o == nil || o.EvidenceIds == nil {
		var ret []string
		return ret
	}
	return o.EvidenceIds
}

// GetEvidenceIdsOk returns a tuple with the EvidenceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEventPayload) GetEvidenceIdsOk() (*[]string, bool) {
	if o == nil || o.EvidenceIds == nil {
		return nil, false
	}
	return &o.EvidenceIds, true
}

// HasEvidenceIds returns a boolean if a field has been set.
func (o *AiAgentEventPayload) HasEvidenceIds() bool {
	return o != nil && o.EvidenceIds != nil
}

// SetEvidenceIds gets a reference to the given []string and assigns it to the EvidenceIds field.
func (o *AiAgentEventPayload) SetEvidenceIds(v []string) {
	o.EvidenceIds = v
}

// GetMissingRequirements returns the MissingRequirements field value if set, zero value otherwise.
func (o *AiAgentEventPayload) GetMissingRequirements() []AiAgentMissingRequirement {
	if o == nil || o.MissingRequirements == nil {
		var ret []AiAgentMissingRequirement
		return ret
	}
	return o.MissingRequirements
}

// GetMissingRequirementsOk returns a tuple with the MissingRequirements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEventPayload) GetMissingRequirementsOk() (*[]AiAgentMissingRequirement, bool) {
	if o == nil || o.MissingRequirements == nil {
		return nil, false
	}
	return &o.MissingRequirements, true
}

// HasMissingRequirements returns a boolean if a field has been set.
func (o *AiAgentEventPayload) HasMissingRequirements() bool {
	return o != nil && o.MissingRequirements != nil
}

// SetMissingRequirements gets a reference to the given []AiAgentMissingRequirement and assigns it to the MissingRequirements field.
func (o *AiAgentEventPayload) SetMissingRequirements(v []AiAgentMissingRequirement) {
	o.MissingRequirements = v
}

// GetToolResponseStatus returns the ToolResponseStatus field value if set, zero value otherwise.
func (o *AiAgentEventPayload) GetToolResponseStatus() AiAgentToolResponseStatus {
	if o == nil || o.ToolResponseStatus == nil {
		var ret AiAgentToolResponseStatus
		return ret
	}
	return *o.ToolResponseStatus
}

// GetToolResponseStatusOk returns a tuple with the ToolResponseStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEventPayload) GetToolResponseStatusOk() (*AiAgentToolResponseStatus, bool) {
	if o == nil || o.ToolResponseStatus == nil {
		return nil, false
	}
	return o.ToolResponseStatus, true
}

// HasToolResponseStatus returns a boolean if a field has been set.
func (o *AiAgentEventPayload) HasToolResponseStatus() bool {
	return o != nil && o.ToolResponseStatus != nil
}

// SetToolResponseStatus gets a reference to the given AiAgentToolResponseStatus and assigns it to the ToolResponseStatus field.
func (o *AiAgentEventPayload) SetToolResponseStatus(v AiAgentToolResponseStatus) {
	o.ToolResponseStatus = &v
}

// GetMissingReason returns the MissingReason field value if set, zero value otherwise.
func (o *AiAgentEventPayload) GetMissingReason() AiAgentMissingReason {
	if o == nil || o.MissingReason == nil {
		var ret AiAgentMissingReason
		return ret
	}
	return *o.MissingReason
}

// GetMissingReasonOk returns a tuple with the MissingReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEventPayload) GetMissingReasonOk() (*AiAgentMissingReason, bool) {
	if o == nil || o.MissingReason == nil {
		return nil, false
	}
	return o.MissingReason, true
}

// HasMissingReason returns a boolean if a field has been set.
func (o *AiAgentEventPayload) HasMissingReason() bool {
	return o != nil && o.MissingReason != nil
}

// SetMissingReason gets a reference to the given AiAgentMissingReason and assigns it to the MissingReason field.
func (o *AiAgentEventPayload) SetMissingReason(v AiAgentMissingReason) {
	o.MissingReason = &v
}

// GetPartialReason returns the PartialReason field value if set, zero value otherwise.
func (o *AiAgentEventPayload) GetPartialReason() AiAgentPartialReason {
	if o == nil || o.PartialReason == nil {
		var ret AiAgentPartialReason
		return ret
	}
	return *o.PartialReason
}

// GetPartialReasonOk returns a tuple with the PartialReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEventPayload) GetPartialReasonOk() (*AiAgentPartialReason, bool) {
	if o == nil || o.PartialReason == nil {
		return nil, false
	}
	return o.PartialReason, true
}

// HasPartialReason returns a boolean if a field has been set.
func (o *AiAgentEventPayload) HasPartialReason() bool {
	return o != nil && o.PartialReason != nil
}

// SetPartialReason gets a reference to the given AiAgentPartialReason and assigns it to the PartialReason field.
func (o *AiAgentEventPayload) SetPartialReason(v AiAgentPartialReason) {
	o.PartialReason = &v
}

// GetStatusReason returns the StatusReason field value if set, zero value otherwise.
func (o *AiAgentEventPayload) GetStatusReason() string {
	if o == nil || o.StatusReason == nil {
		var ret string
		return ret
	}
	return *o.StatusReason
}

// GetStatusReasonOk returns a tuple with the StatusReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEventPayload) GetStatusReasonOk() (*string, bool) {
	if o == nil || o.StatusReason == nil {
		return nil, false
	}
	return o.StatusReason, true
}

// HasStatusReason returns a boolean if a field has been set.
func (o *AiAgentEventPayload) HasStatusReason() bool {
	return o != nil && o.StatusReason != nil
}

// SetStatusReason gets a reference to the given string and assigns it to the StatusReason field.
func (o *AiAgentEventPayload) SetStatusReason(v string) {
	o.StatusReason = &v
}

// GetPrimaryMissingReason returns the PrimaryMissingReason field value if set, zero value otherwise.
func (o *AiAgentEventPayload) GetPrimaryMissingReason() string {
	if o == nil || o.PrimaryMissingReason == nil {
		var ret string
		return ret
	}
	return *o.PrimaryMissingReason
}

// GetPrimaryMissingReasonOk returns a tuple with the PrimaryMissingReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEventPayload) GetPrimaryMissingReasonOk() (*string, bool) {
	if o == nil || o.PrimaryMissingReason == nil {
		return nil, false
	}
	return o.PrimaryMissingReason, true
}

// HasPrimaryMissingReason returns a boolean if a field has been set.
func (o *AiAgentEventPayload) HasPrimaryMissingReason() bool {
	return o != nil && o.PrimaryMissingReason != nil
}

// SetPrimaryMissingReason gets a reference to the given string and assigns it to the PrimaryMissingReason field.
func (o *AiAgentEventPayload) SetPrimaryMissingReason(v string) {
	o.PrimaryMissingReason = &v
}

// GetEvidenceCompleteness returns the EvidenceCompleteness field value if set, zero value otherwise.
func (o *AiAgentEventPayload) GetEvidenceCompleteness() AiAgentEvidenceCompleteness {
	if o == nil || o.EvidenceCompleteness == nil {
		var ret AiAgentEvidenceCompleteness
		return ret
	}
	return *o.EvidenceCompleteness
}

// GetEvidenceCompletenessOk returns a tuple with the EvidenceCompleteness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEventPayload) GetEvidenceCompletenessOk() (*AiAgentEvidenceCompleteness, bool) {
	if o == nil || o.EvidenceCompleteness == nil {
		return nil, false
	}
	return o.EvidenceCompleteness, true
}

// HasEvidenceCompleteness returns a boolean if a field has been set.
func (o *AiAgentEventPayload) HasEvidenceCompleteness() bool {
	return o != nil && o.EvidenceCompleteness != nil
}

// SetEvidenceCompleteness gets a reference to the given AiAgentEvidenceCompleteness and assigns it to the EvidenceCompleteness field.
func (o *AiAgentEventPayload) SetEvidenceCompleteness(v AiAgentEvidenceCompleteness) {
	o.EvidenceCompleteness = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *AiAgentEventPayload) GetSummary() string {
	if o == nil || o.Summary == nil {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEventPayload) GetSummaryOk() (*string, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *AiAgentEventPayload) HasSummary() bool {
	return o != nil && o.Summary != nil
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *AiAgentEventPayload) SetSummary(v string) {
	o.Summary = &v
}

// GetReportId returns the ReportId field value if set, zero value otherwise.
func (o *AiAgentEventPayload) GetReportId() string {
	if o == nil || o.ReportId == nil {
		var ret string
		return ret
	}
	return *o.ReportId
}

// GetReportIdOk returns a tuple with the ReportId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEventPayload) GetReportIdOk() (*string, bool) {
	if o == nil || o.ReportId == nil {
		return nil, false
	}
	return o.ReportId, true
}

// HasReportId returns a boolean if a field has been set.
func (o *AiAgentEventPayload) HasReportId() bool {
	return o != nil && o.ReportId != nil
}

// SetReportId gets a reference to the given string and assigns it to the ReportId field.
func (o *AiAgentEventPayload) SetReportId(v string) {
	o.ReportId = &v
}

// GetReportUrl returns the ReportUrl field value if set, zero value otherwise.
func (o *AiAgentEventPayload) GetReportUrl() string {
	if o == nil || o.ReportUrl == nil {
		var ret string
		return ret
	}
	return *o.ReportUrl
}

// GetReportUrlOk returns a tuple with the ReportUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEventPayload) GetReportUrlOk() (*string, bool) {
	if o == nil || o.ReportUrl == nil {
		return nil, false
	}
	return o.ReportUrl, true
}

// HasReportUrl returns a boolean if a field has been set.
func (o *AiAgentEventPayload) HasReportUrl() bool {
	return o != nil && o.ReportUrl != nil
}

// SetReportUrl gets a reference to the given string and assigns it to the ReportUrl field.
func (o *AiAgentEventPayload) SetReportUrl(v string) {
	o.ReportUrl = &v
}

// GetActionPlanId returns the ActionPlanId field value if set, zero value otherwise.
func (o *AiAgentEventPayload) GetActionPlanId() string {
	if o == nil || o.ActionPlanId == nil {
		var ret string
		return ret
	}
	return *o.ActionPlanId
}

// GetActionPlanIdOk returns a tuple with the ActionPlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEventPayload) GetActionPlanIdOk() (*string, bool) {
	if o == nil || o.ActionPlanId == nil {
		return nil, false
	}
	return o.ActionPlanId, true
}

// HasActionPlanId returns a boolean if a field has been set.
func (o *AiAgentEventPayload) HasActionPlanId() bool {
	return o != nil && o.ActionPlanId != nil
}

// SetActionPlanId gets a reference to the given string and assigns it to the ActionPlanId field.
func (o *AiAgentEventPayload) SetActionPlanId(v string) {
	o.ActionPlanId = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AiAgentEventPayload) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEventPayload) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AiAgentEventPayload) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AiAgentEventPayload) SetTitle(v string) {
	o.Title = &v
}

// GetRiskLevel returns the RiskLevel field value if set, zero value otherwise.
func (o *AiAgentEventPayload) GetRiskLevel() string {
	if o == nil || o.RiskLevel == nil {
		var ret string
		return ret
	}
	return *o.RiskLevel
}

// GetRiskLevelOk returns a tuple with the RiskLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEventPayload) GetRiskLevelOk() (*string, bool) {
	if o == nil || o.RiskLevel == nil {
		return nil, false
	}
	return o.RiskLevel, true
}

// HasRiskLevel returns a boolean if a field has been set.
func (o *AiAgentEventPayload) HasRiskLevel() bool {
	return o != nil && o.RiskLevel != nil
}

// SetRiskLevel gets a reference to the given string and assigns it to the RiskLevel field.
func (o *AiAgentEventPayload) SetRiskLevel(v string) {
	o.RiskLevel = &v
}

// GetImpactSummary returns the ImpactSummary field value if set, zero value otherwise.
func (o *AiAgentEventPayload) GetImpactSummary() string {
	if o == nil || o.ImpactSummary == nil {
		var ret string
		return ret
	}
	return *o.ImpactSummary
}

// GetImpactSummaryOk returns a tuple with the ImpactSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEventPayload) GetImpactSummaryOk() (*string, bool) {
	if o == nil || o.ImpactSummary == nil {
		return nil, false
	}
	return o.ImpactSummary, true
}

// HasImpactSummary returns a boolean if a field has been set.
func (o *AiAgentEventPayload) HasImpactSummary() bool {
	return o != nil && o.ImpactSummary != nil
}

// SetImpactSummary gets a reference to the given string and assigns it to the ImpactSummary field.
func (o *AiAgentEventPayload) SetImpactSummary(v string) {
	o.ImpactSummary = &v
}

// GetQuestion returns the Question field value if set, zero value otherwise.
func (o *AiAgentEventPayload) GetQuestion() string {
	if o == nil || o.Question == nil {
		var ret string
		return ret
	}
	return *o.Question
}

// GetQuestionOk returns a tuple with the Question field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEventPayload) GetQuestionOk() (*string, bool) {
	if o == nil || o.Question == nil {
		return nil, false
	}
	return o.Question, true
}

// HasQuestion returns a boolean if a field has been set.
func (o *AiAgentEventPayload) HasQuestion() bool {
	return o != nil && o.Question != nil
}

// SetQuestion gets a reference to the given string and assigns it to the Question field.
func (o *AiAgentEventPayload) SetQuestion(v string) {
	o.Question = &v
}

// GetRequiredContextFields returns the RequiredContextFields field value if set, zero value otherwise.
func (o *AiAgentEventPayload) GetRequiredContextFields() []string {
	if o == nil || o.RequiredContextFields == nil {
		var ret []string
		return ret
	}
	return o.RequiredContextFields
}

// GetRequiredContextFieldsOk returns a tuple with the RequiredContextFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEventPayload) GetRequiredContextFieldsOk() (*[]string, bool) {
	if o == nil || o.RequiredContextFields == nil {
		return nil, false
	}
	return &o.RequiredContextFields, true
}

// HasRequiredContextFields returns a boolean if a field has been set.
func (o *AiAgentEventPayload) HasRequiredContextFields() bool {
	return o != nil && o.RequiredContextFields != nil
}

// SetRequiredContextFields gets a reference to the given []string and assigns it to the RequiredContextFields field.
func (o *AiAgentEventPayload) SetRequiredContextFields(v []string) {
	o.RequiredContextFields = v
}

// GetSuggestedReplies returns the SuggestedReplies field value if set, zero value otherwise.
func (o *AiAgentEventPayload) GetSuggestedReplies() []AiAgentSuggestedReply {
	if o == nil || o.SuggestedReplies == nil {
		var ret []AiAgentSuggestedReply
		return ret
	}
	return o.SuggestedReplies
}

// GetSuggestedRepliesOk returns a tuple with the SuggestedReplies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEventPayload) GetSuggestedRepliesOk() (*[]AiAgentSuggestedReply, bool) {
	if o == nil || o.SuggestedReplies == nil {
		return nil, false
	}
	return &o.SuggestedReplies, true
}

// HasSuggestedReplies returns a boolean if a field has been set.
func (o *AiAgentEventPayload) HasSuggestedReplies() bool {
	return o != nil && o.SuggestedReplies != nil
}

// SetSuggestedReplies gets a reference to the given []AiAgentSuggestedReply and assigns it to the SuggestedReplies field.
func (o *AiAgentEventPayload) SetSuggestedReplies(v []AiAgentSuggestedReply) {
	o.SuggestedReplies = v
}

// GetRecoverActions returns the RecoverActions field value if set, zero value otherwise.
func (o *AiAgentEventPayload) GetRecoverActions() []AiAgentRecoverAction {
	if o == nil || o.RecoverActions == nil {
		var ret []AiAgentRecoverAction
		return ret
	}
	return o.RecoverActions
}

// GetRecoverActionsOk returns a tuple with the RecoverActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEventPayload) GetRecoverActionsOk() (*[]AiAgentRecoverAction, bool) {
	if o == nil || o.RecoverActions == nil {
		return nil, false
	}
	return &o.RecoverActions, true
}

// HasRecoverActions returns a boolean if a field has been set.
func (o *AiAgentEventPayload) HasRecoverActions() bool {
	return o != nil && o.RecoverActions != nil
}

// SetRecoverActions gets a reference to the given []AiAgentRecoverAction and assigns it to the RecoverActions field.
func (o *AiAgentEventPayload) SetRecoverActions(v []AiAgentRecoverAction) {
	o.RecoverActions = v
}

// GetCostClass returns the CostClass field value if set, zero value otherwise.
func (o *AiAgentEventPayload) GetCostClass() AiAgentCostClass {
	if o == nil || o.CostClass == nil {
		var ret AiAgentCostClass
		return ret
	}
	return *o.CostClass
}

// GetCostClassOk returns a tuple with the CostClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEventPayload) GetCostClassOk() (*AiAgentCostClass, bool) {
	if o == nil || o.CostClass == nil {
		return nil, false
	}
	return o.CostClass, true
}

// HasCostClass returns a boolean if a field has been set.
func (o *AiAgentEventPayload) HasCostClass() bool {
	return o != nil && o.CostClass != nil
}

// SetCostClass gets a reference to the given AiAgentCostClass and assigns it to the CostClass field.
func (o *AiAgentEventPayload) SetCostClass(v AiAgentCostClass) {
	o.CostClass = &v
}

// GetBudgetStatus returns the BudgetStatus field value if set, zero value otherwise.
func (o *AiAgentEventPayload) GetBudgetStatus() string {
	if o == nil || o.BudgetStatus == nil {
		var ret string
		return ret
	}
	return *o.BudgetStatus
}

// GetBudgetStatusOk returns a tuple with the BudgetStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEventPayload) GetBudgetStatusOk() (*string, bool) {
	if o == nil || o.BudgetStatus == nil {
		return nil, false
	}
	return o.BudgetStatus, true
}

// HasBudgetStatus returns a boolean if a field has been set.
func (o *AiAgentEventPayload) HasBudgetStatus() bool {
	return o != nil && o.BudgetStatus != nil
}

// SetBudgetStatus gets a reference to the given string and assigns it to the BudgetStatus field.
func (o *AiAgentEventPayload) SetBudgetStatus(v string) {
	o.BudgetStatus = &v
}

// GetRemainingBudgetClass returns the RemainingBudgetClass field value if set, zero value otherwise.
func (o *AiAgentEventPayload) GetRemainingBudgetClass() AiAgentCostClass {
	if o == nil || o.RemainingBudgetClass == nil {
		var ret AiAgentCostClass
		return ret
	}
	return *o.RemainingBudgetClass
}

// GetRemainingBudgetClassOk returns a tuple with the RemainingBudgetClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEventPayload) GetRemainingBudgetClassOk() (*AiAgentCostClass, bool) {
	if o == nil || o.RemainingBudgetClass == nil {
		return nil, false
	}
	return o.RemainingBudgetClass, true
}

// HasRemainingBudgetClass returns a boolean if a field has been set.
func (o *AiAgentEventPayload) HasRemainingBudgetClass() bool {
	return o != nil && o.RemainingBudgetClass != nil
}

// SetRemainingBudgetClass gets a reference to the given AiAgentCostClass and assigns it to the RemainingBudgetClass field.
func (o *AiAgentEventPayload) SetRemainingBudgetClass(v AiAgentCostClass) {
	o.RemainingBudgetClass = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *AiAgentEventPayload) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEventPayload) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *AiAgentEventPayload) HasReason() bool {
	return o != nil && o.Reason != nil
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *AiAgentEventPayload) SetReason(v string) {
	o.Reason = &v
}

// GetPlanTitle returns the PlanTitle field value if set, zero value otherwise.
func (o *AiAgentEventPayload) GetPlanTitle() string {
	if o == nil || o.PlanTitle == nil {
		var ret string
		return ret
	}
	return *o.PlanTitle
}

// GetPlanTitleOk returns a tuple with the PlanTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEventPayload) GetPlanTitleOk() (*string, bool) {
	if o == nil || o.PlanTitle == nil {
		return nil, false
	}
	return o.PlanTitle, true
}

// HasPlanTitle returns a boolean if a field has been set.
func (o *AiAgentEventPayload) HasPlanTitle() bool {
	return o != nil && o.PlanTitle != nil
}

// SetPlanTitle gets a reference to the given string and assigns it to the PlanTitle field.
func (o *AiAgentEventPayload) SetPlanTitle(v string) {
	o.PlanTitle = &v
}

// GetPlanSteps returns the PlanSteps field value if set, zero value otherwise.
func (o *AiAgentEventPayload) GetPlanSteps() []string {
	if o == nil || o.PlanSteps == nil {
		var ret []string
		return ret
	}
	return o.PlanSteps
}

// GetPlanStepsOk returns a tuple with the PlanSteps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEventPayload) GetPlanStepsOk() (*[]string, bool) {
	if o == nil || o.PlanSteps == nil {
		return nil, false
	}
	return &o.PlanSteps, true
}

// HasPlanSteps returns a boolean if a field has been set.
func (o *AiAgentEventPayload) HasPlanSteps() bool {
	return o != nil && o.PlanSteps != nil
}

// SetPlanSteps gets a reference to the given []string and assigns it to the PlanSteps field.
func (o *AiAgentEventPayload) SetPlanSteps(v []string) {
	o.PlanSteps = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentEventPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.MessageId != nil {
		toSerialize["messageId"] = o.MessageId
	}
	if o.PartId != nil {
		toSerialize["partId"] = o.PartId
	}
	if o.DeltaKind != nil {
		toSerialize["deltaKind"] = o.DeltaKind
	}
	if o.Sequence != nil {
		toSerialize["sequence"] = o.Sequence
	}
	if o.Delta != nil {
		toSerialize["delta"] = o.Delta
	}
	if o.IsFinal != nil {
		toSerialize["isFinal"] = o.IsFinal
	}
	if o.StageId != nil {
		toSerialize["stageId"] = o.StageId
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.DisplayLabel != nil {
		toSerialize["displayLabel"] = o.DisplayLabel
	}
	if o.DisplayDescription != nil {
		toSerialize["displayDescription"] = o.DisplayDescription
	}
	if o.UiVisibility != nil {
		toSerialize["uiVisibility"] = o.UiVisibility
	}
	if o.ToolCallId != nil {
		toSerialize["toolCallId"] = o.ToolCallId
	}
	if o.AgentToolId != nil {
		toSerialize["agentToolId"] = o.AgentToolId
	}
	if o.RequirementIds != nil {
		toSerialize["requirementIds"] = o.RequirementIds
	}
	if o.EvidenceIds != nil {
		toSerialize["evidenceIds"] = o.EvidenceIds
	}
	if o.MissingRequirements != nil {
		toSerialize["missingRequirements"] = o.MissingRequirements
	}
	if o.ToolResponseStatus != nil {
		toSerialize["toolResponseStatus"] = o.ToolResponseStatus
	}
	if o.MissingReason != nil {
		toSerialize["missingReason"] = o.MissingReason
	}
	if o.PartialReason != nil {
		toSerialize["partialReason"] = o.PartialReason
	}
	if o.StatusReason != nil {
		toSerialize["statusReason"] = o.StatusReason
	}
	if o.PrimaryMissingReason != nil {
		toSerialize["primaryMissingReason"] = o.PrimaryMissingReason
	}
	if o.EvidenceCompleteness != nil {
		toSerialize["evidenceCompleteness"] = o.EvidenceCompleteness
	}
	if o.Summary != nil {
		toSerialize["summary"] = o.Summary
	}
	if o.ReportId != nil {
		toSerialize["reportId"] = o.ReportId
	}
	if o.ReportUrl != nil {
		toSerialize["reportUrl"] = o.ReportUrl
	}
	if o.ActionPlanId != nil {
		toSerialize["actionPlanId"] = o.ActionPlanId
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.RiskLevel != nil {
		toSerialize["riskLevel"] = o.RiskLevel
	}
	if o.ImpactSummary != nil {
		toSerialize["impactSummary"] = o.ImpactSummary
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
	if o.RecoverActions != nil {
		toSerialize["recoverActions"] = o.RecoverActions
	}
	if o.CostClass != nil {
		toSerialize["costClass"] = o.CostClass
	}
	if o.BudgetStatus != nil {
		toSerialize["budgetStatus"] = o.BudgetStatus
	}
	if o.RemainingBudgetClass != nil {
		toSerialize["remainingBudgetClass"] = o.RemainingBudgetClass
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.PlanTitle != nil {
		toSerialize["planTitle"] = o.PlanTitle
	}
	if o.PlanSteps != nil {
		toSerialize["planSteps"] = o.PlanSteps
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentEventPayload) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MessageId             *string                      `json:"messageId,omitempty"`
		PartId                *string                      `json:"partId,omitempty"`
		DeltaKind             *AiAgentDeltaKind            `json:"deltaKind,omitempty"`
		Sequence              *int64                       `json:"sequence,omitempty"`
		Delta                 *string                      `json:"delta,omitempty"`
		IsFinal               *bool                        `json:"isFinal,omitempty"`
		StageId               *string                      `json:"stageId,omitempty"`
		State                 *AiAgentStatusDisplay        `json:"state,omitempty"`
		DisplayLabel          *string                      `json:"displayLabel,omitempty"`
		DisplayDescription    *string                      `json:"displayDescription,omitempty"`
		UiVisibility          *AiAgentUIVisibility         `json:"uiVisibility,omitempty"`
		ToolCallId            *string                      `json:"toolCallId,omitempty"`
		AgentToolId           *string                      `json:"agentToolId,omitempty"`
		RequirementIds        []string                     `json:"requirementIds,omitempty"`
		EvidenceIds           []string                     `json:"evidenceIds,omitempty"`
		MissingRequirements   []AiAgentMissingRequirement  `json:"missingRequirements,omitempty"`
		ToolResponseStatus    *AiAgentToolResponseStatus   `json:"toolResponseStatus,omitempty"`
		MissingReason         *AiAgentMissingReason        `json:"missingReason,omitempty"`
		PartialReason         *AiAgentPartialReason        `json:"partialReason,omitempty"`
		StatusReason          *string                      `json:"statusReason,omitempty"`
		PrimaryMissingReason  *string                      `json:"primaryMissingReason,omitempty"`
		EvidenceCompleteness  *AiAgentEvidenceCompleteness `json:"evidenceCompleteness,omitempty"`
		Summary               *string                      `json:"summary,omitempty"`
		ReportId              *string                      `json:"reportId,omitempty"`
		ReportUrl             *string                      `json:"reportUrl,omitempty"`
		ActionPlanId          *string                      `json:"actionPlanId,omitempty"`
		Title                 *string                      `json:"title,omitempty"`
		RiskLevel             *string                      `json:"riskLevel,omitempty"`
		ImpactSummary         *string                      `json:"impactSummary,omitempty"`
		Question              *string                      `json:"question,omitempty"`
		RequiredContextFields []string                     `json:"requiredContextFields,omitempty"`
		SuggestedReplies      []AiAgentSuggestedReply      `json:"suggestedReplies,omitempty"`
		RecoverActions        []AiAgentRecoverAction       `json:"recoverActions,omitempty"`
		CostClass             *AiAgentCostClass            `json:"costClass,omitempty"`
		BudgetStatus          *string                      `json:"budgetStatus,omitempty"`
		RemainingBudgetClass  *AiAgentCostClass            `json:"remainingBudgetClass,omitempty"`
		Reason                *string                      `json:"reason,omitempty"`
		PlanTitle             *string                      `json:"planTitle,omitempty"`
		PlanSteps             []string                     `json:"planSteps,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"messageId", "partId", "deltaKind", "sequence", "delta", "isFinal", "stageId", "state", "displayLabel", "displayDescription", "uiVisibility", "toolCallId", "agentToolId", "requirementIds", "evidenceIds", "missingRequirements", "toolResponseStatus", "missingReason", "partialReason", "statusReason", "primaryMissingReason", "evidenceCompleteness", "summary", "reportId", "reportUrl", "actionPlanId", "title", "riskLevel", "impactSummary", "question", "requiredContextFields", "suggestedReplies", "recoverActions", "costClass", "budgetStatus", "remainingBudgetClass", "reason", "planTitle", "planSteps"})
	} else {
		return err
	}

	hasInvalidField := false
	o.MessageId = all.MessageId
	o.PartId = all.PartId
	if all.DeltaKind != nil && !all.DeltaKind.IsValid() {
		hasInvalidField = true
	} else {
		o.DeltaKind = all.DeltaKind
	}
	o.Sequence = all.Sequence
	o.Delta = all.Delta
	o.IsFinal = all.IsFinal
	o.StageId = all.StageId
	if all.State != nil && all.State.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.State = all.State
	o.DisplayLabel = all.DisplayLabel
	o.DisplayDescription = all.DisplayDescription
	if all.UiVisibility != nil && !all.UiVisibility.IsValid() {
		hasInvalidField = true
	} else {
		o.UiVisibility = all.UiVisibility
	}
	o.ToolCallId = all.ToolCallId
	o.AgentToolId = all.AgentToolId
	o.RequirementIds = all.RequirementIds
	o.EvidenceIds = all.EvidenceIds
	o.MissingRequirements = all.MissingRequirements
	if all.ToolResponseStatus != nil && !all.ToolResponseStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.ToolResponseStatus = all.ToolResponseStatus
	}
	if all.MissingReason != nil && !all.MissingReason.IsValid() {
		hasInvalidField = true
	} else {
		o.MissingReason = all.MissingReason
	}
	if all.PartialReason != nil && !all.PartialReason.IsValid() {
		hasInvalidField = true
	} else {
		o.PartialReason = all.PartialReason
	}
	o.StatusReason = all.StatusReason
	o.PrimaryMissingReason = all.PrimaryMissingReason
	if all.EvidenceCompleteness != nil && !all.EvidenceCompleteness.IsValid() {
		hasInvalidField = true
	} else {
		o.EvidenceCompleteness = all.EvidenceCompleteness
	}
	o.Summary = all.Summary
	o.ReportId = all.ReportId
	o.ReportUrl = all.ReportUrl
	o.ActionPlanId = all.ActionPlanId
	o.Title = all.Title
	o.RiskLevel = all.RiskLevel
	o.ImpactSummary = all.ImpactSummary
	o.Question = all.Question
	o.RequiredContextFields = all.RequiredContextFields
	o.SuggestedReplies = all.SuggestedReplies
	o.RecoverActions = all.RecoverActions
	if all.CostClass != nil && !all.CostClass.IsValid() {
		hasInvalidField = true
	} else {
		o.CostClass = all.CostClass
	}
	o.BudgetStatus = all.BudgetStatus
	if all.RemainingBudgetClass != nil && !all.RemainingBudgetClass.IsValid() {
		hasInvalidField = true
	} else {
		o.RemainingBudgetClass = all.RemainingBudgetClass
	}
	o.Reason = all.Reason
	o.PlanTitle = all.PlanTitle
	o.PlanSteps = all.PlanSteps

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
