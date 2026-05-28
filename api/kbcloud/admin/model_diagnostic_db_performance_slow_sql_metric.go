// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticDBPerformanceSlowSQLMetric Richer slow SQL performance metric. Inlines the generic metric fields and adds
// Sprint 9 slow SQL specific result fields. Cross-field consistency between
// slowSqlState, dataQuality, reasonCode and datasetSummary.totalCount is enforced
// by backend loader Validate (see Sprint 9 v0.1 contract).
type DiagnosticDBPerformanceSlowSQLMetric struct {
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
	// Data source used for this metric, such as pg_stat_statements.
	Source string `json:"source"`
	// Data quality status for one database performance metric.
	DataQuality DiagnosticMetricDataQuality `json:"dataQuality"`
	// Structured reason code for unavailable or degraded database performance metrics.
	ReasonCode *DiagnosticDBPerformanceMetricReasonCode `json:"reasonCode,omitempty"`
	// Reason when the metric is unavailable or degraded.
	Reason string `json:"reason"`
	// Business-result axis for the slow SQL metric. Sprint 9 v0.1 four-state contract.
	// Separate from dataQuality (collection-quality axis) so the Console can render the
	// correct main text without overloading dataQuality. Each enum value pairs with a
	// fixed (dataQuality, reasonCode, count) tuple enforced by backend Validate.
	//
	SlowSqlState DiagnosticDBPerformanceSlowSQLState `json:"slowSqlState"`
	// Summary of the slow SQL sample set. REQUIRED when slowSqlState is
	// source_present_with_data or source_present_no_data_in_window;
	// FORBIDDEN otherwise. Cross-field consistency enforced by
	// ValidateSlowSQLStateConsistency.
	//
	DatasetSummary *DiagnosticDBPerformanceSlowSQLDatasetSummary `json:"datasetSummary,omitempty"`
	// Collection time window for one slow SQL sample set. Required when slowSqlState is source_present_with_data or source_present_no_data_in_window.
	TimeWindow *DiagnosticDBPerformanceSlowSQLTimeWindow `json:"timeWindow,omitempty"`
	// Structured collection error detail when slowSqlState is source_present_collection_failed.
	CollectionError *DiagnosticDBPerformanceSlowSQLCollectionError `json:"collectionError,omitempty"`
	// User-facing detail and DBA guidance for one database performance card.
	// It explains why the metric matters, what evidence was checked, and the
	// next action without requiring Console to infer meaning from rawEvidence.
	// SQL text in this structure must be redacted or summarized.
	//
	Detail DiagnosticDBPerformanceMetricDetail `json:"detail"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticDBPerformanceSlowSQLMetric instantiates a new DiagnosticDBPerformanceSlowSQLMetric object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticDBPerformanceSlowSQLMetric(severity DiagnosticCheckSeverity, source string, dataQuality DiagnosticMetricDataQuality, reason string, slowSqlState DiagnosticDBPerformanceSlowSQLState, detail DiagnosticDBPerformanceMetricDetail) *DiagnosticDBPerformanceSlowSQLMetric {
	this := DiagnosticDBPerformanceSlowSQLMetric{}
	this.Severity = severity
	this.Source = source
	this.DataQuality = dataQuality
	this.Reason = reason
	this.SlowSqlState = slowSqlState
	this.Detail = detail
	return &this
}

// NewDiagnosticDBPerformanceSlowSQLMetricWithDefaults instantiates a new DiagnosticDBPerformanceSlowSQLMetric object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticDBPerformanceSlowSQLMetricWithDefaults() *DiagnosticDBPerformanceSlowSQLMetric {
	this := DiagnosticDBPerformanceSlowSQLMetric{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DiagnosticDBPerformanceSlowSQLMetric) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceSlowSQLMetric) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DiagnosticDBPerformanceSlowSQLMetric) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *DiagnosticDBPerformanceSlowSQLMetric) SetValue(v string) {
	o.Value = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *DiagnosticDBPerformanceSlowSQLMetric) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceSlowSQLMetric) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *DiagnosticDBPerformanceSlowSQLMetric) HasUnit() bool {
	return o != nil && o.Unit != nil
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *DiagnosticDBPerformanceSlowSQLMetric) SetUnit(v string) {
	o.Unit = &v
}

// GetRatio returns the Ratio field value if set, zero value otherwise.
func (o *DiagnosticDBPerformanceSlowSQLMetric) GetRatio() string {
	if o == nil || o.Ratio == nil {
		var ret string
		return ret
	}
	return *o.Ratio
}

// GetRatioOk returns a tuple with the Ratio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceSlowSQLMetric) GetRatioOk() (*string, bool) {
	if o == nil || o.Ratio == nil {
		return nil, false
	}
	return o.Ratio, true
}

// HasRatio returns a boolean if a field has been set.
func (o *DiagnosticDBPerformanceSlowSQLMetric) HasRatio() bool {
	return o != nil && o.Ratio != nil
}

// SetRatio gets a reference to the given string and assigns it to the Ratio field.
func (o *DiagnosticDBPerformanceSlowSQLMetric) SetRatio(v string) {
	o.Ratio = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *DiagnosticDBPerformanceSlowSQLMetric) GetThreshold() string {
	if o == nil || o.Threshold == nil {
		var ret string
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceSlowSQLMetric) GetThresholdOk() (*string, bool) {
	if o == nil || o.Threshold == nil {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *DiagnosticDBPerformanceSlowSQLMetric) HasThreshold() bool {
	return o != nil && o.Threshold != nil
}

// SetThreshold gets a reference to the given string and assigns it to the Threshold field.
func (o *DiagnosticDBPerformanceSlowSQLMetric) SetThreshold(v string) {
	o.Threshold = &v
}

// GetSeverity returns the Severity field value.
func (o *DiagnosticDBPerformanceSlowSQLMetric) GetSeverity() DiagnosticCheckSeverity {
	if o == nil {
		var ret DiagnosticCheckSeverity
		return ret
	}
	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceSlowSQLMetric) GetSeverityOk() (*DiagnosticCheckSeverity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value.
func (o *DiagnosticDBPerformanceSlowSQLMetric) SetSeverity(v DiagnosticCheckSeverity) {
	o.Severity = v
}

// GetSource returns the Source field value.
func (o *DiagnosticDBPerformanceSlowSQLMetric) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceSlowSQLMetric) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *DiagnosticDBPerformanceSlowSQLMetric) SetSource(v string) {
	o.Source = v
}

// GetDataQuality returns the DataQuality field value.
func (o *DiagnosticDBPerformanceSlowSQLMetric) GetDataQuality() DiagnosticMetricDataQuality {
	if o == nil {
		var ret DiagnosticMetricDataQuality
		return ret
	}
	return o.DataQuality
}

// GetDataQualityOk returns a tuple with the DataQuality field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceSlowSQLMetric) GetDataQualityOk() (*DiagnosticMetricDataQuality, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataQuality, true
}

// SetDataQuality sets field value.
func (o *DiagnosticDBPerformanceSlowSQLMetric) SetDataQuality(v DiagnosticMetricDataQuality) {
	o.DataQuality = v
}

// GetReasonCode returns the ReasonCode field value if set, zero value otherwise.
func (o *DiagnosticDBPerformanceSlowSQLMetric) GetReasonCode() DiagnosticDBPerformanceMetricReasonCode {
	if o == nil || o.ReasonCode == nil {
		var ret DiagnosticDBPerformanceMetricReasonCode
		return ret
	}
	return *o.ReasonCode
}

// GetReasonCodeOk returns a tuple with the ReasonCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceSlowSQLMetric) GetReasonCodeOk() (*DiagnosticDBPerformanceMetricReasonCode, bool) {
	if o == nil || o.ReasonCode == nil {
		return nil, false
	}
	return o.ReasonCode, true
}

// HasReasonCode returns a boolean if a field has been set.
func (o *DiagnosticDBPerformanceSlowSQLMetric) HasReasonCode() bool {
	return o != nil && o.ReasonCode != nil
}

// SetReasonCode gets a reference to the given DiagnosticDBPerformanceMetricReasonCode and assigns it to the ReasonCode field.
func (o *DiagnosticDBPerformanceSlowSQLMetric) SetReasonCode(v DiagnosticDBPerformanceMetricReasonCode) {
	o.ReasonCode = &v
}

// GetReason returns the Reason field value.
func (o *DiagnosticDBPerformanceSlowSQLMetric) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceSlowSQLMetric) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value.
func (o *DiagnosticDBPerformanceSlowSQLMetric) SetReason(v string) {
	o.Reason = v
}

// GetSlowSqlState returns the SlowSqlState field value.
func (o *DiagnosticDBPerformanceSlowSQLMetric) GetSlowSqlState() DiagnosticDBPerformanceSlowSQLState {
	if o == nil {
		var ret DiagnosticDBPerformanceSlowSQLState
		return ret
	}
	return o.SlowSqlState
}

// GetSlowSqlStateOk returns a tuple with the SlowSqlState field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceSlowSQLMetric) GetSlowSqlStateOk() (*DiagnosticDBPerformanceSlowSQLState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SlowSqlState, true
}

// SetSlowSqlState sets field value.
func (o *DiagnosticDBPerformanceSlowSQLMetric) SetSlowSqlState(v DiagnosticDBPerformanceSlowSQLState) {
	o.SlowSqlState = v
}

// GetDatasetSummary returns the DatasetSummary field value if set, zero value otherwise.
func (o *DiagnosticDBPerformanceSlowSQLMetric) GetDatasetSummary() DiagnosticDBPerformanceSlowSQLDatasetSummary {
	if o == nil || o.DatasetSummary == nil {
		var ret DiagnosticDBPerformanceSlowSQLDatasetSummary
		return ret
	}
	return *o.DatasetSummary
}

// GetDatasetSummaryOk returns a tuple with the DatasetSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceSlowSQLMetric) GetDatasetSummaryOk() (*DiagnosticDBPerformanceSlowSQLDatasetSummary, bool) {
	if o == nil || o.DatasetSummary == nil {
		return nil, false
	}
	return o.DatasetSummary, true
}

// HasDatasetSummary returns a boolean if a field has been set.
func (o *DiagnosticDBPerformanceSlowSQLMetric) HasDatasetSummary() bool {
	return o != nil && o.DatasetSummary != nil
}

// SetDatasetSummary gets a reference to the given DiagnosticDBPerformanceSlowSQLDatasetSummary and assigns it to the DatasetSummary field.
func (o *DiagnosticDBPerformanceSlowSQLMetric) SetDatasetSummary(v DiagnosticDBPerformanceSlowSQLDatasetSummary) {
	o.DatasetSummary = &v
}

// GetTimeWindow returns the TimeWindow field value if set, zero value otherwise.
func (o *DiagnosticDBPerformanceSlowSQLMetric) GetTimeWindow() DiagnosticDBPerformanceSlowSQLTimeWindow {
	if o == nil || o.TimeWindow == nil {
		var ret DiagnosticDBPerformanceSlowSQLTimeWindow
		return ret
	}
	return *o.TimeWindow
}

// GetTimeWindowOk returns a tuple with the TimeWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceSlowSQLMetric) GetTimeWindowOk() (*DiagnosticDBPerformanceSlowSQLTimeWindow, bool) {
	if o == nil || o.TimeWindow == nil {
		return nil, false
	}
	return o.TimeWindow, true
}

// HasTimeWindow returns a boolean if a field has been set.
func (o *DiagnosticDBPerformanceSlowSQLMetric) HasTimeWindow() bool {
	return o != nil && o.TimeWindow != nil
}

// SetTimeWindow gets a reference to the given DiagnosticDBPerformanceSlowSQLTimeWindow and assigns it to the TimeWindow field.
func (o *DiagnosticDBPerformanceSlowSQLMetric) SetTimeWindow(v DiagnosticDBPerformanceSlowSQLTimeWindow) {
	o.TimeWindow = &v
}

// GetCollectionError returns the CollectionError field value if set, zero value otherwise.
func (o *DiagnosticDBPerformanceSlowSQLMetric) GetCollectionError() DiagnosticDBPerformanceSlowSQLCollectionError {
	if o == nil || o.CollectionError == nil {
		var ret DiagnosticDBPerformanceSlowSQLCollectionError
		return ret
	}
	return *o.CollectionError
}

// GetCollectionErrorOk returns a tuple with the CollectionError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceSlowSQLMetric) GetCollectionErrorOk() (*DiagnosticDBPerformanceSlowSQLCollectionError, bool) {
	if o == nil || o.CollectionError == nil {
		return nil, false
	}
	return o.CollectionError, true
}

// HasCollectionError returns a boolean if a field has been set.
func (o *DiagnosticDBPerformanceSlowSQLMetric) HasCollectionError() bool {
	return o != nil && o.CollectionError != nil
}

// SetCollectionError gets a reference to the given DiagnosticDBPerformanceSlowSQLCollectionError and assigns it to the CollectionError field.
func (o *DiagnosticDBPerformanceSlowSQLMetric) SetCollectionError(v DiagnosticDBPerformanceSlowSQLCollectionError) {
	o.CollectionError = &v
}

// GetDetail returns the Detail field value.
func (o *DiagnosticDBPerformanceSlowSQLMetric) GetDetail() DiagnosticDBPerformanceMetricDetail {
	if o == nil {
		var ret DiagnosticDBPerformanceMetricDetail
		return ret
	}
	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceSlowSQLMetric) GetDetailOk() (*DiagnosticDBPerformanceMetricDetail, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Detail, true
}

// SetDetail sets field value.
func (o *DiagnosticDBPerformanceSlowSQLMetric) SetDetail(v DiagnosticDBPerformanceMetricDetail) {
	o.Detail = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticDBPerformanceSlowSQLMetric) MarshalJSON() ([]byte, error) {
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
	toSerialize["slowSqlState"] = o.SlowSqlState
	if o.DatasetSummary != nil {
		toSerialize["datasetSummary"] = o.DatasetSummary
	}
	if o.TimeWindow != nil {
		toSerialize["timeWindow"] = o.TimeWindow
	}
	if o.CollectionError != nil {
		toSerialize["collectionError"] = o.CollectionError
	}
	toSerialize["detail"] = o.Detail

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticDBPerformanceSlowSQLMetric) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Value           *string                                        `json:"value,omitempty"`
		Unit            *string                                        `json:"unit,omitempty"`
		Ratio           *string                                        `json:"ratio,omitempty"`
		Threshold       *string                                        `json:"threshold,omitempty"`
		Severity        *DiagnosticCheckSeverity                       `json:"severity"`
		Source          *string                                        `json:"source"`
		DataQuality     *DiagnosticMetricDataQuality                   `json:"dataQuality"`
		ReasonCode      *DiagnosticDBPerformanceMetricReasonCode       `json:"reasonCode,omitempty"`
		Reason          *string                                        `json:"reason"`
		SlowSqlState    *DiagnosticDBPerformanceSlowSQLState           `json:"slowSqlState"`
		DatasetSummary  *DiagnosticDBPerformanceSlowSQLDatasetSummary  `json:"datasetSummary,omitempty"`
		TimeWindow      *DiagnosticDBPerformanceSlowSQLTimeWindow      `json:"timeWindow,omitempty"`
		CollectionError *DiagnosticDBPerformanceSlowSQLCollectionError `json:"collectionError,omitempty"`
		Detail          *DiagnosticDBPerformanceMetricDetail           `json:"detail"`
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
	if all.SlowSqlState == nil {
		return fmt.Errorf("required field slowSqlState missing")
	}
	if all.Detail == nil {
		return fmt.Errorf("required field detail missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"value", "unit", "ratio", "threshold", "severity", "source", "dataQuality", "reasonCode", "reason", "slowSqlState", "datasetSummary", "timeWindow", "collectionError", "detail"})
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
	if !all.SlowSqlState.IsValid() {
		hasInvalidField = true
	} else {
		o.SlowSqlState = *all.SlowSqlState
	}
	if all.DatasetSummary != nil && all.DatasetSummary.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DatasetSummary = all.DatasetSummary
	if all.TimeWindow != nil && all.TimeWindow.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TimeWindow = all.TimeWindow
	if all.CollectionError != nil && all.CollectionError.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CollectionError = all.CollectionError
	if all.Detail.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Detail = *all.Detail

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
