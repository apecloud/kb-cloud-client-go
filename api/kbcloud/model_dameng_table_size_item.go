// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DamengTableSizeItem struct {
	// Schema (owner) name.
	Schema string `json:"schema"`
	// Table name.
	TableName string `json:"tableName"`
	// Total segment size in bytes for this table.
	SizeBytes int64 `json:"sizeBytes"`
	// Estimated number of rows from DBA_TABLES.NUM_ROWS.
	RowCount *int64 `json:"rowCount,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDamengTableSizeItem instantiates a new DamengTableSizeItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDamengTableSizeItem(schema string, tableName string, sizeBytes int64) *DamengTableSizeItem {
	this := DamengTableSizeItem{}
	this.Schema = schema
	this.TableName = tableName
	this.SizeBytes = sizeBytes
	return &this
}

// NewDamengTableSizeItemWithDefaults instantiates a new DamengTableSizeItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDamengTableSizeItemWithDefaults() *DamengTableSizeItem {
	this := DamengTableSizeItem{}
	return &this
}

// GetSchema returns the Schema field value.
func (o *DamengTableSizeItem) GetSchema() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value
// and a boolean to check if the value has been set.
func (o *DamengTableSizeItem) GetSchemaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schema, true
}

// SetSchema sets field value.
func (o *DamengTableSizeItem) SetSchema(v string) {
	o.Schema = v
}

// GetTableName returns the TableName field value.
func (o *DamengTableSizeItem) GetTableName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TableName
}

// GetTableNameOk returns a tuple with the TableName field value
// and a boolean to check if the value has been set.
func (o *DamengTableSizeItem) GetTableNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TableName, true
}

// SetTableName sets field value.
func (o *DamengTableSizeItem) SetTableName(v string) {
	o.TableName = v
}

// GetSizeBytes returns the SizeBytes field value.
func (o *DamengTableSizeItem) GetSizeBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.SizeBytes
}

// GetSizeBytesOk returns a tuple with the SizeBytes field value
// and a boolean to check if the value has been set.
func (o *DamengTableSizeItem) GetSizeBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SizeBytes, true
}

// SetSizeBytes sets field value.
func (o *DamengTableSizeItem) SetSizeBytes(v int64) {
	o.SizeBytes = v
}

// GetRowCount returns the RowCount field value if set, zero value otherwise.
func (o *DamengTableSizeItem) GetRowCount() int64 {
	if o == nil || o.RowCount == nil {
		var ret int64
		return ret
	}
	return *o.RowCount
}

// GetRowCountOk returns a tuple with the RowCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamengTableSizeItem) GetRowCountOk() (*int64, bool) {
	if o == nil || o.RowCount == nil {
		return nil, false
	}
	return o.RowCount, true
}

// HasRowCount returns a boolean if a field has been set.
func (o *DamengTableSizeItem) HasRowCount() bool {
	return o != nil && o.RowCount != nil
}

// SetRowCount gets a reference to the given int64 and assigns it to the RowCount field.
func (o *DamengTableSizeItem) SetRowCount(v int64) {
	o.RowCount = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DamengTableSizeItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["schema"] = o.Schema
	toSerialize["tableName"] = o.TableName
	toSerialize["sizeBytes"] = o.SizeBytes
	if o.RowCount != nil {
		toSerialize["rowCount"] = o.RowCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DamengTableSizeItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Schema    *string `json:"schema"`
		TableName *string `json:"tableName"`
		SizeBytes *int64  `json:"sizeBytes"`
		RowCount  *int64  `json:"rowCount,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Schema == nil {
		return fmt.Errorf("required field schema missing")
	}
	if all.TableName == nil {
		return fmt.Errorf("required field tableName missing")
	}
	if all.SizeBytes == nil {
		return fmt.Errorf("required field sizeBytes missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"schema", "tableName", "sizeBytes", "rowCount"})
	} else {
		return err
	}
	o.Schema = *all.Schema
	o.TableName = *all.TableName
	o.SizeBytes = *all.SizeBytes
	o.RowCount = all.RowCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
