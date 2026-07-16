// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DamengLargeMemSQLEntry struct {
	// SQL text from V$SYSTEM_LARGE_MEM_SQLS.
	SqlText string `json:"sqlText"`
	// Memory used in kilobytes.
	MemUsedKb int64 `json:"memUsedKb"`
	// Session ID, when available.
	SessionId *string `json:"sessionId,omitempty"`
	// SQL ID, when available.
	SqlId *string `json:"sqlId,omitempty"`
	// SQL finish time in UTC, when available.
	FinishTime *string `json:"finishTime,omitempty"`
	// Number of executions.
	NRuns int64 `json:"nRuns"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDamengLargeMemSQLEntry instantiates a new DamengLargeMemSQLEntry object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDamengLargeMemSQLEntry(sqlText string, memUsedKb int64, nRuns int64) *DamengLargeMemSQLEntry {
	this := DamengLargeMemSQLEntry{}
	this.SqlText = sqlText
	this.MemUsedKb = memUsedKb
	this.NRuns = nRuns
	return &this
}

// NewDamengLargeMemSQLEntryWithDefaults instantiates a new DamengLargeMemSQLEntry object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDamengLargeMemSQLEntryWithDefaults() *DamengLargeMemSQLEntry {
	this := DamengLargeMemSQLEntry{}
	return &this
}

// GetSqlText returns the SqlText field value.
func (o *DamengLargeMemSQLEntry) GetSqlText() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SqlText
}

// GetSqlTextOk returns a tuple with the SqlText field value
// and a boolean to check if the value has been set.
func (o *DamengLargeMemSQLEntry) GetSqlTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SqlText, true
}

// SetSqlText sets field value.
func (o *DamengLargeMemSQLEntry) SetSqlText(v string) {
	o.SqlText = v
}

// GetMemUsedKb returns the MemUsedKb field value.
func (o *DamengLargeMemSQLEntry) GetMemUsedKb() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.MemUsedKb
}

// GetMemUsedKbOk returns a tuple with the MemUsedKb field value
// and a boolean to check if the value has been set.
func (o *DamengLargeMemSQLEntry) GetMemUsedKbOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MemUsedKb, true
}

// SetMemUsedKb sets field value.
func (o *DamengLargeMemSQLEntry) SetMemUsedKb(v int64) {
	o.MemUsedKb = v
}

// GetSessionId returns the SessionId field value if set, zero value otherwise.
func (o *DamengLargeMemSQLEntry) GetSessionId() string {
	if o == nil || o.SessionId == nil {
		var ret string
		return ret
	}
	return *o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamengLargeMemSQLEntry) GetSessionIdOk() (*string, bool) {
	if o == nil || o.SessionId == nil {
		return nil, false
	}
	return o.SessionId, true
}

// HasSessionId returns a boolean if a field has been set.
func (o *DamengLargeMemSQLEntry) HasSessionId() bool {
	return o != nil && o.SessionId != nil
}

// SetSessionId gets a reference to the given string and assigns it to the SessionId field.
func (o *DamengLargeMemSQLEntry) SetSessionId(v string) {
	o.SessionId = &v
}

// GetSqlId returns the SqlId field value if set, zero value otherwise.
func (o *DamengLargeMemSQLEntry) GetSqlId() string {
	if o == nil || o.SqlId == nil {
		var ret string
		return ret
	}
	return *o.SqlId
}

// GetSqlIdOk returns a tuple with the SqlId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamengLargeMemSQLEntry) GetSqlIdOk() (*string, bool) {
	if o == nil || o.SqlId == nil {
		return nil, false
	}
	return o.SqlId, true
}

// HasSqlId returns a boolean if a field has been set.
func (o *DamengLargeMemSQLEntry) HasSqlId() bool {
	return o != nil && o.SqlId != nil
}

// SetSqlId gets a reference to the given string and assigns it to the SqlId field.
func (o *DamengLargeMemSQLEntry) SetSqlId(v string) {
	o.SqlId = &v
}

// GetFinishTime returns the FinishTime field value if set, zero value otherwise.
func (o *DamengLargeMemSQLEntry) GetFinishTime() string {
	if o == nil || o.FinishTime == nil {
		var ret string
		return ret
	}
	return *o.FinishTime
}

// GetFinishTimeOk returns a tuple with the FinishTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamengLargeMemSQLEntry) GetFinishTimeOk() (*string, bool) {
	if o == nil || o.FinishTime == nil {
		return nil, false
	}
	return o.FinishTime, true
}

// HasFinishTime returns a boolean if a field has been set.
func (o *DamengLargeMemSQLEntry) HasFinishTime() bool {
	return o != nil && o.FinishTime != nil
}

// SetFinishTime gets a reference to the given string and assigns it to the FinishTime field.
func (o *DamengLargeMemSQLEntry) SetFinishTime(v string) {
	o.FinishTime = &v
}

// GetNRuns returns the NRuns field value.
func (o *DamengLargeMemSQLEntry) GetNRuns() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.NRuns
}

// GetNRunsOk returns a tuple with the NRuns field value
// and a boolean to check if the value has been set.
func (o *DamengLargeMemSQLEntry) GetNRunsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NRuns, true
}

// SetNRuns sets field value.
func (o *DamengLargeMemSQLEntry) SetNRuns(v int64) {
	o.NRuns = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DamengLargeMemSQLEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["sqlText"] = o.SqlText
	toSerialize["memUsedKb"] = o.MemUsedKb
	if o.SessionId != nil {
		toSerialize["sessionId"] = o.SessionId
	}
	if o.SqlId != nil {
		toSerialize["sqlId"] = o.SqlId
	}
	if o.FinishTime != nil {
		toSerialize["finishTime"] = o.FinishTime
	}
	toSerialize["nRuns"] = o.NRuns

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DamengLargeMemSQLEntry) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SqlText    *string `json:"sqlText"`
		MemUsedKb  *int64  `json:"memUsedKb"`
		SessionId  *string `json:"sessionId,omitempty"`
		SqlId      *string `json:"sqlId,omitempty"`
		FinishTime *string `json:"finishTime,omitempty"`
		NRuns      *int64  `json:"nRuns"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.SqlText == nil {
		return fmt.Errorf("required field sqlText missing")
	}
	if all.MemUsedKb == nil {
		return fmt.Errorf("required field memUsedKb missing")
	}
	if all.NRuns == nil {
		return fmt.Errorf("required field nRuns missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"sqlText", "memUsedKb", "sessionId", "sqlId", "finishTime", "nRuns"})
	} else {
		return err
	}
	o.SqlText = *all.SqlText
	o.MemUsedKb = *all.MemUsedKb
	o.SessionId = all.SessionId
	o.SqlId = all.SqlId
	o.FinishTime = all.FinishTime
	o.NRuns = *all.NRuns

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
