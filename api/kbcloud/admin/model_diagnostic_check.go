// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticCheck One executed diagnostic check result shown in the self-diagnosis report.
type DiagnosticCheck struct {
	// Stable check ID, such as event-analysis or db-key-metrics.
	Id string `json:"id"`
	// Human-readable check name.
	Name string `json:"name"`
	// Diagnostic report layer.
	Layer DiagnosticLayer `json:"layer"`
	// Result status for one executed diagnostic check.
	Status DiagnosticCheckStatus `json:"status"`
	// User-facing severity and priority bucket for one diagnostic check.
	Severity DiagnosticCheckSeverity `json:"severity"`
	// Deterministic priority, where 1 is the highest.
	Priority int64 `json:"priority"`
	// Whether this check is on the critical diagnostic path for health score calculation.
	Criticality DiagnosticCheckCriticality `json:"criticality"`
	// Whether this check describes database health or KBE control-plane risk.
	RiskDomain DiagnosticCheckRiskDomain `json:"riskDomain"`
	// Whether this check has proven relation to the current database object.
	RelationStatus DiagnosticCheckRelationStatus `json:"relationStatus"`
	// One-sentence check result.
	Summary string `json:"summary"`
	// Human-readable explanation of what this check means.
	Explanation string `json:"explanation"`
	// Impact of this check on the database cluster.
	Impact           string                      `json:"impact"`
	RelatedResources []DiagnosticRelatedResource `json:"relatedResources"`
	// Optional controller-log explanation cards. Only populated for the kb-controller-logs check.
	ControllerLogRelations []DiagnosticControllerLogRelation `json:"controllerLogRelations,omitempty"`
	// Human-readable evidence summary for the default page.
	EvidenceSummary string                 `json:"evidenceSummary"`
	NextActions     []DiagnosticNextAction `json:"nextActions"`
	RawEvidence     []DiagnosticEvidence   `json:"rawEvidence"`
	// Data-source quality for one diagnostic check.
	DataQuality DiagnosticDataQuality `json:"dataQuality"`
	Evidence    []DiagnosticEvidence  `json:"evidence"`
	// Deterministic judgment method used by this check.
	Algorithm string `json:"algorithm"`
	// Deterministic next action for the user or first responder.
	ActionHint string `json:"actionHint"`
	// Suggested first responder, such as platform, kbe-kb, or database team.
	OwnerHint string `json:"ownerHint"`
	// Failure, collection-failed, or not-applicable reason. Empty when no extra reason is needed.
	Reason string `json:"reason"`
	// Check execution time.
	CheckedAt time.Time `json:"checkedAt"`
	// Whether all attached evidence has been redacted.
	Redacted bool `json:"redacted"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticCheck instantiates a new DiagnosticCheck object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticCheck(id string, name string, layer DiagnosticLayer, status DiagnosticCheckStatus, severity DiagnosticCheckSeverity, priority int64, criticality DiagnosticCheckCriticality, riskDomain DiagnosticCheckRiskDomain, relationStatus DiagnosticCheckRelationStatus, summary string, explanation string, impact string, relatedResources []DiagnosticRelatedResource, evidenceSummary string, nextActions []DiagnosticNextAction, rawEvidence []DiagnosticEvidence, dataQuality DiagnosticDataQuality, evidence []DiagnosticEvidence, algorithm string, actionHint string, ownerHint string, reason string, checkedAt time.Time, redacted bool) *DiagnosticCheck {
	this := DiagnosticCheck{}
	this.Id = id
	this.Name = name
	this.Layer = layer
	this.Status = status
	this.Severity = severity
	this.Priority = priority
	this.Criticality = criticality
	this.RiskDomain = riskDomain
	this.RelationStatus = relationStatus
	this.Summary = summary
	this.Explanation = explanation
	this.Impact = impact
	this.RelatedResources = relatedResources
	this.EvidenceSummary = evidenceSummary
	this.NextActions = nextActions
	this.RawEvidence = rawEvidence
	this.DataQuality = dataQuality
	this.Evidence = evidence
	this.Algorithm = algorithm
	this.ActionHint = actionHint
	this.OwnerHint = ownerHint
	this.Reason = reason
	this.CheckedAt = checkedAt
	this.Redacted = redacted
	return &this
}

// NewDiagnosticCheckWithDefaults instantiates a new DiagnosticCheck object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticCheckWithDefaults() *DiagnosticCheck {
	this := DiagnosticCheck{}
	return &this
}

// GetId returns the Id field value.
func (o *DiagnosticCheck) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCheck) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *DiagnosticCheck) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value.
func (o *DiagnosticCheck) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCheck) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *DiagnosticCheck) SetName(v string) {
	o.Name = v
}

// GetLayer returns the Layer field value.
func (o *DiagnosticCheck) GetLayer() DiagnosticLayer {
	if o == nil {
		var ret DiagnosticLayer
		return ret
	}
	return o.Layer
}

// GetLayerOk returns a tuple with the Layer field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCheck) GetLayerOk() (*DiagnosticLayer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Layer, true
}

// SetLayer sets field value.
func (o *DiagnosticCheck) SetLayer(v DiagnosticLayer) {
	o.Layer = v
}

// GetStatus returns the Status field value.
func (o *DiagnosticCheck) GetStatus() DiagnosticCheckStatus {
	if o == nil {
		var ret DiagnosticCheckStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCheck) GetStatusOk() (*DiagnosticCheckStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *DiagnosticCheck) SetStatus(v DiagnosticCheckStatus) {
	o.Status = v
}

// GetSeverity returns the Severity field value.
func (o *DiagnosticCheck) GetSeverity() DiagnosticCheckSeverity {
	if o == nil {
		var ret DiagnosticCheckSeverity
		return ret
	}
	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCheck) GetSeverityOk() (*DiagnosticCheckSeverity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value.
func (o *DiagnosticCheck) SetSeverity(v DiagnosticCheckSeverity) {
	o.Severity = v
}

// GetPriority returns the Priority field value.
func (o *DiagnosticCheck) GetPriority() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCheck) GetPriorityOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value.
func (o *DiagnosticCheck) SetPriority(v int64) {
	o.Priority = v
}

// GetCriticality returns the Criticality field value.
func (o *DiagnosticCheck) GetCriticality() DiagnosticCheckCriticality {
	if o == nil {
		var ret DiagnosticCheckCriticality
		return ret
	}
	return o.Criticality
}

// GetCriticalityOk returns a tuple with the Criticality field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCheck) GetCriticalityOk() (*DiagnosticCheckCriticality, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Criticality, true
}

// SetCriticality sets field value.
func (o *DiagnosticCheck) SetCriticality(v DiagnosticCheckCriticality) {
	o.Criticality = v
}

// GetRiskDomain returns the RiskDomain field value.
func (o *DiagnosticCheck) GetRiskDomain() DiagnosticCheckRiskDomain {
	if o == nil {
		var ret DiagnosticCheckRiskDomain
		return ret
	}
	return o.RiskDomain
}

// GetRiskDomainOk returns a tuple with the RiskDomain field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCheck) GetRiskDomainOk() (*DiagnosticCheckRiskDomain, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RiskDomain, true
}

// SetRiskDomain sets field value.
func (o *DiagnosticCheck) SetRiskDomain(v DiagnosticCheckRiskDomain) {
	o.RiskDomain = v
}

// GetRelationStatus returns the RelationStatus field value.
func (o *DiagnosticCheck) GetRelationStatus() DiagnosticCheckRelationStatus {
	if o == nil {
		var ret DiagnosticCheckRelationStatus
		return ret
	}
	return o.RelationStatus
}

// GetRelationStatusOk returns a tuple with the RelationStatus field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCheck) GetRelationStatusOk() (*DiagnosticCheckRelationStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelationStatus, true
}

// SetRelationStatus sets field value.
func (o *DiagnosticCheck) SetRelationStatus(v DiagnosticCheckRelationStatus) {
	o.RelationStatus = v
}

// GetSummary returns the Summary field value.
func (o *DiagnosticCheck) GetSummary() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCheck) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value.
func (o *DiagnosticCheck) SetSummary(v string) {
	o.Summary = v
}

// GetExplanation returns the Explanation field value.
func (o *DiagnosticCheck) GetExplanation() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Explanation
}

// GetExplanationOk returns a tuple with the Explanation field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCheck) GetExplanationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Explanation, true
}

// SetExplanation sets field value.
func (o *DiagnosticCheck) SetExplanation(v string) {
	o.Explanation = v
}

// GetImpact returns the Impact field value.
func (o *DiagnosticCheck) GetImpact() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Impact
}

// GetImpactOk returns a tuple with the Impact field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCheck) GetImpactOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Impact, true
}

// SetImpact sets field value.
func (o *DiagnosticCheck) SetImpact(v string) {
	o.Impact = v
}

// GetRelatedResources returns the RelatedResources field value.
func (o *DiagnosticCheck) GetRelatedResources() []DiagnosticRelatedResource {
	if o == nil {
		var ret []DiagnosticRelatedResource
		return ret
	}
	return o.RelatedResources
}

// GetRelatedResourcesOk returns a tuple with the RelatedResources field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCheck) GetRelatedResourcesOk() (*[]DiagnosticRelatedResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelatedResources, true
}

// SetRelatedResources sets field value.
func (o *DiagnosticCheck) SetRelatedResources(v []DiagnosticRelatedResource) {
	o.RelatedResources = v
}

// GetControllerLogRelations returns the ControllerLogRelations field value if set, zero value otherwise.
func (o *DiagnosticCheck) GetControllerLogRelations() []DiagnosticControllerLogRelation {
	if o == nil || o.ControllerLogRelations == nil {
		var ret []DiagnosticControllerLogRelation
		return ret
	}
	return o.ControllerLogRelations
}

// GetControllerLogRelationsOk returns a tuple with the ControllerLogRelations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticCheck) GetControllerLogRelationsOk() (*[]DiagnosticControllerLogRelation, bool) {
	if o == nil || o.ControllerLogRelations == nil {
		return nil, false
	}
	return &o.ControllerLogRelations, true
}

// HasControllerLogRelations returns a boolean if a field has been set.
func (o *DiagnosticCheck) HasControllerLogRelations() bool {
	return o != nil && o.ControllerLogRelations != nil
}

// SetControllerLogRelations gets a reference to the given []DiagnosticControllerLogRelation and assigns it to the ControllerLogRelations field.
func (o *DiagnosticCheck) SetControllerLogRelations(v []DiagnosticControllerLogRelation) {
	o.ControllerLogRelations = v
}

// GetEvidenceSummary returns the EvidenceSummary field value.
func (o *DiagnosticCheck) GetEvidenceSummary() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EvidenceSummary
}

// GetEvidenceSummaryOk returns a tuple with the EvidenceSummary field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCheck) GetEvidenceSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EvidenceSummary, true
}

// SetEvidenceSummary sets field value.
func (o *DiagnosticCheck) SetEvidenceSummary(v string) {
	o.EvidenceSummary = v
}

// GetNextActions returns the NextActions field value.
func (o *DiagnosticCheck) GetNextActions() []DiagnosticNextAction {
	if o == nil {
		var ret []DiagnosticNextAction
		return ret
	}
	return o.NextActions
}

// GetNextActionsOk returns a tuple with the NextActions field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCheck) GetNextActionsOk() (*[]DiagnosticNextAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextActions, true
}

// SetNextActions sets field value.
func (o *DiagnosticCheck) SetNextActions(v []DiagnosticNextAction) {
	o.NextActions = v
}

// GetRawEvidence returns the RawEvidence field value.
func (o *DiagnosticCheck) GetRawEvidence() []DiagnosticEvidence {
	if o == nil {
		var ret []DiagnosticEvidence
		return ret
	}
	return o.RawEvidence
}

// GetRawEvidenceOk returns a tuple with the RawEvidence field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCheck) GetRawEvidenceOk() (*[]DiagnosticEvidence, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RawEvidence, true
}

// SetRawEvidence sets field value.
func (o *DiagnosticCheck) SetRawEvidence(v []DiagnosticEvidence) {
	o.RawEvidence = v
}

// GetDataQuality returns the DataQuality field value.
func (o *DiagnosticCheck) GetDataQuality() DiagnosticDataQuality {
	if o == nil {
		var ret DiagnosticDataQuality
		return ret
	}
	return o.DataQuality
}

// GetDataQualityOk returns a tuple with the DataQuality field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCheck) GetDataQualityOk() (*DiagnosticDataQuality, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataQuality, true
}

// SetDataQuality sets field value.
func (o *DiagnosticCheck) SetDataQuality(v DiagnosticDataQuality) {
	o.DataQuality = v
}

// GetEvidence returns the Evidence field value.
func (o *DiagnosticCheck) GetEvidence() []DiagnosticEvidence {
	if o == nil {
		var ret []DiagnosticEvidence
		return ret
	}
	return o.Evidence
}

// GetEvidenceOk returns a tuple with the Evidence field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCheck) GetEvidenceOk() (*[]DiagnosticEvidence, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Evidence, true
}

// SetEvidence sets field value.
func (o *DiagnosticCheck) SetEvidence(v []DiagnosticEvidence) {
	o.Evidence = v
}

// GetAlgorithm returns the Algorithm field value.
func (o *DiagnosticCheck) GetAlgorithm() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCheck) GetAlgorithmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Algorithm, true
}

// SetAlgorithm sets field value.
func (o *DiagnosticCheck) SetAlgorithm(v string) {
	o.Algorithm = v
}

// GetActionHint returns the ActionHint field value.
func (o *DiagnosticCheck) GetActionHint() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ActionHint
}

// GetActionHintOk returns a tuple with the ActionHint field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCheck) GetActionHintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionHint, true
}

// SetActionHint sets field value.
func (o *DiagnosticCheck) SetActionHint(v string) {
	o.ActionHint = v
}

// GetOwnerHint returns the OwnerHint field value.
func (o *DiagnosticCheck) GetOwnerHint() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OwnerHint
}

// GetOwnerHintOk returns a tuple with the OwnerHint field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCheck) GetOwnerHintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerHint, true
}

// SetOwnerHint sets field value.
func (o *DiagnosticCheck) SetOwnerHint(v string) {
	o.OwnerHint = v
}

// GetReason returns the Reason field value.
func (o *DiagnosticCheck) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCheck) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value.
func (o *DiagnosticCheck) SetReason(v string) {
	o.Reason = v
}

// GetCheckedAt returns the CheckedAt field value.
func (o *DiagnosticCheck) GetCheckedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CheckedAt
}

// GetCheckedAtOk returns a tuple with the CheckedAt field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCheck) GetCheckedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CheckedAt, true
}

// SetCheckedAt sets field value.
func (o *DiagnosticCheck) SetCheckedAt(v time.Time) {
	o.CheckedAt = v
}

// GetRedacted returns the Redacted field value.
func (o *DiagnosticCheck) GetRedacted() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Redacted
}

// GetRedactedOk returns a tuple with the Redacted field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCheck) GetRedactedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Redacted, true
}

// SetRedacted sets field value.
func (o *DiagnosticCheck) SetRedacted(v bool) {
	o.Redacted = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticCheck) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["layer"] = o.Layer
	toSerialize["status"] = o.Status
	toSerialize["severity"] = o.Severity
	toSerialize["priority"] = o.Priority
	toSerialize["criticality"] = o.Criticality
	toSerialize["riskDomain"] = o.RiskDomain
	toSerialize["relationStatus"] = o.RelationStatus
	toSerialize["summary"] = o.Summary
	toSerialize["explanation"] = o.Explanation
	toSerialize["impact"] = o.Impact
	toSerialize["relatedResources"] = o.RelatedResources
	if o.ControllerLogRelations != nil {
		toSerialize["controllerLogRelations"] = o.ControllerLogRelations
	}
	toSerialize["evidenceSummary"] = o.EvidenceSummary
	toSerialize["nextActions"] = o.NextActions
	toSerialize["rawEvidence"] = o.RawEvidence
	toSerialize["dataQuality"] = o.DataQuality
	toSerialize["evidence"] = o.Evidence
	toSerialize["algorithm"] = o.Algorithm
	toSerialize["actionHint"] = o.ActionHint
	toSerialize["ownerHint"] = o.OwnerHint
	toSerialize["reason"] = o.Reason
	if o.CheckedAt.Nanosecond() == 0 {
		toSerialize["checkedAt"] = o.CheckedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["checkedAt"] = o.CheckedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["redacted"] = o.Redacted

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticCheck) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id                     *string                           `json:"id"`
		Name                   *string                           `json:"name"`
		Layer                  *DiagnosticLayer                  `json:"layer"`
		Status                 *DiagnosticCheckStatus            `json:"status"`
		Severity               *DiagnosticCheckSeverity          `json:"severity"`
		Priority               *int64                            `json:"priority"`
		Criticality            *DiagnosticCheckCriticality       `json:"criticality"`
		RiskDomain             *DiagnosticCheckRiskDomain        `json:"riskDomain"`
		RelationStatus         *DiagnosticCheckRelationStatus    `json:"relationStatus"`
		Summary                *string                           `json:"summary"`
		Explanation            *string                           `json:"explanation"`
		Impact                 *string                           `json:"impact"`
		RelatedResources       *[]DiagnosticRelatedResource      `json:"relatedResources"`
		ControllerLogRelations []DiagnosticControllerLogRelation `json:"controllerLogRelations,omitempty"`
		EvidenceSummary        *string                           `json:"evidenceSummary"`
		NextActions            *[]DiagnosticNextAction           `json:"nextActions"`
		RawEvidence            *[]DiagnosticEvidence             `json:"rawEvidence"`
		DataQuality            *DiagnosticDataQuality            `json:"dataQuality"`
		Evidence               *[]DiagnosticEvidence             `json:"evidence"`
		Algorithm              *string                           `json:"algorithm"`
		ActionHint             *string                           `json:"actionHint"`
		OwnerHint              *string                           `json:"ownerHint"`
		Reason                 *string                           `json:"reason"`
		CheckedAt              *time.Time                        `json:"checkedAt"`
		Redacted               *bool                             `json:"redacted"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Layer == nil {
		return fmt.Errorf("required field layer missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.Severity == nil {
		return fmt.Errorf("required field severity missing")
	}
	if all.Priority == nil {
		return fmt.Errorf("required field priority missing")
	}
	if all.Criticality == nil {
		return fmt.Errorf("required field criticality missing")
	}
	if all.RiskDomain == nil {
		return fmt.Errorf("required field riskDomain missing")
	}
	if all.RelationStatus == nil {
		return fmt.Errorf("required field relationStatus missing")
	}
	if all.Summary == nil {
		return fmt.Errorf("required field summary missing")
	}
	if all.Explanation == nil {
		return fmt.Errorf("required field explanation missing")
	}
	if all.Impact == nil {
		return fmt.Errorf("required field impact missing")
	}
	if all.RelatedResources == nil {
		return fmt.Errorf("required field relatedResources missing")
	}
	if all.EvidenceSummary == nil {
		return fmt.Errorf("required field evidenceSummary missing")
	}
	if all.NextActions == nil {
		return fmt.Errorf("required field nextActions missing")
	}
	if all.RawEvidence == nil {
		return fmt.Errorf("required field rawEvidence missing")
	}
	if all.DataQuality == nil {
		return fmt.Errorf("required field dataQuality missing")
	}
	if all.Evidence == nil {
		return fmt.Errorf("required field evidence missing")
	}
	if all.Algorithm == nil {
		return fmt.Errorf("required field algorithm missing")
	}
	if all.ActionHint == nil {
		return fmt.Errorf("required field actionHint missing")
	}
	if all.OwnerHint == nil {
		return fmt.Errorf("required field ownerHint missing")
	}
	if all.Reason == nil {
		return fmt.Errorf("required field reason missing")
	}
	if all.CheckedAt == nil {
		return fmt.Errorf("required field checkedAt missing")
	}
	if all.Redacted == nil {
		return fmt.Errorf("required field redacted missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "name", "layer", "status", "severity", "priority", "criticality", "riskDomain", "relationStatus", "summary", "explanation", "impact", "relatedResources", "controllerLogRelations", "evidenceSummary", "nextActions", "rawEvidence", "dataQuality", "evidence", "algorithm", "actionHint", "ownerHint", "reason", "checkedAt", "redacted"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = *all.Id
	o.Name = *all.Name
	if !all.Layer.IsValid() {
		hasInvalidField = true
	} else {
		o.Layer = *all.Layer
	}
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	if !all.Severity.IsValid() {
		hasInvalidField = true
	} else {
		o.Severity = *all.Severity
	}
	o.Priority = *all.Priority
	if !all.Criticality.IsValid() {
		hasInvalidField = true
	} else {
		o.Criticality = *all.Criticality
	}
	if !all.RiskDomain.IsValid() {
		hasInvalidField = true
	} else {
		o.RiskDomain = *all.RiskDomain
	}
	if !all.RelationStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.RelationStatus = *all.RelationStatus
	}
	o.Summary = *all.Summary
	o.Explanation = *all.Explanation
	o.Impact = *all.Impact
	o.RelatedResources = *all.RelatedResources
	o.ControllerLogRelations = all.ControllerLogRelations
	o.EvidenceSummary = *all.EvidenceSummary
	o.NextActions = *all.NextActions
	o.RawEvidence = *all.RawEvidence
	if all.DataQuality.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DataQuality = *all.DataQuality
	o.Evidence = *all.Evidence
	o.Algorithm = *all.Algorithm
	o.ActionHint = *all.ActionHint
	o.OwnerHint = *all.OwnerHint
	o.Reason = *all.Reason
	o.CheckedAt = *all.CheckedAt
	o.Redacted = *all.Redacted

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
