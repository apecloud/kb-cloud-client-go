// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type AiAgentArtifactBatch struct {
	Reports      map[string]AiAgentReport         `json:"reports,omitempty"`
	Evidence     map[string]AiAgentEvidenceDetail `json:"evidence,omitempty"`
	ActionPlans  map[string]AiAgentActionPlan     `json:"actionPlans,omitempty"`
	RunSummaries map[string]AiAgentRun            `json:"runSummaries,omitempty"`
	MetricSeries map[string]AiAgentMetricSeries   `json:"metricSeries,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentArtifactBatch instantiates a new AiAgentArtifactBatch object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentArtifactBatch() *AiAgentArtifactBatch {
	this := AiAgentArtifactBatch{}
	return &this
}

// NewAiAgentArtifactBatchWithDefaults instantiates a new AiAgentArtifactBatch object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentArtifactBatchWithDefaults() *AiAgentArtifactBatch {
	this := AiAgentArtifactBatch{}
	return &this
}

// GetReports returns the Reports field value if set, zero value otherwise.
func (o *AiAgentArtifactBatch) GetReports() map[string]AiAgentReport {
	if o == nil || o.Reports == nil {
		var ret map[string]AiAgentReport
		return ret
	}
	return o.Reports
}

// GetReportsOk returns a tuple with the Reports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentArtifactBatch) GetReportsOk() (*map[string]AiAgentReport, bool) {
	if o == nil || o.Reports == nil {
		return nil, false
	}
	return &o.Reports, true
}

// HasReports returns a boolean if a field has been set.
func (o *AiAgentArtifactBatch) HasReports() bool {
	return o != nil && o.Reports != nil
}

// SetReports gets a reference to the given map[string]AiAgentReport and assigns it to the Reports field.
func (o *AiAgentArtifactBatch) SetReports(v map[string]AiAgentReport) {
	o.Reports = v
}

// GetEvidence returns the Evidence field value if set, zero value otherwise.
func (o *AiAgentArtifactBatch) GetEvidence() map[string]AiAgentEvidenceDetail {
	if o == nil || o.Evidence == nil {
		var ret map[string]AiAgentEvidenceDetail
		return ret
	}
	return o.Evidence
}

// GetEvidenceOk returns a tuple with the Evidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentArtifactBatch) GetEvidenceOk() (*map[string]AiAgentEvidenceDetail, bool) {
	if o == nil || o.Evidence == nil {
		return nil, false
	}
	return &o.Evidence, true
}

// HasEvidence returns a boolean if a field has been set.
func (o *AiAgentArtifactBatch) HasEvidence() bool {
	return o != nil && o.Evidence != nil
}

// SetEvidence gets a reference to the given map[string]AiAgentEvidenceDetail and assigns it to the Evidence field.
func (o *AiAgentArtifactBatch) SetEvidence(v map[string]AiAgentEvidenceDetail) {
	o.Evidence = v
}

// GetActionPlans returns the ActionPlans field value if set, zero value otherwise.
func (o *AiAgentArtifactBatch) GetActionPlans() map[string]AiAgentActionPlan {
	if o == nil || o.ActionPlans == nil {
		var ret map[string]AiAgentActionPlan
		return ret
	}
	return o.ActionPlans
}

// GetActionPlansOk returns a tuple with the ActionPlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentArtifactBatch) GetActionPlansOk() (*map[string]AiAgentActionPlan, bool) {
	if o == nil || o.ActionPlans == nil {
		return nil, false
	}
	return &o.ActionPlans, true
}

// HasActionPlans returns a boolean if a field has been set.
func (o *AiAgentArtifactBatch) HasActionPlans() bool {
	return o != nil && o.ActionPlans != nil
}

// SetActionPlans gets a reference to the given map[string]AiAgentActionPlan and assigns it to the ActionPlans field.
func (o *AiAgentArtifactBatch) SetActionPlans(v map[string]AiAgentActionPlan) {
	o.ActionPlans = v
}

// GetRunSummaries returns the RunSummaries field value if set, zero value otherwise.
func (o *AiAgentArtifactBatch) GetRunSummaries() map[string]AiAgentRun {
	if o == nil || o.RunSummaries == nil {
		var ret map[string]AiAgentRun
		return ret
	}
	return o.RunSummaries
}

// GetRunSummariesOk returns a tuple with the RunSummaries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentArtifactBatch) GetRunSummariesOk() (*map[string]AiAgentRun, bool) {
	if o == nil || o.RunSummaries == nil {
		return nil, false
	}
	return &o.RunSummaries, true
}

// HasRunSummaries returns a boolean if a field has been set.
func (o *AiAgentArtifactBatch) HasRunSummaries() bool {
	return o != nil && o.RunSummaries != nil
}

// SetRunSummaries gets a reference to the given map[string]AiAgentRun and assigns it to the RunSummaries field.
func (o *AiAgentArtifactBatch) SetRunSummaries(v map[string]AiAgentRun) {
	o.RunSummaries = v
}

// GetMetricSeries returns the MetricSeries field value if set, zero value otherwise.
func (o *AiAgentArtifactBatch) GetMetricSeries() map[string]AiAgentMetricSeries {
	if o == nil || o.MetricSeries == nil {
		var ret map[string]AiAgentMetricSeries
		return ret
	}
	return o.MetricSeries
}

// GetMetricSeriesOk returns a tuple with the MetricSeries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentArtifactBatch) GetMetricSeriesOk() (*map[string]AiAgentMetricSeries, bool) {
	if o == nil || o.MetricSeries == nil {
		return nil, false
	}
	return &o.MetricSeries, true
}

// HasMetricSeries returns a boolean if a field has been set.
func (o *AiAgentArtifactBatch) HasMetricSeries() bool {
	return o != nil && o.MetricSeries != nil
}

// SetMetricSeries gets a reference to the given map[string]AiAgentMetricSeries and assigns it to the MetricSeries field.
func (o *AiAgentArtifactBatch) SetMetricSeries(v map[string]AiAgentMetricSeries) {
	o.MetricSeries = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentArtifactBatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Reports != nil {
		toSerialize["reports"] = o.Reports
	}
	if o.Evidence != nil {
		toSerialize["evidence"] = o.Evidence
	}
	if o.ActionPlans != nil {
		toSerialize["actionPlans"] = o.ActionPlans
	}
	if o.RunSummaries != nil {
		toSerialize["runSummaries"] = o.RunSummaries
	}
	if o.MetricSeries != nil {
		toSerialize["metricSeries"] = o.MetricSeries
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentArtifactBatch) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Reports      map[string]AiAgentReport         `json:"reports,omitempty"`
		Evidence     map[string]AiAgentEvidenceDetail `json:"evidence,omitempty"`
		ActionPlans  map[string]AiAgentActionPlan     `json:"actionPlans,omitempty"`
		RunSummaries map[string]AiAgentRun            `json:"runSummaries,omitempty"`
		MetricSeries map[string]AiAgentMetricSeries   `json:"metricSeries,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"reports", "evidence", "actionPlans", "runSummaries", "metricSeries"})
	} else {
		return err
	}
	o.Reports = all.Reports
	o.Evidence = all.Evidence
	o.ActionPlans = all.ActionPlans
	o.RunSummaries = all.RunSummaries
	o.MetricSeries = all.MetricSeries

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
