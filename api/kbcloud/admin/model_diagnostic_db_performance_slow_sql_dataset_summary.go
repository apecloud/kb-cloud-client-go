// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticDBPerformanceSlowSQLDatasetSummary Summary of the slow SQL sample set. REQUIRED when slowSqlState is
// source_present_with_data or source_present_no_data_in_window;
// FORBIDDEN otherwise. Cross-field consistency enforced by
// ValidateSlowSQLStateConsistency.
type DiagnosticDBPerformanceSlowSQLDatasetSummary struct {
	// Number of distinct slow SQL statements / fingerprints matched by
	// the threshold inside the time window. NOT a call count — repeated
	// executions of the same statement contribute to a single fingerprint
	// and increment topSlowSqls[].callCount instead. Must be > 0 when
	// slowSqlState=source_present_with_data and == 0 when
	// slowSqlState=source_present_no_data_in_window.
	//
	TotalCount int64 `json:"totalCount"`
	// Threshold rule used to pick slow SQL, such as "mean_time >= 200ms".
	Threshold string `json:"threshold"`
	// Top-N slow SQL fingerprints inside the window. Empty when
	// slowSqlState=source_present_no_data_in_window. May also be empty
	// when slowSqlState=source_present_with_data if redaction removed
	// every digest.
	//
	TopSlowSqls []DiagnosticDBPerformanceSlowSQLSample `json:"topSlowSqls,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticDBPerformanceSlowSQLDatasetSummary instantiates a new DiagnosticDBPerformanceSlowSQLDatasetSummary object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticDBPerformanceSlowSQLDatasetSummary(totalCount int64, threshold string) *DiagnosticDBPerformanceSlowSQLDatasetSummary {
	this := DiagnosticDBPerformanceSlowSQLDatasetSummary{}
	this.TotalCount = totalCount
	this.Threshold = threshold
	return &this
}

// NewDiagnosticDBPerformanceSlowSQLDatasetSummaryWithDefaults instantiates a new DiagnosticDBPerformanceSlowSQLDatasetSummary object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticDBPerformanceSlowSQLDatasetSummaryWithDefaults() *DiagnosticDBPerformanceSlowSQLDatasetSummary {
	this := DiagnosticDBPerformanceSlowSQLDatasetSummary{}
	return &this
}

// GetTotalCount returns the TotalCount field value.
func (o *DiagnosticDBPerformanceSlowSQLDatasetSummary) GetTotalCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceSlowSQLDatasetSummary) GetTotalCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value.
func (o *DiagnosticDBPerformanceSlowSQLDatasetSummary) SetTotalCount(v int64) {
	o.TotalCount = v
}

// GetThreshold returns the Threshold field value.
func (o *DiagnosticDBPerformanceSlowSQLDatasetSummary) GetThreshold() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceSlowSQLDatasetSummary) GetThresholdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Threshold, true
}

// SetThreshold sets field value.
func (o *DiagnosticDBPerformanceSlowSQLDatasetSummary) SetThreshold(v string) {
	o.Threshold = v
}

// GetTopSlowSqls returns the TopSlowSqls field value if set, zero value otherwise.
func (o *DiagnosticDBPerformanceSlowSQLDatasetSummary) GetTopSlowSqls() []DiagnosticDBPerformanceSlowSQLSample {
	if o == nil || o.TopSlowSqls == nil {
		var ret []DiagnosticDBPerformanceSlowSQLSample
		return ret
	}
	return o.TopSlowSqls
}

// GetTopSlowSqlsOk returns a tuple with the TopSlowSqls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceSlowSQLDatasetSummary) GetTopSlowSqlsOk() (*[]DiagnosticDBPerformanceSlowSQLSample, bool) {
	if o == nil || o.TopSlowSqls == nil {
		return nil, false
	}
	return &o.TopSlowSqls, true
}

// HasTopSlowSqls returns a boolean if a field has been set.
func (o *DiagnosticDBPerformanceSlowSQLDatasetSummary) HasTopSlowSqls() bool {
	return o != nil && o.TopSlowSqls != nil
}

// SetTopSlowSqls gets a reference to the given []DiagnosticDBPerformanceSlowSQLSample and assigns it to the TopSlowSqls field.
func (o *DiagnosticDBPerformanceSlowSQLDatasetSummary) SetTopSlowSqls(v []DiagnosticDBPerformanceSlowSQLSample) {
	o.TopSlowSqls = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticDBPerformanceSlowSQLDatasetSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["totalCount"] = o.TotalCount
	toSerialize["threshold"] = o.Threshold
	if o.TopSlowSqls != nil {
		toSerialize["topSlowSqls"] = o.TopSlowSqls
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticDBPerformanceSlowSQLDatasetSummary) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TotalCount  *int64                                 `json:"totalCount"`
		Threshold   *string                                `json:"threshold"`
		TopSlowSqls []DiagnosticDBPerformanceSlowSQLSample `json:"topSlowSqls,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.TotalCount == nil {
		return fmt.Errorf("required field totalCount missing")
	}
	if all.Threshold == nil {
		return fmt.Errorf("required field threshold missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"totalCount", "threshold", "topSlowSqls"})
	} else {
		return err
	}
	o.TotalCount = *all.TotalCount
	o.Threshold = *all.Threshold
	o.TopSlowSqls = all.TopSlowSqls

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
