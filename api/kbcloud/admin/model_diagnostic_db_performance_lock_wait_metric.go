// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticDBPerformanceLockWaitMetric PostgreSQL lock-wait performance metric. It exposes a three-state live
// collector contract (has data / no data / collection failed) plus the
// source_missing compatibility state when collection is disabled or not
// integrated. Full SQL text is never exposed; only digests are returned.
type DiagnosticDBPerformanceLockWaitMetric struct {
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
	// Data source used for this metric, such as pg_locks.
	Source string `json:"source"`
	// Data quality status for one database performance metric.
	DataQuality DiagnosticMetricDataQuality `json:"dataQuality"`
	// Structured reason code for unavailable or degraded database performance metrics.
	ReasonCode *DiagnosticDBPerformanceMetricReasonCode `json:"reasonCode,omitempty"`
	// Reason when the metric is unavailable or degraded.
	Reason string `json:"reason"`
	// Business-result axis for the PostgreSQL lock-wait metric. Live collection
	// produces the three states source_present_with_data, source_present_no_data,
	// and source_present_collection_failed. source_missing is the compatibility
	// state used when the collector is not enabled or no PostgreSQL pod is
	// available.
	//
	LockWaitState DiagnosticDBPerformanceLockWaitState `json:"lockWaitState"`
	// Summary of the lock-wait sample set.
	DatasetSummary *DiagnosticDBPerformanceLockWaitDatasetSummary `json:"datasetSummary,omitempty"`
	// Structured collection error detail when lockWaitState is source_present_collection_failed.
	CollectionError *DiagnosticDBPerformanceLockWaitCollectionError `json:"collectionError,omitempty"`
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

// NewDiagnosticDBPerformanceLockWaitMetric instantiates a new DiagnosticDBPerformanceLockWaitMetric object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticDBPerformanceLockWaitMetric(severity DiagnosticCheckSeverity, source string, dataQuality DiagnosticMetricDataQuality, reason string, lockWaitState DiagnosticDBPerformanceLockWaitState, detail DiagnosticDBPerformanceMetricDetail) *DiagnosticDBPerformanceLockWaitMetric {
	this := DiagnosticDBPerformanceLockWaitMetric{}
	this.Severity = severity
	this.Source = source
	this.DataQuality = dataQuality
	this.Reason = reason
	this.LockWaitState = lockWaitState
	this.Detail = detail
	return &this
}

// NewDiagnosticDBPerformanceLockWaitMetricWithDefaults instantiates a new DiagnosticDBPerformanceLockWaitMetric object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticDBPerformanceLockWaitMetricWithDefaults() *DiagnosticDBPerformanceLockWaitMetric {
	this := DiagnosticDBPerformanceLockWaitMetric{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DiagnosticDBPerformanceLockWaitMetric) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLockWaitMetric) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DiagnosticDBPerformanceLockWaitMetric) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *DiagnosticDBPerformanceLockWaitMetric) SetValue(v string) {
	o.Value = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *DiagnosticDBPerformanceLockWaitMetric) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLockWaitMetric) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *DiagnosticDBPerformanceLockWaitMetric) HasUnit() bool {
	return o != nil && o.Unit != nil
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *DiagnosticDBPerformanceLockWaitMetric) SetUnit(v string) {
	o.Unit = &v
}

// GetRatio returns the Ratio field value if set, zero value otherwise.
func (o *DiagnosticDBPerformanceLockWaitMetric) GetRatio() string {
	if o == nil || o.Ratio == nil {
		var ret string
		return ret
	}
	return *o.Ratio
}

// GetRatioOk returns a tuple with the Ratio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLockWaitMetric) GetRatioOk() (*string, bool) {
	if o == nil || o.Ratio == nil {
		return nil, false
	}
	return o.Ratio, true
}

// HasRatio returns a boolean if a field has been set.
func (o *DiagnosticDBPerformanceLockWaitMetric) HasRatio() bool {
	return o != nil && o.Ratio != nil
}

// SetRatio gets a reference to the given string and assigns it to the Ratio field.
func (o *DiagnosticDBPerformanceLockWaitMetric) SetRatio(v string) {
	o.Ratio = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *DiagnosticDBPerformanceLockWaitMetric) GetThreshold() string {
	if o == nil || o.Threshold == nil {
		var ret string
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLockWaitMetric) GetThresholdOk() (*string, bool) {
	if o == nil || o.Threshold == nil {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *DiagnosticDBPerformanceLockWaitMetric) HasThreshold() bool {
	return o != nil && o.Threshold != nil
}

// SetThreshold gets a reference to the given string and assigns it to the Threshold field.
func (o *DiagnosticDBPerformanceLockWaitMetric) SetThreshold(v string) {
	o.Threshold = &v
}

// GetSeverity returns the Severity field value.
func (o *DiagnosticDBPerformanceLockWaitMetric) GetSeverity() DiagnosticCheckSeverity {
	if o == nil {
		var ret DiagnosticCheckSeverity
		return ret
	}
	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLockWaitMetric) GetSeverityOk() (*DiagnosticCheckSeverity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value.
func (o *DiagnosticDBPerformanceLockWaitMetric) SetSeverity(v DiagnosticCheckSeverity) {
	o.Severity = v
}

// GetSource returns the Source field value.
func (o *DiagnosticDBPerformanceLockWaitMetric) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLockWaitMetric) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *DiagnosticDBPerformanceLockWaitMetric) SetSource(v string) {
	o.Source = v
}

// GetDataQuality returns the DataQuality field value.
func (o *DiagnosticDBPerformanceLockWaitMetric) GetDataQuality() DiagnosticMetricDataQuality {
	if o == nil {
		var ret DiagnosticMetricDataQuality
		return ret
	}
	return o.DataQuality
}

// GetDataQualityOk returns a tuple with the DataQuality field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLockWaitMetric) GetDataQualityOk() (*DiagnosticMetricDataQuality, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataQuality, true
}

// SetDataQuality sets field value.
func (o *DiagnosticDBPerformanceLockWaitMetric) SetDataQuality(v DiagnosticMetricDataQuality) {
	o.DataQuality = v
}

// GetReasonCode returns the ReasonCode field value if set, zero value otherwise.
func (o *DiagnosticDBPerformanceLockWaitMetric) GetReasonCode() DiagnosticDBPerformanceMetricReasonCode {
	if o == nil || o.ReasonCode == nil {
		var ret DiagnosticDBPerformanceMetricReasonCode
		return ret
	}
	return *o.ReasonCode
}

// GetReasonCodeOk returns a tuple with the ReasonCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLockWaitMetric) GetReasonCodeOk() (*DiagnosticDBPerformanceMetricReasonCode, bool) {
	if o == nil || o.ReasonCode == nil {
		return nil, false
	}
	return o.ReasonCode, true
}

// HasReasonCode returns a boolean if a field has been set.
func (o *DiagnosticDBPerformanceLockWaitMetric) HasReasonCode() bool {
	return o != nil && o.ReasonCode != nil
}

// SetReasonCode gets a reference to the given DiagnosticDBPerformanceMetricReasonCode and assigns it to the ReasonCode field.
func (o *DiagnosticDBPerformanceLockWaitMetric) SetReasonCode(v DiagnosticDBPerformanceMetricReasonCode) {
	o.ReasonCode = &v
}

// GetReason returns the Reason field value.
func (o *DiagnosticDBPerformanceLockWaitMetric) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLockWaitMetric) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value.
func (o *DiagnosticDBPerformanceLockWaitMetric) SetReason(v string) {
	o.Reason = v
}

// GetLockWaitState returns the LockWaitState field value.
func (o *DiagnosticDBPerformanceLockWaitMetric) GetLockWaitState() DiagnosticDBPerformanceLockWaitState {
	if o == nil {
		var ret DiagnosticDBPerformanceLockWaitState
		return ret
	}
	return o.LockWaitState
}

// GetLockWaitStateOk returns a tuple with the LockWaitState field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLockWaitMetric) GetLockWaitStateOk() (*DiagnosticDBPerformanceLockWaitState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LockWaitState, true
}

// SetLockWaitState sets field value.
func (o *DiagnosticDBPerformanceLockWaitMetric) SetLockWaitState(v DiagnosticDBPerformanceLockWaitState) {
	o.LockWaitState = v
}

// GetDatasetSummary returns the DatasetSummary field value if set, zero value otherwise.
func (o *DiagnosticDBPerformanceLockWaitMetric) GetDatasetSummary() DiagnosticDBPerformanceLockWaitDatasetSummary {
	if o == nil || o.DatasetSummary == nil {
		var ret DiagnosticDBPerformanceLockWaitDatasetSummary
		return ret
	}
	return *o.DatasetSummary
}

// GetDatasetSummaryOk returns a tuple with the DatasetSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLockWaitMetric) GetDatasetSummaryOk() (*DiagnosticDBPerformanceLockWaitDatasetSummary, bool) {
	if o == nil || o.DatasetSummary == nil {
		return nil, false
	}
	return o.DatasetSummary, true
}

// HasDatasetSummary returns a boolean if a field has been set.
func (o *DiagnosticDBPerformanceLockWaitMetric) HasDatasetSummary() bool {
	return o != nil && o.DatasetSummary != nil
}

// SetDatasetSummary gets a reference to the given DiagnosticDBPerformanceLockWaitDatasetSummary and assigns it to the DatasetSummary field.
func (o *DiagnosticDBPerformanceLockWaitMetric) SetDatasetSummary(v DiagnosticDBPerformanceLockWaitDatasetSummary) {
	o.DatasetSummary = &v
}

// GetCollectionError returns the CollectionError field value if set, zero value otherwise.
func (o *DiagnosticDBPerformanceLockWaitMetric) GetCollectionError() DiagnosticDBPerformanceLockWaitCollectionError {
	if o == nil || o.CollectionError == nil {
		var ret DiagnosticDBPerformanceLockWaitCollectionError
		return ret
	}
	return *o.CollectionError
}

// GetCollectionErrorOk returns a tuple with the CollectionError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLockWaitMetric) GetCollectionErrorOk() (*DiagnosticDBPerformanceLockWaitCollectionError, bool) {
	if o == nil || o.CollectionError == nil {
		return nil, false
	}
	return o.CollectionError, true
}

// HasCollectionError returns a boolean if a field has been set.
func (o *DiagnosticDBPerformanceLockWaitMetric) HasCollectionError() bool {
	return o != nil && o.CollectionError != nil
}

// SetCollectionError gets a reference to the given DiagnosticDBPerformanceLockWaitCollectionError and assigns it to the CollectionError field.
func (o *DiagnosticDBPerformanceLockWaitMetric) SetCollectionError(v DiagnosticDBPerformanceLockWaitCollectionError) {
	o.CollectionError = &v
}

// GetDetail returns the Detail field value.
func (o *DiagnosticDBPerformanceLockWaitMetric) GetDetail() DiagnosticDBPerformanceMetricDetail {
	if o == nil {
		var ret DiagnosticDBPerformanceMetricDetail
		return ret
	}
	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLockWaitMetric) GetDetailOk() (*DiagnosticDBPerformanceMetricDetail, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Detail, true
}

// SetDetail sets field value.
func (o *DiagnosticDBPerformanceLockWaitMetric) SetDetail(v DiagnosticDBPerformanceMetricDetail) {
	o.Detail = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticDBPerformanceLockWaitMetric) MarshalJSON() ([]byte, error) {
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
	toSerialize["lockWaitState"] = o.LockWaitState
	if o.DatasetSummary != nil {
		toSerialize["datasetSummary"] = o.DatasetSummary
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
func (o *DiagnosticDBPerformanceLockWaitMetric) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Value           *string                                         `json:"value,omitempty"`
		Unit            *string                                         `json:"unit,omitempty"`
		Ratio           *string                                         `json:"ratio,omitempty"`
		Threshold       *string                                         `json:"threshold,omitempty"`
		Severity        *DiagnosticCheckSeverity                        `json:"severity"`
		Source          *string                                         `json:"source"`
		DataQuality     *DiagnosticMetricDataQuality                    `json:"dataQuality"`
		ReasonCode      *DiagnosticDBPerformanceMetricReasonCode        `json:"reasonCode,omitempty"`
		Reason          *string                                         `json:"reason"`
		LockWaitState   *DiagnosticDBPerformanceLockWaitState           `json:"lockWaitState"`
		DatasetSummary  *DiagnosticDBPerformanceLockWaitDatasetSummary  `json:"datasetSummary,omitempty"`
		CollectionError *DiagnosticDBPerformanceLockWaitCollectionError `json:"collectionError,omitempty"`
		Detail          *DiagnosticDBPerformanceMetricDetail            `json:"detail"`
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
	if all.LockWaitState == nil {
		return fmt.Errorf("required field lockWaitState missing")
	}
	if all.Detail == nil {
		return fmt.Errorf("required field detail missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"value", "unit", "ratio", "threshold", "severity", "source", "dataQuality", "reasonCode", "reason", "lockWaitState", "datasetSummary", "collectionError", "detail"})
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
	if !all.LockWaitState.IsValid() {
		hasInvalidField = true
	} else {
		o.LockWaitState = *all.LockWaitState
	}
	if all.DatasetSummary != nil && all.DatasetSummary.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DatasetSummary = all.DatasetSummary
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
