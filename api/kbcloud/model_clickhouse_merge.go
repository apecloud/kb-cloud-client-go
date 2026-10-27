// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type ClickhouseMerge struct {
	// Database name.
	Database *string `json:"database,omitempty"`
	// Table name.
	Table *string `json:"table,omitempty"`
	// ClickHouse server host running the merge.
	HostName *string `json:"hostName,omitempty"`
	// Seconds elapsed since the merge started.
	ElapsedSeconds *float64 `json:"elapsedSeconds,omitempty"`
	// Merge progress (0-1).
	Progress *float64 `json:"progress,omitempty"`
	// Number of source parts.
	NumParts *int64 `json:"numParts,omitempty"`
	// Resulting part name.
	ResultPartName *string `json:"resultPartName,omitempty"`
	// Partition being merged.
	PartitionId *string `json:"partitionId,omitempty"`
	// Whether this merge also applies a mutation.
	IsMutation *bool `json:"isMutation,omitempty"`
	// Memory usage in bytes.
	MemoryUsageBytes *int64 `json:"memoryUsageBytes,omitempty"`
	// Rows read by the merge.
	RowsRead *int64 `json:"rowsRead,omitempty"`
	// Rows written by the merge.
	RowsWritten *int64 `json:"rowsWritten,omitempty"`
	// Compressed size of the source parts.
	TotalSizeBytesCompressed *int64 `json:"totalSizeBytesCompressed,omitempty"`
	// Disk used for the merge.
	DiskName *string `json:"diskName,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClickhouseMerge instantiates a new ClickhouseMerge object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClickhouseMerge() *ClickhouseMerge {
	this := ClickhouseMerge{}
	return &this
}

// NewClickhouseMergeWithDefaults instantiates a new ClickhouseMerge object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClickhouseMergeWithDefaults() *ClickhouseMerge {
	this := ClickhouseMerge{}
	return &this
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *ClickhouseMerge) GetDatabase() string {
	if o == nil || o.Database == nil {
		var ret string
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseMerge) GetDatabaseOk() (*string, bool) {
	if o == nil || o.Database == nil {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *ClickhouseMerge) HasDatabase() bool {
	return o != nil && o.Database != nil
}

// SetDatabase gets a reference to the given string and assigns it to the Database field.
func (o *ClickhouseMerge) SetDatabase(v string) {
	o.Database = &v
}

// GetTable returns the Table field value if set, zero value otherwise.
func (o *ClickhouseMerge) GetTable() string {
	if o == nil || o.Table == nil {
		var ret string
		return ret
	}
	return *o.Table
}

// GetTableOk returns a tuple with the Table field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseMerge) GetTableOk() (*string, bool) {
	if o == nil || o.Table == nil {
		return nil, false
	}
	return o.Table, true
}

// HasTable returns a boolean if a field has been set.
func (o *ClickhouseMerge) HasTable() bool {
	return o != nil && o.Table != nil
}

// SetTable gets a reference to the given string and assigns it to the Table field.
func (o *ClickhouseMerge) SetTable(v string) {
	o.Table = &v
}

// GetHostName returns the HostName field value if set, zero value otherwise.
func (o *ClickhouseMerge) GetHostName() string {
	if o == nil || o.HostName == nil {
		var ret string
		return ret
	}
	return *o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseMerge) GetHostNameOk() (*string, bool) {
	if o == nil || o.HostName == nil {
		return nil, false
	}
	return o.HostName, true
}

// HasHostName returns a boolean if a field has been set.
func (o *ClickhouseMerge) HasHostName() bool {
	return o != nil && o.HostName != nil
}

// SetHostName gets a reference to the given string and assigns it to the HostName field.
func (o *ClickhouseMerge) SetHostName(v string) {
	o.HostName = &v
}

// GetElapsedSeconds returns the ElapsedSeconds field value if set, zero value otherwise.
func (o *ClickhouseMerge) GetElapsedSeconds() float64 {
	if o == nil || o.ElapsedSeconds == nil {
		var ret float64
		return ret
	}
	return *o.ElapsedSeconds
}

// GetElapsedSecondsOk returns a tuple with the ElapsedSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseMerge) GetElapsedSecondsOk() (*float64, bool) {
	if o == nil || o.ElapsedSeconds == nil {
		return nil, false
	}
	return o.ElapsedSeconds, true
}

// HasElapsedSeconds returns a boolean if a field has been set.
func (o *ClickhouseMerge) HasElapsedSeconds() bool {
	return o != nil && o.ElapsedSeconds != nil
}

// SetElapsedSeconds gets a reference to the given float64 and assigns it to the ElapsedSeconds field.
func (o *ClickhouseMerge) SetElapsedSeconds(v float64) {
	o.ElapsedSeconds = &v
}

// GetProgress returns the Progress field value if set, zero value otherwise.
func (o *ClickhouseMerge) GetProgress() float64 {
	if o == nil || o.Progress == nil {
		var ret float64
		return ret
	}
	return *o.Progress
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseMerge) GetProgressOk() (*float64, bool) {
	if o == nil || o.Progress == nil {
		return nil, false
	}
	return o.Progress, true
}

// HasProgress returns a boolean if a field has been set.
func (o *ClickhouseMerge) HasProgress() bool {
	return o != nil && o.Progress != nil
}

// SetProgress gets a reference to the given float64 and assigns it to the Progress field.
func (o *ClickhouseMerge) SetProgress(v float64) {
	o.Progress = &v
}

// GetNumParts returns the NumParts field value if set, zero value otherwise.
func (o *ClickhouseMerge) GetNumParts() int64 {
	if o == nil || o.NumParts == nil {
		var ret int64
		return ret
	}
	return *o.NumParts
}

// GetNumPartsOk returns a tuple with the NumParts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseMerge) GetNumPartsOk() (*int64, bool) {
	if o == nil || o.NumParts == nil {
		return nil, false
	}
	return o.NumParts, true
}

// HasNumParts returns a boolean if a field has been set.
func (o *ClickhouseMerge) HasNumParts() bool {
	return o != nil && o.NumParts != nil
}

// SetNumParts gets a reference to the given int64 and assigns it to the NumParts field.
func (o *ClickhouseMerge) SetNumParts(v int64) {
	o.NumParts = &v
}

// GetResultPartName returns the ResultPartName field value if set, zero value otherwise.
func (o *ClickhouseMerge) GetResultPartName() string {
	if o == nil || o.ResultPartName == nil {
		var ret string
		return ret
	}
	return *o.ResultPartName
}

// GetResultPartNameOk returns a tuple with the ResultPartName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseMerge) GetResultPartNameOk() (*string, bool) {
	if o == nil || o.ResultPartName == nil {
		return nil, false
	}
	return o.ResultPartName, true
}

// HasResultPartName returns a boolean if a field has been set.
func (o *ClickhouseMerge) HasResultPartName() bool {
	return o != nil && o.ResultPartName != nil
}

// SetResultPartName gets a reference to the given string and assigns it to the ResultPartName field.
func (o *ClickhouseMerge) SetResultPartName(v string) {
	o.ResultPartName = &v
}

// GetPartitionId returns the PartitionId field value if set, zero value otherwise.
func (o *ClickhouseMerge) GetPartitionId() string {
	if o == nil || o.PartitionId == nil {
		var ret string
		return ret
	}
	return *o.PartitionId
}

// GetPartitionIdOk returns a tuple with the PartitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseMerge) GetPartitionIdOk() (*string, bool) {
	if o == nil || o.PartitionId == nil {
		return nil, false
	}
	return o.PartitionId, true
}

// HasPartitionId returns a boolean if a field has been set.
func (o *ClickhouseMerge) HasPartitionId() bool {
	return o != nil && o.PartitionId != nil
}

// SetPartitionId gets a reference to the given string and assigns it to the PartitionId field.
func (o *ClickhouseMerge) SetPartitionId(v string) {
	o.PartitionId = &v
}

// GetIsMutation returns the IsMutation field value if set, zero value otherwise.
func (o *ClickhouseMerge) GetIsMutation() bool {
	if o == nil || o.IsMutation == nil {
		var ret bool
		return ret
	}
	return *o.IsMutation
}

// GetIsMutationOk returns a tuple with the IsMutation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseMerge) GetIsMutationOk() (*bool, bool) {
	if o == nil || o.IsMutation == nil {
		return nil, false
	}
	return o.IsMutation, true
}

// HasIsMutation returns a boolean if a field has been set.
func (o *ClickhouseMerge) HasIsMutation() bool {
	return o != nil && o.IsMutation != nil
}

// SetIsMutation gets a reference to the given bool and assigns it to the IsMutation field.
func (o *ClickhouseMerge) SetIsMutation(v bool) {
	o.IsMutation = &v
}

// GetMemoryUsageBytes returns the MemoryUsageBytes field value if set, zero value otherwise.
func (o *ClickhouseMerge) GetMemoryUsageBytes() int64 {
	if o == nil || o.MemoryUsageBytes == nil {
		var ret int64
		return ret
	}
	return *o.MemoryUsageBytes
}

// GetMemoryUsageBytesOk returns a tuple with the MemoryUsageBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseMerge) GetMemoryUsageBytesOk() (*int64, bool) {
	if o == nil || o.MemoryUsageBytes == nil {
		return nil, false
	}
	return o.MemoryUsageBytes, true
}

// HasMemoryUsageBytes returns a boolean if a field has been set.
func (o *ClickhouseMerge) HasMemoryUsageBytes() bool {
	return o != nil && o.MemoryUsageBytes != nil
}

// SetMemoryUsageBytes gets a reference to the given int64 and assigns it to the MemoryUsageBytes field.
func (o *ClickhouseMerge) SetMemoryUsageBytes(v int64) {
	o.MemoryUsageBytes = &v
}

// GetRowsRead returns the RowsRead field value if set, zero value otherwise.
func (o *ClickhouseMerge) GetRowsRead() int64 {
	if o == nil || o.RowsRead == nil {
		var ret int64
		return ret
	}
	return *o.RowsRead
}

// GetRowsReadOk returns a tuple with the RowsRead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseMerge) GetRowsReadOk() (*int64, bool) {
	if o == nil || o.RowsRead == nil {
		return nil, false
	}
	return o.RowsRead, true
}

// HasRowsRead returns a boolean if a field has been set.
func (o *ClickhouseMerge) HasRowsRead() bool {
	return o != nil && o.RowsRead != nil
}

// SetRowsRead gets a reference to the given int64 and assigns it to the RowsRead field.
func (o *ClickhouseMerge) SetRowsRead(v int64) {
	o.RowsRead = &v
}

// GetRowsWritten returns the RowsWritten field value if set, zero value otherwise.
func (o *ClickhouseMerge) GetRowsWritten() int64 {
	if o == nil || o.RowsWritten == nil {
		var ret int64
		return ret
	}
	return *o.RowsWritten
}

// GetRowsWrittenOk returns a tuple with the RowsWritten field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseMerge) GetRowsWrittenOk() (*int64, bool) {
	if o == nil || o.RowsWritten == nil {
		return nil, false
	}
	return o.RowsWritten, true
}

// HasRowsWritten returns a boolean if a field has been set.
func (o *ClickhouseMerge) HasRowsWritten() bool {
	return o != nil && o.RowsWritten != nil
}

// SetRowsWritten gets a reference to the given int64 and assigns it to the RowsWritten field.
func (o *ClickhouseMerge) SetRowsWritten(v int64) {
	o.RowsWritten = &v
}

// GetTotalSizeBytesCompressed returns the TotalSizeBytesCompressed field value if set, zero value otherwise.
func (o *ClickhouseMerge) GetTotalSizeBytesCompressed() int64 {
	if o == nil || o.TotalSizeBytesCompressed == nil {
		var ret int64
		return ret
	}
	return *o.TotalSizeBytesCompressed
}

// GetTotalSizeBytesCompressedOk returns a tuple with the TotalSizeBytesCompressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseMerge) GetTotalSizeBytesCompressedOk() (*int64, bool) {
	if o == nil || o.TotalSizeBytesCompressed == nil {
		return nil, false
	}
	return o.TotalSizeBytesCompressed, true
}

// HasTotalSizeBytesCompressed returns a boolean if a field has been set.
func (o *ClickhouseMerge) HasTotalSizeBytesCompressed() bool {
	return o != nil && o.TotalSizeBytesCompressed != nil
}

// SetTotalSizeBytesCompressed gets a reference to the given int64 and assigns it to the TotalSizeBytesCompressed field.
func (o *ClickhouseMerge) SetTotalSizeBytesCompressed(v int64) {
	o.TotalSizeBytesCompressed = &v
}

// GetDiskName returns the DiskName field value if set, zero value otherwise.
func (o *ClickhouseMerge) GetDiskName() string {
	if o == nil || o.DiskName == nil {
		var ret string
		return ret
	}
	return *o.DiskName
}

// GetDiskNameOk returns a tuple with the DiskName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseMerge) GetDiskNameOk() (*string, bool) {
	if o == nil || o.DiskName == nil {
		return nil, false
	}
	return o.DiskName, true
}

// HasDiskName returns a boolean if a field has been set.
func (o *ClickhouseMerge) HasDiskName() bool {
	return o != nil && o.DiskName != nil
}

// SetDiskName gets a reference to the given string and assigns it to the DiskName field.
func (o *ClickhouseMerge) SetDiskName(v string) {
	o.DiskName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClickhouseMerge) MarshalJSON() ([]byte, error) {
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
	if o.HostName != nil {
		toSerialize["hostName"] = o.HostName
	}
	if o.ElapsedSeconds != nil {
		toSerialize["elapsedSeconds"] = o.ElapsedSeconds
	}
	if o.Progress != nil {
		toSerialize["progress"] = o.Progress
	}
	if o.NumParts != nil {
		toSerialize["numParts"] = o.NumParts
	}
	if o.ResultPartName != nil {
		toSerialize["resultPartName"] = o.ResultPartName
	}
	if o.PartitionId != nil {
		toSerialize["partitionId"] = o.PartitionId
	}
	if o.IsMutation != nil {
		toSerialize["isMutation"] = o.IsMutation
	}
	if o.MemoryUsageBytes != nil {
		toSerialize["memoryUsageBytes"] = o.MemoryUsageBytes
	}
	if o.RowsRead != nil {
		toSerialize["rowsRead"] = o.RowsRead
	}
	if o.RowsWritten != nil {
		toSerialize["rowsWritten"] = o.RowsWritten
	}
	if o.TotalSizeBytesCompressed != nil {
		toSerialize["totalSizeBytesCompressed"] = o.TotalSizeBytesCompressed
	}
	if o.DiskName != nil {
		toSerialize["diskName"] = o.DiskName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClickhouseMerge) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Database                 *string  `json:"database,omitempty"`
		Table                    *string  `json:"table,omitempty"`
		HostName                 *string  `json:"hostName,omitempty"`
		ElapsedSeconds           *float64 `json:"elapsedSeconds,omitempty"`
		Progress                 *float64 `json:"progress,omitempty"`
		NumParts                 *int64   `json:"numParts,omitempty"`
		ResultPartName           *string  `json:"resultPartName,omitempty"`
		PartitionId              *string  `json:"partitionId,omitempty"`
		IsMutation               *bool    `json:"isMutation,omitempty"`
		MemoryUsageBytes         *int64   `json:"memoryUsageBytes,omitempty"`
		RowsRead                 *int64   `json:"rowsRead,omitempty"`
		RowsWritten              *int64   `json:"rowsWritten,omitempty"`
		TotalSizeBytesCompressed *int64   `json:"totalSizeBytesCompressed,omitempty"`
		DiskName                 *string  `json:"diskName,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"database", "table", "hostName", "elapsedSeconds", "progress", "numParts", "resultPartName", "partitionId", "isMutation", "memoryUsageBytes", "rowsRead", "rowsWritten", "totalSizeBytesCompressed", "diskName"})
	} else {
		return err
	}
	o.Database = all.Database
	o.Table = all.Table
	o.HostName = all.HostName
	o.ElapsedSeconds = all.ElapsedSeconds
	o.Progress = all.Progress
	o.NumParts = all.NumParts
	o.ResultPartName = all.ResultPartName
	o.PartitionId = all.PartitionId
	o.IsMutation = all.IsMutation
	o.MemoryUsageBytes = all.MemoryUsageBytes
	o.RowsRead = all.RowsRead
	o.RowsWritten = all.RowsWritten
	o.TotalSizeBytesCompressed = all.TotalSizeBytesCompressed
	o.DiskName = all.DiskName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
