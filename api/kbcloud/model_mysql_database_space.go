// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type MysqlDatabaseSpace struct {
	Name       string `json:"name"`
	DataBytes  int64  `json:"dataBytes"`
	IndexBytes int64  `json:"indexBytes"`
	FreeBytes  int64  `json:"freeBytes"`
	TotalBytes int64  `json:"totalBytes"`
	TableCount int64  `json:"tableCount"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMysqlDatabaseSpace instantiates a new MysqlDatabaseSpace object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMysqlDatabaseSpace(name string, dataBytes int64, indexBytes int64, freeBytes int64, totalBytes int64, tableCount int64) *MysqlDatabaseSpace {
	this := MysqlDatabaseSpace{}
	this.Name = name
	this.DataBytes = dataBytes
	this.IndexBytes = indexBytes
	this.FreeBytes = freeBytes
	this.TotalBytes = totalBytes
	this.TableCount = tableCount
	return &this
}

// NewMysqlDatabaseSpaceWithDefaults instantiates a new MysqlDatabaseSpace object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMysqlDatabaseSpaceWithDefaults() *MysqlDatabaseSpace {
	this := MysqlDatabaseSpace{}
	return &this
}

// GetName returns the Name field value.
func (o *MysqlDatabaseSpace) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MysqlDatabaseSpace) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *MysqlDatabaseSpace) SetName(v string) {
	o.Name = v
}

// GetDataBytes returns the DataBytes field value.
func (o *MysqlDatabaseSpace) GetDataBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.DataBytes
}

// GetDataBytesOk returns a tuple with the DataBytes field value
// and a boolean to check if the value has been set.
func (o *MysqlDatabaseSpace) GetDataBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataBytes, true
}

// SetDataBytes sets field value.
func (o *MysqlDatabaseSpace) SetDataBytes(v int64) {
	o.DataBytes = v
}

// GetIndexBytes returns the IndexBytes field value.
func (o *MysqlDatabaseSpace) GetIndexBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.IndexBytes
}

// GetIndexBytesOk returns a tuple with the IndexBytes field value
// and a boolean to check if the value has been set.
func (o *MysqlDatabaseSpace) GetIndexBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IndexBytes, true
}

// SetIndexBytes sets field value.
func (o *MysqlDatabaseSpace) SetIndexBytes(v int64) {
	o.IndexBytes = v
}

// GetFreeBytes returns the FreeBytes field value.
func (o *MysqlDatabaseSpace) GetFreeBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.FreeBytes
}

// GetFreeBytesOk returns a tuple with the FreeBytes field value
// and a boolean to check if the value has been set.
func (o *MysqlDatabaseSpace) GetFreeBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FreeBytes, true
}

// SetFreeBytes sets field value.
func (o *MysqlDatabaseSpace) SetFreeBytes(v int64) {
	o.FreeBytes = v
}

// GetTotalBytes returns the TotalBytes field value.
func (o *MysqlDatabaseSpace) GetTotalBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TotalBytes
}

// GetTotalBytesOk returns a tuple with the TotalBytes field value
// and a boolean to check if the value has been set.
func (o *MysqlDatabaseSpace) GetTotalBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalBytes, true
}

// SetTotalBytes sets field value.
func (o *MysqlDatabaseSpace) SetTotalBytes(v int64) {
	o.TotalBytes = v
}

// GetTableCount returns the TableCount field value.
func (o *MysqlDatabaseSpace) GetTableCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TableCount
}

// GetTableCountOk returns a tuple with the TableCount field value
// and a boolean to check if the value has been set.
func (o *MysqlDatabaseSpace) GetTableCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TableCount, true
}

// SetTableCount sets field value.
func (o *MysqlDatabaseSpace) SetTableCount(v int64) {
	o.TableCount = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MysqlDatabaseSpace) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["dataBytes"] = o.DataBytes
	toSerialize["indexBytes"] = o.IndexBytes
	toSerialize["freeBytes"] = o.FreeBytes
	toSerialize["totalBytes"] = o.TotalBytes
	toSerialize["tableCount"] = o.TableCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MysqlDatabaseSpace) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name       *string `json:"name"`
		DataBytes  *int64  `json:"dataBytes"`
		IndexBytes *int64  `json:"indexBytes"`
		FreeBytes  *int64  `json:"freeBytes"`
		TotalBytes *int64  `json:"totalBytes"`
		TableCount *int64  `json:"tableCount"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
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
	if all.TotalBytes == nil {
		return fmt.Errorf("required field totalBytes missing")
	}
	if all.TableCount == nil {
		return fmt.Errorf("required field tableCount missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "dataBytes", "indexBytes", "freeBytes", "totalBytes", "tableCount"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.DataBytes = *all.DataBytes
	o.IndexBytes = *all.IndexBytes
	o.FreeBytes = *all.FreeBytes
	o.TotalBytes = *all.TotalBytes
	o.TableCount = *all.TableCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
