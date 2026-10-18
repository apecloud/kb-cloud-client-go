// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ObjectPrivilegeResetStep One ordered, engine-generated step of an account privilege reset.
type ObjectPrivilegeResetStep struct {
	// Connection database in which this statement must run. Omitted when no particular database is required, as with MySQL account-wide REVOKE. This is execution context, not the scope of objects affected by the SQL.
	DatabaseName *string `json:"databaseName,omitempty"`
	// Engine-generated SQL for display. Callers cannot submit or modify these steps for execution; reset rebuilds the plan from current grants.
	Sql string `json:"sql"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObjectPrivilegeResetStep instantiates a new ObjectPrivilegeResetStep object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObjectPrivilegeResetStep(sql string) *ObjectPrivilegeResetStep {
	this := ObjectPrivilegeResetStep{}
	this.Sql = sql
	return &this
}

// NewObjectPrivilegeResetStepWithDefaults instantiates a new ObjectPrivilegeResetStep object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObjectPrivilegeResetStepWithDefaults() *ObjectPrivilegeResetStep {
	this := ObjectPrivilegeResetStep{}
	return &this
}

// GetDatabaseName returns the DatabaseName field value if set, zero value otherwise.
func (o *ObjectPrivilegeResetStep) GetDatabaseName() string {
	if o == nil || o.DatabaseName == nil {
		var ret string
		return ret
	}
	return *o.DatabaseName
}

// GetDatabaseNameOk returns a tuple with the DatabaseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeResetStep) GetDatabaseNameOk() (*string, bool) {
	if o == nil || o.DatabaseName == nil {
		return nil, false
	}
	return o.DatabaseName, true
}

// HasDatabaseName returns a boolean if a field has been set.
func (o *ObjectPrivilegeResetStep) HasDatabaseName() bool {
	return o != nil && o.DatabaseName != nil
}

// SetDatabaseName gets a reference to the given string and assigns it to the DatabaseName field.
func (o *ObjectPrivilegeResetStep) SetDatabaseName(v string) {
	o.DatabaseName = &v
}

// GetSql returns the Sql field value.
func (o *ObjectPrivilegeResetStep) GetSql() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Sql
}

// GetSqlOk returns a tuple with the Sql field value
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeResetStep) GetSqlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sql, true
}

// SetSql sets field value.
func (o *ObjectPrivilegeResetStep) SetSql(v string) {
	o.Sql = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObjectPrivilegeResetStep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.DatabaseName != nil {
		toSerialize["databaseName"] = o.DatabaseName
	}
	toSerialize["sql"] = o.Sql

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObjectPrivilegeResetStep) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DatabaseName *string `json:"databaseName,omitempty"`
		Sql          *string `json:"sql"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Sql == nil {
		return fmt.Errorf("required field sql missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"databaseName", "sql"})
	} else {
		return err
	}
	o.DatabaseName = all.DatabaseName
	o.Sql = *all.Sql

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
