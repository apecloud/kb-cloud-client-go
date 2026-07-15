// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DamengLongExecSQLEntry struct {
	// SQL text from V$SYSTEM_LONG_EXEC_SQLS.
	SqlText string `json:"sqlText"`
	// Execution time in milliseconds.
	ExecTimeMs int64 `json:"execTimeMs"`
	// Session ID, when available.
	SessionId *string `json:"sessionId,omitempty"`
	// SQL ID, when available.
	SqlId *string `json:"sqlId,omitempty"`
	// Number of executions.
	NRuns int64 `json:"nRuns"`
	// SQL finish time in UTC, when available.
	FinishTime *string `json:"finishTime,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDamengLongExecSQLEntry instantiates a new DamengLongExecSQLEntry object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDamengLongExecSQLEntry(sqlText string, execTimeMs int64, nRuns int64) *DamengLongExecSQLEntry {
	this := DamengLongExecSQLEntry{}
	this.SqlText = sqlText
	this.ExecTimeMs = execTimeMs
	this.NRuns = nRuns
	return &this
}

// NewDamengLongExecSQLEntryWithDefaults instantiates a new DamengLongExecSQLEntry object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDamengLongExecSQLEntryWithDefaults() *DamengLongExecSQLEntry {
	this := DamengLongExecSQLEntry{}
	return &this
}

// GetSqlText returns the SqlText field value.
func (o *DamengLongExecSQLEntry) GetSqlText() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SqlText
}

// GetSqlTextOk returns a tuple with the SqlText field value
// and a boolean to check if the value has been set.
func (o *DamengLongExecSQLEntry) GetSqlTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SqlText, true
}

// SetSqlText sets field value.
func (o *DamengLongExecSQLEntry) SetSqlText(v string) {
	o.SqlText = v
}

// GetExecTimeMs returns the ExecTimeMs field value.
func (o *DamengLongExecSQLEntry) GetExecTimeMs() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ExecTimeMs
}

// GetExecTimeMsOk returns a tuple with the ExecTimeMs field value
// and a boolean to check if the value has been set.
func (o *DamengLongExecSQLEntry) GetExecTimeMsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecTimeMs, true
}

// SetExecTimeMs sets field value.
func (o *DamengLongExecSQLEntry) SetExecTimeMs(v int64) {
	o.ExecTimeMs = v
}

// GetSessionId returns the SessionId field value if set, zero value otherwise.
func (o *DamengLongExecSQLEntry) GetSessionId() string {
	if o == nil || o.SessionId == nil {
		var ret string
		return ret
	}
	return *o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamengLongExecSQLEntry) GetSessionIdOk() (*string, bool) {
	if o == nil || o.SessionId == nil {
		return nil, false
	}
	return o.SessionId, true
}

// HasSessionId returns a boolean if a field has been set.
func (o *DamengLongExecSQLEntry) HasSessionId() bool {
	return o != nil && o.SessionId != nil
}

// SetSessionId gets a reference to the given string and assigns it to the SessionId field.
func (o *DamengLongExecSQLEntry) SetSessionId(v string) {
	o.SessionId = &v
}

// GetSqlId returns the SqlId field value if set, zero value otherwise.
func (o *DamengLongExecSQLEntry) GetSqlId() string {
	if o == nil || o.SqlId == nil {
		var ret string
		return ret
	}
	return *o.SqlId
}

// GetSqlIdOk returns a tuple with the SqlId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamengLongExecSQLEntry) GetSqlIdOk() (*string, bool) {
	if o == nil || o.SqlId == nil {
		return nil, false
	}
	return o.SqlId, true
}

// HasSqlId returns a boolean if a field has been set.
func (o *DamengLongExecSQLEntry) HasSqlId() bool {
	return o != nil && o.SqlId != nil
}

// SetSqlId gets a reference to the given string and assigns it to the SqlId field.
func (o *DamengLongExecSQLEntry) SetSqlId(v string) {
	o.SqlId = &v
}

// GetNRuns returns the NRuns field value.
func (o *DamengLongExecSQLEntry) GetNRuns() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.NRuns
}

// GetNRunsOk returns a tuple with the NRuns field value
// and a boolean to check if the value has been set.
func (o *DamengLongExecSQLEntry) GetNRunsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NRuns, true
}

// SetNRuns sets field value.
func (o *DamengLongExecSQLEntry) SetNRuns(v int64) {
	o.NRuns = v
}

// GetFinishTime returns the FinishTime field value if set, zero value otherwise.
func (o *DamengLongExecSQLEntry) GetFinishTime() string {
	if o == nil || o.FinishTime == nil {
		var ret string
		return ret
	}
	return *o.FinishTime
}

// GetFinishTimeOk returns a tuple with the FinishTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamengLongExecSQLEntry) GetFinishTimeOk() (*string, bool) {
	if o == nil || o.FinishTime == nil {
		return nil, false
	}
	return o.FinishTime, true
}

// HasFinishTime returns a boolean if a field has been set.
func (o *DamengLongExecSQLEntry) HasFinishTime() bool {
	return o != nil && o.FinishTime != nil
}

// SetFinishTime gets a reference to the given string and assigns it to the FinishTime field.
func (o *DamengLongExecSQLEntry) SetFinishTime(v string) {
	o.FinishTime = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DamengLongExecSQLEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["sqlText"] = o.SqlText
	toSerialize["execTimeMs"] = o.ExecTimeMs
	if o.SessionId != nil {
		toSerialize["sessionId"] = o.SessionId
	}
	if o.SqlId != nil {
		toSerialize["sqlId"] = o.SqlId
	}
	toSerialize["nRuns"] = o.NRuns
	if o.FinishTime != nil {
		toSerialize["finishTime"] = o.FinishTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DamengLongExecSQLEntry) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SqlText    *string `json:"sqlText"`
		ExecTimeMs *int64  `json:"execTimeMs"`
		SessionId  *string `json:"sessionId,omitempty"`
		SqlId      *string `json:"sqlId,omitempty"`
		NRuns      *int64  `json:"nRuns"`
		FinishTime *string `json:"finishTime,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.SqlText == nil {
		return fmt.Errorf("required field sqlText missing")
	}
	if all.ExecTimeMs == nil {
		return fmt.Errorf("required field execTimeMs missing")
	}
	if all.NRuns == nil {
		return fmt.Errorf("required field nRuns missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"sqlText", "execTimeMs", "sessionId", "sqlId", "nRuns", "finishTime"})
	} else {
		return err
	}
	o.SqlText = *all.SqlText
	o.ExecTimeMs = *all.ExecTimeMs
	o.SessionId = all.SessionId
	o.SqlId = all.SqlId
	o.NRuns = *all.NRuns
	o.FinishTime = all.FinishTime

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
