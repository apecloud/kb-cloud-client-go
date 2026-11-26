// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type MssqlSQLStatement struct {
	Id             string  `json:"id"`
	QueryHash      string  `json:"queryHash"`
	Database       string  `json:"database"`
	Sql            string  `json:"sql"`
	SqlTruncated   bool    `json:"sqlTruncated"`
	CreatedAt      string  `json:"createdAt"`
	LastExecutedAt string  `json:"lastExecutedAt"`
	Executions     int64   `json:"executions"`
	TotalElapsedMs float64 `json:"totalElapsedMs"`
	MeanElapsedMs  float64 `json:"meanElapsedMs"`
	MaxElapsedMs   float64 `json:"maxElapsedMs"`
	TotalCpuMs     float64 `json:"totalCpuMs"`
	MeanCpuMs      float64 `json:"meanCpuMs"`
	LogicalReads   int64   `json:"logicalReads"`
	PhysicalReads  int64   `json:"physicalReads"`
	LogicalWrites  int64   `json:"logicalWrites"`
	Rows           int64   `json:"rows"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMssqlSQLStatement instantiates a new MssqlSQLStatement object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMssqlSQLStatement(id string, queryHash string, database string, sql string, sqlTruncated bool, createdAt string, lastExecutedAt string, executions int64, totalElapsedMs float64, meanElapsedMs float64, maxElapsedMs float64, totalCpuMs float64, meanCpuMs float64, logicalReads int64, physicalReads int64, logicalWrites int64, rows int64) *MssqlSQLStatement {
	this := MssqlSQLStatement{}
	this.Id = id
	this.QueryHash = queryHash
	this.Database = database
	this.Sql = sql
	this.SqlTruncated = sqlTruncated
	this.CreatedAt = createdAt
	this.LastExecutedAt = lastExecutedAt
	this.Executions = executions
	this.TotalElapsedMs = totalElapsedMs
	this.MeanElapsedMs = meanElapsedMs
	this.MaxElapsedMs = maxElapsedMs
	this.TotalCpuMs = totalCpuMs
	this.MeanCpuMs = meanCpuMs
	this.LogicalReads = logicalReads
	this.PhysicalReads = physicalReads
	this.LogicalWrites = logicalWrites
	this.Rows = rows
	return &this
}

// NewMssqlSQLStatementWithDefaults instantiates a new MssqlSQLStatement object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMssqlSQLStatementWithDefaults() *MssqlSQLStatement {
	this := MssqlSQLStatement{}
	return &this
}

// GetId returns the Id field value.
func (o *MssqlSQLStatement) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MssqlSQLStatement) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *MssqlSQLStatement) SetId(v string) {
	o.Id = v
}

// GetQueryHash returns the QueryHash field value.
func (o *MssqlSQLStatement) GetQueryHash() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.QueryHash
}

// GetQueryHashOk returns a tuple with the QueryHash field value
// and a boolean to check if the value has been set.
func (o *MssqlSQLStatement) GetQueryHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryHash, true
}

// SetQueryHash sets field value.
func (o *MssqlSQLStatement) SetQueryHash(v string) {
	o.QueryHash = v
}

// GetDatabase returns the Database field value.
func (o *MssqlSQLStatement) GetDatabase() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Database
}

// GetDatabaseOk returns a tuple with the Database field value
// and a boolean to check if the value has been set.
func (o *MssqlSQLStatement) GetDatabaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Database, true
}

// SetDatabase sets field value.
func (o *MssqlSQLStatement) SetDatabase(v string) {
	o.Database = v
}

// GetSql returns the Sql field value.
func (o *MssqlSQLStatement) GetSql() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Sql
}

// GetSqlOk returns a tuple with the Sql field value
// and a boolean to check if the value has been set.
func (o *MssqlSQLStatement) GetSqlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sql, true
}

// SetSql sets field value.
func (o *MssqlSQLStatement) SetSql(v string) {
	o.Sql = v
}

// GetSqlTruncated returns the SqlTruncated field value.
func (o *MssqlSQLStatement) GetSqlTruncated() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.SqlTruncated
}

// GetSqlTruncatedOk returns a tuple with the SqlTruncated field value
// and a boolean to check if the value has been set.
func (o *MssqlSQLStatement) GetSqlTruncatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SqlTruncated, true
}

// SetSqlTruncated sets field value.
func (o *MssqlSQLStatement) SetSqlTruncated(v bool) {
	o.SqlTruncated = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *MssqlSQLStatement) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *MssqlSQLStatement) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *MssqlSQLStatement) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetLastExecutedAt returns the LastExecutedAt field value.
func (o *MssqlSQLStatement) GetLastExecutedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.LastExecutedAt
}

// GetLastExecutedAtOk returns a tuple with the LastExecutedAt field value
// and a boolean to check if the value has been set.
func (o *MssqlSQLStatement) GetLastExecutedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastExecutedAt, true
}

// SetLastExecutedAt sets field value.
func (o *MssqlSQLStatement) SetLastExecutedAt(v string) {
	o.LastExecutedAt = v
}

// GetExecutions returns the Executions field value.
func (o *MssqlSQLStatement) GetExecutions() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Executions
}

// GetExecutionsOk returns a tuple with the Executions field value
// and a boolean to check if the value has been set.
func (o *MssqlSQLStatement) GetExecutionsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Executions, true
}

// SetExecutions sets field value.
func (o *MssqlSQLStatement) SetExecutions(v int64) {
	o.Executions = v
}

// GetTotalElapsedMs returns the TotalElapsedMs field value.
func (o *MssqlSQLStatement) GetTotalElapsedMs() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.TotalElapsedMs
}

// GetTotalElapsedMsOk returns a tuple with the TotalElapsedMs field value
// and a boolean to check if the value has been set.
func (o *MssqlSQLStatement) GetTotalElapsedMsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalElapsedMs, true
}

// SetTotalElapsedMs sets field value.
func (o *MssqlSQLStatement) SetTotalElapsedMs(v float64) {
	o.TotalElapsedMs = v
}

// GetMeanElapsedMs returns the MeanElapsedMs field value.
func (o *MssqlSQLStatement) GetMeanElapsedMs() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.MeanElapsedMs
}

// GetMeanElapsedMsOk returns a tuple with the MeanElapsedMs field value
// and a boolean to check if the value has been set.
func (o *MssqlSQLStatement) GetMeanElapsedMsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MeanElapsedMs, true
}

// SetMeanElapsedMs sets field value.
func (o *MssqlSQLStatement) SetMeanElapsedMs(v float64) {
	o.MeanElapsedMs = v
}

// GetMaxElapsedMs returns the MaxElapsedMs field value.
func (o *MssqlSQLStatement) GetMaxElapsedMs() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.MaxElapsedMs
}

// GetMaxElapsedMsOk returns a tuple with the MaxElapsedMs field value
// and a boolean to check if the value has been set.
func (o *MssqlSQLStatement) GetMaxElapsedMsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxElapsedMs, true
}

// SetMaxElapsedMs sets field value.
func (o *MssqlSQLStatement) SetMaxElapsedMs(v float64) {
	o.MaxElapsedMs = v
}

// GetTotalCpuMs returns the TotalCpuMs field value.
func (o *MssqlSQLStatement) GetTotalCpuMs() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.TotalCpuMs
}

// GetTotalCpuMsOk returns a tuple with the TotalCpuMs field value
// and a boolean to check if the value has been set.
func (o *MssqlSQLStatement) GetTotalCpuMsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCpuMs, true
}

// SetTotalCpuMs sets field value.
func (o *MssqlSQLStatement) SetTotalCpuMs(v float64) {
	o.TotalCpuMs = v
}

// GetMeanCpuMs returns the MeanCpuMs field value.
func (o *MssqlSQLStatement) GetMeanCpuMs() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.MeanCpuMs
}

// GetMeanCpuMsOk returns a tuple with the MeanCpuMs field value
// and a boolean to check if the value has been set.
func (o *MssqlSQLStatement) GetMeanCpuMsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MeanCpuMs, true
}

// SetMeanCpuMs sets field value.
func (o *MssqlSQLStatement) SetMeanCpuMs(v float64) {
	o.MeanCpuMs = v
}

// GetLogicalReads returns the LogicalReads field value.
func (o *MssqlSQLStatement) GetLogicalReads() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.LogicalReads
}

// GetLogicalReadsOk returns a tuple with the LogicalReads field value
// and a boolean to check if the value has been set.
func (o *MssqlSQLStatement) GetLogicalReadsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogicalReads, true
}

// SetLogicalReads sets field value.
func (o *MssqlSQLStatement) SetLogicalReads(v int64) {
	o.LogicalReads = v
}

// GetPhysicalReads returns the PhysicalReads field value.
func (o *MssqlSQLStatement) GetPhysicalReads() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.PhysicalReads
}

// GetPhysicalReadsOk returns a tuple with the PhysicalReads field value
// and a boolean to check if the value has been set.
func (o *MssqlSQLStatement) GetPhysicalReadsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PhysicalReads, true
}

// SetPhysicalReads sets field value.
func (o *MssqlSQLStatement) SetPhysicalReads(v int64) {
	o.PhysicalReads = v
}

// GetLogicalWrites returns the LogicalWrites field value.
func (o *MssqlSQLStatement) GetLogicalWrites() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.LogicalWrites
}

// GetLogicalWritesOk returns a tuple with the LogicalWrites field value
// and a boolean to check if the value has been set.
func (o *MssqlSQLStatement) GetLogicalWritesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogicalWrites, true
}

// SetLogicalWrites sets field value.
func (o *MssqlSQLStatement) SetLogicalWrites(v int64) {
	o.LogicalWrites = v
}

// GetRows returns the Rows field value.
func (o *MssqlSQLStatement) GetRows() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value
// and a boolean to check if the value has been set.
func (o *MssqlSQLStatement) GetRowsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rows, true
}

// SetRows sets field value.
func (o *MssqlSQLStatement) SetRows(v int64) {
	o.Rows = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MssqlSQLStatement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["queryHash"] = o.QueryHash
	toSerialize["database"] = o.Database
	toSerialize["sql"] = o.Sql
	toSerialize["sqlTruncated"] = o.SqlTruncated
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["lastExecutedAt"] = o.LastExecutedAt
	toSerialize["executions"] = o.Executions
	toSerialize["totalElapsedMs"] = o.TotalElapsedMs
	toSerialize["meanElapsedMs"] = o.MeanElapsedMs
	toSerialize["maxElapsedMs"] = o.MaxElapsedMs
	toSerialize["totalCpuMs"] = o.TotalCpuMs
	toSerialize["meanCpuMs"] = o.MeanCpuMs
	toSerialize["logicalReads"] = o.LogicalReads
	toSerialize["physicalReads"] = o.PhysicalReads
	toSerialize["logicalWrites"] = o.LogicalWrites
	toSerialize["rows"] = o.Rows

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MssqlSQLStatement) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id             *string  `json:"id"`
		QueryHash      *string  `json:"queryHash"`
		Database       *string  `json:"database"`
		Sql            *string  `json:"sql"`
		SqlTruncated   *bool    `json:"sqlTruncated"`
		CreatedAt      *string  `json:"createdAt"`
		LastExecutedAt *string  `json:"lastExecutedAt"`
		Executions     *int64   `json:"executions"`
		TotalElapsedMs *float64 `json:"totalElapsedMs"`
		MeanElapsedMs  *float64 `json:"meanElapsedMs"`
		MaxElapsedMs   *float64 `json:"maxElapsedMs"`
		TotalCpuMs     *float64 `json:"totalCpuMs"`
		MeanCpuMs      *float64 `json:"meanCpuMs"`
		LogicalReads   *int64   `json:"logicalReads"`
		PhysicalReads  *int64   `json:"physicalReads"`
		LogicalWrites  *int64   `json:"logicalWrites"`
		Rows           *int64   `json:"rows"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.QueryHash == nil {
		return fmt.Errorf("required field queryHash missing")
	}
	if all.Database == nil {
		return fmt.Errorf("required field database missing")
	}
	if all.Sql == nil {
		return fmt.Errorf("required field sql missing")
	}
	if all.SqlTruncated == nil {
		return fmt.Errorf("required field sqlTruncated missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field createdAt missing")
	}
	if all.LastExecutedAt == nil {
		return fmt.Errorf("required field lastExecutedAt missing")
	}
	if all.Executions == nil {
		return fmt.Errorf("required field executions missing")
	}
	if all.TotalElapsedMs == nil {
		return fmt.Errorf("required field totalElapsedMs missing")
	}
	if all.MeanElapsedMs == nil {
		return fmt.Errorf("required field meanElapsedMs missing")
	}
	if all.MaxElapsedMs == nil {
		return fmt.Errorf("required field maxElapsedMs missing")
	}
	if all.TotalCpuMs == nil {
		return fmt.Errorf("required field totalCpuMs missing")
	}
	if all.MeanCpuMs == nil {
		return fmt.Errorf("required field meanCpuMs missing")
	}
	if all.LogicalReads == nil {
		return fmt.Errorf("required field logicalReads missing")
	}
	if all.PhysicalReads == nil {
		return fmt.Errorf("required field physicalReads missing")
	}
	if all.LogicalWrites == nil {
		return fmt.Errorf("required field logicalWrites missing")
	}
	if all.Rows == nil {
		return fmt.Errorf("required field rows missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "queryHash", "database", "sql", "sqlTruncated", "createdAt", "lastExecutedAt", "executions", "totalElapsedMs", "meanElapsedMs", "maxElapsedMs", "totalCpuMs", "meanCpuMs", "logicalReads", "physicalReads", "logicalWrites", "rows"})
	} else {
		return err
	}
	o.Id = *all.Id
	o.QueryHash = *all.QueryHash
	o.Database = *all.Database
	o.Sql = *all.Sql
	o.SqlTruncated = *all.SqlTruncated
	o.CreatedAt = *all.CreatedAt
	o.LastExecutedAt = *all.LastExecutedAt
	o.Executions = *all.Executions
	o.TotalElapsedMs = *all.TotalElapsedMs
	o.MeanElapsedMs = *all.MeanElapsedMs
	o.MaxElapsedMs = *all.MaxElapsedMs
	o.TotalCpuMs = *all.TotalCpuMs
	o.MeanCpuMs = *all.MeanCpuMs
	o.LogicalReads = *all.LogicalReads
	o.PhysicalReads = *all.PhysicalReads
	o.LogicalWrites = *all.LogicalWrites
	o.Rows = *all.Rows

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
