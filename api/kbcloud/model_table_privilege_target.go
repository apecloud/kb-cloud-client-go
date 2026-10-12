// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type TablePrivilegeTarget struct {
	// Exact database name; no wildcard expansion.
	DatabaseName string `json:"databaseName"`
	// Required for PostgreSQL; omitted for MySQL.
	SchemaName *string `json:"schemaName,omitempty"`
	// Exact existing table or view name.
	TableName string `json:"tableName"`
	// MySQL account host, defaults to %. Omitted for PostgreSQL.
	Host *string `json:"host,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTablePrivilegeTarget instantiates a new TablePrivilegeTarget object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTablePrivilegeTarget(databaseName string, tableName string) *TablePrivilegeTarget {
	this := TablePrivilegeTarget{}
	this.DatabaseName = databaseName
	this.TableName = tableName
	return &this
}

// NewTablePrivilegeTargetWithDefaults instantiates a new TablePrivilegeTarget object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTablePrivilegeTargetWithDefaults() *TablePrivilegeTarget {
	this := TablePrivilegeTarget{}
	return &this
}

// GetDatabaseName returns the DatabaseName field value.
func (o *TablePrivilegeTarget) GetDatabaseName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DatabaseName
}

// GetDatabaseNameOk returns a tuple with the DatabaseName field value
// and a boolean to check if the value has been set.
func (o *TablePrivilegeTarget) GetDatabaseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatabaseName, true
}

// SetDatabaseName sets field value.
func (o *TablePrivilegeTarget) SetDatabaseName(v string) {
	o.DatabaseName = v
}

// GetSchemaName returns the SchemaName field value if set, zero value otherwise.
func (o *TablePrivilegeTarget) GetSchemaName() string {
	if o == nil || o.SchemaName == nil {
		var ret string
		return ret
	}
	return *o.SchemaName
}

// GetSchemaNameOk returns a tuple with the SchemaName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TablePrivilegeTarget) GetSchemaNameOk() (*string, bool) {
	if o == nil || o.SchemaName == nil {
		return nil, false
	}
	return o.SchemaName, true
}

// HasSchemaName returns a boolean if a field has been set.
func (o *TablePrivilegeTarget) HasSchemaName() bool {
	return o != nil && o.SchemaName != nil
}

// SetSchemaName gets a reference to the given string and assigns it to the SchemaName field.
func (o *TablePrivilegeTarget) SetSchemaName(v string) {
	o.SchemaName = &v
}

// GetTableName returns the TableName field value.
func (o *TablePrivilegeTarget) GetTableName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TableName
}

// GetTableNameOk returns a tuple with the TableName field value
// and a boolean to check if the value has been set.
func (o *TablePrivilegeTarget) GetTableNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TableName, true
}

// SetTableName sets field value.
func (o *TablePrivilegeTarget) SetTableName(v string) {
	o.TableName = v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *TablePrivilegeTarget) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TablePrivilegeTarget) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *TablePrivilegeTarget) HasHost() bool {
	return o != nil && o.Host != nil
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *TablePrivilegeTarget) SetHost(v string) {
	o.Host = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TablePrivilegeTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["databaseName"] = o.DatabaseName
	if o.SchemaName != nil {
		toSerialize["schemaName"] = o.SchemaName
	}
	toSerialize["tableName"] = o.TableName
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TablePrivilegeTarget) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DatabaseName *string `json:"databaseName"`
		SchemaName   *string `json:"schemaName,omitempty"`
		TableName    *string `json:"tableName"`
		Host         *string `json:"host,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.DatabaseName == nil {
		return fmt.Errorf("required field databaseName missing")
	}
	if all.TableName == nil {
		return fmt.Errorf("required field tableName missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"databaseName", "schemaName", "tableName", "host"})
	} else {
		return err
	}
	o.DatabaseName = *all.DatabaseName
	o.SchemaName = all.SchemaName
	o.TableName = *all.TableName
	o.Host = all.Host

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
