// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// DmsForeignKeyReference The reference details of the foreign key
type DmsForeignKeyReference struct {
	// The referenced database
	Database *string `json:"database,omitempty"`
	// The referenced schema
	Schema *string `json:"schema,omitempty"`
	// The referenced table
	Table *string `json:"table,omitempty"`
	// The referenced columns
	Columns []string `json:"columns,omitempty"`
	// The match type for the foreign key (e.g., SIMPLE, FULL)
	MatchType common.NullableString `json:"matchType,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmsForeignKeyReference instantiates a new DmsForeignKeyReference object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsForeignKeyReference() *DmsForeignKeyReference {
	this := DmsForeignKeyReference{}
	return &this
}

// NewDmsForeignKeyReferenceWithDefaults instantiates a new DmsForeignKeyReference object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmsForeignKeyReferenceWithDefaults() *DmsForeignKeyReference {
	this := DmsForeignKeyReference{}
	return &this
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *DmsForeignKeyReference) GetDatabase() string {
	if o == nil || o.Database == nil {
		var ret string
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsForeignKeyReference) GetDatabaseOk() (*string, bool) {
	if o == nil || o.Database == nil {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *DmsForeignKeyReference) HasDatabase() bool {
	return o != nil && o.Database != nil
}

// SetDatabase gets a reference to the given string and assigns it to the Database field.
func (o *DmsForeignKeyReference) SetDatabase(v string) {
	o.Database = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *DmsForeignKeyReference) GetSchema() string {
	if o == nil || o.Schema == nil {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsForeignKeyReference) GetSchemaOk() (*string, bool) {
	if o == nil || o.Schema == nil {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *DmsForeignKeyReference) HasSchema() bool {
	return o != nil && o.Schema != nil
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *DmsForeignKeyReference) SetSchema(v string) {
	o.Schema = &v
}

// GetTable returns the Table field value if set, zero value otherwise.
func (o *DmsForeignKeyReference) GetTable() string {
	if o == nil || o.Table == nil {
		var ret string
		return ret
	}
	return *o.Table
}

// GetTableOk returns a tuple with the Table field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsForeignKeyReference) GetTableOk() (*string, bool) {
	if o == nil || o.Table == nil {
		return nil, false
	}
	return o.Table, true
}

// HasTable returns a boolean if a field has been set.
func (o *DmsForeignKeyReference) HasTable() bool {
	return o != nil && o.Table != nil
}

// SetTable gets a reference to the given string and assigns it to the Table field.
func (o *DmsForeignKeyReference) SetTable(v string) {
	o.Table = &v
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *DmsForeignKeyReference) GetColumns() []string {
	if o == nil || o.Columns == nil {
		var ret []string
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsForeignKeyReference) GetColumnsOk() (*[]string, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return &o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *DmsForeignKeyReference) HasColumns() bool {
	return o != nil && o.Columns != nil
}

// SetColumns gets a reference to the given []string and assigns it to the Columns field.
func (o *DmsForeignKeyReference) SetColumns(v []string) {
	o.Columns = v
}

// GetMatchType returns the MatchType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsForeignKeyReference) GetMatchType() string {
	if o == nil || o.MatchType.Get() == nil {
		var ret string
		return ret
	}
	return *o.MatchType.Get()
}

// GetMatchTypeOk returns a tuple with the MatchType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsForeignKeyReference) GetMatchTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MatchType.Get(), o.MatchType.IsSet()
}

// HasMatchType returns a boolean if a field has been set.
func (o *DmsForeignKeyReference) HasMatchType() bool {
	return o != nil && o.MatchType.IsSet()
}

// SetMatchType gets a reference to the given common.NullableString and assigns it to the MatchType field.
func (o *DmsForeignKeyReference) SetMatchType(v string) {
	o.MatchType.Set(&v)
}

// SetMatchTypeNil sets the value for MatchType to be an explicit nil.
func (o *DmsForeignKeyReference) SetMatchTypeNil() {
	o.MatchType.Set(nil)
}

// UnsetMatchType ensures that no value is present for MatchType, not even an explicit nil.
func (o *DmsForeignKeyReference) UnsetMatchType() {
	o.MatchType.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o DmsForeignKeyReference) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
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
	if o.MatchType.IsSet() {
		toSerialize["matchType"] = o.MatchType.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmsForeignKeyReference) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Database  *string               `json:"database,omitempty"`
		Schema    *string               `json:"schema,omitempty"`
		Table     *string               `json:"table,omitempty"`
		Columns   []string              `json:"columns,omitempty"`
		MatchType common.NullableString `json:"matchType,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"database", "schema", "table", "columns", "matchType"})
	} else {
		return err
	}
	o.Database = all.Database
	o.Schema = all.Schema
	o.Table = all.Table
	o.Columns = all.Columns
	o.MatchType = all.MatchType

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
