// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DataDisk struct {
	DataDiskCapacity     *string `json:"data_disk_capacity,omitempty"`
	DataDiskAllocated    *string `json:"data_disk_allocated,omitempty"`
	DataDiskInUse        *string `json:"data_disk_in_use,omitempty"`
	DataDiskHealthStatus *string `json:"data_disk_health_status,omitempty"`
	DataDiskAbnormalTime *string `json:"data_disk_abnormal_time,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDataDisk instantiates a new DataDisk object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDataDisk() *DataDisk {
	this := DataDisk{}
	return &this
}

// NewDataDiskWithDefaults instantiates a new DataDisk object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDataDiskWithDefaults() *DataDisk {
	this := DataDisk{}
	return &this
}

// GetDataDiskCapacity returns the DataDiskCapacity field value if set, zero value otherwise.
func (o *DataDisk) GetDataDiskCapacity() string {
	if o == nil || o.DataDiskCapacity == nil {
		var ret string
		return ret
	}
	return *o.DataDiskCapacity
}

// GetDataDiskCapacityOk returns a tuple with the DataDiskCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataDisk) GetDataDiskCapacityOk() (*string, bool) {
	if o == nil || o.DataDiskCapacity == nil {
		return nil, false
	}
	return o.DataDiskCapacity, true
}

// HasDataDiskCapacity returns a boolean if a field has been set.
func (o *DataDisk) HasDataDiskCapacity() bool {
	return o != nil && o.DataDiskCapacity != nil
}

// SetDataDiskCapacity gets a reference to the given string and assigns it to the DataDiskCapacity field.
func (o *DataDisk) SetDataDiskCapacity(v string) {
	o.DataDiskCapacity = &v
}

// GetDataDiskAllocated returns the DataDiskAllocated field value if set, zero value otherwise.
func (o *DataDisk) GetDataDiskAllocated() string {
	if o == nil || o.DataDiskAllocated == nil {
		var ret string
		return ret
	}
	return *o.DataDiskAllocated
}

// GetDataDiskAllocatedOk returns a tuple with the DataDiskAllocated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataDisk) GetDataDiskAllocatedOk() (*string, bool) {
	if o == nil || o.DataDiskAllocated == nil {
		return nil, false
	}
	return o.DataDiskAllocated, true
}

// HasDataDiskAllocated returns a boolean if a field has been set.
func (o *DataDisk) HasDataDiskAllocated() bool {
	return o != nil && o.DataDiskAllocated != nil
}

// SetDataDiskAllocated gets a reference to the given string and assigns it to the DataDiskAllocated field.
func (o *DataDisk) SetDataDiskAllocated(v string) {
	o.DataDiskAllocated = &v
}

// GetDataDiskInUse returns the DataDiskInUse field value if set, zero value otherwise.
func (o *DataDisk) GetDataDiskInUse() string {
	if o == nil || o.DataDiskInUse == nil {
		var ret string
		return ret
	}
	return *o.DataDiskInUse
}

// GetDataDiskInUseOk returns a tuple with the DataDiskInUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataDisk) GetDataDiskInUseOk() (*string, bool) {
	if o == nil || o.DataDiskInUse == nil {
		return nil, false
	}
	return o.DataDiskInUse, true
}

// HasDataDiskInUse returns a boolean if a field has been set.
func (o *DataDisk) HasDataDiskInUse() bool {
	return o != nil && o.DataDiskInUse != nil
}

// SetDataDiskInUse gets a reference to the given string and assigns it to the DataDiskInUse field.
func (o *DataDisk) SetDataDiskInUse(v string) {
	o.DataDiskInUse = &v
}

// GetDataDiskHealthStatus returns the DataDiskHealthStatus field value if set, zero value otherwise.
func (o *DataDisk) GetDataDiskHealthStatus() string {
	if o == nil || o.DataDiskHealthStatus == nil {
		var ret string
		return ret
	}
	return *o.DataDiskHealthStatus
}

// GetDataDiskHealthStatusOk returns a tuple with the DataDiskHealthStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataDisk) GetDataDiskHealthStatusOk() (*string, bool) {
	if o == nil || o.DataDiskHealthStatus == nil {
		return nil, false
	}
	return o.DataDiskHealthStatus, true
}

// HasDataDiskHealthStatus returns a boolean if a field has been set.
func (o *DataDisk) HasDataDiskHealthStatus() bool {
	return o != nil && o.DataDiskHealthStatus != nil
}

// SetDataDiskHealthStatus gets a reference to the given string and assigns it to the DataDiskHealthStatus field.
func (o *DataDisk) SetDataDiskHealthStatus(v string) {
	o.DataDiskHealthStatus = &v
}

// GetDataDiskAbnormalTime returns the DataDiskAbnormalTime field value if set, zero value otherwise.
func (o *DataDisk) GetDataDiskAbnormalTime() string {
	if o == nil || o.DataDiskAbnormalTime == nil {
		var ret string
		return ret
	}
	return *o.DataDiskAbnormalTime
}

// GetDataDiskAbnormalTimeOk returns a tuple with the DataDiskAbnormalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataDisk) GetDataDiskAbnormalTimeOk() (*string, bool) {
	if o == nil || o.DataDiskAbnormalTime == nil {
		return nil, false
	}
	return o.DataDiskAbnormalTime, true
}

// HasDataDiskAbnormalTime returns a boolean if a field has been set.
func (o *DataDisk) HasDataDiskAbnormalTime() bool {
	return o != nil && o.DataDiskAbnormalTime != nil
}

// SetDataDiskAbnormalTime gets a reference to the given string and assigns it to the DataDiskAbnormalTime field.
func (o *DataDisk) SetDataDiskAbnormalTime(v string) {
	o.DataDiskAbnormalTime = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DataDisk) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.DataDiskCapacity != nil {
		toSerialize["data_disk_capacity"] = o.DataDiskCapacity
	}
	if o.DataDiskAllocated != nil {
		toSerialize["data_disk_allocated"] = o.DataDiskAllocated
	}
	if o.DataDiskInUse != nil {
		toSerialize["data_disk_in_use"] = o.DataDiskInUse
	}
	if o.DataDiskHealthStatus != nil {
		toSerialize["data_disk_health_status"] = o.DataDiskHealthStatus
	}
	if o.DataDiskAbnormalTime != nil {
		toSerialize["data_disk_abnormal_time"] = o.DataDiskAbnormalTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DataDisk) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataDiskCapacity     *string `json:"data_disk_capacity,omitempty"`
		DataDiskAllocated    *string `json:"data_disk_allocated,omitempty"`
		DataDiskInUse        *string `json:"data_disk_in_use,omitempty"`
		DataDiskHealthStatus *string `json:"data_disk_health_status,omitempty"`
		DataDiskAbnormalTime *string `json:"data_disk_abnormal_time,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"data_disk_capacity", "data_disk_allocated", "data_disk_in_use", "data_disk_health_status", "data_disk_abnormal_time"})
	} else {
		return err
	}
	o.DataDiskCapacity = all.DataDiskCapacity
	o.DataDiskAllocated = all.DataDiskAllocated
	o.DataDiskInUse = all.DataDiskInUse
	o.DataDiskHealthStatus = all.DataDiskHealthStatus
	o.DataDiskAbnormalTime = all.DataDiskAbnormalTime

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
