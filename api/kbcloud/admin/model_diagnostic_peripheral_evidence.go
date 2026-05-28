// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticPeripheralEvidence Sprint 24 K8s / Operator / Event external evidence explanation layer.
type DiagnosticPeripheralEvidence struct {
	// Collection-result axis for the Sprint 24 K8s / Operator / Event explanation layer.
	// with_data means at least one external evidence explanation is available.
	// The remaining states explain why explanations are empty or incomplete without
	// inventing a database-internal conclusion.
	//
	State DiagnosticPeripheralEvidenceState `json:"state"`
	// Stable reason code for unavailable or degraded peripheral evidence.
	ReasonCode DiagnosticPeripheralEvidenceReasonCode `json:"reasonCode"`
	// First-screen summary for peripheral evidence explanations.
	Summary      string                                    `json:"summary"`
	Explanations []DiagnosticPeripheralEvidenceExplanation `json:"explanations"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticPeripheralEvidence instantiates a new DiagnosticPeripheralEvidence object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticPeripheralEvidence(state DiagnosticPeripheralEvidenceState, reasonCode DiagnosticPeripheralEvidenceReasonCode, summary string, explanations []DiagnosticPeripheralEvidenceExplanation) *DiagnosticPeripheralEvidence {
	this := DiagnosticPeripheralEvidence{}
	this.State = state
	this.ReasonCode = reasonCode
	this.Summary = summary
	this.Explanations = explanations
	return &this
}

// NewDiagnosticPeripheralEvidenceWithDefaults instantiates a new DiagnosticPeripheralEvidence object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticPeripheralEvidenceWithDefaults() *DiagnosticPeripheralEvidence {
	this := DiagnosticPeripheralEvidence{}
	return &this
}

// GetState returns the State field value.
func (o *DiagnosticPeripheralEvidence) GetState() DiagnosticPeripheralEvidenceState {
	if o == nil {
		var ret DiagnosticPeripheralEvidenceState
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *DiagnosticPeripheralEvidence) GetStateOk() (*DiagnosticPeripheralEvidenceState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value.
func (o *DiagnosticPeripheralEvidence) SetState(v DiagnosticPeripheralEvidenceState) {
	o.State = v
}

// GetReasonCode returns the ReasonCode field value.
func (o *DiagnosticPeripheralEvidence) GetReasonCode() DiagnosticPeripheralEvidenceReasonCode {
	if o == nil {
		var ret DiagnosticPeripheralEvidenceReasonCode
		return ret
	}
	return o.ReasonCode
}

// GetReasonCodeOk returns a tuple with the ReasonCode field value
// and a boolean to check if the value has been set.
func (o *DiagnosticPeripheralEvidence) GetReasonCodeOk() (*DiagnosticPeripheralEvidenceReasonCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReasonCode, true
}

// SetReasonCode sets field value.
func (o *DiagnosticPeripheralEvidence) SetReasonCode(v DiagnosticPeripheralEvidenceReasonCode) {
	o.ReasonCode = v
}

// GetSummary returns the Summary field value.
func (o *DiagnosticPeripheralEvidence) GetSummary() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *DiagnosticPeripheralEvidence) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value.
func (o *DiagnosticPeripheralEvidence) SetSummary(v string) {
	o.Summary = v
}

// GetExplanations returns the Explanations field value.
func (o *DiagnosticPeripheralEvidence) GetExplanations() []DiagnosticPeripheralEvidenceExplanation {
	if o == nil {
		var ret []DiagnosticPeripheralEvidenceExplanation
		return ret
	}
	return o.Explanations
}

// GetExplanationsOk returns a tuple with the Explanations field value
// and a boolean to check if the value has been set.
func (o *DiagnosticPeripheralEvidence) GetExplanationsOk() (*[]DiagnosticPeripheralEvidenceExplanation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Explanations, true
}

// SetExplanations sets field value.
func (o *DiagnosticPeripheralEvidence) SetExplanations(v []DiagnosticPeripheralEvidenceExplanation) {
	o.Explanations = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticPeripheralEvidence) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["state"] = o.State
	toSerialize["reasonCode"] = o.ReasonCode
	toSerialize["summary"] = o.Summary
	toSerialize["explanations"] = o.Explanations

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticPeripheralEvidence) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		State        *DiagnosticPeripheralEvidenceState         `json:"state"`
		ReasonCode   *DiagnosticPeripheralEvidenceReasonCode    `json:"reasonCode"`
		Summary      *string                                    `json:"summary"`
		Explanations *[]DiagnosticPeripheralEvidenceExplanation `json:"explanations"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.State == nil {
		return fmt.Errorf("required field state missing")
	}
	if all.ReasonCode == nil {
		return fmt.Errorf("required field reasonCode missing")
	}
	if all.Summary == nil {
		return fmt.Errorf("required field summary missing")
	}
	if all.Explanations == nil {
		return fmt.Errorf("required field explanations missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"state", "reasonCode", "summary", "explanations"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.State.IsValid() {
		hasInvalidField = true
	} else {
		o.State = *all.State
	}
	if !all.ReasonCode.IsValid() {
		hasInvalidField = true
	} else {
		o.ReasonCode = *all.ReasonCode
	}
	o.Summary = *all.Summary
	o.Explanations = *all.Explanations

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
