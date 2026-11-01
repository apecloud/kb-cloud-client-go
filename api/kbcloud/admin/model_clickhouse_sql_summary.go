// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type ClickhouseSQLSummary struct {
	// Finished and failed queries in the selected range.
	QueryCount *int64 `json:"queryCount,omitempty"`
	// Queries that ended with ExceptionWhileProcessing.
	FailedCount *int64 `json:"failedCount,omitempty"`
	// Queries whose duration is at least 1000 ms.
	SlowQueryCount *int64 `json:"slowQueryCount,omitempty"`
	// Sum of query_duration_ms for all finished and failed queries.
	TotalDurationMs *float64 `json:"totalDurationMs,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClickhouseSQLSummary instantiates a new ClickhouseSQLSummary object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClickhouseSQLSummary() *ClickhouseSQLSummary {
	this := ClickhouseSQLSummary{}
	return &this
}

// NewClickhouseSQLSummaryWithDefaults instantiates a new ClickhouseSQLSummary object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClickhouseSQLSummaryWithDefaults() *ClickhouseSQLSummary {
	this := ClickhouseSQLSummary{}
	return &this
}

// GetQueryCount returns the QueryCount field value if set, zero value otherwise.
func (o *ClickhouseSQLSummary) GetQueryCount() int64 {
	if o == nil || o.QueryCount == nil {
		var ret int64
		return ret
	}
	return *o.QueryCount
}

// GetQueryCountOk returns a tuple with the QueryCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSQLSummary) GetQueryCountOk() (*int64, bool) {
	if o == nil || o.QueryCount == nil {
		return nil, false
	}
	return o.QueryCount, true
}

// HasQueryCount returns a boolean if a field has been set.
func (o *ClickhouseSQLSummary) HasQueryCount() bool {
	return o != nil && o.QueryCount != nil
}

// SetQueryCount gets a reference to the given int64 and assigns it to the QueryCount field.
func (o *ClickhouseSQLSummary) SetQueryCount(v int64) {
	o.QueryCount = &v
}

// GetFailedCount returns the FailedCount field value if set, zero value otherwise.
func (o *ClickhouseSQLSummary) GetFailedCount() int64 {
	if o == nil || o.FailedCount == nil {
		var ret int64
		return ret
	}
	return *o.FailedCount
}

// GetFailedCountOk returns a tuple with the FailedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSQLSummary) GetFailedCountOk() (*int64, bool) {
	if o == nil || o.FailedCount == nil {
		return nil, false
	}
	return o.FailedCount, true
}

// HasFailedCount returns a boolean if a field has been set.
func (o *ClickhouseSQLSummary) HasFailedCount() bool {
	return o != nil && o.FailedCount != nil
}

// SetFailedCount gets a reference to the given int64 and assigns it to the FailedCount field.
func (o *ClickhouseSQLSummary) SetFailedCount(v int64) {
	o.FailedCount = &v
}

// GetSlowQueryCount returns the SlowQueryCount field value if set, zero value otherwise.
func (o *ClickhouseSQLSummary) GetSlowQueryCount() int64 {
	if o == nil || o.SlowQueryCount == nil {
		var ret int64
		return ret
	}
	return *o.SlowQueryCount
}

// GetSlowQueryCountOk returns a tuple with the SlowQueryCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSQLSummary) GetSlowQueryCountOk() (*int64, bool) {
	if o == nil || o.SlowQueryCount == nil {
		return nil, false
	}
	return o.SlowQueryCount, true
}

// HasSlowQueryCount returns a boolean if a field has been set.
func (o *ClickhouseSQLSummary) HasSlowQueryCount() bool {
	return o != nil && o.SlowQueryCount != nil
}

// SetSlowQueryCount gets a reference to the given int64 and assigns it to the SlowQueryCount field.
func (o *ClickhouseSQLSummary) SetSlowQueryCount(v int64) {
	o.SlowQueryCount = &v
}

// GetTotalDurationMs returns the TotalDurationMs field value if set, zero value otherwise.
func (o *ClickhouseSQLSummary) GetTotalDurationMs() float64 {
	if o == nil || o.TotalDurationMs == nil {
		var ret float64
		return ret
	}
	return *o.TotalDurationMs
}

// GetTotalDurationMsOk returns a tuple with the TotalDurationMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSQLSummary) GetTotalDurationMsOk() (*float64, bool) {
	if o == nil || o.TotalDurationMs == nil {
		return nil, false
	}
	return o.TotalDurationMs, true
}

// HasTotalDurationMs returns a boolean if a field has been set.
func (o *ClickhouseSQLSummary) HasTotalDurationMs() bool {
	return o != nil && o.TotalDurationMs != nil
}

// SetTotalDurationMs gets a reference to the given float64 and assigns it to the TotalDurationMs field.
func (o *ClickhouseSQLSummary) SetTotalDurationMs(v float64) {
	o.TotalDurationMs = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClickhouseSQLSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.QueryCount != nil {
		toSerialize["queryCount"] = o.QueryCount
	}
	if o.FailedCount != nil {
		toSerialize["failedCount"] = o.FailedCount
	}
	if o.SlowQueryCount != nil {
		toSerialize["slowQueryCount"] = o.SlowQueryCount
	}
	if o.TotalDurationMs != nil {
		toSerialize["totalDurationMs"] = o.TotalDurationMs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClickhouseSQLSummary) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		QueryCount      *int64   `json:"queryCount,omitempty"`
		FailedCount     *int64   `json:"failedCount,omitempty"`
		SlowQueryCount  *int64   `json:"slowQueryCount,omitempty"`
		TotalDurationMs *float64 `json:"totalDurationMs,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"queryCount", "failedCount", "slowQueryCount", "totalDurationMs"})
	} else {
		return err
	}
	o.QueryCount = all.QueryCount
	o.FailedCount = all.FailedCount
	o.SlowQueryCount = all.SlowQueryCount
	o.TotalDurationMs = all.TotalDurationMs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
