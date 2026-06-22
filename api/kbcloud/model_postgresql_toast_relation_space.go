// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlToastRelationSpace struct {
	Database          *string `json:"database,omitempty"`
	Schema            *string `json:"schema,omitempty"`
	Table             *string `json:"table,omitempty"`
	TableRelationOid  *int64  `json:"tableRelationOid,omitempty"`
	ToastRelationOid  *int64  `json:"toastRelationOid,omitempty"`
	ToastBytes        *int64  `json:"toastBytes,omitempty"`
	NoHistoryYet      bool    `json:"noHistoryYet"`
	CannotProveGrowth bool    `json:"cannotProveGrowth"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlToastRelationSpace instantiates a new PostgresqlToastRelationSpace object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlToastRelationSpace(noHistoryYet bool, cannotProveGrowth bool) *PostgresqlToastRelationSpace {
	this := PostgresqlToastRelationSpace{}
	this.NoHistoryYet = noHistoryYet
	this.CannotProveGrowth = cannotProveGrowth
	return &this
}

// NewPostgresqlToastRelationSpaceWithDefaults instantiates a new PostgresqlToastRelationSpace object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlToastRelationSpaceWithDefaults() *PostgresqlToastRelationSpace {
	this := PostgresqlToastRelationSpace{}
	return &this
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *PostgresqlToastRelationSpace) GetDatabase() string {
	if o == nil || o.Database == nil {
		var ret string
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlToastRelationSpace) GetDatabaseOk() (*string, bool) {
	if o == nil || o.Database == nil {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *PostgresqlToastRelationSpace) HasDatabase() bool {
	return o != nil && o.Database != nil
}

// SetDatabase gets a reference to the given string and assigns it to the Database field.
func (o *PostgresqlToastRelationSpace) SetDatabase(v string) {
	o.Database = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *PostgresqlToastRelationSpace) GetSchema() string {
	if o == nil || o.Schema == nil {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlToastRelationSpace) GetSchemaOk() (*string, bool) {
	if o == nil || o.Schema == nil {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *PostgresqlToastRelationSpace) HasSchema() bool {
	return o != nil && o.Schema != nil
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *PostgresqlToastRelationSpace) SetSchema(v string) {
	o.Schema = &v
}

// GetTable returns the Table field value if set, zero value otherwise.
func (o *PostgresqlToastRelationSpace) GetTable() string {
	if o == nil || o.Table == nil {
		var ret string
		return ret
	}
	return *o.Table
}

// GetTableOk returns a tuple with the Table field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlToastRelationSpace) GetTableOk() (*string, bool) {
	if o == nil || o.Table == nil {
		return nil, false
	}
	return o.Table, true
}

// HasTable returns a boolean if a field has been set.
func (o *PostgresqlToastRelationSpace) HasTable() bool {
	return o != nil && o.Table != nil
}

// SetTable gets a reference to the given string and assigns it to the Table field.
func (o *PostgresqlToastRelationSpace) SetTable(v string) {
	o.Table = &v
}

// GetTableRelationOid returns the TableRelationOid field value if set, zero value otherwise.
func (o *PostgresqlToastRelationSpace) GetTableRelationOid() int64 {
	if o == nil || o.TableRelationOid == nil {
		var ret int64
		return ret
	}
	return *o.TableRelationOid
}

// GetTableRelationOidOk returns a tuple with the TableRelationOid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlToastRelationSpace) GetTableRelationOidOk() (*int64, bool) {
	if o == nil || o.TableRelationOid == nil {
		return nil, false
	}
	return o.TableRelationOid, true
}

// HasTableRelationOid returns a boolean if a field has been set.
func (o *PostgresqlToastRelationSpace) HasTableRelationOid() bool {
	return o != nil && o.TableRelationOid != nil
}

// SetTableRelationOid gets a reference to the given int64 and assigns it to the TableRelationOid field.
func (o *PostgresqlToastRelationSpace) SetTableRelationOid(v int64) {
	o.TableRelationOid = &v
}

// GetToastRelationOid returns the ToastRelationOid field value if set, zero value otherwise.
func (o *PostgresqlToastRelationSpace) GetToastRelationOid() int64 {
	if o == nil || o.ToastRelationOid == nil {
		var ret int64
		return ret
	}
	return *o.ToastRelationOid
}

// GetToastRelationOidOk returns a tuple with the ToastRelationOid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlToastRelationSpace) GetToastRelationOidOk() (*int64, bool) {
	if o == nil || o.ToastRelationOid == nil {
		return nil, false
	}
	return o.ToastRelationOid, true
}

// HasToastRelationOid returns a boolean if a field has been set.
func (o *PostgresqlToastRelationSpace) HasToastRelationOid() bool {
	return o != nil && o.ToastRelationOid != nil
}

// SetToastRelationOid gets a reference to the given int64 and assigns it to the ToastRelationOid field.
func (o *PostgresqlToastRelationSpace) SetToastRelationOid(v int64) {
	o.ToastRelationOid = &v
}

// GetToastBytes returns the ToastBytes field value if set, zero value otherwise.
func (o *PostgresqlToastRelationSpace) GetToastBytes() int64 {
	if o == nil || o.ToastBytes == nil {
		var ret int64
		return ret
	}
	return *o.ToastBytes
}

// GetToastBytesOk returns a tuple with the ToastBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlToastRelationSpace) GetToastBytesOk() (*int64, bool) {
	if o == nil || o.ToastBytes == nil {
		return nil, false
	}
	return o.ToastBytes, true
}

// HasToastBytes returns a boolean if a field has been set.
func (o *PostgresqlToastRelationSpace) HasToastBytes() bool {
	return o != nil && o.ToastBytes != nil
}

// SetToastBytes gets a reference to the given int64 and assigns it to the ToastBytes field.
func (o *PostgresqlToastRelationSpace) SetToastBytes(v int64) {
	o.ToastBytes = &v
}

// GetNoHistoryYet returns the NoHistoryYet field value.
func (o *PostgresqlToastRelationSpace) GetNoHistoryYet() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.NoHistoryYet
}

// GetNoHistoryYetOk returns a tuple with the NoHistoryYet field value
// and a boolean to check if the value has been set.
func (o *PostgresqlToastRelationSpace) GetNoHistoryYetOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NoHistoryYet, true
}

// SetNoHistoryYet sets field value.
func (o *PostgresqlToastRelationSpace) SetNoHistoryYet(v bool) {
	o.NoHistoryYet = v
}

// GetCannotProveGrowth returns the CannotProveGrowth field value.
func (o *PostgresqlToastRelationSpace) GetCannotProveGrowth() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.CannotProveGrowth
}

// GetCannotProveGrowthOk returns a tuple with the CannotProveGrowth field value
// and a boolean to check if the value has been set.
func (o *PostgresqlToastRelationSpace) GetCannotProveGrowthOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CannotProveGrowth, true
}

// SetCannotProveGrowth sets field value.
func (o *PostgresqlToastRelationSpace) SetCannotProveGrowth(v bool) {
	o.CannotProveGrowth = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlToastRelationSpace) MarshalJSON() ([]byte, error) {
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
	if o.TableRelationOid != nil {
		toSerialize["tableRelationOid"] = o.TableRelationOid
	}
	if o.ToastRelationOid != nil {
		toSerialize["toastRelationOid"] = o.ToastRelationOid
	}
	if o.ToastBytes != nil {
		toSerialize["toastBytes"] = o.ToastBytes
	}
	toSerialize["noHistoryYet"] = o.NoHistoryYet
	toSerialize["cannotProveGrowth"] = o.CannotProveGrowth

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostgresqlToastRelationSpace) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Database          *string `json:"database,omitempty"`
		Schema            *string `json:"schema,omitempty"`
		Table             *string `json:"table,omitempty"`
		TableRelationOid  *int64  `json:"tableRelationOid,omitempty"`
		ToastRelationOid  *int64  `json:"toastRelationOid,omitempty"`
		ToastBytes        *int64  `json:"toastBytes,omitempty"`
		NoHistoryYet      *bool   `json:"noHistoryYet"`
		CannotProveGrowth *bool   `json:"cannotProveGrowth"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.NoHistoryYet == nil {
		return fmt.Errorf("required field noHistoryYet missing")
	}
	if all.CannotProveGrowth == nil {
		return fmt.Errorf("required field cannotProveGrowth missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"database", "schema", "table", "tableRelationOid", "toastRelationOid", "toastBytes", "noHistoryYet", "cannotProveGrowth"})
	} else {
		return err
	}
	o.Database = all.Database
	o.Schema = all.Schema
	o.Table = all.Table
	o.TableRelationOid = all.TableRelationOid
	o.ToastRelationOid = all.ToastRelationOid
	o.ToastBytes = all.ToastBytes
	o.NoHistoryYet = *all.NoHistoryYet
	o.CannotProveGrowth = *all.CannotProveGrowth

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
