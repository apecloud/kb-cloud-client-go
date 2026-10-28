// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ClickhouseSession struct {
	// ClickHouse query_id. Use it with the KILL QUERY endpoint.
	QueryId string `json:"queryId"`
	// query_id of the initial (distributed) query when this row is a remote sub-query.
	InitialQueryId *string `json:"initialQueryId,omitempty"`
	// User that started the query.
	User *string `json:"user,omitempty"`
	// Client address.
	Address *string `json:"address,omitempty"`
	// Client port.
	Port *int64 `json:"port,omitempty"`
	// Current database of the query.
	Database *string `json:"database,omitempty"`
	// ClickHouse server host that reported the query.
	HostName *string `json:"hostName,omitempty"`
	// Seconds elapsed since the query started.
	ElapsedSeconds *float64 `json:"elapsedSeconds,omitempty"`
	// Rows read so far.
	ReadRows *int64 `json:"readRows,omitempty"`
	// Bytes read so far.
	ReadBytes *int64 `json:"readBytes,omitempty"`
	// Rows written so far.
	WrittenRows *int64 `json:"writtenRows,omitempty"`
	// Bytes written so far.
	WrittenBytes *int64 `json:"writtenBytes,omitempty"`
	// Current memory usage in bytes.
	MemoryUsageBytes *int64 `json:"memoryUsageBytes,omitempty"`
	// max_memory_usage from the query Settings map (only explicitly modified settings). Zero means the query did not set a limit, so the server default applies.
	MemoryLimitBytes *int64 `json:"memoryLimitBytes,omitempty"`
	// memory_usage / memoryLimitBytes (0-1). Zero when the query did not explicitly set max_memory_usage. Values close to 1 indicate the query is about to hit its memory limit.
	MemoryUsageRatio *float64 `json:"memoryUsageRatio,omitempty"`
	// Peak memory usage in bytes. Zero when the server version does not expose peak_memory_usage.
	PeakMemoryUsageBytes *int64 `json:"peakMemoryUsageBytes,omitempty"`
	// Whether this row is the initial query rather than a remote sub-query.
	IsInitialQuery *bool `json:"isInitialQuery,omitempty"`
	// Estimated query start time in RFC3339 UTC, computed from elapsed because system.processes has no start timestamp.
	QueryStartTime *string `json:"queryStartTime,omitempty"`
	// Full query text, truncated by the client for display when necessary.
	Query *string `json:"query,omitempty"`
	// Single-line query summary suitable for list rendering.
	QuerySummary *string `json:"querySummary,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClickhouseSession instantiates a new ClickhouseSession object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClickhouseSession(queryId string) *ClickhouseSession {
	this := ClickhouseSession{}
	this.QueryId = queryId
	return &this
}

// NewClickhouseSessionWithDefaults instantiates a new ClickhouseSession object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClickhouseSessionWithDefaults() *ClickhouseSession {
	this := ClickhouseSession{}
	return &this
}

// GetQueryId returns the QueryId field value.
func (o *ClickhouseSession) GetQueryId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.QueryId
}

// GetQueryIdOk returns a tuple with the QueryId field value
// and a boolean to check if the value has been set.
func (o *ClickhouseSession) GetQueryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryId, true
}

// SetQueryId sets field value.
func (o *ClickhouseSession) SetQueryId(v string) {
	o.QueryId = v
}

// GetInitialQueryId returns the InitialQueryId field value if set, zero value otherwise.
func (o *ClickhouseSession) GetInitialQueryId() string {
	if o == nil || o.InitialQueryId == nil {
		var ret string
		return ret
	}
	return *o.InitialQueryId
}

// GetInitialQueryIdOk returns a tuple with the InitialQueryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSession) GetInitialQueryIdOk() (*string, bool) {
	if o == nil || o.InitialQueryId == nil {
		return nil, false
	}
	return o.InitialQueryId, true
}

// HasInitialQueryId returns a boolean if a field has been set.
func (o *ClickhouseSession) HasInitialQueryId() bool {
	return o != nil && o.InitialQueryId != nil
}

// SetInitialQueryId gets a reference to the given string and assigns it to the InitialQueryId field.
func (o *ClickhouseSession) SetInitialQueryId(v string) {
	o.InitialQueryId = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *ClickhouseSession) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSession) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *ClickhouseSession) HasUser() bool {
	return o != nil && o.User != nil
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *ClickhouseSession) SetUser(v string) {
	o.User = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ClickhouseSession) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSession) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ClickhouseSession) HasAddress() bool {
	return o != nil && o.Address != nil
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *ClickhouseSession) SetAddress(v string) {
	o.Address = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *ClickhouseSession) GetPort() int64 {
	if o == nil || o.Port == nil {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSession) GetPortOk() (*int64, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *ClickhouseSession) HasPort() bool {
	return o != nil && o.Port != nil
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *ClickhouseSession) SetPort(v int64) {
	o.Port = &v
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *ClickhouseSession) GetDatabase() string {
	if o == nil || o.Database == nil {
		var ret string
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSession) GetDatabaseOk() (*string, bool) {
	if o == nil || o.Database == nil {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *ClickhouseSession) HasDatabase() bool {
	return o != nil && o.Database != nil
}

// SetDatabase gets a reference to the given string and assigns it to the Database field.
func (o *ClickhouseSession) SetDatabase(v string) {
	o.Database = &v
}

// GetHostName returns the HostName field value if set, zero value otherwise.
func (o *ClickhouseSession) GetHostName() string {
	if o == nil || o.HostName == nil {
		var ret string
		return ret
	}
	return *o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSession) GetHostNameOk() (*string, bool) {
	if o == nil || o.HostName == nil {
		return nil, false
	}
	return o.HostName, true
}

// HasHostName returns a boolean if a field has been set.
func (o *ClickhouseSession) HasHostName() bool {
	return o != nil && o.HostName != nil
}

// SetHostName gets a reference to the given string and assigns it to the HostName field.
func (o *ClickhouseSession) SetHostName(v string) {
	o.HostName = &v
}

// GetElapsedSeconds returns the ElapsedSeconds field value if set, zero value otherwise.
func (o *ClickhouseSession) GetElapsedSeconds() float64 {
	if o == nil || o.ElapsedSeconds == nil {
		var ret float64
		return ret
	}
	return *o.ElapsedSeconds
}

// GetElapsedSecondsOk returns a tuple with the ElapsedSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSession) GetElapsedSecondsOk() (*float64, bool) {
	if o == nil || o.ElapsedSeconds == nil {
		return nil, false
	}
	return o.ElapsedSeconds, true
}

// HasElapsedSeconds returns a boolean if a field has been set.
func (o *ClickhouseSession) HasElapsedSeconds() bool {
	return o != nil && o.ElapsedSeconds != nil
}

// SetElapsedSeconds gets a reference to the given float64 and assigns it to the ElapsedSeconds field.
func (o *ClickhouseSession) SetElapsedSeconds(v float64) {
	o.ElapsedSeconds = &v
}

// GetReadRows returns the ReadRows field value if set, zero value otherwise.
func (o *ClickhouseSession) GetReadRows() int64 {
	if o == nil || o.ReadRows == nil {
		var ret int64
		return ret
	}
	return *o.ReadRows
}

// GetReadRowsOk returns a tuple with the ReadRows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSession) GetReadRowsOk() (*int64, bool) {
	if o == nil || o.ReadRows == nil {
		return nil, false
	}
	return o.ReadRows, true
}

// HasReadRows returns a boolean if a field has been set.
func (o *ClickhouseSession) HasReadRows() bool {
	return o != nil && o.ReadRows != nil
}

// SetReadRows gets a reference to the given int64 and assigns it to the ReadRows field.
func (o *ClickhouseSession) SetReadRows(v int64) {
	o.ReadRows = &v
}

// GetReadBytes returns the ReadBytes field value if set, zero value otherwise.
func (o *ClickhouseSession) GetReadBytes() int64 {
	if o == nil || o.ReadBytes == nil {
		var ret int64
		return ret
	}
	return *o.ReadBytes
}

// GetReadBytesOk returns a tuple with the ReadBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSession) GetReadBytesOk() (*int64, bool) {
	if o == nil || o.ReadBytes == nil {
		return nil, false
	}
	return o.ReadBytes, true
}

// HasReadBytes returns a boolean if a field has been set.
func (o *ClickhouseSession) HasReadBytes() bool {
	return o != nil && o.ReadBytes != nil
}

// SetReadBytes gets a reference to the given int64 and assigns it to the ReadBytes field.
func (o *ClickhouseSession) SetReadBytes(v int64) {
	o.ReadBytes = &v
}

// GetWrittenRows returns the WrittenRows field value if set, zero value otherwise.
func (o *ClickhouseSession) GetWrittenRows() int64 {
	if o == nil || o.WrittenRows == nil {
		var ret int64
		return ret
	}
	return *o.WrittenRows
}

// GetWrittenRowsOk returns a tuple with the WrittenRows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSession) GetWrittenRowsOk() (*int64, bool) {
	if o == nil || o.WrittenRows == nil {
		return nil, false
	}
	return o.WrittenRows, true
}

// HasWrittenRows returns a boolean if a field has been set.
func (o *ClickhouseSession) HasWrittenRows() bool {
	return o != nil && o.WrittenRows != nil
}

// SetWrittenRows gets a reference to the given int64 and assigns it to the WrittenRows field.
func (o *ClickhouseSession) SetWrittenRows(v int64) {
	o.WrittenRows = &v
}

// GetWrittenBytes returns the WrittenBytes field value if set, zero value otherwise.
func (o *ClickhouseSession) GetWrittenBytes() int64 {
	if o == nil || o.WrittenBytes == nil {
		var ret int64
		return ret
	}
	return *o.WrittenBytes
}

// GetWrittenBytesOk returns a tuple with the WrittenBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSession) GetWrittenBytesOk() (*int64, bool) {
	if o == nil || o.WrittenBytes == nil {
		return nil, false
	}
	return o.WrittenBytes, true
}

// HasWrittenBytes returns a boolean if a field has been set.
func (o *ClickhouseSession) HasWrittenBytes() bool {
	return o != nil && o.WrittenBytes != nil
}

// SetWrittenBytes gets a reference to the given int64 and assigns it to the WrittenBytes field.
func (o *ClickhouseSession) SetWrittenBytes(v int64) {
	o.WrittenBytes = &v
}

// GetMemoryUsageBytes returns the MemoryUsageBytes field value if set, zero value otherwise.
func (o *ClickhouseSession) GetMemoryUsageBytes() int64 {
	if o == nil || o.MemoryUsageBytes == nil {
		var ret int64
		return ret
	}
	return *o.MemoryUsageBytes
}

// GetMemoryUsageBytesOk returns a tuple with the MemoryUsageBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSession) GetMemoryUsageBytesOk() (*int64, bool) {
	if o == nil || o.MemoryUsageBytes == nil {
		return nil, false
	}
	return o.MemoryUsageBytes, true
}

// HasMemoryUsageBytes returns a boolean if a field has been set.
func (o *ClickhouseSession) HasMemoryUsageBytes() bool {
	return o != nil && o.MemoryUsageBytes != nil
}

// SetMemoryUsageBytes gets a reference to the given int64 and assigns it to the MemoryUsageBytes field.
func (o *ClickhouseSession) SetMemoryUsageBytes(v int64) {
	o.MemoryUsageBytes = &v
}

// GetMemoryLimitBytes returns the MemoryLimitBytes field value if set, zero value otherwise.
func (o *ClickhouseSession) GetMemoryLimitBytes() int64 {
	if o == nil || o.MemoryLimitBytes == nil {
		var ret int64
		return ret
	}
	return *o.MemoryLimitBytes
}

// GetMemoryLimitBytesOk returns a tuple with the MemoryLimitBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSession) GetMemoryLimitBytesOk() (*int64, bool) {
	if o == nil || o.MemoryLimitBytes == nil {
		return nil, false
	}
	return o.MemoryLimitBytes, true
}

// HasMemoryLimitBytes returns a boolean if a field has been set.
func (o *ClickhouseSession) HasMemoryLimitBytes() bool {
	return o != nil && o.MemoryLimitBytes != nil
}

// SetMemoryLimitBytes gets a reference to the given int64 and assigns it to the MemoryLimitBytes field.
func (o *ClickhouseSession) SetMemoryLimitBytes(v int64) {
	o.MemoryLimitBytes = &v
}

// GetMemoryUsageRatio returns the MemoryUsageRatio field value if set, zero value otherwise.
func (o *ClickhouseSession) GetMemoryUsageRatio() float64 {
	if o == nil || o.MemoryUsageRatio == nil {
		var ret float64
		return ret
	}
	return *o.MemoryUsageRatio
}

// GetMemoryUsageRatioOk returns a tuple with the MemoryUsageRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSession) GetMemoryUsageRatioOk() (*float64, bool) {
	if o == nil || o.MemoryUsageRatio == nil {
		return nil, false
	}
	return o.MemoryUsageRatio, true
}

// HasMemoryUsageRatio returns a boolean if a field has been set.
func (o *ClickhouseSession) HasMemoryUsageRatio() bool {
	return o != nil && o.MemoryUsageRatio != nil
}

// SetMemoryUsageRatio gets a reference to the given float64 and assigns it to the MemoryUsageRatio field.
func (o *ClickhouseSession) SetMemoryUsageRatio(v float64) {
	o.MemoryUsageRatio = &v
}

// GetPeakMemoryUsageBytes returns the PeakMemoryUsageBytes field value if set, zero value otherwise.
func (o *ClickhouseSession) GetPeakMemoryUsageBytes() int64 {
	if o == nil || o.PeakMemoryUsageBytes == nil {
		var ret int64
		return ret
	}
	return *o.PeakMemoryUsageBytes
}

// GetPeakMemoryUsageBytesOk returns a tuple with the PeakMemoryUsageBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSession) GetPeakMemoryUsageBytesOk() (*int64, bool) {
	if o == nil || o.PeakMemoryUsageBytes == nil {
		return nil, false
	}
	return o.PeakMemoryUsageBytes, true
}

// HasPeakMemoryUsageBytes returns a boolean if a field has been set.
func (o *ClickhouseSession) HasPeakMemoryUsageBytes() bool {
	return o != nil && o.PeakMemoryUsageBytes != nil
}

// SetPeakMemoryUsageBytes gets a reference to the given int64 and assigns it to the PeakMemoryUsageBytes field.
func (o *ClickhouseSession) SetPeakMemoryUsageBytes(v int64) {
	o.PeakMemoryUsageBytes = &v
}

// GetIsInitialQuery returns the IsInitialQuery field value if set, zero value otherwise.
func (o *ClickhouseSession) GetIsInitialQuery() bool {
	if o == nil || o.IsInitialQuery == nil {
		var ret bool
		return ret
	}
	return *o.IsInitialQuery
}

// GetIsInitialQueryOk returns a tuple with the IsInitialQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSession) GetIsInitialQueryOk() (*bool, bool) {
	if o == nil || o.IsInitialQuery == nil {
		return nil, false
	}
	return o.IsInitialQuery, true
}

// HasIsInitialQuery returns a boolean if a field has been set.
func (o *ClickhouseSession) HasIsInitialQuery() bool {
	return o != nil && o.IsInitialQuery != nil
}

// SetIsInitialQuery gets a reference to the given bool and assigns it to the IsInitialQuery field.
func (o *ClickhouseSession) SetIsInitialQuery(v bool) {
	o.IsInitialQuery = &v
}

// GetQueryStartTime returns the QueryStartTime field value if set, zero value otherwise.
func (o *ClickhouseSession) GetQueryStartTime() string {
	if o == nil || o.QueryStartTime == nil {
		var ret string
		return ret
	}
	return *o.QueryStartTime
}

// GetQueryStartTimeOk returns a tuple with the QueryStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSession) GetQueryStartTimeOk() (*string, bool) {
	if o == nil || o.QueryStartTime == nil {
		return nil, false
	}
	return o.QueryStartTime, true
}

// HasQueryStartTime returns a boolean if a field has been set.
func (o *ClickhouseSession) HasQueryStartTime() bool {
	return o != nil && o.QueryStartTime != nil
}

// SetQueryStartTime gets a reference to the given string and assigns it to the QueryStartTime field.
func (o *ClickhouseSession) SetQueryStartTime(v string) {
	o.QueryStartTime = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *ClickhouseSession) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSession) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *ClickhouseSession) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *ClickhouseSession) SetQuery(v string) {
	o.Query = &v
}

// GetQuerySummary returns the QuerySummary field value if set, zero value otherwise.
func (o *ClickhouseSession) GetQuerySummary() string {
	if o == nil || o.QuerySummary == nil {
		var ret string
		return ret
	}
	return *o.QuerySummary
}

// GetQuerySummaryOk returns a tuple with the QuerySummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSession) GetQuerySummaryOk() (*string, bool) {
	if o == nil || o.QuerySummary == nil {
		return nil, false
	}
	return o.QuerySummary, true
}

// HasQuerySummary returns a boolean if a field has been set.
func (o *ClickhouseSession) HasQuerySummary() bool {
	return o != nil && o.QuerySummary != nil
}

// SetQuerySummary gets a reference to the given string and assigns it to the QuerySummary field.
func (o *ClickhouseSession) SetQuerySummary(v string) {
	o.QuerySummary = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClickhouseSession) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["queryId"] = o.QueryId
	if o.InitialQueryId != nil {
		toSerialize["initialQueryId"] = o.InitialQueryId
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.Database != nil {
		toSerialize["database"] = o.Database
	}
	if o.HostName != nil {
		toSerialize["hostName"] = o.HostName
	}
	if o.ElapsedSeconds != nil {
		toSerialize["elapsedSeconds"] = o.ElapsedSeconds
	}
	if o.ReadRows != nil {
		toSerialize["readRows"] = o.ReadRows
	}
	if o.ReadBytes != nil {
		toSerialize["readBytes"] = o.ReadBytes
	}
	if o.WrittenRows != nil {
		toSerialize["writtenRows"] = o.WrittenRows
	}
	if o.WrittenBytes != nil {
		toSerialize["writtenBytes"] = o.WrittenBytes
	}
	if o.MemoryUsageBytes != nil {
		toSerialize["memoryUsageBytes"] = o.MemoryUsageBytes
	}
	if o.MemoryLimitBytes != nil {
		toSerialize["memoryLimitBytes"] = o.MemoryLimitBytes
	}
	if o.MemoryUsageRatio != nil {
		toSerialize["memoryUsageRatio"] = o.MemoryUsageRatio
	}
	if o.PeakMemoryUsageBytes != nil {
		toSerialize["peakMemoryUsageBytes"] = o.PeakMemoryUsageBytes
	}
	if o.IsInitialQuery != nil {
		toSerialize["isInitialQuery"] = o.IsInitialQuery
	}
	if o.QueryStartTime != nil {
		toSerialize["queryStartTime"] = o.QueryStartTime
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.QuerySummary != nil {
		toSerialize["querySummary"] = o.QuerySummary
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClickhouseSession) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		QueryId              *string  `json:"queryId"`
		InitialQueryId       *string  `json:"initialQueryId,omitempty"`
		User                 *string  `json:"user,omitempty"`
		Address              *string  `json:"address,omitempty"`
		Port                 *int64   `json:"port,omitempty"`
		Database             *string  `json:"database,omitempty"`
		HostName             *string  `json:"hostName,omitempty"`
		ElapsedSeconds       *float64 `json:"elapsedSeconds,omitempty"`
		ReadRows             *int64   `json:"readRows,omitempty"`
		ReadBytes            *int64   `json:"readBytes,omitempty"`
		WrittenRows          *int64   `json:"writtenRows,omitempty"`
		WrittenBytes         *int64   `json:"writtenBytes,omitempty"`
		MemoryUsageBytes     *int64   `json:"memoryUsageBytes,omitempty"`
		MemoryLimitBytes     *int64   `json:"memoryLimitBytes,omitempty"`
		MemoryUsageRatio     *float64 `json:"memoryUsageRatio,omitempty"`
		PeakMemoryUsageBytes *int64   `json:"peakMemoryUsageBytes,omitempty"`
		IsInitialQuery       *bool    `json:"isInitialQuery,omitempty"`
		QueryStartTime       *string  `json:"queryStartTime,omitempty"`
		Query                *string  `json:"query,omitempty"`
		QuerySummary         *string  `json:"querySummary,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.QueryId == nil {
		return fmt.Errorf("required field queryId missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"queryId", "initialQueryId", "user", "address", "port", "database", "hostName", "elapsedSeconds", "readRows", "readBytes", "writtenRows", "writtenBytes", "memoryUsageBytes", "memoryLimitBytes", "memoryUsageRatio", "peakMemoryUsageBytes", "isInitialQuery", "queryStartTime", "query", "querySummary"})
	} else {
		return err
	}
	o.QueryId = *all.QueryId
	o.InitialQueryId = all.InitialQueryId
	o.User = all.User
	o.Address = all.Address
	o.Port = all.Port
	o.Database = all.Database
	o.HostName = all.HostName
	o.ElapsedSeconds = all.ElapsedSeconds
	o.ReadRows = all.ReadRows
	o.ReadBytes = all.ReadBytes
	o.WrittenRows = all.WrittenRows
	o.WrittenBytes = all.WrittenBytes
	o.MemoryUsageBytes = all.MemoryUsageBytes
	o.MemoryLimitBytes = all.MemoryLimitBytes
	o.MemoryUsageRatio = all.MemoryUsageRatio
	o.PeakMemoryUsageBytes = all.PeakMemoryUsageBytes
	o.IsInitialQuery = all.IsInitialQuery
	o.QueryStartTime = all.QueryStartTime
	o.Query = all.Query
	o.QuerySummary = all.QuerySummary

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
