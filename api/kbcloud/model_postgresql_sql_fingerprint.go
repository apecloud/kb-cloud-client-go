// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlSQLFingerprint struct {
	// PostgreSQL pg_stat_statements queryid represented as a string.
	QueryId string `json:"queryID"`
	// Stable SQL fingerprint identifier for UI grouping. Currently aligned with PostgreSQL pg_stat_statements queryid.
	Fingerprint string `json:"fingerprint"`
	// Redacted SQL summary. Full raw SQL is intentionally not returned.
	QuerySummary string `json:"querySummary"`
	// Number of executions accumulated by pg_stat_statements since statsReset.
	Calls int64 `json:"calls"`
	// Total execution time in milliseconds accumulated by pg_stat_statements since statsReset.
	TotalTimeMs float64 `json:"totalTimeMs"`
	// Mean execution time in milliseconds.
	MeanTimeMs float64 `json:"meanTimeMs"`
	// Max execution time in milliseconds.
	MaxTimeMs float64 `json:"maxTimeMs"`
	// Rows returned or affected by the fingerprint, as reported by pg_stat_statements.
	Rows int64 `json:"rows"`
	// Database name resolved from pg_stat_statements.dbid when visible.
	Database string `json:"database"`
	// Database user name resolved from pg_stat_statements.userid when visible.
	User string `json:"user"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlSQLFingerprint instantiates a new PostgresqlSQLFingerprint object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlSQLFingerprint(queryId string, fingerprint string, querySummary string, calls int64, totalTimeMs float64, meanTimeMs float64, maxTimeMs float64, rows int64, database string, user string) *PostgresqlSQLFingerprint {
	this := PostgresqlSQLFingerprint{}
	this.QueryId = queryId
	this.Fingerprint = fingerprint
	this.QuerySummary = querySummary
	this.Calls = calls
	this.TotalTimeMs = totalTimeMs
	this.MeanTimeMs = meanTimeMs
	this.MaxTimeMs = maxTimeMs
	this.Rows = rows
	this.Database = database
	this.User = user
	return &this
}

// NewPostgresqlSQLFingerprintWithDefaults instantiates a new PostgresqlSQLFingerprint object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlSQLFingerprintWithDefaults() *PostgresqlSQLFingerprint {
	this := PostgresqlSQLFingerprint{}
	return &this
}

// GetQueryId returns the QueryId field value.
func (o *PostgresqlSQLFingerprint) GetQueryId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.QueryId
}

// GetQueryIdOk returns a tuple with the QueryId field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLFingerprint) GetQueryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryId, true
}

// SetQueryId sets field value.
func (o *PostgresqlSQLFingerprint) SetQueryId(v string) {
	o.QueryId = v
}

// GetFingerprint returns the Fingerprint field value.
func (o *PostgresqlSQLFingerprint) GetFingerprint() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLFingerprint) GetFingerprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fingerprint, true
}

// SetFingerprint sets field value.
func (o *PostgresqlSQLFingerprint) SetFingerprint(v string) {
	o.Fingerprint = v
}

// GetQuerySummary returns the QuerySummary field value.
func (o *PostgresqlSQLFingerprint) GetQuerySummary() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.QuerySummary
}

// GetQuerySummaryOk returns a tuple with the QuerySummary field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLFingerprint) GetQuerySummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QuerySummary, true
}

// SetQuerySummary sets field value.
func (o *PostgresqlSQLFingerprint) SetQuerySummary(v string) {
	o.QuerySummary = v
}

// GetCalls returns the Calls field value.
func (o *PostgresqlSQLFingerprint) GetCalls() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Calls
}

// GetCallsOk returns a tuple with the Calls field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLFingerprint) GetCallsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Calls, true
}

// SetCalls sets field value.
func (o *PostgresqlSQLFingerprint) SetCalls(v int64) {
	o.Calls = v
}

// GetTotalTimeMs returns the TotalTimeMs field value.
func (o *PostgresqlSQLFingerprint) GetTotalTimeMs() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.TotalTimeMs
}

// GetTotalTimeMsOk returns a tuple with the TotalTimeMs field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLFingerprint) GetTotalTimeMsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalTimeMs, true
}

// SetTotalTimeMs sets field value.
func (o *PostgresqlSQLFingerprint) SetTotalTimeMs(v float64) {
	o.TotalTimeMs = v
}

// GetMeanTimeMs returns the MeanTimeMs field value.
func (o *PostgresqlSQLFingerprint) GetMeanTimeMs() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.MeanTimeMs
}

// GetMeanTimeMsOk returns a tuple with the MeanTimeMs field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLFingerprint) GetMeanTimeMsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MeanTimeMs, true
}

// SetMeanTimeMs sets field value.
func (o *PostgresqlSQLFingerprint) SetMeanTimeMs(v float64) {
	o.MeanTimeMs = v
}

// GetMaxTimeMs returns the MaxTimeMs field value.
func (o *PostgresqlSQLFingerprint) GetMaxTimeMs() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.MaxTimeMs
}

// GetMaxTimeMsOk returns a tuple with the MaxTimeMs field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLFingerprint) GetMaxTimeMsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxTimeMs, true
}

// SetMaxTimeMs sets field value.
func (o *PostgresqlSQLFingerprint) SetMaxTimeMs(v float64) {
	o.MaxTimeMs = v
}

// GetRows returns the Rows field value.
func (o *PostgresqlSQLFingerprint) GetRows() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLFingerprint) GetRowsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rows, true
}

// SetRows sets field value.
func (o *PostgresqlSQLFingerprint) SetRows(v int64) {
	o.Rows = v
}

// GetDatabase returns the Database field value.
func (o *PostgresqlSQLFingerprint) GetDatabase() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Database
}

// GetDatabaseOk returns a tuple with the Database field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLFingerprint) GetDatabaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Database, true
}

// SetDatabase sets field value.
func (o *PostgresqlSQLFingerprint) SetDatabase(v string) {
	o.Database = v
}

// GetUser returns the User field value.
func (o *PostgresqlSQLFingerprint) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLFingerprint) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value.
func (o *PostgresqlSQLFingerprint) SetUser(v string) {
	o.User = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlSQLFingerprint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["queryID"] = o.QueryId
	toSerialize["fingerprint"] = o.Fingerprint
	toSerialize["querySummary"] = o.QuerySummary
	toSerialize["calls"] = o.Calls
	toSerialize["totalTimeMs"] = o.TotalTimeMs
	toSerialize["meanTimeMs"] = o.MeanTimeMs
	toSerialize["maxTimeMs"] = o.MaxTimeMs
	toSerialize["rows"] = o.Rows
	toSerialize["database"] = o.Database
	toSerialize["user"] = o.User

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostgresqlSQLFingerprint) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		QueryId      *string  `json:"queryID"`
		Fingerprint  *string  `json:"fingerprint"`
		QuerySummary *string  `json:"querySummary"`
		Calls        *int64   `json:"calls"`
		TotalTimeMs  *float64 `json:"totalTimeMs"`
		MeanTimeMs   *float64 `json:"meanTimeMs"`
		MaxTimeMs    *float64 `json:"maxTimeMs"`
		Rows         *int64   `json:"rows"`
		Database     *string  `json:"database"`
		User         *string  `json:"user"`
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
	if all.QuerySummary == nil {
		return fmt.Errorf("required field querySummary missing")
	}
	if all.Calls == nil {
		return fmt.Errorf("required field calls missing")
	}
	if all.TotalTimeMs == nil {
		return fmt.Errorf("required field totalTimeMs missing")
	}
	if all.MeanTimeMs == nil {
		return fmt.Errorf("required field meanTimeMs missing")
	}
	if all.MaxTimeMs == nil {
		return fmt.Errorf("required field maxTimeMs missing")
	}
	if all.Rows == nil {
		return fmt.Errorf("required field rows missing")
	}
	if all.Database == nil {
		return fmt.Errorf("required field database missing")
	}
	if all.User == nil {
		return fmt.Errorf("required field user missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"queryID", "fingerprint", "querySummary", "calls", "totalTimeMs", "meanTimeMs", "maxTimeMs", "rows", "database", "user"})
	} else {
		return err
	}
	o.QueryId = *all.QueryId
	o.Fingerprint = *all.Fingerprint
	o.QuerySummary = *all.QuerySummary
	o.Calls = *all.Calls
	o.TotalTimeMs = *all.TotalTimeMs
	o.MeanTimeMs = *all.MeanTimeMs
	o.MaxTimeMs = *all.MaxTimeMs
	o.Rows = *all.Rows
	o.Database = *all.Database
	o.User = *all.User

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
