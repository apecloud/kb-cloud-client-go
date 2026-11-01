// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type ClickhouseSpaceSummary struct {
	// Number of databases with active parts.
	DatabaseCount *int64 `json:"databaseCount,omitempty"`
	// Number of tables with active parts.
	TableCount *int64 `json:"tableCount,omitempty"`
	// Physical bytes_on_disk across all hosts.
	TotalBytes *int64 `json:"totalBytes,omitempty"`
	// Sum of data_compressed_bytes.
	DataCompressedBytes *int64 `json:"dataCompressedBytes,omitempty"`
	// Sum of data_uncompressed_bytes.
	DataUncompressedBytes *int64 `json:"dataUncompressedBytes,omitempty"`
	// data_compressed_bytes / data_uncompressed_bytes. Lower is better; zero when uncompressed size is unknown.
	CompressionRatio *float64 `json:"compressionRatio,omitempty"`
	// Sum of active part rows across all hosts.
	TotalRows *int64 `json:"totalRows,omitempty"`
	// Number of active parts across all hosts.
	ActiveParts *int64 `json:"activeParts,omitempty"`
	// Sum of per-table partition counts. Replicated partitions can be counted more than once.
	PartitionCount *int64 `json:"partitionCount,omitempty"`
	// Tables whose active part count reached parts_to_delay_insert.
	WarningTables *int64 `json:"warningTables,omitempty"`
	// Tables whose active part count reached parts_to_throw_insert.
	CriticalTables *int64 `json:"criticalTables,omitempty"`
	// Effective parts_to_delay_insert from system.merge_tree_settings (default 150).
	PartsToDelayInsert *int64 `json:"partsToDelayInsert,omitempty"`
	// Effective parts_to_throw_insert from system.merge_tree_settings (default 300).
	PartsToThrowInsert *int64 `json:"partsToThrowInsert,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClickhouseSpaceSummary instantiates a new ClickhouseSpaceSummary object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClickhouseSpaceSummary() *ClickhouseSpaceSummary {
	this := ClickhouseSpaceSummary{}
	return &this
}

// NewClickhouseSpaceSummaryWithDefaults instantiates a new ClickhouseSpaceSummary object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClickhouseSpaceSummaryWithDefaults() *ClickhouseSpaceSummary {
	this := ClickhouseSpaceSummary{}
	return &this
}

// GetDatabaseCount returns the DatabaseCount field value if set, zero value otherwise.
func (o *ClickhouseSpaceSummary) GetDatabaseCount() int64 {
	if o == nil || o.DatabaseCount == nil {
		var ret int64
		return ret
	}
	return *o.DatabaseCount
}

// GetDatabaseCountOk returns a tuple with the DatabaseCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSpaceSummary) GetDatabaseCountOk() (*int64, bool) {
	if o == nil || o.DatabaseCount == nil {
		return nil, false
	}
	return o.DatabaseCount, true
}

// HasDatabaseCount returns a boolean if a field has been set.
func (o *ClickhouseSpaceSummary) HasDatabaseCount() bool {
	return o != nil && o.DatabaseCount != nil
}

// SetDatabaseCount gets a reference to the given int64 and assigns it to the DatabaseCount field.
func (o *ClickhouseSpaceSummary) SetDatabaseCount(v int64) {
	o.DatabaseCount = &v
}

// GetTableCount returns the TableCount field value if set, zero value otherwise.
func (o *ClickhouseSpaceSummary) GetTableCount() int64 {
	if o == nil || o.TableCount == nil {
		var ret int64
		return ret
	}
	return *o.TableCount
}

// GetTableCountOk returns a tuple with the TableCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSpaceSummary) GetTableCountOk() (*int64, bool) {
	if o == nil || o.TableCount == nil {
		return nil, false
	}
	return o.TableCount, true
}

// HasTableCount returns a boolean if a field has been set.
func (o *ClickhouseSpaceSummary) HasTableCount() bool {
	return o != nil && o.TableCount != nil
}

// SetTableCount gets a reference to the given int64 and assigns it to the TableCount field.
func (o *ClickhouseSpaceSummary) SetTableCount(v int64) {
	o.TableCount = &v
}

// GetTotalBytes returns the TotalBytes field value if set, zero value otherwise.
func (o *ClickhouseSpaceSummary) GetTotalBytes() int64 {
	if o == nil || o.TotalBytes == nil {
		var ret int64
		return ret
	}
	return *o.TotalBytes
}

// GetTotalBytesOk returns a tuple with the TotalBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSpaceSummary) GetTotalBytesOk() (*int64, bool) {
	if o == nil || o.TotalBytes == nil {
		return nil, false
	}
	return o.TotalBytes, true
}

// HasTotalBytes returns a boolean if a field has been set.
func (o *ClickhouseSpaceSummary) HasTotalBytes() bool {
	return o != nil && o.TotalBytes != nil
}

// SetTotalBytes gets a reference to the given int64 and assigns it to the TotalBytes field.
func (o *ClickhouseSpaceSummary) SetTotalBytes(v int64) {
	o.TotalBytes = &v
}

// GetDataCompressedBytes returns the DataCompressedBytes field value if set, zero value otherwise.
func (o *ClickhouseSpaceSummary) GetDataCompressedBytes() int64 {
	if o == nil || o.DataCompressedBytes == nil {
		var ret int64
		return ret
	}
	return *o.DataCompressedBytes
}

// GetDataCompressedBytesOk returns a tuple with the DataCompressedBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSpaceSummary) GetDataCompressedBytesOk() (*int64, bool) {
	if o == nil || o.DataCompressedBytes == nil {
		return nil, false
	}
	return o.DataCompressedBytes, true
}

// HasDataCompressedBytes returns a boolean if a field has been set.
func (o *ClickhouseSpaceSummary) HasDataCompressedBytes() bool {
	return o != nil && o.DataCompressedBytes != nil
}

// SetDataCompressedBytes gets a reference to the given int64 and assigns it to the DataCompressedBytes field.
func (o *ClickhouseSpaceSummary) SetDataCompressedBytes(v int64) {
	o.DataCompressedBytes = &v
}

// GetDataUncompressedBytes returns the DataUncompressedBytes field value if set, zero value otherwise.
func (o *ClickhouseSpaceSummary) GetDataUncompressedBytes() int64 {
	if o == nil || o.DataUncompressedBytes == nil {
		var ret int64
		return ret
	}
	return *o.DataUncompressedBytes
}

// GetDataUncompressedBytesOk returns a tuple with the DataUncompressedBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSpaceSummary) GetDataUncompressedBytesOk() (*int64, bool) {
	if o == nil || o.DataUncompressedBytes == nil {
		return nil, false
	}
	return o.DataUncompressedBytes, true
}

// HasDataUncompressedBytes returns a boolean if a field has been set.
func (o *ClickhouseSpaceSummary) HasDataUncompressedBytes() bool {
	return o != nil && o.DataUncompressedBytes != nil
}

// SetDataUncompressedBytes gets a reference to the given int64 and assigns it to the DataUncompressedBytes field.
func (o *ClickhouseSpaceSummary) SetDataUncompressedBytes(v int64) {
	o.DataUncompressedBytes = &v
}

// GetCompressionRatio returns the CompressionRatio field value if set, zero value otherwise.
func (o *ClickhouseSpaceSummary) GetCompressionRatio() float64 {
	if o == nil || o.CompressionRatio == nil {
		var ret float64
		return ret
	}
	return *o.CompressionRatio
}

// GetCompressionRatioOk returns a tuple with the CompressionRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSpaceSummary) GetCompressionRatioOk() (*float64, bool) {
	if o == nil || o.CompressionRatio == nil {
		return nil, false
	}
	return o.CompressionRatio, true
}

// HasCompressionRatio returns a boolean if a field has been set.
func (o *ClickhouseSpaceSummary) HasCompressionRatio() bool {
	return o != nil && o.CompressionRatio != nil
}

// SetCompressionRatio gets a reference to the given float64 and assigns it to the CompressionRatio field.
func (o *ClickhouseSpaceSummary) SetCompressionRatio(v float64) {
	o.CompressionRatio = &v
}

// GetTotalRows returns the TotalRows field value if set, zero value otherwise.
func (o *ClickhouseSpaceSummary) GetTotalRows() int64 {
	if o == nil || o.TotalRows == nil {
		var ret int64
		return ret
	}
	return *o.TotalRows
}

// GetTotalRowsOk returns a tuple with the TotalRows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSpaceSummary) GetTotalRowsOk() (*int64, bool) {
	if o == nil || o.TotalRows == nil {
		return nil, false
	}
	return o.TotalRows, true
}

// HasTotalRows returns a boolean if a field has been set.
func (o *ClickhouseSpaceSummary) HasTotalRows() bool {
	return o != nil && o.TotalRows != nil
}

// SetTotalRows gets a reference to the given int64 and assigns it to the TotalRows field.
func (o *ClickhouseSpaceSummary) SetTotalRows(v int64) {
	o.TotalRows = &v
}

// GetActiveParts returns the ActiveParts field value if set, zero value otherwise.
func (o *ClickhouseSpaceSummary) GetActiveParts() int64 {
	if o == nil || o.ActiveParts == nil {
		var ret int64
		return ret
	}
	return *o.ActiveParts
}

// GetActivePartsOk returns a tuple with the ActiveParts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSpaceSummary) GetActivePartsOk() (*int64, bool) {
	if o == nil || o.ActiveParts == nil {
		return nil, false
	}
	return o.ActiveParts, true
}

// HasActiveParts returns a boolean if a field has been set.
func (o *ClickhouseSpaceSummary) HasActiveParts() bool {
	return o != nil && o.ActiveParts != nil
}

// SetActiveParts gets a reference to the given int64 and assigns it to the ActiveParts field.
func (o *ClickhouseSpaceSummary) SetActiveParts(v int64) {
	o.ActiveParts = &v
}

// GetPartitionCount returns the PartitionCount field value if set, zero value otherwise.
func (o *ClickhouseSpaceSummary) GetPartitionCount() int64 {
	if o == nil || o.PartitionCount == nil {
		var ret int64
		return ret
	}
	return *o.PartitionCount
}

// GetPartitionCountOk returns a tuple with the PartitionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSpaceSummary) GetPartitionCountOk() (*int64, bool) {
	if o == nil || o.PartitionCount == nil {
		return nil, false
	}
	return o.PartitionCount, true
}

// HasPartitionCount returns a boolean if a field has been set.
func (o *ClickhouseSpaceSummary) HasPartitionCount() bool {
	return o != nil && o.PartitionCount != nil
}

// SetPartitionCount gets a reference to the given int64 and assigns it to the PartitionCount field.
func (o *ClickhouseSpaceSummary) SetPartitionCount(v int64) {
	o.PartitionCount = &v
}

// GetWarningTables returns the WarningTables field value if set, zero value otherwise.
func (o *ClickhouseSpaceSummary) GetWarningTables() int64 {
	if o == nil || o.WarningTables == nil {
		var ret int64
		return ret
	}
	return *o.WarningTables
}

// GetWarningTablesOk returns a tuple with the WarningTables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSpaceSummary) GetWarningTablesOk() (*int64, bool) {
	if o == nil || o.WarningTables == nil {
		return nil, false
	}
	return o.WarningTables, true
}

// HasWarningTables returns a boolean if a field has been set.
func (o *ClickhouseSpaceSummary) HasWarningTables() bool {
	return o != nil && o.WarningTables != nil
}

// SetWarningTables gets a reference to the given int64 and assigns it to the WarningTables field.
func (o *ClickhouseSpaceSummary) SetWarningTables(v int64) {
	o.WarningTables = &v
}

// GetCriticalTables returns the CriticalTables field value if set, zero value otherwise.
func (o *ClickhouseSpaceSummary) GetCriticalTables() int64 {
	if o == nil || o.CriticalTables == nil {
		var ret int64
		return ret
	}
	return *o.CriticalTables
}

// GetCriticalTablesOk returns a tuple with the CriticalTables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSpaceSummary) GetCriticalTablesOk() (*int64, bool) {
	if o == nil || o.CriticalTables == nil {
		return nil, false
	}
	return o.CriticalTables, true
}

// HasCriticalTables returns a boolean if a field has been set.
func (o *ClickhouseSpaceSummary) HasCriticalTables() bool {
	return o != nil && o.CriticalTables != nil
}

// SetCriticalTables gets a reference to the given int64 and assigns it to the CriticalTables field.
func (o *ClickhouseSpaceSummary) SetCriticalTables(v int64) {
	o.CriticalTables = &v
}

// GetPartsToDelayInsert returns the PartsToDelayInsert field value if set, zero value otherwise.
func (o *ClickhouseSpaceSummary) GetPartsToDelayInsert() int64 {
	if o == nil || o.PartsToDelayInsert == nil {
		var ret int64
		return ret
	}
	return *o.PartsToDelayInsert
}

// GetPartsToDelayInsertOk returns a tuple with the PartsToDelayInsert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSpaceSummary) GetPartsToDelayInsertOk() (*int64, bool) {
	if o == nil || o.PartsToDelayInsert == nil {
		return nil, false
	}
	return o.PartsToDelayInsert, true
}

// HasPartsToDelayInsert returns a boolean if a field has been set.
func (o *ClickhouseSpaceSummary) HasPartsToDelayInsert() bool {
	return o != nil && o.PartsToDelayInsert != nil
}

// SetPartsToDelayInsert gets a reference to the given int64 and assigns it to the PartsToDelayInsert field.
func (o *ClickhouseSpaceSummary) SetPartsToDelayInsert(v int64) {
	o.PartsToDelayInsert = &v
}

// GetPartsToThrowInsert returns the PartsToThrowInsert field value if set, zero value otherwise.
func (o *ClickhouseSpaceSummary) GetPartsToThrowInsert() int64 {
	if o == nil || o.PartsToThrowInsert == nil {
		var ret int64
		return ret
	}
	return *o.PartsToThrowInsert
}

// GetPartsToThrowInsertOk returns a tuple with the PartsToThrowInsert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSpaceSummary) GetPartsToThrowInsertOk() (*int64, bool) {
	if o == nil || o.PartsToThrowInsert == nil {
		return nil, false
	}
	return o.PartsToThrowInsert, true
}

// HasPartsToThrowInsert returns a boolean if a field has been set.
func (o *ClickhouseSpaceSummary) HasPartsToThrowInsert() bool {
	return o != nil && o.PartsToThrowInsert != nil
}

// SetPartsToThrowInsert gets a reference to the given int64 and assigns it to the PartsToThrowInsert field.
func (o *ClickhouseSpaceSummary) SetPartsToThrowInsert(v int64) {
	o.PartsToThrowInsert = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClickhouseSpaceSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.DatabaseCount != nil {
		toSerialize["databaseCount"] = o.DatabaseCount
	}
	if o.TableCount != nil {
		toSerialize["tableCount"] = o.TableCount
	}
	if o.TotalBytes != nil {
		toSerialize["totalBytes"] = o.TotalBytes
	}
	if o.DataCompressedBytes != nil {
		toSerialize["dataCompressedBytes"] = o.DataCompressedBytes
	}
	if o.DataUncompressedBytes != nil {
		toSerialize["dataUncompressedBytes"] = o.DataUncompressedBytes
	}
	if o.CompressionRatio != nil {
		toSerialize["compressionRatio"] = o.CompressionRatio
	}
	if o.TotalRows != nil {
		toSerialize["totalRows"] = o.TotalRows
	}
	if o.ActiveParts != nil {
		toSerialize["activeParts"] = o.ActiveParts
	}
	if o.PartitionCount != nil {
		toSerialize["partitionCount"] = o.PartitionCount
	}
	if o.WarningTables != nil {
		toSerialize["warningTables"] = o.WarningTables
	}
	if o.CriticalTables != nil {
		toSerialize["criticalTables"] = o.CriticalTables
	}
	if o.PartsToDelayInsert != nil {
		toSerialize["partsToDelayInsert"] = o.PartsToDelayInsert
	}
	if o.PartsToThrowInsert != nil {
		toSerialize["partsToThrowInsert"] = o.PartsToThrowInsert
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClickhouseSpaceSummary) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DatabaseCount         *int64   `json:"databaseCount,omitempty"`
		TableCount            *int64   `json:"tableCount,omitempty"`
		TotalBytes            *int64   `json:"totalBytes,omitempty"`
		DataCompressedBytes   *int64   `json:"dataCompressedBytes,omitempty"`
		DataUncompressedBytes *int64   `json:"dataUncompressedBytes,omitempty"`
		CompressionRatio      *float64 `json:"compressionRatio,omitempty"`
		TotalRows             *int64   `json:"totalRows,omitempty"`
		ActiveParts           *int64   `json:"activeParts,omitempty"`
		PartitionCount        *int64   `json:"partitionCount,omitempty"`
		WarningTables         *int64   `json:"warningTables,omitempty"`
		CriticalTables        *int64   `json:"criticalTables,omitempty"`
		PartsToDelayInsert    *int64   `json:"partsToDelayInsert,omitempty"`
		PartsToThrowInsert    *int64   `json:"partsToThrowInsert,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"databaseCount", "tableCount", "totalBytes", "dataCompressedBytes", "dataUncompressedBytes", "compressionRatio", "totalRows", "activeParts", "partitionCount", "warningTables", "criticalTables", "partsToDelayInsert", "partsToThrowInsert"})
	} else {
		return err
	}
	o.DatabaseCount = all.DatabaseCount
	o.TableCount = all.TableCount
	o.TotalBytes = all.TotalBytes
	o.DataCompressedBytes = all.DataCompressedBytes
	o.DataUncompressedBytes = all.DataUncompressedBytes
	o.CompressionRatio = all.CompressionRatio
	o.TotalRows = all.TotalRows
	o.ActiveParts = all.ActiveParts
	o.PartitionCount = all.PartitionCount
	o.WarningTables = all.WarningTables
	o.CriticalTables = all.CriticalTables
	o.PartsToDelayInsert = all.PartsToDelayInsert
	o.PartsToThrowInsert = all.PartsToThrowInsert

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
