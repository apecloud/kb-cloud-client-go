// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlPerformanceTrendSummary struct {
	Metric      string `json:"metric"`
	DisplayName string `json:"displayName"`
	// Metric category. Values are connectionsSessions, workload, resourcePressure, or storagePVC.
	Category string  `json:"category"`
	Unit     string  `json:"unit"`
	Min      float64 `json:"min"`
	Max      float64 `json:"max"`
	Avg      float64 `json:"avg"`
	Latest   float64 `json:"latest"`
	// Change direction when at least two samples exist. Values are increasing, decreasing, or stable.
	ChangeDirection *string `json:"changeDirection,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlPerformanceTrendSummary instantiates a new PostgresqlPerformanceTrendSummary object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlPerformanceTrendSummary(metric string, displayName string, category string, unit string, min float64, max float64, avg float64, latest float64) *PostgresqlPerformanceTrendSummary {
	this := PostgresqlPerformanceTrendSummary{}
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

// NewPostgresqlPerformanceTrendSummaryWithDefaults instantiates a new PostgresqlPerformanceTrendSummary object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlPerformanceTrendSummaryWithDefaults() *PostgresqlPerformanceTrendSummary {
	this := PostgresqlPerformanceTrendSummary{}
	return &this
}

// GetMetric returns the Metric field value.
func (o *PostgresqlPerformanceTrendSummary) GetMetric() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Metric
}

// GetMetricOk returns a tuple with the Metric field value
// and a boolean to check if the value has been set.
func (o *PostgresqlPerformanceTrendSummary) GetMetricOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metric, true
}

// SetMetric sets field value.
func (o *PostgresqlPerformanceTrendSummary) SetMetric(v string) {
	o.Metric = v
}

// GetDisplayName returns the DisplayName field value.
func (o *PostgresqlPerformanceTrendSummary) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *PostgresqlPerformanceTrendSummary) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value.
func (o *PostgresqlPerformanceTrendSummary) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetCategory returns the Category field value.
func (o *PostgresqlPerformanceTrendSummary) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *PostgresqlPerformanceTrendSummary) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value.
func (o *PostgresqlPerformanceTrendSummary) SetCategory(v string) {
	o.Category = v
}

// GetUnit returns the Unit field value.
func (o *PostgresqlPerformanceTrendSummary) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *PostgresqlPerformanceTrendSummary) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value.
func (o *PostgresqlPerformanceTrendSummary) SetUnit(v string) {
	o.Unit = v
}

// GetMin returns the Min field value.
func (o *PostgresqlPerformanceTrendSummary) GetMin() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Min
}

// GetMinOk returns a tuple with the Min field value
// and a boolean to check if the value has been set.
func (o *PostgresqlPerformanceTrendSummary) GetMinOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Min, true
}

// SetMin sets field value.
func (o *PostgresqlPerformanceTrendSummary) SetMin(v float64) {
	o.Min = v
}

// GetMax returns the Max field value.
func (o *PostgresqlPerformanceTrendSummary) GetMax() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Max
}

// GetMaxOk returns a tuple with the Max field value
// and a boolean to check if the value has been set.
func (o *PostgresqlPerformanceTrendSummary) GetMaxOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Max, true
}

// SetMax sets field value.
func (o *PostgresqlPerformanceTrendSummary) SetMax(v float64) {
	o.Max = v
}

// GetAvg returns the Avg field value.
func (o *PostgresqlPerformanceTrendSummary) GetAvg() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Avg
}

// GetAvgOk returns a tuple with the Avg field value
// and a boolean to check if the value has been set.
func (o *PostgresqlPerformanceTrendSummary) GetAvgOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Avg, true
}

// SetAvg sets field value.
func (o *PostgresqlPerformanceTrendSummary) SetAvg(v float64) {
	o.Avg = v
}

// GetLatest returns the Latest field value.
func (o *PostgresqlPerformanceTrendSummary) GetLatest() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Latest
}

// GetLatestOk returns a tuple with the Latest field value
// and a boolean to check if the value has been set.
func (o *PostgresqlPerformanceTrendSummary) GetLatestOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Latest, true
}

// SetLatest sets field value.
func (o *PostgresqlPerformanceTrendSummary) SetLatest(v float64) {
	o.Latest = v
}

// GetChangeDirection returns the ChangeDirection field value if set, zero value otherwise.
func (o *PostgresqlPerformanceTrendSummary) GetChangeDirection() string {
	if o == nil || o.ChangeDirection == nil {
		var ret string
		return ret
	}
	return *o.ChangeDirection
}

// GetChangeDirectionOk returns a tuple with the ChangeDirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlPerformanceTrendSummary) GetChangeDirectionOk() (*string, bool) {
	if o == nil || o.ChangeDirection == nil {
		return nil, false
	}
	return o.ChangeDirection, true
}

// HasChangeDirection returns a boolean if a field has been set.
func (o *PostgresqlPerformanceTrendSummary) HasChangeDirection() bool {
	return o != nil && o.ChangeDirection != nil
}

// SetChangeDirection gets a reference to the given string and assigns it to the ChangeDirection field.
func (o *PostgresqlPerformanceTrendSummary) SetChangeDirection(v string) {
	o.ChangeDirection = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlPerformanceTrendSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["metric"] = o.Metric
	toSerialize["displayName"] = o.DisplayName
	toSerialize["category"] = o.Category
	toSerialize["unit"] = o.Unit
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
func (o *PostgresqlPerformanceTrendSummary) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Metric          *string  `json:"metric"`
		DisplayName     *string  `json:"displayName"`
		Category        *string  `json:"category"`
		Unit            *string  `json:"unit"`
		Min             *float64 `json:"min"`
		Max             *float64 `json:"max"`
		Avg             *float64 `json:"avg"`
		Latest          *float64 `json:"latest"`
		ChangeDirection *string  `json:"changeDirection,omitempty"`
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
		common.DeleteKeys(additionalProperties, &[]string{"metric", "displayName", "category", "unit", "min", "max", "avg", "latest", "changeDirection"})
	} else {
		return err
	}
	o.Metric = *all.Metric
	o.DisplayName = *all.DisplayName
	o.Category = *all.Category
	o.Unit = *all.Unit
	o.Min = *all.Min
	o.Max = *all.Max
	o.Avg = *all.Avg
	o.Latest = *all.Latest
	o.ChangeDirection = all.ChangeDirection

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
