// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlPerformanceTrends struct {
	Range PostgresqlPerformanceTrendRange `json:"range"`
	// Controlled backend query step.
	Granularity string `json:"granularity"`
	// Per-series numeric summary. changeDirection is present only when at least two points exist.
	Summary []PostgresqlPerformanceTrendSummary `json:"summary"`
	// Successfully collected and displayable metric series only. Unsupported, empty, or failed metrics are not represented as per-metric entries.
	Series []PostgresqlPerformanceTrendSeries `json:"series"`
	// Coarse collection source status. Messages are sanitized and do not expose PromQL, endpoints, credentials, or raw internal errors.
	Sources  []PostgresqlPerformanceTrendSource `json:"sources"`
	Warnings []string                           `json:"warnings"`
	// Backend collection timestamp in UTC.
	CollectedAt string `json:"collectedAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlPerformanceTrends instantiates a new PostgresqlPerformanceTrends object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlPerformanceTrends(rangeVar PostgresqlPerformanceTrendRange, granularity string, summary []PostgresqlPerformanceTrendSummary, series []PostgresqlPerformanceTrendSeries, sources []PostgresqlPerformanceTrendSource, warnings []string, collectedAt string) *PostgresqlPerformanceTrends {
	this := PostgresqlPerformanceTrends{}
	this.Range = rangeVar
	this.Granularity = granularity
	this.Summary = summary
	this.Series = series
	this.Sources = sources
	this.Warnings = warnings
	this.CollectedAt = collectedAt
	return &this
}

// NewPostgresqlPerformanceTrendsWithDefaults instantiates a new PostgresqlPerformanceTrends object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlPerformanceTrendsWithDefaults() *PostgresqlPerformanceTrends {
	this := PostgresqlPerformanceTrends{}
	return &this
}

// GetRange returns the Range field value.
func (o *PostgresqlPerformanceTrends) GetRange() PostgresqlPerformanceTrendRange {
	if o == nil {
		var ret PostgresqlPerformanceTrendRange
		return ret
	}
	return o.Range
}

// GetRangeOk returns a tuple with the Range field value
// and a boolean to check if the value has been set.
func (o *PostgresqlPerformanceTrends) GetRangeOk() (*PostgresqlPerformanceTrendRange, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Range, true
}

// SetRange sets field value.
func (o *PostgresqlPerformanceTrends) SetRange(v PostgresqlPerformanceTrendRange) {
	o.Range = v
}

// GetGranularity returns the Granularity field value.
func (o *PostgresqlPerformanceTrends) GetGranularity() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Granularity
}

// GetGranularityOk returns a tuple with the Granularity field value
// and a boolean to check if the value has been set.
func (o *PostgresqlPerformanceTrends) GetGranularityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Granularity, true
}

// SetGranularity sets field value.
func (o *PostgresqlPerformanceTrends) SetGranularity(v string) {
	o.Granularity = v
}

// GetSummary returns the Summary field value.
func (o *PostgresqlPerformanceTrends) GetSummary() []PostgresqlPerformanceTrendSummary {
	if o == nil {
		var ret []PostgresqlPerformanceTrendSummary
		return ret
	}
	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *PostgresqlPerformanceTrends) GetSummaryOk() (*[]PostgresqlPerformanceTrendSummary, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value.
func (o *PostgresqlPerformanceTrends) SetSummary(v []PostgresqlPerformanceTrendSummary) {
	o.Summary = v
}

// GetSeries returns the Series field value.
func (o *PostgresqlPerformanceTrends) GetSeries() []PostgresqlPerformanceTrendSeries {
	if o == nil {
		var ret []PostgresqlPerformanceTrendSeries
		return ret
	}
	return o.Series
}

// GetSeriesOk returns a tuple with the Series field value
// and a boolean to check if the value has been set.
func (o *PostgresqlPerformanceTrends) GetSeriesOk() (*[]PostgresqlPerformanceTrendSeries, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Series, true
}

// SetSeries sets field value.
func (o *PostgresqlPerformanceTrends) SetSeries(v []PostgresqlPerformanceTrendSeries) {
	o.Series = v
}

// GetSources returns the Sources field value.
func (o *PostgresqlPerformanceTrends) GetSources() []PostgresqlPerformanceTrendSource {
	if o == nil {
		var ret []PostgresqlPerformanceTrendSource
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value
// and a boolean to check if the value has been set.
func (o *PostgresqlPerformanceTrends) GetSourcesOk() (*[]PostgresqlPerformanceTrendSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sources, true
}

// SetSources sets field value.
func (o *PostgresqlPerformanceTrends) SetSources(v []PostgresqlPerformanceTrendSource) {
	o.Sources = v
}

// GetWarnings returns the Warnings field value.
func (o *PostgresqlPerformanceTrends) GetWarnings() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value
// and a boolean to check if the value has been set.
func (o *PostgresqlPerformanceTrends) GetWarningsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Warnings, true
}

// SetWarnings sets field value.
func (o *PostgresqlPerformanceTrends) SetWarnings(v []string) {
	o.Warnings = v
}

// GetCollectedAt returns the CollectedAt field value.
func (o *PostgresqlPerformanceTrends) GetCollectedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CollectedAt
}

// GetCollectedAtOk returns a tuple with the CollectedAt field value
// and a boolean to check if the value has been set.
func (o *PostgresqlPerformanceTrends) GetCollectedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectedAt, true
}

// SetCollectedAt sets field value.
func (o *PostgresqlPerformanceTrends) SetCollectedAt(v string) {
	o.CollectedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlPerformanceTrends) MarshalJSON() ([]byte, error) {
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
func (o *PostgresqlPerformanceTrends) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Range       *PostgresqlPerformanceTrendRange     `json:"range"`
		Granularity *string                              `json:"granularity"`
		Summary     *[]PostgresqlPerformanceTrendSummary `json:"summary"`
		Series      *[]PostgresqlPerformanceTrendSeries  `json:"series"`
		Sources     *[]PostgresqlPerformanceTrendSource  `json:"sources"`
		Warnings    *[]string                            `json:"warnings"`
		CollectedAt *string                              `json:"collectedAt"`
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
