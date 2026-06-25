// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PerformanceTrendSeries struct {
	Metric      string               `json:"metric"`
	DisplayName LocalizedDescription `json:"displayName"`
	// Metric category. Values are connectionsSessions, workload, resourcePressure, storagePVC, availability, maintenance, performance, or capacity.
	Category string `json:"category"`
	// Optional series dimension. Values include cluster, instance, pvc, and database.
	Dimension *string `json:"dimension,omitempty"`
	// Optional dimension labels, for example instance, pvc, role, or datname.
	Labels map[string]string `json:"labels,omitempty"`
	Unit   string            `json:"unit"`
	// Optional warning threshold in the same unit as this metric.
	WarnThreshold *float64 `json:"warnThreshold,omitempty"`
	// Optional critical threshold in the same unit as this metric.
	CritThreshold *float64 `json:"critThreshold,omitempty"`
	// Optional threshold direction. Values are above, below, or boolean.
	ThresholdDirection *string                 `json:"thresholdDirection,omitempty"`
	Points             []PerformanceTrendPoint `json:"points"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPerformanceTrendSeries instantiates a new PerformanceTrendSeries object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPerformanceTrendSeries(metric string, displayName LocalizedDescription, category string, unit string, points []PerformanceTrendPoint) *PerformanceTrendSeries {
	this := PerformanceTrendSeries{}
	this.Metric = metric
	this.DisplayName = displayName
	this.Category = category
	this.Unit = unit
	this.Points = points
	return &this
}

// NewPerformanceTrendSeriesWithDefaults instantiates a new PerformanceTrendSeries object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPerformanceTrendSeriesWithDefaults() *PerformanceTrendSeries {
	this := PerformanceTrendSeries{}
	return &this
}

// GetMetric returns the Metric field value.
func (o *PerformanceTrendSeries) GetMetric() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Metric
}

// GetMetricOk returns a tuple with the Metric field value
// and a boolean to check if the value has been set.
func (o *PerformanceTrendSeries) GetMetricOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metric, true
}

// SetMetric sets field value.
func (o *PerformanceTrendSeries) SetMetric(v string) {
	o.Metric = v
}

// GetDisplayName returns the DisplayName field value.
func (o *PerformanceTrendSeries) GetDisplayName() LocalizedDescription {
	if o == nil {
		var ret LocalizedDescription
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *PerformanceTrendSeries) GetDisplayNameOk() (*LocalizedDescription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value.
func (o *PerformanceTrendSeries) SetDisplayName(v LocalizedDescription) {
	o.DisplayName = v
}

// GetCategory returns the Category field value.
func (o *PerformanceTrendSeries) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *PerformanceTrendSeries) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value.
func (o *PerformanceTrendSeries) SetCategory(v string) {
	o.Category = v
}

// GetDimension returns the Dimension field value if set, zero value otherwise.
func (o *PerformanceTrendSeries) GetDimension() string {
	if o == nil || o.Dimension == nil {
		var ret string
		return ret
	}
	return *o.Dimension
}

// GetDimensionOk returns a tuple with the Dimension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceTrendSeries) GetDimensionOk() (*string, bool) {
	if o == nil || o.Dimension == nil {
		return nil, false
	}
	return o.Dimension, true
}

// HasDimension returns a boolean if a field has been set.
func (o *PerformanceTrendSeries) HasDimension() bool {
	return o != nil && o.Dimension != nil
}

// SetDimension gets a reference to the given string and assigns it to the Dimension field.
func (o *PerformanceTrendSeries) SetDimension(v string) {
	o.Dimension = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *PerformanceTrendSeries) GetLabels() map[string]string {
	if o == nil || o.Labels == nil {
		var ret map[string]string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceTrendSeries) GetLabelsOk() (*map[string]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return &o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *PerformanceTrendSeries) HasLabels() bool {
	return o != nil && o.Labels != nil
}

// SetLabels gets a reference to the given map[string]string and assigns it to the Labels field.
func (o *PerformanceTrendSeries) SetLabels(v map[string]string) {
	o.Labels = v
}

// GetUnit returns the Unit field value.
func (o *PerformanceTrendSeries) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *PerformanceTrendSeries) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value.
func (o *PerformanceTrendSeries) SetUnit(v string) {
	o.Unit = v
}

// GetWarnThreshold returns the WarnThreshold field value if set, zero value otherwise.
func (o *PerformanceTrendSeries) GetWarnThreshold() float64 {
	if o == nil || o.WarnThreshold == nil {
		var ret float64
		return ret
	}
	return *o.WarnThreshold
}

// GetWarnThresholdOk returns a tuple with the WarnThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceTrendSeries) GetWarnThresholdOk() (*float64, bool) {
	if o == nil || o.WarnThreshold == nil {
		return nil, false
	}
	return o.WarnThreshold, true
}

// HasWarnThreshold returns a boolean if a field has been set.
func (o *PerformanceTrendSeries) HasWarnThreshold() bool {
	return o != nil && o.WarnThreshold != nil
}

// SetWarnThreshold gets a reference to the given float64 and assigns it to the WarnThreshold field.
func (o *PerformanceTrendSeries) SetWarnThreshold(v float64) {
	o.WarnThreshold = &v
}

// GetCritThreshold returns the CritThreshold field value if set, zero value otherwise.
func (o *PerformanceTrendSeries) GetCritThreshold() float64 {
	if o == nil || o.CritThreshold == nil {
		var ret float64
		return ret
	}
	return *o.CritThreshold
}

// GetCritThresholdOk returns a tuple with the CritThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceTrendSeries) GetCritThresholdOk() (*float64, bool) {
	if o == nil || o.CritThreshold == nil {
		return nil, false
	}
	return o.CritThreshold, true
}

// HasCritThreshold returns a boolean if a field has been set.
func (o *PerformanceTrendSeries) HasCritThreshold() bool {
	return o != nil && o.CritThreshold != nil
}

// SetCritThreshold gets a reference to the given float64 and assigns it to the CritThreshold field.
func (o *PerformanceTrendSeries) SetCritThreshold(v float64) {
	o.CritThreshold = &v
}

// GetThresholdDirection returns the ThresholdDirection field value if set, zero value otherwise.
func (o *PerformanceTrendSeries) GetThresholdDirection() string {
	if o == nil || o.ThresholdDirection == nil {
		var ret string
		return ret
	}
	return *o.ThresholdDirection
}

// GetThresholdDirectionOk returns a tuple with the ThresholdDirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceTrendSeries) GetThresholdDirectionOk() (*string, bool) {
	if o == nil || o.ThresholdDirection == nil {
		return nil, false
	}
	return o.ThresholdDirection, true
}

// HasThresholdDirection returns a boolean if a field has been set.
func (o *PerformanceTrendSeries) HasThresholdDirection() bool {
	return o != nil && o.ThresholdDirection != nil
}

// SetThresholdDirection gets a reference to the given string and assigns it to the ThresholdDirection field.
func (o *PerformanceTrendSeries) SetThresholdDirection(v string) {
	o.ThresholdDirection = &v
}

// GetPoints returns the Points field value.
func (o *PerformanceTrendSeries) GetPoints() []PerformanceTrendPoint {
	if o == nil {
		var ret []PerformanceTrendPoint
		return ret
	}
	return o.Points
}

// GetPointsOk returns a tuple with the Points field value
// and a boolean to check if the value has been set.
func (o *PerformanceTrendSeries) GetPointsOk() (*[]PerformanceTrendPoint, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Points, true
}

// SetPoints sets field value.
func (o *PerformanceTrendSeries) SetPoints(v []PerformanceTrendPoint) {
	o.Points = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PerformanceTrendSeries) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["metric"] = o.Metric
	toSerialize["displayName"] = o.DisplayName
	toSerialize["category"] = o.Category
	if o.Dimension != nil {
		toSerialize["dimension"] = o.Dimension
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	toSerialize["unit"] = o.Unit
	if o.WarnThreshold != nil {
		toSerialize["warnThreshold"] = o.WarnThreshold
	}
	if o.CritThreshold != nil {
		toSerialize["critThreshold"] = o.CritThreshold
	}
	if o.ThresholdDirection != nil {
		toSerialize["thresholdDirection"] = o.ThresholdDirection
	}
	toSerialize["points"] = o.Points

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PerformanceTrendSeries) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Metric             *string                  `json:"metric"`
		DisplayName        *LocalizedDescription    `json:"displayName"`
		Category           *string                  `json:"category"`
		Dimension          *string                  `json:"dimension,omitempty"`
		Labels             map[string]string        `json:"labels,omitempty"`
		Unit               *string                  `json:"unit"`
		WarnThreshold      *float64                 `json:"warnThreshold,omitempty"`
		CritThreshold      *float64                 `json:"critThreshold,omitempty"`
		ThresholdDirection *string                  `json:"thresholdDirection,omitempty"`
		Points             *[]PerformanceTrendPoint `json:"points"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Metric == nil {
		return fmt.Errorf("required field metric missing")
	}
	if all.DisplayName == nil {
		return fmt.Errorf("required field displayName missing")
	}
	if all.Category == nil {
		return fmt.Errorf("required field category missing")
	}
	if all.Unit == nil {
		return fmt.Errorf("required field unit missing")
	}
	if all.Points == nil {
		return fmt.Errorf("required field points missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"metric", "displayName", "category", "dimension", "labels", "unit", "warnThreshold", "critThreshold", "thresholdDirection", "points"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Metric = *all.Metric
	if all.DisplayName.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DisplayName = *all.DisplayName
	o.Category = *all.Category
	o.Dimension = all.Dimension
	o.Labels = all.Labels
	o.Unit = *all.Unit
	o.WarnThreshold = all.WarnThreshold
	o.CritThreshold = all.CritThreshold
	o.ThresholdDirection = all.ThresholdDirection
	o.Points = *all.Points

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
