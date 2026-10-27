// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type ClickhouseFailedQuery struct {
	// query_log event time in RFC3339 UTC.
	EventTime *string `json:"eventTime,omitempty"`
	// ClickHouse query_id.
	QueryId *string `json:"queryId,omitempty"`
	// User that ran the query.
	User *string `json:"user,omitempty"`
	// Current database at execution time.
	Database *string `json:"database,omitempty"`
	// Full query text.
	Query *string `json:"query,omitempty"`
	// ClickHouse exception code.
	ExceptionCode *int64 `json:"exceptionCode,omitempty"`
	// Best-effort category such as memory_limit_exceeded, timeout, too_many_parts, readonly, unknown_table or syntax_error.
	ExceptionCategory *string `json:"exceptionCategory,omitempty"`
	// Full exception text.
	Exception *string `json:"exception,omitempty"`
	// query_duration_ms before the failure.
	DurationMs *float64 `json:"durationMs,omitempty"`
	// Memory usage in bytes.
	MemoryUsageBytes *int64 `json:"memoryUsageBytes,omitempty"`
	// ClickHouse server host that ran the query.
	HostName *string `json:"hostName,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClickhouseFailedQuery instantiates a new ClickhouseFailedQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClickhouseFailedQuery() *ClickhouseFailedQuery {
	this := ClickhouseFailedQuery{}
	return &this
}

// NewClickhouseFailedQueryWithDefaults instantiates a new ClickhouseFailedQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClickhouseFailedQueryWithDefaults() *ClickhouseFailedQuery {
	this := ClickhouseFailedQuery{}
	return &this
}

// GetEventTime returns the EventTime field value if set, zero value otherwise.
func (o *ClickhouseFailedQuery) GetEventTime() string {
	if o == nil || o.EventTime == nil {
		var ret string
		return ret
	}
	return *o.EventTime
}

// GetEventTimeOk returns a tuple with the EventTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseFailedQuery) GetEventTimeOk() (*string, bool) {
	if o == nil || o.EventTime == nil {
		return nil, false
	}
	return o.EventTime, true
}

// HasEventTime returns a boolean if a field has been set.
func (o *ClickhouseFailedQuery) HasEventTime() bool {
	return o != nil && o.EventTime != nil
}

// SetEventTime gets a reference to the given string and assigns it to the EventTime field.
func (o *ClickhouseFailedQuery) SetEventTime(v string) {
	o.EventTime = &v
}

// GetQueryId returns the QueryId field value if set, zero value otherwise.
func (o *ClickhouseFailedQuery) GetQueryId() string {
	if o == nil || o.QueryId == nil {
		var ret string
		return ret
	}
	return *o.QueryId
}

// GetQueryIdOk returns a tuple with the QueryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseFailedQuery) GetQueryIdOk() (*string, bool) {
	if o == nil || o.QueryId == nil {
		return nil, false
	}
	return o.QueryId, true
}

// HasQueryId returns a boolean if a field has been set.
func (o *ClickhouseFailedQuery) HasQueryId() bool {
	return o != nil && o.QueryId != nil
}

// SetQueryId gets a reference to the given string and assigns it to the QueryId field.
func (o *ClickhouseFailedQuery) SetQueryId(v string) {
	o.QueryId = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *ClickhouseFailedQuery) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseFailedQuery) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *ClickhouseFailedQuery) HasUser() bool {
	return o != nil && o.User != nil
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *ClickhouseFailedQuery) SetUser(v string) {
	o.User = &v
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *ClickhouseFailedQuery) GetDatabase() string {
	if o == nil || o.Database == nil {
		var ret string
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseFailedQuery) GetDatabaseOk() (*string, bool) {
	if o == nil || o.Database == nil {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *ClickhouseFailedQuery) HasDatabase() bool {
	return o != nil && o.Database != nil
}

// SetDatabase gets a reference to the given string and assigns it to the Database field.
func (o *ClickhouseFailedQuery) SetDatabase(v string) {
	o.Database = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *ClickhouseFailedQuery) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseFailedQuery) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *ClickhouseFailedQuery) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *ClickhouseFailedQuery) SetQuery(v string) {
	o.Query = &v
}

// GetExceptionCode returns the ExceptionCode field value if set, zero value otherwise.
func (o *ClickhouseFailedQuery) GetExceptionCode() int64 {
	if o == nil || o.ExceptionCode == nil {
		var ret int64
		return ret
	}
	return *o.ExceptionCode
}

// GetExceptionCodeOk returns a tuple with the ExceptionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseFailedQuery) GetExceptionCodeOk() (*int64, bool) {
	if o == nil || o.ExceptionCode == nil {
		return nil, false
	}
	return o.ExceptionCode, true
}

// HasExceptionCode returns a boolean if a field has been set.
func (o *ClickhouseFailedQuery) HasExceptionCode() bool {
	return o != nil && o.ExceptionCode != nil
}

// SetExceptionCode gets a reference to the given int64 and assigns it to the ExceptionCode field.
func (o *ClickhouseFailedQuery) SetExceptionCode(v int64) {
	o.ExceptionCode = &v
}

// GetExceptionCategory returns the ExceptionCategory field value if set, zero value otherwise.
func (o *ClickhouseFailedQuery) GetExceptionCategory() string {
	if o == nil || o.ExceptionCategory == nil {
		var ret string
		return ret
	}
	return *o.ExceptionCategory
}

// GetExceptionCategoryOk returns a tuple with the ExceptionCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseFailedQuery) GetExceptionCategoryOk() (*string, bool) {
	if o == nil || o.ExceptionCategory == nil {
		return nil, false
	}
	return o.ExceptionCategory, true
}

// HasExceptionCategory returns a boolean if a field has been set.
func (o *ClickhouseFailedQuery) HasExceptionCategory() bool {
	return o != nil && o.ExceptionCategory != nil
}

// SetExceptionCategory gets a reference to the given string and assigns it to the ExceptionCategory field.
func (o *ClickhouseFailedQuery) SetExceptionCategory(v string) {
	o.ExceptionCategory = &v
}

// GetException returns the Exception field value if set, zero value otherwise.
func (o *ClickhouseFailedQuery) GetException() string {
	if o == nil || o.Exception == nil {
		var ret string
		return ret
	}
	return *o.Exception
}

// GetExceptionOk returns a tuple with the Exception field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseFailedQuery) GetExceptionOk() (*string, bool) {
	if o == nil || o.Exception == nil {
		return nil, false
	}
	return o.Exception, true
}

// HasException returns a boolean if a field has been set.
func (o *ClickhouseFailedQuery) HasException() bool {
	return o != nil && o.Exception != nil
}

// SetException gets a reference to the given string and assigns it to the Exception field.
func (o *ClickhouseFailedQuery) SetException(v string) {
	o.Exception = &v
}

// GetDurationMs returns the DurationMs field value if set, zero value otherwise.
func (o *ClickhouseFailedQuery) GetDurationMs() float64 {
	if o == nil || o.DurationMs == nil {
		var ret float64
		return ret
	}
	return *o.DurationMs
}

// GetDurationMsOk returns a tuple with the DurationMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseFailedQuery) GetDurationMsOk() (*float64, bool) {
	if o == nil || o.DurationMs == nil {
		return nil, false
	}
	return o.DurationMs, true
}

// HasDurationMs returns a boolean if a field has been set.
func (o *ClickhouseFailedQuery) HasDurationMs() bool {
	return o != nil && o.DurationMs != nil
}

// SetDurationMs gets a reference to the given float64 and assigns it to the DurationMs field.
func (o *ClickhouseFailedQuery) SetDurationMs(v float64) {
	o.DurationMs = &v
}

// GetMemoryUsageBytes returns the MemoryUsageBytes field value if set, zero value otherwise.
func (o *ClickhouseFailedQuery) GetMemoryUsageBytes() int64 {
	if o == nil || o.MemoryUsageBytes == nil {
		var ret int64
		return ret
	}
	return *o.MemoryUsageBytes
}

// GetMemoryUsageBytesOk returns a tuple with the MemoryUsageBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseFailedQuery) GetMemoryUsageBytesOk() (*int64, bool) {
	if o == nil || o.MemoryUsageBytes == nil {
		return nil, false
	}
	return o.MemoryUsageBytes, true
}

// HasMemoryUsageBytes returns a boolean if a field has been set.
func (o *ClickhouseFailedQuery) HasMemoryUsageBytes() bool {
	return o != nil && o.MemoryUsageBytes != nil
}

// SetMemoryUsageBytes gets a reference to the given int64 and assigns it to the MemoryUsageBytes field.
func (o *ClickhouseFailedQuery) SetMemoryUsageBytes(v int64) {
	o.MemoryUsageBytes = &v
}

// GetHostName returns the HostName field value if set, zero value otherwise.
func (o *ClickhouseFailedQuery) GetHostName() string {
	if o == nil || o.HostName == nil {
		var ret string
		return ret
	}
	return *o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseFailedQuery) GetHostNameOk() (*string, bool) {
	if o == nil || o.HostName == nil {
		return nil, false
	}
	return o.HostName, true
}

// HasHostName returns a boolean if a field has been set.
func (o *ClickhouseFailedQuery) HasHostName() bool {
	return o != nil && o.HostName != nil
}

// SetHostName gets a reference to the given string and assigns it to the HostName field.
func (o *ClickhouseFailedQuery) SetHostName(v string) {
	o.HostName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClickhouseFailedQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.EventTime != nil {
		toSerialize["eventTime"] = o.EventTime
	}
	if o.QueryId != nil {
		toSerialize["queryId"] = o.QueryId
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Database != nil {
		toSerialize["database"] = o.Database
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.ExceptionCode != nil {
		toSerialize["exceptionCode"] = o.ExceptionCode
	}
	if o.ExceptionCategory != nil {
		toSerialize["exceptionCategory"] = o.ExceptionCategory
	}
	if o.Exception != nil {
		toSerialize["exception"] = o.Exception
	}
	if o.DurationMs != nil {
		toSerialize["durationMs"] = o.DurationMs
	}
	if o.MemoryUsageBytes != nil {
		toSerialize["memoryUsageBytes"] = o.MemoryUsageBytes
	}
	if o.HostName != nil {
		toSerialize["hostName"] = o.HostName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClickhouseFailedQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EventTime         *string  `json:"eventTime,omitempty"`
		QueryId           *string  `json:"queryId,omitempty"`
		User              *string  `json:"user,omitempty"`
		Database          *string  `json:"database,omitempty"`
		Query             *string  `json:"query,omitempty"`
		ExceptionCode     *int64   `json:"exceptionCode,omitempty"`
		ExceptionCategory *string  `json:"exceptionCategory,omitempty"`
		Exception         *string  `json:"exception,omitempty"`
		DurationMs        *float64 `json:"durationMs,omitempty"`
		MemoryUsageBytes  *int64   `json:"memoryUsageBytes,omitempty"`
		HostName          *string  `json:"hostName,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"eventTime", "queryId", "user", "database", "query", "exceptionCode", "exceptionCategory", "exception", "durationMs", "memoryUsageBytes", "hostName"})
	} else {
		return err
	}
	o.EventTime = all.EventTime
	o.QueryId = all.QueryId
	o.User = all.User
	o.Database = all.Database
	o.Query = all.Query
	o.ExceptionCode = all.ExceptionCode
	o.ExceptionCategory = all.ExceptionCategory
	o.Exception = all.Exception
	o.DurationMs = all.DurationMs
	o.MemoryUsageBytes = all.MemoryUsageBytes
	o.HostName = all.HostName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
