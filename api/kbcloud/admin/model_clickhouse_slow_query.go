// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type ClickhouseSlowQuery struct {
	// ClickHouse query_id.
	QueryId *string `json:"queryId,omitempty"`
	// Initial query_id for a distributed query.
	InitialQueryId *string `json:"initialQueryId,omitempty"`
	// query_log type (QueryFinish or ExceptionWhileProcessing).
	Type *string `json:"type,omitempty"`
	// User that ran the query.
	User *string `json:"user,omitempty"`
	// Current database at execution time.
	Database *string `json:"database,omitempty"`
	// Query kind reported by ClickHouse.
	QueryKind *string `json:"queryKind,omitempty"`
	// Full query text.
	Query *string `json:"query,omitempty"`
	// normalized_query_hash as a string to avoid precision loss in clients.
	NormalizedQueryHash *string `json:"normalizedQueryHash,omitempty"`
	// query_duration_ms.
	DurationMs *float64 `json:"durationMs,omitempty"`
	// Rows read by the query.
	ReadRows *int64 `json:"readRows,omitempty"`
	// Bytes read by the query.
	ReadBytes *int64 `json:"readBytes,omitempty"`
	// Memory usage in bytes.
	MemoryUsageBytes *int64 `json:"memoryUsageBytes,omitempty"`
	// Rows returned by the query.
	ResultRows *int64 `json:"resultRows,omitempty"`
	// Bytes returned by the query.
	ResultBytes *int64 `json:"resultBytes,omitempty"`
	// ClickHouse exception code, zero for successful queries.
	ExceptionCode *int64 `json:"exceptionCode,omitempty"`
	// Exception text for failed queries.
	Exception *string `json:"exception,omitempty"`
	// query_log event time in RFC3339 UTC.
	EventTime *string `json:"eventTime,omitempty"`
	// ClickHouse server host that ran the query.
	HostName *string `json:"hostName,omitempty"`
	// Client name reported by ClickHouse.
	ClientName *string `json:"clientName,omitempty"`
	// Whether this is the initial query rather than a remote sub-query.
	IsInitialQuery *bool `json:"isInitialQuery,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClickhouseSlowQuery instantiates a new ClickhouseSlowQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClickhouseSlowQuery() *ClickhouseSlowQuery {
	this := ClickhouseSlowQuery{}
	return &this
}

// NewClickhouseSlowQueryWithDefaults instantiates a new ClickhouseSlowQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClickhouseSlowQueryWithDefaults() *ClickhouseSlowQuery {
	this := ClickhouseSlowQuery{}
	return &this
}

// GetQueryId returns the QueryId field value if set, zero value otherwise.
func (o *ClickhouseSlowQuery) GetQueryId() string {
	if o == nil || o.QueryId == nil {
		var ret string
		return ret
	}
	return *o.QueryId
}

// GetQueryIdOk returns a tuple with the QueryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSlowQuery) GetQueryIdOk() (*string, bool) {
	if o == nil || o.QueryId == nil {
		return nil, false
	}
	return o.QueryId, true
}

// HasQueryId returns a boolean if a field has been set.
func (o *ClickhouseSlowQuery) HasQueryId() bool {
	return o != nil && o.QueryId != nil
}

// SetQueryId gets a reference to the given string and assigns it to the QueryId field.
func (o *ClickhouseSlowQuery) SetQueryId(v string) {
	o.QueryId = &v
}

// GetInitialQueryId returns the InitialQueryId field value if set, zero value otherwise.
func (o *ClickhouseSlowQuery) GetInitialQueryId() string {
	if o == nil || o.InitialQueryId == nil {
		var ret string
		return ret
	}
	return *o.InitialQueryId
}

// GetInitialQueryIdOk returns a tuple with the InitialQueryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSlowQuery) GetInitialQueryIdOk() (*string, bool) {
	if o == nil || o.InitialQueryId == nil {
		return nil, false
	}
	return o.InitialQueryId, true
}

// HasInitialQueryId returns a boolean if a field has been set.
func (o *ClickhouseSlowQuery) HasInitialQueryId() bool {
	return o != nil && o.InitialQueryId != nil
}

// SetInitialQueryId gets a reference to the given string and assigns it to the InitialQueryId field.
func (o *ClickhouseSlowQuery) SetInitialQueryId(v string) {
	o.InitialQueryId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ClickhouseSlowQuery) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSlowQuery) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ClickhouseSlowQuery) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ClickhouseSlowQuery) SetType(v string) {
	o.Type = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *ClickhouseSlowQuery) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSlowQuery) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *ClickhouseSlowQuery) HasUser() bool {
	return o != nil && o.User != nil
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *ClickhouseSlowQuery) SetUser(v string) {
	o.User = &v
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *ClickhouseSlowQuery) GetDatabase() string {
	if o == nil || o.Database == nil {
		var ret string
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSlowQuery) GetDatabaseOk() (*string, bool) {
	if o == nil || o.Database == nil {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *ClickhouseSlowQuery) HasDatabase() bool {
	return o != nil && o.Database != nil
}

// SetDatabase gets a reference to the given string and assigns it to the Database field.
func (o *ClickhouseSlowQuery) SetDatabase(v string) {
	o.Database = &v
}

// GetQueryKind returns the QueryKind field value if set, zero value otherwise.
func (o *ClickhouseSlowQuery) GetQueryKind() string {
	if o == nil || o.QueryKind == nil {
		var ret string
		return ret
	}
	return *o.QueryKind
}

// GetQueryKindOk returns a tuple with the QueryKind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSlowQuery) GetQueryKindOk() (*string, bool) {
	if o == nil || o.QueryKind == nil {
		return nil, false
	}
	return o.QueryKind, true
}

// HasQueryKind returns a boolean if a field has been set.
func (o *ClickhouseSlowQuery) HasQueryKind() bool {
	return o != nil && o.QueryKind != nil
}

// SetQueryKind gets a reference to the given string and assigns it to the QueryKind field.
func (o *ClickhouseSlowQuery) SetQueryKind(v string) {
	o.QueryKind = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *ClickhouseSlowQuery) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSlowQuery) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *ClickhouseSlowQuery) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *ClickhouseSlowQuery) SetQuery(v string) {
	o.Query = &v
}

// GetNormalizedQueryHash returns the NormalizedQueryHash field value if set, zero value otherwise.
func (o *ClickhouseSlowQuery) GetNormalizedQueryHash() string {
	if o == nil || o.NormalizedQueryHash == nil {
		var ret string
		return ret
	}
	return *o.NormalizedQueryHash
}

// GetNormalizedQueryHashOk returns a tuple with the NormalizedQueryHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSlowQuery) GetNormalizedQueryHashOk() (*string, bool) {
	if o == nil || o.NormalizedQueryHash == nil {
		return nil, false
	}
	return o.NormalizedQueryHash, true
}

// HasNormalizedQueryHash returns a boolean if a field has been set.
func (o *ClickhouseSlowQuery) HasNormalizedQueryHash() bool {
	return o != nil && o.NormalizedQueryHash != nil
}

// SetNormalizedQueryHash gets a reference to the given string and assigns it to the NormalizedQueryHash field.
func (o *ClickhouseSlowQuery) SetNormalizedQueryHash(v string) {
	o.NormalizedQueryHash = &v
}

// GetDurationMs returns the DurationMs field value if set, zero value otherwise.
func (o *ClickhouseSlowQuery) GetDurationMs() float64 {
	if o == nil || o.DurationMs == nil {
		var ret float64
		return ret
	}
	return *o.DurationMs
}

// GetDurationMsOk returns a tuple with the DurationMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSlowQuery) GetDurationMsOk() (*float64, bool) {
	if o == nil || o.DurationMs == nil {
		return nil, false
	}
	return o.DurationMs, true
}

// HasDurationMs returns a boolean if a field has been set.
func (o *ClickhouseSlowQuery) HasDurationMs() bool {
	return o != nil && o.DurationMs != nil
}

// SetDurationMs gets a reference to the given float64 and assigns it to the DurationMs field.
func (o *ClickhouseSlowQuery) SetDurationMs(v float64) {
	o.DurationMs = &v
}

// GetReadRows returns the ReadRows field value if set, zero value otherwise.
func (o *ClickhouseSlowQuery) GetReadRows() int64 {
	if o == nil || o.ReadRows == nil {
		var ret int64
		return ret
	}
	return *o.ReadRows
}

// GetReadRowsOk returns a tuple with the ReadRows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSlowQuery) GetReadRowsOk() (*int64, bool) {
	if o == nil || o.ReadRows == nil {
		return nil, false
	}
	return o.ReadRows, true
}

// HasReadRows returns a boolean if a field has been set.
func (o *ClickhouseSlowQuery) HasReadRows() bool {
	return o != nil && o.ReadRows != nil
}

// SetReadRows gets a reference to the given int64 and assigns it to the ReadRows field.
func (o *ClickhouseSlowQuery) SetReadRows(v int64) {
	o.ReadRows = &v
}

// GetReadBytes returns the ReadBytes field value if set, zero value otherwise.
func (o *ClickhouseSlowQuery) GetReadBytes() int64 {
	if o == nil || o.ReadBytes == nil {
		var ret int64
		return ret
	}
	return *o.ReadBytes
}

// GetReadBytesOk returns a tuple with the ReadBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSlowQuery) GetReadBytesOk() (*int64, bool) {
	if o == nil || o.ReadBytes == nil {
		return nil, false
	}
	return o.ReadBytes, true
}

// HasReadBytes returns a boolean if a field has been set.
func (o *ClickhouseSlowQuery) HasReadBytes() bool {
	return o != nil && o.ReadBytes != nil
}

// SetReadBytes gets a reference to the given int64 and assigns it to the ReadBytes field.
func (o *ClickhouseSlowQuery) SetReadBytes(v int64) {
	o.ReadBytes = &v
}

// GetMemoryUsageBytes returns the MemoryUsageBytes field value if set, zero value otherwise.
func (o *ClickhouseSlowQuery) GetMemoryUsageBytes() int64 {
	if o == nil || o.MemoryUsageBytes == nil {
		var ret int64
		return ret
	}
	return *o.MemoryUsageBytes
}

// GetMemoryUsageBytesOk returns a tuple with the MemoryUsageBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSlowQuery) GetMemoryUsageBytesOk() (*int64, bool) {
	if o == nil || o.MemoryUsageBytes == nil {
		return nil, false
	}
	return o.MemoryUsageBytes, true
}

// HasMemoryUsageBytes returns a boolean if a field has been set.
func (o *ClickhouseSlowQuery) HasMemoryUsageBytes() bool {
	return o != nil && o.MemoryUsageBytes != nil
}

// SetMemoryUsageBytes gets a reference to the given int64 and assigns it to the MemoryUsageBytes field.
func (o *ClickhouseSlowQuery) SetMemoryUsageBytes(v int64) {
	o.MemoryUsageBytes = &v
}

// GetResultRows returns the ResultRows field value if set, zero value otherwise.
func (o *ClickhouseSlowQuery) GetResultRows() int64 {
	if o == nil || o.ResultRows == nil {
		var ret int64
		return ret
	}
	return *o.ResultRows
}

// GetResultRowsOk returns a tuple with the ResultRows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSlowQuery) GetResultRowsOk() (*int64, bool) {
	if o == nil || o.ResultRows == nil {
		return nil, false
	}
	return o.ResultRows, true
}

// HasResultRows returns a boolean if a field has been set.
func (o *ClickhouseSlowQuery) HasResultRows() bool {
	return o != nil && o.ResultRows != nil
}

// SetResultRows gets a reference to the given int64 and assigns it to the ResultRows field.
func (o *ClickhouseSlowQuery) SetResultRows(v int64) {
	o.ResultRows = &v
}

// GetResultBytes returns the ResultBytes field value if set, zero value otherwise.
func (o *ClickhouseSlowQuery) GetResultBytes() int64 {
	if o == nil || o.ResultBytes == nil {
		var ret int64
		return ret
	}
	return *o.ResultBytes
}

// GetResultBytesOk returns a tuple with the ResultBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSlowQuery) GetResultBytesOk() (*int64, bool) {
	if o == nil || o.ResultBytes == nil {
		return nil, false
	}
	return o.ResultBytes, true
}

// HasResultBytes returns a boolean if a field has been set.
func (o *ClickhouseSlowQuery) HasResultBytes() bool {
	return o != nil && o.ResultBytes != nil
}

// SetResultBytes gets a reference to the given int64 and assigns it to the ResultBytes field.
func (o *ClickhouseSlowQuery) SetResultBytes(v int64) {
	o.ResultBytes = &v
}

// GetExceptionCode returns the ExceptionCode field value if set, zero value otherwise.
func (o *ClickhouseSlowQuery) GetExceptionCode() int64 {
	if o == nil || o.ExceptionCode == nil {
		var ret int64
		return ret
	}
	return *o.ExceptionCode
}

// GetExceptionCodeOk returns a tuple with the ExceptionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSlowQuery) GetExceptionCodeOk() (*int64, bool) {
	if o == nil || o.ExceptionCode == nil {
		return nil, false
	}
	return o.ExceptionCode, true
}

// HasExceptionCode returns a boolean if a field has been set.
func (o *ClickhouseSlowQuery) HasExceptionCode() bool {
	return o != nil && o.ExceptionCode != nil
}

// SetExceptionCode gets a reference to the given int64 and assigns it to the ExceptionCode field.
func (o *ClickhouseSlowQuery) SetExceptionCode(v int64) {
	o.ExceptionCode = &v
}

// GetException returns the Exception field value if set, zero value otherwise.
func (o *ClickhouseSlowQuery) GetException() string {
	if o == nil || o.Exception == nil {
		var ret string
		return ret
	}
	return *o.Exception
}

// GetExceptionOk returns a tuple with the Exception field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSlowQuery) GetExceptionOk() (*string, bool) {
	if o == nil || o.Exception == nil {
		return nil, false
	}
	return o.Exception, true
}

// HasException returns a boolean if a field has been set.
func (o *ClickhouseSlowQuery) HasException() bool {
	return o != nil && o.Exception != nil
}

// SetException gets a reference to the given string and assigns it to the Exception field.
func (o *ClickhouseSlowQuery) SetException(v string) {
	o.Exception = &v
}

// GetEventTime returns the EventTime field value if set, zero value otherwise.
func (o *ClickhouseSlowQuery) GetEventTime() string {
	if o == nil || o.EventTime == nil {
		var ret string
		return ret
	}
	return *o.EventTime
}

// GetEventTimeOk returns a tuple with the EventTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSlowQuery) GetEventTimeOk() (*string, bool) {
	if o == nil || o.EventTime == nil {
		return nil, false
	}
	return o.EventTime, true
}

// HasEventTime returns a boolean if a field has been set.
func (o *ClickhouseSlowQuery) HasEventTime() bool {
	return o != nil && o.EventTime != nil
}

// SetEventTime gets a reference to the given string and assigns it to the EventTime field.
func (o *ClickhouseSlowQuery) SetEventTime(v string) {
	o.EventTime = &v
}

// GetHostName returns the HostName field value if set, zero value otherwise.
func (o *ClickhouseSlowQuery) GetHostName() string {
	if o == nil || o.HostName == nil {
		var ret string
		return ret
	}
	return *o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSlowQuery) GetHostNameOk() (*string, bool) {
	if o == nil || o.HostName == nil {
		return nil, false
	}
	return o.HostName, true
}

// HasHostName returns a boolean if a field has been set.
func (o *ClickhouseSlowQuery) HasHostName() bool {
	return o != nil && o.HostName != nil
}

// SetHostName gets a reference to the given string and assigns it to the HostName field.
func (o *ClickhouseSlowQuery) SetHostName(v string) {
	o.HostName = &v
}

// GetClientName returns the ClientName field value if set, zero value otherwise.
func (o *ClickhouseSlowQuery) GetClientName() string {
	if o == nil || o.ClientName == nil {
		var ret string
		return ret
	}
	return *o.ClientName
}

// GetClientNameOk returns a tuple with the ClientName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSlowQuery) GetClientNameOk() (*string, bool) {
	if o == nil || o.ClientName == nil {
		return nil, false
	}
	return o.ClientName, true
}

// HasClientName returns a boolean if a field has been set.
func (o *ClickhouseSlowQuery) HasClientName() bool {
	return o != nil && o.ClientName != nil
}

// SetClientName gets a reference to the given string and assigns it to the ClientName field.
func (o *ClickhouseSlowQuery) SetClientName(v string) {
	o.ClientName = &v
}

// GetIsInitialQuery returns the IsInitialQuery field value if set, zero value otherwise.
func (o *ClickhouseSlowQuery) GetIsInitialQuery() bool {
	if o == nil || o.IsInitialQuery == nil {
		var ret bool
		return ret
	}
	return *o.IsInitialQuery
}

// GetIsInitialQueryOk returns a tuple with the IsInitialQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSlowQuery) GetIsInitialQueryOk() (*bool, bool) {
	if o == nil || o.IsInitialQuery == nil {
		return nil, false
	}
	return o.IsInitialQuery, true
}

// HasIsInitialQuery returns a boolean if a field has been set.
func (o *ClickhouseSlowQuery) HasIsInitialQuery() bool {
	return o != nil && o.IsInitialQuery != nil
}

// SetIsInitialQuery gets a reference to the given bool and assigns it to the IsInitialQuery field.
func (o *ClickhouseSlowQuery) SetIsInitialQuery(v bool) {
	o.IsInitialQuery = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClickhouseSlowQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.QueryId != nil {
		toSerialize["queryId"] = o.QueryId
	}
	if o.InitialQueryId != nil {
		toSerialize["initialQueryId"] = o.InitialQueryId
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Database != nil {
		toSerialize["database"] = o.Database
	}
	if o.QueryKind != nil {
		toSerialize["queryKind"] = o.QueryKind
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.NormalizedQueryHash != nil {
		toSerialize["normalizedQueryHash"] = o.NormalizedQueryHash
	}
	if o.DurationMs != nil {
		toSerialize["durationMs"] = o.DurationMs
	}
	if o.ReadRows != nil {
		toSerialize["readRows"] = o.ReadRows
	}
	if o.ReadBytes != nil {
		toSerialize["readBytes"] = o.ReadBytes
	}
	if o.MemoryUsageBytes != nil {
		toSerialize["memoryUsageBytes"] = o.MemoryUsageBytes
	}
	if o.ResultRows != nil {
		toSerialize["resultRows"] = o.ResultRows
	}
	if o.ResultBytes != nil {
		toSerialize["resultBytes"] = o.ResultBytes
	}
	if o.ExceptionCode != nil {
		toSerialize["exceptionCode"] = o.ExceptionCode
	}
	if o.Exception != nil {
		toSerialize["exception"] = o.Exception
	}
	if o.EventTime != nil {
		toSerialize["eventTime"] = o.EventTime
	}
	if o.HostName != nil {
		toSerialize["hostName"] = o.HostName
	}
	if o.ClientName != nil {
		toSerialize["clientName"] = o.ClientName
	}
	if o.IsInitialQuery != nil {
		toSerialize["isInitialQuery"] = o.IsInitialQuery
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClickhouseSlowQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		QueryId             *string  `json:"queryId,omitempty"`
		InitialQueryId      *string  `json:"initialQueryId,omitempty"`
		Type                *string  `json:"type,omitempty"`
		User                *string  `json:"user,omitempty"`
		Database            *string  `json:"database,omitempty"`
		QueryKind           *string  `json:"queryKind,omitempty"`
		Query               *string  `json:"query,omitempty"`
		NormalizedQueryHash *string  `json:"normalizedQueryHash,omitempty"`
		DurationMs          *float64 `json:"durationMs,omitempty"`
		ReadRows            *int64   `json:"readRows,omitempty"`
		ReadBytes           *int64   `json:"readBytes,omitempty"`
		MemoryUsageBytes    *int64   `json:"memoryUsageBytes,omitempty"`
		ResultRows          *int64   `json:"resultRows,omitempty"`
		ResultBytes         *int64   `json:"resultBytes,omitempty"`
		ExceptionCode       *int64   `json:"exceptionCode,omitempty"`
		Exception           *string  `json:"exception,omitempty"`
		EventTime           *string  `json:"eventTime,omitempty"`
		HostName            *string  `json:"hostName,omitempty"`
		ClientName          *string  `json:"clientName,omitempty"`
		IsInitialQuery      *bool    `json:"isInitialQuery,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"queryId", "initialQueryId", "type", "user", "database", "queryKind", "query", "normalizedQueryHash", "durationMs", "readRows", "readBytes", "memoryUsageBytes", "resultRows", "resultBytes", "exceptionCode", "exception", "eventTime", "hostName", "clientName", "isInitialQuery"})
	} else {
		return err
	}
	o.QueryId = all.QueryId
	o.InitialQueryId = all.InitialQueryId
	o.Type = all.Type
	o.User = all.User
	o.Database = all.Database
	o.QueryKind = all.QueryKind
	o.Query = all.Query
	o.NormalizedQueryHash = all.NormalizedQueryHash
	o.DurationMs = all.DurationMs
	o.ReadRows = all.ReadRows
	o.ReadBytes = all.ReadBytes
	o.MemoryUsageBytes = all.MemoryUsageBytes
	o.ResultRows = all.ResultRows
	o.ResultBytes = all.ResultBytes
	o.ExceptionCode = all.ExceptionCode
	o.Exception = all.Exception
	o.EventTime = all.EventTime
	o.HostName = all.HostName
	o.ClientName = all.ClientName
	o.IsInitialQuery = all.IsInitialQuery

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
