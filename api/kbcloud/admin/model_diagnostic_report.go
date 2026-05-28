// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticReport Layered diagnostic report.
type DiagnosticReport struct {
	// Diagnostic task ID.
	TaskId string `json:"taskId"`
	// Report generation time.
	GeneratedAt time.Time `json:"generatedAt"`
	// Result status for one diagnostic rule.
	OverallStatus DiagnosticRuleStatus `json:"overallStatus"`
	// Human-readable top-level report summary.
	Summary string `json:"summary"`
	// Diagnostic report layer.
	RecommendedLayer DiagnosticLayer `json:"recommendedLayer"`
	// Suggested first responder for the highest-priority finding.
	OwnerHint string `json:"ownerHint"`
	// Whether a redacted diagnostic package can be downloaded for this task.
	ExportAvailable bool `json:"exportAvailable"`
	// Deterministic 0-100 health score for the first screen.
	HealthScore int64 `json:"healthScore"`
	// Health score bucket for the report first screen.
	HealthBucket DiagnosticHealthBucket `json:"healthBucket"`
	// Top prioritized issues for the first screen.
	TopIssues []DiagnosticTopIssue `json:"topIssues"`
	// Time-ordered diagnostic timeline.
	Timeline []DiagnosticTimelineItem `json:"timeline"`
	// First-screen database performance summary.
	DbPerformance DiagnosticDBPerformance `json:"dbPerformance"`
	// Sprint 24 K8s / Operator / Event external evidence explanation layer.
	PeripheralEvidence *DiagnosticPeripheralEvidence `json:"peripheralEvidence,omitempty"`
	// Sprint 25 K8s environment diagnosis page contract. It groups the Sprint 24
	// peripheral evidence explanations by environment layer without changing the
	// peripheralEvidence compatibility field.
	//
	EnvironmentDiagnosis *DiagnosticEnvironmentDiagnosis `json:"environmentDiagnosis,omitempty"`
	// Sprint 26 PostgreSQL readonly space analysis contract. Database Top objects
	// and PVC/PV capacity evidence are intentionally separated so Console does
	// not imply that a PVC capacity symptom proves one table or index is large.
	//
	SpaceAnalysis *DiagnosticSpaceAnalysis `json:"spaceAnalysis,omitempty"`
	// Sprint 27 PostgreSQL readonly CPU analysis contract. CPU, SQL, session, and
	// Kubernetes evidence are intentionally separated so Console does not turn a
	// correlation into a root-cause proof.
	//
	CpuAnalysis *DiagnosticCPUAnalysis `json:"cpuAnalysis,omitempty"`
	Findings    []DiagnosticFinding    `json:"findings"`
	Checks      []DiagnosticCheck      `json:"checks"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticReport instantiates a new DiagnosticReport object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticReport(taskId string, generatedAt time.Time, overallStatus DiagnosticRuleStatus, summary string, recommendedLayer DiagnosticLayer, ownerHint string, exportAvailable bool, healthScore int64, healthBucket DiagnosticHealthBucket, topIssues []DiagnosticTopIssue, timeline []DiagnosticTimelineItem, dbPerformance DiagnosticDBPerformance, findings []DiagnosticFinding, checks []DiagnosticCheck) *DiagnosticReport {
	this := DiagnosticReport{}
	this.TaskId = taskId
	this.GeneratedAt = generatedAt
	this.OverallStatus = overallStatus
	this.Summary = summary
	this.RecommendedLayer = recommendedLayer
	this.OwnerHint = ownerHint
	this.ExportAvailable = exportAvailable
	this.HealthScore = healthScore
	this.HealthBucket = healthBucket
	this.TopIssues = topIssues
	this.Timeline = timeline
	this.DbPerformance = dbPerformance
	this.Findings = findings
	this.Checks = checks
	return &this
}

// NewDiagnosticReportWithDefaults instantiates a new DiagnosticReport object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticReportWithDefaults() *DiagnosticReport {
	this := DiagnosticReport{}
	return &this
}

// GetTaskId returns the TaskId field value.
func (o *DiagnosticReport) GetTaskId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetTaskIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaskId, true
}

// SetTaskId sets field value.
func (o *DiagnosticReport) SetTaskId(v string) {
	o.TaskId = v
}

// GetGeneratedAt returns the GeneratedAt field value.
func (o *DiagnosticReport) GetGeneratedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.GeneratedAt
}

// GetGeneratedAtOk returns a tuple with the GeneratedAt field value
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetGeneratedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GeneratedAt, true
}

// SetGeneratedAt sets field value.
func (o *DiagnosticReport) SetGeneratedAt(v time.Time) {
	o.GeneratedAt = v
}

// GetOverallStatus returns the OverallStatus field value.
func (o *DiagnosticReport) GetOverallStatus() DiagnosticRuleStatus {
	if o == nil {
		var ret DiagnosticRuleStatus
		return ret
	}
	return o.OverallStatus
}

// GetOverallStatusOk returns a tuple with the OverallStatus field value
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetOverallStatusOk() (*DiagnosticRuleStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OverallStatus, true
}

// SetOverallStatus sets field value.
func (o *DiagnosticReport) SetOverallStatus(v DiagnosticRuleStatus) {
	o.OverallStatus = v
}

// GetSummary returns the Summary field value.
func (o *DiagnosticReport) GetSummary() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value.
func (o *DiagnosticReport) SetSummary(v string) {
	o.Summary = v
}

// GetRecommendedLayer returns the RecommendedLayer field value.
func (o *DiagnosticReport) GetRecommendedLayer() DiagnosticLayer {
	if o == nil {
		var ret DiagnosticLayer
		return ret
	}
	return o.RecommendedLayer
}

// GetRecommendedLayerOk returns a tuple with the RecommendedLayer field value
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetRecommendedLayerOk() (*DiagnosticLayer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecommendedLayer, true
}

// SetRecommendedLayer sets field value.
func (o *DiagnosticReport) SetRecommendedLayer(v DiagnosticLayer) {
	o.RecommendedLayer = v
}

// GetOwnerHint returns the OwnerHint field value.
func (o *DiagnosticReport) GetOwnerHint() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OwnerHint
}

// GetOwnerHintOk returns a tuple with the OwnerHint field value
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetOwnerHintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerHint, true
}

// SetOwnerHint sets field value.
func (o *DiagnosticReport) SetOwnerHint(v string) {
	o.OwnerHint = v
}

// GetExportAvailable returns the ExportAvailable field value.
func (o *DiagnosticReport) GetExportAvailable() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.ExportAvailable
}

// GetExportAvailableOk returns a tuple with the ExportAvailable field value
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetExportAvailableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExportAvailable, true
}

// SetExportAvailable sets field value.
func (o *DiagnosticReport) SetExportAvailable(v bool) {
	o.ExportAvailable = v
}

// GetHealthScore returns the HealthScore field value.
func (o *DiagnosticReport) GetHealthScore() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.HealthScore
}

// GetHealthScoreOk returns a tuple with the HealthScore field value
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetHealthScoreOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HealthScore, true
}

// SetHealthScore sets field value.
func (o *DiagnosticReport) SetHealthScore(v int64) {
	o.HealthScore = v
}

// GetHealthBucket returns the HealthBucket field value.
func (o *DiagnosticReport) GetHealthBucket() DiagnosticHealthBucket {
	if o == nil {
		var ret DiagnosticHealthBucket
		return ret
	}
	return o.HealthBucket
}

// GetHealthBucketOk returns a tuple with the HealthBucket field value
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetHealthBucketOk() (*DiagnosticHealthBucket, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HealthBucket, true
}

// SetHealthBucket sets field value.
func (o *DiagnosticReport) SetHealthBucket(v DiagnosticHealthBucket) {
	o.HealthBucket = v
}

// GetTopIssues returns the TopIssues field value.
func (o *DiagnosticReport) GetTopIssues() []DiagnosticTopIssue {
	if o == nil {
		var ret []DiagnosticTopIssue
		return ret
	}
	return o.TopIssues
}

// GetTopIssuesOk returns a tuple with the TopIssues field value
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetTopIssuesOk() (*[]DiagnosticTopIssue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TopIssues, true
}

// SetTopIssues sets field value.
func (o *DiagnosticReport) SetTopIssues(v []DiagnosticTopIssue) {
	o.TopIssues = v
}

// GetTimeline returns the Timeline field value.
func (o *DiagnosticReport) GetTimeline() []DiagnosticTimelineItem {
	if o == nil {
		var ret []DiagnosticTimelineItem
		return ret
	}
	return o.Timeline
}

// GetTimelineOk returns a tuple with the Timeline field value
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetTimelineOk() (*[]DiagnosticTimelineItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timeline, true
}

// SetTimeline sets field value.
func (o *DiagnosticReport) SetTimeline(v []DiagnosticTimelineItem) {
	o.Timeline = v
}

// GetDbPerformance returns the DbPerformance field value.
func (o *DiagnosticReport) GetDbPerformance() DiagnosticDBPerformance {
	if o == nil {
		var ret DiagnosticDBPerformance
		return ret
	}
	return o.DbPerformance
}

// GetDbPerformanceOk returns a tuple with the DbPerformance field value
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetDbPerformanceOk() (*DiagnosticDBPerformance, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DbPerformance, true
}

// SetDbPerformance sets field value.
func (o *DiagnosticReport) SetDbPerformance(v DiagnosticDBPerformance) {
	o.DbPerformance = v
}

// GetPeripheralEvidence returns the PeripheralEvidence field value if set, zero value otherwise.
func (o *DiagnosticReport) GetPeripheralEvidence() DiagnosticPeripheralEvidence {
	if o == nil || o.PeripheralEvidence == nil {
		var ret DiagnosticPeripheralEvidence
		return ret
	}
	return *o.PeripheralEvidence
}

// GetPeripheralEvidenceOk returns a tuple with the PeripheralEvidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetPeripheralEvidenceOk() (*DiagnosticPeripheralEvidence, bool) {
	if o == nil || o.PeripheralEvidence == nil {
		return nil, false
	}
	return o.PeripheralEvidence, true
}

// HasPeripheralEvidence returns a boolean if a field has been set.
func (o *DiagnosticReport) HasPeripheralEvidence() bool {
	return o != nil && o.PeripheralEvidence != nil
}

// SetPeripheralEvidence gets a reference to the given DiagnosticPeripheralEvidence and assigns it to the PeripheralEvidence field.
func (o *DiagnosticReport) SetPeripheralEvidence(v DiagnosticPeripheralEvidence) {
	o.PeripheralEvidence = &v
}

// GetEnvironmentDiagnosis returns the EnvironmentDiagnosis field value if set, zero value otherwise.
func (o *DiagnosticReport) GetEnvironmentDiagnosis() DiagnosticEnvironmentDiagnosis {
	if o == nil || o.EnvironmentDiagnosis == nil {
		var ret DiagnosticEnvironmentDiagnosis
		return ret
	}
	return *o.EnvironmentDiagnosis
}

// GetEnvironmentDiagnosisOk returns a tuple with the EnvironmentDiagnosis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetEnvironmentDiagnosisOk() (*DiagnosticEnvironmentDiagnosis, bool) {
	if o == nil || o.EnvironmentDiagnosis == nil {
		return nil, false
	}
	return o.EnvironmentDiagnosis, true
}

// HasEnvironmentDiagnosis returns a boolean if a field has been set.
func (o *DiagnosticReport) HasEnvironmentDiagnosis() bool {
	return o != nil && o.EnvironmentDiagnosis != nil
}

// SetEnvironmentDiagnosis gets a reference to the given DiagnosticEnvironmentDiagnosis and assigns it to the EnvironmentDiagnosis field.
func (o *DiagnosticReport) SetEnvironmentDiagnosis(v DiagnosticEnvironmentDiagnosis) {
	o.EnvironmentDiagnosis = &v
}

// GetSpaceAnalysis returns the SpaceAnalysis field value if set, zero value otherwise.
func (o *DiagnosticReport) GetSpaceAnalysis() DiagnosticSpaceAnalysis {
	if o == nil || o.SpaceAnalysis == nil {
		var ret DiagnosticSpaceAnalysis
		return ret
	}
	return *o.SpaceAnalysis
}

// GetSpaceAnalysisOk returns a tuple with the SpaceAnalysis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetSpaceAnalysisOk() (*DiagnosticSpaceAnalysis, bool) {
	if o == nil || o.SpaceAnalysis == nil {
		return nil, false
	}
	return o.SpaceAnalysis, true
}

// HasSpaceAnalysis returns a boolean if a field has been set.
func (o *DiagnosticReport) HasSpaceAnalysis() bool {
	return o != nil && o.SpaceAnalysis != nil
}

// SetSpaceAnalysis gets a reference to the given DiagnosticSpaceAnalysis and assigns it to the SpaceAnalysis field.
func (o *DiagnosticReport) SetSpaceAnalysis(v DiagnosticSpaceAnalysis) {
	o.SpaceAnalysis = &v
}

// GetCpuAnalysis returns the CpuAnalysis field value if set, zero value otherwise.
func (o *DiagnosticReport) GetCpuAnalysis() DiagnosticCPUAnalysis {
	if o == nil || o.CpuAnalysis == nil {
		var ret DiagnosticCPUAnalysis
		return ret
	}
	return *o.CpuAnalysis
}

// GetCpuAnalysisOk returns a tuple with the CpuAnalysis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetCpuAnalysisOk() (*DiagnosticCPUAnalysis, bool) {
	if o == nil || o.CpuAnalysis == nil {
		return nil, false
	}
	return o.CpuAnalysis, true
}

// HasCpuAnalysis returns a boolean if a field has been set.
func (o *DiagnosticReport) HasCpuAnalysis() bool {
	return o != nil && o.CpuAnalysis != nil
}

// SetCpuAnalysis gets a reference to the given DiagnosticCPUAnalysis and assigns it to the CpuAnalysis field.
func (o *DiagnosticReport) SetCpuAnalysis(v DiagnosticCPUAnalysis) {
	o.CpuAnalysis = &v
}

// GetFindings returns the Findings field value.
func (o *DiagnosticReport) GetFindings() []DiagnosticFinding {
	if o == nil {
		var ret []DiagnosticFinding
		return ret
	}
	return o.Findings
}

// GetFindingsOk returns a tuple with the Findings field value
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetFindingsOk() (*[]DiagnosticFinding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Findings, true
}

// SetFindings sets field value.
func (o *DiagnosticReport) SetFindings(v []DiagnosticFinding) {
	o.Findings = v
}

// GetChecks returns the Checks field value.
func (o *DiagnosticReport) GetChecks() []DiagnosticCheck {
	if o == nil {
		var ret []DiagnosticCheck
		return ret
	}
	return o.Checks
}

// GetChecksOk returns a tuple with the Checks field value
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetChecksOk() (*[]DiagnosticCheck, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Checks, true
}

// SetChecks sets field value.
func (o *DiagnosticReport) SetChecks(v []DiagnosticCheck) {
	o.Checks = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticReport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["taskId"] = o.TaskId
	if o.GeneratedAt.Nanosecond() == 0 {
		toSerialize["generatedAt"] = o.GeneratedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["generatedAt"] = o.GeneratedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["overallStatus"] = o.OverallStatus
	toSerialize["summary"] = o.Summary
	toSerialize["recommendedLayer"] = o.RecommendedLayer
	toSerialize["ownerHint"] = o.OwnerHint
	toSerialize["exportAvailable"] = o.ExportAvailable
	toSerialize["healthScore"] = o.HealthScore
	toSerialize["healthBucket"] = o.HealthBucket
	toSerialize["topIssues"] = o.TopIssues
	toSerialize["timeline"] = o.Timeline
	toSerialize["dbPerformance"] = o.DbPerformance
	if o.PeripheralEvidence != nil {
		toSerialize["peripheralEvidence"] = o.PeripheralEvidence
	}
	if o.EnvironmentDiagnosis != nil {
		toSerialize["environmentDiagnosis"] = o.EnvironmentDiagnosis
	}
	if o.SpaceAnalysis != nil {
		toSerialize["spaceAnalysis"] = o.SpaceAnalysis
	}
	if o.CpuAnalysis != nil {
		toSerialize["cpuAnalysis"] = o.CpuAnalysis
	}
	toSerialize["findings"] = o.Findings
	toSerialize["checks"] = o.Checks

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticReport) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TaskId               *string                         `json:"taskId"`
		GeneratedAt          *time.Time                      `json:"generatedAt"`
		OverallStatus        *DiagnosticRuleStatus           `json:"overallStatus"`
		Summary              *string                         `json:"summary"`
		RecommendedLayer     *DiagnosticLayer                `json:"recommendedLayer"`
		OwnerHint            *string                         `json:"ownerHint"`
		ExportAvailable      *bool                           `json:"exportAvailable"`
		HealthScore          *int64                          `json:"healthScore"`
		HealthBucket         *DiagnosticHealthBucket         `json:"healthBucket"`
		TopIssues            *[]DiagnosticTopIssue           `json:"topIssues"`
		Timeline             *[]DiagnosticTimelineItem       `json:"timeline"`
		DbPerformance        *DiagnosticDBPerformance        `json:"dbPerformance"`
		PeripheralEvidence   *DiagnosticPeripheralEvidence   `json:"peripheralEvidence,omitempty"`
		EnvironmentDiagnosis *DiagnosticEnvironmentDiagnosis `json:"environmentDiagnosis,omitempty"`
		SpaceAnalysis        *DiagnosticSpaceAnalysis        `json:"spaceAnalysis,omitempty"`
		CpuAnalysis          *DiagnosticCPUAnalysis          `json:"cpuAnalysis,omitempty"`
		Findings             *[]DiagnosticFinding            `json:"findings"`
		Checks               *[]DiagnosticCheck              `json:"checks"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.TaskId == nil {
		return fmt.Errorf("required field taskId missing")
	}
	if all.GeneratedAt == nil {
		return fmt.Errorf("required field generatedAt missing")
	}
	if all.OverallStatus == nil {
		return fmt.Errorf("required field overallStatus missing")
	}
	if all.Summary == nil {
		return fmt.Errorf("required field summary missing")
	}
	if all.RecommendedLayer == nil {
		return fmt.Errorf("required field recommendedLayer missing")
	}
	if all.OwnerHint == nil {
		return fmt.Errorf("required field ownerHint missing")
	}
	if all.ExportAvailable == nil {
		return fmt.Errorf("required field exportAvailable missing")
	}
	if all.HealthScore == nil {
		return fmt.Errorf("required field healthScore missing")
	}
	if all.HealthBucket == nil {
		return fmt.Errorf("required field healthBucket missing")
	}
	if all.TopIssues == nil {
		return fmt.Errorf("required field topIssues missing")
	}
	if all.Timeline == nil {
		return fmt.Errorf("required field timeline missing")
	}
	if all.DbPerformance == nil {
		return fmt.Errorf("required field dbPerformance missing")
	}
	if all.Findings == nil {
		return fmt.Errorf("required field findings missing")
	}
	if all.Checks == nil {
		return fmt.Errorf("required field checks missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"taskId", "generatedAt", "overallStatus", "summary", "recommendedLayer", "ownerHint", "exportAvailable", "healthScore", "healthBucket", "topIssues", "timeline", "dbPerformance", "peripheralEvidence", "environmentDiagnosis", "spaceAnalysis", "cpuAnalysis", "findings", "checks"})
	} else {
		return err
	}

	hasInvalidField := false
	o.TaskId = *all.TaskId
	o.GeneratedAt = *all.GeneratedAt
	if !all.OverallStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.OverallStatus = *all.OverallStatus
	}
	o.Summary = *all.Summary
	if !all.RecommendedLayer.IsValid() {
		hasInvalidField = true
	} else {
		o.RecommendedLayer = *all.RecommendedLayer
	}
	o.OwnerHint = *all.OwnerHint
	o.ExportAvailable = *all.ExportAvailable
	o.HealthScore = *all.HealthScore
	if !all.HealthBucket.IsValid() {
		hasInvalidField = true
	} else {
		o.HealthBucket = *all.HealthBucket
	}
	o.TopIssues = *all.TopIssues
	o.Timeline = *all.Timeline
	if all.DbPerformance.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DbPerformance = *all.DbPerformance
	if all.PeripheralEvidence != nil && all.PeripheralEvidence.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.PeripheralEvidence = all.PeripheralEvidence
	if all.EnvironmentDiagnosis != nil && all.EnvironmentDiagnosis.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.EnvironmentDiagnosis = all.EnvironmentDiagnosis
	if all.SpaceAnalysis != nil && all.SpaceAnalysis.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SpaceAnalysis = all.SpaceAnalysis
	if all.CpuAnalysis != nil && all.CpuAnalysis.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CpuAnalysis = all.CpuAnalysis
	o.Findings = *all.Findings
	o.Checks = *all.Checks

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
