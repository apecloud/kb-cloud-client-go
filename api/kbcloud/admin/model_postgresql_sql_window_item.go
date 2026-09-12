// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// PostgresqlSQLWindowItem Delta sums after delivery deduplication. Calls and rows at or above 2^53 are unavailable because the storage aggregation uses float64. Summary contains only collector-redacted SQL; meanTimeMs is null for zero calls.
type PostgresqlSQLWindowItem struct {
	SqlId *string `json:"sqlId,omitempty"`
	// Optional sums of SQL resource deltas. Missing values are unknown, never zero. Block counters are PostgreSQL blocks, WAL bytes are bytes, and timing fields are milliseconds. Legacy blk timing is the older PostgreSQL block I/O timing; it is not interchangeable with the PostgreSQL 17 split fields. Counters at or above 2^53 are unavailable to avoid storage aggregation precision loss.
	Resources *PostgresqlSQLWindowResources `json:"resources,omitempty"`
	// Resource field name to unavailable reason (unsupported, disabled, configuration_changed, invalid_value, not_collected, or partial). A resource sum is present only when every contributing deduplicated SQL observation has a valid value. This does not override collection-level partial coverage.
	UnavailableMetrics map[string]string                `json:"unavailableMetrics,omitempty"`
	QueryId            *string                          `json:"queryId,omitempty"`
	DatabaseId         *string                          `json:"databaseId,omitempty"`
	DatabaseName       *string                          `json:"databaseName,omitempty"`
	UserId             *string                          `json:"userId,omitempty"`
	UserName           *string                          `json:"userName,omitempty"`
	TopLevel           *PostgresqlSQLWindowTopLevel     `json:"topLevel,omitempty"`
	Summary            *string                          `json:"summary,omitempty"`
	SummaryState       *PostgresqlSQLWindowSummaryState `json:"summaryState,omitempty"`
	Calls              *int64                           `json:"calls,omitempty"`
	TotalTimeMs        *float64                         `json:"totalTimeMs,omitempty"`
	MeanTimeMs         common.NullableFloat64           `json:"meanTimeMs,omitempty"`
	Rows               *int64                           `json:"rows,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlSQLWindowItem instantiates a new PostgresqlSQLWindowItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlSQLWindowItem() *PostgresqlSQLWindowItem {
	this := PostgresqlSQLWindowItem{}
	return &this
}

// NewPostgresqlSQLWindowItemWithDefaults instantiates a new PostgresqlSQLWindowItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlSQLWindowItemWithDefaults() *PostgresqlSQLWindowItem {
	this := PostgresqlSQLWindowItem{}
	return &this
}

// GetSqlId returns the SqlId field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowItem) GetSqlId() string {
	if o == nil || o.SqlId == nil {
		var ret string
		return ret
	}
	return *o.SqlId
}

// GetSqlIdOk returns a tuple with the SqlId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowItem) GetSqlIdOk() (*string, bool) {
	if o == nil || o.SqlId == nil {
		return nil, false
	}
	return o.SqlId, true
}

// HasSqlId returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowItem) HasSqlId() bool {
	return o != nil && o.SqlId != nil
}

// SetSqlId gets a reference to the given string and assigns it to the SqlId field.
func (o *PostgresqlSQLWindowItem) SetSqlId(v string) {
	o.SqlId = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowItem) GetResources() PostgresqlSQLWindowResources {
	if o == nil || o.Resources == nil {
		var ret PostgresqlSQLWindowResources
		return ret
	}
	return *o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowItem) GetResourcesOk() (*PostgresqlSQLWindowResources, bool) {
	if o == nil || o.Resources == nil {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowItem) HasResources() bool {
	return o != nil && o.Resources != nil
}

// SetResources gets a reference to the given PostgresqlSQLWindowResources and assigns it to the Resources field.
func (o *PostgresqlSQLWindowItem) SetResources(v PostgresqlSQLWindowResources) {
	o.Resources = &v
}

// GetUnavailableMetrics returns the UnavailableMetrics field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowItem) GetUnavailableMetrics() map[string]string {
	if o == nil || o.UnavailableMetrics == nil {
		var ret map[string]string
		return ret
	}
	return o.UnavailableMetrics
}

// GetUnavailableMetricsOk returns a tuple with the UnavailableMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowItem) GetUnavailableMetricsOk() (*map[string]string, bool) {
	if o == nil || o.UnavailableMetrics == nil {
		return nil, false
	}
	return &o.UnavailableMetrics, true
}

// HasUnavailableMetrics returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowItem) HasUnavailableMetrics() bool {
	return o != nil && o.UnavailableMetrics != nil
}

// SetUnavailableMetrics gets a reference to the given map[string]string and assigns it to the UnavailableMetrics field.
func (o *PostgresqlSQLWindowItem) SetUnavailableMetrics(v map[string]string) {
	o.UnavailableMetrics = v
}

// GetQueryId returns the QueryId field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowItem) GetQueryId() string {
	if o == nil || o.QueryId == nil {
		var ret string
		return ret
	}
	return *o.QueryId
}

// GetQueryIdOk returns a tuple with the QueryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowItem) GetQueryIdOk() (*string, bool) {
	if o == nil || o.QueryId == nil {
		return nil, false
	}
	return o.QueryId, true
}

// HasQueryId returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowItem) HasQueryId() bool {
	return o != nil && o.QueryId != nil
}

// SetQueryId gets a reference to the given string and assigns it to the QueryId field.
func (o *PostgresqlSQLWindowItem) SetQueryId(v string) {
	o.QueryId = &v
}

// GetDatabaseId returns the DatabaseId field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowItem) GetDatabaseId() string {
	if o == nil || o.DatabaseId == nil {
		var ret string
		return ret
	}
	return *o.DatabaseId
}

// GetDatabaseIdOk returns a tuple with the DatabaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowItem) GetDatabaseIdOk() (*string, bool) {
	if o == nil || o.DatabaseId == nil {
		return nil, false
	}
	return o.DatabaseId, true
}

// HasDatabaseId returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowItem) HasDatabaseId() bool {
	return o != nil && o.DatabaseId != nil
}

// SetDatabaseId gets a reference to the given string and assigns it to the DatabaseId field.
func (o *PostgresqlSQLWindowItem) SetDatabaseId(v string) {
	o.DatabaseId = &v
}

// GetDatabaseName returns the DatabaseName field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowItem) GetDatabaseName() string {
	if o == nil || o.DatabaseName == nil {
		var ret string
		return ret
	}
	return *o.DatabaseName
}

// GetDatabaseNameOk returns a tuple with the DatabaseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowItem) GetDatabaseNameOk() (*string, bool) {
	if o == nil || o.DatabaseName == nil {
		return nil, false
	}
	return o.DatabaseName, true
}

// HasDatabaseName returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowItem) HasDatabaseName() bool {
	return o != nil && o.DatabaseName != nil
}

// SetDatabaseName gets a reference to the given string and assigns it to the DatabaseName field.
func (o *PostgresqlSQLWindowItem) SetDatabaseName(v string) {
	o.DatabaseName = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowItem) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowItem) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowItem) HasUserId() bool {
	return o != nil && o.UserId != nil
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *PostgresqlSQLWindowItem) SetUserId(v string) {
	o.UserId = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowItem) GetUserName() string {
	if o == nil || o.UserName == nil {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowItem) GetUserNameOk() (*string, bool) {
	if o == nil || o.UserName == nil {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowItem) HasUserName() bool {
	return o != nil && o.UserName != nil
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *PostgresqlSQLWindowItem) SetUserName(v string) {
	o.UserName = &v
}

// GetTopLevel returns the TopLevel field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowItem) GetTopLevel() PostgresqlSQLWindowTopLevel {
	if o == nil || o.TopLevel == nil {
		var ret PostgresqlSQLWindowTopLevel
		return ret
	}
	return *o.TopLevel
}

// GetTopLevelOk returns a tuple with the TopLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowItem) GetTopLevelOk() (*PostgresqlSQLWindowTopLevel, bool) {
	if o == nil || o.TopLevel == nil {
		return nil, false
	}
	return o.TopLevel, true
}

// HasTopLevel returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowItem) HasTopLevel() bool {
	return o != nil && o.TopLevel != nil
}

// SetTopLevel gets a reference to the given PostgresqlSQLWindowTopLevel and assigns it to the TopLevel field.
func (o *PostgresqlSQLWindowItem) SetTopLevel(v PostgresqlSQLWindowTopLevel) {
	o.TopLevel = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowItem) GetSummary() string {
	if o == nil || o.Summary == nil {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowItem) GetSummaryOk() (*string, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowItem) HasSummary() bool {
	return o != nil && o.Summary != nil
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *PostgresqlSQLWindowItem) SetSummary(v string) {
	o.Summary = &v
}

// GetSummaryState returns the SummaryState field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowItem) GetSummaryState() PostgresqlSQLWindowSummaryState {
	if o == nil || o.SummaryState == nil {
		var ret PostgresqlSQLWindowSummaryState
		return ret
	}
	return *o.SummaryState
}

// GetSummaryStateOk returns a tuple with the SummaryState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowItem) GetSummaryStateOk() (*PostgresqlSQLWindowSummaryState, bool) {
	if o == nil || o.SummaryState == nil {
		return nil, false
	}
	return o.SummaryState, true
}

// HasSummaryState returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowItem) HasSummaryState() bool {
	return o != nil && o.SummaryState != nil
}

// SetSummaryState gets a reference to the given PostgresqlSQLWindowSummaryState and assigns it to the SummaryState field.
func (o *PostgresqlSQLWindowItem) SetSummaryState(v PostgresqlSQLWindowSummaryState) {
	o.SummaryState = &v
}

// GetCalls returns the Calls field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowItem) GetCalls() int64 {
	if o == nil || o.Calls == nil {
		var ret int64
		return ret
	}
	return *o.Calls
}

// GetCallsOk returns a tuple with the Calls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowItem) GetCallsOk() (*int64, bool) {
	if o == nil || o.Calls == nil {
		return nil, false
	}
	return o.Calls, true
}

// HasCalls returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowItem) HasCalls() bool {
	return o != nil && o.Calls != nil
}

// SetCalls gets a reference to the given int64 and assigns it to the Calls field.
func (o *PostgresqlSQLWindowItem) SetCalls(v int64) {
	o.Calls = &v
}

// GetTotalTimeMs returns the TotalTimeMs field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowItem) GetTotalTimeMs() float64 {
	if o == nil || o.TotalTimeMs == nil {
		var ret float64
		return ret
	}
	return *o.TotalTimeMs
}

// GetTotalTimeMsOk returns a tuple with the TotalTimeMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowItem) GetTotalTimeMsOk() (*float64, bool) {
	if o == nil || o.TotalTimeMs == nil {
		return nil, false
	}
	return o.TotalTimeMs, true
}

// HasTotalTimeMs returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowItem) HasTotalTimeMs() bool {
	return o != nil && o.TotalTimeMs != nil
}

// SetTotalTimeMs gets a reference to the given float64 and assigns it to the TotalTimeMs field.
func (o *PostgresqlSQLWindowItem) SetTotalTimeMs(v float64) {
	o.TotalTimeMs = &v
}

// GetMeanTimeMs returns the MeanTimeMs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostgresqlSQLWindowItem) GetMeanTimeMs() float64 {
	if o == nil || o.MeanTimeMs.Get() == nil {
		var ret float64
		return ret
	}
	return *o.MeanTimeMs.Get()
}

// GetMeanTimeMsOk returns a tuple with the MeanTimeMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PostgresqlSQLWindowItem) GetMeanTimeMsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MeanTimeMs.Get(), o.MeanTimeMs.IsSet()
}

// HasMeanTimeMs returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowItem) HasMeanTimeMs() bool {
	return o != nil && o.MeanTimeMs.IsSet()
}

// SetMeanTimeMs gets a reference to the given common.NullableFloat64 and assigns it to the MeanTimeMs field.
func (o *PostgresqlSQLWindowItem) SetMeanTimeMs(v float64) {
	o.MeanTimeMs.Set(&v)
}

// SetMeanTimeMsNil sets the value for MeanTimeMs to be an explicit nil.
func (o *PostgresqlSQLWindowItem) SetMeanTimeMsNil() {
	o.MeanTimeMs.Set(nil)
}

// UnsetMeanTimeMs ensures that no value is present for MeanTimeMs, not even an explicit nil.
func (o *PostgresqlSQLWindowItem) UnsetMeanTimeMs() {
	o.MeanTimeMs.Unset()
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowItem) GetRows() int64 {
	if o == nil || o.Rows == nil {
		var ret int64
		return ret
	}
	return *o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowItem) GetRowsOk() (*int64, bool) {
	if o == nil || o.Rows == nil {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowItem) HasRows() bool {
	return o != nil && o.Rows != nil
}

// SetRows gets a reference to the given int64 and assigns it to the Rows field.
func (o *PostgresqlSQLWindowItem) SetRows(v int64) {
	o.Rows = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlSQLWindowItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.SqlId != nil {
		toSerialize["sqlId"] = o.SqlId
	}
	if o.Resources != nil {
		toSerialize["resources"] = o.Resources
	}
	if o.UnavailableMetrics != nil {
		toSerialize["unavailableMetrics"] = o.UnavailableMetrics
	}
	if o.QueryId != nil {
		toSerialize["queryId"] = o.QueryId
	}
	if o.DatabaseId != nil {
		toSerialize["databaseId"] = o.DatabaseId
	}
	if o.DatabaseName != nil {
		toSerialize["databaseName"] = o.DatabaseName
	}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	if o.UserName != nil {
		toSerialize["userName"] = o.UserName
	}
	if o.TopLevel != nil {
		toSerialize["topLevel"] = o.TopLevel
	}
	if o.Summary != nil {
		toSerialize["summary"] = o.Summary
	}
	if o.SummaryState != nil {
		toSerialize["summaryState"] = o.SummaryState
	}
	if o.Calls != nil {
		toSerialize["calls"] = o.Calls
	}
	if o.TotalTimeMs != nil {
		toSerialize["totalTimeMs"] = o.TotalTimeMs
	}
	if o.MeanTimeMs.IsSet() {
		toSerialize["meanTimeMs"] = o.MeanTimeMs.Get()
	}
	if o.Rows != nil {
		toSerialize["rows"] = o.Rows
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostgresqlSQLWindowItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SqlId              *string                          `json:"sqlId,omitempty"`
		Resources          *PostgresqlSQLWindowResources    `json:"resources,omitempty"`
		UnavailableMetrics map[string]string                `json:"unavailableMetrics,omitempty"`
		QueryId            *string                          `json:"queryId,omitempty"`
		DatabaseId         *string                          `json:"databaseId,omitempty"`
		DatabaseName       *string                          `json:"databaseName,omitempty"`
		UserId             *string                          `json:"userId,omitempty"`
		UserName           *string                          `json:"userName,omitempty"`
		TopLevel           *PostgresqlSQLWindowTopLevel     `json:"topLevel,omitempty"`
		Summary            *string                          `json:"summary,omitempty"`
		SummaryState       *PostgresqlSQLWindowSummaryState `json:"summaryState,omitempty"`
		Calls              *int64                           `json:"calls,omitempty"`
		TotalTimeMs        *float64                         `json:"totalTimeMs,omitempty"`
		MeanTimeMs         common.NullableFloat64           `json:"meanTimeMs,omitempty"`
		Rows               *int64                           `json:"rows,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"sqlId", "resources", "unavailableMetrics", "queryId", "databaseId", "databaseName", "userId", "userName", "topLevel", "summary", "summaryState", "calls", "totalTimeMs", "meanTimeMs", "rows"})
	} else {
		return err
	}

	hasInvalidField := false
	o.SqlId = all.SqlId
	if all.Resources != nil && all.Resources.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Resources = all.Resources
	o.UnavailableMetrics = all.UnavailableMetrics
	o.QueryId = all.QueryId
	o.DatabaseId = all.DatabaseId
	o.DatabaseName = all.DatabaseName
	o.UserId = all.UserId
	o.UserName = all.UserName
	if all.TopLevel != nil && !all.TopLevel.IsValid() {
		hasInvalidField = true
	} else {
		o.TopLevel = all.TopLevel
	}
	o.Summary = all.Summary
	if all.SummaryState != nil && !all.SummaryState.IsValid() {
		hasInvalidField = true
	} else {
		o.SummaryState = all.SummaryState
	}
	o.Calls = all.Calls
	o.TotalTimeMs = all.TotalTimeMs
	o.MeanTimeMs = all.MeanTimeMs
	o.Rows = all.Rows

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
