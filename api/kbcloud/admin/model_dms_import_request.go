// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	_io "io"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DmsImportRequest the data of the import task
type DmsImportRequest struct {
	// whether the first line of the file is column name
	FirstLineColumn *bool `json:"firstLineColumn,omitempty"`
	// the delimiter of fields
	FieldsTerminalBy *string `json:"fieldsTerminalBy,omitempty"`
	// the field names
	FieldNames common.NullableList[string] `json:"fieldNames,omitempty"`
	// whether ignore foreign key when import data or not
	IgnoreForeignKey *bool  `json:"ignoreForeignKey,omitempty"`
	FileType         string `json:"fileType"`
	// the data file, csv or other format
	File *_io.Reader `json:"file,omitempty"`
	// the target table name
	TableName string `json:"tableName"`
	// the database of the specific table, required for the engine which don't need to specify database when create database, such as mysql
	Database string `json:"database"`
	// whether skip the row which has null value when import data
	SkipNull *bool `json:"skipNull,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmsImportRequest instantiates a new DmsImportRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsImportRequest(fileType string, tableName string, database string) *DmsImportRequest {
	this := DmsImportRequest{}
	var firstLineColumn bool = true
	this.FirstLineColumn = &firstLineColumn
	var fieldsTerminalBy string = ","
	this.FieldsTerminalBy = &fieldsTerminalBy
	var ignoreForeignKey bool = false
	this.IgnoreForeignKey = &ignoreForeignKey
	this.FileType = fileType
	this.TableName = tableName
	this.Database = database
	var skipNull bool = true
	this.SkipNull = &skipNull
	return &this
}

// NewDmsImportRequestWithDefaults instantiates a new DmsImportRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmsImportRequestWithDefaults() *DmsImportRequest {
	this := DmsImportRequest{}
	var firstLineColumn bool = true
	this.FirstLineColumn = &firstLineColumn
	var fieldsTerminalBy string = ","
	this.FieldsTerminalBy = &fieldsTerminalBy
	var ignoreForeignKey bool = false
	this.IgnoreForeignKey = &ignoreForeignKey
	var skipNull bool = true
	this.SkipNull = &skipNull
	return &this
}

// GetFirstLineColumn returns the FirstLineColumn field value if set, zero value otherwise.
func (o *DmsImportRequest) GetFirstLineColumn() bool {
	if o == nil || o.FirstLineColumn == nil {
		var ret bool
		return ret
	}
	return *o.FirstLineColumn
}

// GetFirstLineColumnOk returns a tuple with the FirstLineColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsImportRequest) GetFirstLineColumnOk() (*bool, bool) {
	if o == nil || o.FirstLineColumn == nil {
		return nil, false
	}
	return o.FirstLineColumn, true
}

// HasFirstLineColumn returns a boolean if a field has been set.
func (o *DmsImportRequest) HasFirstLineColumn() bool {
	return o != nil && o.FirstLineColumn != nil
}

// SetFirstLineColumn gets a reference to the given bool and assigns it to the FirstLineColumn field.
func (o *DmsImportRequest) SetFirstLineColumn(v bool) {
	o.FirstLineColumn = &v
}

// GetFieldsTerminalBy returns the FieldsTerminalBy field value if set, zero value otherwise.
func (o *DmsImportRequest) GetFieldsTerminalBy() string {
	if o == nil || o.FieldsTerminalBy == nil {
		var ret string
		return ret
	}
	return *o.FieldsTerminalBy
}

// GetFieldsTerminalByOk returns a tuple with the FieldsTerminalBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsImportRequest) GetFieldsTerminalByOk() (*string, bool) {
	if o == nil || o.FieldsTerminalBy == nil {
		return nil, false
	}
	return o.FieldsTerminalBy, true
}

// HasFieldsTerminalBy returns a boolean if a field has been set.
func (o *DmsImportRequest) HasFieldsTerminalBy() bool {
	return o != nil && o.FieldsTerminalBy != nil
}

// SetFieldsTerminalBy gets a reference to the given string and assigns it to the FieldsTerminalBy field.
func (o *DmsImportRequest) SetFieldsTerminalBy(v string) {
	o.FieldsTerminalBy = &v
}

// GetFieldNames returns the FieldNames field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsImportRequest) GetFieldNames() []string {
	if o == nil || o.FieldNames.Get() == nil {
		var ret []string
		return ret
	}
	return *o.FieldNames.Get()
}

// GetFieldNamesOk returns a tuple with the FieldNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsImportRequest) GetFieldNamesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FieldNames.Get(), o.FieldNames.IsSet()
}

// HasFieldNames returns a boolean if a field has been set.
func (o *DmsImportRequest) HasFieldNames() bool {
	return o != nil && o.FieldNames.IsSet()
}

// SetFieldNames gets a reference to the given common.NullableList[string] and assigns it to the FieldNames field.
func (o *DmsImportRequest) SetFieldNames(v []string) {
	o.FieldNames.Set(&v)
}

// SetFieldNamesNil sets the value for FieldNames to be an explicit nil.
func (o *DmsImportRequest) SetFieldNamesNil() {
	o.FieldNames.Set(nil)
}

// UnsetFieldNames ensures that no value is present for FieldNames, not even an explicit nil.
func (o *DmsImportRequest) UnsetFieldNames() {
	o.FieldNames.Unset()
}

// GetIgnoreForeignKey returns the IgnoreForeignKey field value if set, zero value otherwise.
func (o *DmsImportRequest) GetIgnoreForeignKey() bool {
	if o == nil || o.IgnoreForeignKey == nil {
		var ret bool
		return ret
	}
	return *o.IgnoreForeignKey
}

// GetIgnoreForeignKeyOk returns a tuple with the IgnoreForeignKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsImportRequest) GetIgnoreForeignKeyOk() (*bool, bool) {
	if o == nil || o.IgnoreForeignKey == nil {
		return nil, false
	}
	return o.IgnoreForeignKey, true
}

// HasIgnoreForeignKey returns a boolean if a field has been set.
func (o *DmsImportRequest) HasIgnoreForeignKey() bool {
	return o != nil && o.IgnoreForeignKey != nil
}

// SetIgnoreForeignKey gets a reference to the given bool and assigns it to the IgnoreForeignKey field.
func (o *DmsImportRequest) SetIgnoreForeignKey(v bool) {
	o.IgnoreForeignKey = &v
}

// GetFileType returns the FileType field value.
func (o *DmsImportRequest) GetFileType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.FileType
}

// GetFileTypeOk returns a tuple with the FileType field value
// and a boolean to check if the value has been set.
func (o *DmsImportRequest) GetFileTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileType, true
}

// SetFileType sets field value.
func (o *DmsImportRequest) SetFileType(v string) {
	o.FileType = v
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *DmsImportRequest) GetFile() _io.Reader {
	if o == nil || o.File == nil {
		var ret _io.Reader
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsImportRequest) GetFileOk() (*_io.Reader, bool) {
	if o == nil || o.File == nil {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *DmsImportRequest) HasFile() bool {
	return o != nil && o.File != nil
}

// SetFile gets a reference to the given _io.Reader and assigns it to the File field.
func (o *DmsImportRequest) SetFile(v _io.Reader) {
	o.File = &v
}

// GetTableName returns the TableName field value.
func (o *DmsImportRequest) GetTableName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TableName
}

// GetTableNameOk returns a tuple with the TableName field value
// and a boolean to check if the value has been set.
func (o *DmsImportRequest) GetTableNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TableName, true
}

// SetTableName sets field value.
func (o *DmsImportRequest) SetTableName(v string) {
	o.TableName = v
}

// GetDatabase returns the Database field value.
func (o *DmsImportRequest) GetDatabase() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Database
}

// GetDatabaseOk returns a tuple with the Database field value
// and a boolean to check if the value has been set.
func (o *DmsImportRequest) GetDatabaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Database, true
}

// SetDatabase sets field value.
func (o *DmsImportRequest) SetDatabase(v string) {
	o.Database = v
}

// GetSkipNull returns the SkipNull field value if set, zero value otherwise.
func (o *DmsImportRequest) GetSkipNull() bool {
	if o == nil || o.SkipNull == nil {
		var ret bool
		return ret
	}
	return *o.SkipNull
}

// GetSkipNullOk returns a tuple with the SkipNull field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsImportRequest) GetSkipNullOk() (*bool, bool) {
	if o == nil || o.SkipNull == nil {
		return nil, false
	}
	return o.SkipNull, true
}

// HasSkipNull returns a boolean if a field has been set.
func (o *DmsImportRequest) HasSkipNull() bool {
	return o != nil && o.SkipNull != nil
}

// SetSkipNull gets a reference to the given bool and assigns it to the SkipNull field.
func (o *DmsImportRequest) SetSkipNull(v bool) {
	o.SkipNull = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DmsImportRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.FirstLineColumn != nil {
		toSerialize["firstLineColumn"] = o.FirstLineColumn
	}
	if o.FieldsTerminalBy != nil {
		toSerialize["fieldsTerminalBy"] = o.FieldsTerminalBy
	}
	if o.FieldNames.IsSet() {
		toSerialize["fieldNames"] = o.FieldNames.Get()
	}
	if o.IgnoreForeignKey != nil {
		toSerialize["ignoreForeignKey"] = o.IgnoreForeignKey
	}
	toSerialize["fileType"] = o.FileType
	if o.File != nil {
		toSerialize["file"] = o.File
	}
	toSerialize["tableName"] = o.TableName
	toSerialize["database"] = o.Database
	if o.SkipNull != nil {
		toSerialize["skipNull"] = o.SkipNull
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmsImportRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FirstLineColumn  *bool                       `json:"firstLineColumn,omitempty"`
		FieldsTerminalBy *string                     `json:"fieldsTerminalBy,omitempty"`
		FieldNames       common.NullableList[string] `json:"fieldNames,omitempty"`
		IgnoreForeignKey *bool                       `json:"ignoreForeignKey,omitempty"`
		FileType         *string                     `json:"fileType"`
		File             *_io.Reader                 `json:"file,omitempty"`
		TableName        *string                     `json:"tableName"`
		Database         *string                     `json:"database"`
		SkipNull         *bool                       `json:"skipNull,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.FileType == nil {
		return fmt.Errorf("required field fileType missing")
	}
	if all.TableName == nil {
		return fmt.Errorf("required field tableName missing")
	}
	if all.Database == nil {
		return fmt.Errorf("required field database missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"firstLineColumn", "fieldsTerminalBy", "fieldNames", "ignoreForeignKey", "fileType", "file", "tableName", "database", "skipNull"})
	} else {
		return err
	}
	o.FirstLineColumn = all.FirstLineColumn
	o.FieldsTerminalBy = all.FieldsTerminalBy
	o.FieldNames = all.FieldNames
	o.IgnoreForeignKey = all.IgnoreForeignKey
	o.FileType = *all.FileType
	o.File = all.File
	o.TableName = *all.TableName
	o.Database = *all.Database
	o.SkipNull = all.SkipNull

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
