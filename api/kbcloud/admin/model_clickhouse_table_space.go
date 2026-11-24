// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type ClickhouseTableSpace struct {
	// Database name.
	Database *string `json:"database,omitempty"`
	// Table name.
	Table *string `json:"table,omitempty"`
	// Table engine.
	Engine *string `json:"engine,omitempty"`
	// Physical bytes_on_disk across all hosts and disks.
	BytesOnDisk *int64 `json:"bytesOnDisk,omitempty"`
	// Sum of data_compressed_bytes.
	DataCompressedBytes *int64 `json:"dataCompressedBytes,omitempty"`
	// Sum of data_uncompressed_bytes.
	DataUncompressedBytes *int64 `json:"dataUncompressedBytes,omitempty"`
	// data_compressed_bytes / data_uncompressed_bytes.
	CompressionRatio *float64 `json:"compressionRatio,omitempty"`
	// Sum of active part rows.
	Rows *int64 `json:"rows,omitempty"`
	// Active part count used for the Too many parts risk check.
	ActiveParts *int64 `json:"activeParts,omitempty"`
	// Distinct partitions across disks and hosts.
	PartitionCount *int64 `json:"partitionCount,omitempty"`
	// Number of hosts that own active parts for this table.
	Hosts *int64 `json:"hosts,omitempty"`
	// Latest part modification time in RFC3339 UTC.
	LastModifiedAt *string `json:"lastModifiedAt,omitempty"`
	// normal, warning (parts >= parts_to_delay_insert) or critical (parts >= parts_to_throw_insert).
	Risk *string `json:"risk,omitempty"`
	// Per-disk and per-host breakdown.
	Disks []ClickhouseDiskUsage `json:"disks,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClickhouseTableSpace instantiates a new ClickhouseTableSpace object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClickhouseTableSpace() *ClickhouseTableSpace {
	this := ClickhouseTableSpace{}
	return &this
}

// NewClickhouseTableSpaceWithDefaults instantiates a new ClickhouseTableSpace object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClickhouseTableSpaceWithDefaults() *ClickhouseTableSpace {
	this := ClickhouseTableSpace{}
	return &this
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *ClickhouseTableSpace) GetDatabase() string {
	if o == nil || o.Database == nil {
		var ret string
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseTableSpace) GetDatabaseOk() (*string, bool) {
	if o == nil || o.Database == nil {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *ClickhouseTableSpace) HasDatabase() bool {
	return o != nil && o.Database != nil
}

// SetDatabase gets a reference to the given string and assigns it to the Database field.
func (o *ClickhouseTableSpace) SetDatabase(v string) {
	o.Database = &v
}

// GetTable returns the Table field value if set, zero value otherwise.
func (o *ClickhouseTableSpace) GetTable() string {
	if o == nil || o.Table == nil {
		var ret string
		return ret
	}
	return *o.Table
}

// GetTableOk returns a tuple with the Table field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseTableSpace) GetTableOk() (*string, bool) {
	if o == nil || o.Table == nil {
		return nil, false
	}
	return o.Table, true
}

// HasTable returns a boolean if a field has been set.
func (o *ClickhouseTableSpace) HasTable() bool {
	return o != nil && o.Table != nil
}

// SetTable gets a reference to the given string and assigns it to the Table field.
func (o *ClickhouseTableSpace) SetTable(v string) {
	o.Table = &v
}

// GetEngine returns the Engine field value if set, zero value otherwise.
func (o *ClickhouseTableSpace) GetEngine() string {
	if o == nil || o.Engine == nil {
		var ret string
		return ret
	}
	return *o.Engine
}

// GetEngineOk returns a tuple with the Engine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseTableSpace) GetEngineOk() (*string, bool) {
	if o == nil || o.Engine == nil {
		return nil, false
	}
	return o.Engine, true
}

// HasEngine returns a boolean if a field has been set.
func (o *ClickhouseTableSpace) HasEngine() bool {
	return o != nil && o.Engine != nil
}

// SetEngine gets a reference to the given string and assigns it to the Engine field.
func (o *ClickhouseTableSpace) SetEngine(v string) {
	o.Engine = &v
}

// GetBytesOnDisk returns the BytesOnDisk field value if set, zero value otherwise.
func (o *ClickhouseTableSpace) GetBytesOnDisk() int64 {
	if o == nil || o.BytesOnDisk == nil {
		var ret int64
		return ret
	}
	return *o.BytesOnDisk
}

// GetBytesOnDiskOk returns a tuple with the BytesOnDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseTableSpace) GetBytesOnDiskOk() (*int64, bool) {
	if o == nil || o.BytesOnDisk == nil {
		return nil, false
	}
	return o.BytesOnDisk, true
}

// HasBytesOnDisk returns a boolean if a field has been set.
func (o *ClickhouseTableSpace) HasBytesOnDisk() bool {
	return o != nil && o.BytesOnDisk != nil
}

// SetBytesOnDisk gets a reference to the given int64 and assigns it to the BytesOnDisk field.
func (o *ClickhouseTableSpace) SetBytesOnDisk(v int64) {
	o.BytesOnDisk = &v
}

// GetDataCompressedBytes returns the DataCompressedBytes field value if set, zero value otherwise.
func (o *ClickhouseTableSpace) GetDataCompressedBytes() int64 {
	if o == nil || o.DataCompressedBytes == nil {
		var ret int64
		return ret
	}
	return *o.DataCompressedBytes
}

// GetDataCompressedBytesOk returns a tuple with the DataCompressedBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseTableSpace) GetDataCompressedBytesOk() (*int64, bool) {
	if o == nil || o.DataCompressedBytes == nil {
		return nil, false
	}
	return o.DataCompressedBytes, true
}

// HasDataCompressedBytes returns a boolean if a field has been set.
func (o *ClickhouseTableSpace) HasDataCompressedBytes() bool {
	return o != nil && o.DataCompressedBytes != nil
}

// SetDataCompressedBytes gets a reference to the given int64 and assigns it to the DataCompressedBytes field.
func (o *ClickhouseTableSpace) SetDataCompressedBytes(v int64) {
	o.DataCompressedBytes = &v
}

// GetDataUncompressedBytes returns the DataUncompressedBytes field value if set, zero value otherwise.
func (o *ClickhouseTableSpace) GetDataUncompressedBytes() int64 {
	if o == nil || o.DataUncompressedBytes == nil {
		var ret int64
		return ret
	}
	return *o.DataUncompressedBytes
}

// GetDataUncompressedBytesOk returns a tuple with the DataUncompressedBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseTableSpace) GetDataUncompressedBytesOk() (*int64, bool) {
	if o == nil || o.DataUncompressedBytes == nil {
		return nil, false
	}
	return o.DataUncompressedBytes, true
}

// HasDataUncompressedBytes returns a boolean if a field has been set.
func (o *ClickhouseTableSpace) HasDataUncompressedBytes() bool {
	return o != nil && o.DataUncompressedBytes != nil
}

// SetDataUncompressedBytes gets a reference to the given int64 and assigns it to the DataUncompressedBytes field.
func (o *ClickhouseTableSpace) SetDataUncompressedBytes(v int64) {
	o.DataUncompressedBytes = &v
}

// GetCompressionRatio returns the CompressionRatio field value if set, zero value otherwise.
func (o *ClickhouseTableSpace) GetCompressionRatio() float64 {
	if o == nil || o.CompressionRatio == nil {
		var ret float64
		return ret
	}
	return *o.CompressionRatio
}

// GetCompressionRatioOk returns a tuple with the CompressionRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseTableSpace) GetCompressionRatioOk() (*float64, bool) {
	if o == nil || o.CompressionRatio == nil {
		return nil, false
	}
	return o.CompressionRatio, true
}

// HasCompressionRatio returns a boolean if a field has been set.
func (o *ClickhouseTableSpace) HasCompressionRatio() bool {
	return o != nil && o.CompressionRatio != nil
}

// SetCompressionRatio gets a reference to the given float64 and assigns it to the CompressionRatio field.
func (o *ClickhouseTableSpace) SetCompressionRatio(v float64) {
	o.CompressionRatio = &v
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *ClickhouseTableSpace) GetRows() int64 {
	if o == nil || o.Rows == nil {
		var ret int64
		return ret
	}
	return *o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseTableSpace) GetRowsOk() (*int64, bool) {
	if o == nil || o.Rows == nil {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *ClickhouseTableSpace) HasRows() bool {
	return o != nil && o.Rows != nil
}

// SetRows gets a reference to the given int64 and assigns it to the Rows field.
func (o *ClickhouseTableSpace) SetRows(v int64) {
	o.Rows = &v
}

// GetActiveParts returns the ActiveParts field value if set, zero value otherwise.
func (o *ClickhouseTableSpace) GetActiveParts() int64 {
	if o == nil || o.ActiveParts == nil {
		var ret int64
		return ret
	}
	return *o.ActiveParts
}

// GetActivePartsOk returns a tuple with the ActiveParts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseTableSpace) GetActivePartsOk() (*int64, bool) {
	if o == nil || o.ActiveParts == nil {
		return nil, false
	}
	return o.ActiveParts, true
}

// HasActiveParts returns a boolean if a field has been set.
func (o *ClickhouseTableSpace) HasActiveParts() bool {
	return o != nil && o.ActiveParts != nil
}

// SetActiveParts gets a reference to the given int64 and assigns it to the ActiveParts field.
func (o *ClickhouseTableSpace) SetActiveParts(v int64) {
	o.ActiveParts = &v
}

// GetPartitionCount returns the PartitionCount field value if set, zero value otherwise.
func (o *ClickhouseTableSpace) GetPartitionCount() int64 {
	if o == nil || o.PartitionCount == nil {
		var ret int64
		return ret
	}
	return *o.PartitionCount
}

// GetPartitionCountOk returns a tuple with the PartitionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseTableSpace) GetPartitionCountOk() (*int64, bool) {
	if o == nil || o.PartitionCount == nil {
		return nil, false
	}
	return o.PartitionCount, true
}

// HasPartitionCount returns a boolean if a field has been set.
func (o *ClickhouseTableSpace) HasPartitionCount() bool {
	return o != nil && o.PartitionCount != nil
}

// SetPartitionCount gets a reference to the given int64 and assigns it to the PartitionCount field.
func (o *ClickhouseTableSpace) SetPartitionCount(v int64) {
	o.PartitionCount = &v
}

// GetHosts returns the Hosts field value if set, zero value otherwise.
func (o *ClickhouseTableSpace) GetHosts() int64 {
	if o == nil || o.Hosts == nil {
		var ret int64
		return ret
	}
	return *o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseTableSpace) GetHostsOk() (*int64, bool) {
	if o == nil || o.Hosts == nil {
		return nil, false
	}
	return o.Hosts, true
}

// HasHosts returns a boolean if a field has been set.
func (o *ClickhouseTableSpace) HasHosts() bool {
	return o != nil && o.Hosts != nil
}

// SetHosts gets a reference to the given int64 and assigns it to the Hosts field.
func (o *ClickhouseTableSpace) SetHosts(v int64) {
	o.Hosts = &v
}

// GetLastModifiedAt returns the LastModifiedAt field value if set, zero value otherwise.
func (o *ClickhouseTableSpace) GetLastModifiedAt() string {
	if o == nil || o.LastModifiedAt == nil {
		var ret string
		return ret
	}
	return *o.LastModifiedAt
}

// GetLastModifiedAtOk returns a tuple with the LastModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseTableSpace) GetLastModifiedAtOk() (*string, bool) {
	if o == nil || o.LastModifiedAt == nil {
		return nil, false
	}
	return o.LastModifiedAt, true
}

// HasLastModifiedAt returns a boolean if a field has been set.
func (o *ClickhouseTableSpace) HasLastModifiedAt() bool {
	return o != nil && o.LastModifiedAt != nil
}

// SetLastModifiedAt gets a reference to the given string and assigns it to the LastModifiedAt field.
func (o *ClickhouseTableSpace) SetLastModifiedAt(v string) {
	o.LastModifiedAt = &v
}

// GetRisk returns the Risk field value if set, zero value otherwise.
func (o *ClickhouseTableSpace) GetRisk() string {
	if o == nil || o.Risk == nil {
		var ret string
		return ret
	}
	return *o.Risk
}

// GetRiskOk returns a tuple with the Risk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseTableSpace) GetRiskOk() (*string, bool) {
	if o == nil || o.Risk == nil {
		return nil, false
	}
	return o.Risk, true
}

// HasRisk returns a boolean if a field has been set.
func (o *ClickhouseTableSpace) HasRisk() bool {
	return o != nil && o.Risk != nil
}

// SetRisk gets a reference to the given string and assigns it to the Risk field.
func (o *ClickhouseTableSpace) SetRisk(v string) {
	o.Risk = &v
}

// GetDisks returns the Disks field value if set, zero value otherwise.
func (o *ClickhouseTableSpace) GetDisks() []ClickhouseDiskUsage {
	if o == nil || o.Disks == nil {
		var ret []ClickhouseDiskUsage
		return ret
	}
	return o.Disks
}

// GetDisksOk returns a tuple with the Disks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseTableSpace) GetDisksOk() (*[]ClickhouseDiskUsage, bool) {
	if o == nil || o.Disks == nil {
		return nil, false
	}
	return &o.Disks, true
}

// HasDisks returns a boolean if a field has been set.
func (o *ClickhouseTableSpace) HasDisks() bool {
	return o != nil && o.Disks != nil
}

// SetDisks gets a reference to the given []ClickhouseDiskUsage and assigns it to the Disks field.
func (o *ClickhouseTableSpace) SetDisks(v []ClickhouseDiskUsage) {
	o.Disks = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClickhouseTableSpace) MarshalJSON() ([]byte, error) {
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
	if o.Engine != nil {
		toSerialize["engine"] = o.Engine
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
	if o.CompressionRatio != nil {
		toSerialize["compressionRatio"] = o.CompressionRatio
	}
	if o.Rows != nil {
		toSerialize["rows"] = o.Rows
	}
	if o.ActiveParts != nil {
		toSerialize["activeParts"] = o.ActiveParts
	}
	if o.PartitionCount != nil {
		toSerialize["partitionCount"] = o.PartitionCount
	}
	if o.Hosts != nil {
		toSerialize["hosts"] = o.Hosts
	}
	if o.LastModifiedAt != nil {
		toSerialize["lastModifiedAt"] = o.LastModifiedAt
	}
	if o.Risk != nil {
		toSerialize["risk"] = o.Risk
	}
	if o.Disks != nil {
		toSerialize["disks"] = o.Disks
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClickhouseTableSpace) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Database              *string               `json:"database,omitempty"`
		Table                 *string               `json:"table,omitempty"`
		Engine                *string               `json:"engine,omitempty"`
		BytesOnDisk           *int64                `json:"bytesOnDisk,omitempty"`
		DataCompressedBytes   *int64                `json:"dataCompressedBytes,omitempty"`
		DataUncompressedBytes *int64                `json:"dataUncompressedBytes,omitempty"`
		CompressionRatio      *float64              `json:"compressionRatio,omitempty"`
		Rows                  *int64                `json:"rows,omitempty"`
		ActiveParts           *int64                `json:"activeParts,omitempty"`
		PartitionCount        *int64                `json:"partitionCount,omitempty"`
		Hosts                 *int64                `json:"hosts,omitempty"`
		LastModifiedAt        *string               `json:"lastModifiedAt,omitempty"`
		Risk                  *string               `json:"risk,omitempty"`
		Disks                 []ClickhouseDiskUsage `json:"disks,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"database", "table", "engine", "bytesOnDisk", "dataCompressedBytes", "dataUncompressedBytes", "compressionRatio", "rows", "activeParts", "partitionCount", "hosts", "lastModifiedAt", "risk", "disks"})
	} else {
		return err
	}
	o.Database = all.Database
	o.Table = all.Table
	o.Engine = all.Engine
	o.BytesOnDisk = all.BytesOnDisk
	o.DataCompressedBytes = all.DataCompressedBytes
	o.DataUncompressedBytes = all.DataUncompressedBytes
	o.CompressionRatio = all.CompressionRatio
	o.Rows = all.Rows
	o.ActiveParts = all.ActiveParts
	o.PartitionCount = all.PartitionCount
	o.Hosts = all.Hosts
	o.LastModifiedAt = all.LastModifiedAt
	o.Risk = all.Risk
	o.Disks = all.Disks

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
