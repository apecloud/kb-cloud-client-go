// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type ClickhousePartitionSpace struct {
	// Database name.
	Database *string `json:"database,omitempty"`
	// Table name.
	Table *string `json:"table,omitempty"`
	// partition_id from system.parts.
	PartitionId *string `json:"partitionId,omitempty"`
	// Human-readable partition expression.
	Partition *string `json:"partition,omitempty"`
	// Sum of rows across hosts.
	Rows *int64 `json:"rows,omitempty"`
	// Physical bytes_on_disk across hosts.
	BytesOnDisk *int64 `json:"bytesOnDisk,omitempty"`
	// Sum of data_compressed_bytes.
	DataCompressedBytes *int64 `json:"dataCompressedBytes,omitempty"`
	// Sum of data_uncompressed_bytes.
	DataUncompressedBytes *int64 `json:"dataUncompressedBytes,omitempty"`
	// Active part count across hosts.
	ActiveParts *int64 `json:"activeParts,omitempty"`
	// Number of hosts that own this partition.
	Hosts *int64 `json:"hosts,omitempty"`
	// Minimum min_date in the partition.
	MinDate *string `json:"minDate,omitempty"`
	// Maximum max_date in the partition.
	MaxDate *string `json:"maxDate,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClickhousePartitionSpace instantiates a new ClickhousePartitionSpace object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClickhousePartitionSpace() *ClickhousePartitionSpace {
	this := ClickhousePartitionSpace{}
	return &this
}

// NewClickhousePartitionSpaceWithDefaults instantiates a new ClickhousePartitionSpace object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClickhousePartitionSpaceWithDefaults() *ClickhousePartitionSpace {
	this := ClickhousePartitionSpace{}
	return &this
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *ClickhousePartitionSpace) GetDatabase() string {
	if o == nil || o.Database == nil {
		var ret string
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhousePartitionSpace) GetDatabaseOk() (*string, bool) {
	if o == nil || o.Database == nil {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *ClickhousePartitionSpace) HasDatabase() bool {
	return o != nil && o.Database != nil
}

// SetDatabase gets a reference to the given string and assigns it to the Database field.
func (o *ClickhousePartitionSpace) SetDatabase(v string) {
	o.Database = &v
}

// GetTable returns the Table field value if set, zero value otherwise.
func (o *ClickhousePartitionSpace) GetTable() string {
	if o == nil || o.Table == nil {
		var ret string
		return ret
	}
	return *o.Table
}

// GetTableOk returns a tuple with the Table field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhousePartitionSpace) GetTableOk() (*string, bool) {
	if o == nil || o.Table == nil {
		return nil, false
	}
	return o.Table, true
}

// HasTable returns a boolean if a field has been set.
func (o *ClickhousePartitionSpace) HasTable() bool {
	return o != nil && o.Table != nil
}

// SetTable gets a reference to the given string and assigns it to the Table field.
func (o *ClickhousePartitionSpace) SetTable(v string) {
	o.Table = &v
}

// GetPartitionId returns the PartitionId field value if set, zero value otherwise.
func (o *ClickhousePartitionSpace) GetPartitionId() string {
	if o == nil || o.PartitionId == nil {
		var ret string
		return ret
	}
	return *o.PartitionId
}

// GetPartitionIdOk returns a tuple with the PartitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhousePartitionSpace) GetPartitionIdOk() (*string, bool) {
	if o == nil || o.PartitionId == nil {
		return nil, false
	}
	return o.PartitionId, true
}

// HasPartitionId returns a boolean if a field has been set.
func (o *ClickhousePartitionSpace) HasPartitionId() bool {
	return o != nil && o.PartitionId != nil
}

// SetPartitionId gets a reference to the given string and assigns it to the PartitionId field.
func (o *ClickhousePartitionSpace) SetPartitionId(v string) {
	o.PartitionId = &v
}

// GetPartition returns the Partition field value if set, zero value otherwise.
func (o *ClickhousePartitionSpace) GetPartition() string {
	if o == nil || o.Partition == nil {
		var ret string
		return ret
	}
	return *o.Partition
}

// GetPartitionOk returns a tuple with the Partition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhousePartitionSpace) GetPartitionOk() (*string, bool) {
	if o == nil || o.Partition == nil {
		return nil, false
	}
	return o.Partition, true
}

// HasPartition returns a boolean if a field has been set.
func (o *ClickhousePartitionSpace) HasPartition() bool {
	return o != nil && o.Partition != nil
}

// SetPartition gets a reference to the given string and assigns it to the Partition field.
func (o *ClickhousePartitionSpace) SetPartition(v string) {
	o.Partition = &v
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *ClickhousePartitionSpace) GetRows() int64 {
	if o == nil || o.Rows == nil {
		var ret int64
		return ret
	}
	return *o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhousePartitionSpace) GetRowsOk() (*int64, bool) {
	if o == nil || o.Rows == nil {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *ClickhousePartitionSpace) HasRows() bool {
	return o != nil && o.Rows != nil
}

// SetRows gets a reference to the given int64 and assigns it to the Rows field.
func (o *ClickhousePartitionSpace) SetRows(v int64) {
	o.Rows = &v
}

// GetBytesOnDisk returns the BytesOnDisk field value if set, zero value otherwise.
func (o *ClickhousePartitionSpace) GetBytesOnDisk() int64 {
	if o == nil || o.BytesOnDisk == nil {
		var ret int64
		return ret
	}
	return *o.BytesOnDisk
}

// GetBytesOnDiskOk returns a tuple with the BytesOnDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhousePartitionSpace) GetBytesOnDiskOk() (*int64, bool) {
	if o == nil || o.BytesOnDisk == nil {
		return nil, false
	}
	return o.BytesOnDisk, true
}

// HasBytesOnDisk returns a boolean if a field has been set.
func (o *ClickhousePartitionSpace) HasBytesOnDisk() bool {
	return o != nil && o.BytesOnDisk != nil
}

// SetBytesOnDisk gets a reference to the given int64 and assigns it to the BytesOnDisk field.
func (o *ClickhousePartitionSpace) SetBytesOnDisk(v int64) {
	o.BytesOnDisk = &v
}

// GetDataCompressedBytes returns the DataCompressedBytes field value if set, zero value otherwise.
func (o *ClickhousePartitionSpace) GetDataCompressedBytes() int64 {
	if o == nil || o.DataCompressedBytes == nil {
		var ret int64
		return ret
	}
	return *o.DataCompressedBytes
}

// GetDataCompressedBytesOk returns a tuple with the DataCompressedBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhousePartitionSpace) GetDataCompressedBytesOk() (*int64, bool) {
	if o == nil || o.DataCompressedBytes == nil {
		return nil, false
	}
	return o.DataCompressedBytes, true
}

// HasDataCompressedBytes returns a boolean if a field has been set.
func (o *ClickhousePartitionSpace) HasDataCompressedBytes() bool {
	return o != nil && o.DataCompressedBytes != nil
}

// SetDataCompressedBytes gets a reference to the given int64 and assigns it to the DataCompressedBytes field.
func (o *ClickhousePartitionSpace) SetDataCompressedBytes(v int64) {
	o.DataCompressedBytes = &v
}

// GetDataUncompressedBytes returns the DataUncompressedBytes field value if set, zero value otherwise.
func (o *ClickhousePartitionSpace) GetDataUncompressedBytes() int64 {
	if o == nil || o.DataUncompressedBytes == nil {
		var ret int64
		return ret
	}
	return *o.DataUncompressedBytes
}

// GetDataUncompressedBytesOk returns a tuple with the DataUncompressedBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhousePartitionSpace) GetDataUncompressedBytesOk() (*int64, bool) {
	if o == nil || o.DataUncompressedBytes == nil {
		return nil, false
	}
	return o.DataUncompressedBytes, true
}

// HasDataUncompressedBytes returns a boolean if a field has been set.
func (o *ClickhousePartitionSpace) HasDataUncompressedBytes() bool {
	return o != nil && o.DataUncompressedBytes != nil
}

// SetDataUncompressedBytes gets a reference to the given int64 and assigns it to the DataUncompressedBytes field.
func (o *ClickhousePartitionSpace) SetDataUncompressedBytes(v int64) {
	o.DataUncompressedBytes = &v
}

// GetActiveParts returns the ActiveParts field value if set, zero value otherwise.
func (o *ClickhousePartitionSpace) GetActiveParts() int64 {
	if o == nil || o.ActiveParts == nil {
		var ret int64
		return ret
	}
	return *o.ActiveParts
}

// GetActivePartsOk returns a tuple with the ActiveParts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhousePartitionSpace) GetActivePartsOk() (*int64, bool) {
	if o == nil || o.ActiveParts == nil {
		return nil, false
	}
	return o.ActiveParts, true
}

// HasActiveParts returns a boolean if a field has been set.
func (o *ClickhousePartitionSpace) HasActiveParts() bool {
	return o != nil && o.ActiveParts != nil
}

// SetActiveParts gets a reference to the given int64 and assigns it to the ActiveParts field.
func (o *ClickhousePartitionSpace) SetActiveParts(v int64) {
	o.ActiveParts = &v
}

// GetHosts returns the Hosts field value if set, zero value otherwise.
func (o *ClickhousePartitionSpace) GetHosts() int64 {
	if o == nil || o.Hosts == nil {
		var ret int64
		return ret
	}
	return *o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhousePartitionSpace) GetHostsOk() (*int64, bool) {
	if o == nil || o.Hosts == nil {
		return nil, false
	}
	return o.Hosts, true
}

// HasHosts returns a boolean if a field has been set.
func (o *ClickhousePartitionSpace) HasHosts() bool {
	return o != nil && o.Hosts != nil
}

// SetHosts gets a reference to the given int64 and assigns it to the Hosts field.
func (o *ClickhousePartitionSpace) SetHosts(v int64) {
	o.Hosts = &v
}

// GetMinDate returns the MinDate field value if set, zero value otherwise.
func (o *ClickhousePartitionSpace) GetMinDate() string {
	if o == nil || o.MinDate == nil {
		var ret string
		return ret
	}
	return *o.MinDate
}

// GetMinDateOk returns a tuple with the MinDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhousePartitionSpace) GetMinDateOk() (*string, bool) {
	if o == nil || o.MinDate == nil {
		return nil, false
	}
	return o.MinDate, true
}

// HasMinDate returns a boolean if a field has been set.
func (o *ClickhousePartitionSpace) HasMinDate() bool {
	return o != nil && o.MinDate != nil
}

// SetMinDate gets a reference to the given string and assigns it to the MinDate field.
func (o *ClickhousePartitionSpace) SetMinDate(v string) {
	o.MinDate = &v
}

// GetMaxDate returns the MaxDate field value if set, zero value otherwise.
func (o *ClickhousePartitionSpace) GetMaxDate() string {
	if o == nil || o.MaxDate == nil {
		var ret string
		return ret
	}
	return *o.MaxDate
}

// GetMaxDateOk returns a tuple with the MaxDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhousePartitionSpace) GetMaxDateOk() (*string, bool) {
	if o == nil || o.MaxDate == nil {
		return nil, false
	}
	return o.MaxDate, true
}

// HasMaxDate returns a boolean if a field has been set.
func (o *ClickhousePartitionSpace) HasMaxDate() bool {
	return o != nil && o.MaxDate != nil
}

// SetMaxDate gets a reference to the given string and assigns it to the MaxDate field.
func (o *ClickhousePartitionSpace) SetMaxDate(v string) {
	o.MaxDate = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClickhousePartitionSpace) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Database != nil {
		toSerialize["database"] = o.Database
	}
	if o.Table != nil {
		toSerialize["table"] = o.Table
	}
	if o.PartitionId != nil {
		toSerialize["partitionId"] = o.PartitionId
	}
	if o.Partition != nil {
		toSerialize["partition"] = o.Partition
	}
	if o.Rows != nil {
		toSerialize["rows"] = o.Rows
	}
	if o.BytesOnDisk != nil {
		toSerialize["bytesOnDisk"] = o.BytesOnDisk
	}
	if o.DataCompressedBytes != nil {
		toSerialize["dataCompressedBytes"] = o.DataCompressedBytes
	}
	if o.DataUncompressedBytes != nil {
		toSerialize["dataUncompressedBytes"] = o.DataUncompressedBytes
	}
	if o.ActiveParts != nil {
		toSerialize["activeParts"] = o.ActiveParts
	}
	if o.Hosts != nil {
		toSerialize["hosts"] = o.Hosts
	}
	if o.MinDate != nil {
		toSerialize["minDate"] = o.MinDate
	}
	if o.MaxDate != nil {
		toSerialize["maxDate"] = o.MaxDate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClickhousePartitionSpace) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Database              *string `json:"database,omitempty"`
		Table                 *string `json:"table,omitempty"`
		PartitionId           *string `json:"partitionId,omitempty"`
		Partition             *string `json:"partition,omitempty"`
		Rows                  *int64  `json:"rows,omitempty"`
		BytesOnDisk           *int64  `json:"bytesOnDisk,omitempty"`
		DataCompressedBytes   *int64  `json:"dataCompressedBytes,omitempty"`
		DataUncompressedBytes *int64  `json:"dataUncompressedBytes,omitempty"`
		ActiveParts           *int64  `json:"activeParts,omitempty"`
		Hosts                 *int64  `json:"hosts,omitempty"`
		MinDate               *string `json:"minDate,omitempty"`
		MaxDate               *string `json:"maxDate,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"database", "table", "partitionId", "partition", "rows", "bytesOnDisk", "dataCompressedBytes", "dataUncompressedBytes", "activeParts", "hosts", "minDate", "maxDate"})
	} else {
		return err
	}
	o.Database = all.Database
	o.Table = all.Table
	o.PartitionId = all.PartitionId
	o.Partition = all.Partition
	o.Rows = all.Rows
	o.BytesOnDisk = all.BytesOnDisk
	o.DataCompressedBytes = all.DataCompressedBytes
	o.DataUncompressedBytes = all.DataUncompressedBytes
	o.ActiveParts = all.ActiveParts
	o.Hosts = all.Hosts
	o.MinDate = all.MinDate
	o.MaxDate = all.MaxDate

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
