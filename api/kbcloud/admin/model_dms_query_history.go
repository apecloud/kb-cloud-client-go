// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DmsQueryHistory struct {
	// history record id
	Id *int64 `json:"id,omitempty"`
	// organization name
	OrgName *string `json:"orgName,omitempty"`
	// cluster name
	ClusterName *string `json:"clusterName,omitempty"`
	// datasource id
	DatasourceId *int64 `json:"datasourceId,omitempty"`
	// datasource display name
	DatasourceName *string `json:"datasourceName,omitempty"`
	// user id that created the history record
	UserId *string `json:"userId,omitempty"`
	// datasource engine
	Engine *string `json:"engine,omitempty"`
	// operation history type, for example sql or shell
	HistoryType *string `json:"historyType,omitempty"`
	// shell session id when historyType is shell
	SessionId *string `json:"sessionId,omitempty"`
	// shell prompt after evaluation
	Prompt *string `json:"prompt,omitempty"`
	// current database inferred from shell prompt
	CurrentDb *string `json:"currentDb,omitempty"`
	// executed command, generalized from sql for shell history
	Command *string `json:"command,omitempty"`
	// command status, such as success, error, interrupted, or timeout
	Status *string `json:"status,omitempty"`
	// truncated and redacted result summary
	ResultSummary *string `json:"resultSummary,omitempty"`
	// truncated and redacted error message
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// executed sql statements
	Sql *string `json:"sql,omitempty"`
	// sql executed massage
	ErrMassage *string `json:"errMassage,omitempty"`
	// sql executed time
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// sql executed duration
	Duration *int64 `json:"duration,omitempty"`
	// shell command duration in milliseconds
	DurationMs *int64 `json:"durationMs,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmsQueryHistory instantiates a new DmsQueryHistory object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsQueryHistory() *DmsQueryHistory {
	this := DmsQueryHistory{}
	return &this
}

// NewDmsQueryHistoryWithDefaults instantiates a new DmsQueryHistory object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmsQueryHistoryWithDefaults() *DmsQueryHistory {
	this := DmsQueryHistory{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DmsQueryHistory) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsQueryHistory) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DmsQueryHistory) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *DmsQueryHistory) SetId(v int64) {
	o.Id = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *DmsQueryHistory) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsQueryHistory) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *DmsQueryHistory) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *DmsQueryHistory) SetOrgName(v string) {
	o.OrgName = &v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *DmsQueryHistory) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsQueryHistory) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *DmsQueryHistory) HasClusterName() bool {
	return o != nil && o.ClusterName != nil
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *DmsQueryHistory) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetDatasourceId returns the DatasourceId field value if set, zero value otherwise.
func (o *DmsQueryHistory) GetDatasourceId() int64 {
	if o == nil || o.DatasourceId == nil {
		var ret int64
		return ret
	}
	return *o.DatasourceId
}

// GetDatasourceIdOk returns a tuple with the DatasourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsQueryHistory) GetDatasourceIdOk() (*int64, bool) {
	if o == nil || o.DatasourceId == nil {
		return nil, false
	}
	return o.DatasourceId, true
}

// HasDatasourceId returns a boolean if a field has been set.
func (o *DmsQueryHistory) HasDatasourceId() bool {
	return o != nil && o.DatasourceId != nil
}

// SetDatasourceId gets a reference to the given int64 and assigns it to the DatasourceId field.
func (o *DmsQueryHistory) SetDatasourceId(v int64) {
	o.DatasourceId = &v
}

// GetDatasourceName returns the DatasourceName field value if set, zero value otherwise.
func (o *DmsQueryHistory) GetDatasourceName() string {
	if o == nil || o.DatasourceName == nil {
		var ret string
		return ret
	}
	return *o.DatasourceName
}

// GetDatasourceNameOk returns a tuple with the DatasourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsQueryHistory) GetDatasourceNameOk() (*string, bool) {
	if o == nil || o.DatasourceName == nil {
		return nil, false
	}
	return o.DatasourceName, true
}

// HasDatasourceName returns a boolean if a field has been set.
func (o *DmsQueryHistory) HasDatasourceName() bool {
	return o != nil && o.DatasourceName != nil
}

// SetDatasourceName gets a reference to the given string and assigns it to the DatasourceName field.
func (o *DmsQueryHistory) SetDatasourceName(v string) {
	o.DatasourceName = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *DmsQueryHistory) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsQueryHistory) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *DmsQueryHistory) HasUserId() bool {
	return o != nil && o.UserId != nil
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *DmsQueryHistory) SetUserId(v string) {
	o.UserId = &v
}

// GetEngine returns the Engine field value if set, zero value otherwise.
func (o *DmsQueryHistory) GetEngine() string {
	if o == nil || o.Engine == nil {
		var ret string
		return ret
	}
	return *o.Engine
}

// GetEngineOk returns a tuple with the Engine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsQueryHistory) GetEngineOk() (*string, bool) {
	if o == nil || o.Engine == nil {
		return nil, false
	}
	return o.Engine, true
}

// HasEngine returns a boolean if a field has been set.
func (o *DmsQueryHistory) HasEngine() bool {
	return o != nil && o.Engine != nil
}

// SetEngine gets a reference to the given string and assigns it to the Engine field.
func (o *DmsQueryHistory) SetEngine(v string) {
	o.Engine = &v
}

// GetHistoryType returns the HistoryType field value if set, zero value otherwise.
func (o *DmsQueryHistory) GetHistoryType() string {
	if o == nil || o.HistoryType == nil {
		var ret string
		return ret
	}
	return *o.HistoryType
}

// GetHistoryTypeOk returns a tuple with the HistoryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsQueryHistory) GetHistoryTypeOk() (*string, bool) {
	if o == nil || o.HistoryType == nil {
		return nil, false
	}
	return o.HistoryType, true
}

// HasHistoryType returns a boolean if a field has been set.
func (o *DmsQueryHistory) HasHistoryType() bool {
	return o != nil && o.HistoryType != nil
}

// SetHistoryType gets a reference to the given string and assigns it to the HistoryType field.
func (o *DmsQueryHistory) SetHistoryType(v string) {
	o.HistoryType = &v
}

// GetSessionId returns the SessionId field value if set, zero value otherwise.
func (o *DmsQueryHistory) GetSessionId() string {
	if o == nil || o.SessionId == nil {
		var ret string
		return ret
	}
	return *o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsQueryHistory) GetSessionIdOk() (*string, bool) {
	if o == nil || o.SessionId == nil {
		return nil, false
	}
	return o.SessionId, true
}

// HasSessionId returns a boolean if a field has been set.
func (o *DmsQueryHistory) HasSessionId() bool {
	return o != nil && o.SessionId != nil
}

// SetSessionId gets a reference to the given string and assigns it to the SessionId field.
func (o *DmsQueryHistory) SetSessionId(v string) {
	o.SessionId = &v
}

// GetPrompt returns the Prompt field value if set, zero value otherwise.
func (o *DmsQueryHistory) GetPrompt() string {
	if o == nil || o.Prompt == nil {
		var ret string
		return ret
	}
	return *o.Prompt
}

// GetPromptOk returns a tuple with the Prompt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsQueryHistory) GetPromptOk() (*string, bool) {
	if o == nil || o.Prompt == nil {
		return nil, false
	}
	return o.Prompt, true
}

// HasPrompt returns a boolean if a field has been set.
func (o *DmsQueryHistory) HasPrompt() bool {
	return o != nil && o.Prompt != nil
}

// SetPrompt gets a reference to the given string and assigns it to the Prompt field.
func (o *DmsQueryHistory) SetPrompt(v string) {
	o.Prompt = &v
}

// GetCurrentDb returns the CurrentDb field value if set, zero value otherwise.
func (o *DmsQueryHistory) GetCurrentDb() string {
	if o == nil || o.CurrentDb == nil {
		var ret string
		return ret
	}
	return *o.CurrentDb
}

// GetCurrentDbOk returns a tuple with the CurrentDb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsQueryHistory) GetCurrentDbOk() (*string, bool) {
	if o == nil || o.CurrentDb == nil {
		return nil, false
	}
	return o.CurrentDb, true
}

// HasCurrentDb returns a boolean if a field has been set.
func (o *DmsQueryHistory) HasCurrentDb() bool {
	return o != nil && o.CurrentDb != nil
}

// SetCurrentDb gets a reference to the given string and assigns it to the CurrentDb field.
func (o *DmsQueryHistory) SetCurrentDb(v string) {
	o.CurrentDb = &v
}

// GetCommand returns the Command field value if set, zero value otherwise.
func (o *DmsQueryHistory) GetCommand() string {
	if o == nil || o.Command == nil {
		var ret string
		return ret
	}
	return *o.Command
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsQueryHistory) GetCommandOk() (*string, bool) {
	if o == nil || o.Command == nil {
		return nil, false
	}
	return o.Command, true
}

// HasCommand returns a boolean if a field has been set.
func (o *DmsQueryHistory) HasCommand() bool {
	return o != nil && o.Command != nil
}

// SetCommand gets a reference to the given string and assigns it to the Command field.
func (o *DmsQueryHistory) SetCommand(v string) {
	o.Command = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DmsQueryHistory) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsQueryHistory) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DmsQueryHistory) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DmsQueryHistory) SetStatus(v string) {
	o.Status = &v
}

// GetResultSummary returns the ResultSummary field value if set, zero value otherwise.
func (o *DmsQueryHistory) GetResultSummary() string {
	if o == nil || o.ResultSummary == nil {
		var ret string
		return ret
	}
	return *o.ResultSummary
}

// GetResultSummaryOk returns a tuple with the ResultSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsQueryHistory) GetResultSummaryOk() (*string, bool) {
	if o == nil || o.ResultSummary == nil {
		return nil, false
	}
	return o.ResultSummary, true
}

// HasResultSummary returns a boolean if a field has been set.
func (o *DmsQueryHistory) HasResultSummary() bool {
	return o != nil && o.ResultSummary != nil
}

// SetResultSummary gets a reference to the given string and assigns it to the ResultSummary field.
func (o *DmsQueryHistory) SetResultSummary(v string) {
	o.ResultSummary = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *DmsQueryHistory) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsQueryHistory) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *DmsQueryHistory) HasErrorMessage() bool {
	return o != nil && o.ErrorMessage != nil
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *DmsQueryHistory) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetSql returns the Sql field value if set, zero value otherwise.
func (o *DmsQueryHistory) GetSql() string {
	if o == nil || o.Sql == nil {
		var ret string
		return ret
	}
	return *o.Sql
}

// GetSqlOk returns a tuple with the Sql field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsQueryHistory) GetSqlOk() (*string, bool) {
	if o == nil || o.Sql == nil {
		return nil, false
	}
	return o.Sql, true
}

// HasSql returns a boolean if a field has been set.
func (o *DmsQueryHistory) HasSql() bool {
	return o != nil && o.Sql != nil
}

// SetSql gets a reference to the given string and assigns it to the Sql field.
func (o *DmsQueryHistory) SetSql(v string) {
	o.Sql = &v
}

// GetErrMassage returns the ErrMassage field value if set, zero value otherwise.
func (o *DmsQueryHistory) GetErrMassage() string {
	if o == nil || o.ErrMassage == nil {
		var ret string
		return ret
	}
	return *o.ErrMassage
}

// GetErrMassageOk returns a tuple with the ErrMassage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsQueryHistory) GetErrMassageOk() (*string, bool) {
	if o == nil || o.ErrMassage == nil {
		return nil, false
	}
	return o.ErrMassage, true
}

// HasErrMassage returns a boolean if a field has been set.
func (o *DmsQueryHistory) HasErrMassage() bool {
	return o != nil && o.ErrMassage != nil
}

// SetErrMassage gets a reference to the given string and assigns it to the ErrMassage field.
func (o *DmsQueryHistory) SetErrMassage(v string) {
	o.ErrMassage = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DmsQueryHistory) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsQueryHistory) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DmsQueryHistory) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *DmsQueryHistory) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *DmsQueryHistory) GetDuration() int64 {
	if o == nil || o.Duration == nil {
		var ret int64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsQueryHistory) GetDurationOk() (*int64, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *DmsQueryHistory) HasDuration() bool {
	return o != nil && o.Duration != nil
}

// SetDuration gets a reference to the given int64 and assigns it to the Duration field.
func (o *DmsQueryHistory) SetDuration(v int64) {
	o.Duration = &v
}

// GetDurationMs returns the DurationMs field value if set, zero value otherwise.
func (o *DmsQueryHistory) GetDurationMs() int64 {
	if o == nil || o.DurationMs == nil {
		var ret int64
		return ret
	}
	return *o.DurationMs
}

// GetDurationMsOk returns a tuple with the DurationMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsQueryHistory) GetDurationMsOk() (*int64, bool) {
	if o == nil || o.DurationMs == nil {
		return nil, false
	}
	return o.DurationMs, true
}

// HasDurationMs returns a boolean if a field has been set.
func (o *DmsQueryHistory) HasDurationMs() bool {
	return o != nil && o.DurationMs != nil
}

// SetDurationMs gets a reference to the given int64 and assigns it to the DurationMs field.
func (o *DmsQueryHistory) SetDurationMs(v int64) {
	o.DurationMs = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DmsQueryHistory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.OrgName != nil {
		toSerialize["orgName"] = o.OrgName
	}
	if o.ClusterName != nil {
		toSerialize["clusterName"] = o.ClusterName
	}
	if o.DatasourceId != nil {
		toSerialize["datasourceId"] = o.DatasourceId
	}
	if o.DatasourceName != nil {
		toSerialize["datasourceName"] = o.DatasourceName
	}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	if o.Engine != nil {
		toSerialize["engine"] = o.Engine
	}
	if o.HistoryType != nil {
		toSerialize["historyType"] = o.HistoryType
	}
	if o.SessionId != nil {
		toSerialize["sessionId"] = o.SessionId
	}
	if o.Prompt != nil {
		toSerialize["prompt"] = o.Prompt
	}
	if o.CurrentDb != nil {
		toSerialize["currentDb"] = o.CurrentDb
	}
	if o.Command != nil {
		toSerialize["command"] = o.Command
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.ResultSummary != nil {
		toSerialize["resultSummary"] = o.ResultSummary
	}
	if o.ErrorMessage != nil {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if o.Sql != nil {
		toSerialize["sql"] = o.Sql
	}
	if o.ErrMassage != nil {
		toSerialize["errMassage"] = o.ErrMassage
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	if o.DurationMs != nil {
		toSerialize["durationMs"] = o.DurationMs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmsQueryHistory) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id             *int64     `json:"id,omitempty"`
		OrgName        *string    `json:"orgName,omitempty"`
		ClusterName    *string    `json:"clusterName,omitempty"`
		DatasourceId   *int64     `json:"datasourceId,omitempty"`
		DatasourceName *string    `json:"datasourceName,omitempty"`
		UserId         *string    `json:"userId,omitempty"`
		Engine         *string    `json:"engine,omitempty"`
		HistoryType    *string    `json:"historyType,omitempty"`
		SessionId      *string    `json:"sessionId,omitempty"`
		Prompt         *string    `json:"prompt,omitempty"`
		CurrentDb      *string    `json:"currentDb,omitempty"`
		Command        *string    `json:"command,omitempty"`
		Status         *string    `json:"status,omitempty"`
		ResultSummary  *string    `json:"resultSummary,omitempty"`
		ErrorMessage   *string    `json:"errorMessage,omitempty"`
		Sql            *string    `json:"sql,omitempty"`
		ErrMassage     *string    `json:"errMassage,omitempty"`
		CreatedAt      *time.Time `json:"createdAt,omitempty"`
		Duration       *int64     `json:"duration,omitempty"`
		DurationMs     *int64     `json:"durationMs,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "orgName", "clusterName", "datasourceId", "datasourceName", "userId", "engine", "historyType", "sessionId", "prompt", "currentDb", "command", "status", "resultSummary", "errorMessage", "sql", "errMassage", "createdAt", "duration", "durationMs"})
	} else {
		return err
	}
	o.Id = all.Id
	o.OrgName = all.OrgName
	o.ClusterName = all.ClusterName
	o.DatasourceId = all.DatasourceId
	o.DatasourceName = all.DatasourceName
	o.UserId = all.UserId
	o.Engine = all.Engine
	o.HistoryType = all.HistoryType
	o.SessionId = all.SessionId
	o.Prompt = all.Prompt
	o.CurrentDb = all.CurrentDb
	o.Command = all.Command
	o.Status = all.Status
	o.ResultSummary = all.ResultSummary
	o.ErrorMessage = all.ErrorMessage
	o.Sql = all.Sql
	o.ErrMassage = all.ErrMassage
	o.CreatedAt = all.CreatedAt
	o.Duration = all.Duration
	o.DurationMs = all.DurationMs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
