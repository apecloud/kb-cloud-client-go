// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type AiAgentArtifactBatchGetRequest struct {
	Reports      []string `json:"reports,omitempty"`
	Evidence     []string `json:"evidence,omitempty"`
	ActionPlans  []string `json:"actionPlans,omitempty"`
	RunSummaries []string `json:"runSummaries,omitempty"`
	MetricSeries []string `json:"metricSeries,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentArtifactBatchGetRequest instantiates a new AiAgentArtifactBatchGetRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentArtifactBatchGetRequest() *AiAgentArtifactBatchGetRequest {
	this := AiAgentArtifactBatchGetRequest{}
	return &this
}

// NewAiAgentArtifactBatchGetRequestWithDefaults instantiates a new AiAgentArtifactBatchGetRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentArtifactBatchGetRequestWithDefaults() *AiAgentArtifactBatchGetRequest {
	this := AiAgentArtifactBatchGetRequest{}
	return &this
}

// GetReports returns the Reports field value if set, zero value otherwise.
func (o *AiAgentArtifactBatchGetRequest) GetReports() []string {
	if o == nil || o.Reports == nil {
		var ret []string
		return ret
	}
	return o.Reports
}

// GetReportsOk returns a tuple with the Reports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentArtifactBatchGetRequest) GetReportsOk() (*[]string, bool) {
	if o == nil || o.Reports == nil {
		return nil, false
	}
	return &o.Reports, true
}

// HasReports returns a boolean if a field has been set.
func (o *AiAgentArtifactBatchGetRequest) HasReports() bool {
	return o != nil && o.Reports != nil
}

// SetReports gets a reference to the given []string and assigns it to the Reports field.
func (o *AiAgentArtifactBatchGetRequest) SetReports(v []string) {
	o.Reports = v
}

// GetEvidence returns the Evidence field value if set, zero value otherwise.
func (o *AiAgentArtifactBatchGetRequest) GetEvidence() []string {
	if o == nil || o.Evidence == nil {
		var ret []string
		return ret
	}
	return o.Evidence
}

// GetEvidenceOk returns a tuple with the Evidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentArtifactBatchGetRequest) GetEvidenceOk() (*[]string, bool) {
	if o == nil || o.Evidence == nil {
		return nil, false
	}
	return &o.Evidence, true
}

// HasEvidence returns a boolean if a field has been set.
func (o *AiAgentArtifactBatchGetRequest) HasEvidence() bool {
	return o != nil && o.Evidence != nil
}

// SetEvidence gets a reference to the given []string and assigns it to the Evidence field.
func (o *AiAgentArtifactBatchGetRequest) SetEvidence(v []string) {
	o.Evidence = v
}

// GetActionPlans returns the ActionPlans field value if set, zero value otherwise.
func (o *AiAgentArtifactBatchGetRequest) GetActionPlans() []string {
	if o == nil || o.ActionPlans == nil {
		var ret []string
		return ret
	}
	return o.ActionPlans
}

// GetActionPlansOk returns a tuple with the ActionPlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentArtifactBatchGetRequest) GetActionPlansOk() (*[]string, bool) {
	if o == nil || o.ActionPlans == nil {
		return nil, false
	}
	return &o.ActionPlans, true
}

// HasActionPlans returns a boolean if a field has been set.
func (o *AiAgentArtifactBatchGetRequest) HasActionPlans() bool {
	return o != nil && o.ActionPlans != nil
}

// SetActionPlans gets a reference to the given []string and assigns it to the ActionPlans field.
func (o *AiAgentArtifactBatchGetRequest) SetActionPlans(v []string) {
	o.ActionPlans = v
}

// GetRunSummaries returns the RunSummaries field value if set, zero value otherwise.
func (o *AiAgentArtifactBatchGetRequest) GetRunSummaries() []string {
	if o == nil || o.RunSummaries == nil {
		var ret []string
		return ret
	}
	return o.RunSummaries
}

// GetRunSummariesOk returns a tuple with the RunSummaries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentArtifactBatchGetRequest) GetRunSummariesOk() (*[]string, bool) {
	if o == nil || o.RunSummaries == nil {
		return nil, false
	}
	return &o.RunSummaries, true
}

// HasRunSummaries returns a boolean if a field has been set.
func (o *AiAgentArtifactBatchGetRequest) HasRunSummaries() bool {
	return o != nil && o.RunSummaries != nil
}

// SetRunSummaries gets a reference to the given []string and assigns it to the RunSummaries field.
func (o *AiAgentArtifactBatchGetRequest) SetRunSummaries(v []string) {
	o.RunSummaries = v
}

// GetMetricSeries returns the MetricSeries field value if set, zero value otherwise.
func (o *AiAgentArtifactBatchGetRequest) GetMetricSeries() []string {
	if o == nil || o.MetricSeries == nil {
		var ret []string
		return ret
	}
	return o.MetricSeries
}

// GetMetricSeriesOk returns a tuple with the MetricSeries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentArtifactBatchGetRequest) GetMetricSeriesOk() (*[]string, bool) {
	if o == nil || o.MetricSeries == nil {
		return nil, false
	}
	return &o.MetricSeries, true
}

// HasMetricSeries returns a boolean if a field has been set.
func (o *AiAgentArtifactBatchGetRequest) HasMetricSeries() bool {
	return o != nil && o.MetricSeries != nil
}

// SetMetricSeries gets a reference to the given []string and assigns it to the MetricSeries field.
func (o *AiAgentArtifactBatchGetRequest) SetMetricSeries(v []string) {
	o.MetricSeries = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentArtifactBatchGetRequest) MarshalJSON() ([]byte, error) {
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
func (o *AiAgentArtifactBatchGetRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Reports      []string `json:"reports,omitempty"`
		Evidence     []string `json:"evidence,omitempty"`
		ActionPlans  []string `json:"actionPlans,omitempty"`
		RunSummaries []string `json:"runSummaries,omitempty"`
		MetricSeries []string `json:"metricSeries,omitempty"`
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
