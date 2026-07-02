// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlSQLFingerprintExplainResponse struct {
	// PostgreSQL pg_stat_statements queryid represented as a string.
	QueryId string `json:"queryID"`
	// Stable SQL fingerprint identifier for UI grouping. M1 uses queryID.
	Fingerprint string `json:"fingerprint"`
	// Database name from the ranking row.
	Database string `json:"database"`
	// Database user from the ranking row.
	User string `json:"user"`
	// Terminal EXPLAIN status. Expected values are success, unavailable, unsupported, permission_denied, and failed.
	Status string `json:"status"`
	// Reason when status is not success, such as no_explainable_sql_sample, non_select, permission_denied, or dms_explain_failed.
	UnavailableReason *string `json:"unavailableReason,omitempty"`
	// Server-side sample source when available. Raw SQL is not returned.
	SampleSource *string `json:"sampleSource,omitempty"`
	// Sample collection timestamp when available.
	SampleCollectedAt *string `json:"sampleCollectedAt,omitempty"`
	// Redacted sample SQL summary when available. Full raw SQL is intentionally not returned.
	SampleSqlSummary *string `json:"sampleSQLSummary,omitempty"`
	// DMS SQLExplainV1 result when status is success.
	ExplainResult interface{} `json:"explainResult,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlSQLFingerprintExplainResponse instantiates a new PostgresqlSQLFingerprintExplainResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlSQLFingerprintExplainResponse(queryId string, fingerprint string, database string, user string, status string) *PostgresqlSQLFingerprintExplainResponse {
	this := PostgresqlSQLFingerprintExplainResponse{}
	this.QueryId = queryId
	this.Fingerprint = fingerprint
	this.Database = database
	this.User = user
	this.Status = status
	return &this
}

// NewPostgresqlSQLFingerprintExplainResponseWithDefaults instantiates a new PostgresqlSQLFingerprintExplainResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlSQLFingerprintExplainResponseWithDefaults() *PostgresqlSQLFingerprintExplainResponse {
	this := PostgresqlSQLFingerprintExplainResponse{}
	return &this
}

// GetQueryId returns the QueryId field value.
func (o *PostgresqlSQLFingerprintExplainResponse) GetQueryId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.QueryId
}

// GetQueryIdOk returns a tuple with the QueryId field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLFingerprintExplainResponse) GetQueryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryId, true
}

// SetQueryId sets field value.
func (o *PostgresqlSQLFingerprintExplainResponse) SetQueryId(v string) {
	o.QueryId = v
}

// GetFingerprint returns the Fingerprint field value.
func (o *PostgresqlSQLFingerprintExplainResponse) GetFingerprint() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLFingerprintExplainResponse) GetFingerprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fingerprint, true
}

// SetFingerprint sets field value.
func (o *PostgresqlSQLFingerprintExplainResponse) SetFingerprint(v string) {
	o.Fingerprint = v
}

// GetDatabase returns the Database field value.
func (o *PostgresqlSQLFingerprintExplainResponse) GetDatabase() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Database
}

// GetDatabaseOk returns a tuple with the Database field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLFingerprintExplainResponse) GetDatabaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Database, true
}

// SetDatabase sets field value.
func (o *PostgresqlSQLFingerprintExplainResponse) SetDatabase(v string) {
	o.Database = v
}

// GetUser returns the User field value.
func (o *PostgresqlSQLFingerprintExplainResponse) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLFingerprintExplainResponse) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value.
func (o *PostgresqlSQLFingerprintExplainResponse) SetUser(v string) {
	o.User = v
}

// GetStatus returns the Status field value.
func (o *PostgresqlSQLFingerprintExplainResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLFingerprintExplainResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *PostgresqlSQLFingerprintExplainResponse) SetStatus(v string) {
	o.Status = v
}

// GetUnavailableReason returns the UnavailableReason field value if set, zero value otherwise.
func (o *PostgresqlSQLFingerprintExplainResponse) GetUnavailableReason() string {
	if o == nil || o.UnavailableReason == nil {
		var ret string
		return ret
	}
	return *o.UnavailableReason
}

// GetUnavailableReasonOk returns a tuple with the UnavailableReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLFingerprintExplainResponse) GetUnavailableReasonOk() (*string, bool) {
	if o == nil || o.UnavailableReason == nil {
		return nil, false
	}
	return o.UnavailableReason, true
}

// HasUnavailableReason returns a boolean if a field has been set.
func (o *PostgresqlSQLFingerprintExplainResponse) HasUnavailableReason() bool {
	return o != nil && o.UnavailableReason != nil
}

// SetUnavailableReason gets a reference to the given string and assigns it to the UnavailableReason field.
func (o *PostgresqlSQLFingerprintExplainResponse) SetUnavailableReason(v string) {
	o.UnavailableReason = &v
}

// GetSampleSource returns the SampleSource field value if set, zero value otherwise.
func (o *PostgresqlSQLFingerprintExplainResponse) GetSampleSource() string {
	if o == nil || o.SampleSource == nil {
		var ret string
		return ret
	}
	return *o.SampleSource
}

// GetSampleSourceOk returns a tuple with the SampleSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLFingerprintExplainResponse) GetSampleSourceOk() (*string, bool) {
	if o == nil || o.SampleSource == nil {
		return nil, false
	}
	return o.SampleSource, true
}

// HasSampleSource returns a boolean if a field has been set.
func (o *PostgresqlSQLFingerprintExplainResponse) HasSampleSource() bool {
	return o != nil && o.SampleSource != nil
}

// SetSampleSource gets a reference to the given string and assigns it to the SampleSource field.
func (o *PostgresqlSQLFingerprintExplainResponse) SetSampleSource(v string) {
	o.SampleSource = &v
}

// GetSampleCollectedAt returns the SampleCollectedAt field value if set, zero value otherwise.
func (o *PostgresqlSQLFingerprintExplainResponse) GetSampleCollectedAt() string {
	if o == nil || o.SampleCollectedAt == nil {
		var ret string
		return ret
	}
	return *o.SampleCollectedAt
}

// GetSampleCollectedAtOk returns a tuple with the SampleCollectedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLFingerprintExplainResponse) GetSampleCollectedAtOk() (*string, bool) {
	if o == nil || o.SampleCollectedAt == nil {
		return nil, false
	}
	return o.SampleCollectedAt, true
}

// HasSampleCollectedAt returns a boolean if a field has been set.
func (o *PostgresqlSQLFingerprintExplainResponse) HasSampleCollectedAt() bool {
	return o != nil && o.SampleCollectedAt != nil
}

// SetSampleCollectedAt gets a reference to the given string and assigns it to the SampleCollectedAt field.
func (o *PostgresqlSQLFingerprintExplainResponse) SetSampleCollectedAt(v string) {
	o.SampleCollectedAt = &v
}

// GetSampleSqlSummary returns the SampleSqlSummary field value if set, zero value otherwise.
func (o *PostgresqlSQLFingerprintExplainResponse) GetSampleSqlSummary() string {
	if o == nil || o.SampleSqlSummary == nil {
		var ret string
		return ret
	}
	return *o.SampleSqlSummary
}

// GetSampleSqlSummaryOk returns a tuple with the SampleSqlSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLFingerprintExplainResponse) GetSampleSqlSummaryOk() (*string, bool) {
	if o == nil || o.SampleSqlSummary == nil {
		return nil, false
	}
	return o.SampleSqlSummary, true
}

// HasSampleSqlSummary returns a boolean if a field has been set.
func (o *PostgresqlSQLFingerprintExplainResponse) HasSampleSqlSummary() bool {
	return o != nil && o.SampleSqlSummary != nil
}

// SetSampleSqlSummary gets a reference to the given string and assigns it to the SampleSqlSummary field.
func (o *PostgresqlSQLFingerprintExplainResponse) SetSampleSqlSummary(v string) {
	o.SampleSqlSummary = &v
}

// GetExplainResult returns the ExplainResult field value if set, zero value otherwise.
func (o *PostgresqlSQLFingerprintExplainResponse) GetExplainResult() interface{} {
	if o == nil || o.ExplainResult == nil {
		var ret interface{}
		return ret
	}
	return o.ExplainResult
}

// GetExplainResultOk returns a tuple with the ExplainResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLFingerprintExplainResponse) GetExplainResultOk() (*interface{}, bool) {
	if o == nil || o.ExplainResult == nil {
		return nil, false
	}
	return &o.ExplainResult, true
}

// HasExplainResult returns a boolean if a field has been set.
func (o *PostgresqlSQLFingerprintExplainResponse) HasExplainResult() bool {
	return o != nil && o.ExplainResult != nil
}

// SetExplainResult gets a reference to the given interface{} and assigns it to the ExplainResult field.
func (o *PostgresqlSQLFingerprintExplainResponse) SetExplainResult(v interface{}) {
	o.ExplainResult = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlSQLFingerprintExplainResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["queryID"] = o.QueryId
	toSerialize["fingerprint"] = o.Fingerprint
	toSerialize["database"] = o.Database
	toSerialize["user"] = o.User
	toSerialize["status"] = o.Status
	if o.UnavailableReason != nil {
		toSerialize["unavailableReason"] = o.UnavailableReason
	}
	if o.SampleSource != nil {
		toSerialize["sampleSource"] = o.SampleSource
	}
	if o.SampleCollectedAt != nil {
		toSerialize["sampleCollectedAt"] = o.SampleCollectedAt
	}
	if o.SampleSqlSummary != nil {
		toSerialize["sampleSQLSummary"] = o.SampleSqlSummary
	}
	if o.ExplainResult != nil {
		toSerialize["explainResult"] = o.ExplainResult
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostgresqlSQLFingerprintExplainResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		QueryId           *string     `json:"queryID"`
		Fingerprint       *string     `json:"fingerprint"`
		Database          *string     `json:"database"`
		User              *string     `json:"user"`
		Status            *string     `json:"status"`
		UnavailableReason *string     `json:"unavailableReason,omitempty"`
		SampleSource      *string     `json:"sampleSource,omitempty"`
		SampleCollectedAt *string     `json:"sampleCollectedAt,omitempty"`
		SampleSqlSummary  *string     `json:"sampleSQLSummary,omitempty"`
		ExplainResult     interface{} `json:"explainResult,omitempty"`
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
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"queryID", "fingerprint", "database", "user", "status", "unavailableReason", "sampleSource", "sampleCollectedAt", "sampleSQLSummary", "explainResult"})
	} else {
		return err
	}
	o.QueryId = *all.QueryId
	o.Fingerprint = *all.Fingerprint
	o.Database = *all.Database
	o.User = *all.User
	o.Status = *all.Status
	o.UnavailableReason = all.UnavailableReason
	o.SampleSource = all.SampleSource
	o.SampleCollectedAt = all.SampleCollectedAt
	o.SampleSqlSummary = all.SampleSqlSummary
	o.ExplainResult = all.ExplainResult

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
