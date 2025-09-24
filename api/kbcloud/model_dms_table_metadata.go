// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DmsTableMetadata struct {
	// The name of the table
	Name *string `json:"name,omitempty"`
	// The schema of the table
	Schema common.NullableString `json:"schema,omitempty"`
	// The database where the table resides
	Database common.NullableString `json:"database,omitempty"`
	// A comment describing the table
	Comment common.NullableString `json:"comment,omitempty"`
	// List of columns in the table
	Columns []DmsTableColumn `json:"columns,omitempty"`
	// List of indexes on the table
	Indexes    []DmsTableIndex `json:"indexes,omitempty"`
	PrimaryKey *DmsPrimaryKey  `json:"primaryKey,omitempty"`
	// List of foreign keys on the table
	ForeignKeys []DmsForeignKey `json:"foreignKeys,omitempty"`
	// List of unique keys on the table
	UniqueKeys []DmsUniqueKey `json:"uniqueKeys,omitempty"`
	// List of check constraints on the table
	CheckConstraints []DmsCheckConstraint `json:"checkConstraints,omitempty"`
	// List of exclude constraints on the table
	ExcludeConstraints []DmsExcludeConstraint `json:"excludeConstraints,omitempty"`
	Partitioning       *DmsTablePartitioning  `json:"partitioning,omitempty"`
	Options            *DmsTableOptions       `json:"options,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmsTableMetadata instantiates a new DmsTableMetadata object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsTableMetadata() *DmsTableMetadata {
	this := DmsTableMetadata{}
	return &this
}

// NewDmsTableMetadataWithDefaults instantiates a new DmsTableMetadata object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmsTableMetadataWithDefaults() *DmsTableMetadata {
	this := DmsTableMetadata{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DmsTableMetadata) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsTableMetadata) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DmsTableMetadata) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DmsTableMetadata) SetName(v string) {
	o.Name = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsTableMetadata) GetSchema() string {
	if o == nil || o.Schema.Get() == nil {
		var ret string
		return ret
	}
	return *o.Schema.Get()
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsTableMetadata) GetSchemaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schema.Get(), o.Schema.IsSet()
}

// HasSchema returns a boolean if a field has been set.
func (o *DmsTableMetadata) HasSchema() bool {
	return o != nil && o.Schema.IsSet()
}

// SetSchema gets a reference to the given common.NullableString and assigns it to the Schema field.
func (o *DmsTableMetadata) SetSchema(v string) {
	o.Schema.Set(&v)
}

// SetSchemaNil sets the value for Schema to be an explicit nil.
func (o *DmsTableMetadata) SetSchemaNil() {
	o.Schema.Set(nil)
}

// UnsetSchema ensures that no value is present for Schema, not even an explicit nil.
func (o *DmsTableMetadata) UnsetSchema() {
	o.Schema.Unset()
}

// GetDatabase returns the Database field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsTableMetadata) GetDatabase() string {
	if o == nil || o.Database.Get() == nil {
		var ret string
		return ret
	}
	return *o.Database.Get()
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsTableMetadata) GetDatabaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Database.Get(), o.Database.IsSet()
}

// HasDatabase returns a boolean if a field has been set.
func (o *DmsTableMetadata) HasDatabase() bool {
	return o != nil && o.Database.IsSet()
}

// SetDatabase gets a reference to the given common.NullableString and assigns it to the Database field.
func (o *DmsTableMetadata) SetDatabase(v string) {
	o.Database.Set(&v)
}

// SetDatabaseNil sets the value for Database to be an explicit nil.
func (o *DmsTableMetadata) SetDatabaseNil() {
	o.Database.Set(nil)
}

// UnsetDatabase ensures that no value is present for Database, not even an explicit nil.
func (o *DmsTableMetadata) UnsetDatabase() {
	o.Database.Unset()
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsTableMetadata) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsTableMetadata) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *DmsTableMetadata) HasComment() bool {
	return o != nil && o.Comment.IsSet()
}

// SetComment gets a reference to the given common.NullableString and assigns it to the Comment field.
func (o *DmsTableMetadata) SetComment(v string) {
	o.Comment.Set(&v)
}

// SetCommentNil sets the value for Comment to be an explicit nil.
func (o *DmsTableMetadata) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil.
func (o *DmsTableMetadata) UnsetComment() {
	o.Comment.Unset()
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *DmsTableMetadata) GetColumns() []DmsTableColumn {
	if o == nil || o.Columns == nil {
		var ret []DmsTableColumn
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsTableMetadata) GetColumnsOk() (*[]DmsTableColumn, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return &o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *DmsTableMetadata) HasColumns() bool {
	return o != nil && o.Columns != nil
}

// SetColumns gets a reference to the given []DmsTableColumn and assigns it to the Columns field.
func (o *DmsTableMetadata) SetColumns(v []DmsTableColumn) {
	o.Columns = v
}

// GetIndexes returns the Indexes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsTableMetadata) GetIndexes() []DmsTableIndex {
	if o == nil {
		var ret []DmsTableIndex
		return ret
	}
	return o.Indexes
}

// GetIndexesOk returns a tuple with the Indexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsTableMetadata) GetIndexesOk() (*[]DmsTableIndex, bool) {
	if o == nil || o.Indexes == nil {
		return nil, false
	}
	return &o.Indexes, true
}

// HasIndexes returns a boolean if a field has been set.
func (o *DmsTableMetadata) HasIndexes() bool {
	return o != nil && o.Indexes != nil
}

// SetIndexes gets a reference to the given []DmsTableIndex and assigns it to the Indexes field.
func (o *DmsTableMetadata) SetIndexes(v []DmsTableIndex) {
	o.Indexes = v
}

// GetPrimaryKey returns the PrimaryKey field value if set, zero value otherwise.
func (o *DmsTableMetadata) GetPrimaryKey() DmsPrimaryKey {
	if o == nil || o.PrimaryKey == nil {
		var ret DmsPrimaryKey
		return ret
	}
	return *o.PrimaryKey
}

// GetPrimaryKeyOk returns a tuple with the PrimaryKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsTableMetadata) GetPrimaryKeyOk() (*DmsPrimaryKey, bool) {
	if o == nil || o.PrimaryKey == nil {
		return nil, false
	}
	return o.PrimaryKey, true
}

// HasPrimaryKey returns a boolean if a field has been set.
func (o *DmsTableMetadata) HasPrimaryKey() bool {
	return o != nil && o.PrimaryKey != nil
}

// SetPrimaryKey gets a reference to the given DmsPrimaryKey and assigns it to the PrimaryKey field.
func (o *DmsTableMetadata) SetPrimaryKey(v DmsPrimaryKey) {
	o.PrimaryKey = &v
}

// GetForeignKeys returns the ForeignKeys field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsTableMetadata) GetForeignKeys() []DmsForeignKey {
	if o == nil {
		var ret []DmsForeignKey
		return ret
	}
	return o.ForeignKeys
}

// GetForeignKeysOk returns a tuple with the ForeignKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsTableMetadata) GetForeignKeysOk() (*[]DmsForeignKey, bool) {
	if o == nil || o.ForeignKeys == nil {
		return nil, false
	}
	return &o.ForeignKeys, true
}

// HasForeignKeys returns a boolean if a field has been set.
func (o *DmsTableMetadata) HasForeignKeys() bool {
	return o != nil && o.ForeignKeys != nil
}

// SetForeignKeys gets a reference to the given []DmsForeignKey and assigns it to the ForeignKeys field.
func (o *DmsTableMetadata) SetForeignKeys(v []DmsForeignKey) {
	o.ForeignKeys = v
}

// GetUniqueKeys returns the UniqueKeys field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsTableMetadata) GetUniqueKeys() []DmsUniqueKey {
	if o == nil {
		var ret []DmsUniqueKey
		return ret
	}
	return o.UniqueKeys
}

// GetUniqueKeysOk returns a tuple with the UniqueKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsTableMetadata) GetUniqueKeysOk() (*[]DmsUniqueKey, bool) {
	if o == nil || o.UniqueKeys == nil {
		return nil, false
	}
	return &o.UniqueKeys, true
}

// HasUniqueKeys returns a boolean if a field has been set.
func (o *DmsTableMetadata) HasUniqueKeys() bool {
	return o != nil && o.UniqueKeys != nil
}

// SetUniqueKeys gets a reference to the given []DmsUniqueKey and assigns it to the UniqueKeys field.
func (o *DmsTableMetadata) SetUniqueKeys(v []DmsUniqueKey) {
	o.UniqueKeys = v
}

// GetCheckConstraints returns the CheckConstraints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsTableMetadata) GetCheckConstraints() []DmsCheckConstraint {
	if o == nil {
		var ret []DmsCheckConstraint
		return ret
	}
	return o.CheckConstraints
}

// GetCheckConstraintsOk returns a tuple with the CheckConstraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsTableMetadata) GetCheckConstraintsOk() (*[]DmsCheckConstraint, bool) {
	if o == nil || o.CheckConstraints == nil {
		return nil, false
	}
	return &o.CheckConstraints, true
}

// HasCheckConstraints returns a boolean if a field has been set.
func (o *DmsTableMetadata) HasCheckConstraints() bool {
	return o != nil && o.CheckConstraints != nil
}

// SetCheckConstraints gets a reference to the given []DmsCheckConstraint and assigns it to the CheckConstraints field.
func (o *DmsTableMetadata) SetCheckConstraints(v []DmsCheckConstraint) {
	o.CheckConstraints = v
}

// GetExcludeConstraints returns the ExcludeConstraints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsTableMetadata) GetExcludeConstraints() []DmsExcludeConstraint {
	if o == nil {
		var ret []DmsExcludeConstraint
		return ret
	}
	return o.ExcludeConstraints
}

// GetExcludeConstraintsOk returns a tuple with the ExcludeConstraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsTableMetadata) GetExcludeConstraintsOk() (*[]DmsExcludeConstraint, bool) {
	if o == nil || o.ExcludeConstraints == nil {
		return nil, false
	}
	return &o.ExcludeConstraints, true
}

// HasExcludeConstraints returns a boolean if a field has been set.
func (o *DmsTableMetadata) HasExcludeConstraints() bool {
	return o != nil && o.ExcludeConstraints != nil
}

// SetExcludeConstraints gets a reference to the given []DmsExcludeConstraint and assigns it to the ExcludeConstraints field.
func (o *DmsTableMetadata) SetExcludeConstraints(v []DmsExcludeConstraint) {
	o.ExcludeConstraints = v
}

// GetPartitioning returns the Partitioning field value if set, zero value otherwise.
func (o *DmsTableMetadata) GetPartitioning() DmsTablePartitioning {
	if o == nil || o.Partitioning == nil {
		var ret DmsTablePartitioning
		return ret
	}
	return *o.Partitioning
}

// GetPartitioningOk returns a tuple with the Partitioning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsTableMetadata) GetPartitioningOk() (*DmsTablePartitioning, bool) {
	if o == nil || o.Partitioning == nil {
		return nil, false
	}
	return o.Partitioning, true
}

// HasPartitioning returns a boolean if a field has been set.
func (o *DmsTableMetadata) HasPartitioning() bool {
	return o != nil && o.Partitioning != nil
}

// SetPartitioning gets a reference to the given DmsTablePartitioning and assigns it to the Partitioning field.
func (o *DmsTableMetadata) SetPartitioning(v DmsTablePartitioning) {
	o.Partitioning = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *DmsTableMetadata) GetOptions() DmsTableOptions {
	if o == nil || o.Options == nil {
		var ret DmsTableOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsTableMetadata) GetOptionsOk() (*DmsTableOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *DmsTableMetadata) HasOptions() bool {
	return o != nil && o.Options != nil
}

// SetOptions gets a reference to the given DmsTableOptions and assigns it to the Options field.
func (o *DmsTableMetadata) SetOptions(v DmsTableOptions) {
	o.Options = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DmsTableMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Schema.IsSet() {
		toSerialize["schema"] = o.Schema.Get()
	}
	if o.Database.IsSet() {
		toSerialize["database"] = o.Database.Get()
	}
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	if o.Indexes != nil {
		toSerialize["indexes"] = o.Indexes
	}
	if o.PrimaryKey != nil {
		toSerialize["primaryKey"] = o.PrimaryKey
	}
	if o.ForeignKeys != nil {
		toSerialize["foreignKeys"] = o.ForeignKeys
	}
	if o.UniqueKeys != nil {
		toSerialize["uniqueKeys"] = o.UniqueKeys
	}
	if o.CheckConstraints != nil {
		toSerialize["checkConstraints"] = o.CheckConstraints
	}
	if o.ExcludeConstraints != nil {
		toSerialize["excludeConstraints"] = o.ExcludeConstraints
	}
	if o.Partitioning != nil {
		toSerialize["partitioning"] = o.Partitioning
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmsTableMetadata) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name               *string                `json:"name,omitempty"`
		Schema             common.NullableString  `json:"schema,omitempty"`
		Database           common.NullableString  `json:"database,omitempty"`
		Comment            common.NullableString  `json:"comment,omitempty"`
		Columns            []DmsTableColumn       `json:"columns,omitempty"`
		Indexes            []DmsTableIndex        `json:"indexes,omitempty"`
		PrimaryKey         *DmsPrimaryKey         `json:"primaryKey,omitempty"`
		ForeignKeys        []DmsForeignKey        `json:"foreignKeys,omitempty"`
		UniqueKeys         []DmsUniqueKey         `json:"uniqueKeys,omitempty"`
		CheckConstraints   []DmsCheckConstraint   `json:"checkConstraints,omitempty"`
		ExcludeConstraints []DmsExcludeConstraint `json:"excludeConstraints,omitempty"`
		Partitioning       *DmsTablePartitioning  `json:"partitioning,omitempty"`
		Options            *DmsTableOptions       `json:"options,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "schema", "database", "comment", "columns", "indexes", "primaryKey", "foreignKeys", "uniqueKeys", "checkConstraints", "excludeConstraints", "partitioning", "options"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = all.Name
	o.Schema = all.Schema
	o.Database = all.Database
	o.Comment = all.Comment
	o.Columns = all.Columns
	o.Indexes = all.Indexes
	if all.PrimaryKey != nil && all.PrimaryKey.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.PrimaryKey = all.PrimaryKey
	o.ForeignKeys = all.ForeignKeys
	o.UniqueKeys = all.UniqueKeys
	o.CheckConstraints = all.CheckConstraints
	o.ExcludeConstraints = all.ExcludeConstraints
	if all.Partitioning != nil && all.Partitioning.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Partitioning = all.Partitioning
	if all.Options != nil && all.Options.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Options = all.Options

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
