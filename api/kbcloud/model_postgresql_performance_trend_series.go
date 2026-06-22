// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlPerformanceTrendSeries struct {
	Metric      string `json:"metric"`
	DisplayName string `json:"displayName"`
	// Metric category. Values are connectionsSessions, workload, resourcePressure, or storagePVC.
	Category string                            `json:"category"`
	Unit     string                            `json:"unit"`
	Points   []PostgresqlPerformanceTrendPoint `json:"points"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlPerformanceTrendSeries instantiates a new PostgresqlPerformanceTrendSeries object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlPerformanceTrendSeries(metric string, displayName string, category string, unit string, points []PostgresqlPerformanceTrendPoint) *PostgresqlPerformanceTrendSeries {
	this := PostgresqlPerformanceTrendSeries{}
	this.Metric = metric
	this.DisplayName = displayName
	this.Category = category
	this.Unit = unit
	this.Points = points
	return &this
}

// NewPostgresqlPerformanceTrendSeriesWithDefaults instantiates a new PostgresqlPerformanceTrendSeries object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlPerformanceTrendSeriesWithDefaults() *PostgresqlPerformanceTrendSeries {
	this := PostgresqlPerformanceTrendSeries{}
	return &this
}

// GetMetric returns the Metric field value.
func (o *PostgresqlPerformanceTrendSeries) GetMetric() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Metric
}

// GetMetricOk returns a tuple with the Metric field value
// and a boolean to check if the value has been set.
func (o *PostgresqlPerformanceTrendSeries) GetMetricOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metric, true
}

// SetMetric sets field value.
func (o *PostgresqlPerformanceTrendSeries) SetMetric(v string) {
	o.Metric = v
}

// GetDisplayName returns the DisplayName field value.
func (o *PostgresqlPerformanceTrendSeries) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *PostgresqlPerformanceTrendSeries) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value.
func (o *PostgresqlPerformanceTrendSeries) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetCategory returns the Category field value.
func (o *PostgresqlPerformanceTrendSeries) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *PostgresqlPerformanceTrendSeries) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value.
func (o *PostgresqlPerformanceTrendSeries) SetCategory(v string) {
	o.Category = v
}

// GetUnit returns the Unit field value.
func (o *PostgresqlPerformanceTrendSeries) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *PostgresqlPerformanceTrendSeries) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value.
func (o *PostgresqlPerformanceTrendSeries) SetUnit(v string) {
	o.Unit = v
}

// GetPoints returns the Points field value.
func (o *PostgresqlPerformanceTrendSeries) GetPoints() []PostgresqlPerformanceTrendPoint {
	if o == nil {
		var ret []PostgresqlPerformanceTrendPoint
		return ret
	}
	return o.Points
}

// GetPointsOk returns a tuple with the Points field value
// and a boolean to check if the value has been set.
func (o *PostgresqlPerformanceTrendSeries) GetPointsOk() (*[]PostgresqlPerformanceTrendPoint, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Points, true
}

// SetPoints sets field value.
func (o *PostgresqlPerformanceTrendSeries) SetPoints(v []PostgresqlPerformanceTrendPoint) {
	o.Points = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlPerformanceTrendSeries) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["metric"] = o.Metric
	toSerialize["displayName"] = o.DisplayName
	toSerialize["category"] = o.Category
	toSerialize["unit"] = o.Unit
	toSerialize["points"] = o.Points

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostgresqlPerformanceTrendSeries) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Metric      *string                            `json:"metric"`
		DisplayName *string                            `json:"displayName"`
		Category    *string                            `json:"category"`
		Unit        *string                            `json:"unit"`
		Points      *[]PostgresqlPerformanceTrendPoint `json:"points"`
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
		common.DeleteKeys(additionalProperties, &[]string{"metric", "displayName", "category", "unit", "points"})
	} else {
		return err
	}
	o.Metric = *all.Metric
	o.DisplayName = *all.DisplayName
	o.Category = *all.Category
	o.Unit = *all.Unit
	o.Points = *all.Points

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
