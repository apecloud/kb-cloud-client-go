// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type MysqlSpaceSummary struct {
	DatabaseName      string `json:"databaseName"`
	DatabaseSizeBytes int64  `json:"databaseSizeBytes"`
	DataBytes         int64  `json:"dataBytes"`
	IndexBytes        int64  `json:"indexBytes"`
	// DATA_FREE reported by information_schema. Its per-table meaning is most reliable with innodb_file_per_table.
	FreeBytes          int64   `json:"freeBytes"`
	TableCount         int64   `json:"tableCount"`
	TableListTruncated bool    `json:"tableListTruncated"`
	IndexListTruncated bool    `json:"indexListTruncated"`
	IndexSizeSource    *string `json:"indexSizeSource,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMysqlSpaceSummary instantiates a new MysqlSpaceSummary object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMysqlSpaceSummary(databaseName string, databaseSizeBytes int64, dataBytes int64, indexBytes int64, freeBytes int64, tableCount int64, tableListTruncated bool, indexListTruncated bool) *MysqlSpaceSummary {
	this := MysqlSpaceSummary{}
	this.DatabaseName = databaseName
	this.DatabaseSizeBytes = databaseSizeBytes
	this.DataBytes = dataBytes
	this.IndexBytes = indexBytes
	this.FreeBytes = freeBytes
	this.TableCount = tableCount
	this.TableListTruncated = tableListTruncated
	this.IndexListTruncated = indexListTruncated
	return &this
}

// NewMysqlSpaceSummaryWithDefaults instantiates a new MysqlSpaceSummary object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMysqlSpaceSummaryWithDefaults() *MysqlSpaceSummary {
	this := MysqlSpaceSummary{}
	return &this
}

// GetDatabaseName returns the DatabaseName field value.
func (o *MysqlSpaceSummary) GetDatabaseName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DatabaseName
}

// GetDatabaseNameOk returns a tuple with the DatabaseName field value
// and a boolean to check if the value has been set.
func (o *MysqlSpaceSummary) GetDatabaseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatabaseName, true
}

// SetDatabaseName sets field value.
func (o *MysqlSpaceSummary) SetDatabaseName(v string) {
	o.DatabaseName = v
}

// GetDatabaseSizeBytes returns the DatabaseSizeBytes field value.
func (o *MysqlSpaceSummary) GetDatabaseSizeBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.DatabaseSizeBytes
}

// GetDatabaseSizeBytesOk returns a tuple with the DatabaseSizeBytes field value
// and a boolean to check if the value has been set.
func (o *MysqlSpaceSummary) GetDatabaseSizeBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatabaseSizeBytes, true
}

// SetDatabaseSizeBytes sets field value.
func (o *MysqlSpaceSummary) SetDatabaseSizeBytes(v int64) {
	o.DatabaseSizeBytes = v
}

// GetDataBytes returns the DataBytes field value.
func (o *MysqlSpaceSummary) GetDataBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.DataBytes
}

// GetDataBytesOk returns a tuple with the DataBytes field value
// and a boolean to check if the value has been set.
func (o *MysqlSpaceSummary) GetDataBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataBytes, true
}

// SetDataBytes sets field value.
func (o *MysqlSpaceSummary) SetDataBytes(v int64) {
	o.DataBytes = v
}

// GetIndexBytes returns the IndexBytes field value.
func (o *MysqlSpaceSummary) GetIndexBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.IndexBytes
}

// GetIndexBytesOk returns a tuple with the IndexBytes field value
// and a boolean to check if the value has been set.
func (o *MysqlSpaceSummary) GetIndexBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IndexBytes, true
}

// SetIndexBytes sets field value.
func (o *MysqlSpaceSummary) SetIndexBytes(v int64) {
	o.IndexBytes = v
}

// GetFreeBytes returns the FreeBytes field value.
func (o *MysqlSpaceSummary) GetFreeBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.FreeBytes
}

// GetFreeBytesOk returns a tuple with the FreeBytes field value
// and a boolean to check if the value has been set.
func (o *MysqlSpaceSummary) GetFreeBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FreeBytes, true
}

// SetFreeBytes sets field value.
func (o *MysqlSpaceSummary) SetFreeBytes(v int64) {
	o.FreeBytes = v
}

// GetTableCount returns the TableCount field value.
func (o *MysqlSpaceSummary) GetTableCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TableCount
}

// GetTableCountOk returns a tuple with the TableCount field value
// and a boolean to check if the value has been set.
func (o *MysqlSpaceSummary) GetTableCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TableCount, true
}

// SetTableCount sets field value.
func (o *MysqlSpaceSummary) SetTableCount(v int64) {
	o.TableCount = v
}

// GetTableListTruncated returns the TableListTruncated field value.
func (o *MysqlSpaceSummary) GetTableListTruncated() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.TableListTruncated
}

// GetTableListTruncatedOk returns a tuple with the TableListTruncated field value
// and a boolean to check if the value has been set.
func (o *MysqlSpaceSummary) GetTableListTruncatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TableListTruncated, true
}

// SetTableListTruncated sets field value.
func (o *MysqlSpaceSummary) SetTableListTruncated(v bool) {
	o.TableListTruncated = v
}

// GetIndexListTruncated returns the IndexListTruncated field value.
func (o *MysqlSpaceSummary) GetIndexListTruncated() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IndexListTruncated
}

// GetIndexListTruncatedOk returns a tuple with the IndexListTruncated field value
// and a boolean to check if the value has been set.
func (o *MysqlSpaceSummary) GetIndexListTruncatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IndexListTruncated, true
}

// SetIndexListTruncated sets field value.
func (o *MysqlSpaceSummary) SetIndexListTruncated(v bool) {
	o.IndexListTruncated = v
}

// GetIndexSizeSource returns the IndexSizeSource field value if set, zero value otherwise.
func (o *MysqlSpaceSummary) GetIndexSizeSource() string {
	if o == nil || o.IndexSizeSource == nil {
		var ret string
		return ret
	}
	return *o.IndexSizeSource
}

// GetIndexSizeSourceOk returns a tuple with the IndexSizeSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlSpaceSummary) GetIndexSizeSourceOk() (*string, bool) {
	if o == nil || o.IndexSizeSource == nil {
		return nil, false
	}
	return o.IndexSizeSource, true
}

// HasIndexSizeSource returns a boolean if a field has been set.
func (o *MysqlSpaceSummary) HasIndexSizeSource() bool {
	return o != nil && o.IndexSizeSource != nil
}

// SetIndexSizeSource gets a reference to the given string and assigns it to the IndexSizeSource field.
func (o *MysqlSpaceSummary) SetIndexSizeSource(v string) {
	o.IndexSizeSource = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MysqlSpaceSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["databaseName"] = o.DatabaseName
	toSerialize["databaseSizeBytes"] = o.DatabaseSizeBytes
	toSerialize["dataBytes"] = o.DataBytes
	toSerialize["indexBytes"] = o.IndexBytes
	toSerialize["freeBytes"] = o.FreeBytes
	toSerialize["tableCount"] = o.TableCount
	toSerialize["tableListTruncated"] = o.TableListTruncated
	toSerialize["indexListTruncated"] = o.IndexListTruncated
	if o.IndexSizeSource != nil {
		toSerialize["indexSizeSource"] = o.IndexSizeSource
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MysqlSpaceSummary) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DatabaseName       *string `json:"databaseName"`
		DatabaseSizeBytes  *int64  `json:"databaseSizeBytes"`
		DataBytes          *int64  `json:"dataBytes"`
		IndexBytes         *int64  `json:"indexBytes"`
		FreeBytes          *int64  `json:"freeBytes"`
		TableCount         *int64  `json:"tableCount"`
		TableListTruncated *bool   `json:"tableListTruncated"`
		IndexListTruncated *bool   `json:"indexListTruncated"`
		IndexSizeSource    *string `json:"indexSizeSource,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.DatabaseName == nil {
		return fmt.Errorf("required field databaseName missing")
	}
	if all.DatabaseSizeBytes == nil {
		return fmt.Errorf("required field databaseSizeBytes missing")
	}
	if all.DataBytes == nil {
		return fmt.Errorf("required field dataBytes missing")
	}
	if all.IndexBytes == nil {
		return fmt.Errorf("required field indexBytes missing")
	}
	if all.FreeBytes == nil {
		return fmt.Errorf("required field freeBytes missing")
	}
	if all.TableCount == nil {
		return fmt.Errorf("required field tableCount missing")
	}
	if all.TableListTruncated == nil {
		return fmt.Errorf("required field tableListTruncated missing")
	}
	if all.IndexListTruncated == nil {
		return fmt.Errorf("required field indexListTruncated missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"databaseName", "databaseSizeBytes", "dataBytes", "indexBytes", "freeBytes", "tableCount", "tableListTruncated", "indexListTruncated", "indexSizeSource"})
	} else {
		return err
	}
	o.DatabaseName = *all.DatabaseName
	o.DatabaseSizeBytes = *all.DatabaseSizeBytes
	o.DataBytes = *all.DataBytes
	o.IndexBytes = *all.IndexBytes
	o.FreeBytes = *all.FreeBytes
	o.TableCount = *all.TableCount
	o.TableListTruncated = *all.TableListTruncated
	o.IndexListTruncated = *all.IndexListTruncated
	o.IndexSizeSource = all.IndexSizeSource

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
