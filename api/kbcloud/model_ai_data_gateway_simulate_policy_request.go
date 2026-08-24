// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiDataGatewaySimulatePolicyRequest struct {
	DatasourceId *string  `json:"datasourceId,omitempty"`
	Database     *string  `json:"database,omitempty"`
	Schema       *string  `json:"schema,omitempty"`
	Table        *string  `json:"table,omitempty"`
	Columns      []string `json:"columns,omitempty"`
	SqlType      string   `json:"sqlType"`
	Sql          *string  `json:"sql,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiDataGatewaySimulatePolicyRequest instantiates a new AiDataGatewaySimulatePolicyRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiDataGatewaySimulatePolicyRequest(sqlType string) *AiDataGatewaySimulatePolicyRequest {
	this := AiDataGatewaySimulatePolicyRequest{}
	this.SqlType = sqlType
	return &this
}

// NewAiDataGatewaySimulatePolicyRequestWithDefaults instantiates a new AiDataGatewaySimulatePolicyRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiDataGatewaySimulatePolicyRequestWithDefaults() *AiDataGatewaySimulatePolicyRequest {
	this := AiDataGatewaySimulatePolicyRequest{}
	return &this
}

// GetDatasourceId returns the DatasourceId field value if set, zero value otherwise.
func (o *AiDataGatewaySimulatePolicyRequest) GetDatasourceId() string {
	if o == nil || o.DatasourceId == nil {
		var ret string
		return ret
	}
	return *o.DatasourceId
}

// GetDatasourceIdOk returns a tuple with the DatasourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewaySimulatePolicyRequest) GetDatasourceIdOk() (*string, bool) {
	if o == nil || o.DatasourceId == nil {
		return nil, false
	}
	return o.DatasourceId, true
}

// HasDatasourceId returns a boolean if a field has been set.
func (o *AiDataGatewaySimulatePolicyRequest) HasDatasourceId() bool {
	return o != nil && o.DatasourceId != nil
}

// SetDatasourceId gets a reference to the given string and assigns it to the DatasourceId field.
func (o *AiDataGatewaySimulatePolicyRequest) SetDatasourceId(v string) {
	o.DatasourceId = &v
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *AiDataGatewaySimulatePolicyRequest) GetDatabase() string {
	if o == nil || o.Database == nil {
		var ret string
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewaySimulatePolicyRequest) GetDatabaseOk() (*string, bool) {
	if o == nil || o.Database == nil {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *AiDataGatewaySimulatePolicyRequest) HasDatabase() bool {
	return o != nil && o.Database != nil
}

// SetDatabase gets a reference to the given string and assigns it to the Database field.
func (o *AiDataGatewaySimulatePolicyRequest) SetDatabase(v string) {
	o.Database = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *AiDataGatewaySimulatePolicyRequest) GetSchema() string {
	if o == nil || o.Schema == nil {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewaySimulatePolicyRequest) GetSchemaOk() (*string, bool) {
	if o == nil || o.Schema == nil {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *AiDataGatewaySimulatePolicyRequest) HasSchema() bool {
	return o != nil && o.Schema != nil
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *AiDataGatewaySimulatePolicyRequest) SetSchema(v string) {
	o.Schema = &v
}

// GetTable returns the Table field value if set, zero value otherwise.
func (o *AiDataGatewaySimulatePolicyRequest) GetTable() string {
	if o == nil || o.Table == nil {
		var ret string
		return ret
	}
	return *o.Table
}

// GetTableOk returns a tuple with the Table field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewaySimulatePolicyRequest) GetTableOk() (*string, bool) {
	if o == nil || o.Table == nil {
		return nil, false
	}
	return o.Table, true
}

// HasTable returns a boolean if a field has been set.
func (o *AiDataGatewaySimulatePolicyRequest) HasTable() bool {
	return o != nil && o.Table != nil
}

// SetTable gets a reference to the given string and assigns it to the Table field.
func (o *AiDataGatewaySimulatePolicyRequest) SetTable(v string) {
	o.Table = &v
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *AiDataGatewaySimulatePolicyRequest) GetColumns() []string {
	if o == nil || o.Columns == nil {
		var ret []string
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewaySimulatePolicyRequest) GetColumnsOk() (*[]string, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return &o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *AiDataGatewaySimulatePolicyRequest) HasColumns() bool {
	return o != nil && o.Columns != nil
}

// SetColumns gets a reference to the given []string and assigns it to the Columns field.
func (o *AiDataGatewaySimulatePolicyRequest) SetColumns(v []string) {
	o.Columns = v
}

// GetSqlType returns the SqlType field value.
func (o *AiDataGatewaySimulatePolicyRequest) GetSqlType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SqlType
}

// GetSqlTypeOk returns a tuple with the SqlType field value
// and a boolean to check if the value has been set.
func (o *AiDataGatewaySimulatePolicyRequest) GetSqlTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SqlType, true
}

// SetSqlType sets field value.
func (o *AiDataGatewaySimulatePolicyRequest) SetSqlType(v string) {
	o.SqlType = v
}

// GetSql returns the Sql field value if set, zero value otherwise.
func (o *AiDataGatewaySimulatePolicyRequest) GetSql() string {
	if o == nil || o.Sql == nil {
		var ret string
		return ret
	}
	return *o.Sql
}

// GetSqlOk returns a tuple with the Sql field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewaySimulatePolicyRequest) GetSqlOk() (*string, bool) {
	if o == nil || o.Sql == nil {
		return nil, false
	}
	return o.Sql, true
}

// HasSql returns a boolean if a field has been set.
func (o *AiDataGatewaySimulatePolicyRequest) HasSql() bool {
	return o != nil && o.Sql != nil
}

// SetSql gets a reference to the given string and assigns it to the Sql field.
func (o *AiDataGatewaySimulatePolicyRequest) SetSql(v string) {
	o.Sql = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiDataGatewaySimulatePolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.DatasourceId != nil {
		toSerialize["datasourceId"] = o.DatasourceId
	}
	if o.Database != nil {
		toSerialize["database"] = o.Database
	}
	if o.Schema != nil {
		toSerialize["schema"] = o.Schema
	}
	if o.Table != nil {
		toSerialize["table"] = o.Table
	}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	toSerialize["sqlType"] = o.SqlType
	if o.Sql != nil {
		toSerialize["sql"] = o.Sql
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiDataGatewaySimulatePolicyRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DatasourceId *string  `json:"datasourceId,omitempty"`
		Database     *string  `json:"database,omitempty"`
		Schema       *string  `json:"schema,omitempty"`
		Table        *string  `json:"table,omitempty"`
		Columns      []string `json:"columns,omitempty"`
		SqlType      *string  `json:"sqlType"`
		Sql          *string  `json:"sql,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.SqlType == nil {
		return fmt.Errorf("required field sqlType missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"datasourceId", "database", "schema", "table", "columns", "sqlType", "sql"})
	} else {
		return err
	}
	o.DatasourceId = all.DatasourceId
	o.Database = all.Database
	o.Schema = all.Schema
	o.Table = all.Table
	o.Columns = all.Columns
	o.SqlType = *all.SqlType
	o.Sql = all.Sql

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
