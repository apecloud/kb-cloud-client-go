// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type ClickhouseQueryPattern struct {
	// normalized_query_hash as a string.
	NormalizedQueryHash *string `json:"normalizedQueryHash,omitempty"`
	// One representative query text for the pattern, with string and numeric literals and comments masked as ?.
	SampleQuery *string `json:"sampleQuery,omitempty"`
	// Number of executions of the pattern.
	Calls *int64 `json:"calls,omitempty"`
	// Sum of query_duration_ms across all executions.
	TotalDurationMs *float64 `json:"totalDurationMs,omitempty"`
	// Average query_duration_ms.
	MeanDurationMs *float64 `json:"meanDurationMs,omitempty"`
	// Maximum query_duration_ms.
	MaxDurationMs *float64 `json:"maxDurationMs,omitempty"`
	// Sum of read_rows.
	ReadRows *int64 `json:"readRows,omitempty"`
	// Sum of memory_usage in bytes.
	MemoryUsageBytes *int64 `json:"memoryUsageBytes,omitempty"`
	// Executions that ended with ExceptionWhileProcessing.
	FailureCount *int64 `json:"failureCount,omitempty"`
	// Executions whose duration is at least 1000 ms.
	SlowCallCount *int64 `json:"slowCallCount,omitempty"`
	// First event_time in RFC3339 UTC.
	FirstSeen *string `json:"firstSeen,omitempty"`
	// Last event_time in RFC3339 UTC.
	LastSeen *string `json:"lastSeen,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClickhouseQueryPattern instantiates a new ClickhouseQueryPattern object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClickhouseQueryPattern() *ClickhouseQueryPattern {
	this := ClickhouseQueryPattern{}
	return &this
}

// NewClickhouseQueryPatternWithDefaults instantiates a new ClickhouseQueryPattern object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClickhouseQueryPatternWithDefaults() *ClickhouseQueryPattern {
	this := ClickhouseQueryPattern{}
	return &this
}

// GetNormalizedQueryHash returns the NormalizedQueryHash field value if set, zero value otherwise.
func (o *ClickhouseQueryPattern) GetNormalizedQueryHash() string {
	if o == nil || o.NormalizedQueryHash == nil {
		var ret string
		return ret
	}
	return *o.NormalizedQueryHash
}

// GetNormalizedQueryHashOk returns a tuple with the NormalizedQueryHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseQueryPattern) GetNormalizedQueryHashOk() (*string, bool) {
	if o == nil || o.NormalizedQueryHash == nil {
		return nil, false
	}
	return o.NormalizedQueryHash, true
}

// HasNormalizedQueryHash returns a boolean if a field has been set.
func (o *ClickhouseQueryPattern) HasNormalizedQueryHash() bool {
	return o != nil && o.NormalizedQueryHash != nil
}

// SetNormalizedQueryHash gets a reference to the given string and assigns it to the NormalizedQueryHash field.
func (o *ClickhouseQueryPattern) SetNormalizedQueryHash(v string) {
	o.NormalizedQueryHash = &v
}

// GetSampleQuery returns the SampleQuery field value if set, zero value otherwise.
func (o *ClickhouseQueryPattern) GetSampleQuery() string {
	if o == nil || o.SampleQuery == nil {
		var ret string
		return ret
	}
	return *o.SampleQuery
}

// GetSampleQueryOk returns a tuple with the SampleQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseQueryPattern) GetSampleQueryOk() (*string, bool) {
	if o == nil || o.SampleQuery == nil {
		return nil, false
	}
	return o.SampleQuery, true
}

// HasSampleQuery returns a boolean if a field has been set.
func (o *ClickhouseQueryPattern) HasSampleQuery() bool {
	return o != nil && o.SampleQuery != nil
}

// SetSampleQuery gets a reference to the given string and assigns it to the SampleQuery field.
func (o *ClickhouseQueryPattern) SetSampleQuery(v string) {
	o.SampleQuery = &v
}

// GetCalls returns the Calls field value if set, zero value otherwise.
func (o *ClickhouseQueryPattern) GetCalls() int64 {
	if o == nil || o.Calls == nil {
		var ret int64
		return ret
	}
	return *o.Calls
}

// GetCallsOk returns a tuple with the Calls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseQueryPattern) GetCallsOk() (*int64, bool) {
	if o == nil || o.Calls == nil {
		return nil, false
	}
	return o.Calls, true
}

// HasCalls returns a boolean if a field has been set.
func (o *ClickhouseQueryPattern) HasCalls() bool {
	return o != nil && o.Calls != nil
}

// SetCalls gets a reference to the given int64 and assigns it to the Calls field.
func (o *ClickhouseQueryPattern) SetCalls(v int64) {
	o.Calls = &v
}

// GetTotalDurationMs returns the TotalDurationMs field value if set, zero value otherwise.
func (o *ClickhouseQueryPattern) GetTotalDurationMs() float64 {
	if o == nil || o.TotalDurationMs == nil {
		var ret float64
		return ret
	}
	return *o.TotalDurationMs
}

// GetTotalDurationMsOk returns a tuple with the TotalDurationMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseQueryPattern) GetTotalDurationMsOk() (*float64, bool) {
	if o == nil || o.TotalDurationMs == nil {
		return nil, false
	}
	return o.TotalDurationMs, true
}

// HasTotalDurationMs returns a boolean if a field has been set.
func (o *ClickhouseQueryPattern) HasTotalDurationMs() bool {
	return o != nil && o.TotalDurationMs != nil
}

// SetTotalDurationMs gets a reference to the given float64 and assigns it to the TotalDurationMs field.
func (o *ClickhouseQueryPattern) SetTotalDurationMs(v float64) {
	o.TotalDurationMs = &v
}

// GetMeanDurationMs returns the MeanDurationMs field value if set, zero value otherwise.
func (o *ClickhouseQueryPattern) GetMeanDurationMs() float64 {
	if o == nil || o.MeanDurationMs == nil {
		var ret float64
		return ret
	}
	return *o.MeanDurationMs
}

// GetMeanDurationMsOk returns a tuple with the MeanDurationMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseQueryPattern) GetMeanDurationMsOk() (*float64, bool) {
	if o == nil || o.MeanDurationMs == nil {
		return nil, false
	}
	return o.MeanDurationMs, true
}

// HasMeanDurationMs returns a boolean if a field has been set.
func (o *ClickhouseQueryPattern) HasMeanDurationMs() bool {
	return o != nil && o.MeanDurationMs != nil
}

// SetMeanDurationMs gets a reference to the given float64 and assigns it to the MeanDurationMs field.
func (o *ClickhouseQueryPattern) SetMeanDurationMs(v float64) {
	o.MeanDurationMs = &v
}

// GetMaxDurationMs returns the MaxDurationMs field value if set, zero value otherwise.
func (o *ClickhouseQueryPattern) GetMaxDurationMs() float64 {
	if o == nil || o.MaxDurationMs == nil {
		var ret float64
		return ret
	}
	return *o.MaxDurationMs
}

// GetMaxDurationMsOk returns a tuple with the MaxDurationMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseQueryPattern) GetMaxDurationMsOk() (*float64, bool) {
	if o == nil || o.MaxDurationMs == nil {
		return nil, false
	}
	return o.MaxDurationMs, true
}

// HasMaxDurationMs returns a boolean if a field has been set.
func (o *ClickhouseQueryPattern) HasMaxDurationMs() bool {
	return o != nil && o.MaxDurationMs != nil
}

// SetMaxDurationMs gets a reference to the given float64 and assigns it to the MaxDurationMs field.
func (o *ClickhouseQueryPattern) SetMaxDurationMs(v float64) {
	o.MaxDurationMs = &v
}

// GetReadRows returns the ReadRows field value if set, zero value otherwise.
func (o *ClickhouseQueryPattern) GetReadRows() int64 {
	if o == nil || o.ReadRows == nil {
		var ret int64
		return ret
	}
	return *o.ReadRows
}

// GetReadRowsOk returns a tuple with the ReadRows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseQueryPattern) GetReadRowsOk() (*int64, bool) {
	if o == nil || o.ReadRows == nil {
		return nil, false
	}
	return o.ReadRows, true
}

// HasReadRows returns a boolean if a field has been set.
func (o *ClickhouseQueryPattern) HasReadRows() bool {
	return o != nil && o.ReadRows != nil
}

// SetReadRows gets a reference to the given int64 and assigns it to the ReadRows field.
func (o *ClickhouseQueryPattern) SetReadRows(v int64) {
	o.ReadRows = &v
}

// GetMemoryUsageBytes returns the MemoryUsageBytes field value if set, zero value otherwise.
func (o *ClickhouseQueryPattern) GetMemoryUsageBytes() int64 {
	if o == nil || o.MemoryUsageBytes == nil {
		var ret int64
		return ret
	}
	return *o.MemoryUsageBytes
}

// GetMemoryUsageBytesOk returns a tuple with the MemoryUsageBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseQueryPattern) GetMemoryUsageBytesOk() (*int64, bool) {
	if o == nil || o.MemoryUsageBytes == nil {
		return nil, false
	}
	return o.MemoryUsageBytes, true
}

// HasMemoryUsageBytes returns a boolean if a field has been set.
func (o *ClickhouseQueryPattern) HasMemoryUsageBytes() bool {
	return o != nil && o.MemoryUsageBytes != nil
}

// SetMemoryUsageBytes gets a reference to the given int64 and assigns it to the MemoryUsageBytes field.
func (o *ClickhouseQueryPattern) SetMemoryUsageBytes(v int64) {
	o.MemoryUsageBytes = &v
}

// GetFailureCount returns the FailureCount field value if set, zero value otherwise.
func (o *ClickhouseQueryPattern) GetFailureCount() int64 {
	if o == nil || o.FailureCount == nil {
		var ret int64
		return ret
	}
	return *o.FailureCount
}

// GetFailureCountOk returns a tuple with the FailureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseQueryPattern) GetFailureCountOk() (*int64, bool) {
	if o == nil || o.FailureCount == nil {
		return nil, false
	}
	return o.FailureCount, true
}

// HasFailureCount returns a boolean if a field has been set.
func (o *ClickhouseQueryPattern) HasFailureCount() bool {
	return o != nil && o.FailureCount != nil
}

// SetFailureCount gets a reference to the given int64 and assigns it to the FailureCount field.
func (o *ClickhouseQueryPattern) SetFailureCount(v int64) {
	o.FailureCount = &v
}

// GetSlowCallCount returns the SlowCallCount field value if set, zero value otherwise.
func (o *ClickhouseQueryPattern) GetSlowCallCount() int64 {
	if o == nil || o.SlowCallCount == nil {
		var ret int64
		return ret
	}
	return *o.SlowCallCount
}

// GetSlowCallCountOk returns a tuple with the SlowCallCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseQueryPattern) GetSlowCallCountOk() (*int64, bool) {
	if o == nil || o.SlowCallCount == nil {
		return nil, false
	}
	return o.SlowCallCount, true
}

// HasSlowCallCount returns a boolean if a field has been set.
func (o *ClickhouseQueryPattern) HasSlowCallCount() bool {
	return o != nil && o.SlowCallCount != nil
}

// SetSlowCallCount gets a reference to the given int64 and assigns it to the SlowCallCount field.
func (o *ClickhouseQueryPattern) SetSlowCallCount(v int64) {
	o.SlowCallCount = &v
}

// GetFirstSeen returns the FirstSeen field value if set, zero value otherwise.
func (o *ClickhouseQueryPattern) GetFirstSeen() string {
	if o == nil || o.FirstSeen == nil {
		var ret string
		return ret
	}
	return *o.FirstSeen
}

// GetFirstSeenOk returns a tuple with the FirstSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseQueryPattern) GetFirstSeenOk() (*string, bool) {
	if o == nil || o.FirstSeen == nil {
		return nil, false
	}
	return o.FirstSeen, true
}

// HasFirstSeen returns a boolean if a field has been set.
func (o *ClickhouseQueryPattern) HasFirstSeen() bool {
	return o != nil && o.FirstSeen != nil
}

// SetFirstSeen gets a reference to the given string and assigns it to the FirstSeen field.
func (o *ClickhouseQueryPattern) SetFirstSeen(v string) {
	o.FirstSeen = &v
}

// GetLastSeen returns the LastSeen field value if set, zero value otherwise.
func (o *ClickhouseQueryPattern) GetLastSeen() string {
	if o == nil || o.LastSeen == nil {
		var ret string
		return ret
	}
	return *o.LastSeen
}

// GetLastSeenOk returns a tuple with the LastSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseQueryPattern) GetLastSeenOk() (*string, bool) {
	if o == nil || o.LastSeen == nil {
		return nil, false
	}
	return o.LastSeen, true
}

// HasLastSeen returns a boolean if a field has been set.
func (o *ClickhouseQueryPattern) HasLastSeen() bool {
	return o != nil && o.LastSeen != nil
}

// SetLastSeen gets a reference to the given string and assigns it to the LastSeen field.
func (o *ClickhouseQueryPattern) SetLastSeen(v string) {
	o.LastSeen = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClickhouseQueryPattern) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.NormalizedQueryHash != nil {
		toSerialize["normalizedQueryHash"] = o.NormalizedQueryHash
	}
	if o.SampleQuery != nil {
		toSerialize["sampleQuery"] = o.SampleQuery
	}
	if o.Calls != nil {
		toSerialize["calls"] = o.Calls
	}
	if o.TotalDurationMs != nil {
		toSerialize["totalDurationMs"] = o.TotalDurationMs
	}
	if o.MeanDurationMs != nil {
		toSerialize["meanDurationMs"] = o.MeanDurationMs
	}
	if o.MaxDurationMs != nil {
		toSerialize["maxDurationMs"] = o.MaxDurationMs
	}
	if o.ReadRows != nil {
		toSerialize["readRows"] = o.ReadRows
	}
	if o.MemoryUsageBytes != nil {
		toSerialize["memoryUsageBytes"] = o.MemoryUsageBytes
	}
	if o.FailureCount != nil {
		toSerialize["failureCount"] = o.FailureCount
	}
	if o.SlowCallCount != nil {
		toSerialize["slowCallCount"] = o.SlowCallCount
	}
	if o.FirstSeen != nil {
		toSerialize["firstSeen"] = o.FirstSeen
	}
	if o.LastSeen != nil {
		toSerialize["lastSeen"] = o.LastSeen
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClickhouseQueryPattern) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NormalizedQueryHash *string  `json:"normalizedQueryHash,omitempty"`
		SampleQuery         *string  `json:"sampleQuery,omitempty"`
		Calls               *int64   `json:"calls,omitempty"`
		TotalDurationMs     *float64 `json:"totalDurationMs,omitempty"`
		MeanDurationMs      *float64 `json:"meanDurationMs,omitempty"`
		MaxDurationMs       *float64 `json:"maxDurationMs,omitempty"`
		ReadRows            *int64   `json:"readRows,omitempty"`
		MemoryUsageBytes    *int64   `json:"memoryUsageBytes,omitempty"`
		FailureCount        *int64   `json:"failureCount,omitempty"`
		SlowCallCount       *int64   `json:"slowCallCount,omitempty"`
		FirstSeen           *string  `json:"firstSeen,omitempty"`
		LastSeen            *string  `json:"lastSeen,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"normalizedQueryHash", "sampleQuery", "calls", "totalDurationMs", "meanDurationMs", "maxDurationMs", "readRows", "memoryUsageBytes", "failureCount", "slowCallCount", "firstSeen", "lastSeen"})
	} else {
		return err
	}
	o.NormalizedQueryHash = all.NormalizedQueryHash
	o.SampleQuery = all.SampleQuery
	o.Calls = all.Calls
	o.TotalDurationMs = all.TotalDurationMs
	o.MeanDurationMs = all.MeanDurationMs
	o.MaxDurationMs = all.MaxDurationMs
	o.ReadRows = all.ReadRows
	o.MemoryUsageBytes = all.MemoryUsageBytes
	o.FailureCount = all.FailureCount
	o.SlowCallCount = all.SlowCallCount
	o.FirstSeen = all.FirstSeen
	o.LastSeen = all.LastSeen

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
