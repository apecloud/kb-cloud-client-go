// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package kbcloud

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)



// DmsImportFormData the data of the import task 
type DmsImportFormData struct {
	// the data file, csv or other format
	File _io.Reader `json:"file"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewDmsImportFormData instantiates a new DmsImportFormData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsImportFormData(file _io.Reader) *DmsImportFormData {
	this := DmsImportFormData{}
	this.File = file
	return &this
}

// NewDmsImportFormDataWithDefaults instantiates a new DmsImportFormData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmsImportFormDataWithDefaults() *DmsImportFormData {
	this := DmsImportFormData{}
	return &this
}
// GetFile returns the File field value.
func (o *DmsImportFormData) GetFile() _io.Reader {
	if o == nil {
		var ret _io.Reader
		return ret
	}
	return o.File
}

// GetFileOk returns a tuple with the File field value
// and a boolean to check if the value has been set.
func (o *DmsImportFormData) GetFileOk() (*_io.Reader, bool) {
	if o == nil {
		return nil, false
	}
	return &o.File, true
}

// SetFile sets field value.
func (o *DmsImportFormData) SetFile(v _io.Reader) {
	o.File = v
}



// MarshalJSON serializes the struct using spec logic.
func (o DmsImportFormData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["file"] = o.File

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmsImportFormData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		File *_io.Reader `json:"file"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.File == nil {
		return fmt.Errorf("required field file missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{ "file",  })
	} else {
		return err
	}
	o.File = *all.File

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
