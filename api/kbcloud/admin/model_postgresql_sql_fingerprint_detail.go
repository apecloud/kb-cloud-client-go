// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlSQLFingerprintDetail struct {
	// PostgreSQL pg_stat_statements queryid represented as a string.
	QueryId string `json:"queryID"`
	// Stable SQL fingerprint identifier for UI grouping. M1 uses queryID.
	Fingerprint string `json:"fingerprint"`
	// Redacted SQL summary. Full raw SQL is intentionally not returned.
	QuerySummary *string `json:"querySummary,omitempty"`
	// Number of executions accumulated by pg_stat_statements since statsReset.
	Calls *int64 `json:"calls,omitempty"`
	// Total execution time in milliseconds accumulated by pg_stat_statements since statsReset.
	TotalTimeMs *float64 `json:"totalTimeMs,omitempty"`
	// Mean execution time in milliseconds.
	MeanTimeMs *float64 `json:"meanTimeMs,omitempty"`
	// Max execution time in milliseconds.
	MaxTimeMs *float64 `json:"maxTimeMs,omitempty"`
	// Rows returned or affected by the fingerprint, as reported by pg_stat_statements.
	Rows *int64 `json:"rows,omitempty"`
	// Database name resolved from pg_stat_statements.dbid when visible.
	Database string `json:"database"`
	// Database user name resolved from pg_stat_statements.userid when visible.
	User string `json:"user"`
	// Data source. M1 uses pg_stat_statements only.
	Source string `json:"source"`
	// Detail status. Expected values are available or unavailable.
	Status string `json:"status"`
	// Reason when status is unavailable, such as not_found or fingerprint_mismatch.
	UnavailableReason *string `json:"unavailableReason,omitempty"`
	// UTC timestamp reported by pg_stat_statements_info.stats_reset when available.
	StatsReset *string `json:"statsReset,omitempty"`
	// Backend collection timestamp in UTC.
	CollectedAt string `json:"collectedAt"`
	// Statistics accumulation scope. M1 uses pg_stat_statements_since_stats_reset.
	StatisticsScope string                         `json:"statisticsScope"`
	Explain         PostgresqlSQLExplainCapability `json:"explain"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlSQLFingerprintDetail instantiates a new PostgresqlSQLFingerprintDetail object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlSQLFingerprintDetail(queryId string, fingerprint string, database string, user string, source string, status string, collectedAt string, statisticsScope string, explain PostgresqlSQLExplainCapability) *PostgresqlSQLFingerprintDetail {
	this := PostgresqlSQLFingerprintDetail{}
	this.QueryId = queryId
	this.Fingerprint = fingerprint
	this.Database = database
	this.User = user
	this.Source = source
	this.Status = status
	this.CollectedAt = collectedAt
	this.StatisticsScope = statisticsScope
	this.Explain = explain
	return &this
}

// NewPostgresqlSQLFingerprintDetailWithDefaults instantiates a new PostgresqlSQLFingerprintDetail object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlSQLFingerprintDetailWithDefaults() *PostgresqlSQLFingerprintDetail {
	this := PostgresqlSQLFingerprintDetail{}
	return &this
}

// GetQueryId returns the QueryId field value.
func (o *PostgresqlSQLFingerprintDetail) GetQueryId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.QueryId
}

// GetQueryIdOk returns a tuple with the QueryId field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLFingerprintDetail) GetQueryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryId, true
}

// SetQueryId sets field value.
func (o *PostgresqlSQLFingerprintDetail) SetQueryId(v string) {
	o.QueryId = v
}

// GetFingerprint returns the Fingerprint field value.
func (o *PostgresqlSQLFingerprintDetail) GetFingerprint() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLFingerprintDetail) GetFingerprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fingerprint, true
}

// SetFingerprint sets field value.
func (o *PostgresqlSQLFingerprintDetail) SetFingerprint(v string) {
	o.Fingerprint = v
}

// GetQuerySummary returns the QuerySummary field value if set, zero value otherwise.
func (o *PostgresqlSQLFingerprintDetail) GetQuerySummary() string {
	if o == nil || o.QuerySummary == nil {
		var ret string
		return ret
	}
	return *o.QuerySummary
}

// GetQuerySummaryOk returns a tuple with the QuerySummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLFingerprintDetail) GetQuerySummaryOk() (*string, bool) {
	if o == nil || o.QuerySummary == nil {
		return nil, false
	}
	return o.QuerySummary, true
}

// HasQuerySummary returns a boolean if a field has been set.
func (o *PostgresqlSQLFingerprintDetail) HasQuerySummary() bool {
	return o != nil && o.QuerySummary != nil
}

// SetQuerySummary gets a reference to the given string and assigns it to the QuerySummary field.
func (o *PostgresqlSQLFingerprintDetail) SetQuerySummary(v string) {
	o.QuerySummary = &v
}

// GetCalls returns the Calls field value if set, zero value otherwise.
func (o *PostgresqlSQLFingerprintDetail) GetCalls() int64 {
	if o == nil || o.Calls == nil {
		var ret int64
		return ret
	}
	return *o.Calls
}

// GetCallsOk returns a tuple with the Calls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLFingerprintDetail) GetCallsOk() (*int64, bool) {
	if o == nil || o.Calls == nil {
		return nil, false
	}
	return o.Calls, true
}

// HasCalls returns a boolean if a field has been set.
func (o *PostgresqlSQLFingerprintDetail) HasCalls() bool {
	return o != nil && o.Calls != nil
}

// SetCalls gets a reference to the given int64 and assigns it to the Calls field.
func (o *PostgresqlSQLFingerprintDetail) SetCalls(v int64) {
	o.Calls = &v
}

// GetTotalTimeMs returns the TotalTimeMs field value if set, zero value otherwise.
func (o *PostgresqlSQLFingerprintDetail) GetTotalTimeMs() float64 {
	if o == nil || o.TotalTimeMs == nil {
		var ret float64
		return ret
	}
	return *o.TotalTimeMs
}

// GetTotalTimeMsOk returns a tuple with the TotalTimeMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLFingerprintDetail) GetTotalTimeMsOk() (*float64, bool) {
	if o == nil || o.TotalTimeMs == nil {
		return nil, false
	}
	return o.TotalTimeMs, true
}

// HasTotalTimeMs returns a boolean if a field has been set.
func (o *PostgresqlSQLFingerprintDetail) HasTotalTimeMs() bool {
	return o != nil && o.TotalTimeMs != nil
}

// SetTotalTimeMs gets a reference to the given float64 and assigns it to the TotalTimeMs field.
func (o *PostgresqlSQLFingerprintDetail) SetTotalTimeMs(v float64) {
	o.TotalTimeMs = &v
}

// GetMeanTimeMs returns the MeanTimeMs field value if set, zero value otherwise.
func (o *PostgresqlSQLFingerprintDetail) GetMeanTimeMs() float64 {
	if o == nil || o.MeanTimeMs == nil {
		var ret float64
		return ret
	}
	return *o.MeanTimeMs
}

// GetMeanTimeMsOk returns a tuple with the MeanTimeMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLFingerprintDetail) GetMeanTimeMsOk() (*float64, bool) {
	if o == nil || o.MeanTimeMs == nil {
		return nil, false
	}
	return o.MeanTimeMs, true
}

// HasMeanTimeMs returns a boolean if a field has been set.
func (o *PostgresqlSQLFingerprintDetail) HasMeanTimeMs() bool {
	return o != nil && o.MeanTimeMs != nil
}

// SetMeanTimeMs gets a reference to the given float64 and assigns it to the MeanTimeMs field.
func (o *PostgresqlSQLFingerprintDetail) SetMeanTimeMs(v float64) {
	o.MeanTimeMs = &v
}

// GetMaxTimeMs returns the MaxTimeMs field value if set, zero value otherwise.
func (o *PostgresqlSQLFingerprintDetail) GetMaxTimeMs() float64 {
	if o == nil || o.MaxTimeMs == nil {
		var ret float64
		return ret
	}
	return *o.MaxTimeMs
}

// GetMaxTimeMsOk returns a tuple with the MaxTimeMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLFingerprintDetail) GetMaxTimeMsOk() (*float64, bool) {
	if o == nil || o.MaxTimeMs == nil {
		return nil, false
	}
	return o.MaxTimeMs, true
}

// HasMaxTimeMs returns a boolean if a field has been set.
func (o *PostgresqlSQLFingerprintDetail) HasMaxTimeMs() bool {
	return o != nil && o.MaxTimeMs != nil
}

// SetMaxTimeMs gets a reference to the given float64 and assigns it to the MaxTimeMs field.
func (o *PostgresqlSQLFingerprintDetail) SetMaxTimeMs(v float64) {
	o.MaxTimeMs = &v
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *PostgresqlSQLFingerprintDetail) GetRows() int64 {
	if o == nil || o.Rows == nil {
		var ret int64
		return ret
	}
	return *o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLFingerprintDetail) GetRowsOk() (*int64, bool) {
	if o == nil || o.Rows == nil {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *PostgresqlSQLFingerprintDetail) HasRows() bool {
	return o != nil && o.Rows != nil
}

// SetRows gets a reference to the given int64 and assigns it to the Rows field.
func (o *PostgresqlSQLFingerprintDetail) SetRows(v int64) {
	o.Rows = &v
}

// GetDatabase returns the Database field value.
func (o *PostgresqlSQLFingerprintDetail) GetDatabase() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Database
}

// GetDatabaseOk returns a tuple with the Database field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLFingerprintDetail) GetDatabaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Database, true
}

// SetDatabase sets field value.
func (o *PostgresqlSQLFingerprintDetail) SetDatabase(v string) {
	o.Database = v
}

// GetUser returns the User field value.
func (o *PostgresqlSQLFingerprintDetail) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLFingerprintDetail) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value.
func (o *PostgresqlSQLFingerprintDetail) SetUser(v string) {
	o.User = v
}

// GetSource returns the Source field value.
func (o *PostgresqlSQLFingerprintDetail) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLFingerprintDetail) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *PostgresqlSQLFingerprintDetail) SetSource(v string) {
	o.Source = v
}

// GetStatus returns the Status field value.
func (o *PostgresqlSQLFingerprintDetail) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLFingerprintDetail) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *PostgresqlSQLFingerprintDetail) SetStatus(v string) {
	o.Status = v
}

// GetUnavailableReason returns the UnavailableReason field value if set, zero value otherwise.
func (o *PostgresqlSQLFingerprintDetail) GetUnavailableReason() string {
	if o == nil || o.UnavailableReason == nil {
		var ret string
		return ret
	}
	return *o.UnavailableReason
}

// GetUnavailableReasonOk returns a tuple with the UnavailableReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLFingerprintDetail) GetUnavailableReasonOk() (*string, bool) {
	if o == nil || o.UnavailableReason == nil {
		return nil, false
	}
	return o.UnavailableReason, true
}

// HasUnavailableReason returns a boolean if a field has been set.
func (o *PostgresqlSQLFingerprintDetail) HasUnavailableReason() bool {
	return o != nil && o.UnavailableReason != nil
}

// SetUnavailableReason gets a reference to the given string and assigns it to the UnavailableReason field.
func (o *PostgresqlSQLFingerprintDetail) SetUnavailableReason(v string) {
	o.UnavailableReason = &v
}

// GetStatsReset returns the StatsReset field value if set, zero value otherwise.
func (o *PostgresqlSQLFingerprintDetail) GetStatsReset() string {
	if o == nil || o.StatsReset == nil {
		var ret string
		return ret
	}
	return *o.StatsReset
}

// GetStatsResetOk returns a tuple with the StatsReset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLFingerprintDetail) GetStatsResetOk() (*string, bool) {
	if o == nil || o.StatsReset == nil {
		return nil, false
	}
	return o.StatsReset, true
}

// HasStatsReset returns a boolean if a field has been set.
func (o *PostgresqlSQLFingerprintDetail) HasStatsReset() bool {
	return o != nil && o.StatsReset != nil
}

// SetStatsReset gets a reference to the given string and assigns it to the StatsReset field.
func (o *PostgresqlSQLFingerprintDetail) SetStatsReset(v string) {
	o.StatsReset = &v
}

// GetCollectedAt returns the CollectedAt field value.
func (o *PostgresqlSQLFingerprintDetail) GetCollectedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CollectedAt
}

// GetCollectedAtOk returns a tuple with the CollectedAt field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLFingerprintDetail) GetCollectedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectedAt, true
}

// SetCollectedAt sets field value.
func (o *PostgresqlSQLFingerprintDetail) SetCollectedAt(v string) {
	o.CollectedAt = v
}

// GetStatisticsScope returns the StatisticsScope field value.
func (o *PostgresqlSQLFingerprintDetail) GetStatisticsScope() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.StatisticsScope
}

// GetStatisticsScopeOk returns a tuple with the StatisticsScope field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLFingerprintDetail) GetStatisticsScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatisticsScope, true
}

// SetStatisticsScope sets field value.
func (o *PostgresqlSQLFingerprintDetail) SetStatisticsScope(v string) {
	o.StatisticsScope = v
}

// GetExplain returns the Explain field value.
func (o *PostgresqlSQLFingerprintDetail) GetExplain() PostgresqlSQLExplainCapability {
	if o == nil {
		var ret PostgresqlSQLExplainCapability
		return ret
	}
	return o.Explain
}

// GetExplainOk returns a tuple with the Explain field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLFingerprintDetail) GetExplainOk() (*PostgresqlSQLExplainCapability, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Explain, true
}

// SetExplain sets field value.
func (o *PostgresqlSQLFingerprintDetail) SetExplain(v PostgresqlSQLExplainCapability) {
	o.Explain = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlSQLFingerprintDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["queryID"] = o.QueryId
	toSerialize["fingerprint"] = o.Fingerprint
	if o.QuerySummary != nil {
		toSerialize["querySummary"] = o.QuerySummary
	}
	if o.Calls != nil {
		toSerialize["calls"] = o.Calls
	}
	if o.TotalTimeMs != nil {
		toSerialize["totalTimeMs"] = o.TotalTimeMs
	}
	if o.MeanTimeMs != nil {
		toSerialize["meanTimeMs"] = o.MeanTimeMs
	}
	if o.MaxTimeMs != nil {
		toSerialize["maxTimeMs"] = o.MaxTimeMs
	}
	if o.Rows != nil {
		toSerialize["rows"] = o.Rows
	}
	toSerialize["database"] = o.Database
	toSerialize["user"] = o.User
	toSerialize["source"] = o.Source
	toSerialize["status"] = o.Status
	if o.UnavailableReason != nil {
		toSerialize["unavailableReason"] = o.UnavailableReason
	}
	if o.StatsReset != nil {
		toSerialize["statsReset"] = o.StatsReset
	}
	toSerialize["collectedAt"] = o.CollectedAt
	toSerialize["statisticsScope"] = o.StatisticsScope
	toSerialize["explain"] = o.Explain

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostgresqlSQLFingerprintDetail) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		QueryId           *string                         `json:"queryID"`
		Fingerprint       *string                         `json:"fingerprint"`
		QuerySummary      *string                         `json:"querySummary,omitempty"`
		Calls             *int64                          `json:"calls,omitempty"`
		TotalTimeMs       *float64                        `json:"totalTimeMs,omitempty"`
		MeanTimeMs        *float64                        `json:"meanTimeMs,omitempty"`
		MaxTimeMs         *float64                        `json:"maxTimeMs,omitempty"`
		Rows              *int64                          `json:"rows,omitempty"`
		Database          *string                         `json:"database"`
		User              *string                         `json:"user"`
		Source            *string                         `json:"source"`
		Status            *string                         `json:"status"`
		UnavailableReason *string                         `json:"unavailableReason,omitempty"`
		StatsReset        *string                         `json:"statsReset,omitempty"`
		CollectedAt       *string                         `json:"collectedAt"`
		StatisticsScope   *string                         `json:"statisticsScope"`
		Explain           *PostgresqlSQLExplainCapability `json:"explain"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.QueryId == nil {
		return fmt.Errorf("required field queryID missing")
	}
	if all.Fingerprint == nil {
		return fmt.Errorf("required field fingerprint missing")
	}
	if all.Database == nil {
		return fmt.Errorf("required field database missing")
	}
	if all.User == nil {
		return fmt.Errorf("required field user missing")
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.CollectedAt == nil {
		return fmt.Errorf("required field collectedAt missing")
	}
	if all.StatisticsScope == nil {
		return fmt.Errorf("required field statisticsScope missing")
	}
	if all.Explain == nil {
		return fmt.Errorf("required field explain missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"queryID", "fingerprint", "querySummary", "calls", "totalTimeMs", "meanTimeMs", "maxTimeMs", "rows", "database", "user", "source", "status", "unavailableReason", "statsReset", "collectedAt", "statisticsScope", "explain"})
	} else {
		return err
	}

	hasInvalidField := false
	o.QueryId = *all.QueryId
	o.Fingerprint = *all.Fingerprint
	o.QuerySummary = all.QuerySummary
	o.Calls = all.Calls
	o.TotalTimeMs = all.TotalTimeMs
	o.MeanTimeMs = all.MeanTimeMs
	o.MaxTimeMs = all.MaxTimeMs
	o.Rows = all.Rows
	o.Database = *all.Database
	o.User = *all.User
	o.Source = *all.Source
	o.Status = *all.Status
	o.UnavailableReason = all.UnavailableReason
	o.StatsReset = all.StatsReset
	o.CollectedAt = *all.CollectedAt
	o.StatisticsScope = *all.StatisticsScope
	if all.Explain.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Explain = *all.Explain

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
