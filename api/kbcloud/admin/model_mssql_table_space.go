// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type MssqlTableSpace struct {
	Schema        string `json:"schema"`
	Name          string `json:"name"`
	Rows          int64  `json:"rows"`
	ReservedBytes int64  `json:"reservedBytes"`
	UsedBytes     int64  `json:"usedBytes"`
	IndexBytes    int64  `json:"indexBytes"`
	UnusedBytes   int64  `json:"unusedBytes"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMssqlTableSpace instantiates a new MssqlTableSpace object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMssqlTableSpace(schema string, name string, rows int64, reservedBytes int64, usedBytes int64, indexBytes int64, unusedBytes int64) *MssqlTableSpace {
	this := MssqlTableSpace{}
	this.Schema = schema
	this.Name = name
	this.Rows = rows
	this.ReservedBytes = reservedBytes
	this.UsedBytes = usedBytes
	this.IndexBytes = indexBytes
	this.UnusedBytes = unusedBytes
	return &this
}

// NewMssqlTableSpaceWithDefaults instantiates a new MssqlTableSpace object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMssqlTableSpaceWithDefaults() *MssqlTableSpace {
	this := MssqlTableSpace{}
	return &this
}

// GetSchema returns the Schema field value.
func (o *MssqlTableSpace) GetSchema() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value
// and a boolean to check if the value has been set.
func (o *MssqlTableSpace) GetSchemaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schema, true
}

// SetSchema sets field value.
func (o *MssqlTableSpace) SetSchema(v string) {
	o.Schema = v
}

// GetName returns the Name field value.
func (o *MssqlTableSpace) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MssqlTableSpace) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *MssqlTableSpace) SetName(v string) {
	o.Name = v
}

// GetRows returns the Rows field value.
func (o *MssqlTableSpace) GetRows() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value
// and a boolean to check if the value has been set.
func (o *MssqlTableSpace) GetRowsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rows, true
}

// SetRows sets field value.
func (o *MssqlTableSpace) SetRows(v int64) {
	o.Rows = v
}

// GetReservedBytes returns the ReservedBytes field value.
func (o *MssqlTableSpace) GetReservedBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ReservedBytes
}

// GetReservedBytesOk returns a tuple with the ReservedBytes field value
// and a boolean to check if the value has been set.
func (o *MssqlTableSpace) GetReservedBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReservedBytes, true
}

// SetReservedBytes sets field value.
func (o *MssqlTableSpace) SetReservedBytes(v int64) {
	o.ReservedBytes = v
}

// GetUsedBytes returns the UsedBytes field value.
func (o *MssqlTableSpace) GetUsedBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.UsedBytes
}

// GetUsedBytesOk returns a tuple with the UsedBytes field value
// and a boolean to check if the value has been set.
func (o *MssqlTableSpace) GetUsedBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsedBytes, true
}

// SetUsedBytes sets field value.
func (o *MssqlTableSpace) SetUsedBytes(v int64) {
	o.UsedBytes = v
}

// GetIndexBytes returns the IndexBytes field value.
func (o *MssqlTableSpace) GetIndexBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.IndexBytes
}

// GetIndexBytesOk returns a tuple with the IndexBytes field value
// and a boolean to check if the value has been set.
func (o *MssqlTableSpace) GetIndexBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IndexBytes, true
}

// SetIndexBytes sets field value.
func (o *MssqlTableSpace) SetIndexBytes(v int64) {
	o.IndexBytes = v
}

// GetUnusedBytes returns the UnusedBytes field value.
func (o *MssqlTableSpace) GetUnusedBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.UnusedBytes
}

// GetUnusedBytesOk returns a tuple with the UnusedBytes field value
// and a boolean to check if the value has been set.
func (o *MssqlTableSpace) GetUnusedBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnusedBytes, true
}

// SetUnusedBytes sets field value.
func (o *MssqlTableSpace) SetUnusedBytes(v int64) {
	o.UnusedBytes = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MssqlTableSpace) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["schema"] = o.Schema
	toSerialize["name"] = o.Name
	toSerialize["rows"] = o.Rows
	toSerialize["reservedBytes"] = o.ReservedBytes
	toSerialize["usedBytes"] = o.UsedBytes
	toSerialize["indexBytes"] = o.IndexBytes
	toSerialize["unusedBytes"] = o.UnusedBytes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MssqlTableSpace) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Schema        *string `json:"schema"`
		Name          *string `json:"name"`
		Rows          *int64  `json:"rows"`
		ReservedBytes *int64  `json:"reservedBytes"`
		UsedBytes     *int64  `json:"usedBytes"`
		IndexBytes    *int64  `json:"indexBytes"`
		UnusedBytes   *int64  `json:"unusedBytes"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Schema == nil {
		return fmt.Errorf("required field schema missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Rows == nil {
		return fmt.Errorf("required field rows missing")
	}
	if all.ReservedBytes == nil {
		return fmt.Errorf("required field reservedBytes missing")
	}
	if all.UsedBytes == nil {
		return fmt.Errorf("required field usedBytes missing")
	}
	if all.IndexBytes == nil {
		return fmt.Errorf("required field indexBytes missing")
	}
	if all.UnusedBytes == nil {
		return fmt.Errorf("required field unusedBytes missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"schema", "name", "rows", "reservedBytes", "usedBytes", "indexBytes", "unusedBytes"})
	} else {
		return err
	}
	o.Schema = *all.Schema
	o.Name = *all.Name
	o.Rows = *all.Rows
	o.ReservedBytes = *all.ReservedBytes
	o.UsedBytes = *all.UsedBytes
	o.IndexBytes = *all.IndexBytes
	o.UnusedBytes = *all.UnusedBytes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
