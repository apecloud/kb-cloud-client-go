// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DmsAuditStatement struct {
	Sequence     int32          `json:"sequence"`
	Statement    string         `json:"statement"`
	Database     string         `json:"database"`
	TableNames   []string       `json:"tableNames"`
	SessionId    *string        `json:"sessionId,omitempty"`
	Risk         DmsAuditRisk   `json:"risk"`
	Status       DmsAuditStatus `json:"status"`
	ErrorMessage string         `json:"errorMessage"`
	DurationMs   *int64         `json:"durationMs,omitempty"`
	RowsAffected *int64         `json:"rowsAffected,omitempty"`
	StartedAt    *time.Time     `json:"startedAt,omitempty"`
	FinishedAt   *time.Time     `json:"finishedAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmsAuditStatement instantiates a new DmsAuditStatement object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsAuditStatement(sequence int32, statement string, database string, tableNames []string, risk DmsAuditRisk, status DmsAuditStatus, errorMessage string) *DmsAuditStatement {
	this := DmsAuditStatement{}
	this.Sequence = sequence
	this.Statement = statement
	this.Database = database
	this.TableNames = tableNames
	this.Risk = risk
	this.Status = status
	this.ErrorMessage = errorMessage
	return &this
}

// NewDmsAuditStatementWithDefaults instantiates a new DmsAuditStatement object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmsAuditStatementWithDefaults() *DmsAuditStatement {
	this := DmsAuditStatement{}
	return &this
}

// GetSequence returns the Sequence field value.
func (o *DmsAuditStatement) GetSequence() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value
// and a boolean to check if the value has been set.
func (o *DmsAuditStatement) GetSequenceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sequence, true
}

// SetSequence sets field value.
func (o *DmsAuditStatement) SetSequence(v int32) {
	o.Sequence = v
}

// GetStatement returns the Statement field value.
func (o *DmsAuditStatement) GetStatement() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Statement
}

// GetStatementOk returns a tuple with the Statement field value
// and a boolean to check if the value has been set.
func (o *DmsAuditStatement) GetStatementOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Statement, true
}

// SetStatement sets field value.
func (o *DmsAuditStatement) SetStatement(v string) {
	o.Statement = v
}

// GetDatabase returns the Database field value.
func (o *DmsAuditStatement) GetDatabase() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Database
}

// GetDatabaseOk returns a tuple with the Database field value
// and a boolean to check if the value has been set.
func (o *DmsAuditStatement) GetDatabaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Database, true
}

// SetDatabase sets field value.
func (o *DmsAuditStatement) SetDatabase(v string) {
	o.Database = v
}

// GetTableNames returns the TableNames field value.
func (o *DmsAuditStatement) GetTableNames() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TableNames
}

// GetTableNamesOk returns a tuple with the TableNames field value
// and a boolean to check if the value has been set.
func (o *DmsAuditStatement) GetTableNamesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TableNames, true
}

// SetTableNames sets field value.
func (o *DmsAuditStatement) SetTableNames(v []string) {
	o.TableNames = v
}

// GetSessionId returns the SessionId field value if set, zero value otherwise.
func (o *DmsAuditStatement) GetSessionId() string {
	if o == nil || o.SessionId == nil {
		var ret string
		return ret
	}
	return *o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsAuditStatement) GetSessionIdOk() (*string, bool) {
	if o == nil || o.SessionId == nil {
		return nil, false
	}
	return o.SessionId, true
}

// HasSessionId returns a boolean if a field has been set.
func (o *DmsAuditStatement) HasSessionId() bool {
	return o != nil && o.SessionId != nil
}

// SetSessionId gets a reference to the given string and assigns it to the SessionId field.
func (o *DmsAuditStatement) SetSessionId(v string) {
	o.SessionId = &v
}

// GetRisk returns the Risk field value.
func (o *DmsAuditStatement) GetRisk() DmsAuditRisk {
	if o == nil {
		var ret DmsAuditRisk
		return ret
	}
	return o.Risk
}

// GetRiskOk returns a tuple with the Risk field value
// and a boolean to check if the value has been set.
func (o *DmsAuditStatement) GetRiskOk() (*DmsAuditRisk, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Risk, true
}

// SetRisk sets field value.
func (o *DmsAuditStatement) SetRisk(v DmsAuditRisk) {
	o.Risk = v
}

// GetStatus returns the Status field value.
func (o *DmsAuditStatement) GetStatus() DmsAuditStatus {
	if o == nil {
		var ret DmsAuditStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DmsAuditStatement) GetStatusOk() (*DmsAuditStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *DmsAuditStatement) SetStatus(v DmsAuditStatus) {
	o.Status = v
}

// GetErrorMessage returns the ErrorMessage field value.
func (o *DmsAuditStatement) GetErrorMessage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value
// and a boolean to check if the value has been set.
func (o *DmsAuditStatement) GetErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorMessage, true
}

// SetErrorMessage sets field value.
func (o *DmsAuditStatement) SetErrorMessage(v string) {
	o.ErrorMessage = v
}

// GetDurationMs returns the DurationMs field value if set, zero value otherwise.
func (o *DmsAuditStatement) GetDurationMs() int64 {
	if o == nil || o.DurationMs == nil {
		var ret int64
		return ret
	}
	return *o.DurationMs
}

// GetDurationMsOk returns a tuple with the DurationMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsAuditStatement) GetDurationMsOk() (*int64, bool) {
	if o == nil || o.DurationMs == nil {
		return nil, false
	}
	return o.DurationMs, true
}

// HasDurationMs returns a boolean if a field has been set.
func (o *DmsAuditStatement) HasDurationMs() bool {
	return o != nil && o.DurationMs != nil
}

// SetDurationMs gets a reference to the given int64 and assigns it to the DurationMs field.
func (o *DmsAuditStatement) SetDurationMs(v int64) {
	o.DurationMs = &v
}

// GetRowsAffected returns the RowsAffected field value if set, zero value otherwise.
func (o *DmsAuditStatement) GetRowsAffected() int64 {
	if o == nil || o.RowsAffected == nil {
		var ret int64
		return ret
	}
	return *o.RowsAffected
}

// GetRowsAffectedOk returns a tuple with the RowsAffected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsAuditStatement) GetRowsAffectedOk() (*int64, bool) {
	if o == nil || o.RowsAffected == nil {
		return nil, false
	}
	return o.RowsAffected, true
}

// HasRowsAffected returns a boolean if a field has been set.
func (o *DmsAuditStatement) HasRowsAffected() bool {
	return o != nil && o.RowsAffected != nil
}

// SetRowsAffected gets a reference to the given int64 and assigns it to the RowsAffected field.
func (o *DmsAuditStatement) SetRowsAffected(v int64) {
	o.RowsAffected = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *DmsAuditStatement) GetStartedAt() time.Time {
	if o == nil || o.StartedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsAuditStatement) GetStartedAtOk() (*time.Time, bool) {
	if o == nil || o.StartedAt == nil {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *DmsAuditStatement) HasStartedAt() bool {
	return o != nil && o.StartedAt != nil
}

// SetStartedAt gets a reference to the given time.Time and assigns it to the StartedAt field.
func (o *DmsAuditStatement) SetStartedAt(v time.Time) {
	o.StartedAt = &v
}

// GetFinishedAt returns the FinishedAt field value if set, zero value otherwise.
func (o *DmsAuditStatement) GetFinishedAt() time.Time {
	if o == nil || o.FinishedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.FinishedAt
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsAuditStatement) GetFinishedAtOk() (*time.Time, bool) {
	if o == nil || o.FinishedAt == nil {
		return nil, false
	}
	return o.FinishedAt, true
}

// HasFinishedAt returns a boolean if a field has been set.
func (o *DmsAuditStatement) HasFinishedAt() bool {
	return o != nil && o.FinishedAt != nil
}

// SetFinishedAt gets a reference to the given time.Time and assigns it to the FinishedAt field.
func (o *DmsAuditStatement) SetFinishedAt(v time.Time) {
	o.FinishedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DmsAuditStatement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["sequence"] = o.Sequence
	toSerialize["statement"] = o.Statement
	toSerialize["database"] = o.Database
	toSerialize["tableNames"] = o.TableNames
	if o.SessionId != nil {
		toSerialize["sessionId"] = o.SessionId
	}
	toSerialize["risk"] = o.Risk
	toSerialize["status"] = o.Status
	toSerialize["errorMessage"] = o.ErrorMessage
	if o.DurationMs != nil {
		toSerialize["durationMs"] = o.DurationMs
	}
	if o.RowsAffected != nil {
		toSerialize["rowsAffected"] = o.RowsAffected
	}
	if o.StartedAt != nil {
		if o.StartedAt.Nanosecond() == 0 {
			toSerialize["startedAt"] = o.StartedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["startedAt"] = o.StartedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.FinishedAt != nil {
		if o.FinishedAt.Nanosecond() == 0 {
			toSerialize["finishedAt"] = o.FinishedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["finishedAt"] = o.FinishedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmsAuditStatement) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Sequence     *int32          `json:"sequence"`
		Statement    *string         `json:"statement"`
		Database     *string         `json:"database"`
		TableNames   *[]string       `json:"tableNames"`
		SessionId    *string         `json:"sessionId,omitempty"`
		Risk         *DmsAuditRisk   `json:"risk"`
		Status       *DmsAuditStatus `json:"status"`
		ErrorMessage *string         `json:"errorMessage"`
		DurationMs   *int64          `json:"durationMs,omitempty"`
		RowsAffected *int64          `json:"rowsAffected,omitempty"`
		StartedAt    *time.Time      `json:"startedAt,omitempty"`
		FinishedAt   *time.Time      `json:"finishedAt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Sequence == nil {
		return fmt.Errorf("required field sequence missing")
	}
	if all.Statement == nil {
		return fmt.Errorf("required field statement missing")
	}
	if all.Database == nil {
		return fmt.Errorf("required field database missing")
	}
	if all.TableNames == nil {
		return fmt.Errorf("required field tableNames missing")
	}
	if all.Risk == nil {
		return fmt.Errorf("required field risk missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.ErrorMessage == nil {
		return fmt.Errorf("required field errorMessage missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"sequence", "statement", "database", "tableNames", "sessionId", "risk", "status", "errorMessage", "durationMs", "rowsAffected", "startedAt", "finishedAt"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Sequence = *all.Sequence
	o.Statement = *all.Statement
	o.Database = *all.Database
	o.TableNames = *all.TableNames
	o.SessionId = all.SessionId
	if !all.Risk.IsValid() {
		hasInvalidField = true
	} else {
		o.Risk = *all.Risk
	}
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.ErrorMessage = *all.ErrorMessage
	o.DurationMs = all.DurationMs
	o.RowsAffected = all.RowsAffected
	o.StartedAt = all.StartedAt
	o.FinishedAt = all.FinishedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
