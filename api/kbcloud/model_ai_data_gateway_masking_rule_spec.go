// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiDataGatewayMaskingRuleSpec struct {
	// Optional datasource scope. Empty means all data sources in the gateway.
	DatasourceIds []string `json:"datasourceIds,omitempty"`
	// Optional database scope. Empty means all databases.
	Databases []string `json:"databases,omitempty"`
	// Optional schema scope. Empty means all schemas.
	Schemas []string `json:"schemas,omitempty"`
	// Optional table scope. Empty means all tables.
	Tables []string `json:"tables,omitempty"`
	// Columns to partially mask with the system default masking algorithm.
	Columns []string `json:"columns"`
	Enabled *bool    `json:"enabled,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiDataGatewayMaskingRuleSpec instantiates a new AiDataGatewayMaskingRuleSpec object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiDataGatewayMaskingRuleSpec(columns []string) *AiDataGatewayMaskingRuleSpec {
	this := AiDataGatewayMaskingRuleSpec{}
	this.Columns = columns
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// NewAiDataGatewayMaskingRuleSpecWithDefaults instantiates a new AiDataGatewayMaskingRuleSpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiDataGatewayMaskingRuleSpecWithDefaults() *AiDataGatewayMaskingRuleSpec {
	this := AiDataGatewayMaskingRuleSpec{}
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// GetDatasourceIds returns the DatasourceIds field value if set, zero value otherwise.
func (o *AiDataGatewayMaskingRuleSpec) GetDatasourceIds() []string {
	if o == nil || o.DatasourceIds == nil {
		var ret []string
		return ret
	}
	return o.DatasourceIds
}

// GetDatasourceIdsOk returns a tuple with the DatasourceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayMaskingRuleSpec) GetDatasourceIdsOk() (*[]string, bool) {
	if o == nil || o.DatasourceIds == nil {
		return nil, false
	}
	return &o.DatasourceIds, true
}

// HasDatasourceIds returns a boolean if a field has been set.
func (o *AiDataGatewayMaskingRuleSpec) HasDatasourceIds() bool {
	return o != nil && o.DatasourceIds != nil
}

// SetDatasourceIds gets a reference to the given []string and assigns it to the DatasourceIds field.
func (o *AiDataGatewayMaskingRuleSpec) SetDatasourceIds(v []string) {
	o.DatasourceIds = v
}

// GetDatabases returns the Databases field value if set, zero value otherwise.
func (o *AiDataGatewayMaskingRuleSpec) GetDatabases() []string {
	if o == nil || o.Databases == nil {
		var ret []string
		return ret
	}
	return o.Databases
}

// GetDatabasesOk returns a tuple with the Databases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayMaskingRuleSpec) GetDatabasesOk() (*[]string, bool) {
	if o == nil || o.Databases == nil {
		return nil, false
	}
	return &o.Databases, true
}

// HasDatabases returns a boolean if a field has been set.
func (o *AiDataGatewayMaskingRuleSpec) HasDatabases() bool {
	return o != nil && o.Databases != nil
}

// SetDatabases gets a reference to the given []string and assigns it to the Databases field.
func (o *AiDataGatewayMaskingRuleSpec) SetDatabases(v []string) {
	o.Databases = v
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *AiDataGatewayMaskingRuleSpec) GetSchemas() []string {
	if o == nil || o.Schemas == nil {
		var ret []string
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayMaskingRuleSpec) GetSchemasOk() (*[]string, bool) {
	if o == nil || o.Schemas == nil {
		return nil, false
	}
	return &o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *AiDataGatewayMaskingRuleSpec) HasSchemas() bool {
	return o != nil && o.Schemas != nil
}

// SetSchemas gets a reference to the given []string and assigns it to the Schemas field.
func (o *AiDataGatewayMaskingRuleSpec) SetSchemas(v []string) {
	o.Schemas = v
}

// GetTables returns the Tables field value if set, zero value otherwise.
func (o *AiDataGatewayMaskingRuleSpec) GetTables() []string {
	if o == nil || o.Tables == nil {
		var ret []string
		return ret
	}
	return o.Tables
}

// GetTablesOk returns a tuple with the Tables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayMaskingRuleSpec) GetTablesOk() (*[]string, bool) {
	if o == nil || o.Tables == nil {
		return nil, false
	}
	return &o.Tables, true
}

// HasTables returns a boolean if a field has been set.
func (o *AiDataGatewayMaskingRuleSpec) HasTables() bool {
	return o != nil && o.Tables != nil
}

// SetTables gets a reference to the given []string and assigns it to the Tables field.
func (o *AiDataGatewayMaskingRuleSpec) SetTables(v []string) {
	o.Tables = v
}

// GetColumns returns the Columns field value.
func (o *AiDataGatewayMaskingRuleSpec) GetColumns() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value
// and a boolean to check if the value has been set.
func (o *AiDataGatewayMaskingRuleSpec) GetColumnsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Columns, true
}

// SetColumns sets field value.
func (o *AiDataGatewayMaskingRuleSpec) SetColumns(v []string) {
	o.Columns = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AiDataGatewayMaskingRuleSpec) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayMaskingRuleSpec) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AiDataGatewayMaskingRuleSpec) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AiDataGatewayMaskingRuleSpec) SetEnabled(v bool) {
	o.Enabled = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiDataGatewayMaskingRuleSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.DatasourceIds != nil {
		toSerialize["datasourceIds"] = o.DatasourceIds
	}
	if o.Databases != nil {
		toSerialize["databases"] = o.Databases
	}
	if o.Schemas != nil {
		toSerialize["schemas"] = o.Schemas
	}
	if o.Tables != nil {
		toSerialize["tables"] = o.Tables
	}
	toSerialize["columns"] = o.Columns
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiDataGatewayMaskingRuleSpec) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DatasourceIds []string  `json:"datasourceIds,omitempty"`
		Databases     []string  `json:"databases,omitempty"`
		Schemas       []string  `json:"schemas,omitempty"`
		Tables        []string  `json:"tables,omitempty"`
		Columns       *[]string `json:"columns"`
		Enabled       *bool     `json:"enabled,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Columns == nil {
		return fmt.Errorf("required field columns missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"datasourceIds", "databases", "schemas", "tables", "columns", "enabled"})
	} else {
		return err
	}
	o.DatasourceIds = all.DatasourceIds
	o.Databases = all.Databases
	o.Schemas = all.Schemas
	o.Tables = all.Tables
	o.Columns = *all.Columns
	o.Enabled = all.Enabled

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
