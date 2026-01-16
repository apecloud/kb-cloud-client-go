// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	_io "io"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// VulnUploadFormData the data of the upload task
type VulnUploadFormData struct {
	// the data file, zip or other format
	File _io.Reader `json:"file"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewVulnUploadFormData instantiates a new VulnUploadFormData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewVulnUploadFormData(file _io.Reader) *VulnUploadFormData {
	this := VulnUploadFormData{}
	this.File = file
	return &this
}

// NewVulnUploadFormDataWithDefaults instantiates a new VulnUploadFormData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewVulnUploadFormDataWithDefaults() *VulnUploadFormData {
	this := VulnUploadFormData{}
	return &this
}

// GetFile returns the File field value.
func (o *VulnUploadFormData) GetFile() _io.Reader {
	if o == nil {
		var ret _io.Reader
		return ret
	}
	return o.File
}

// GetFileOk returns a tuple with the File field value
// and a boolean to check if the value has been set.
func (o *VulnUploadFormData) GetFileOk() (*_io.Reader, bool) {
	if o == nil {
		return nil, false
	}
	return &o.File, true
}

// SetFile sets field value.
func (o *VulnUploadFormData) SetFile(v _io.Reader) {
	o.File = v
}

// MarshalJSON serializes the struct using spec logic.
func (o VulnUploadFormData) MarshalJSON() ([]byte, error) {
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
func (o *VulnUploadFormData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		File *_io.Reader `json:"file"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.File == nil {
		return fmt.Errorf("required field file missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"file"})
	} else {
		return err
	}
	o.File = *all.File

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
