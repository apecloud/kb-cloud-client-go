// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlSQLFingerprintExplainResponse struct {
	// PostgreSQL pg_stat_statements queryid represented as a string.
	QueryId string `json:"queryID"`
	// Stable SQL fingerprint identifier for UI grouping. Currently aligned with PostgreSQL pg_stat_statements queryid.
	Fingerprint string `json:"fingerprint"`
	// Database name from the ranking row.
	Database string `json:"database"`
	// Database user from the ranking row.
	User string `json:"user"`
	// Top-level identity from the matched ranking row.
	TopLevel bool `json:"topLevel"`
	// Server-side source used to resolve the parameterized SQL statement. Raw SQL is not returned.
	StatementSource *string `json:"statementSource,omitempty"`
	// Timestamp when the parameterized SQL statement was resolved.
	StatementResolvedAt *string `json:"statementResolvedAt,omitempty"`
	// Redacted SQL summary from the matched ranking row. The resolved statement is intentionally not returned.
	QuerySummary *string                      `json:"querySummary,omitempty"`
	PlanMode     DmsExecutionPlanPlanningMode `json:"planMode"`
	// Whether the estimated plan was produced without concrete parameter values.
	Parameterized bool `json:"parameterized"`
	// Number of parameters when DMS can determine it. It may be omitted for PostgreSQL generic-plan paths that do not expose the count.
	ParameterCount common.NullableInt64   `json:"parameterCount,omitempty"`
	ExplainResult  DmsExecutionPlanResult `json:"explainResult"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlSQLFingerprintExplainResponse instantiates a new PostgresqlSQLFingerprintExplainResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlSQLFingerprintExplainResponse(queryId string, fingerprint string, database string, user string, topLevel bool, planMode DmsExecutionPlanPlanningMode, parameterized bool, explainResult DmsExecutionPlanResult) *PostgresqlSQLFingerprintExplainResponse {
	this := PostgresqlSQLFingerprintExplainResponse{}
	this.QueryId = queryId
	this.Fingerprint = fingerprint
	this.Database = database
	this.User = user
	this.TopLevel = topLevel
	this.PlanMode = planMode
	this.Parameterized = parameterized
	this.ExplainResult = explainResult
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

// GetTopLevel returns the TopLevel field value.
func (o *PostgresqlSQLFingerprintExplainResponse) GetTopLevel() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.TopLevel
}

// GetTopLevelOk returns a tuple with the TopLevel field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLFingerprintExplainResponse) GetTopLevelOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TopLevel, true
}

// SetTopLevel sets field value.
func (o *PostgresqlSQLFingerprintExplainResponse) SetTopLevel(v bool) {
	o.TopLevel = v
}

// GetStatementSource returns the StatementSource field value if set, zero value otherwise.
func (o *PostgresqlSQLFingerprintExplainResponse) GetStatementSource() string {
	if o == nil || o.StatementSource == nil {
		var ret string
		return ret
	}
	return *o.StatementSource
}

// GetStatementSourceOk returns a tuple with the StatementSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLFingerprintExplainResponse) GetStatementSourceOk() (*string, bool) {
	if o == nil || o.StatementSource == nil {
		return nil, false
	}
	return o.StatementSource, true
}

// HasStatementSource returns a boolean if a field has been set.
func (o *PostgresqlSQLFingerprintExplainResponse) HasStatementSource() bool {
	return o != nil && o.StatementSource != nil
}

// SetStatementSource gets a reference to the given string and assigns it to the StatementSource field.
func (o *PostgresqlSQLFingerprintExplainResponse) SetStatementSource(v string) {
	o.StatementSource = &v
}

// GetStatementResolvedAt returns the StatementResolvedAt field value if set, zero value otherwise.
func (o *PostgresqlSQLFingerprintExplainResponse) GetStatementResolvedAt() string {
	if o == nil || o.StatementResolvedAt == nil {
		var ret string
		return ret
	}
	return *o.StatementResolvedAt
}

// GetStatementResolvedAtOk returns a tuple with the StatementResolvedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLFingerprintExplainResponse) GetStatementResolvedAtOk() (*string, bool) {
	if o == nil || o.StatementResolvedAt == nil {
		return nil, false
	}
	return o.StatementResolvedAt, true
}

// HasStatementResolvedAt returns a boolean if a field has been set.
func (o *PostgresqlSQLFingerprintExplainResponse) HasStatementResolvedAt() bool {
	return o != nil && o.StatementResolvedAt != nil
}

// SetStatementResolvedAt gets a reference to the given string and assigns it to the StatementResolvedAt field.
func (o *PostgresqlSQLFingerprintExplainResponse) SetStatementResolvedAt(v string) {
	o.StatementResolvedAt = &v
}

// GetQuerySummary returns the QuerySummary field value if set, zero value otherwise.
func (o *PostgresqlSQLFingerprintExplainResponse) GetQuerySummary() string {
	if o == nil || o.QuerySummary == nil {
		var ret string
		return ret
	}
	return *o.QuerySummary
}

// GetQuerySummaryOk returns a tuple with the QuerySummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLFingerprintExplainResponse) GetQuerySummaryOk() (*string, bool) {
	if o == nil || o.QuerySummary == nil {
		return nil, false
	}
	return o.QuerySummary, true
}

// HasQuerySummary returns a boolean if a field has been set.
func (o *PostgresqlSQLFingerprintExplainResponse) HasQuerySummary() bool {
	return o != nil && o.QuerySummary != nil
}

// SetQuerySummary gets a reference to the given string and assigns it to the QuerySummary field.
func (o *PostgresqlSQLFingerprintExplainResponse) SetQuerySummary(v string) {
	o.QuerySummary = &v
}

// GetPlanMode returns the PlanMode field value.
func (o *PostgresqlSQLFingerprintExplainResponse) GetPlanMode() DmsExecutionPlanPlanningMode {
	if o == nil {
		var ret DmsExecutionPlanPlanningMode
		return ret
	}
	return o.PlanMode
}

// GetPlanModeOk returns a tuple with the PlanMode field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLFingerprintExplainResponse) GetPlanModeOk() (*DmsExecutionPlanPlanningMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlanMode, true
}

// SetPlanMode sets field value.
func (o *PostgresqlSQLFingerprintExplainResponse) SetPlanMode(v DmsExecutionPlanPlanningMode) {
	o.PlanMode = v
}

// GetParameterized returns the Parameterized field value.
func (o *PostgresqlSQLFingerprintExplainResponse) GetParameterized() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Parameterized
}

// GetParameterizedOk returns a tuple with the Parameterized field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLFingerprintExplainResponse) GetParameterizedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parameterized, true
}

// SetParameterized sets field value.
func (o *PostgresqlSQLFingerprintExplainResponse) SetParameterized(v bool) {
	o.Parameterized = v
}

// GetParameterCount returns the ParameterCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostgresqlSQLFingerprintExplainResponse) GetParameterCount() int64 {
	if o == nil || o.ParameterCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ParameterCount.Get()
}

// GetParameterCountOk returns a tuple with the ParameterCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PostgresqlSQLFingerprintExplainResponse) GetParameterCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParameterCount.Get(), o.ParameterCount.IsSet()
}

// HasParameterCount returns a boolean if a field has been set.
func (o *PostgresqlSQLFingerprintExplainResponse) HasParameterCount() bool {
	return o != nil && o.ParameterCount.IsSet()
}

// SetParameterCount gets a reference to the given common.NullableInt64 and assigns it to the ParameterCount field.
func (o *PostgresqlSQLFingerprintExplainResponse) SetParameterCount(v int64) {
	o.ParameterCount.Set(&v)
}

// SetParameterCountNil sets the value for ParameterCount to be an explicit nil.
func (o *PostgresqlSQLFingerprintExplainResponse) SetParameterCountNil() {
	o.ParameterCount.Set(nil)
}

// UnsetParameterCount ensures that no value is present for ParameterCount, not even an explicit nil.
func (o *PostgresqlSQLFingerprintExplainResponse) UnsetParameterCount() {
	o.ParameterCount.Unset()
}

// GetExplainResult returns the ExplainResult field value.
func (o *PostgresqlSQLFingerprintExplainResponse) GetExplainResult() DmsExecutionPlanResult {
	if o == nil {
		var ret DmsExecutionPlanResult
		return ret
	}
	return o.ExplainResult
}

// GetExplainResultOk returns a tuple with the ExplainResult field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLFingerprintExplainResponse) GetExplainResultOk() (*DmsExecutionPlanResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExplainResult, true
}

// SetExplainResult sets field value.
func (o *PostgresqlSQLFingerprintExplainResponse) SetExplainResult(v DmsExecutionPlanResult) {
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
	toSerialize["topLevel"] = o.TopLevel
	if o.StatementSource != nil {
		toSerialize["statementSource"] = o.StatementSource
	}
	if o.StatementResolvedAt != nil {
		toSerialize["statementResolvedAt"] = o.StatementResolvedAt
	}
	if o.QuerySummary != nil {
		toSerialize["querySummary"] = o.QuerySummary
	}
	toSerialize["planMode"] = o.PlanMode
	toSerialize["parameterized"] = o.Parameterized
	if o.ParameterCount.IsSet() {
		toSerialize["parameterCount"] = o.ParameterCount.Get()
	}
	toSerialize["explainResult"] = o.ExplainResult

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostgresqlSQLFingerprintExplainResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		QueryId             *string                       `json:"queryID"`
		Fingerprint         *string                       `json:"fingerprint"`
		Database            *string                       `json:"database"`
		User                *string                       `json:"user"`
		TopLevel            *bool                         `json:"topLevel"`
		StatementSource     *string                       `json:"statementSource,omitempty"`
		StatementResolvedAt *string                       `json:"statementResolvedAt,omitempty"`
		QuerySummary        *string                       `json:"querySummary,omitempty"`
		PlanMode            *DmsExecutionPlanPlanningMode `json:"planMode"`
		Parameterized       *bool                         `json:"parameterized"`
		ParameterCount      common.NullableInt64          `json:"parameterCount,omitempty"`
		ExplainResult       *DmsExecutionPlanResult       `json:"explainResult"`
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
	if all.TopLevel == nil {
		return fmt.Errorf("required field topLevel missing")
	}
	if all.PlanMode == nil {
		return fmt.Errorf("required field planMode missing")
	}
	if all.Parameterized == nil {
		return fmt.Errorf("required field parameterized missing")
	}
	if all.ExplainResult == nil {
		return fmt.Errorf("required field explainResult missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"queryID", "fingerprint", "database", "user", "topLevel", "statementSource", "statementResolvedAt", "querySummary", "planMode", "parameterized", "parameterCount", "explainResult"})
	} else {
		return err
	}

	hasInvalidField := false
	o.QueryId = *all.QueryId
	o.Fingerprint = *all.Fingerprint
	o.Database = *all.Database
	o.User = *all.User
	o.TopLevel = *all.TopLevel
	o.StatementSource = all.StatementSource
	o.StatementResolvedAt = all.StatementResolvedAt
	o.QuerySummary = all.QuerySummary
	if !all.PlanMode.IsValid() {
		hasInvalidField = true
	} else {
		o.PlanMode = *all.PlanMode
	}
	o.Parameterized = *all.Parameterized
	o.ParameterCount = all.ParameterCount
	if all.ExplainResult.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ExplainResult = *all.ExplainResult

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
