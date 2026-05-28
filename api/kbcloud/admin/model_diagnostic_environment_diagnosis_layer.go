// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticEnvironmentDiagnosisLayer One grouped layer in the K8s environment diagnosis page.
type DiagnosticEnvironmentDiagnosisLayer struct {
	// K8s environment diagnosis display layer.
	Layer DiagnosticEnvironmentDiagnosisLayerType `json:"layer"`
	// Human-readable layer title.
	Title string `json:"title"`
	// Collection-result axis for the Sprint 24 K8s / Operator / Event explanation layer.
	// with_data means at least one external evidence explanation is available.
	// The remaining states explain why explanations are empty or incomplete without
	// inventing a database-internal conclusion.
	//
	State DiagnosticPeripheralEvidenceState `json:"state"`
	// Stable reason code for unavailable or degraded peripheral evidence.
	ReasonCode *DiagnosticPeripheralEvidenceReasonCode `json:"reasonCode,omitempty"`
	// Summary of this environment layer.
	Summary string `json:"summary"`
	// Number of explanation cards in this layer.
	ExplanationCount int64 `json:"explanationCount"`
	// Peripheral evidence explanations grouped into this environment layer.
	Explanations []DiagnosticPeripheralEvidenceExplanation `json:"explanations"`
	// Number of direct object evidence records in this layer.
	DirectEvidenceCount *int64 `json:"directEvidenceCount,omitempty"`
	// Sprint 28 Network/Storage direct object evidence attached to this layer.
	DirectEvidence []DiagnosticEnvironmentDirectEvidence `json:"directEvidence,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticEnvironmentDiagnosisLayer instantiates a new DiagnosticEnvironmentDiagnosisLayer object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticEnvironmentDiagnosisLayer(layer DiagnosticEnvironmentDiagnosisLayerType, title string, state DiagnosticPeripheralEvidenceState, summary string, explanationCount int64, explanations []DiagnosticPeripheralEvidenceExplanation) *DiagnosticEnvironmentDiagnosisLayer {
	this := DiagnosticEnvironmentDiagnosisLayer{}
	this.Layer = layer
	this.Title = title
	this.State = state
	this.Summary = summary
	this.ExplanationCount = explanationCount
	this.Explanations = explanations
	return &this
}

// NewDiagnosticEnvironmentDiagnosisLayerWithDefaults instantiates a new DiagnosticEnvironmentDiagnosisLayer object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticEnvironmentDiagnosisLayerWithDefaults() *DiagnosticEnvironmentDiagnosisLayer {
	this := DiagnosticEnvironmentDiagnosisLayer{}
	return &this
}

// GetLayer returns the Layer field value.
func (o *DiagnosticEnvironmentDiagnosisLayer) GetLayer() DiagnosticEnvironmentDiagnosisLayerType {
	if o == nil {
		var ret DiagnosticEnvironmentDiagnosisLayerType
		return ret
	}
	return o.Layer
}

// GetLayerOk returns a tuple with the Layer field value
// and a boolean to check if the value has been set.
func (o *DiagnosticEnvironmentDiagnosisLayer) GetLayerOk() (*DiagnosticEnvironmentDiagnosisLayerType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Layer, true
}

// SetLayer sets field value.
func (o *DiagnosticEnvironmentDiagnosisLayer) SetLayer(v DiagnosticEnvironmentDiagnosisLayerType) {
	o.Layer = v
}

// GetTitle returns the Title field value.
func (o *DiagnosticEnvironmentDiagnosisLayer) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *DiagnosticEnvironmentDiagnosisLayer) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *DiagnosticEnvironmentDiagnosisLayer) SetTitle(v string) {
	o.Title = v
}

// GetState returns the State field value.
func (o *DiagnosticEnvironmentDiagnosisLayer) GetState() DiagnosticPeripheralEvidenceState {
	if o == nil {
		var ret DiagnosticPeripheralEvidenceState
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *DiagnosticEnvironmentDiagnosisLayer) GetStateOk() (*DiagnosticPeripheralEvidenceState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value.
func (o *DiagnosticEnvironmentDiagnosisLayer) SetState(v DiagnosticPeripheralEvidenceState) {
	o.State = v
}

// GetReasonCode returns the ReasonCode field value if set, zero value otherwise.
func (o *DiagnosticEnvironmentDiagnosisLayer) GetReasonCode() DiagnosticPeripheralEvidenceReasonCode {
	if o == nil || o.ReasonCode == nil {
		var ret DiagnosticPeripheralEvidenceReasonCode
		return ret
	}
	return *o.ReasonCode
}

// GetReasonCodeOk returns a tuple with the ReasonCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticEnvironmentDiagnosisLayer) GetReasonCodeOk() (*DiagnosticPeripheralEvidenceReasonCode, bool) {
	if o == nil || o.ReasonCode == nil {
		return nil, false
	}
	return o.ReasonCode, true
}

// HasReasonCode returns a boolean if a field has been set.
func (o *DiagnosticEnvironmentDiagnosisLayer) HasReasonCode() bool {
	return o != nil && o.ReasonCode != nil
}

// SetReasonCode gets a reference to the given DiagnosticPeripheralEvidenceReasonCode and assigns it to the ReasonCode field.
func (o *DiagnosticEnvironmentDiagnosisLayer) SetReasonCode(v DiagnosticPeripheralEvidenceReasonCode) {
	o.ReasonCode = &v
}

// GetSummary returns the Summary field value.
func (o *DiagnosticEnvironmentDiagnosisLayer) GetSummary() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *DiagnosticEnvironmentDiagnosisLayer) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value.
func (o *DiagnosticEnvironmentDiagnosisLayer) SetSummary(v string) {
	o.Summary = v
}

// GetExplanationCount returns the ExplanationCount field value.
func (o *DiagnosticEnvironmentDiagnosisLayer) GetExplanationCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ExplanationCount
}

// GetExplanationCountOk returns a tuple with the ExplanationCount field value
// and a boolean to check if the value has been set.
func (o *DiagnosticEnvironmentDiagnosisLayer) GetExplanationCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExplanationCount, true
}

// SetExplanationCount sets field value.
func (o *DiagnosticEnvironmentDiagnosisLayer) SetExplanationCount(v int64) {
	o.ExplanationCount = v
}

// GetExplanations returns the Explanations field value.
func (o *DiagnosticEnvironmentDiagnosisLayer) GetExplanations() []DiagnosticPeripheralEvidenceExplanation {
	if o == nil {
		var ret []DiagnosticPeripheralEvidenceExplanation
		return ret
	}
	return o.Explanations
}

// GetExplanationsOk returns a tuple with the Explanations field value
// and a boolean to check if the value has been set.
func (o *DiagnosticEnvironmentDiagnosisLayer) GetExplanationsOk() (*[]DiagnosticPeripheralEvidenceExplanation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Explanations, true
}

// SetExplanations sets field value.
func (o *DiagnosticEnvironmentDiagnosisLayer) SetExplanations(v []DiagnosticPeripheralEvidenceExplanation) {
	o.Explanations = v
}

// GetDirectEvidenceCount returns the DirectEvidenceCount field value if set, zero value otherwise.
func (o *DiagnosticEnvironmentDiagnosisLayer) GetDirectEvidenceCount() int64 {
	if o == nil || o.DirectEvidenceCount == nil {
		var ret int64
		return ret
	}
	return *o.DirectEvidenceCount
}

// GetDirectEvidenceCountOk returns a tuple with the DirectEvidenceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticEnvironmentDiagnosisLayer) GetDirectEvidenceCountOk() (*int64, bool) {
	if o == nil || o.DirectEvidenceCount == nil {
		return nil, false
	}
	return o.DirectEvidenceCount, true
}

// HasDirectEvidenceCount returns a boolean if a field has been set.
func (o *DiagnosticEnvironmentDiagnosisLayer) HasDirectEvidenceCount() bool {
	return o != nil && o.DirectEvidenceCount != nil
}

// SetDirectEvidenceCount gets a reference to the given int64 and assigns it to the DirectEvidenceCount field.
func (o *DiagnosticEnvironmentDiagnosisLayer) SetDirectEvidenceCount(v int64) {
	o.DirectEvidenceCount = &v
}

// GetDirectEvidence returns the DirectEvidence field value if set, zero value otherwise.
func (o *DiagnosticEnvironmentDiagnosisLayer) GetDirectEvidence() []DiagnosticEnvironmentDirectEvidence {
	if o == nil || o.DirectEvidence == nil {
		var ret []DiagnosticEnvironmentDirectEvidence
		return ret
	}
	return o.DirectEvidence
}

// GetDirectEvidenceOk returns a tuple with the DirectEvidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticEnvironmentDiagnosisLayer) GetDirectEvidenceOk() (*[]DiagnosticEnvironmentDirectEvidence, bool) {
	if o == nil || o.DirectEvidence == nil {
		return nil, false
	}
	return &o.DirectEvidence, true
}

// HasDirectEvidence returns a boolean if a field has been set.
func (o *DiagnosticEnvironmentDiagnosisLayer) HasDirectEvidence() bool {
	return o != nil && o.DirectEvidence != nil
}

// SetDirectEvidence gets a reference to the given []DiagnosticEnvironmentDirectEvidence and assigns it to the DirectEvidence field.
func (o *DiagnosticEnvironmentDiagnosisLayer) SetDirectEvidence(v []DiagnosticEnvironmentDirectEvidence) {
	o.DirectEvidence = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticEnvironmentDiagnosisLayer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["layer"] = o.Layer
	toSerialize["title"] = o.Title
	toSerialize["state"] = o.State
	if o.ReasonCode != nil {
		toSerialize["reasonCode"] = o.ReasonCode
	}
	toSerialize["summary"] = o.Summary
	toSerialize["explanationCount"] = o.ExplanationCount
	toSerialize["explanations"] = o.Explanations
	if o.DirectEvidenceCount != nil {
		toSerialize["directEvidenceCount"] = o.DirectEvidenceCount
	}
	if o.DirectEvidence != nil {
		toSerialize["directEvidence"] = o.DirectEvidence
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticEnvironmentDiagnosisLayer) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Layer               *DiagnosticEnvironmentDiagnosisLayerType   `json:"layer"`
		Title               *string                                    `json:"title"`
		State               *DiagnosticPeripheralEvidenceState         `json:"state"`
		ReasonCode          *DiagnosticPeripheralEvidenceReasonCode    `json:"reasonCode,omitempty"`
		Summary             *string                                    `json:"summary"`
		ExplanationCount    *int64                                     `json:"explanationCount"`
		Explanations        *[]DiagnosticPeripheralEvidenceExplanation `json:"explanations"`
		DirectEvidenceCount *int64                                     `json:"directEvidenceCount,omitempty"`
		DirectEvidence      []DiagnosticEnvironmentDirectEvidence      `json:"directEvidence,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Layer == nil {
		return fmt.Errorf("required field layer missing")
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	if all.State == nil {
		return fmt.Errorf("required field state missing")
	}
	if all.Summary == nil {
		return fmt.Errorf("required field summary missing")
	}
	if all.ExplanationCount == nil {
		return fmt.Errorf("required field explanationCount missing")
	}
	if all.Explanations == nil {
		return fmt.Errorf("required field explanations missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"layer", "title", "state", "reasonCode", "summary", "explanationCount", "explanations", "directEvidenceCount", "directEvidence"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Layer.IsValid() {
		hasInvalidField = true
	} else {
		o.Layer = *all.Layer
	}
	o.Title = *all.Title
	if !all.State.IsValid() {
		hasInvalidField = true
	} else {
		o.State = *all.State
	}
	if all.ReasonCode != nil && !all.ReasonCode.IsValid() {
		hasInvalidField = true
	} else {
		o.ReasonCode = all.ReasonCode
	}
	o.Summary = *all.Summary
	o.ExplanationCount = *all.ExplanationCount
	o.Explanations = *all.Explanations
	o.DirectEvidenceCount = all.DirectEvidenceCount
	o.DirectEvidence = all.DirectEvidence

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
