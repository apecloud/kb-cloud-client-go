// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ChatRequest Chat message request
type ChatRequest struct {
	// Organization name
	OrgName string `json:"orgName"`
	// Cluster name
	ClusterName string `json:"clusterName"`
	// Session ID
	SessionId string `json:"sessionId"`
	// Message ID
	MessageId *string `json:"messageID,omitempty"`
	// Query content
	Query string `json:"query"`
	// LLM model
	Model string `json:"model"`
	// Datasource name
	DatasourceName *string `json:"datasourceName,omitempty"`
	// Database involved in the chat
	Database *string `json:"database,omitempty"`
	// Schema involved in the chat
	Schema *string `json:"schema,omitempty"`
	// Tables involved in the chat
	Tables []string `json:"tables,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewChatRequest instantiates a new ChatRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewChatRequest(orgName string, clusterName string, sessionId string, query string, model string) *ChatRequest {
	this := ChatRequest{}
	this.OrgName = orgName
	this.ClusterName = clusterName
	this.SessionId = sessionId
	this.Query = query
	this.Model = model
	return &this
}

// NewChatRequestWithDefaults instantiates a new ChatRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewChatRequestWithDefaults() *ChatRequest {
	this := ChatRequest{}
	return &this
}

// GetOrgName returns the OrgName field value.
func (o *ChatRequest) GetOrgName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value
// and a boolean to check if the value has been set.
func (o *ChatRequest) GetOrgNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgName, true
}

// SetOrgName sets field value.
func (o *ChatRequest) SetOrgName(v string) {
	o.OrgName = v
}

// GetClusterName returns the ClusterName field value.
func (o *ChatRequest) GetClusterName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value
// and a boolean to check if the value has been set.
func (o *ChatRequest) GetClusterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterName, true
}

// SetClusterName sets field value.
func (o *ChatRequest) SetClusterName(v string) {
	o.ClusterName = v
}

// GetSessionId returns the SessionId field value.
func (o *ChatRequest) GetSessionId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value
// and a boolean to check if the value has been set.
func (o *ChatRequest) GetSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionId, true
}

// SetSessionId sets field value.
func (o *ChatRequest) SetSessionId(v string) {
	o.SessionId = v
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *ChatRequest) GetMessageId() string {
	if o == nil || o.MessageId == nil {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatRequest) GetMessageIdOk() (*string, bool) {
	if o == nil || o.MessageId == nil {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *ChatRequest) HasMessageId() bool {
	return o != nil && o.MessageId != nil
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *ChatRequest) SetMessageId(v string) {
	o.MessageId = &v
}

// GetQuery returns the Query field value.
func (o *ChatRequest) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *ChatRequest) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *ChatRequest) SetQuery(v string) {
	o.Query = v
}

// GetModel returns the Model field value.
func (o *ChatRequest) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *ChatRequest) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value.
func (o *ChatRequest) SetModel(v string) {
	o.Model = v
}

// GetDatasourceName returns the DatasourceName field value if set, zero value otherwise.
func (o *ChatRequest) GetDatasourceName() string {
	if o == nil || o.DatasourceName == nil {
		var ret string
		return ret
	}
	return *o.DatasourceName
}

// GetDatasourceNameOk returns a tuple with the DatasourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatRequest) GetDatasourceNameOk() (*string, bool) {
	if o == nil || o.DatasourceName == nil {
		return nil, false
	}
	return o.DatasourceName, true
}

// HasDatasourceName returns a boolean if a field has been set.
func (o *ChatRequest) HasDatasourceName() bool {
	return o != nil && o.DatasourceName != nil
}

// SetDatasourceName gets a reference to the given string and assigns it to the DatasourceName field.
func (o *ChatRequest) SetDatasourceName(v string) {
	o.DatasourceName = &v
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *ChatRequest) GetDatabase() string {
	if o == nil || o.Database == nil {
		var ret string
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatRequest) GetDatabaseOk() (*string, bool) {
	if o == nil || o.Database == nil {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *ChatRequest) HasDatabase() bool {
	return o != nil && o.Database != nil
}

// SetDatabase gets a reference to the given string and assigns it to the Database field.
func (o *ChatRequest) SetDatabase(v string) {
	o.Database = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *ChatRequest) GetSchema() string {
	if o == nil || o.Schema == nil {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatRequest) GetSchemaOk() (*string, bool) {
	if o == nil || o.Schema == nil {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *ChatRequest) HasSchema() bool {
	return o != nil && o.Schema != nil
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *ChatRequest) SetSchema(v string) {
	o.Schema = &v
}

// GetTables returns the Tables field value if set, zero value otherwise.
func (o *ChatRequest) GetTables() []string {
	if o == nil || o.Tables == nil {
		var ret []string
		return ret
	}
	return o.Tables
}

// GetTablesOk returns a tuple with the Tables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatRequest) GetTablesOk() (*[]string, bool) {
	if o == nil || o.Tables == nil {
		return nil, false
	}
	return &o.Tables, true
}

// HasTables returns a boolean if a field has been set.
func (o *ChatRequest) HasTables() bool {
	return o != nil && o.Tables != nil
}

// SetTables gets a reference to the given []string and assigns it to the Tables field.
func (o *ChatRequest) SetTables(v []string) {
	o.Tables = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ChatRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["orgName"] = o.OrgName
	toSerialize["clusterName"] = o.ClusterName
	toSerialize["sessionId"] = o.SessionId
	if o.MessageId != nil {
		toSerialize["messageID"] = o.MessageId
	}
	toSerialize["query"] = o.Query
	toSerialize["model"] = o.Model
	if o.DatasourceName != nil {
		toSerialize["datasourceName"] = o.DatasourceName
	}
	if o.Database != nil {
		toSerialize["database"] = o.Database
	}
	if o.Schema != nil {
		toSerialize["schema"] = o.Schema
	}
	if o.Tables != nil {
		toSerialize["tables"] = o.Tables
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ChatRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OrgName        *string  `json:"orgName"`
		ClusterName    *string  `json:"clusterName"`
		SessionId      *string  `json:"sessionId"`
		MessageId      *string  `json:"messageID,omitempty"`
		Query          *string  `json:"query"`
		Model          *string  `json:"model"`
		DatasourceName *string  `json:"datasourceName,omitempty"`
		Database       *string  `json:"database,omitempty"`
		Schema         *string  `json:"schema,omitempty"`
		Tables         []string `json:"tables,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.OrgName == nil {
		return fmt.Errorf("required field orgName missing")
	}
	if all.ClusterName == nil {
		return fmt.Errorf("required field clusterName missing")
	}
	if all.SessionId == nil {
		return fmt.Errorf("required field sessionId missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	if all.Model == nil {
		return fmt.Errorf("required field model missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"orgName", "clusterName", "sessionId", "messageID", "query", "model", "datasourceName", "database", "schema", "tables"})
	} else {
		return err
	}
	o.OrgName = *all.OrgName
	o.ClusterName = *all.ClusterName
	o.SessionId = *all.SessionId
	o.MessageId = all.MessageId
	o.Query = *all.Query
	o.Model = *all.Model
	o.DatasourceName = all.DatasourceName
	o.Database = all.Database
	o.Schema = all.Schema
	o.Tables = all.Tables

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
