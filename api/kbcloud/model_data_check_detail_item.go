// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DataCheckDetailItem struct {
	// Struct check object key. Present for struct details; omitted for snapshot row details.
	Key *string `json:"key,omitempty"`
	// Source schema or database. Present for snapshot details; omitted for struct details.
	Schema *string `json:"schema,omitempty"`
	// Source table. Present for snapshot details; omitted for struct details.
	Tb *string `json:"tb,omitempty"`
	// Target schema or database for snapshot details when routing changes the target object.
	TargetSchema common.NullableString `json:"target_schema,omitempty"`
	// Target table for snapshot details when routing changes the target object.
	TargetTb      common.NullableString         `json:"target_tb,omitempty"`
	IdColValues   map[string]interface{}        `json:"id_col_values,omitempty"`
	SrcSql        common.NullableString         `json:"src_sql,omitempty"`
	DiffColValues map[string]DataCheckDiffValue `json:"diff_col_values,omitempty"`
	SrcRow        map[string]interface{}        `json:"src_row,omitempty"`
	DstRow        map[string]interface{}        `json:"dst_row,omitempty"`
	DstSql        common.NullableString         `json:"dst_sql,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDataCheckDetailItem instantiates a new DataCheckDetailItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDataCheckDetailItem() *DataCheckDetailItem {
	this := DataCheckDetailItem{}
	return &this
}

// NewDataCheckDetailItemWithDefaults instantiates a new DataCheckDetailItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDataCheckDetailItemWithDefaults() *DataCheckDetailItem {
	this := DataCheckDetailItem{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *DataCheckDetailItem) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCheckDetailItem) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *DataCheckDetailItem) HasKey() bool {
	return o != nil && o.Key != nil
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *DataCheckDetailItem) SetKey(v string) {
	o.Key = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *DataCheckDetailItem) GetSchema() string {
	if o == nil || o.Schema == nil {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCheckDetailItem) GetSchemaOk() (*string, bool) {
	if o == nil || o.Schema == nil {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *DataCheckDetailItem) HasSchema() bool {
	return o != nil && o.Schema != nil
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *DataCheckDetailItem) SetSchema(v string) {
	o.Schema = &v
}

// GetTb returns the Tb field value if set, zero value otherwise.
func (o *DataCheckDetailItem) GetTb() string {
	if o == nil || o.Tb == nil {
		var ret string
		return ret
	}
	return *o.Tb
}

// GetTbOk returns a tuple with the Tb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCheckDetailItem) GetTbOk() (*string, bool) {
	if o == nil || o.Tb == nil {
		return nil, false
	}
	return o.Tb, true
}

// HasTb returns a boolean if a field has been set.
func (o *DataCheckDetailItem) HasTb() bool {
	return o != nil && o.Tb != nil
}

// SetTb gets a reference to the given string and assigns it to the Tb field.
func (o *DataCheckDetailItem) SetTb(v string) {
	o.Tb = &v
}

// GetTargetSchema returns the TargetSchema field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataCheckDetailItem) GetTargetSchema() string {
	if o == nil || o.TargetSchema.Get() == nil {
		var ret string
		return ret
	}
	return *o.TargetSchema.Get()
}

// GetTargetSchemaOk returns a tuple with the TargetSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DataCheckDetailItem) GetTargetSchemaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetSchema.Get(), o.TargetSchema.IsSet()
}

// HasTargetSchema returns a boolean if a field has been set.
func (o *DataCheckDetailItem) HasTargetSchema() bool {
	return o != nil && o.TargetSchema.IsSet()
}

// SetTargetSchema gets a reference to the given common.NullableString and assigns it to the TargetSchema field.
func (o *DataCheckDetailItem) SetTargetSchema(v string) {
	o.TargetSchema.Set(&v)
}

// SetTargetSchemaNil sets the value for TargetSchema to be an explicit nil.
func (o *DataCheckDetailItem) SetTargetSchemaNil() {
	o.TargetSchema.Set(nil)
}

// UnsetTargetSchema ensures that no value is present for TargetSchema, not even an explicit nil.
func (o *DataCheckDetailItem) UnsetTargetSchema() {
	o.TargetSchema.Unset()
}

// GetTargetTb returns the TargetTb field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataCheckDetailItem) GetTargetTb() string {
	if o == nil || o.TargetTb.Get() == nil {
		var ret string
		return ret
	}
	return *o.TargetTb.Get()
}

// GetTargetTbOk returns a tuple with the TargetTb field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DataCheckDetailItem) GetTargetTbOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetTb.Get(), o.TargetTb.IsSet()
}

// HasTargetTb returns a boolean if a field has been set.
func (o *DataCheckDetailItem) HasTargetTb() bool {
	return o != nil && o.TargetTb.IsSet()
}

// SetTargetTb gets a reference to the given common.NullableString and assigns it to the TargetTb field.
func (o *DataCheckDetailItem) SetTargetTb(v string) {
	o.TargetTb.Set(&v)
}

// SetTargetTbNil sets the value for TargetTb to be an explicit nil.
func (o *DataCheckDetailItem) SetTargetTbNil() {
	o.TargetTb.Set(nil)
}

// UnsetTargetTb ensures that no value is present for TargetTb, not even an explicit nil.
func (o *DataCheckDetailItem) UnsetTargetTb() {
	o.TargetTb.Unset()
}

// GetIdColValues returns the IdColValues field value if set, zero value otherwise.
func (o *DataCheckDetailItem) GetIdColValues() map[string]interface{} {
	if o == nil || o.IdColValues == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.IdColValues
}

// GetIdColValuesOk returns a tuple with the IdColValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCheckDetailItem) GetIdColValuesOk() (*map[string]interface{}, bool) {
	if o == nil || o.IdColValues == nil {
		return nil, false
	}
	return &o.IdColValues, true
}

// HasIdColValues returns a boolean if a field has been set.
func (o *DataCheckDetailItem) HasIdColValues() bool {
	return o != nil && o.IdColValues != nil
}

// SetIdColValues gets a reference to the given map[string]interface{} and assigns it to the IdColValues field.
func (o *DataCheckDetailItem) SetIdColValues(v map[string]interface{}) {
	o.IdColValues = v
}

// GetSrcSql returns the SrcSql field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataCheckDetailItem) GetSrcSql() string {
	if o == nil || o.SrcSql.Get() == nil {
		var ret string
		return ret
	}
	return *o.SrcSql.Get()
}

// GetSrcSqlOk returns a tuple with the SrcSql field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DataCheckDetailItem) GetSrcSqlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SrcSql.Get(), o.SrcSql.IsSet()
}

// HasSrcSql returns a boolean if a field has been set.
func (o *DataCheckDetailItem) HasSrcSql() bool {
	return o != nil && o.SrcSql.IsSet()
}

// SetSrcSql gets a reference to the given common.NullableString and assigns it to the SrcSql field.
func (o *DataCheckDetailItem) SetSrcSql(v string) {
	o.SrcSql.Set(&v)
}

// SetSrcSqlNil sets the value for SrcSql to be an explicit nil.
func (o *DataCheckDetailItem) SetSrcSqlNil() {
	o.SrcSql.Set(nil)
}

// UnsetSrcSql ensures that no value is present for SrcSql, not even an explicit nil.
func (o *DataCheckDetailItem) UnsetSrcSql() {
	o.SrcSql.Unset()
}

// GetDiffColValues returns the DiffColValues field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataCheckDetailItem) GetDiffColValues() map[string]DataCheckDiffValue {
	if o == nil {
		var ret map[string]DataCheckDiffValue
		return ret
	}
	return o.DiffColValues
}

// GetDiffColValuesOk returns a tuple with the DiffColValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DataCheckDetailItem) GetDiffColValuesOk() (*map[string]DataCheckDiffValue, bool) {
	if o == nil || o.DiffColValues == nil {
		return nil, false
	}
	return &o.DiffColValues, true
}

// HasDiffColValues returns a boolean if a field has been set.
func (o *DataCheckDetailItem) HasDiffColValues() bool {
	return o != nil && o.DiffColValues != nil
}

// SetDiffColValues gets a reference to the given map[string]DataCheckDiffValue and assigns it to the DiffColValues field.
func (o *DataCheckDetailItem) SetDiffColValues(v map[string]DataCheckDiffValue) {
	o.DiffColValues = v
}

// GetSrcRow returns the SrcRow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataCheckDetailItem) GetSrcRow() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.SrcRow
}

// GetSrcRowOk returns a tuple with the SrcRow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DataCheckDetailItem) GetSrcRowOk() (*map[string]interface{}, bool) {
	if o == nil || o.SrcRow == nil {
		return nil, false
	}
	return &o.SrcRow, true
}

// HasSrcRow returns a boolean if a field has been set.
func (o *DataCheckDetailItem) HasSrcRow() bool {
	return o != nil && o.SrcRow != nil
}

// SetSrcRow gets a reference to the given map[string]interface{} and assigns it to the SrcRow field.
func (o *DataCheckDetailItem) SetSrcRow(v map[string]interface{}) {
	o.SrcRow = v
}

// GetDstRow returns the DstRow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataCheckDetailItem) GetDstRow() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.DstRow
}

// GetDstRowOk returns a tuple with the DstRow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DataCheckDetailItem) GetDstRowOk() (*map[string]interface{}, bool) {
	if o == nil || o.DstRow == nil {
		return nil, false
	}
	return &o.DstRow, true
}

// HasDstRow returns a boolean if a field has been set.
func (o *DataCheckDetailItem) HasDstRow() bool {
	return o != nil && o.DstRow != nil
}

// SetDstRow gets a reference to the given map[string]interface{} and assigns it to the DstRow field.
func (o *DataCheckDetailItem) SetDstRow(v map[string]interface{}) {
	o.DstRow = v
}

// GetDstSql returns the DstSql field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataCheckDetailItem) GetDstSql() string {
	if o == nil || o.DstSql.Get() == nil {
		var ret string
		return ret
	}
	return *o.DstSql.Get()
}

// GetDstSqlOk returns a tuple with the DstSql field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DataCheckDetailItem) GetDstSqlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DstSql.Get(), o.DstSql.IsSet()
}

// HasDstSql returns a boolean if a field has been set.
func (o *DataCheckDetailItem) HasDstSql() bool {
	return o != nil && o.DstSql.IsSet()
}

// SetDstSql gets a reference to the given common.NullableString and assigns it to the DstSql field.
func (o *DataCheckDetailItem) SetDstSql(v string) {
	o.DstSql.Set(&v)
}

// SetDstSqlNil sets the value for DstSql to be an explicit nil.
func (o *DataCheckDetailItem) SetDstSqlNil() {
	o.DstSql.Set(nil)
}

// UnsetDstSql ensures that no value is present for DstSql, not even an explicit nil.
func (o *DataCheckDetailItem) UnsetDstSql() {
	o.DstSql.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o DataCheckDetailItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Schema != nil {
		toSerialize["schema"] = o.Schema
	}
	if o.Tb != nil {
		toSerialize["tb"] = o.Tb
	}
	if o.TargetSchema.IsSet() {
		toSerialize["target_schema"] = o.TargetSchema.Get()
	}
	if o.TargetTb.IsSet() {
		toSerialize["target_tb"] = o.TargetTb.Get()
	}
	if o.IdColValues != nil {
		toSerialize["id_col_values"] = o.IdColValues
	}
	if o.SrcSql.IsSet() {
		toSerialize["src_sql"] = o.SrcSql.Get()
	}
	if o.DiffColValues != nil {
		toSerialize["diff_col_values"] = o.DiffColValues
	}
	if o.SrcRow != nil {
		toSerialize["src_row"] = o.SrcRow
	}
	if o.DstRow != nil {
		toSerialize["dst_row"] = o.DstRow
	}
	if o.DstSql.IsSet() {
		toSerialize["dst_sql"] = o.DstSql.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DataCheckDetailItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Key           *string                       `json:"key,omitempty"`
		Schema        *string                       `json:"schema,omitempty"`
		Tb            *string                       `json:"tb,omitempty"`
		TargetSchema  common.NullableString         `json:"target_schema,omitempty"`
		TargetTb      common.NullableString         `json:"target_tb,omitempty"`
		IdColValues   map[string]interface{}        `json:"id_col_values,omitempty"`
		SrcSql        common.NullableString         `json:"src_sql,omitempty"`
		DiffColValues map[string]DataCheckDiffValue `json:"diff_col_values,omitempty"`
		SrcRow        map[string]interface{}        `json:"src_row,omitempty"`
		DstRow        map[string]interface{}        `json:"dst_row,omitempty"`
		DstSql        common.NullableString         `json:"dst_sql,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"key", "schema", "tb", "target_schema", "target_tb", "id_col_values", "src_sql", "diff_col_values", "src_row", "dst_row", "dst_sql"})
	} else {
		return err
	}
	o.Key = all.Key
	o.Schema = all.Schema
	o.Tb = all.Tb
	o.TargetSchema = all.TargetSchema
	o.TargetTb = all.TargetTb
	o.IdColValues = all.IdColValues
	o.SrcSql = all.SrcSql
	o.DiffColValues = all.DiffColValues
	o.SrcRow = all.SrcRow
	o.DstRow = all.DstRow
	o.DstSql = all.DstSql

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
