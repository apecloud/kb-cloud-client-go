// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type PostgresqlIndexSpace struct {
	Database         *string `json:"database,omitempty"`
	Schema           *string `json:"schema,omitempty"`
	Table            *string `json:"table,omitempty"`
	Name             *string `json:"name,omitempty"`
	TableRelationOid *int64  `json:"tableRelationOid,omitempty"`
	RelationOid      *int64  `json:"relationOid,omitempty"`
	SizeBytes        *int64  `json:"sizeBytes,omitempty"`
	IsPrimary        *bool   `json:"isPrimary,omitempty"`
	IsUnique         *bool   `json:"isUnique,omitempty"`
	ScanCount        *int64  `json:"scanCount,omitempty"`
	TuplesRead       *int64  `json:"tuplesRead,omitempty"`
	TuplesFetch      *int64  `json:"tuplesFetch,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlIndexSpace instantiates a new PostgresqlIndexSpace object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlIndexSpace() *PostgresqlIndexSpace {
	this := PostgresqlIndexSpace{}
	return &this
}

// NewPostgresqlIndexSpaceWithDefaults instantiates a new PostgresqlIndexSpace object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlIndexSpaceWithDefaults() *PostgresqlIndexSpace {
	this := PostgresqlIndexSpace{}
	return &this
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *PostgresqlIndexSpace) GetDatabase() string {
	if o == nil || o.Database == nil {
		var ret string
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlIndexSpace) GetDatabaseOk() (*string, bool) {
	if o == nil || o.Database == nil {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *PostgresqlIndexSpace) HasDatabase() bool {
	return o != nil && o.Database != nil
}

// SetDatabase gets a reference to the given string and assigns it to the Database field.
func (o *PostgresqlIndexSpace) SetDatabase(v string) {
	o.Database = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *PostgresqlIndexSpace) GetSchema() string {
	if o == nil || o.Schema == nil {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlIndexSpace) GetSchemaOk() (*string, bool) {
	if o == nil || o.Schema == nil {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *PostgresqlIndexSpace) HasSchema() bool {
	return o != nil && o.Schema != nil
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *PostgresqlIndexSpace) SetSchema(v string) {
	o.Schema = &v
}

// GetTable returns the Table field value if set, zero value otherwise.
func (o *PostgresqlIndexSpace) GetTable() string {
	if o == nil || o.Table == nil {
		var ret string
		return ret
	}
	return *o.Table
}

// GetTableOk returns a tuple with the Table field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlIndexSpace) GetTableOk() (*string, bool) {
	if o == nil || o.Table == nil {
		return nil, false
	}
	return o.Table, true
}

// HasTable returns a boolean if a field has been set.
func (o *PostgresqlIndexSpace) HasTable() bool {
	return o != nil && o.Table != nil
}

// SetTable gets a reference to the given string and assigns it to the Table field.
func (o *PostgresqlIndexSpace) SetTable(v string) {
	o.Table = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PostgresqlIndexSpace) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlIndexSpace) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PostgresqlIndexSpace) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PostgresqlIndexSpace) SetName(v string) {
	o.Name = &v
}

// GetTableRelationOid returns the TableRelationOid field value if set, zero value otherwise.
func (o *PostgresqlIndexSpace) GetTableRelationOid() int64 {
	if o == nil || o.TableRelationOid == nil {
		var ret int64
		return ret
	}
	return *o.TableRelationOid
}

// GetTableRelationOidOk returns a tuple with the TableRelationOid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlIndexSpace) GetTableRelationOidOk() (*int64, bool) {
	if o == nil || o.TableRelationOid == nil {
		return nil, false
	}
	return o.TableRelationOid, true
}

// HasTableRelationOid returns a boolean if a field has been set.
func (o *PostgresqlIndexSpace) HasTableRelationOid() bool {
	return o != nil && o.TableRelationOid != nil
}

// SetTableRelationOid gets a reference to the given int64 and assigns it to the TableRelationOid field.
func (o *PostgresqlIndexSpace) SetTableRelationOid(v int64) {
	o.TableRelationOid = &v
}

// GetRelationOid returns the RelationOid field value if set, zero value otherwise.
func (o *PostgresqlIndexSpace) GetRelationOid() int64 {
	if o == nil || o.RelationOid == nil {
		var ret int64
		return ret
	}
	return *o.RelationOid
}

// GetRelationOidOk returns a tuple with the RelationOid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlIndexSpace) GetRelationOidOk() (*int64, bool) {
	if o == nil || o.RelationOid == nil {
		return nil, false
	}
	return o.RelationOid, true
}

// HasRelationOid returns a boolean if a field has been set.
func (o *PostgresqlIndexSpace) HasRelationOid() bool {
	return o != nil && o.RelationOid != nil
}

// SetRelationOid gets a reference to the given int64 and assigns it to the RelationOid field.
func (o *PostgresqlIndexSpace) SetRelationOid(v int64) {
	o.RelationOid = &v
}

// GetSizeBytes returns the SizeBytes field value if set, zero value otherwise.
func (o *PostgresqlIndexSpace) GetSizeBytes() int64 {
	if o == nil || o.SizeBytes == nil {
		var ret int64
		return ret
	}
	return *o.SizeBytes
}

// GetSizeBytesOk returns a tuple with the SizeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlIndexSpace) GetSizeBytesOk() (*int64, bool) {
	if o == nil || o.SizeBytes == nil {
		return nil, false
	}
	return o.SizeBytes, true
}

// HasSizeBytes returns a boolean if a field has been set.
func (o *PostgresqlIndexSpace) HasSizeBytes() bool {
	return o != nil && o.SizeBytes != nil
}

// SetSizeBytes gets a reference to the given int64 and assigns it to the SizeBytes field.
func (o *PostgresqlIndexSpace) SetSizeBytes(v int64) {
	o.SizeBytes = &v
}

// GetIsPrimary returns the IsPrimary field value if set, zero value otherwise.
func (o *PostgresqlIndexSpace) GetIsPrimary() bool {
	if o == nil || o.IsPrimary == nil {
		var ret bool
		return ret
	}
	return *o.IsPrimary
}

// GetIsPrimaryOk returns a tuple with the IsPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlIndexSpace) GetIsPrimaryOk() (*bool, bool) {
	if o == nil || o.IsPrimary == nil {
		return nil, false
	}
	return o.IsPrimary, true
}

// HasIsPrimary returns a boolean if a field has been set.
func (o *PostgresqlIndexSpace) HasIsPrimary() bool {
	return o != nil && o.IsPrimary != nil
}

// SetIsPrimary gets a reference to the given bool and assigns it to the IsPrimary field.
func (o *PostgresqlIndexSpace) SetIsPrimary(v bool) {
	o.IsPrimary = &v
}

// GetIsUnique returns the IsUnique field value if set, zero value otherwise.
func (o *PostgresqlIndexSpace) GetIsUnique() bool {
	if o == nil || o.IsUnique == nil {
		var ret bool
		return ret
	}
	return *o.IsUnique
}

// GetIsUniqueOk returns a tuple with the IsUnique field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlIndexSpace) GetIsUniqueOk() (*bool, bool) {
	if o == nil || o.IsUnique == nil {
		return nil, false
	}
	return o.IsUnique, true
}

// HasIsUnique returns a boolean if a field has been set.
func (o *PostgresqlIndexSpace) HasIsUnique() bool {
	return o != nil && o.IsUnique != nil
}

// SetIsUnique gets a reference to the given bool and assigns it to the IsUnique field.
func (o *PostgresqlIndexSpace) SetIsUnique(v bool) {
	o.IsUnique = &v
}

// GetScanCount returns the ScanCount field value if set, zero value otherwise.
func (o *PostgresqlIndexSpace) GetScanCount() int64 {
	if o == nil || o.ScanCount == nil {
		var ret int64
		return ret
	}
	return *o.ScanCount
}

// GetScanCountOk returns a tuple with the ScanCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlIndexSpace) GetScanCountOk() (*int64, bool) {
	if o == nil || o.ScanCount == nil {
		return nil, false
	}
	return o.ScanCount, true
}

// HasScanCount returns a boolean if a field has been set.
func (o *PostgresqlIndexSpace) HasScanCount() bool {
	return o != nil && o.ScanCount != nil
}

// SetScanCount gets a reference to the given int64 and assigns it to the ScanCount field.
func (o *PostgresqlIndexSpace) SetScanCount(v int64) {
	o.ScanCount = &v
}

// GetTuplesRead returns the TuplesRead field value if set, zero value otherwise.
func (o *PostgresqlIndexSpace) GetTuplesRead() int64 {
	if o == nil || o.TuplesRead == nil {
		var ret int64
		return ret
	}
	return *o.TuplesRead
}

// GetTuplesReadOk returns a tuple with the TuplesRead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlIndexSpace) GetTuplesReadOk() (*int64, bool) {
	if o == nil || o.TuplesRead == nil {
		return nil, false
	}
	return o.TuplesRead, true
}

// HasTuplesRead returns a boolean if a field has been set.
func (o *PostgresqlIndexSpace) HasTuplesRead() bool {
	return o != nil && o.TuplesRead != nil
}

// SetTuplesRead gets a reference to the given int64 and assigns it to the TuplesRead field.
func (o *PostgresqlIndexSpace) SetTuplesRead(v int64) {
	o.TuplesRead = &v
}

// GetTuplesFetch returns the TuplesFetch field value if set, zero value otherwise.
func (o *PostgresqlIndexSpace) GetTuplesFetch() int64 {
	if o == nil || o.TuplesFetch == nil {
		var ret int64
		return ret
	}
	return *o.TuplesFetch
}

// GetTuplesFetchOk returns a tuple with the TuplesFetch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlIndexSpace) GetTuplesFetchOk() (*int64, bool) {
	if o == nil || o.TuplesFetch == nil {
		return nil, false
	}
	return o.TuplesFetch, true
}

// HasTuplesFetch returns a boolean if a field has been set.
func (o *PostgresqlIndexSpace) HasTuplesFetch() bool {
	return o != nil && o.TuplesFetch != nil
}

// SetTuplesFetch gets a reference to the given int64 and assigns it to the TuplesFetch field.
func (o *PostgresqlIndexSpace) SetTuplesFetch(v int64) {
	o.TuplesFetch = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlIndexSpace) MarshalJSON() ([]byte, error) {
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
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.TableRelationOid != nil {
		toSerialize["tableRelationOid"] = o.TableRelationOid
	}
	if o.RelationOid != nil {
		toSerialize["relationOid"] = o.RelationOid
	}
	if o.SizeBytes != nil {
		toSerialize["sizeBytes"] = o.SizeBytes
	}
	if o.IsPrimary != nil {
		toSerialize["isPrimary"] = o.IsPrimary
	}
	if o.IsUnique != nil {
		toSerialize["isUnique"] = o.IsUnique
	}
	if o.ScanCount != nil {
		toSerialize["scanCount"] = o.ScanCount
	}
	if o.TuplesRead != nil {
		toSerialize["tuplesRead"] = o.TuplesRead
	}
	if o.TuplesFetch != nil {
		toSerialize["tuplesFetch"] = o.TuplesFetch
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostgresqlIndexSpace) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Database         *string `json:"database,omitempty"`
		Schema           *string `json:"schema,omitempty"`
		Table            *string `json:"table,omitempty"`
		Name             *string `json:"name,omitempty"`
		TableRelationOid *int64  `json:"tableRelationOid,omitempty"`
		RelationOid      *int64  `json:"relationOid,omitempty"`
		SizeBytes        *int64  `json:"sizeBytes,omitempty"`
		IsPrimary        *bool   `json:"isPrimary,omitempty"`
		IsUnique         *bool   `json:"isUnique,omitempty"`
		ScanCount        *int64  `json:"scanCount,omitempty"`
		TuplesRead       *int64  `json:"tuplesRead,omitempty"`
		TuplesFetch      *int64  `json:"tuplesFetch,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"database", "schema", "table", "name", "tableRelationOid", "relationOid", "sizeBytes", "isPrimary", "isUnique", "scanCount", "tuplesRead", "tuplesFetch"})
	} else {
		return err
	}
	o.Database = all.Database
	o.Schema = all.Schema
	o.Table = all.Table
	o.Name = all.Name
	o.TableRelationOid = all.TableRelationOid
	o.RelationOid = all.RelationOid
	o.SizeBytes = all.SizeBytes
	o.IsPrimary = all.IsPrimary
	o.IsUnique = all.IsUnique
	o.ScanCount = all.ScanCount
	o.TuplesRead = all.TuplesRead
	o.TuplesFetch = all.TuplesFetch

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
