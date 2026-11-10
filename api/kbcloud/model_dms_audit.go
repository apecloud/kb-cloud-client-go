// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DmsAudit struct {
	Id           string  `json:"id"`
	OrgName      string  `json:"orgName"`
	ClusterId    string  `json:"clusterId"`
	ClusterName  string  `json:"clusterName"`
	DatasourceId *string `json:"datasourceId,omitempty"`
	// Absent for migrated history with unknown actor.
	UserId    *string `json:"userId,omitempty"`
	UserName  string  `json:"userName"`
	Engine    string  `json:"engine"`
	Database  string  `json:"database"`
	Operation string  `json:"operation"`
	Statement string  `json:"statement"`
	ClientIp  *string `json:"clientIp,omitempty"`
	// Native database session ID; see individual statements when a batch uses multiple connections.
	SessionId    *string             `json:"sessionId,omitempty"`
	Risk         *DmsAuditRisk       `json:"risk,omitempty"`
	TableNames   []string            `json:"tableNames,omitempty"`
	ErrorMessage *string             `json:"errorMessage,omitempty"`
	RowsAffected *int64              `json:"rowsAffected,omitempty"`
	Statements   []DmsAuditStatement `json:"statements,omitempty"`
	Status       DmsAuditStatus      `json:"status"`
	CreatedAt    time.Time           `json:"createdAt"`
	FinishedAt   *time.Time          `json:"finishedAt,omitempty"`
	DurationMs   *int64              `json:"durationMs,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmsAudit instantiates a new DmsAudit object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsAudit(id string, orgName string, clusterId string, clusterName string, userName string, engine string, database string, operation string, statement string, status DmsAuditStatus, createdAt time.Time) *DmsAudit {
	this := DmsAudit{}
	this.Id = id
	this.OrgName = orgName
	this.ClusterId = clusterId
	this.ClusterName = clusterName
	this.UserName = userName
	this.Engine = engine
	this.Database = database
	this.Operation = operation
	this.Statement = statement
	this.Status = status
	this.CreatedAt = createdAt
	return &this
}

// NewDmsAuditWithDefaults instantiates a new DmsAudit object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmsAuditWithDefaults() *DmsAudit {
	this := DmsAudit{}
	return &this
}

// GetId returns the Id field value.
func (o *DmsAudit) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DmsAudit) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *DmsAudit) SetId(v string) {
	o.Id = v
}

// GetOrgName returns the OrgName field value.
func (o *DmsAudit) GetOrgName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value
// and a boolean to check if the value has been set.
func (o *DmsAudit) GetOrgNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgName, true
}

// SetOrgName sets field value.
func (o *DmsAudit) SetOrgName(v string) {
	o.OrgName = v
}

// GetClusterId returns the ClusterId field value.
func (o *DmsAudit) GetClusterId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value
// and a boolean to check if the value has been set.
func (o *DmsAudit) GetClusterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterId, true
}

// SetClusterId sets field value.
func (o *DmsAudit) SetClusterId(v string) {
	o.ClusterId = v
}

// GetClusterName returns the ClusterName field value.
func (o *DmsAudit) GetClusterName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value
// and a boolean to check if the value has been set.
func (o *DmsAudit) GetClusterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterName, true
}

// SetClusterName sets field value.
func (o *DmsAudit) SetClusterName(v string) {
	o.ClusterName = v
}

// GetDatasourceId returns the DatasourceId field value if set, zero value otherwise.
func (o *DmsAudit) GetDatasourceId() string {
	if o == nil || o.DatasourceId == nil {
		var ret string
		return ret
	}
	return *o.DatasourceId
}

// GetDatasourceIdOk returns a tuple with the DatasourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsAudit) GetDatasourceIdOk() (*string, bool) {
	if o == nil || o.DatasourceId == nil {
		return nil, false
	}
	return o.DatasourceId, true
}

// HasDatasourceId returns a boolean if a field has been set.
func (o *DmsAudit) HasDatasourceId() bool {
	return o != nil && o.DatasourceId != nil
}

// SetDatasourceId gets a reference to the given string and assigns it to the DatasourceId field.
func (o *DmsAudit) SetDatasourceId(v string) {
	o.DatasourceId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *DmsAudit) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsAudit) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *DmsAudit) HasUserId() bool {
	return o != nil && o.UserId != nil
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *DmsAudit) SetUserId(v string) {
	o.UserId = &v
}

// GetUserName returns the UserName field value.
func (o *DmsAudit) GetUserName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
func (o *DmsAudit) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserName, true
}

// SetUserName sets field value.
func (o *DmsAudit) SetUserName(v string) {
	o.UserName = v
}

// GetEngine returns the Engine field value.
func (o *DmsAudit) GetEngine() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Engine
}

// GetEngineOk returns a tuple with the Engine field value
// and a boolean to check if the value has been set.
func (o *DmsAudit) GetEngineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Engine, true
}

// SetEngine sets field value.
func (o *DmsAudit) SetEngine(v string) {
	o.Engine = v
}

// GetDatabase returns the Database field value.
func (o *DmsAudit) GetDatabase() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Database
}

// GetDatabaseOk returns a tuple with the Database field value
// and a boolean to check if the value has been set.
func (o *DmsAudit) GetDatabaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Database, true
}

// SetDatabase sets field value.
func (o *DmsAudit) SetDatabase(v string) {
	o.Database = v
}

// GetOperation returns the Operation field value.
func (o *DmsAudit) GetOperation() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *DmsAudit) GetOperationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value.
func (o *DmsAudit) SetOperation(v string) {
	o.Operation = v
}

// GetStatement returns the Statement field value.
func (o *DmsAudit) GetStatement() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Statement
}

// GetStatementOk returns a tuple with the Statement field value
// and a boolean to check if the value has been set.
func (o *DmsAudit) GetStatementOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Statement, true
}

// SetStatement sets field value.
func (o *DmsAudit) SetStatement(v string) {
	o.Statement = v
}

// GetClientIp returns the ClientIp field value if set, zero value otherwise.
func (o *DmsAudit) GetClientIp() string {
	if o == nil || o.ClientIp == nil {
		var ret string
		return ret
	}
	return *o.ClientIp
}

// GetClientIpOk returns a tuple with the ClientIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsAudit) GetClientIpOk() (*string, bool) {
	if o == nil || o.ClientIp == nil {
		return nil, false
	}
	return o.ClientIp, true
}

// HasClientIp returns a boolean if a field has been set.
func (o *DmsAudit) HasClientIp() bool {
	return o != nil && o.ClientIp != nil
}

// SetClientIp gets a reference to the given string and assigns it to the ClientIp field.
func (o *DmsAudit) SetClientIp(v string) {
	o.ClientIp = &v
}

// GetSessionId returns the SessionId field value if set, zero value otherwise.
func (o *DmsAudit) GetSessionId() string {
	if o == nil || o.SessionId == nil {
		var ret string
		return ret
	}
	return *o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsAudit) GetSessionIdOk() (*string, bool) {
	if o == nil || o.SessionId == nil {
		return nil, false
	}
	return o.SessionId, true
}

// HasSessionId returns a boolean if a field has been set.
func (o *DmsAudit) HasSessionId() bool {
	return o != nil && o.SessionId != nil
}

// SetSessionId gets a reference to the given string and assigns it to the SessionId field.
func (o *DmsAudit) SetSessionId(v string) {
	o.SessionId = &v
}

// GetRisk returns the Risk field value if set, zero value otherwise.
func (o *DmsAudit) GetRisk() DmsAuditRisk {
	if o == nil || o.Risk == nil {
		var ret DmsAuditRisk
		return ret
	}
	return *o.Risk
}

// GetRiskOk returns a tuple with the Risk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsAudit) GetRiskOk() (*DmsAuditRisk, bool) {
	if o == nil || o.Risk == nil {
		return nil, false
	}
	return o.Risk, true
}

// HasRisk returns a boolean if a field has been set.
func (o *DmsAudit) HasRisk() bool {
	return o != nil && o.Risk != nil
}

// SetRisk gets a reference to the given DmsAuditRisk and assigns it to the Risk field.
func (o *DmsAudit) SetRisk(v DmsAuditRisk) {
	o.Risk = &v
}

// GetTableNames returns the TableNames field value if set, zero value otherwise.
func (o *DmsAudit) GetTableNames() []string {
	if o == nil || o.TableNames == nil {
		var ret []string
		return ret
	}
	return o.TableNames
}

// GetTableNamesOk returns a tuple with the TableNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsAudit) GetTableNamesOk() (*[]string, bool) {
	if o == nil || o.TableNames == nil {
		return nil, false
	}
	return &o.TableNames, true
}

// HasTableNames returns a boolean if a field has been set.
func (o *DmsAudit) HasTableNames() bool {
	return o != nil && o.TableNames != nil
}

// SetTableNames gets a reference to the given []string and assigns it to the TableNames field.
func (o *DmsAudit) SetTableNames(v []string) {
	o.TableNames = v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *DmsAudit) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsAudit) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *DmsAudit) HasErrorMessage() bool {
	return o != nil && o.ErrorMessage != nil
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *DmsAudit) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetRowsAffected returns the RowsAffected field value if set, zero value otherwise.
func (o *DmsAudit) GetRowsAffected() int64 {
	if o == nil || o.RowsAffected == nil {
		var ret int64
		return ret
	}
	return *o.RowsAffected
}

// GetRowsAffectedOk returns a tuple with the RowsAffected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsAudit) GetRowsAffectedOk() (*int64, bool) {
	if o == nil || o.RowsAffected == nil {
		return nil, false
	}
	return o.RowsAffected, true
}

// HasRowsAffected returns a boolean if a field has been set.
func (o *DmsAudit) HasRowsAffected() bool {
	return o != nil && o.RowsAffected != nil
}

// SetRowsAffected gets a reference to the given int64 and assigns it to the RowsAffected field.
func (o *DmsAudit) SetRowsAffected(v int64) {
	o.RowsAffected = &v
}

// GetStatements returns the Statements field value if set, zero value otherwise.
func (o *DmsAudit) GetStatements() []DmsAuditStatement {
	if o == nil || o.Statements == nil {
		var ret []DmsAuditStatement
		return ret
	}
	return o.Statements
}

// GetStatementsOk returns a tuple with the Statements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsAudit) GetStatementsOk() (*[]DmsAuditStatement, bool) {
	if o == nil || o.Statements == nil {
		return nil, false
	}
	return &o.Statements, true
}

// HasStatements returns a boolean if a field has been set.
func (o *DmsAudit) HasStatements() bool {
	return o != nil && o.Statements != nil
}

// SetStatements gets a reference to the given []DmsAuditStatement and assigns it to the Statements field.
func (o *DmsAudit) SetStatements(v []DmsAuditStatement) {
	o.Statements = v
}

// GetStatus returns the Status field value.
func (o *DmsAudit) GetStatus() DmsAuditStatus {
	if o == nil {
		var ret DmsAuditStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DmsAudit) GetStatusOk() (*DmsAuditStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *DmsAudit) SetStatus(v DmsAuditStatus) {
	o.Status = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *DmsAudit) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DmsAudit) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *DmsAudit) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetFinishedAt returns the FinishedAt field value if set, zero value otherwise.
func (o *DmsAudit) GetFinishedAt() time.Time {
	if o == nil || o.FinishedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.FinishedAt
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsAudit) GetFinishedAtOk() (*time.Time, bool) {
	if o == nil || o.FinishedAt == nil {
		return nil, false
	}
	return o.FinishedAt, true
}

// HasFinishedAt returns a boolean if a field has been set.
func (o *DmsAudit) HasFinishedAt() bool {
	return o != nil && o.FinishedAt != nil
}

// SetFinishedAt gets a reference to the given time.Time and assigns it to the FinishedAt field.
func (o *DmsAudit) SetFinishedAt(v time.Time) {
	o.FinishedAt = &v
}

// GetDurationMs returns the DurationMs field value if set, zero value otherwise.
func (o *DmsAudit) GetDurationMs() int64 {
	if o == nil || o.DurationMs == nil {
		var ret int64
		return ret
	}
	return *o.DurationMs
}

// GetDurationMsOk returns a tuple with the DurationMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsAudit) GetDurationMsOk() (*int64, bool) {
	if o == nil || o.DurationMs == nil {
		return nil, false
	}
	return o.DurationMs, true
}

// HasDurationMs returns a boolean if a field has been set.
func (o *DmsAudit) HasDurationMs() bool {
	return o != nil && o.DurationMs != nil
}

// SetDurationMs gets a reference to the given int64 and assigns it to the DurationMs field.
func (o *DmsAudit) SetDurationMs(v int64) {
	o.DurationMs = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DmsAudit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["orgName"] = o.OrgName
	toSerialize["clusterId"] = o.ClusterId
	toSerialize["clusterName"] = o.ClusterName
	if o.DatasourceId != nil {
		toSerialize["datasourceId"] = o.DatasourceId
	}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	toSerialize["userName"] = o.UserName
	toSerialize["engine"] = o.Engine
	toSerialize["database"] = o.Database
	toSerialize["operation"] = o.Operation
	toSerialize["statement"] = o.Statement
	if o.ClientIp != nil {
		toSerialize["clientIp"] = o.ClientIp
	}
	if o.SessionId != nil {
		toSerialize["sessionId"] = o.SessionId
	}
	if o.Risk != nil {
		toSerialize["risk"] = o.Risk
	}
	if o.TableNames != nil {
		toSerialize["tableNames"] = o.TableNames
	}
	if o.ErrorMessage != nil {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if o.RowsAffected != nil {
		toSerialize["rowsAffected"] = o.RowsAffected
	}
	if o.Statements != nil {
		toSerialize["statements"] = o.Statements
	}
	toSerialize["status"] = o.Status
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.FinishedAt != nil {
		if o.FinishedAt.Nanosecond() == 0 {
			toSerialize["finishedAt"] = o.FinishedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["finishedAt"] = o.FinishedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
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
func (o *DmsAudit) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id           *string             `json:"id"`
		OrgName      *string             `json:"orgName"`
		ClusterId    *string             `json:"clusterId"`
		ClusterName  *string             `json:"clusterName"`
		DatasourceId *string             `json:"datasourceId,omitempty"`
		UserId       *string             `json:"userId,omitempty"`
		UserName     *string             `json:"userName"`
		Engine       *string             `json:"engine"`
		Database     *string             `json:"database"`
		Operation    *string             `json:"operation"`
		Statement    *string             `json:"statement"`
		ClientIp     *string             `json:"clientIp,omitempty"`
		SessionId    *string             `json:"sessionId,omitempty"`
		Risk         *DmsAuditRisk       `json:"risk,omitempty"`
		TableNames   []string            `json:"tableNames,omitempty"`
		ErrorMessage *string             `json:"errorMessage,omitempty"`
		RowsAffected *int64              `json:"rowsAffected,omitempty"`
		Statements   []DmsAuditStatement `json:"statements,omitempty"`
		Status       *DmsAuditStatus     `json:"status"`
		CreatedAt    *time.Time          `json:"createdAt"`
		FinishedAt   *time.Time          `json:"finishedAt,omitempty"`
		DurationMs   *int64              `json:"durationMs,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.OrgName == nil {
		return fmt.Errorf("required field orgName missing")
	}
	if all.ClusterId == nil {
		return fmt.Errorf("required field clusterId missing")
	}
	if all.ClusterName == nil {
		return fmt.Errorf("required field clusterName missing")
	}
	if all.UserName == nil {
		return fmt.Errorf("required field userName missing")
	}
	if all.Engine == nil {
		return fmt.Errorf("required field engine missing")
	}
	if all.Database == nil {
		return fmt.Errorf("required field database missing")
	}
	if all.Operation == nil {
		return fmt.Errorf("required field operation missing")
	}
	if all.Statement == nil {
		return fmt.Errorf("required field statement missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field createdAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "orgName", "clusterId", "clusterName", "datasourceId", "userId", "userName", "engine", "database", "operation", "statement", "clientIp", "sessionId", "risk", "tableNames", "errorMessage", "rowsAffected", "statements", "status", "createdAt", "finishedAt", "durationMs"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = *all.Id
	o.OrgName = *all.OrgName
	o.ClusterId = *all.ClusterId
	o.ClusterName = *all.ClusterName
	o.DatasourceId = all.DatasourceId
	o.UserId = all.UserId
	o.UserName = *all.UserName
	o.Engine = *all.Engine
	o.Database = *all.Database
	o.Operation = *all.Operation
	o.Statement = *all.Statement
	o.ClientIp = all.ClientIp
	o.SessionId = all.SessionId
	if all.Risk != nil && !all.Risk.IsValid() {
		hasInvalidField = true
	} else {
		o.Risk = all.Risk
	}
	o.TableNames = all.TableNames
	o.ErrorMessage = all.ErrorMessage
	o.RowsAffected = all.RowsAffected
	o.Statements = all.Statements
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.CreatedAt = *all.CreatedAt
	o.FinishedAt = all.FinishedAt
	o.DurationMs = all.DurationMs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
