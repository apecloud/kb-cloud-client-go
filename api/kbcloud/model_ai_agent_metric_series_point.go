// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// AiAgentMetricSeriesPoint Chart-ready point. x/y are the stable frontend plotting fields; timestamp/value may mirror them for time-series and numeric readers.
type AiAgentMetricSeriesPoint struct {
	Timestamp *time.Time `json:"timestamp,omitempty"`
	X         string     `json:"x"`
	Value     *float64   `json:"value,omitempty"`
	Y         float64    `json:"y"`
	Series    *string    `json:"series,omitempty"`
	Category  *string    `json:"category,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentMetricSeriesPoint instantiates a new AiAgentMetricSeriesPoint object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentMetricSeriesPoint(x string, y float64) *AiAgentMetricSeriesPoint {
	this := AiAgentMetricSeriesPoint{}
	this.X = x
	this.Y = y
	return &this
}

// NewAiAgentMetricSeriesPointWithDefaults instantiates a new AiAgentMetricSeriesPoint object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentMetricSeriesPointWithDefaults() *AiAgentMetricSeriesPoint {
	this := AiAgentMetricSeriesPoint{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *AiAgentMetricSeriesPoint) GetTimestamp() time.Time {
	if o == nil || o.Timestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentMetricSeriesPoint) GetTimestampOk() (*time.Time, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *AiAgentMetricSeriesPoint) HasTimestamp() bool {
	return o != nil && o.Timestamp != nil
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *AiAgentMetricSeriesPoint) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetX returns the X field value.
func (o *AiAgentMetricSeriesPoint) GetX() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.X
}

// GetXOk returns a tuple with the X field value
// and a boolean to check if the value has been set.
func (o *AiAgentMetricSeriesPoint) GetXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.X, true
}

// SetX sets field value.
func (o *AiAgentMetricSeriesPoint) SetX(v string) {
	o.X = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *AiAgentMetricSeriesPoint) GetValue() float64 {
	if o == nil || o.Value == nil {
		var ret float64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentMetricSeriesPoint) GetValueOk() (*float64, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *AiAgentMetricSeriesPoint) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given float64 and assigns it to the Value field.
func (o *AiAgentMetricSeriesPoint) SetValue(v float64) {
	o.Value = &v
}

// GetY returns the Y field value.
func (o *AiAgentMetricSeriesPoint) GetY() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Y
}

// GetYOk returns a tuple with the Y field value
// and a boolean to check if the value has been set.
func (o *AiAgentMetricSeriesPoint) GetYOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Y, true
}

// SetY sets field value.
func (o *AiAgentMetricSeriesPoint) SetY(v float64) {
	o.Y = v
}

// GetSeries returns the Series field value if set, zero value otherwise.
func (o *AiAgentMetricSeriesPoint) GetSeries() string {
	if o == nil || o.Series == nil {
		var ret string
		return ret
	}
	return *o.Series
}

// GetSeriesOk returns a tuple with the Series field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentMetricSeriesPoint) GetSeriesOk() (*string, bool) {
	if o == nil || o.Series == nil {
		return nil, false
	}
	return o.Series, true
}

// HasSeries returns a boolean if a field has been set.
func (o *AiAgentMetricSeriesPoint) HasSeries() bool {
	return o != nil && o.Series != nil
}

// SetSeries gets a reference to the given string and assigns it to the Series field.
func (o *AiAgentMetricSeriesPoint) SetSeries(v string) {
	o.Series = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *AiAgentMetricSeriesPoint) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentMetricSeriesPoint) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *AiAgentMetricSeriesPoint) HasCategory() bool {
	return o != nil && o.Category != nil
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *AiAgentMetricSeriesPoint) SetCategory(v string) {
	o.Category = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentMetricSeriesPoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Timestamp != nil {
		if o.Timestamp.Nanosecond() == 0 {
			toSerialize["timestamp"] = o.Timestamp.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["timestamp"] = o.Timestamp.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	toSerialize["x"] = o.X
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	toSerialize["y"] = o.Y
	if o.Series != nil {
		toSerialize["series"] = o.Series
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentMetricSeriesPoint) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Timestamp *time.Time `json:"timestamp,omitempty"`
		X         *string    `json:"x"`
		Value     *float64   `json:"value,omitempty"`
		Y         *float64   `json:"y"`
		Series    *string    `json:"series,omitempty"`
		Category  *string    `json:"category,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.X == nil {
		return fmt.Errorf("required field x missing")
	}
	if all.Y == nil {
		return fmt.Errorf("required field y missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"timestamp", "x", "value", "y", "series", "category"})
	} else {
		return err
	}
	o.Timestamp = all.Timestamp
	o.X = *all.X
	o.Value = all.Value
	o.Y = *all.Y
	o.Series = all.Series
	o.Category = all.Category

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
