// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticDBPerformanceMetric One database performance metric with threshold and source quality.
type DiagnosticDBPerformanceMetric struct {
	// Display value when available.
	Value *string `json:"value,omitempty"`
	// Metric unit.
	Unit *string `json:"unit,omitempty"`
	// Utilization ratio when available.
	Ratio *string `json:"ratio,omitempty"`
	// Threshold rule used for this metric.
	Threshold *string `json:"threshold,omitempty"`
	// User-facing severity and priority bucket for one diagnostic check.
	Severity DiagnosticCheckSeverity `json:"severity"`
	// Data source used for this metric.
	Source string `json:"source"`
	// Data quality status for one database performance metric.
	DataQuality DiagnosticMetricDataQuality `json:"dataQuality"`
	// Structured reason code for unavailable or degraded database performance metrics.
	ReasonCode *DiagnosticDBPerformanceMetricReasonCode `json:"reasonCode,omitempty"`
	// Reason when the metric is unavailable or degraded.
	Reason string `json:"reason"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticDBPerformanceMetric instantiates a new DiagnosticDBPerformanceMetric object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticDBPerformanceMetric(severity DiagnosticCheckSeverity, source string, dataQuality DiagnosticMetricDataQuality, reason string) *DiagnosticDBPerformanceMetric {
	this := DiagnosticDBPerformanceMetric{}
	this.Severity = severity
	this.Source = source
	this.DataQuality = dataQuality
	this.Reason = reason
	return &this
}

// NewDiagnosticDBPerformanceMetricWithDefaults instantiates a new DiagnosticDBPerformanceMetric object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticDBPerformanceMetricWithDefaults() *DiagnosticDBPerformanceMetric {
	this := DiagnosticDBPerformanceMetric{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DiagnosticDBPerformanceMetric) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceMetric) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DiagnosticDBPerformanceMetric) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *DiagnosticDBPerformanceMetric) SetValue(v string) {
	o.Value = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *DiagnosticDBPerformanceMetric) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceMetric) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *DiagnosticDBPerformanceMetric) HasUnit() bool {
	return o != nil && o.Unit != nil
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *DiagnosticDBPerformanceMetric) SetUnit(v string) {
	o.Unit = &v
}

// GetRatio returns the Ratio field value if set, zero value otherwise.
func (o *DiagnosticDBPerformanceMetric) GetRatio() string {
	if o == nil || o.Ratio == nil {
		var ret string
		return ret
	}
	return *o.Ratio
}

// GetRatioOk returns a tuple with the Ratio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceMetric) GetRatioOk() (*string, bool) {
	if o == nil || o.Ratio == nil {
		return nil, false
	}
	return o.Ratio, true
}

// HasRatio returns a boolean if a field has been set.
func (o *DiagnosticDBPerformanceMetric) HasRatio() bool {
	return o != nil && o.Ratio != nil
}

// SetRatio gets a reference to the given string and assigns it to the Ratio field.
func (o *DiagnosticDBPerformanceMetric) SetRatio(v string) {
	o.Ratio = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *DiagnosticDBPerformanceMetric) GetThreshold() string {
	if o == nil || o.Threshold == nil {
		var ret string
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceMetric) GetThresholdOk() (*string, bool) {
	if o == nil || o.Threshold == nil {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *DiagnosticDBPerformanceMetric) HasThreshold() bool {
	return o != nil && o.Threshold != nil
}

// SetThreshold gets a reference to the given string and assigns it to the Threshold field.
func (o *DiagnosticDBPerformanceMetric) SetThreshold(v string) {
	o.Threshold = &v
}

// GetSeverity returns the Severity field value.
func (o *DiagnosticDBPerformanceMetric) GetSeverity() DiagnosticCheckSeverity {
	if o == nil {
		var ret DiagnosticCheckSeverity
		return ret
	}
	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceMetric) GetSeverityOk() (*DiagnosticCheckSeverity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value.
func (o *DiagnosticDBPerformanceMetric) SetSeverity(v DiagnosticCheckSeverity) {
	o.Severity = v
}

// GetSource returns the Source field value.
func (o *DiagnosticDBPerformanceMetric) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceMetric) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *DiagnosticDBPerformanceMetric) SetSource(v string) {
	o.Source = v
}

// GetDataQuality returns the DataQuality field value.
func (o *DiagnosticDBPerformanceMetric) GetDataQuality() DiagnosticMetricDataQuality {
	if o == nil {
		var ret DiagnosticMetricDataQuality
		return ret
	}
	return o.DataQuality
}

// GetDataQualityOk returns a tuple with the DataQuality field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceMetric) GetDataQualityOk() (*DiagnosticMetricDataQuality, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataQuality, true
}

// SetDataQuality sets field value.
func (o *DiagnosticDBPerformanceMetric) SetDataQuality(v DiagnosticMetricDataQuality) {
	o.DataQuality = v
}

// GetReasonCode returns the ReasonCode field value if set, zero value otherwise.
func (o *DiagnosticDBPerformanceMetric) GetReasonCode() DiagnosticDBPerformanceMetricReasonCode {
	if o == nil || o.ReasonCode == nil {
		var ret DiagnosticDBPerformanceMetricReasonCode
		return ret
	}
	return *o.ReasonCode
}

// GetReasonCodeOk returns a tuple with the ReasonCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceMetric) GetReasonCodeOk() (*DiagnosticDBPerformanceMetricReasonCode, bool) {
	if o == nil || o.ReasonCode == nil {
		return nil, false
	}
	return o.ReasonCode, true
}

// HasReasonCode returns a boolean if a field has been set.
func (o *DiagnosticDBPerformanceMetric) HasReasonCode() bool {
	return o != nil && o.ReasonCode != nil
}

// SetReasonCode gets a reference to the given DiagnosticDBPerformanceMetricReasonCode and assigns it to the ReasonCode field.
func (o *DiagnosticDBPerformanceMetric) SetReasonCode(v DiagnosticDBPerformanceMetricReasonCode) {
	o.ReasonCode = &v
}

// GetReason returns the Reason field value.
func (o *DiagnosticDBPerformanceMetric) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceMetric) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value.
func (o *DiagnosticDBPerformanceMetric) SetReason(v string) {
	o.Reason = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticDBPerformanceMetric) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}
	if o.Ratio != nil {
		toSerialize["ratio"] = o.Ratio
	}
	if o.Threshold != nil {
		toSerialize["threshold"] = o.Threshold
	}
	toSerialize["severity"] = o.Severity
	toSerialize["source"] = o.Source
	toSerialize["dataQuality"] = o.DataQuality
	if o.ReasonCode != nil {
		toSerialize["reasonCode"] = o.ReasonCode
	}
	toSerialize["reason"] = o.Reason

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticDBPerformanceMetric) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Value       *string                                  `json:"value,omitempty"`
		Unit        *string                                  `json:"unit,omitempty"`
		Ratio       *string                                  `json:"ratio,omitempty"`
		Threshold   *string                                  `json:"threshold,omitempty"`
		Severity    *DiagnosticCheckSeverity                 `json:"severity"`
		Source      *string                                  `json:"source"`
		DataQuality *DiagnosticMetricDataQuality             `json:"dataQuality"`
		ReasonCode  *DiagnosticDBPerformanceMetricReasonCode `json:"reasonCode,omitempty"`
		Reason      *string                                  `json:"reason"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Severity == nil {
		return fmt.Errorf("required field severity missing")
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	if all.DataQuality == nil {
		return fmt.Errorf("required field dataQuality missing")
	}
	if all.Reason == nil {
		return fmt.Errorf("required field reason missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"value", "unit", "ratio", "threshold", "severity", "source", "dataQuality", "reasonCode", "reason"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Value = all.Value
	o.Unit = all.Unit
	o.Ratio = all.Ratio
	o.Threshold = all.Threshold
	if !all.Severity.IsValid() {
		hasInvalidField = true
	} else {
		o.Severity = *all.Severity
	}
	o.Source = *all.Source
	if !all.DataQuality.IsValid() {
		hasInvalidField = true
	} else {
		o.DataQuality = *all.DataQuality
	}
	if all.ReasonCode != nil && !all.ReasonCode.IsValid() {
		hasInvalidField = true
	} else {
		o.ReasonCode = all.ReasonCode
	}
	o.Reason = *all.Reason

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
