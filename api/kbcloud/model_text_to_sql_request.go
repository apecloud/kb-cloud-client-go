// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type TextToSQLRequest struct {
	// the sql query in natural language
	Query *string `json:"query,omitempty"`
	// the database name
	Database *string  `json:"database,omitempty"`
	Tables   []string `json:"tables,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTextToSQLRequest instantiates a new TextToSQLRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTextToSQLRequest() *TextToSQLRequest {
	this := TextToSQLRequest{}
	return &this
}

// NewTextToSQLRequestWithDefaults instantiates a new TextToSQLRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTextToSQLRequestWithDefaults() *TextToSQLRequest {
	this := TextToSQLRequest{}
	return &this
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *TextToSQLRequest) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextToSQLRequest) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *TextToSQLRequest) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *TextToSQLRequest) SetQuery(v string) {
	o.Query = &v
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *TextToSQLRequest) GetDatabase() string {
	if o == nil || o.Database == nil {
		var ret string
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextToSQLRequest) GetDatabaseOk() (*string, bool) {
	if o == nil || o.Database == nil {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *TextToSQLRequest) HasDatabase() bool {
	return o != nil && o.Database != nil
}

// SetDatabase gets a reference to the given string and assigns it to the Database field.
func (o *TextToSQLRequest) SetDatabase(v string) {
	o.Database = &v
}

// GetTables returns the Tables field value if set, zero value otherwise.
func (o *TextToSQLRequest) GetTables() []string {
	if o == nil || o.Tables == nil {
		var ret []string
		return ret
	}
	return o.Tables
}

// GetTablesOk returns a tuple with the Tables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextToSQLRequest) GetTablesOk() (*[]string, bool) {
	if o == nil || o.Tables == nil {
		return nil, false
	}
	return &o.Tables, true
}

// HasTables returns a boolean if a field has been set.
func (o *TextToSQLRequest) HasTables() bool {
	return o != nil && o.Tables != nil
}

// SetTables gets a reference to the given []string and assigns it to the Tables field.
func (o *TextToSQLRequest) SetTables(v []string) {
	o.Tables = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TextToSQLRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.Database != nil {
		toSerialize["database"] = o.Database
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
func (o *TextToSQLRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Query    *string  `json:"query,omitempty"`
		Database *string  `json:"database,omitempty"`
		Tables   []string `json:"tables,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"query", "database", "tables"})
	} else {
		return err
	}
	o.Query = all.Query
	o.Database = all.Database
	o.Tables = all.Tables

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
