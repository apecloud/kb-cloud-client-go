// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type MysqlIndexSpace struct {
	Database string `json:"database"`
	Table    string `json:"table"`
	Name     string `json:"name"`
	// Estimated from mysql.innodb_index_stats pages multiplied by innodb_page_size.
	SizeBytes   common.NullableInt64 `json:"sizeBytes,omitempty"`
	IsUnique    bool                 `json:"isUnique"`
	IsPrimary   bool                 `json:"isPrimary"`
	Cardinality int64                `json:"cardinality"`
	// COUNT_READ from performance_schema; resets with server statistics.
	ScanCount      common.NullableInt64 `json:"scanCount,omitempty"`
	LastStatUpdate *string              `json:"lastStatUpdate,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMysqlIndexSpace instantiates a new MysqlIndexSpace object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMysqlIndexSpace(database string, table string, name string, isUnique bool, isPrimary bool, cardinality int64) *MysqlIndexSpace {
	this := MysqlIndexSpace{}
	this.Database = database
	this.Table = table
	this.Name = name
	this.IsUnique = isUnique
	this.IsPrimary = isPrimary
	this.Cardinality = cardinality
	return &this
}

// NewMysqlIndexSpaceWithDefaults instantiates a new MysqlIndexSpace object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMysqlIndexSpaceWithDefaults() *MysqlIndexSpace {
	this := MysqlIndexSpace{}
	return &this
}

// GetDatabase returns the Database field value.
func (o *MysqlIndexSpace) GetDatabase() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Database
}

// GetDatabaseOk returns a tuple with the Database field value
// and a boolean to check if the value has been set.
func (o *MysqlIndexSpace) GetDatabaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Database, true
}

// SetDatabase sets field value.
func (o *MysqlIndexSpace) SetDatabase(v string) {
	o.Database = v
}

// GetTable returns the Table field value.
func (o *MysqlIndexSpace) GetTable() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Table
}

// GetTableOk returns a tuple with the Table field value
// and a boolean to check if the value has been set.
func (o *MysqlIndexSpace) GetTableOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Table, true
}

// SetTable sets field value.
func (o *MysqlIndexSpace) SetTable(v string) {
	o.Table = v
}

// GetName returns the Name field value.
func (o *MysqlIndexSpace) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MysqlIndexSpace) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *MysqlIndexSpace) SetName(v string) {
	o.Name = v
}

// GetSizeBytes returns the SizeBytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MysqlIndexSpace) GetSizeBytes() int64 {
	if o == nil || o.SizeBytes.Get() == nil {
		var ret int64
		return ret
	}
	return *o.SizeBytes.Get()
}

// GetSizeBytesOk returns a tuple with the SizeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MysqlIndexSpace) GetSizeBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SizeBytes.Get(), o.SizeBytes.IsSet()
}

// HasSizeBytes returns a boolean if a field has been set.
func (o *MysqlIndexSpace) HasSizeBytes() bool {
	return o != nil && o.SizeBytes.IsSet()
}

// SetSizeBytes gets a reference to the given common.NullableInt64 and assigns it to the SizeBytes field.
func (o *MysqlIndexSpace) SetSizeBytes(v int64) {
	o.SizeBytes.Set(&v)
}

// SetSizeBytesNil sets the value for SizeBytes to be an explicit nil.
func (o *MysqlIndexSpace) SetSizeBytesNil() {
	o.SizeBytes.Set(nil)
}

// UnsetSizeBytes ensures that no value is present for SizeBytes, not even an explicit nil.
func (o *MysqlIndexSpace) UnsetSizeBytes() {
	o.SizeBytes.Unset()
}

// GetIsUnique returns the IsUnique field value.
func (o *MysqlIndexSpace) GetIsUnique() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsUnique
}

// GetIsUniqueOk returns a tuple with the IsUnique field value
// and a boolean to check if the value has been set.
func (o *MysqlIndexSpace) GetIsUniqueOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsUnique, true
}

// SetIsUnique sets field value.
func (o *MysqlIndexSpace) SetIsUnique(v bool) {
	o.IsUnique = v
}

// GetIsPrimary returns the IsPrimary field value.
func (o *MysqlIndexSpace) GetIsPrimary() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsPrimary
}

// GetIsPrimaryOk returns a tuple with the IsPrimary field value
// and a boolean to check if the value has been set.
func (o *MysqlIndexSpace) GetIsPrimaryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPrimary, true
}

// SetIsPrimary sets field value.
func (o *MysqlIndexSpace) SetIsPrimary(v bool) {
	o.IsPrimary = v
}

// GetCardinality returns the Cardinality field value.
func (o *MysqlIndexSpace) GetCardinality() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Cardinality
}

// GetCardinalityOk returns a tuple with the Cardinality field value
// and a boolean to check if the value has been set.
func (o *MysqlIndexSpace) GetCardinalityOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cardinality, true
}

// SetCardinality sets field value.
func (o *MysqlIndexSpace) SetCardinality(v int64) {
	o.Cardinality = v
}

// GetScanCount returns the ScanCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MysqlIndexSpace) GetScanCount() int64 {
	if o == nil || o.ScanCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ScanCount.Get()
}

// GetScanCountOk returns a tuple with the ScanCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MysqlIndexSpace) GetScanCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScanCount.Get(), o.ScanCount.IsSet()
}

// HasScanCount returns a boolean if a field has been set.
func (o *MysqlIndexSpace) HasScanCount() bool {
	return o != nil && o.ScanCount.IsSet()
}

// SetScanCount gets a reference to the given common.NullableInt64 and assigns it to the ScanCount field.
func (o *MysqlIndexSpace) SetScanCount(v int64) {
	o.ScanCount.Set(&v)
}

// SetScanCountNil sets the value for ScanCount to be an explicit nil.
func (o *MysqlIndexSpace) SetScanCountNil() {
	o.ScanCount.Set(nil)
}

// UnsetScanCount ensures that no value is present for ScanCount, not even an explicit nil.
func (o *MysqlIndexSpace) UnsetScanCount() {
	o.ScanCount.Unset()
}

// GetLastStatUpdate returns the LastStatUpdate field value if set, zero value otherwise.
func (o *MysqlIndexSpace) GetLastStatUpdate() string {
	if o == nil || o.LastStatUpdate == nil {
		var ret string
		return ret
	}
	return *o.LastStatUpdate
}

// GetLastStatUpdateOk returns a tuple with the LastStatUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlIndexSpace) GetLastStatUpdateOk() (*string, bool) {
	if o == nil || o.LastStatUpdate == nil {
		return nil, false
	}
	return o.LastStatUpdate, true
}

// HasLastStatUpdate returns a boolean if a field has been set.
func (o *MysqlIndexSpace) HasLastStatUpdate() bool {
	return o != nil && o.LastStatUpdate != nil
}

// SetLastStatUpdate gets a reference to the given string and assigns it to the LastStatUpdate field.
func (o *MysqlIndexSpace) SetLastStatUpdate(v string) {
	o.LastStatUpdate = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MysqlIndexSpace) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["database"] = o.Database
	toSerialize["table"] = o.Table
	toSerialize["name"] = o.Name
	if o.SizeBytes.IsSet() {
		toSerialize["sizeBytes"] = o.SizeBytes.Get()
	}
	toSerialize["isUnique"] = o.IsUnique
	toSerialize["isPrimary"] = o.IsPrimary
	toSerialize["cardinality"] = o.Cardinality
	if o.ScanCount.IsSet() {
		toSerialize["scanCount"] = o.ScanCount.Get()
	}
	if o.LastStatUpdate != nil {
		toSerialize["lastStatUpdate"] = o.LastStatUpdate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MysqlIndexSpace) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Database       *string              `json:"database"`
		Table          *string              `json:"table"`
		Name           *string              `json:"name"`
		SizeBytes      common.NullableInt64 `json:"sizeBytes,omitempty"`
		IsUnique       *bool                `json:"isUnique"`
		IsPrimary      *bool                `json:"isPrimary"`
		Cardinality    *int64               `json:"cardinality"`
		ScanCount      common.NullableInt64 `json:"scanCount,omitempty"`
		LastStatUpdate *string              `json:"lastStatUpdate,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Database == nil {
		return fmt.Errorf("required field database missing")
	}
	if all.Table == nil {
		return fmt.Errorf("required field table missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.IsUnique == nil {
		return fmt.Errorf("required field isUnique missing")
	}
	if all.IsPrimary == nil {
		return fmt.Errorf("required field isPrimary missing")
	}
	if all.Cardinality == nil {
		return fmt.Errorf("required field cardinality missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"database", "table", "name", "sizeBytes", "isUnique", "isPrimary", "cardinality", "scanCount", "lastStatUpdate"})
	} else {
		return err
	}
	o.Database = *all.Database
	o.Table = *all.Table
	o.Name = *all.Name
	o.SizeBytes = all.SizeBytes
	o.IsUnique = *all.IsUnique
	o.IsPrimary = *all.IsPrimary
	o.Cardinality = *all.Cardinality
	o.ScanCount = all.ScanCount
	o.LastStatUpdate = all.LastStatUpdate

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
