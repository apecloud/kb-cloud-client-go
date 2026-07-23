// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DamengSQLAnalysis struct {
	// Top long-running SQL entries from V$SYSTEM_LONG_EXEC_SQLS ordered by execution time descending.
	TopLongExecSql []DamengLongExecSQLEntry `json:"topLongExecSql"`
	// Top high-memory SQL entries from V$SYSTEM_LARGE_MEM_SQLS ordered by memory usage descending.
	TopLargeMemSql []DamengLargeMemSQLEntry `json:"topLargeMemSql"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDamengSQLAnalysis instantiates a new DamengSQLAnalysis object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDamengSQLAnalysis(topLongExecSql []DamengLongExecSQLEntry, topLargeMemSql []DamengLargeMemSQLEntry) *DamengSQLAnalysis {
	this := DamengSQLAnalysis{}
	this.TopLongExecSql = topLongExecSql
	this.TopLargeMemSql = topLargeMemSql
	return &this
}

// NewDamengSQLAnalysisWithDefaults instantiates a new DamengSQLAnalysis object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDamengSQLAnalysisWithDefaults() *DamengSQLAnalysis {
	this := DamengSQLAnalysis{}
	return &this
}

// GetTopLongExecSql returns the TopLongExecSql field value.
func (o *DamengSQLAnalysis) GetTopLongExecSql() []DamengLongExecSQLEntry {
	if o == nil {
		var ret []DamengLongExecSQLEntry
		return ret
	}
	return o.TopLongExecSql
}

// GetTopLongExecSqlOk returns a tuple with the TopLongExecSql field value
// and a boolean to check if the value has been set.
func (o *DamengSQLAnalysis) GetTopLongExecSqlOk() (*[]DamengLongExecSQLEntry, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TopLongExecSql, true
}

// SetTopLongExecSql sets field value.
func (o *DamengSQLAnalysis) SetTopLongExecSql(v []DamengLongExecSQLEntry) {
	o.TopLongExecSql = v
}

// GetTopLargeMemSql returns the TopLargeMemSql field value.
func (o *DamengSQLAnalysis) GetTopLargeMemSql() []DamengLargeMemSQLEntry {
	if o == nil {
		var ret []DamengLargeMemSQLEntry
		return ret
	}
	return o.TopLargeMemSql
}

// GetTopLargeMemSqlOk returns a tuple with the TopLargeMemSql field value
// and a boolean to check if the value has been set.
func (o *DamengSQLAnalysis) GetTopLargeMemSqlOk() (*[]DamengLargeMemSQLEntry, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TopLargeMemSql, true
}

// SetTopLargeMemSql sets field value.
func (o *DamengSQLAnalysis) SetTopLargeMemSql(v []DamengLargeMemSQLEntry) {
	o.TopLargeMemSql = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DamengSQLAnalysis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["topLongExecSql"] = o.TopLongExecSql
	toSerialize["topLargeMemSql"] = o.TopLargeMemSql

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DamengSQLAnalysis) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TopLongExecSql *[]DamengLongExecSQLEntry `json:"topLongExecSql"`
		TopLargeMemSql *[]DamengLargeMemSQLEntry `json:"topLargeMemSql"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.TopLongExecSql == nil {
		return fmt.Errorf("required field topLongExecSql missing")
	}
	if all.TopLargeMemSql == nil {
		return fmt.Errorf("required field topLargeMemSql missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"topLongExecSql", "topLargeMemSql"})
	} else {
		return err
	}
	o.TopLongExecSql = *all.TopLongExecSql
	o.TopLargeMemSql = *all.TopLargeMemSql

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
