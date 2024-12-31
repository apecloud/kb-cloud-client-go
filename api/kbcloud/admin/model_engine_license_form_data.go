// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	_io "io"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EngineLicenseFormData the data of the engine license
type EngineLicenseFormData struct {
	// The license file to upload
	LicenseFile _io.Reader `json:"licenseFile"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEngineLicenseFormData instantiates a new EngineLicenseFormData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEngineLicenseFormData(licenseFile _io.Reader) *EngineLicenseFormData {
	this := EngineLicenseFormData{}
	this.LicenseFile = licenseFile
	return &this
}

// NewEngineLicenseFormDataWithDefaults instantiates a new EngineLicenseFormData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEngineLicenseFormDataWithDefaults() *EngineLicenseFormData {
	this := EngineLicenseFormData{}
	return &this
}

// GetLicenseFile returns the LicenseFile field value.
func (o *EngineLicenseFormData) GetLicenseFile() _io.Reader {
	if o == nil {
		var ret _io.Reader
		return ret
	}
	return o.LicenseFile
}

// GetLicenseFileOk returns a tuple with the LicenseFile field value
// and a boolean to check if the value has been set.
func (o *EngineLicenseFormData) GetLicenseFileOk() (*_io.Reader, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LicenseFile, true
}

// SetLicenseFile sets field value.
func (o *EngineLicenseFormData) SetLicenseFile(v _io.Reader) {
	o.LicenseFile = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EngineLicenseFormData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["licenseFile"] = o.LicenseFile

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EngineLicenseFormData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		LicenseFile *_io.Reader `json:"licenseFile"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.LicenseFile == nil {
		return fmt.Errorf("required field licenseFile missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"licenseFile"})
	} else {
		return err
	}
	o.LicenseFile = *all.LicenseFile

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
