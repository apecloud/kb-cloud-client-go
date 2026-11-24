// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ClickhouseSQLAnalysis struct {
	// Collection time in RFC3339 UTC.
	CollectedAt string `json:"collectedAt"`
	// available when system.query_log was collected, unavailable when query logging is disabled or the table is missing.
	Status string `json:"status"`
	// Stable reason code when status is unavailable, for example query_log_unavailable.
	UnavailableReason *string `json:"unavailableReason,omitempty"`
	// Human-readable guidance when the analysis is unavailable.
	Message *string                     `json:"message,omitempty"`
	Range   *ClickhouseSQLAnalysisRange `json:"range,omitempty"`
	Summary *ClickhouseSQLSummary       `json:"summary,omitempty"`
	// Top initial queries ordered by query_duration_ms descending. Distributed child queries are excluded via is_initial_query = 1.
	SlowestQueries []ClickhouseSlowQuery `json:"slowestQueries"`
	// Query patterns grouped by normalized_query_hash and ordered by the requested metric.
	Patterns []ClickhouseQueryPattern `json:"patterns"`
	// Failed queries (type = ExceptionWhileProcessing) ordered by event_time descending.
	Failures []ClickhouseFailedQuery `json:"failures"`
	// Per-source collection status.
	Sources []ClickhouseSource `json:"sources"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClickhouseSQLAnalysis instantiates a new ClickhouseSQLAnalysis object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClickhouseSQLAnalysis(collectedAt string, status string, slowestQueries []ClickhouseSlowQuery, patterns []ClickhouseQueryPattern, failures []ClickhouseFailedQuery, sources []ClickhouseSource) *ClickhouseSQLAnalysis {
	this := ClickhouseSQLAnalysis{}
	this.CollectedAt = collectedAt
	this.Status = status
	this.SlowestQueries = slowestQueries
	this.Patterns = patterns
	this.Failures = failures
	this.Sources = sources
	return &this
}

// NewClickhouseSQLAnalysisWithDefaults instantiates a new ClickhouseSQLAnalysis object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClickhouseSQLAnalysisWithDefaults() *ClickhouseSQLAnalysis {
	this := ClickhouseSQLAnalysis{}
	return &this
}

// GetCollectedAt returns the CollectedAt field value.
func (o *ClickhouseSQLAnalysis) GetCollectedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CollectedAt
}

// GetCollectedAtOk returns a tuple with the CollectedAt field value
// and a boolean to check if the value has been set.
func (o *ClickhouseSQLAnalysis) GetCollectedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectedAt, true
}

// SetCollectedAt sets field value.
func (o *ClickhouseSQLAnalysis) SetCollectedAt(v string) {
	o.CollectedAt = v
}

// GetStatus returns the Status field value.
func (o *ClickhouseSQLAnalysis) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ClickhouseSQLAnalysis) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *ClickhouseSQLAnalysis) SetStatus(v string) {
	o.Status = v
}

// GetUnavailableReason returns the UnavailableReason field value if set, zero value otherwise.
func (o *ClickhouseSQLAnalysis) GetUnavailableReason() string {
	if o == nil || o.UnavailableReason == nil {
		var ret string
		return ret
	}
	return *o.UnavailableReason
}

// GetUnavailableReasonOk returns a tuple with the UnavailableReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSQLAnalysis) GetUnavailableReasonOk() (*string, bool) {
	if o == nil || o.UnavailableReason == nil {
		return nil, false
	}
	return o.UnavailableReason, true
}

// HasUnavailableReason returns a boolean if a field has been set.
func (o *ClickhouseSQLAnalysis) HasUnavailableReason() bool {
	return o != nil && o.UnavailableReason != nil
}

// SetUnavailableReason gets a reference to the given string and assigns it to the UnavailableReason field.
func (o *ClickhouseSQLAnalysis) SetUnavailableReason(v string) {
	o.UnavailableReason = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ClickhouseSQLAnalysis) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSQLAnalysis) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ClickhouseSQLAnalysis) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ClickhouseSQLAnalysis) SetMessage(v string) {
	o.Message = &v
}

// GetRange returns the Range field value if set, zero value otherwise.
func (o *ClickhouseSQLAnalysis) GetRange() ClickhouseSQLAnalysisRange {
	if o == nil || o.Range == nil {
		var ret ClickhouseSQLAnalysisRange
		return ret
	}
	return *o.Range
}

// GetRangeOk returns a tuple with the Range field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSQLAnalysis) GetRangeOk() (*ClickhouseSQLAnalysisRange, bool) {
	if o == nil || o.Range == nil {
		return nil, false
	}
	return o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *ClickhouseSQLAnalysis) HasRange() bool {
	return o != nil && o.Range != nil
}

// SetRange gets a reference to the given ClickhouseSQLAnalysisRange and assigns it to the Range field.
func (o *ClickhouseSQLAnalysis) SetRange(v ClickhouseSQLAnalysisRange) {
	o.Range = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *ClickhouseSQLAnalysis) GetSummary() ClickhouseSQLSummary {
	if o == nil || o.Summary == nil {
		var ret ClickhouseSQLSummary
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSQLAnalysis) GetSummaryOk() (*ClickhouseSQLSummary, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *ClickhouseSQLAnalysis) HasSummary() bool {
	return o != nil && o.Summary != nil
}

// SetSummary gets a reference to the given ClickhouseSQLSummary and assigns it to the Summary field.
func (o *ClickhouseSQLAnalysis) SetSummary(v ClickhouseSQLSummary) {
	o.Summary = &v
}

// GetSlowestQueries returns the SlowestQueries field value.
func (o *ClickhouseSQLAnalysis) GetSlowestQueries() []ClickhouseSlowQuery {
	if o == nil {
		var ret []ClickhouseSlowQuery
		return ret
	}
	return o.SlowestQueries
}

// GetSlowestQueriesOk returns a tuple with the SlowestQueries field value
// and a boolean to check if the value has been set.
func (o *ClickhouseSQLAnalysis) GetSlowestQueriesOk() (*[]ClickhouseSlowQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SlowestQueries, true
}

// SetSlowestQueries sets field value.
func (o *ClickhouseSQLAnalysis) SetSlowestQueries(v []ClickhouseSlowQuery) {
	o.SlowestQueries = v
}

// GetPatterns returns the Patterns field value.
func (o *ClickhouseSQLAnalysis) GetPatterns() []ClickhouseQueryPattern {
	if o == nil {
		var ret []ClickhouseQueryPattern
		return ret
	}
	return o.Patterns
}

// GetPatternsOk returns a tuple with the Patterns field value
// and a boolean to check if the value has been set.
func (o *ClickhouseSQLAnalysis) GetPatternsOk() (*[]ClickhouseQueryPattern, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Patterns, true
}

// SetPatterns sets field value.
func (o *ClickhouseSQLAnalysis) SetPatterns(v []ClickhouseQueryPattern) {
	o.Patterns = v
}

// GetFailures returns the Failures field value.
func (o *ClickhouseSQLAnalysis) GetFailures() []ClickhouseFailedQuery {
	if o == nil {
		var ret []ClickhouseFailedQuery
		return ret
	}
	return o.Failures
}

// GetFailuresOk returns a tuple with the Failures field value
// and a boolean to check if the value has been set.
func (o *ClickhouseSQLAnalysis) GetFailuresOk() (*[]ClickhouseFailedQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Failures, true
}

// SetFailures sets field value.
func (o *ClickhouseSQLAnalysis) SetFailures(v []ClickhouseFailedQuery) {
	o.Failures = v
}

// GetSources returns the Sources field value.
func (o *ClickhouseSQLAnalysis) GetSources() []ClickhouseSource {
	if o == nil {
		var ret []ClickhouseSource
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value
// and a boolean to check if the value has been set.
func (o *ClickhouseSQLAnalysis) GetSourcesOk() (*[]ClickhouseSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sources, true
}

// SetSources sets field value.
func (o *ClickhouseSQLAnalysis) SetSources(v []ClickhouseSource) {
	o.Sources = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClickhouseSQLAnalysis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["collectedAt"] = o.CollectedAt
	toSerialize["status"] = o.Status
	if o.UnavailableReason != nil {
		toSerialize["unavailableReason"] = o.UnavailableReason
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Range != nil {
		toSerialize["range"] = o.Range
	}
	if o.Summary != nil {
		toSerialize["summary"] = o.Summary
	}
	toSerialize["slowestQueries"] = o.SlowestQueries
	toSerialize["patterns"] = o.Patterns
	toSerialize["failures"] = o.Failures
	toSerialize["sources"] = o.Sources

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClickhouseSQLAnalysis) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CollectedAt       *string                     `json:"collectedAt"`
		Status            *string                     `json:"status"`
		UnavailableReason *string                     `json:"unavailableReason,omitempty"`
		Message           *string                     `json:"message,omitempty"`
		Range             *ClickhouseSQLAnalysisRange `json:"range,omitempty"`
		Summary           *ClickhouseSQLSummary       `json:"summary,omitempty"`
		SlowestQueries    *[]ClickhouseSlowQuery      `json:"slowestQueries"`
		Patterns          *[]ClickhouseQueryPattern   `json:"patterns"`
		Failures          *[]ClickhouseFailedQuery    `json:"failures"`
		Sources           *[]ClickhouseSource         `json:"sources"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.CollectedAt == nil {
		return fmt.Errorf("required field collectedAt missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.SlowestQueries == nil {
		return fmt.Errorf("required field slowestQueries missing")
	}
	if all.Patterns == nil {
		return fmt.Errorf("required field patterns missing")
	}
	if all.Failures == nil {
		return fmt.Errorf("required field failures missing")
	}
	if all.Sources == nil {
		return fmt.Errorf("required field sources missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"collectedAt", "status", "unavailableReason", "message", "range", "summary", "slowestQueries", "patterns", "failures", "sources"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CollectedAt = *all.CollectedAt
	o.Status = *all.Status
	o.UnavailableReason = all.UnavailableReason
	o.Message = all.Message
	if all.Range != nil && all.Range.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Range = all.Range
	if all.Summary != nil && all.Summary.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Summary = all.Summary
	o.SlowestQueries = *all.SlowestQueries
	o.Patterns = *all.Patterns
	o.Failures = *all.Failures
	o.Sources = *all.Sources

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
