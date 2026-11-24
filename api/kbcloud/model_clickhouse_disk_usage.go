// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type ClickhouseDiskUsage struct {
	// ClickHouse disk name.
	DiskName *string `json:"diskName,omitempty"`
	// ClickHouse server host.
	HostName *string `json:"hostName,omitempty"`
	// bytes_on_disk for this disk and host.
	BytesOnDisk *int64 `json:"bytesOnDisk,omitempty"`
	// data_compressed_bytes for this disk and host.
	DataCompressedBytes *int64 `json:"dataCompressedBytes,omitempty"`
	// data_uncompressed_bytes for this disk and host.
	DataUncompressedBytes *int64 `json:"dataUncompressedBytes,omitempty"`
	// Rows on this disk and host.
	Rows *int64 `json:"rows,omitempty"`
	// Active parts on this disk and host.
	ActiveParts *int64 `json:"activeParts,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClickhouseDiskUsage instantiates a new ClickhouseDiskUsage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClickhouseDiskUsage() *ClickhouseDiskUsage {
	this := ClickhouseDiskUsage{}
	return &this
}

// NewClickhouseDiskUsageWithDefaults instantiates a new ClickhouseDiskUsage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClickhouseDiskUsageWithDefaults() *ClickhouseDiskUsage {
	this := ClickhouseDiskUsage{}
	return &this
}

// GetDiskName returns the DiskName field value if set, zero value otherwise.
func (o *ClickhouseDiskUsage) GetDiskName() string {
	if o == nil || o.DiskName == nil {
		var ret string
		return ret
	}
	return *o.DiskName
}

// GetDiskNameOk returns a tuple with the DiskName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseDiskUsage) GetDiskNameOk() (*string, bool) {
	if o == nil || o.DiskName == nil {
		return nil, false
	}
	return o.DiskName, true
}

// HasDiskName returns a boolean if a field has been set.
func (o *ClickhouseDiskUsage) HasDiskName() bool {
	return o != nil && o.DiskName != nil
}

// SetDiskName gets a reference to the given string and assigns it to the DiskName field.
func (o *ClickhouseDiskUsage) SetDiskName(v string) {
	o.DiskName = &v
}

// GetHostName returns the HostName field value if set, zero value otherwise.
func (o *ClickhouseDiskUsage) GetHostName() string {
	if o == nil || o.HostName == nil {
		var ret string
		return ret
	}
	return *o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseDiskUsage) GetHostNameOk() (*string, bool) {
	if o == nil || o.HostName == nil {
		return nil, false
	}
	return o.HostName, true
}

// HasHostName returns a boolean if a field has been set.
func (o *ClickhouseDiskUsage) HasHostName() bool {
	return o != nil && o.HostName != nil
}

// SetHostName gets a reference to the given string and assigns it to the HostName field.
func (o *ClickhouseDiskUsage) SetHostName(v string) {
	o.HostName = &v
}

// GetBytesOnDisk returns the BytesOnDisk field value if set, zero value otherwise.
func (o *ClickhouseDiskUsage) GetBytesOnDisk() int64 {
	if o == nil || o.BytesOnDisk == nil {
		var ret int64
		return ret
	}
	return *o.BytesOnDisk
}

// GetBytesOnDiskOk returns a tuple with the BytesOnDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseDiskUsage) GetBytesOnDiskOk() (*int64, bool) {
	if o == nil || o.BytesOnDisk == nil {
		return nil, false
	}
	return o.BytesOnDisk, true
}

// HasBytesOnDisk returns a boolean if a field has been set.
func (o *ClickhouseDiskUsage) HasBytesOnDisk() bool {
	return o != nil && o.BytesOnDisk != nil
}

// SetBytesOnDisk gets a reference to the given int64 and assigns it to the BytesOnDisk field.
func (o *ClickhouseDiskUsage) SetBytesOnDisk(v int64) {
	o.BytesOnDisk = &v
}

// GetDataCompressedBytes returns the DataCompressedBytes field value if set, zero value otherwise.
func (o *ClickhouseDiskUsage) GetDataCompressedBytes() int64 {
	if o == nil || o.DataCompressedBytes == nil {
		var ret int64
		return ret
	}
	return *o.DataCompressedBytes
}

// GetDataCompressedBytesOk returns a tuple with the DataCompressedBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseDiskUsage) GetDataCompressedBytesOk() (*int64, bool) {
	if o == nil || o.DataCompressedBytes == nil {
		return nil, false
	}
	return o.DataCompressedBytes, true
}

// HasDataCompressedBytes returns a boolean if a field has been set.
func (o *ClickhouseDiskUsage) HasDataCompressedBytes() bool {
	return o != nil && o.DataCompressedBytes != nil
}

// SetDataCompressedBytes gets a reference to the given int64 and assigns it to the DataCompressedBytes field.
func (o *ClickhouseDiskUsage) SetDataCompressedBytes(v int64) {
	o.DataCompressedBytes = &v
}

// GetDataUncompressedBytes returns the DataUncompressedBytes field value if set, zero value otherwise.
func (o *ClickhouseDiskUsage) GetDataUncompressedBytes() int64 {
	if o == nil || o.DataUncompressedBytes == nil {
		var ret int64
		return ret
	}
	return *o.DataUncompressedBytes
}

// GetDataUncompressedBytesOk returns a tuple with the DataUncompressedBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseDiskUsage) GetDataUncompressedBytesOk() (*int64, bool) {
	if o == nil || o.DataUncompressedBytes == nil {
		return nil, false
	}
	return o.DataUncompressedBytes, true
}

// HasDataUncompressedBytes returns a boolean if a field has been set.
func (o *ClickhouseDiskUsage) HasDataUncompressedBytes() bool {
	return o != nil && o.DataUncompressedBytes != nil
}

// SetDataUncompressedBytes gets a reference to the given int64 and assigns it to the DataUncompressedBytes field.
func (o *ClickhouseDiskUsage) SetDataUncompressedBytes(v int64) {
	o.DataUncompressedBytes = &v
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *ClickhouseDiskUsage) GetRows() int64 {
	if o == nil || o.Rows == nil {
		var ret int64
		return ret
	}
	return *o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseDiskUsage) GetRowsOk() (*int64, bool) {
	if o == nil || o.Rows == nil {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *ClickhouseDiskUsage) HasRows() bool {
	return o != nil && o.Rows != nil
}

// SetRows gets a reference to the given int64 and assigns it to the Rows field.
func (o *ClickhouseDiskUsage) SetRows(v int64) {
	o.Rows = &v
}

// GetActiveParts returns the ActiveParts field value if set, zero value otherwise.
func (o *ClickhouseDiskUsage) GetActiveParts() int64 {
	if o == nil || o.ActiveParts == nil {
		var ret int64
		return ret
	}
	return *o.ActiveParts
}

// GetActivePartsOk returns a tuple with the ActiveParts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseDiskUsage) GetActivePartsOk() (*int64, bool) {
	if o == nil || o.ActiveParts == nil {
		return nil, false
	}
	return o.ActiveParts, true
}

// HasActiveParts returns a boolean if a field has been set.
func (o *ClickhouseDiskUsage) HasActiveParts() bool {
	return o != nil && o.ActiveParts != nil
}

// SetActiveParts gets a reference to the given int64 and assigns it to the ActiveParts field.
func (o *ClickhouseDiskUsage) SetActiveParts(v int64) {
	o.ActiveParts = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClickhouseDiskUsage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.DiskName != nil {
		toSerialize["diskName"] = o.DiskName
	}
	if o.HostName != nil {
		toSerialize["hostName"] = o.HostName
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
	if o.Rows != nil {
		toSerialize["rows"] = o.Rows
	}
	if o.ActiveParts != nil {
		toSerialize["activeParts"] = o.ActiveParts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClickhouseDiskUsage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DiskName              *string `json:"diskName,omitempty"`
		HostName              *string `json:"hostName,omitempty"`
		BytesOnDisk           *int64  `json:"bytesOnDisk,omitempty"`
		DataCompressedBytes   *int64  `json:"dataCompressedBytes,omitempty"`
		DataUncompressedBytes *int64  `json:"dataUncompressedBytes,omitempty"`
		Rows                  *int64  `json:"rows,omitempty"`
		ActiveParts           *int64  `json:"activeParts,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"diskName", "hostName", "bytesOnDisk", "dataCompressedBytes", "dataUncompressedBytes", "rows", "activeParts"})
	} else {
		return err
	}
	o.DiskName = all.DiskName
	o.HostName = all.HostName
	o.BytesOnDisk = all.BytesOnDisk
	o.DataCompressedBytes = all.DataCompressedBytes
	o.DataUncompressedBytes = all.DataUncompressedBytes
	o.Rows = all.Rows
	o.ActiveParts = all.ActiveParts

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
