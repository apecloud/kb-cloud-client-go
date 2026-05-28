// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticEnvironmentDiagnosis Sprint 25 K8s environment diagnosis page contract. It groups the Sprint 24
// peripheral evidence explanations by environment layer without changing the
// peripheralEvidence compatibility field.
type DiagnosticEnvironmentDiagnosis struct {
	// Collection-result axis for the Sprint 24 K8s / Operator / Event explanation layer.
	// with_data means at least one external evidence explanation is available.
	// The remaining states explain why explanations are empty or incomplete without
	// inventing a database-internal conclusion.
	//
	State DiagnosticPeripheralEvidenceState `json:"state"`
	// Stable reason code for unavailable or degraded peripheral evidence.
	ReasonCode DiagnosticPeripheralEvidenceReasonCode `json:"reasonCode"`
	// First-screen summary for the K8s environment diagnosis page.
	Summary string `json:"summary"`
	// Whether this read-only diagnosis found no risk in the evidence it
	// actually collected. This means "no risk found in this run", not absolute
	// cluster health.
	//
	NoRisk bool `json:"noRisk"`
	// User-facing no-risk boundary summary; must not claim absolute health.
	NoRiskSummary string `json:"noRiskSummary"`
	// Explicit limits of the no-risk / environment-diagnosis conclusion.
	CannotProve []string                              `json:"cannotProve"`
	Layers      []DiagnosticEnvironmentDiagnosisLayer `json:"layers"`
	// Read-only safety boundary for environment diagnosis v0.2.
	ReadonlyBoundary *DiagnosticEnvironmentDiagnosisReadonlyBoundary `json:"readonlyBoundary,omitempty"`
	// Safe diagnosis-level actions for no-risk summary and rerun.
	NextActions []DiagnosticNextAction `json:"nextActions"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticEnvironmentDiagnosis instantiates a new DiagnosticEnvironmentDiagnosis object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticEnvironmentDiagnosis(state DiagnosticPeripheralEvidenceState, reasonCode DiagnosticPeripheralEvidenceReasonCode, summary string, noRisk bool, noRiskSummary string, cannotProve []string, layers []DiagnosticEnvironmentDiagnosisLayer, nextActions []DiagnosticNextAction) *DiagnosticEnvironmentDiagnosis {
	this := DiagnosticEnvironmentDiagnosis{}
	this.State = state
	this.ReasonCode = reasonCode
	this.Summary = summary
	this.NoRisk = noRisk
	this.NoRiskSummary = noRiskSummary
	this.CannotProve = cannotProve
	this.Layers = layers
	this.NextActions = nextActions
	return &this
}

// NewDiagnosticEnvironmentDiagnosisWithDefaults instantiates a new DiagnosticEnvironmentDiagnosis object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticEnvironmentDiagnosisWithDefaults() *DiagnosticEnvironmentDiagnosis {
	this := DiagnosticEnvironmentDiagnosis{}
	return &this
}

// GetState returns the State field value.
func (o *DiagnosticEnvironmentDiagnosis) GetState() DiagnosticPeripheralEvidenceState {
	if o == nil {
		var ret DiagnosticPeripheralEvidenceState
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *DiagnosticEnvironmentDiagnosis) GetStateOk() (*DiagnosticPeripheralEvidenceState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value.
func (o *DiagnosticEnvironmentDiagnosis) SetState(v DiagnosticPeripheralEvidenceState) {
	o.State = v
}

// GetReasonCode returns the ReasonCode field value.
func (o *DiagnosticEnvironmentDiagnosis) GetReasonCode() DiagnosticPeripheralEvidenceReasonCode {
	if o == nil {
		var ret DiagnosticPeripheralEvidenceReasonCode
		return ret
	}
	return o.ReasonCode
}

// GetReasonCodeOk returns a tuple with the ReasonCode field value
// and a boolean to check if the value has been set.
func (o *DiagnosticEnvironmentDiagnosis) GetReasonCodeOk() (*DiagnosticPeripheralEvidenceReasonCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReasonCode, true
}

// SetReasonCode sets field value.
func (o *DiagnosticEnvironmentDiagnosis) SetReasonCode(v DiagnosticPeripheralEvidenceReasonCode) {
	o.ReasonCode = v
}

// GetSummary returns the Summary field value.
func (o *DiagnosticEnvironmentDiagnosis) GetSummary() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *DiagnosticEnvironmentDiagnosis) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value.
func (o *DiagnosticEnvironmentDiagnosis) SetSummary(v string) {
	o.Summary = v
}

// GetNoRisk returns the NoRisk field value.
func (o *DiagnosticEnvironmentDiagnosis) GetNoRisk() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.NoRisk
}

// GetNoRiskOk returns a tuple with the NoRisk field value
// and a boolean to check if the value has been set.
func (o *DiagnosticEnvironmentDiagnosis) GetNoRiskOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NoRisk, true
}

// SetNoRisk sets field value.
func (o *DiagnosticEnvironmentDiagnosis) SetNoRisk(v bool) {
	o.NoRisk = v
}

// GetNoRiskSummary returns the NoRiskSummary field value.
func (o *DiagnosticEnvironmentDiagnosis) GetNoRiskSummary() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.NoRiskSummary
}

// GetNoRiskSummaryOk returns a tuple with the NoRiskSummary field value
// and a boolean to check if the value has been set.
func (o *DiagnosticEnvironmentDiagnosis) GetNoRiskSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NoRiskSummary, true
}

// SetNoRiskSummary sets field value.
func (o *DiagnosticEnvironmentDiagnosis) SetNoRiskSummary(v string) {
	o.NoRiskSummary = v
}

// GetCannotProve returns the CannotProve field value.
func (o *DiagnosticEnvironmentDiagnosis) GetCannotProve() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.CannotProve
}

// GetCannotProveOk returns a tuple with the CannotProve field value
// and a boolean to check if the value has been set.
func (o *DiagnosticEnvironmentDiagnosis) GetCannotProveOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CannotProve, true
}

// SetCannotProve sets field value.
func (o *DiagnosticEnvironmentDiagnosis) SetCannotProve(v []string) {
	o.CannotProve = v
}

// GetLayers returns the Layers field value.
func (o *DiagnosticEnvironmentDiagnosis) GetLayers() []DiagnosticEnvironmentDiagnosisLayer {
	if o == nil {
		var ret []DiagnosticEnvironmentDiagnosisLayer
		return ret
	}
	return o.Layers
}

// GetLayersOk returns a tuple with the Layers field value
// and a boolean to check if the value has been set.
func (o *DiagnosticEnvironmentDiagnosis) GetLayersOk() (*[]DiagnosticEnvironmentDiagnosisLayer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Layers, true
}

// SetLayers sets field value.
func (o *DiagnosticEnvironmentDiagnosis) SetLayers(v []DiagnosticEnvironmentDiagnosisLayer) {
	o.Layers = v
}

// GetReadonlyBoundary returns the ReadonlyBoundary field value if set, zero value otherwise.
func (o *DiagnosticEnvironmentDiagnosis) GetReadonlyBoundary() DiagnosticEnvironmentDiagnosisReadonlyBoundary {
	if o == nil || o.ReadonlyBoundary == nil {
		var ret DiagnosticEnvironmentDiagnosisReadonlyBoundary
		return ret
	}
	return *o.ReadonlyBoundary
}

// GetReadonlyBoundaryOk returns a tuple with the ReadonlyBoundary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticEnvironmentDiagnosis) GetReadonlyBoundaryOk() (*DiagnosticEnvironmentDiagnosisReadonlyBoundary, bool) {
	if o == nil || o.ReadonlyBoundary == nil {
		return nil, false
	}
	return o.ReadonlyBoundary, true
}

// HasReadonlyBoundary returns a boolean if a field has been set.
func (o *DiagnosticEnvironmentDiagnosis) HasReadonlyBoundary() bool {
	return o != nil && o.ReadonlyBoundary != nil
}

// SetReadonlyBoundary gets a reference to the given DiagnosticEnvironmentDiagnosisReadonlyBoundary and assigns it to the ReadonlyBoundary field.
func (o *DiagnosticEnvironmentDiagnosis) SetReadonlyBoundary(v DiagnosticEnvironmentDiagnosisReadonlyBoundary) {
	o.ReadonlyBoundary = &v
}

// GetNextActions returns the NextActions field value.
func (o *DiagnosticEnvironmentDiagnosis) GetNextActions() []DiagnosticNextAction {
	if o == nil {
		var ret []DiagnosticNextAction
		return ret
	}
	return o.NextActions
}

// GetNextActionsOk returns a tuple with the NextActions field value
// and a boolean to check if the value has been set.
func (o *DiagnosticEnvironmentDiagnosis) GetNextActionsOk() (*[]DiagnosticNextAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextActions, true
}

// SetNextActions sets field value.
func (o *DiagnosticEnvironmentDiagnosis) SetNextActions(v []DiagnosticNextAction) {
	o.NextActions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticEnvironmentDiagnosis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["state"] = o.State
	toSerialize["reasonCode"] = o.ReasonCode
	toSerialize["summary"] = o.Summary
	toSerialize["noRisk"] = o.NoRisk
	toSerialize["noRiskSummary"] = o.NoRiskSummary
	toSerialize["cannotProve"] = o.CannotProve
	toSerialize["layers"] = o.Layers
	if o.ReadonlyBoundary != nil {
		toSerialize["readonlyBoundary"] = o.ReadonlyBoundary
	}
	toSerialize["nextActions"] = o.NextActions

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticEnvironmentDiagnosis) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		State            *DiagnosticPeripheralEvidenceState              `json:"state"`
		ReasonCode       *DiagnosticPeripheralEvidenceReasonCode         `json:"reasonCode"`
		Summary          *string                                         `json:"summary"`
		NoRisk           *bool                                           `json:"noRisk"`
		NoRiskSummary    *string                                         `json:"noRiskSummary"`
		CannotProve      *[]string                                       `json:"cannotProve"`
		Layers           *[]DiagnosticEnvironmentDiagnosisLayer          `json:"layers"`
		ReadonlyBoundary *DiagnosticEnvironmentDiagnosisReadonlyBoundary `json:"readonlyBoundary,omitempty"`
		NextActions      *[]DiagnosticNextAction                         `json:"nextActions"`
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
	if all.NoRisk == nil {
		return fmt.Errorf("required field noRisk missing")
	}
	if all.NoRiskSummary == nil {
		return fmt.Errorf("required field noRiskSummary missing")
	}
	if all.CannotProve == nil {
		return fmt.Errorf("required field cannotProve missing")
	}
	if all.Layers == nil {
		return fmt.Errorf("required field layers missing")
	}
	if all.NextActions == nil {
		return fmt.Errorf("required field nextActions missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"state", "reasonCode", "summary", "noRisk", "noRiskSummary", "cannotProve", "layers", "readonlyBoundary", "nextActions"})
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
	o.NoRisk = *all.NoRisk
	o.NoRiskSummary = *all.NoRiskSummary
	o.CannotProve = *all.CannotProve
	o.Layers = *all.Layers
	if all.ReadonlyBoundary != nil && all.ReadonlyBoundary.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ReadonlyBoundary = all.ReadonlyBoundary
	o.NextActions = *all.NextActions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
