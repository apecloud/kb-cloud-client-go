// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiDataGatewayMaskingRule struct {
	// Optional datasource scope. Empty means all data sources in the gateway.
	DatasourceIds []string `json:"datasourceIds,omitempty"`
	// Optional database scope. Empty means all databases.
	Databases []string `json:"databases,omitempty"`
	// Optional schema scope. Empty means all schemas.
	Schemas []string `json:"schemas,omitempty"`
	// Optional table scope. Empty means all tables.
	Tables []string `json:"tables,omitempty"`
	// Columns to partially mask with the system default masking algorithm.
	Columns   []string   `json:"columns,omitempty"`
	Enabled   *bool      `json:"enabled,omitempty"`
	RuleId    *string    `json:"ruleId,omitempty"`
	GatewayId *string    `json:"gatewayId,omitempty"`
	OrgName   *string    `json:"orgName,omitempty"`
	CreatedBy *string    `json:"createdBy,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiDataGatewayMaskingRule instantiates a new AiDataGatewayMaskingRule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiDataGatewayMaskingRule() *AiDataGatewayMaskingRule {
	this := AiDataGatewayMaskingRule{}
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// GetDatasourceIds returns the DatasourceIds field value if set, zero value otherwise.
func (o *AiDataGatewayMaskingRule) GetDatasourceIds() []string {
	if o == nil || o.DatasourceIds == nil {
		var ret []string
		return ret
	}
	return o.DatasourceIds
}

// GetDatasourceIdsOk returns a tuple with the DatasourceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayMaskingRule) GetDatasourceIdsOk() (*[]string, bool) {
	if o == nil || o.DatasourceIds == nil {
		return nil, false
	}
	return &o.DatasourceIds, true
}

// HasDatasourceIds returns a boolean if a field has been set.
func (o *AiDataGatewayMaskingRule) HasDatasourceIds() bool {
	return o != nil && o.DatasourceIds != nil
}

// SetDatasourceIds gets a reference to the given []string and assigns it to the DatasourceIds field.
func (o *AiDataGatewayMaskingRule) SetDatasourceIds(v []string) {
	o.DatasourceIds = v
}

// GetDatabases returns the Databases field value if set, zero value otherwise.
func (o *AiDataGatewayMaskingRule) GetDatabases() []string {
	if o == nil || o.Databases == nil {
		var ret []string
		return ret
	}
	return o.Databases
}

// GetDatabasesOk returns a tuple with the Databases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayMaskingRule) GetDatabasesOk() (*[]string, bool) {
	if o == nil || o.Databases == nil {
		return nil, false
	}
	return &o.Databases, true
}

// HasDatabases returns a boolean if a field has been set.
func (o *AiDataGatewayMaskingRule) HasDatabases() bool {
	return o != nil && o.Databases != nil
}

// SetDatabases gets a reference to the given []string and assigns it to the Databases field.
func (o *AiDataGatewayMaskingRule) SetDatabases(v []string) {
	o.Databases = v
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *AiDataGatewayMaskingRule) GetSchemas() []string {
	if o == nil || o.Schemas == nil {
		var ret []string
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayMaskingRule) GetSchemasOk() (*[]string, bool) {
	if o == nil || o.Schemas == nil {
		return nil, false
	}
	return &o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *AiDataGatewayMaskingRule) HasSchemas() bool {
	return o != nil && o.Schemas != nil
}

// SetSchemas gets a reference to the given []string and assigns it to the Schemas field.
func (o *AiDataGatewayMaskingRule) SetSchemas(v []string) {
	o.Schemas = v
}

// GetTables returns the Tables field value if set, zero value otherwise.
func (o *AiDataGatewayMaskingRule) GetTables() []string {
	if o == nil || o.Tables == nil {
		var ret []string
		return ret
	}
	return o.Tables
}

// GetTablesOk returns a tuple with the Tables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayMaskingRule) GetTablesOk() (*[]string, bool) {
	if o == nil || o.Tables == nil {
		return nil, false
	}
	return &o.Tables, true
}

// HasTables returns a boolean if a field has been set.
func (o *AiDataGatewayMaskingRule) HasTables() bool {
	return o != nil && o.Tables != nil
}

// SetTables gets a reference to the given []string and assigns it to the Tables field.
func (o *AiDataGatewayMaskingRule) SetTables(v []string) {
	o.Tables = v
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *AiDataGatewayMaskingRule) GetColumns() []string {
	if o == nil || o.Columns == nil {
		var ret []string
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayMaskingRule) GetColumnsOk() (*[]string, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return &o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *AiDataGatewayMaskingRule) HasColumns() bool {
	return o != nil && o.Columns != nil
}

// SetColumns gets a reference to the given []string and assigns it to the Columns field.
func (o *AiDataGatewayMaskingRule) SetColumns(v []string) {
	o.Columns = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AiDataGatewayMaskingRule) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayMaskingRule) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AiDataGatewayMaskingRule) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AiDataGatewayMaskingRule) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetRuleId returns the RuleId field value if set, zero value otherwise.
func (o *AiDataGatewayMaskingRule) GetRuleId() string {
	if o == nil || o.RuleId == nil {
		var ret string
		return ret
	}
	return *o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayMaskingRule) GetRuleIdOk() (*string, bool) {
	if o == nil || o.RuleId == nil {
		return nil, false
	}
	return o.RuleId, true
}

// HasRuleId returns a boolean if a field has been set.
func (o *AiDataGatewayMaskingRule) HasRuleId() bool {
	return o != nil && o.RuleId != nil
}

// SetRuleId gets a reference to the given string and assigns it to the RuleId field.
func (o *AiDataGatewayMaskingRule) SetRuleId(v string) {
	o.RuleId = &v
}

// GetGatewayId returns the GatewayId field value if set, zero value otherwise.
func (o *AiDataGatewayMaskingRule) GetGatewayId() string {
	if o == nil || o.GatewayId == nil {
		var ret string
		return ret
	}
	return *o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayMaskingRule) GetGatewayIdOk() (*string, bool) {
	if o == nil || o.GatewayId == nil {
		return nil, false
	}
	return o.GatewayId, true
}

// HasGatewayId returns a boolean if a field has been set.
func (o *AiDataGatewayMaskingRule) HasGatewayId() bool {
	return o != nil && o.GatewayId != nil
}

// SetGatewayId gets a reference to the given string and assigns it to the GatewayId field.
func (o *AiDataGatewayMaskingRule) SetGatewayId(v string) {
	o.GatewayId = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *AiDataGatewayMaskingRule) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayMaskingRule) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *AiDataGatewayMaskingRule) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *AiDataGatewayMaskingRule) SetOrgName(v string) {
	o.OrgName = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *AiDataGatewayMaskingRule) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayMaskingRule) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *AiDataGatewayMaskingRule) HasCreatedBy() bool {
	return o != nil && o.CreatedBy != nil
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *AiDataGatewayMaskingRule) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AiDataGatewayMaskingRule) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayMaskingRule) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AiDataGatewayMaskingRule) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AiDataGatewayMaskingRule) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AiDataGatewayMaskingRule) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayMaskingRule) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AiDataGatewayMaskingRule) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *AiDataGatewayMaskingRule) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiDataGatewayMaskingRule) MarshalJSON() ([]byte, error) {
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
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.RuleId != nil {
		toSerialize["ruleId"] = o.RuleId
	}
	if o.GatewayId != nil {
		toSerialize["gatewayId"] = o.GatewayId
	}
	if o.OrgName != nil {
		toSerialize["orgName"] = o.OrgName
	}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.UpdatedAt != nil {
		if o.UpdatedAt.Nanosecond() == 0 {
			toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiDataGatewayMaskingRule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DatasourceIds []string   `json:"datasourceIds,omitempty"`
		Databases     []string   `json:"databases,omitempty"`
		Schemas       []string   `json:"schemas,omitempty"`
		Tables        []string   `json:"tables,omitempty"`
		Columns       []string   `json:"columns,omitempty"`
		Enabled       *bool      `json:"enabled,omitempty"`
		RuleId        *string    `json:"ruleId,omitempty"`
		GatewayId     *string    `json:"gatewayId,omitempty"`
		OrgName       *string    `json:"orgName,omitempty"`
		CreatedBy     *string    `json:"createdBy,omitempty"`
		CreatedAt     *time.Time `json:"createdAt,omitempty"`
		UpdatedAt     *time.Time `json:"updatedAt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"datasourceIds", "databases", "schemas", "tables", "columns", "enabled", "ruleId", "gatewayId", "orgName", "createdBy", "createdAt", "updatedAt"})
	} else {
		return err
	}
	o.DatasourceIds = all.DatasourceIds
	o.Databases = all.Databases
	o.Schemas = all.Schemas
	o.Tables = all.Tables
	o.Columns = all.Columns
	o.Enabled = all.Enabled
	o.RuleId = all.RuleId
	o.GatewayId = all.GatewayId
	o.OrgName = all.OrgName
	o.CreatedBy = all.CreatedBy
	o.CreatedAt = all.CreatedAt
	o.UpdatedAt = all.UpdatedAt
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
