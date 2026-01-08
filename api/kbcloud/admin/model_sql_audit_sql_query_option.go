// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type SqlAuditSQLQueryOption struct {
	// SQL query to check audit log status
	Sql string `json:"sql"`
	// List of query result values that indicate SQL audit is enabled (case-insensitive comparison). If not provided, defaults to ["1", "true", "on", "enabled", "yes"].
	EnabledValues []string `json:"enabledValues,omitempty"`
	// List of query result values that indicate SQL audit is disabled (case-insensitive comparison). If not provided, defaults to ["0", "false", "off", "disabled", "no"].
	DisabledValues []string `json:"disabledValues,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSqlAuditSQLQueryOption instantiates a new SqlAuditSQLQueryOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSqlAuditSQLQueryOption(sql string) *SqlAuditSQLQueryOption {
	this := SqlAuditSQLQueryOption{}
	this.Sql = sql
	return &this
}

// NewSqlAuditSQLQueryOptionWithDefaults instantiates a new SqlAuditSQLQueryOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSqlAuditSQLQueryOptionWithDefaults() *SqlAuditSQLQueryOption {
	this := SqlAuditSQLQueryOption{}
	return &this
}

// GetSql returns the Sql field value.
func (o *SqlAuditSQLQueryOption) GetSql() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Sql
}

// GetSqlOk returns a tuple with the Sql field value
// and a boolean to check if the value has been set.
func (o *SqlAuditSQLQueryOption) GetSqlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sql, true
}

// SetSql sets field value.
func (o *SqlAuditSQLQueryOption) SetSql(v string) {
	o.Sql = v
}

// GetEnabledValues returns the EnabledValues field value if set, zero value otherwise.
func (o *SqlAuditSQLQueryOption) GetEnabledValues() []string {
	if o == nil || o.EnabledValues == nil {
		var ret []string
		return ret
	}
	return o.EnabledValues
}

// GetEnabledValuesOk returns a tuple with the EnabledValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SqlAuditSQLQueryOption) GetEnabledValuesOk() (*[]string, bool) {
	if o == nil || o.EnabledValues == nil {
		return nil, false
	}
	return &o.EnabledValues, true
}

// HasEnabledValues returns a boolean if a field has been set.
func (o *SqlAuditSQLQueryOption) HasEnabledValues() bool {
	return o != nil && o.EnabledValues != nil
}

// SetEnabledValues gets a reference to the given []string and assigns it to the EnabledValues field.
func (o *SqlAuditSQLQueryOption) SetEnabledValues(v []string) {
	o.EnabledValues = v
}

// GetDisabledValues returns the DisabledValues field value if set, zero value otherwise.
func (o *SqlAuditSQLQueryOption) GetDisabledValues() []string {
	if o == nil || o.DisabledValues == nil {
		var ret []string
		return ret
	}
	return o.DisabledValues
}

// GetDisabledValuesOk returns a tuple with the DisabledValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SqlAuditSQLQueryOption) GetDisabledValuesOk() (*[]string, bool) {
	if o == nil || o.DisabledValues == nil {
		return nil, false
	}
	return &o.DisabledValues, true
}

// HasDisabledValues returns a boolean if a field has been set.
func (o *SqlAuditSQLQueryOption) HasDisabledValues() bool {
	return o != nil && o.DisabledValues != nil
}

// SetDisabledValues gets a reference to the given []string and assigns it to the DisabledValues field.
func (o *SqlAuditSQLQueryOption) SetDisabledValues(v []string) {
	o.DisabledValues = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SqlAuditSQLQueryOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["sql"] = o.Sql
	if o.EnabledValues != nil {
		toSerialize["enabledValues"] = o.EnabledValues
	}
	if o.DisabledValues != nil {
		toSerialize["disabledValues"] = o.DisabledValues
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SqlAuditSQLQueryOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Sql            *string  `json:"sql"`
		EnabledValues  []string `json:"enabledValues,omitempty"`
		DisabledValues []string `json:"disabledValues,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Sql == nil {
		return fmt.Errorf("required field sql missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"sql", "enabledValues", "disabledValues"})
	} else {
		return err
	}
	o.Sql = *all.Sql
	o.EnabledValues = all.EnabledValues
	o.DisabledValues = all.DisabledValues

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
