// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticDBPerformanceMetricDetail User-facing detail and DBA guidance for one database performance card.
// It explains why the metric matters, what evidence was checked, and the
// next action without requiring Console to infer meaning from rawEvidence.
// SQL text in this structure must be redacted or summarized.
type DiagnosticDBPerformanceMetricDetail struct {
	// Short detail summary shown near the expanded card.
	Summary string `json:"summary"`
	// Human-readable summary of the key structured evidence or checked data source.
	EvidenceSummary string `json:"evidenceSummary"`
	// DBA-oriented first action based on this metric state.
	Recommendation string `json:"recommendation"`
	// Structured next actions for this metric detail.
	NextActions []DiagnosticNextAction `json:"nextActions"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticDBPerformanceMetricDetail instantiates a new DiagnosticDBPerformanceMetricDetail object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticDBPerformanceMetricDetail(summary string, evidenceSummary string, recommendation string, nextActions []DiagnosticNextAction) *DiagnosticDBPerformanceMetricDetail {
	this := DiagnosticDBPerformanceMetricDetail{}
	this.Summary = summary
	this.EvidenceSummary = evidenceSummary
	this.Recommendation = recommendation
	this.NextActions = nextActions
	return &this
}

// NewDiagnosticDBPerformanceMetricDetailWithDefaults instantiates a new DiagnosticDBPerformanceMetricDetail object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticDBPerformanceMetricDetailWithDefaults() *DiagnosticDBPerformanceMetricDetail {
	this := DiagnosticDBPerformanceMetricDetail{}
	return &this
}

// GetSummary returns the Summary field value.
func (o *DiagnosticDBPerformanceMetricDetail) GetSummary() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceMetricDetail) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value.
func (o *DiagnosticDBPerformanceMetricDetail) SetSummary(v string) {
	o.Summary = v
}

// GetEvidenceSummary returns the EvidenceSummary field value.
func (o *DiagnosticDBPerformanceMetricDetail) GetEvidenceSummary() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EvidenceSummary
}

// GetEvidenceSummaryOk returns a tuple with the EvidenceSummary field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceMetricDetail) GetEvidenceSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EvidenceSummary, true
}

// SetEvidenceSummary sets field value.
func (o *DiagnosticDBPerformanceMetricDetail) SetEvidenceSummary(v string) {
	o.EvidenceSummary = v
}

// GetRecommendation returns the Recommendation field value.
func (o *DiagnosticDBPerformanceMetricDetail) GetRecommendation() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Recommendation
}

// GetRecommendationOk returns a tuple with the Recommendation field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceMetricDetail) GetRecommendationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recommendation, true
}

// SetRecommendation sets field value.
func (o *DiagnosticDBPerformanceMetricDetail) SetRecommendation(v string) {
	o.Recommendation = v
}

// GetNextActions returns the NextActions field value.
func (o *DiagnosticDBPerformanceMetricDetail) GetNextActions() []DiagnosticNextAction {
	if o == nil {
		var ret []DiagnosticNextAction
		return ret
	}
	return o.NextActions
}

// GetNextActionsOk returns a tuple with the NextActions field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceMetricDetail) GetNextActionsOk() (*[]DiagnosticNextAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextActions, true
}

// SetNextActions sets field value.
func (o *DiagnosticDBPerformanceMetricDetail) SetNextActions(v []DiagnosticNextAction) {
	o.NextActions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticDBPerformanceMetricDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["summary"] = o.Summary
	toSerialize["evidenceSummary"] = o.EvidenceSummary
	toSerialize["recommendation"] = o.Recommendation
	toSerialize["nextActions"] = o.NextActions

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticDBPerformanceMetricDetail) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Summary         *string                 `json:"summary"`
		EvidenceSummary *string                 `json:"evidenceSummary"`
		Recommendation  *string                 `json:"recommendation"`
		NextActions     *[]DiagnosticNextAction `json:"nextActions"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Summary == nil {
		return fmt.Errorf("required field summary missing")
	}
	if all.EvidenceSummary == nil {
		return fmt.Errorf("required field evidenceSummary missing")
	}
	if all.Recommendation == nil {
		return fmt.Errorf("required field recommendation missing")
	}
	if all.NextActions == nil {
		return fmt.Errorf("required field nextActions missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"summary", "evidenceSummary", "recommendation", "nextActions"})
	} else {
		return err
	}
	o.Summary = *all.Summary
	o.EvidenceSummary = *all.EvidenceSummary
	o.Recommendation = *all.Recommendation
	o.NextActions = *all.NextActions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
