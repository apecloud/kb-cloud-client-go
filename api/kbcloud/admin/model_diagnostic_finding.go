// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticFinding One rule result in a diagnostic report.
type DiagnosticFinding struct {
	// Diagnostic report layer.
	Layer DiagnosticLayer `json:"layer"`
	// Stable rule ID, such as node-not-ready or kb-version-mismatch.
	RuleId string `json:"ruleId"`
	// Result status for one diagnostic rule.
	Status DiagnosticRuleStatus `json:"status"`
	// One-sentence result summary.
	Summary  string               `json:"summary"`
	Evidence []DiagnosticEvidence `json:"evidence"`
	// Deterministic next action. Do not include guesses without evidence.
	ActionHint *string `json:"actionHint,omitempty"`
	// Whether all attached evidence has been redacted.
	Redaction bool `json:"redaction"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticFinding instantiates a new DiagnosticFinding object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticFinding(layer DiagnosticLayer, ruleId string, status DiagnosticRuleStatus, summary string, evidence []DiagnosticEvidence, redaction bool) *DiagnosticFinding {
	this := DiagnosticFinding{}
	this.Layer = layer
	this.RuleId = ruleId
	this.Status = status
	this.Summary = summary
	this.Evidence = evidence
	this.Redaction = redaction
	return &this
}

// NewDiagnosticFindingWithDefaults instantiates a new DiagnosticFinding object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticFindingWithDefaults() *DiagnosticFinding {
	this := DiagnosticFinding{}
	return &this
}

// GetLayer returns the Layer field value.
func (o *DiagnosticFinding) GetLayer() DiagnosticLayer {
	if o == nil {
		var ret DiagnosticLayer
		return ret
	}
	return o.Layer
}

// GetLayerOk returns a tuple with the Layer field value
// and a boolean to check if the value has been set.
func (o *DiagnosticFinding) GetLayerOk() (*DiagnosticLayer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Layer, true
}

// SetLayer sets field value.
func (o *DiagnosticFinding) SetLayer(v DiagnosticLayer) {
	o.Layer = v
}

// GetRuleId returns the RuleId field value.
func (o *DiagnosticFinding) GetRuleId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value
// and a boolean to check if the value has been set.
func (o *DiagnosticFinding) GetRuleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleId, true
}

// SetRuleId sets field value.
func (o *DiagnosticFinding) SetRuleId(v string) {
	o.RuleId = v
}

// GetStatus returns the Status field value.
func (o *DiagnosticFinding) GetStatus() DiagnosticRuleStatus {
	if o == nil {
		var ret DiagnosticRuleStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DiagnosticFinding) GetStatusOk() (*DiagnosticRuleStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *DiagnosticFinding) SetStatus(v DiagnosticRuleStatus) {
	o.Status = v
}

// GetSummary returns the Summary field value.
func (o *DiagnosticFinding) GetSummary() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *DiagnosticFinding) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value.
func (o *DiagnosticFinding) SetSummary(v string) {
	o.Summary = v
}

// GetEvidence returns the Evidence field value.
func (o *DiagnosticFinding) GetEvidence() []DiagnosticEvidence {
	if o == nil {
		var ret []DiagnosticEvidence
		return ret
	}
	return o.Evidence
}

// GetEvidenceOk returns a tuple with the Evidence field value
// and a boolean to check if the value has been set.
func (o *DiagnosticFinding) GetEvidenceOk() (*[]DiagnosticEvidence, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Evidence, true
}

// SetEvidence sets field value.
func (o *DiagnosticFinding) SetEvidence(v []DiagnosticEvidence) {
	o.Evidence = v
}

// GetActionHint returns the ActionHint field value if set, zero value otherwise.
func (o *DiagnosticFinding) GetActionHint() string {
	if o == nil || o.ActionHint == nil {
		var ret string
		return ret
	}
	return *o.ActionHint
}

// GetActionHintOk returns a tuple with the ActionHint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticFinding) GetActionHintOk() (*string, bool) {
	if o == nil || o.ActionHint == nil {
		return nil, false
	}
	return o.ActionHint, true
}

// HasActionHint returns a boolean if a field has been set.
func (o *DiagnosticFinding) HasActionHint() bool {
	return o != nil && o.ActionHint != nil
}

// SetActionHint gets a reference to the given string and assigns it to the ActionHint field.
func (o *DiagnosticFinding) SetActionHint(v string) {
	o.ActionHint = &v
}

// GetRedaction returns the Redaction field value.
func (o *DiagnosticFinding) GetRedaction() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Redaction
}

// GetRedactionOk returns a tuple with the Redaction field value
// and a boolean to check if the value has been set.
func (o *DiagnosticFinding) GetRedactionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Redaction, true
}

// SetRedaction sets field value.
func (o *DiagnosticFinding) SetRedaction(v bool) {
	o.Redaction = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticFinding) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["layer"] = o.Layer
	toSerialize["ruleId"] = o.RuleId
	toSerialize["status"] = o.Status
	toSerialize["summary"] = o.Summary
	toSerialize["evidence"] = o.Evidence
	if o.ActionHint != nil {
		toSerialize["actionHint"] = o.ActionHint
	}
	toSerialize["redaction"] = o.Redaction

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticFinding) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Layer      *DiagnosticLayer      `json:"layer"`
		RuleId     *string               `json:"ruleId"`
		Status     *DiagnosticRuleStatus `json:"status"`
		Summary    *string               `json:"summary"`
		Evidence   *[]DiagnosticEvidence `json:"evidence"`
		ActionHint *string               `json:"actionHint,omitempty"`
		Redaction  *bool                 `json:"redaction"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Layer == nil {
		return fmt.Errorf("required field layer missing")
	}
	if all.RuleId == nil {
		return fmt.Errorf("required field ruleId missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.Summary == nil {
		return fmt.Errorf("required field summary missing")
	}
	if all.Evidence == nil {
		return fmt.Errorf("required field evidence missing")
	}
	if all.Redaction == nil {
		return fmt.Errorf("required field redaction missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"layer", "ruleId", "status", "summary", "evidence", "actionHint", "redaction"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Layer.IsValid() {
		hasInvalidField = true
	} else {
		o.Layer = *all.Layer
	}
	o.RuleId = *all.RuleId
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.Summary = *all.Summary
	o.Evidence = *all.Evidence
	o.ActionHint = all.ActionHint
	o.Redaction = *all.Redaction

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
