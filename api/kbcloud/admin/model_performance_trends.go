// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PerformanceTrends struct {
	Range PerformanceTrendRange `json:"range"`
	// Controlled backend query step.
	Granularity string `json:"granularity"`
	// Per-series numeric summary. changeDirection is present only when at least two points exist.
	Summary []PerformanceTrendSummary `json:"summary"`
	// Successfully collected and displayable metric series only. Unsupported, empty, or failed metrics are not represented as per-metric entries.
	Series []PerformanceTrendSeries `json:"series"`
	// Coarse collection source status. Messages are sanitized and do not expose PromQL, endpoints, credentials, or raw internal errors.
	Sources  []PerformanceTrendSource `json:"sources"`
	Warnings []string                 `json:"warnings"`
	// Backend collection timestamp in UTC.
	CollectedAt string `json:"collectedAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPerformanceTrends instantiates a new PerformanceTrends object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPerformanceTrends(rangeVar PerformanceTrendRange, granularity string, summary []PerformanceTrendSummary, series []PerformanceTrendSeries, sources []PerformanceTrendSource, warnings []string, collectedAt string) *PerformanceTrends {
	this := PerformanceTrends{}
	this.Range = rangeVar
	this.Granularity = granularity
	this.Summary = summary
	this.Series = series
	this.Sources = sources
	this.Warnings = warnings
	this.CollectedAt = collectedAt
	return &this
}

// NewPerformanceTrendsWithDefaults instantiates a new PerformanceTrends object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPerformanceTrendsWithDefaults() *PerformanceTrends {
	this := PerformanceTrends{}
	return &this
}

// GetRange returns the Range field value.
func (o *PerformanceTrends) GetRange() PerformanceTrendRange {
	if o == nil {
		var ret PerformanceTrendRange
		return ret
	}
	return o.Range
}

// GetRangeOk returns a tuple with the Range field value
// and a boolean to check if the value has been set.
func (o *PerformanceTrends) GetRangeOk() (*PerformanceTrendRange, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Range, true
}

// SetRange sets field value.
func (o *PerformanceTrends) SetRange(v PerformanceTrendRange) {
	o.Range = v
}

// GetGranularity returns the Granularity field value.
func (o *PerformanceTrends) GetGranularity() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Granularity
}

// GetGranularityOk returns a tuple with the Granularity field value
// and a boolean to check if the value has been set.
func (o *PerformanceTrends) GetGranularityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Granularity, true
}

// SetGranularity sets field value.
func (o *PerformanceTrends) SetGranularity(v string) {
	o.Granularity = v
}

// GetSummary returns the Summary field value.
func (o *PerformanceTrends) GetSummary() []PerformanceTrendSummary {
	if o == nil {
		var ret []PerformanceTrendSummary
		return ret
	}
	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *PerformanceTrends) GetSummaryOk() (*[]PerformanceTrendSummary, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value.
func (o *PerformanceTrends) SetSummary(v []PerformanceTrendSummary) {
	o.Summary = v
}

// GetSeries returns the Series field value.
func (o *PerformanceTrends) GetSeries() []PerformanceTrendSeries {
	if o == nil {
		var ret []PerformanceTrendSeries
		return ret
	}
	return o.Series
}

// GetSeriesOk returns a tuple with the Series field value
// and a boolean to check if the value has been set.
func (o *PerformanceTrends) GetSeriesOk() (*[]PerformanceTrendSeries, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Series, true
}

// SetSeries sets field value.
func (o *PerformanceTrends) SetSeries(v []PerformanceTrendSeries) {
	o.Series = v
}

// GetSources returns the Sources field value.
func (o *PerformanceTrends) GetSources() []PerformanceTrendSource {
	if o == nil {
		var ret []PerformanceTrendSource
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value
// and a boolean to check if the value has been set.
func (o *PerformanceTrends) GetSourcesOk() (*[]PerformanceTrendSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sources, true
}

// SetSources sets field value.
func (o *PerformanceTrends) SetSources(v []PerformanceTrendSource) {
	o.Sources = v
}

// GetWarnings returns the Warnings field value.
func (o *PerformanceTrends) GetWarnings() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value
// and a boolean to check if the value has been set.
func (o *PerformanceTrends) GetWarningsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Warnings, true
}

// SetWarnings sets field value.
func (o *PerformanceTrends) SetWarnings(v []string) {
	o.Warnings = v
}

// GetCollectedAt returns the CollectedAt field value.
func (o *PerformanceTrends) GetCollectedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CollectedAt
}

// GetCollectedAtOk returns a tuple with the CollectedAt field value
// and a boolean to check if the value has been set.
func (o *PerformanceTrends) GetCollectedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectedAt, true
}

// SetCollectedAt sets field value.
func (o *PerformanceTrends) SetCollectedAt(v string) {
	o.CollectedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PerformanceTrends) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["range"] = o.Range
	toSerialize["granularity"] = o.Granularity
	toSerialize["summary"] = o.Summary
	toSerialize["series"] = o.Series
	toSerialize["sources"] = o.Sources
	toSerialize["warnings"] = o.Warnings
	toSerialize["collectedAt"] = o.CollectedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PerformanceTrends) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Range       *PerformanceTrendRange     `json:"range"`
		Granularity *string                    `json:"granularity"`
		Summary     *[]PerformanceTrendSummary `json:"summary"`
		Series      *[]PerformanceTrendSeries  `json:"series"`
		Sources     *[]PerformanceTrendSource  `json:"sources"`
		Warnings    *[]string                  `json:"warnings"`
		CollectedAt *string                    `json:"collectedAt"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Range == nil {
		return fmt.Errorf("required field range missing")
	}
	if all.Granularity == nil {
		return fmt.Errorf("required field granularity missing")
	}
	if all.Summary == nil {
		return fmt.Errorf("required field summary missing")
	}
	if all.Series == nil {
		return fmt.Errorf("required field series missing")
	}
	if all.Sources == nil {
		return fmt.Errorf("required field sources missing")
	}
	if all.Warnings == nil {
		return fmt.Errorf("required field warnings missing")
	}
	if all.CollectedAt == nil {
		return fmt.Errorf("required field collectedAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"range", "granularity", "summary", "series", "sources", "warnings", "collectedAt"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Range.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Range = *all.Range
	o.Granularity = *all.Granularity
	o.Summary = *all.Summary
	o.Series = *all.Series
	o.Sources = *all.Sources
	o.Warnings = *all.Warnings
	o.CollectedAt = *all.CollectedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
