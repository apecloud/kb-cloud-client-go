// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticCPUAnalysis Sprint 27 PostgreSQL readonly CPU analysis contract. CPU, SQL, session, and
// Kubernetes evidence are intentionally separated so Console does not turn a
// correlation into a root-cause proof.
type DiagnosticCPUAnalysis struct {
	// Business-result axis for PostgreSQL readonly CPU analysis. with_data means
	// at least one CPU-related evidence partition was collected. no_data means
	// the collector ran but did not find CPU risk evidence in the current window.
	// collection_failed and source_missing keep unavailable data sources explicit.
	//
	State DiagnosticCPUAnalysisState `json:"state"`
	// Stable reason code for unavailable or degraded CPU analysis.
	ReasonCode DiagnosticCPUAnalysisReasonCode `json:"reasonCode"`
	// Data quality status for one database performance metric.
	DataQuality DiagnosticMetricDataQuality `json:"dataQuality"`
	// First-screen summary for the CPU analysis card.
	Summary string `json:"summary"`
	// Time when the CPU analysis result was assembled.
	CollectedAt time.Time `json:"collectedAt"`
	// Time window represented by the CPU analysis evidence.
	TimeWindow DiagnosticCPUAnalysisTimeWindow `json:"timeWindow"`
	// Best-effort evidence bucket for the first screen. This is not a root-cause
	// proof and must be read with cannotProve.
	//
	PrimaryCause DiagnosticCPUAnalysisPrimaryCause `json:"primaryCause"`
	// Confidence derived from the available readonly evidence.
	Confidence DiagnosticCPUAnalysisConfidence `json:"confidence"`
	// CPU risk level derived from evidence thresholds, not from an automatic fix decision.
	RiskLevel DiagnosticCPUAnalysisRiskLevel `json:"riskLevel"`
	// Pod/container CPU evidence from Kubernetes metrics and resource requests/limits.
	PodCpuEvidence DiagnosticCPUAnalysisPodCPUEvidence `json:"podCPUEvidence"`
	// PostgreSQL session/activity evidence from readonly system views.
	DbActivityEvidence DiagnosticCPUAnalysisDBActivityEvidence `json:"dbActivityEvidence"`
	// SQL pressure evidence using digest/summary only; no raw SQL text.
	SqlPressureEvidence DiagnosticCPUAnalysisSQLPressureEvidence `json:"sqlPressureEvidence"`
	// Kubernetes resource/event evidence in a separate section.
	K8sEvidence []DiagnosticCPUAnalysisK8sEvidence `json:"k8sEvidence"`
	// Explicit safety boundary for the CPU analysis collector.
	ReadonlyBoundary DiagnosticCPUAnalysisReadonlyBoundary `json:"readonlyBoundary"`
	// Structured collection error detail when cpuAnalysisState is source_present_collection_failed.
	CollectionError *DiagnosticCPUAnalysisCollectionError `json:"collectionError,omitempty"`
	// Safe actions only: open detail, copy redacted context, jump to evidence, or rerun diagnosis.
	NextActions []DiagnosticNextAction `json:"nextActions"`
	// Explicit limits of this evidence.
	CannotProve []string `json:"cannotProve"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticCPUAnalysis instantiates a new DiagnosticCPUAnalysis object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticCPUAnalysis(state DiagnosticCPUAnalysisState, reasonCode DiagnosticCPUAnalysisReasonCode, dataQuality DiagnosticMetricDataQuality, summary string, collectedAt time.Time, timeWindow DiagnosticCPUAnalysisTimeWindow, primaryCause DiagnosticCPUAnalysisPrimaryCause, confidence DiagnosticCPUAnalysisConfidence, riskLevel DiagnosticCPUAnalysisRiskLevel, podCpuEvidence DiagnosticCPUAnalysisPodCPUEvidence, dbActivityEvidence DiagnosticCPUAnalysisDBActivityEvidence, sqlPressureEvidence DiagnosticCPUAnalysisSQLPressureEvidence, k8sEvidence []DiagnosticCPUAnalysisK8sEvidence, readonlyBoundary DiagnosticCPUAnalysisReadonlyBoundary, nextActions []DiagnosticNextAction, cannotProve []string) *DiagnosticCPUAnalysis {
	this := DiagnosticCPUAnalysis{}
	this.State = state
	this.ReasonCode = reasonCode
	this.DataQuality = dataQuality
	this.Summary = summary
	this.CollectedAt = collectedAt
	this.TimeWindow = timeWindow
	this.PrimaryCause = primaryCause
	this.Confidence = confidence
	this.RiskLevel = riskLevel
	this.PodCpuEvidence = podCpuEvidence
	this.DbActivityEvidence = dbActivityEvidence
	this.SqlPressureEvidence = sqlPressureEvidence
	this.K8sEvidence = k8sEvidence
	this.ReadonlyBoundary = readonlyBoundary
	this.NextActions = nextActions
	this.CannotProve = cannotProve
	return &this
}

// NewDiagnosticCPUAnalysisWithDefaults instantiates a new DiagnosticCPUAnalysis object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticCPUAnalysisWithDefaults() *DiagnosticCPUAnalysis {
	this := DiagnosticCPUAnalysis{}
	return &this
}

// GetState returns the State field value.
func (o *DiagnosticCPUAnalysis) GetState() DiagnosticCPUAnalysisState {
	if o == nil {
		var ret DiagnosticCPUAnalysisState
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysis) GetStateOk() (*DiagnosticCPUAnalysisState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value.
func (o *DiagnosticCPUAnalysis) SetState(v DiagnosticCPUAnalysisState) {
	o.State = v
}

// GetReasonCode returns the ReasonCode field value.
func (o *DiagnosticCPUAnalysis) GetReasonCode() DiagnosticCPUAnalysisReasonCode {
	if o == nil {
		var ret DiagnosticCPUAnalysisReasonCode
		return ret
	}
	return o.ReasonCode
}

// GetReasonCodeOk returns a tuple with the ReasonCode field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysis) GetReasonCodeOk() (*DiagnosticCPUAnalysisReasonCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReasonCode, true
}

// SetReasonCode sets field value.
func (o *DiagnosticCPUAnalysis) SetReasonCode(v DiagnosticCPUAnalysisReasonCode) {
	o.ReasonCode = v
}

// GetDataQuality returns the DataQuality field value.
func (o *DiagnosticCPUAnalysis) GetDataQuality() DiagnosticMetricDataQuality {
	if o == nil {
		var ret DiagnosticMetricDataQuality
		return ret
	}
	return o.DataQuality
}

// GetDataQualityOk returns a tuple with the DataQuality field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysis) GetDataQualityOk() (*DiagnosticMetricDataQuality, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataQuality, true
}

// SetDataQuality sets field value.
func (o *DiagnosticCPUAnalysis) SetDataQuality(v DiagnosticMetricDataQuality) {
	o.DataQuality = v
}

// GetSummary returns the Summary field value.
func (o *DiagnosticCPUAnalysis) GetSummary() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysis) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value.
func (o *DiagnosticCPUAnalysis) SetSummary(v string) {
	o.Summary = v
}

// GetCollectedAt returns the CollectedAt field value.
func (o *DiagnosticCPUAnalysis) GetCollectedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CollectedAt
}

// GetCollectedAtOk returns a tuple with the CollectedAt field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysis) GetCollectedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectedAt, true
}

// SetCollectedAt sets field value.
func (o *DiagnosticCPUAnalysis) SetCollectedAt(v time.Time) {
	o.CollectedAt = v
}

// GetTimeWindow returns the TimeWindow field value.
func (o *DiagnosticCPUAnalysis) GetTimeWindow() DiagnosticCPUAnalysisTimeWindow {
	if o == nil {
		var ret DiagnosticCPUAnalysisTimeWindow
		return ret
	}
	return o.TimeWindow
}

// GetTimeWindowOk returns a tuple with the TimeWindow field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysis) GetTimeWindowOk() (*DiagnosticCPUAnalysisTimeWindow, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeWindow, true
}

// SetTimeWindow sets field value.
func (o *DiagnosticCPUAnalysis) SetTimeWindow(v DiagnosticCPUAnalysisTimeWindow) {
	o.TimeWindow = v
}

// GetPrimaryCause returns the PrimaryCause field value.
func (o *DiagnosticCPUAnalysis) GetPrimaryCause() DiagnosticCPUAnalysisPrimaryCause {
	if o == nil {
		var ret DiagnosticCPUAnalysisPrimaryCause
		return ret
	}
	return o.PrimaryCause
}

// GetPrimaryCauseOk returns a tuple with the PrimaryCause field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysis) GetPrimaryCauseOk() (*DiagnosticCPUAnalysisPrimaryCause, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrimaryCause, true
}

// SetPrimaryCause sets field value.
func (o *DiagnosticCPUAnalysis) SetPrimaryCause(v DiagnosticCPUAnalysisPrimaryCause) {
	o.PrimaryCause = v
}

// GetConfidence returns the Confidence field value.
func (o *DiagnosticCPUAnalysis) GetConfidence() DiagnosticCPUAnalysisConfidence {
	if o == nil {
		var ret DiagnosticCPUAnalysisConfidence
		return ret
	}
	return o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysis) GetConfidenceOk() (*DiagnosticCPUAnalysisConfidence, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Confidence, true
}

// SetConfidence sets field value.
func (o *DiagnosticCPUAnalysis) SetConfidence(v DiagnosticCPUAnalysisConfidence) {
	o.Confidence = v
}

// GetRiskLevel returns the RiskLevel field value.
func (o *DiagnosticCPUAnalysis) GetRiskLevel() DiagnosticCPUAnalysisRiskLevel {
	if o == nil {
		var ret DiagnosticCPUAnalysisRiskLevel
		return ret
	}
	return o.RiskLevel
}

// GetRiskLevelOk returns a tuple with the RiskLevel field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysis) GetRiskLevelOk() (*DiagnosticCPUAnalysisRiskLevel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RiskLevel, true
}

// SetRiskLevel sets field value.
func (o *DiagnosticCPUAnalysis) SetRiskLevel(v DiagnosticCPUAnalysisRiskLevel) {
	o.RiskLevel = v
}

// GetPodCpuEvidence returns the PodCpuEvidence field value.
func (o *DiagnosticCPUAnalysis) GetPodCpuEvidence() DiagnosticCPUAnalysisPodCPUEvidence {
	if o == nil {
		var ret DiagnosticCPUAnalysisPodCPUEvidence
		return ret
	}
	return o.PodCpuEvidence
}

// GetPodCpuEvidenceOk returns a tuple with the PodCpuEvidence field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysis) GetPodCpuEvidenceOk() (*DiagnosticCPUAnalysisPodCPUEvidence, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PodCpuEvidence, true
}

// SetPodCpuEvidence sets field value.
func (o *DiagnosticCPUAnalysis) SetPodCpuEvidence(v DiagnosticCPUAnalysisPodCPUEvidence) {
	o.PodCpuEvidence = v
}

// GetDbActivityEvidence returns the DbActivityEvidence field value.
func (o *DiagnosticCPUAnalysis) GetDbActivityEvidence() DiagnosticCPUAnalysisDBActivityEvidence {
	if o == nil {
		var ret DiagnosticCPUAnalysisDBActivityEvidence
		return ret
	}
	return o.DbActivityEvidence
}

// GetDbActivityEvidenceOk returns a tuple with the DbActivityEvidence field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysis) GetDbActivityEvidenceOk() (*DiagnosticCPUAnalysisDBActivityEvidence, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DbActivityEvidence, true
}

// SetDbActivityEvidence sets field value.
func (o *DiagnosticCPUAnalysis) SetDbActivityEvidence(v DiagnosticCPUAnalysisDBActivityEvidence) {
	o.DbActivityEvidence = v
}

// GetSqlPressureEvidence returns the SqlPressureEvidence field value.
func (o *DiagnosticCPUAnalysis) GetSqlPressureEvidence() DiagnosticCPUAnalysisSQLPressureEvidence {
	if o == nil {
		var ret DiagnosticCPUAnalysisSQLPressureEvidence
		return ret
	}
	return o.SqlPressureEvidence
}

// GetSqlPressureEvidenceOk returns a tuple with the SqlPressureEvidence field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysis) GetSqlPressureEvidenceOk() (*DiagnosticCPUAnalysisSQLPressureEvidence, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SqlPressureEvidence, true
}

// SetSqlPressureEvidence sets field value.
func (o *DiagnosticCPUAnalysis) SetSqlPressureEvidence(v DiagnosticCPUAnalysisSQLPressureEvidence) {
	o.SqlPressureEvidence = v
}

// GetK8sEvidence returns the K8sEvidence field value.
func (o *DiagnosticCPUAnalysis) GetK8sEvidence() []DiagnosticCPUAnalysisK8sEvidence {
	if o == nil {
		var ret []DiagnosticCPUAnalysisK8sEvidence
		return ret
	}
	return o.K8sEvidence
}

// GetK8sEvidenceOk returns a tuple with the K8sEvidence field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysis) GetK8sEvidenceOk() (*[]DiagnosticCPUAnalysisK8sEvidence, bool) {
	if o == nil {
		return nil, false
	}
	return &o.K8sEvidence, true
}

// SetK8sEvidence sets field value.
func (o *DiagnosticCPUAnalysis) SetK8sEvidence(v []DiagnosticCPUAnalysisK8sEvidence) {
	o.K8sEvidence = v
}

// GetReadonlyBoundary returns the ReadonlyBoundary field value.
func (o *DiagnosticCPUAnalysis) GetReadonlyBoundary() DiagnosticCPUAnalysisReadonlyBoundary {
	if o == nil {
		var ret DiagnosticCPUAnalysisReadonlyBoundary
		return ret
	}
	return o.ReadonlyBoundary
}

// GetReadonlyBoundaryOk returns a tuple with the ReadonlyBoundary field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysis) GetReadonlyBoundaryOk() (*DiagnosticCPUAnalysisReadonlyBoundary, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReadonlyBoundary, true
}

// SetReadonlyBoundary sets field value.
func (o *DiagnosticCPUAnalysis) SetReadonlyBoundary(v DiagnosticCPUAnalysisReadonlyBoundary) {
	o.ReadonlyBoundary = v
}

// GetCollectionError returns the CollectionError field value if set, zero value otherwise.
func (o *DiagnosticCPUAnalysis) GetCollectionError() DiagnosticCPUAnalysisCollectionError {
	if o == nil || o.CollectionError == nil {
		var ret DiagnosticCPUAnalysisCollectionError
		return ret
	}
	return *o.CollectionError
}

// GetCollectionErrorOk returns a tuple with the CollectionError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysis) GetCollectionErrorOk() (*DiagnosticCPUAnalysisCollectionError, bool) {
	if o == nil || o.CollectionError == nil {
		return nil, false
	}
	return o.CollectionError, true
}

// HasCollectionError returns a boolean if a field has been set.
func (o *DiagnosticCPUAnalysis) HasCollectionError() bool {
	return o != nil && o.CollectionError != nil
}

// SetCollectionError gets a reference to the given DiagnosticCPUAnalysisCollectionError and assigns it to the CollectionError field.
func (o *DiagnosticCPUAnalysis) SetCollectionError(v DiagnosticCPUAnalysisCollectionError) {
	o.CollectionError = &v
}

// GetNextActions returns the NextActions field value.
func (o *DiagnosticCPUAnalysis) GetNextActions() []DiagnosticNextAction {
	if o == nil {
		var ret []DiagnosticNextAction
		return ret
	}
	return o.NextActions
}

// GetNextActionsOk returns a tuple with the NextActions field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysis) GetNextActionsOk() (*[]DiagnosticNextAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextActions, true
}

// SetNextActions sets field value.
func (o *DiagnosticCPUAnalysis) SetNextActions(v []DiagnosticNextAction) {
	o.NextActions = v
}

// GetCannotProve returns the CannotProve field value.
func (o *DiagnosticCPUAnalysis) GetCannotProve() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.CannotProve
}

// GetCannotProveOk returns a tuple with the CannotProve field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysis) GetCannotProveOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CannotProve, true
}

// SetCannotProve sets field value.
func (o *DiagnosticCPUAnalysis) SetCannotProve(v []string) {
	o.CannotProve = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticCPUAnalysis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["state"] = o.State
	toSerialize["reasonCode"] = o.ReasonCode
	toSerialize["dataQuality"] = o.DataQuality
	toSerialize["summary"] = o.Summary
	if o.CollectedAt.Nanosecond() == 0 {
		toSerialize["collectedAt"] = o.CollectedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["collectedAt"] = o.CollectedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["timeWindow"] = o.TimeWindow
	toSerialize["primaryCause"] = o.PrimaryCause
	toSerialize["confidence"] = o.Confidence
	toSerialize["riskLevel"] = o.RiskLevel
	toSerialize["podCPUEvidence"] = o.PodCpuEvidence
	toSerialize["dbActivityEvidence"] = o.DbActivityEvidence
	toSerialize["sqlPressureEvidence"] = o.SqlPressureEvidence
	toSerialize["k8sEvidence"] = o.K8sEvidence
	toSerialize["readonlyBoundary"] = o.ReadonlyBoundary
	if o.CollectionError != nil {
		toSerialize["collectionError"] = o.CollectionError
	}
	toSerialize["nextActions"] = o.NextActions
	toSerialize["cannotProve"] = o.CannotProve

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticCPUAnalysis) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		State               *DiagnosticCPUAnalysisState               `json:"state"`
		ReasonCode          *DiagnosticCPUAnalysisReasonCode          `json:"reasonCode"`
		DataQuality         *DiagnosticMetricDataQuality              `json:"dataQuality"`
		Summary             *string                                   `json:"summary"`
		CollectedAt         *time.Time                                `json:"collectedAt"`
		TimeWindow          *DiagnosticCPUAnalysisTimeWindow          `json:"timeWindow"`
		PrimaryCause        *DiagnosticCPUAnalysisPrimaryCause        `json:"primaryCause"`
		Confidence          *DiagnosticCPUAnalysisConfidence          `json:"confidence"`
		RiskLevel           *DiagnosticCPUAnalysisRiskLevel           `json:"riskLevel"`
		PodCpuEvidence      *DiagnosticCPUAnalysisPodCPUEvidence      `json:"podCPUEvidence"`
		DbActivityEvidence  *DiagnosticCPUAnalysisDBActivityEvidence  `json:"dbActivityEvidence"`
		SqlPressureEvidence *DiagnosticCPUAnalysisSQLPressureEvidence `json:"sqlPressureEvidence"`
		K8sEvidence         *[]DiagnosticCPUAnalysisK8sEvidence       `json:"k8sEvidence"`
		ReadonlyBoundary    *DiagnosticCPUAnalysisReadonlyBoundary    `json:"readonlyBoundary"`
		CollectionError     *DiagnosticCPUAnalysisCollectionError     `json:"collectionError,omitempty"`
		NextActions         *[]DiagnosticNextAction                   `json:"nextActions"`
		CannotProve         *[]string                                 `json:"cannotProve"`
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
	if all.DataQuality == nil {
		return fmt.Errorf("required field dataQuality missing")
	}
	if all.Summary == nil {
		return fmt.Errorf("required field summary missing")
	}
	if all.CollectedAt == nil {
		return fmt.Errorf("required field collectedAt missing")
	}
	if all.TimeWindow == nil {
		return fmt.Errorf("required field timeWindow missing")
	}
	if all.PrimaryCause == nil {
		return fmt.Errorf("required field primaryCause missing")
	}
	if all.Confidence == nil {
		return fmt.Errorf("required field confidence missing")
	}
	if all.RiskLevel == nil {
		return fmt.Errorf("required field riskLevel missing")
	}
	if all.PodCpuEvidence == nil {
		return fmt.Errorf("required field podCPUEvidence missing")
	}
	if all.DbActivityEvidence == nil {
		return fmt.Errorf("required field dbActivityEvidence missing")
	}
	if all.SqlPressureEvidence == nil {
		return fmt.Errorf("required field sqlPressureEvidence missing")
	}
	if all.K8sEvidence == nil {
		return fmt.Errorf("required field k8sEvidence missing")
	}
	if all.ReadonlyBoundary == nil {
		return fmt.Errorf("required field readonlyBoundary missing")
	}
	if all.NextActions == nil {
		return fmt.Errorf("required field nextActions missing")
	}
	if all.CannotProve == nil {
		return fmt.Errorf("required field cannotProve missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"state", "reasonCode", "dataQuality", "summary", "collectedAt", "timeWindow", "primaryCause", "confidence", "riskLevel", "podCPUEvidence", "dbActivityEvidence", "sqlPressureEvidence", "k8sEvidence", "readonlyBoundary", "collectionError", "nextActions", "cannotProve"})
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
	if !all.DataQuality.IsValid() {
		hasInvalidField = true
	} else {
		o.DataQuality = *all.DataQuality
	}
	o.Summary = *all.Summary
	o.CollectedAt = *all.CollectedAt
	if all.TimeWindow.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TimeWindow = *all.TimeWindow
	if !all.PrimaryCause.IsValid() {
		hasInvalidField = true
	} else {
		o.PrimaryCause = *all.PrimaryCause
	}
	if !all.Confidence.IsValid() {
		hasInvalidField = true
	} else {
		o.Confidence = *all.Confidence
	}
	if !all.RiskLevel.IsValid() {
		hasInvalidField = true
	} else {
		o.RiskLevel = *all.RiskLevel
	}
	if all.PodCpuEvidence.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.PodCpuEvidence = *all.PodCpuEvidence
	if all.DbActivityEvidence.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DbActivityEvidence = *all.DbActivityEvidence
	if all.SqlPressureEvidence.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SqlPressureEvidence = *all.SqlPressureEvidence
	o.K8sEvidence = *all.K8sEvidence
	if all.ReadonlyBoundary.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ReadonlyBoundary = *all.ReadonlyBoundary
	if all.CollectionError != nil && all.CollectionError.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CollectionError = all.CollectionError
	o.NextActions = *all.NextActions
	o.CannotProve = *all.CannotProve

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
