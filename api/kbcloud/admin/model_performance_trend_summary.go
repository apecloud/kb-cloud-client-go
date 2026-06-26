// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PerformanceTrendSummary struct {
	Metric      string               `json:"metric"`
	DisplayName LocalizedDescription `json:"displayName"`
	// Metric category. Values are availability, performance, connections, resource, capacity, or vacuumWraparound.
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
	ThresholdDirection *string `json:"thresholdDirection,omitempty"`
	Min                float64 `json:"min"`
	Max                float64 `json:"max"`
	Avg                float64 `json:"avg"`
	Latest             float64 `json:"latest"`
	// Change direction when at least two samples exist. Values are increasing, decreasing, or stable.
	ChangeDirection *string `json:"changeDirection,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPerformanceTrendSummary instantiates a new PerformanceTrendSummary object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPerformanceTrendSummary(metric string, displayName LocalizedDescription, category string, unit string, min float64, max float64, avg float64, latest float64) *PerformanceTrendSummary {
	this := PerformanceTrendSummary{}
	this.Metric = metric
	this.DisplayName = displayName
	this.Category = category
	this.Unit = unit
	this.Min = min
	this.Max = max
	this.Avg = avg
	this.Latest = latest
	return &this
}

// NewPerformanceTrendSummaryWithDefaults instantiates a new PerformanceTrendSummary object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPerformanceTrendSummaryWithDefaults() *PerformanceTrendSummary {
	this := PerformanceTrendSummary{}
	return &this
}

// GetMetric returns the Metric field value.
func (o *PerformanceTrendSummary) GetMetric() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Metric
}

// GetMetricOk returns a tuple with the Metric field value
// and a boolean to check if the value has been set.
func (o *PerformanceTrendSummary) GetMetricOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metric, true
}

// SetMetric sets field value.
func (o *PerformanceTrendSummary) SetMetric(v string) {
	o.Metric = v
}

// GetDisplayName returns the DisplayName field value.
func (o *PerformanceTrendSummary) GetDisplayName() LocalizedDescription {
	if o == nil {
		var ret LocalizedDescription
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *PerformanceTrendSummary) GetDisplayNameOk() (*LocalizedDescription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value.
func (o *PerformanceTrendSummary) SetDisplayName(v LocalizedDescription) {
	o.DisplayName = v
}

// GetCategory returns the Category field value.
func (o *PerformanceTrendSummary) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *PerformanceTrendSummary) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value.
func (o *PerformanceTrendSummary) SetCategory(v string) {
	o.Category = v
}

// GetDimension returns the Dimension field value if set, zero value otherwise.
func (o *PerformanceTrendSummary) GetDimension() string {
	if o == nil || o.Dimension == nil {
		var ret string
		return ret
	}
	return *o.Dimension
}

// GetDimensionOk returns a tuple with the Dimension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceTrendSummary) GetDimensionOk() (*string, bool) {
	if o == nil || o.Dimension == nil {
		return nil, false
	}
	return o.Dimension, true
}

// HasDimension returns a boolean if a field has been set.
func (o *PerformanceTrendSummary) HasDimension() bool {
	return o != nil && o.Dimension != nil
}

// SetDimension gets a reference to the given string and assigns it to the Dimension field.
func (o *PerformanceTrendSummary) SetDimension(v string) {
	o.Dimension = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *PerformanceTrendSummary) GetLabels() map[string]string {
	if o == nil || o.Labels == nil {
		var ret map[string]string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceTrendSummary) GetLabelsOk() (*map[string]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return &o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *PerformanceTrendSummary) HasLabels() bool {
	return o != nil && o.Labels != nil
}

// SetLabels gets a reference to the given map[string]string and assigns it to the Labels field.
func (o *PerformanceTrendSummary) SetLabels(v map[string]string) {
	o.Labels = v
}

// GetUnit returns the Unit field value.
func (o *PerformanceTrendSummary) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *PerformanceTrendSummary) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value.
func (o *PerformanceTrendSummary) SetUnit(v string) {
	o.Unit = v
}

// GetWarnThreshold returns the WarnThreshold field value if set, zero value otherwise.
func (o *PerformanceTrendSummary) GetWarnThreshold() float64 {
	if o == nil || o.WarnThreshold == nil {
		var ret float64
		return ret
	}
	return *o.WarnThreshold
}

// GetWarnThresholdOk returns a tuple with the WarnThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceTrendSummary) GetWarnThresholdOk() (*float64, bool) {
	if o == nil || o.WarnThreshold == nil {
		return nil, false
	}
	return o.WarnThreshold, true
}

// HasWarnThreshold returns a boolean if a field has been set.
func (o *PerformanceTrendSummary) HasWarnThreshold() bool {
	return o != nil && o.WarnThreshold != nil
}

// SetWarnThreshold gets a reference to the given float64 and assigns it to the WarnThreshold field.
func (o *PerformanceTrendSummary) SetWarnThreshold(v float64) {
	o.WarnThreshold = &v
}

// GetCritThreshold returns the CritThreshold field value if set, zero value otherwise.
func (o *PerformanceTrendSummary) GetCritThreshold() float64 {
	if o == nil || o.CritThreshold == nil {
		var ret float64
		return ret
	}
	return *o.CritThreshold
}

// GetCritThresholdOk returns a tuple with the CritThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceTrendSummary) GetCritThresholdOk() (*float64, bool) {
	if o == nil || o.CritThreshold == nil {
		return nil, false
	}
	return o.CritThreshold, true
}

// HasCritThreshold returns a boolean if a field has been set.
func (o *PerformanceTrendSummary) HasCritThreshold() bool {
	return o != nil && o.CritThreshold != nil
}

// SetCritThreshold gets a reference to the given float64 and assigns it to the CritThreshold field.
func (o *PerformanceTrendSummary) SetCritThreshold(v float64) {
	o.CritThreshold = &v
}

// GetThresholdDirection returns the ThresholdDirection field value if set, zero value otherwise.
func (o *PerformanceTrendSummary) GetThresholdDirection() string {
	if o == nil || o.ThresholdDirection == nil {
		var ret string
		return ret
	}
	return *o.ThresholdDirection
}

// GetThresholdDirectionOk returns a tuple with the ThresholdDirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceTrendSummary) GetThresholdDirectionOk() (*string, bool) {
	if o == nil || o.ThresholdDirection == nil {
		return nil, false
	}
	return o.ThresholdDirection, true
}

// HasThresholdDirection returns a boolean if a field has been set.
func (o *PerformanceTrendSummary) HasThresholdDirection() bool {
	return o != nil && o.ThresholdDirection != nil
}

// SetThresholdDirection gets a reference to the given string and assigns it to the ThresholdDirection field.
func (o *PerformanceTrendSummary) SetThresholdDirection(v string) {
	o.ThresholdDirection = &v
}

// GetMin returns the Min field value.
func (o *PerformanceTrendSummary) GetMin() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Min
}

// GetMinOk returns a tuple with the Min field value
// and a boolean to check if the value has been set.
func (o *PerformanceTrendSummary) GetMinOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Min, true
}

// SetMin sets field value.
func (o *PerformanceTrendSummary) SetMin(v float64) {
	o.Min = v
}

// GetMax returns the Max field value.
func (o *PerformanceTrendSummary) GetMax() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Max
}

// GetMaxOk returns a tuple with the Max field value
// and a boolean to check if the value has been set.
func (o *PerformanceTrendSummary) GetMaxOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Max, true
}

// SetMax sets field value.
func (o *PerformanceTrendSummary) SetMax(v float64) {
	o.Max = v
}

// GetAvg returns the Avg field value.
func (o *PerformanceTrendSummary) GetAvg() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Avg
}

// GetAvgOk returns a tuple with the Avg field value
// and a boolean to check if the value has been set.
func (o *PerformanceTrendSummary) GetAvgOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Avg, true
}

// SetAvg sets field value.
func (o *PerformanceTrendSummary) SetAvg(v float64) {
	o.Avg = v
}

// GetLatest returns the Latest field value.
func (o *PerformanceTrendSummary) GetLatest() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Latest
}

// GetLatestOk returns a tuple with the Latest field value
// and a boolean to check if the value has been set.
func (o *PerformanceTrendSummary) GetLatestOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Latest, true
}

// SetLatest sets field value.
func (o *PerformanceTrendSummary) SetLatest(v float64) {
	o.Latest = v
}

// GetChangeDirection returns the ChangeDirection field value if set, zero value otherwise.
func (o *PerformanceTrendSummary) GetChangeDirection() string {
	if o == nil || o.ChangeDirection == nil {
		var ret string
		return ret
	}
	return *o.ChangeDirection
}

// GetChangeDirectionOk returns a tuple with the ChangeDirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceTrendSummary) GetChangeDirectionOk() (*string, bool) {
	if o == nil || o.ChangeDirection == nil {
		return nil, false
	}
	return o.ChangeDirection, true
}

// HasChangeDirection returns a boolean if a field has been set.
func (o *PerformanceTrendSummary) HasChangeDirection() bool {
	return o != nil && o.ChangeDirection != nil
}

// SetChangeDirection gets a reference to the given string and assigns it to the ChangeDirection field.
func (o *PerformanceTrendSummary) SetChangeDirection(v string) {
	o.ChangeDirection = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PerformanceTrendSummary) MarshalJSON() ([]byte, error) {
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
	toSerialize["min"] = o.Min
	toSerialize["max"] = o.Max
	toSerialize["avg"] = o.Avg
	toSerialize["latest"] = o.Latest
	if o.ChangeDirection != nil {
		toSerialize["changeDirection"] = o.ChangeDirection
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PerformanceTrendSummary) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Metric             *string               `json:"metric"`
		DisplayName        *LocalizedDescription `json:"displayName"`
		Category           *string               `json:"category"`
		Dimension          *string               `json:"dimension,omitempty"`
		Labels             map[string]string     `json:"labels,omitempty"`
		Unit               *string               `json:"unit"`
		WarnThreshold      *float64              `json:"warnThreshold,omitempty"`
		CritThreshold      *float64              `json:"critThreshold,omitempty"`
		ThresholdDirection *string               `json:"thresholdDirection,omitempty"`
		Min                *float64              `json:"min"`
		Max                *float64              `json:"max"`
		Avg                *float64              `json:"avg"`
		Latest             *float64              `json:"latest"`
		ChangeDirection    *string               `json:"changeDirection,omitempty"`
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
	if all.Min == nil {
		return fmt.Errorf("required field min missing")
	}
	if all.Max == nil {
		return fmt.Errorf("required field max missing")
	}
	if all.Avg == nil {
		return fmt.Errorf("required field avg missing")
	}
	if all.Latest == nil {
		return fmt.Errorf("required field latest missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"metric", "displayName", "category", "dimension", "labels", "unit", "warnThreshold", "critThreshold", "thresholdDirection", "min", "max", "avg", "latest", "changeDirection"})
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
	o.Min = *all.Min
	o.Max = *all.Max
	o.Avg = *all.Avg
	o.Latest = *all.Latest
	o.ChangeDirection = all.ChangeDirection

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
