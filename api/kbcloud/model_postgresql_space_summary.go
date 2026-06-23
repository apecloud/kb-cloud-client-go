// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlSpaceSummary struct {
	DatabaseName       *string `json:"databaseName,omitempty"`
	DatabaseSizeBytes  *int64  `json:"databaseSizeBytes,omitempty"`
	TableCount         int64   `json:"tableCount"`
	IndexCount         int64   `json:"indexCount"`
	ToastRelationCount int64   `json:"toastRelationCount"`
	LargestTableBytes  *int64  `json:"largestTableBytes,omitempty"`
	LargestIndexBytes  *int64  `json:"largestIndexBytes,omitempty"`
	TableListTruncated bool    `json:"tableListTruncated"`
	IndexListTruncated bool    `json:"indexListTruncated"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlSpaceSummary instantiates a new PostgresqlSpaceSummary object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlSpaceSummary(tableCount int64, indexCount int64, toastRelationCount int64, tableListTruncated bool, indexListTruncated bool) *PostgresqlSpaceSummary {
	this := PostgresqlSpaceSummary{}
	this.TableCount = tableCount
	this.IndexCount = indexCount
	this.ToastRelationCount = toastRelationCount
	this.TableListTruncated = tableListTruncated
	this.IndexListTruncated = indexListTruncated
	return &this
}

// NewPostgresqlSpaceSummaryWithDefaults instantiates a new PostgresqlSpaceSummary object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlSpaceSummaryWithDefaults() *PostgresqlSpaceSummary {
	this := PostgresqlSpaceSummary{}
	return &this
}

// GetDatabaseName returns the DatabaseName field value if set, zero value otherwise.
func (o *PostgresqlSpaceSummary) GetDatabaseName() string {
	if o == nil || o.DatabaseName == nil {
		var ret string
		return ret
	}
	return *o.DatabaseName
}

// GetDatabaseNameOk returns a tuple with the DatabaseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSpaceSummary) GetDatabaseNameOk() (*string, bool) {
	if o == nil || o.DatabaseName == nil {
		return nil, false
	}
	return o.DatabaseName, true
}

// HasDatabaseName returns a boolean if a field has been set.
func (o *PostgresqlSpaceSummary) HasDatabaseName() bool {
	return o != nil && o.DatabaseName != nil
}

// SetDatabaseName gets a reference to the given string and assigns it to the DatabaseName field.
func (o *PostgresqlSpaceSummary) SetDatabaseName(v string) {
	o.DatabaseName = &v
}

// GetDatabaseSizeBytes returns the DatabaseSizeBytes field value if set, zero value otherwise.
func (o *PostgresqlSpaceSummary) GetDatabaseSizeBytes() int64 {
	if o == nil || o.DatabaseSizeBytes == nil {
		var ret int64
		return ret
	}
	return *o.DatabaseSizeBytes
}

// GetDatabaseSizeBytesOk returns a tuple with the DatabaseSizeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSpaceSummary) GetDatabaseSizeBytesOk() (*int64, bool) {
	if o == nil || o.DatabaseSizeBytes == nil {
		return nil, false
	}
	return o.DatabaseSizeBytes, true
}

// HasDatabaseSizeBytes returns a boolean if a field has been set.
func (o *PostgresqlSpaceSummary) HasDatabaseSizeBytes() bool {
	return o != nil && o.DatabaseSizeBytes != nil
}

// SetDatabaseSizeBytes gets a reference to the given int64 and assigns it to the DatabaseSizeBytes field.
func (o *PostgresqlSpaceSummary) SetDatabaseSizeBytes(v int64) {
	o.DatabaseSizeBytes = &v
}

// GetTableCount returns the TableCount field value.
func (o *PostgresqlSpaceSummary) GetTableCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TableCount
}

// GetTableCountOk returns a tuple with the TableCount field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSpaceSummary) GetTableCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TableCount, true
}

// SetTableCount sets field value.
func (o *PostgresqlSpaceSummary) SetTableCount(v int64) {
	o.TableCount = v
}

// GetIndexCount returns the IndexCount field value.
func (o *PostgresqlSpaceSummary) GetIndexCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.IndexCount
}

// GetIndexCountOk returns a tuple with the IndexCount field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSpaceSummary) GetIndexCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IndexCount, true
}

// SetIndexCount sets field value.
func (o *PostgresqlSpaceSummary) SetIndexCount(v int64) {
	o.IndexCount = v
}

// GetToastRelationCount returns the ToastRelationCount field value.
func (o *PostgresqlSpaceSummary) GetToastRelationCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ToastRelationCount
}

// GetToastRelationCountOk returns a tuple with the ToastRelationCount field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSpaceSummary) GetToastRelationCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToastRelationCount, true
}

// SetToastRelationCount sets field value.
func (o *PostgresqlSpaceSummary) SetToastRelationCount(v int64) {
	o.ToastRelationCount = v
}

// GetLargestTableBytes returns the LargestTableBytes field value if set, zero value otherwise.
func (o *PostgresqlSpaceSummary) GetLargestTableBytes() int64 {
	if o == nil || o.LargestTableBytes == nil {
		var ret int64
		return ret
	}
	return *o.LargestTableBytes
}

// GetLargestTableBytesOk returns a tuple with the LargestTableBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSpaceSummary) GetLargestTableBytesOk() (*int64, bool) {
	if o == nil || o.LargestTableBytes == nil {
		return nil, false
	}
	return o.LargestTableBytes, true
}

// HasLargestTableBytes returns a boolean if a field has been set.
func (o *PostgresqlSpaceSummary) HasLargestTableBytes() bool {
	return o != nil && o.LargestTableBytes != nil
}

// SetLargestTableBytes gets a reference to the given int64 and assigns it to the LargestTableBytes field.
func (o *PostgresqlSpaceSummary) SetLargestTableBytes(v int64) {
	o.LargestTableBytes = &v
}

// GetLargestIndexBytes returns the LargestIndexBytes field value if set, zero value otherwise.
func (o *PostgresqlSpaceSummary) GetLargestIndexBytes() int64 {
	if o == nil || o.LargestIndexBytes == nil {
		var ret int64
		return ret
	}
	return *o.LargestIndexBytes
}

// GetLargestIndexBytesOk returns a tuple with the LargestIndexBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSpaceSummary) GetLargestIndexBytesOk() (*int64, bool) {
	if o == nil || o.LargestIndexBytes == nil {
		return nil, false
	}
	return o.LargestIndexBytes, true
}

// HasLargestIndexBytes returns a boolean if a field has been set.
func (o *PostgresqlSpaceSummary) HasLargestIndexBytes() bool {
	return o != nil && o.LargestIndexBytes != nil
}

// SetLargestIndexBytes gets a reference to the given int64 and assigns it to the LargestIndexBytes field.
func (o *PostgresqlSpaceSummary) SetLargestIndexBytes(v int64) {
	o.LargestIndexBytes = &v
}

// GetTableListTruncated returns the TableListTruncated field value.
func (o *PostgresqlSpaceSummary) GetTableListTruncated() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.TableListTruncated
}

// GetTableListTruncatedOk returns a tuple with the TableListTruncated field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSpaceSummary) GetTableListTruncatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TableListTruncated, true
}

// SetTableListTruncated sets field value.
func (o *PostgresqlSpaceSummary) SetTableListTruncated(v bool) {
	o.TableListTruncated = v
}

// GetIndexListTruncated returns the IndexListTruncated field value.
func (o *PostgresqlSpaceSummary) GetIndexListTruncated() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IndexListTruncated
}

// GetIndexListTruncatedOk returns a tuple with the IndexListTruncated field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSpaceSummary) GetIndexListTruncatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IndexListTruncated, true
}

// SetIndexListTruncated sets field value.
func (o *PostgresqlSpaceSummary) SetIndexListTruncated(v bool) {
	o.IndexListTruncated = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlSpaceSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.DatabaseName != nil {
		toSerialize["databaseName"] = o.DatabaseName
	}
	if o.DatabaseSizeBytes != nil {
		toSerialize["databaseSizeBytes"] = o.DatabaseSizeBytes
	}
	toSerialize["tableCount"] = o.TableCount
	toSerialize["indexCount"] = o.IndexCount
	toSerialize["toastRelationCount"] = o.ToastRelationCount
	if o.LargestTableBytes != nil {
		toSerialize["largestTableBytes"] = o.LargestTableBytes
	}
	if o.LargestIndexBytes != nil {
		toSerialize["largestIndexBytes"] = o.LargestIndexBytes
	}
	toSerialize["tableListTruncated"] = o.TableListTruncated
	toSerialize["indexListTruncated"] = o.IndexListTruncated

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostgresqlSpaceSummary) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DatabaseName       *string `json:"databaseName,omitempty"`
		DatabaseSizeBytes  *int64  `json:"databaseSizeBytes,omitempty"`
		TableCount         *int64  `json:"tableCount"`
		IndexCount         *int64  `json:"indexCount"`
		ToastRelationCount *int64  `json:"toastRelationCount"`
		LargestTableBytes  *int64  `json:"largestTableBytes,omitempty"`
		LargestIndexBytes  *int64  `json:"largestIndexBytes,omitempty"`
		TableListTruncated *bool   `json:"tableListTruncated"`
		IndexListTruncated *bool   `json:"indexListTruncated"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.TableCount == nil {
		return fmt.Errorf("required field tableCount missing")
	}
	if all.IndexCount == nil {
		return fmt.Errorf("required field indexCount missing")
	}
	if all.ToastRelationCount == nil {
		return fmt.Errorf("required field toastRelationCount missing")
	}
	if all.TableListTruncated == nil {
		return fmt.Errorf("required field tableListTruncated missing")
	}
	if all.IndexListTruncated == nil {
		return fmt.Errorf("required field indexListTruncated missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"databaseName", "databaseSizeBytes", "tableCount", "indexCount", "toastRelationCount", "largestTableBytes", "largestIndexBytes", "tableListTruncated", "indexListTruncated"})
	} else {
		return err
	}
	o.DatabaseName = all.DatabaseName
	o.DatabaseSizeBytes = all.DatabaseSizeBytes
	o.TableCount = *all.TableCount
	o.IndexCount = *all.IndexCount
	o.ToastRelationCount = *all.ToastRelationCount
	o.LargestTableBytes = all.LargestTableBytes
	o.LargestIndexBytes = all.LargestIndexBytes
	o.TableListTruncated = *all.TableListTruncated
	o.IndexListTruncated = *all.IndexListTruncated

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
