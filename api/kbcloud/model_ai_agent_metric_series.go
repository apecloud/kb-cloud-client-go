// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentMetricSeries struct {
	MetricSeriesId    string                     `json:"metricSeriesId"`
	Title             string                     `json:"title"`
	Unit              *string                    `json:"unit,omitempty"`
	TimeRange         *AiAgentTimeRange          `json:"timeRange,omitempty"`
	Points            []AiAgentMetricSeriesPoint `json:"points"`
	SourceEvidenceIds []string                   `json:"sourceEvidenceIds"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentMetricSeries instantiates a new AiAgentMetricSeries object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentMetricSeries(metricSeriesId string, title string, points []AiAgentMetricSeriesPoint, sourceEvidenceIds []string) *AiAgentMetricSeries {
	this := AiAgentMetricSeries{}
	this.MetricSeriesId = metricSeriesId
	this.Title = title
	this.Points = points
	this.SourceEvidenceIds = sourceEvidenceIds
	return &this
}

// NewAiAgentMetricSeriesWithDefaults instantiates a new AiAgentMetricSeries object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentMetricSeriesWithDefaults() *AiAgentMetricSeries {
	this := AiAgentMetricSeries{}
	return &this
}

// GetMetricSeriesId returns the MetricSeriesId field value.
func (o *AiAgentMetricSeries) GetMetricSeriesId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MetricSeriesId
}

// GetMetricSeriesIdOk returns a tuple with the MetricSeriesId field value
// and a boolean to check if the value has been set.
func (o *AiAgentMetricSeries) GetMetricSeriesIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricSeriesId, true
}

// SetMetricSeriesId sets field value.
func (o *AiAgentMetricSeries) SetMetricSeriesId(v string) {
	o.MetricSeriesId = v
}

// GetTitle returns the Title field value.
func (o *AiAgentMetricSeries) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *AiAgentMetricSeries) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *AiAgentMetricSeries) SetTitle(v string) {
	o.Title = v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *AiAgentMetricSeries) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentMetricSeries) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *AiAgentMetricSeries) HasUnit() bool {
	return o != nil && o.Unit != nil
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *AiAgentMetricSeries) SetUnit(v string) {
	o.Unit = &v
}

// GetTimeRange returns the TimeRange field value if set, zero value otherwise.
func (o *AiAgentMetricSeries) GetTimeRange() AiAgentTimeRange {
	if o == nil || o.TimeRange == nil {
		var ret AiAgentTimeRange
		return ret
	}
	return *o.TimeRange
}

// GetTimeRangeOk returns a tuple with the TimeRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentMetricSeries) GetTimeRangeOk() (*AiAgentTimeRange, bool) {
	if o == nil || o.TimeRange == nil {
		return nil, false
	}
	return o.TimeRange, true
}

// HasTimeRange returns a boolean if a field has been set.
func (o *AiAgentMetricSeries) HasTimeRange() bool {
	return o != nil && o.TimeRange != nil
}

// SetTimeRange gets a reference to the given AiAgentTimeRange and assigns it to the TimeRange field.
func (o *AiAgentMetricSeries) SetTimeRange(v AiAgentTimeRange) {
	o.TimeRange = &v
}

// GetPoints returns the Points field value.
func (o *AiAgentMetricSeries) GetPoints() []AiAgentMetricSeriesPoint {
	if o == nil {
		var ret []AiAgentMetricSeriesPoint
		return ret
	}
	return o.Points
}

// GetPointsOk returns a tuple with the Points field value
// and a boolean to check if the value has been set.
func (o *AiAgentMetricSeries) GetPointsOk() (*[]AiAgentMetricSeriesPoint, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Points, true
}

// SetPoints sets field value.
func (o *AiAgentMetricSeries) SetPoints(v []AiAgentMetricSeriesPoint) {
	o.Points = v
}

// GetSourceEvidenceIds returns the SourceEvidenceIds field value.
func (o *AiAgentMetricSeries) GetSourceEvidenceIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SourceEvidenceIds
}

// GetSourceEvidenceIdsOk returns a tuple with the SourceEvidenceIds field value
// and a boolean to check if the value has been set.
func (o *AiAgentMetricSeries) GetSourceEvidenceIdsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceEvidenceIds, true
}

// SetSourceEvidenceIds sets field value.
func (o *AiAgentMetricSeries) SetSourceEvidenceIds(v []string) {
	o.SourceEvidenceIds = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentMetricSeries) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["metricSeriesId"] = o.MetricSeriesId
	toSerialize["title"] = o.Title
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}
	if o.TimeRange != nil {
		toSerialize["timeRange"] = o.TimeRange
	}
	toSerialize["points"] = o.Points
	toSerialize["sourceEvidenceIds"] = o.SourceEvidenceIds

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentMetricSeries) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MetricSeriesId    *string                     `json:"metricSeriesId"`
		Title             *string                     `json:"title"`
		Unit              *string                     `json:"unit,omitempty"`
		TimeRange         *AiAgentTimeRange           `json:"timeRange,omitempty"`
		Points            *[]AiAgentMetricSeriesPoint `json:"points"`
		SourceEvidenceIds *[]string                   `json:"sourceEvidenceIds"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.MetricSeriesId == nil {
		return fmt.Errorf("required field metricSeriesId missing")
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	if all.Points == nil {
		return fmt.Errorf("required field points missing")
	}
	if all.SourceEvidenceIds == nil {
		return fmt.Errorf("required field sourceEvidenceIds missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"metricSeriesId", "title", "unit", "timeRange", "points", "sourceEvidenceIds"})
	} else {
		return err
	}

	hasInvalidField := false
	o.MetricSeriesId = *all.MetricSeriesId
	o.Title = *all.Title
	o.Unit = all.Unit
	if all.TimeRange != nil && all.TimeRange.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TimeRange = all.TimeRange
	o.Points = *all.Points
	o.SourceEvidenceIds = *all.SourceEvidenceIds

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
