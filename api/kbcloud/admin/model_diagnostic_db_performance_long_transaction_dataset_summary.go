// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticDBPerformanceLongTransactionDatasetSummary Summary of the long-transaction sample set.
type DiagnosticDBPerformanceLongTransactionDatasetSummary struct {
	// Number of sessions whose transaction age exceeded the readonly threshold.
	TotalCount int64 `json:"totalCount"`
	// Minimum transaction age, in seconds, used by the readonly collector.
	ThresholdSeconds int64 `json:"thresholdSeconds"`
	// Time when the readonly long-transaction collector produced this sample set.
	CollectedAt time.Time `json:"collectedAt"`
	// Top long transactions sorted by transaction duration.
	ActiveTransactions []DiagnosticDBPerformanceLongTransactionSample `json:"activeTransactions"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticDBPerformanceLongTransactionDatasetSummary instantiates a new DiagnosticDBPerformanceLongTransactionDatasetSummary object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticDBPerformanceLongTransactionDatasetSummary(totalCount int64, thresholdSeconds int64, collectedAt time.Time, activeTransactions []DiagnosticDBPerformanceLongTransactionSample) *DiagnosticDBPerformanceLongTransactionDatasetSummary {
	this := DiagnosticDBPerformanceLongTransactionDatasetSummary{}
	this.TotalCount = totalCount
	this.ThresholdSeconds = thresholdSeconds
	this.CollectedAt = collectedAt
	this.ActiveTransactions = activeTransactions
	return &this
}

// NewDiagnosticDBPerformanceLongTransactionDatasetSummaryWithDefaults instantiates a new DiagnosticDBPerformanceLongTransactionDatasetSummary object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticDBPerformanceLongTransactionDatasetSummaryWithDefaults() *DiagnosticDBPerformanceLongTransactionDatasetSummary {
	this := DiagnosticDBPerformanceLongTransactionDatasetSummary{}
	return &this
}

// GetTotalCount returns the TotalCount field value.
func (o *DiagnosticDBPerformanceLongTransactionDatasetSummary) GetTotalCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLongTransactionDatasetSummary) GetTotalCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value.
func (o *DiagnosticDBPerformanceLongTransactionDatasetSummary) SetTotalCount(v int64) {
	o.TotalCount = v
}

// GetThresholdSeconds returns the ThresholdSeconds field value.
func (o *DiagnosticDBPerformanceLongTransactionDatasetSummary) GetThresholdSeconds() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ThresholdSeconds
}

// GetThresholdSecondsOk returns a tuple with the ThresholdSeconds field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLongTransactionDatasetSummary) GetThresholdSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThresholdSeconds, true
}

// SetThresholdSeconds sets field value.
func (o *DiagnosticDBPerformanceLongTransactionDatasetSummary) SetThresholdSeconds(v int64) {
	o.ThresholdSeconds = v
}

// GetCollectedAt returns the CollectedAt field value.
func (o *DiagnosticDBPerformanceLongTransactionDatasetSummary) GetCollectedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CollectedAt
}

// GetCollectedAtOk returns a tuple with the CollectedAt field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLongTransactionDatasetSummary) GetCollectedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectedAt, true
}

// SetCollectedAt sets field value.
func (o *DiagnosticDBPerformanceLongTransactionDatasetSummary) SetCollectedAt(v time.Time) {
	o.CollectedAt = v
}

// GetActiveTransactions returns the ActiveTransactions field value.
func (o *DiagnosticDBPerformanceLongTransactionDatasetSummary) GetActiveTransactions() []DiagnosticDBPerformanceLongTransactionSample {
	if o == nil {
		var ret []DiagnosticDBPerformanceLongTransactionSample
		return ret
	}
	return o.ActiveTransactions
}

// GetActiveTransactionsOk returns a tuple with the ActiveTransactions field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLongTransactionDatasetSummary) GetActiveTransactionsOk() (*[]DiagnosticDBPerformanceLongTransactionSample, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActiveTransactions, true
}

// SetActiveTransactions sets field value.
func (o *DiagnosticDBPerformanceLongTransactionDatasetSummary) SetActiveTransactions(v []DiagnosticDBPerformanceLongTransactionSample) {
	o.ActiveTransactions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticDBPerformanceLongTransactionDatasetSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["totalCount"] = o.TotalCount
	toSerialize["thresholdSeconds"] = o.ThresholdSeconds
	if o.CollectedAt.Nanosecond() == 0 {
		toSerialize["collectedAt"] = o.CollectedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["collectedAt"] = o.CollectedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["activeTransactions"] = o.ActiveTransactions

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticDBPerformanceLongTransactionDatasetSummary) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TotalCount         *int64                                          `json:"totalCount"`
		ThresholdSeconds   *int64                                          `json:"thresholdSeconds"`
		CollectedAt        *time.Time                                      `json:"collectedAt"`
		ActiveTransactions *[]DiagnosticDBPerformanceLongTransactionSample `json:"activeTransactions"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.TotalCount == nil {
		return fmt.Errorf("required field totalCount missing")
	}
	if all.ThresholdSeconds == nil {
		return fmt.Errorf("required field thresholdSeconds missing")
	}
	if all.CollectedAt == nil {
		return fmt.Errorf("required field collectedAt missing")
	}
	if all.ActiveTransactions == nil {
		return fmt.Errorf("required field activeTransactions missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"totalCount", "thresholdSeconds", "collectedAt", "activeTransactions"})
	} else {
		return err
	}
	o.TotalCount = *all.TotalCount
	o.ThresholdSeconds = *all.ThresholdSeconds
	o.CollectedAt = *all.CollectedAt
	o.ActiveTransactions = *all.ActiveTransactions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
