// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DamengIndexItem struct {
	// Schema (owner) name.
	Schema string `json:"schema"`
	// Index name.
	IndexName string `json:"indexName"`
	// Parent table name.
	TableName string `json:"tableName"`
	// Segment size in bytes from DBA_SEGMENTS.
	SizeBytes int64 `json:"sizeBytes"`
	// Index type (e.g. NORMAL, BITMAP, CLUSTER).
	IndexType *string `json:"indexType,omitempty"`
	// Whether the index enforces uniqueness.
	IsUnique bool `json:"isUnique"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDamengIndexItem instantiates a new DamengIndexItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDamengIndexItem(schema string, indexName string, tableName string, sizeBytes int64, isUnique bool) *DamengIndexItem {
	this := DamengIndexItem{}
	this.Schema = schema
	this.IndexName = indexName
	this.TableName = tableName
	this.SizeBytes = sizeBytes
	this.IsUnique = isUnique
	return &this
}

// NewDamengIndexItemWithDefaults instantiates a new DamengIndexItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDamengIndexItemWithDefaults() *DamengIndexItem {
	this := DamengIndexItem{}
	return &this
}

// GetSchema returns the Schema field value.
func (o *DamengIndexItem) GetSchema() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value
// and a boolean to check if the value has been set.
func (o *DamengIndexItem) GetSchemaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schema, true
}

// SetSchema sets field value.
func (o *DamengIndexItem) SetSchema(v string) {
	o.Schema = v
}

// GetIndexName returns the IndexName field value.
func (o *DamengIndexItem) GetIndexName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.IndexName
}

// GetIndexNameOk returns a tuple with the IndexName field value
// and a boolean to check if the value has been set.
func (o *DamengIndexItem) GetIndexNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IndexName, true
}

// SetIndexName sets field value.
func (o *DamengIndexItem) SetIndexName(v string) {
	o.IndexName = v
}

// GetTableName returns the TableName field value.
func (o *DamengIndexItem) GetTableName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TableName
}

// GetTableNameOk returns a tuple with the TableName field value
// and a boolean to check if the value has been set.
func (o *DamengIndexItem) GetTableNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TableName, true
}

// SetTableName sets field value.
func (o *DamengIndexItem) SetTableName(v string) {
	o.TableName = v
}

// GetSizeBytes returns the SizeBytes field value.
func (o *DamengIndexItem) GetSizeBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.SizeBytes
}

// GetSizeBytesOk returns a tuple with the SizeBytes field value
// and a boolean to check if the value has been set.
func (o *DamengIndexItem) GetSizeBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SizeBytes, true
}

// SetSizeBytes sets field value.
func (o *DamengIndexItem) SetSizeBytes(v int64) {
	o.SizeBytes = v
}

// GetIndexType returns the IndexType field value if set, zero value otherwise.
func (o *DamengIndexItem) GetIndexType() string {
	if o == nil || o.IndexType == nil {
		var ret string
		return ret
	}
	return *o.IndexType
}

// GetIndexTypeOk returns a tuple with the IndexType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamengIndexItem) GetIndexTypeOk() (*string, bool) {
	if o == nil || o.IndexType == nil {
		return nil, false
	}
	return o.IndexType, true
}

// HasIndexType returns a boolean if a field has been set.
func (o *DamengIndexItem) HasIndexType() bool {
	return o != nil && o.IndexType != nil
}

// SetIndexType gets a reference to the given string and assigns it to the IndexType field.
func (o *DamengIndexItem) SetIndexType(v string) {
	o.IndexType = &v
}

// GetIsUnique returns the IsUnique field value.
func (o *DamengIndexItem) GetIsUnique() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsUnique
}

// GetIsUniqueOk returns a tuple with the IsUnique field value
// and a boolean to check if the value has been set.
func (o *DamengIndexItem) GetIsUniqueOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsUnique, true
}

// SetIsUnique sets field value.
func (o *DamengIndexItem) SetIsUnique(v bool) {
	o.IsUnique = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DamengIndexItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["schema"] = o.Schema
	toSerialize["indexName"] = o.IndexName
	toSerialize["tableName"] = o.TableName
	toSerialize["sizeBytes"] = o.SizeBytes
	if o.IndexType != nil {
		toSerialize["indexType"] = o.IndexType
	}
	toSerialize["isUnique"] = o.IsUnique

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DamengIndexItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Schema    *string `json:"schema"`
		IndexName *string `json:"indexName"`
		TableName *string `json:"tableName"`
		SizeBytes *int64  `json:"sizeBytes"`
		IndexType *string `json:"indexType,omitempty"`
		IsUnique  *bool   `json:"isUnique"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Schema == nil {
		return fmt.Errorf("required field schema missing")
	}
	if all.IndexName == nil {
		return fmt.Errorf("required field indexName missing")
	}
	if all.TableName == nil {
		return fmt.Errorf("required field tableName missing")
	}
	if all.SizeBytes == nil {
		return fmt.Errorf("required field sizeBytes missing")
	}
	if all.IsUnique == nil {
		return fmt.Errorf("required field isUnique missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"schema", "indexName", "tableName", "sizeBytes", "indexType", "isUnique"})
	} else {
		return err
	}
	o.Schema = *all.Schema
	o.IndexName = *all.IndexName
	o.TableName = *all.TableName
	o.SizeBytes = *all.SizeBytes
	o.IndexType = all.IndexType
	o.IsUnique = *all.IsUnique

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
