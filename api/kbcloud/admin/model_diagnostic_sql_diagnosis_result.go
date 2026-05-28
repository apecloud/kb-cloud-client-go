// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticSQLDiagnosisResult PostgreSQL read-only SQL diagnosis result.
type DiagnosticSQLDiagnosisResult struct {
	// SQL diagnosis result status.
	Status DiagnosticSQLDiagnosisStatus `json:"status"`
	// Database used by the EXPLAIN request.
	Database *string `json:"database,omitempty"`
	// Redacted SQL summary for display.
	SqlSummary string `json:"sqlSummary"`
	// Redacted normalized SELECT SQL after parameter binding. Parameter values are not returned raw.
	NormalizedSql *string `json:"normalizedSQL,omitempty"`
	// Redacted parameter binding summary for display and copy_context.
	ParameterSummary *string `json:"parameterSummary,omitempty"`
	// Optional source context for a SQL diagnosis request.
	Context *DiagnosticSQLDiagnosisContext `json:"context,omitempty"`
	// PostgreSQL EXPLAIN plan summary for Console display.
	PlanSummary *DiagnosticSQLDiagnosisPlanSummary `json:"planSummary,omitempty"`
	// Safety boundary enforced by the backend for this SQL diagnosis.
	ReadonlyBoundary DiagnosticSQLDiagnosisReadonlyBoundary `json:"readonlyBoundary"`
	// Safe, redacted text payload that Console may copy for DBA/SRE handoff. It must not contain passwords, tokens, cookies, kubeconfig, or full sensitive SQL text.
	CopyPayload *DiagnosticNextActionCopyPayload `json:"copyPayload,omitempty"`
	// Stable error code for SQL diagnosis failures and safety rejections.
	ErrorCode *DiagnosticSQLDiagnosisErrorCode `json:"errorCode,omitempty"`
	// User-facing failure message.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// Collection time.
	CollectedAt time.Time `json:"collectedAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticSQLDiagnosisResult instantiates a new DiagnosticSQLDiagnosisResult object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticSQLDiagnosisResult(status DiagnosticSQLDiagnosisStatus, sqlSummary string, readonlyBoundary DiagnosticSQLDiagnosisReadonlyBoundary, collectedAt time.Time) *DiagnosticSQLDiagnosisResult {
	this := DiagnosticSQLDiagnosisResult{}
	this.Status = status
	this.SqlSummary = sqlSummary
	this.ReadonlyBoundary = readonlyBoundary
	this.CollectedAt = collectedAt
	return &this
}

// NewDiagnosticSQLDiagnosisResultWithDefaults instantiates a new DiagnosticSQLDiagnosisResult object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticSQLDiagnosisResultWithDefaults() *DiagnosticSQLDiagnosisResult {
	this := DiagnosticSQLDiagnosisResult{}
	return &this
}

// GetStatus returns the Status field value.
func (o *DiagnosticSQLDiagnosisResult) GetStatus() DiagnosticSQLDiagnosisStatus {
	if o == nil {
		var ret DiagnosticSQLDiagnosisStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSQLDiagnosisResult) GetStatusOk() (*DiagnosticSQLDiagnosisStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *DiagnosticSQLDiagnosisResult) SetStatus(v DiagnosticSQLDiagnosisStatus) {
	o.Status = v
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *DiagnosticSQLDiagnosisResult) GetDatabase() string {
	if o == nil || o.Database == nil {
		var ret string
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticSQLDiagnosisResult) GetDatabaseOk() (*string, bool) {
	if o == nil || o.Database == nil {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *DiagnosticSQLDiagnosisResult) HasDatabase() bool {
	return o != nil && o.Database != nil
}

// SetDatabase gets a reference to the given string and assigns it to the Database field.
func (o *DiagnosticSQLDiagnosisResult) SetDatabase(v string) {
	o.Database = &v
}

// GetSqlSummary returns the SqlSummary field value.
func (o *DiagnosticSQLDiagnosisResult) GetSqlSummary() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SqlSummary
}

// GetSqlSummaryOk returns a tuple with the SqlSummary field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSQLDiagnosisResult) GetSqlSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SqlSummary, true
}

// SetSqlSummary sets field value.
func (o *DiagnosticSQLDiagnosisResult) SetSqlSummary(v string) {
	o.SqlSummary = v
}

// GetNormalizedSql returns the NormalizedSql field value if set, zero value otherwise.
func (o *DiagnosticSQLDiagnosisResult) GetNormalizedSql() string {
	if o == nil || o.NormalizedSql == nil {
		var ret string
		return ret
	}
	return *o.NormalizedSql
}

// GetNormalizedSqlOk returns a tuple with the NormalizedSql field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticSQLDiagnosisResult) GetNormalizedSqlOk() (*string, bool) {
	if o == nil || o.NormalizedSql == nil {
		return nil, false
	}
	return o.NormalizedSql, true
}

// HasNormalizedSql returns a boolean if a field has been set.
func (o *DiagnosticSQLDiagnosisResult) HasNormalizedSql() bool {
	return o != nil && o.NormalizedSql != nil
}

// SetNormalizedSql gets a reference to the given string and assigns it to the NormalizedSql field.
func (o *DiagnosticSQLDiagnosisResult) SetNormalizedSql(v string) {
	o.NormalizedSql = &v
}

// GetParameterSummary returns the ParameterSummary field value if set, zero value otherwise.
func (o *DiagnosticSQLDiagnosisResult) GetParameterSummary() string {
	if o == nil || o.ParameterSummary == nil {
		var ret string
		return ret
	}
	return *o.ParameterSummary
}

// GetParameterSummaryOk returns a tuple with the ParameterSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticSQLDiagnosisResult) GetParameterSummaryOk() (*string, bool) {
	if o == nil || o.ParameterSummary == nil {
		return nil, false
	}
	return o.ParameterSummary, true
}

// HasParameterSummary returns a boolean if a field has been set.
func (o *DiagnosticSQLDiagnosisResult) HasParameterSummary() bool {
	return o != nil && o.ParameterSummary != nil
}

// SetParameterSummary gets a reference to the given string and assigns it to the ParameterSummary field.
func (o *DiagnosticSQLDiagnosisResult) SetParameterSummary(v string) {
	o.ParameterSummary = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *DiagnosticSQLDiagnosisResult) GetContext() DiagnosticSQLDiagnosisContext {
	if o == nil || o.Context == nil {
		var ret DiagnosticSQLDiagnosisContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticSQLDiagnosisResult) GetContextOk() (*DiagnosticSQLDiagnosisContext, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *DiagnosticSQLDiagnosisResult) HasContext() bool {
	return o != nil && o.Context != nil
}

// SetContext gets a reference to the given DiagnosticSQLDiagnosisContext and assigns it to the Context field.
func (o *DiagnosticSQLDiagnosisResult) SetContext(v DiagnosticSQLDiagnosisContext) {
	o.Context = &v
}

// GetPlanSummary returns the PlanSummary field value if set, zero value otherwise.
func (o *DiagnosticSQLDiagnosisResult) GetPlanSummary() DiagnosticSQLDiagnosisPlanSummary {
	if o == nil || o.PlanSummary == nil {
		var ret DiagnosticSQLDiagnosisPlanSummary
		return ret
	}
	return *o.PlanSummary
}

// GetPlanSummaryOk returns a tuple with the PlanSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticSQLDiagnosisResult) GetPlanSummaryOk() (*DiagnosticSQLDiagnosisPlanSummary, bool) {
	if o == nil || o.PlanSummary == nil {
		return nil, false
	}
	return o.PlanSummary, true
}

// HasPlanSummary returns a boolean if a field has been set.
func (o *DiagnosticSQLDiagnosisResult) HasPlanSummary() bool {
	return o != nil && o.PlanSummary != nil
}

// SetPlanSummary gets a reference to the given DiagnosticSQLDiagnosisPlanSummary and assigns it to the PlanSummary field.
func (o *DiagnosticSQLDiagnosisResult) SetPlanSummary(v DiagnosticSQLDiagnosisPlanSummary) {
	o.PlanSummary = &v
}

// GetReadonlyBoundary returns the ReadonlyBoundary field value.
func (o *DiagnosticSQLDiagnosisResult) GetReadonlyBoundary() DiagnosticSQLDiagnosisReadonlyBoundary {
	if o == nil {
		var ret DiagnosticSQLDiagnosisReadonlyBoundary
		return ret
	}
	return o.ReadonlyBoundary
}

// GetReadonlyBoundaryOk returns a tuple with the ReadonlyBoundary field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSQLDiagnosisResult) GetReadonlyBoundaryOk() (*DiagnosticSQLDiagnosisReadonlyBoundary, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReadonlyBoundary, true
}

// SetReadonlyBoundary sets field value.
func (o *DiagnosticSQLDiagnosisResult) SetReadonlyBoundary(v DiagnosticSQLDiagnosisReadonlyBoundary) {
	o.ReadonlyBoundary = v
}

// GetCopyPayload returns the CopyPayload field value if set, zero value otherwise.
func (o *DiagnosticSQLDiagnosisResult) GetCopyPayload() DiagnosticNextActionCopyPayload {
	if o == nil || o.CopyPayload == nil {
		var ret DiagnosticNextActionCopyPayload
		return ret
	}
	return *o.CopyPayload
}

// GetCopyPayloadOk returns a tuple with the CopyPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticSQLDiagnosisResult) GetCopyPayloadOk() (*DiagnosticNextActionCopyPayload, bool) {
	if o == nil || o.CopyPayload == nil {
		return nil, false
	}
	return o.CopyPayload, true
}

// HasCopyPayload returns a boolean if a field has been set.
func (o *DiagnosticSQLDiagnosisResult) HasCopyPayload() bool {
	return o != nil && o.CopyPayload != nil
}

// SetCopyPayload gets a reference to the given DiagnosticNextActionCopyPayload and assigns it to the CopyPayload field.
func (o *DiagnosticSQLDiagnosisResult) SetCopyPayload(v DiagnosticNextActionCopyPayload) {
	o.CopyPayload = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *DiagnosticSQLDiagnosisResult) GetErrorCode() DiagnosticSQLDiagnosisErrorCode {
	if o == nil || o.ErrorCode == nil {
		var ret DiagnosticSQLDiagnosisErrorCode
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticSQLDiagnosisResult) GetErrorCodeOk() (*DiagnosticSQLDiagnosisErrorCode, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *DiagnosticSQLDiagnosisResult) HasErrorCode() bool {
	return o != nil && o.ErrorCode != nil
}

// SetErrorCode gets a reference to the given DiagnosticSQLDiagnosisErrorCode and assigns it to the ErrorCode field.
func (o *DiagnosticSQLDiagnosisResult) SetErrorCode(v DiagnosticSQLDiagnosisErrorCode) {
	o.ErrorCode = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *DiagnosticSQLDiagnosisResult) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticSQLDiagnosisResult) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *DiagnosticSQLDiagnosisResult) HasErrorMessage() bool {
	return o != nil && o.ErrorMessage != nil
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *DiagnosticSQLDiagnosisResult) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetCollectedAt returns the CollectedAt field value.
func (o *DiagnosticSQLDiagnosisResult) GetCollectedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CollectedAt
}

// GetCollectedAtOk returns a tuple with the CollectedAt field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSQLDiagnosisResult) GetCollectedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectedAt, true
}

// SetCollectedAt sets field value.
func (o *DiagnosticSQLDiagnosisResult) SetCollectedAt(v time.Time) {
	o.CollectedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticSQLDiagnosisResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["status"] = o.Status
	if o.Database != nil {
		toSerialize["database"] = o.Database
	}
	toSerialize["sqlSummary"] = o.SqlSummary
	if o.NormalizedSql != nil {
		toSerialize["normalizedSQL"] = o.NormalizedSql
	}
	if o.ParameterSummary != nil {
		toSerialize["parameterSummary"] = o.ParameterSummary
	}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	if o.PlanSummary != nil {
		toSerialize["planSummary"] = o.PlanSummary
	}
	toSerialize["readonlyBoundary"] = o.ReadonlyBoundary
	if o.CopyPayload != nil {
		toSerialize["copyPayload"] = o.CopyPayload
	}
	if o.ErrorCode != nil {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if o.ErrorMessage != nil {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if o.CollectedAt.Nanosecond() == 0 {
		toSerialize["collectedAt"] = o.CollectedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["collectedAt"] = o.CollectedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticSQLDiagnosisResult) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Status           *DiagnosticSQLDiagnosisStatus           `json:"status"`
		Database         *string                                 `json:"database,omitempty"`
		SqlSummary       *string                                 `json:"sqlSummary"`
		NormalizedSql    *string                                 `json:"normalizedSQL,omitempty"`
		ParameterSummary *string                                 `json:"parameterSummary,omitempty"`
		Context          *DiagnosticSQLDiagnosisContext          `json:"context,omitempty"`
		PlanSummary      *DiagnosticSQLDiagnosisPlanSummary      `json:"planSummary,omitempty"`
		ReadonlyBoundary *DiagnosticSQLDiagnosisReadonlyBoundary `json:"readonlyBoundary"`
		CopyPayload      *DiagnosticNextActionCopyPayload        `json:"copyPayload,omitempty"`
		ErrorCode        *DiagnosticSQLDiagnosisErrorCode        `json:"errorCode,omitempty"`
		ErrorMessage     *string                                 `json:"errorMessage,omitempty"`
		CollectedAt      *time.Time                              `json:"collectedAt"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.SqlSummary == nil {
		return fmt.Errorf("required field sqlSummary missing")
	}
	if all.ReadonlyBoundary == nil {
		return fmt.Errorf("required field readonlyBoundary missing")
	}
	if all.CollectedAt == nil {
		return fmt.Errorf("required field collectedAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"status", "database", "sqlSummary", "normalizedSQL", "parameterSummary", "context", "planSummary", "readonlyBoundary", "copyPayload", "errorCode", "errorMessage", "collectedAt"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.Database = all.Database
	o.SqlSummary = *all.SqlSummary
	o.NormalizedSql = all.NormalizedSql
	o.ParameterSummary = all.ParameterSummary
	if all.Context != nil && all.Context.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Context = all.Context
	if all.PlanSummary != nil && all.PlanSummary.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.PlanSummary = all.PlanSummary
	if all.ReadonlyBoundary.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ReadonlyBoundary = *all.ReadonlyBoundary
	if all.CopyPayload != nil && all.CopyPayload.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CopyPayload = all.CopyPayload
	if all.ErrorCode != nil && !all.ErrorCode.IsValid() {
		hasInvalidField = true
	} else {
		o.ErrorCode = all.ErrorCode
	}
	o.ErrorMessage = all.ErrorMessage
	o.CollectedAt = *all.CollectedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
