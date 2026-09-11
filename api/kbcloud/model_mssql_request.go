// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type MssqlRequest struct {
	RequestId int64 `json:"requestId"`
	// Opaque SQL Server local start time, used to detect session or request reuse. Pass back unchanged; it is not a UTC timestamp.
	StartedAt string `json:"startedAt"`
	Status    string `json:"status"`
	Command   string `json:"command"`
	Database  string `json:"database"`
	// Current statement extracted with byte offsets. Truncated to 8192 characters.
	Sql               string `json:"sql"`
	SqlTruncated      bool   `json:"sqlTruncated"`
	ElapsedMs         int64  `json:"elapsedMs"`
	CpuMs             int64  `json:"cpuMs"`
	Reads             int64  `json:"reads"`
	LogicalReads      int64  `json:"logicalReads"`
	Writes            int64  `json:"writes"`
	WaitType          string `json:"waitType"`
	WaitMs            int64  `json:"waitMs"`
	WaitResource      string `json:"waitResource"`
	BlockingSessionId int64  `json:"blockingSessionId"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMssqlRequest instantiates a new MssqlRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMssqlRequest(requestId int64, startedAt string, status string, command string, database string, sql string, sqlTruncated bool, elapsedMs int64, cpuMs int64, reads int64, logicalReads int64, writes int64, waitType string, waitMs int64, waitResource string, blockingSessionId int64) *MssqlRequest {
	this := MssqlRequest{}
	this.RequestId = requestId
	this.StartedAt = startedAt
	this.Status = status
	this.Command = command
	this.Database = database
	this.Sql = sql
	this.SqlTruncated = sqlTruncated
	this.ElapsedMs = elapsedMs
	this.CpuMs = cpuMs
	this.Reads = reads
	this.LogicalReads = logicalReads
	this.Writes = writes
	this.WaitType = waitType
	this.WaitMs = waitMs
	this.WaitResource = waitResource
	this.BlockingSessionId = blockingSessionId
	return &this
}

// NewMssqlRequestWithDefaults instantiates a new MssqlRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMssqlRequestWithDefaults() *MssqlRequest {
	this := MssqlRequest{}
	return &this
}

// GetRequestId returns the RequestId field value.
func (o *MssqlRequest) GetRequestId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *MssqlRequest) GetRequestIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value.
func (o *MssqlRequest) SetRequestId(v int64) {
	o.RequestId = v
}

// GetStartedAt returns the StartedAt field value.
func (o *MssqlRequest) GetStartedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *MssqlRequest) GetStartedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value.
func (o *MssqlRequest) SetStartedAt(v string) {
	o.StartedAt = v
}

// GetStatus returns the Status field value.
func (o *MssqlRequest) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *MssqlRequest) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *MssqlRequest) SetStatus(v string) {
	o.Status = v
}

// GetCommand returns the Command field value.
func (o *MssqlRequest) GetCommand() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Command
}

// GetCommandOk returns a tuple with the Command field value
// and a boolean to check if the value has been set.
func (o *MssqlRequest) GetCommandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Command, true
}

// SetCommand sets field value.
func (o *MssqlRequest) SetCommand(v string) {
	o.Command = v
}

// GetDatabase returns the Database field value.
func (o *MssqlRequest) GetDatabase() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Database
}

// GetDatabaseOk returns a tuple with the Database field value
// and a boolean to check if the value has been set.
func (o *MssqlRequest) GetDatabaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Database, true
}

// SetDatabase sets field value.
func (o *MssqlRequest) SetDatabase(v string) {
	o.Database = v
}

// GetSql returns the Sql field value.
func (o *MssqlRequest) GetSql() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Sql
}

// GetSqlOk returns a tuple with the Sql field value
// and a boolean to check if the value has been set.
func (o *MssqlRequest) GetSqlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sql, true
}

// SetSql sets field value.
func (o *MssqlRequest) SetSql(v string) {
	o.Sql = v
}

// GetSqlTruncated returns the SqlTruncated field value.
func (o *MssqlRequest) GetSqlTruncated() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.SqlTruncated
}

// GetSqlTruncatedOk returns a tuple with the SqlTruncated field value
// and a boolean to check if the value has been set.
func (o *MssqlRequest) GetSqlTruncatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SqlTruncated, true
}

// SetSqlTruncated sets field value.
func (o *MssqlRequest) SetSqlTruncated(v bool) {
	o.SqlTruncated = v
}

// GetElapsedMs returns the ElapsedMs field value.
func (o *MssqlRequest) GetElapsedMs() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ElapsedMs
}

// GetElapsedMsOk returns a tuple with the ElapsedMs field value
// and a boolean to check if the value has been set.
func (o *MssqlRequest) GetElapsedMsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ElapsedMs, true
}

// SetElapsedMs sets field value.
func (o *MssqlRequest) SetElapsedMs(v int64) {
	o.ElapsedMs = v
}

// GetCpuMs returns the CpuMs field value.
func (o *MssqlRequest) GetCpuMs() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.CpuMs
}

// GetCpuMsOk returns a tuple with the CpuMs field value
// and a boolean to check if the value has been set.
func (o *MssqlRequest) GetCpuMsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CpuMs, true
}

// SetCpuMs sets field value.
func (o *MssqlRequest) SetCpuMs(v int64) {
	o.CpuMs = v
}

// GetReads returns the Reads field value.
func (o *MssqlRequest) GetReads() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Reads
}

// GetReadsOk returns a tuple with the Reads field value
// and a boolean to check if the value has been set.
func (o *MssqlRequest) GetReadsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reads, true
}

// SetReads sets field value.
func (o *MssqlRequest) SetReads(v int64) {
	o.Reads = v
}

// GetLogicalReads returns the LogicalReads field value.
func (o *MssqlRequest) GetLogicalReads() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.LogicalReads
}

// GetLogicalReadsOk returns a tuple with the LogicalReads field value
// and a boolean to check if the value has been set.
func (o *MssqlRequest) GetLogicalReadsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogicalReads, true
}

// SetLogicalReads sets field value.
func (o *MssqlRequest) SetLogicalReads(v int64) {
	o.LogicalReads = v
}

// GetWrites returns the Writes field value.
func (o *MssqlRequest) GetWrites() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Writes
}

// GetWritesOk returns a tuple with the Writes field value
// and a boolean to check if the value has been set.
func (o *MssqlRequest) GetWritesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Writes, true
}

// SetWrites sets field value.
func (o *MssqlRequest) SetWrites(v int64) {
	o.Writes = v
}

// GetWaitType returns the WaitType field value.
func (o *MssqlRequest) GetWaitType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.WaitType
}

// GetWaitTypeOk returns a tuple with the WaitType field value
// and a boolean to check if the value has been set.
func (o *MssqlRequest) GetWaitTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WaitType, true
}

// SetWaitType sets field value.
func (o *MssqlRequest) SetWaitType(v string) {
	o.WaitType = v
}

// GetWaitMs returns the WaitMs field value.
func (o *MssqlRequest) GetWaitMs() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.WaitMs
}

// GetWaitMsOk returns a tuple with the WaitMs field value
// and a boolean to check if the value has been set.
func (o *MssqlRequest) GetWaitMsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WaitMs, true
}

// SetWaitMs sets field value.
func (o *MssqlRequest) SetWaitMs(v int64) {
	o.WaitMs = v
}

// GetWaitResource returns the WaitResource field value.
func (o *MssqlRequest) GetWaitResource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.WaitResource
}

// GetWaitResourceOk returns a tuple with the WaitResource field value
// and a boolean to check if the value has been set.
func (o *MssqlRequest) GetWaitResourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WaitResource, true
}

// SetWaitResource sets field value.
func (o *MssqlRequest) SetWaitResource(v string) {
	o.WaitResource = v
}

// GetBlockingSessionId returns the BlockingSessionId field value.
func (o *MssqlRequest) GetBlockingSessionId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.BlockingSessionId
}

// GetBlockingSessionIdOk returns a tuple with the BlockingSessionId field value
// and a boolean to check if the value has been set.
func (o *MssqlRequest) GetBlockingSessionIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockingSessionId, true
}

// SetBlockingSessionId sets field value.
func (o *MssqlRequest) SetBlockingSessionId(v int64) {
	o.BlockingSessionId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MssqlRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["requestId"] = o.RequestId
	toSerialize["startedAt"] = o.StartedAt
	toSerialize["status"] = o.Status
	toSerialize["command"] = o.Command
	toSerialize["database"] = o.Database
	toSerialize["sql"] = o.Sql
	toSerialize["sqlTruncated"] = o.SqlTruncated
	toSerialize["elapsedMs"] = o.ElapsedMs
	toSerialize["cpuMs"] = o.CpuMs
	toSerialize["reads"] = o.Reads
	toSerialize["logicalReads"] = o.LogicalReads
	toSerialize["writes"] = o.Writes
	toSerialize["waitType"] = o.WaitType
	toSerialize["waitMs"] = o.WaitMs
	toSerialize["waitResource"] = o.WaitResource
	toSerialize["blockingSessionId"] = o.BlockingSessionId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MssqlRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		RequestId         *int64  `json:"requestId"`
		StartedAt         *string `json:"startedAt"`
		Status            *string `json:"status"`
		Command           *string `json:"command"`
		Database          *string `json:"database"`
		Sql               *string `json:"sql"`
		SqlTruncated      *bool   `json:"sqlTruncated"`
		ElapsedMs         *int64  `json:"elapsedMs"`
		CpuMs             *int64  `json:"cpuMs"`
		Reads             *int64  `json:"reads"`
		LogicalReads      *int64  `json:"logicalReads"`
		Writes            *int64  `json:"writes"`
		WaitType          *string `json:"waitType"`
		WaitMs            *int64  `json:"waitMs"`
		WaitResource      *string `json:"waitResource"`
		BlockingSessionId *int64  `json:"blockingSessionId"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.RequestId == nil {
		return fmt.Errorf("required field requestId missing")
	}
	if all.StartedAt == nil {
		return fmt.Errorf("required field startedAt missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.Command == nil {
		return fmt.Errorf("required field command missing")
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
	if all.ElapsedMs == nil {
		return fmt.Errorf("required field elapsedMs missing")
	}
	if all.CpuMs == nil {
		return fmt.Errorf("required field cpuMs missing")
	}
	if all.Reads == nil {
		return fmt.Errorf("required field reads missing")
	}
	if all.LogicalReads == nil {
		return fmt.Errorf("required field logicalReads missing")
	}
	if all.Writes == nil {
		return fmt.Errorf("required field writes missing")
	}
	if all.WaitType == nil {
		return fmt.Errorf("required field waitType missing")
	}
	if all.WaitMs == nil {
		return fmt.Errorf("required field waitMs missing")
	}
	if all.WaitResource == nil {
		return fmt.Errorf("required field waitResource missing")
	}
	if all.BlockingSessionId == nil {
		return fmt.Errorf("required field blockingSessionId missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"requestId", "startedAt", "status", "command", "database", "sql", "sqlTruncated", "elapsedMs", "cpuMs", "reads", "logicalReads", "writes", "waitType", "waitMs", "waitResource", "blockingSessionId"})
	} else {
		return err
	}
	o.RequestId = *all.RequestId
	o.StartedAt = *all.StartedAt
	o.Status = *all.Status
	o.Command = *all.Command
	o.Database = *all.Database
	o.Sql = *all.Sql
	o.SqlTruncated = *all.SqlTruncated
	o.ElapsedMs = *all.ElapsedMs
	o.CpuMs = *all.CpuMs
	o.Reads = *all.Reads
	o.LogicalReads = *all.LogicalReads
	o.Writes = *all.Writes
	o.WaitType = *all.WaitType
	o.WaitMs = *all.WaitMs
	o.WaitResource = *all.WaitResource
	o.BlockingSessionId = *all.BlockingSessionId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
