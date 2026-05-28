// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticTopIssue One prioritized issue shown on the report first screen.
type DiagnosticTopIssue struct {
	// ID of the diagnostic check that produced this issue.
	CheckId string `json:"checkId"`
	// User-facing severity and priority bucket for one diagnostic check.
	Severity DiagnosticCheckSeverity `json:"severity"`
	// Deterministic priority, where 1 is the highest.
	Priority int64 `json:"priority"`
	// One-sentence issue summary.
	Summary string `json:"summary"`
	// Impact of this issue on the database cluster.
	Impact string `json:"impact"`
	// Deterministic next action for the user or first responder.
	ActionHint string `json:"actionHint"`
	// Suggested first responder.
	OwnerHint string `json:"ownerHint"`
	// Optional related resource reference.
	ResourceRef *string `json:"resourceRef,omitempty"`
	// Structured next actions for this top issue. Sprint 7 v0.1 hard gate.
	NextActions []DiagnosticNextAction `json:"nextActions"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticTopIssue instantiates a new DiagnosticTopIssue object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticTopIssue(checkId string, severity DiagnosticCheckSeverity, priority int64, summary string, impact string, actionHint string, ownerHint string, nextActions []DiagnosticNextAction) *DiagnosticTopIssue {
	this := DiagnosticTopIssue{}
	this.CheckId = checkId
	this.Severity = severity
	this.Priority = priority
	this.Summary = summary
	this.Impact = impact
	this.ActionHint = actionHint
	this.OwnerHint = ownerHint
	this.NextActions = nextActions
	return &this
}

// NewDiagnosticTopIssueWithDefaults instantiates a new DiagnosticTopIssue object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticTopIssueWithDefaults() *DiagnosticTopIssue {
	this := DiagnosticTopIssue{}
	return &this
}

// GetCheckId returns the CheckId field value.
func (o *DiagnosticTopIssue) GetCheckId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CheckId
}

// GetCheckIdOk returns a tuple with the CheckId field value
// and a boolean to check if the value has been set.
func (o *DiagnosticTopIssue) GetCheckIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CheckId, true
}

// SetCheckId sets field value.
func (o *DiagnosticTopIssue) SetCheckId(v string) {
	o.CheckId = v
}

// GetSeverity returns the Severity field value.
func (o *DiagnosticTopIssue) GetSeverity() DiagnosticCheckSeverity {
	if o == nil {
		var ret DiagnosticCheckSeverity
		return ret
	}
	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *DiagnosticTopIssue) GetSeverityOk() (*DiagnosticCheckSeverity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value.
func (o *DiagnosticTopIssue) SetSeverity(v DiagnosticCheckSeverity) {
	o.Severity = v
}

// GetPriority returns the Priority field value.
func (o *DiagnosticTopIssue) GetPriority() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *DiagnosticTopIssue) GetPriorityOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value.
func (o *DiagnosticTopIssue) SetPriority(v int64) {
	o.Priority = v
}

// GetSummary returns the Summary field value.
func (o *DiagnosticTopIssue) GetSummary() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *DiagnosticTopIssue) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value.
func (o *DiagnosticTopIssue) SetSummary(v string) {
	o.Summary = v
}

// GetImpact returns the Impact field value.
func (o *DiagnosticTopIssue) GetImpact() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Impact
}

// GetImpactOk returns a tuple with the Impact field value
// and a boolean to check if the value has been set.
func (o *DiagnosticTopIssue) GetImpactOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Impact, true
}

// SetImpact sets field value.
func (o *DiagnosticTopIssue) SetImpact(v string) {
	o.Impact = v
}

// GetActionHint returns the ActionHint field value.
func (o *DiagnosticTopIssue) GetActionHint() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ActionHint
}

// GetActionHintOk returns a tuple with the ActionHint field value
// and a boolean to check if the value has been set.
func (o *DiagnosticTopIssue) GetActionHintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionHint, true
}

// SetActionHint sets field value.
func (o *DiagnosticTopIssue) SetActionHint(v string) {
	o.ActionHint = v
}

// GetOwnerHint returns the OwnerHint field value.
func (o *DiagnosticTopIssue) GetOwnerHint() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OwnerHint
}

// GetOwnerHintOk returns a tuple with the OwnerHint field value
// and a boolean to check if the value has been set.
func (o *DiagnosticTopIssue) GetOwnerHintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerHint, true
}

// SetOwnerHint sets field value.
func (o *DiagnosticTopIssue) SetOwnerHint(v string) {
	o.OwnerHint = v
}

// GetResourceRef returns the ResourceRef field value if set, zero value otherwise.
func (o *DiagnosticTopIssue) GetResourceRef() string {
	if o == nil || o.ResourceRef == nil {
		var ret string
		return ret
	}
	return *o.ResourceRef
}

// GetResourceRefOk returns a tuple with the ResourceRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticTopIssue) GetResourceRefOk() (*string, bool) {
	if o == nil || o.ResourceRef == nil {
		return nil, false
	}
	return o.ResourceRef, true
}

// HasResourceRef returns a boolean if a field has been set.
func (o *DiagnosticTopIssue) HasResourceRef() bool {
	return o != nil && o.ResourceRef != nil
}

// SetResourceRef gets a reference to the given string and assigns it to the ResourceRef field.
func (o *DiagnosticTopIssue) SetResourceRef(v string) {
	o.ResourceRef = &v
}

// GetNextActions returns the NextActions field value.
func (o *DiagnosticTopIssue) GetNextActions() []DiagnosticNextAction {
	if o == nil {
		var ret []DiagnosticNextAction
		return ret
	}
	return o.NextActions
}

// GetNextActionsOk returns a tuple with the NextActions field value
// and a boolean to check if the value has been set.
func (o *DiagnosticTopIssue) GetNextActionsOk() (*[]DiagnosticNextAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextActions, true
}

// SetNextActions sets field value.
func (o *DiagnosticTopIssue) SetNextActions(v []DiagnosticNextAction) {
	o.NextActions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticTopIssue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["checkId"] = o.CheckId
	toSerialize["severity"] = o.Severity
	toSerialize["priority"] = o.Priority
	toSerialize["summary"] = o.Summary
	toSerialize["impact"] = o.Impact
	toSerialize["actionHint"] = o.ActionHint
	toSerialize["ownerHint"] = o.OwnerHint
	if o.ResourceRef != nil {
		toSerialize["resourceRef"] = o.ResourceRef
	}
	toSerialize["nextActions"] = o.NextActions

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticTopIssue) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CheckId     *string                  `json:"checkId"`
		Severity    *DiagnosticCheckSeverity `json:"severity"`
		Priority    *int64                   `json:"priority"`
		Summary     *string                  `json:"summary"`
		Impact      *string                  `json:"impact"`
		ActionHint  *string                  `json:"actionHint"`
		OwnerHint   *string                  `json:"ownerHint"`
		ResourceRef *string                  `json:"resourceRef,omitempty"`
		NextActions *[]DiagnosticNextAction  `json:"nextActions"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.CheckId == nil {
		return fmt.Errorf("required field checkId missing")
	}
	if all.Severity == nil {
		return fmt.Errorf("required field severity missing")
	}
	if all.Priority == nil {
		return fmt.Errorf("required field priority missing")
	}
	if all.Summary == nil {
		return fmt.Errorf("required field summary missing")
	}
	if all.Impact == nil {
		return fmt.Errorf("required field impact missing")
	}
	if all.ActionHint == nil {
		return fmt.Errorf("required field actionHint missing")
	}
	if all.OwnerHint == nil {
		return fmt.Errorf("required field ownerHint missing")
	}
	if all.NextActions == nil {
		return fmt.Errorf("required field nextActions missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"checkId", "severity", "priority", "summary", "impact", "actionHint", "ownerHint", "resourceRef", "nextActions"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CheckId = *all.CheckId
	if !all.Severity.IsValid() {
		hasInvalidField = true
	} else {
		o.Severity = *all.Severity
	}
	o.Priority = *all.Priority
	o.Summary = *all.Summary
	o.Impact = *all.Impact
	o.ActionHint = *all.ActionHint
	o.OwnerHint = *all.OwnerHint
	o.ResourceRef = all.ResourceRef
	o.NextActions = *all.NextActions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
