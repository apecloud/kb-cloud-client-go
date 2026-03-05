// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	_io "io"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type EngineLicenseEntityCreate struct {
	// ID of the engine license
	LicenseId string `json:"licenseId"`
	// Name of the node
	NodeName string `json:"nodeName"`
	// Key of the entity
	Key *string `json:"key,omitempty"`
	// Description of the entity
	Description *string `json:"description,omitempty"`
	// The key file to upload
	LicenseFile _io.Reader `json:"licenseFile"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEngineLicenseEntityCreate instantiates a new EngineLicenseEntityCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEngineLicenseEntityCreate(licenseId string, nodeName string, licenseFile _io.Reader) *EngineLicenseEntityCreate {
	this := EngineLicenseEntityCreate{}
	this.LicenseId = licenseId
	this.NodeName = nodeName
	this.LicenseFile = licenseFile
	return &this
}

// NewEngineLicenseEntityCreateWithDefaults instantiates a new EngineLicenseEntityCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEngineLicenseEntityCreateWithDefaults() *EngineLicenseEntityCreate {
	this := EngineLicenseEntityCreate{}
	return &this
}

// GetLicenseId returns the LicenseId field value.
func (o *EngineLicenseEntityCreate) GetLicenseId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.LicenseId
}

// GetLicenseIdOk returns a tuple with the LicenseId field value
// and a boolean to check if the value has been set.
func (o *EngineLicenseEntityCreate) GetLicenseIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LicenseId, true
}

// SetLicenseId sets field value.
func (o *EngineLicenseEntityCreate) SetLicenseId(v string) {
	o.LicenseId = v
}

// GetNodeName returns the NodeName field value.
func (o *EngineLicenseEntityCreate) GetNodeName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value
// and a boolean to check if the value has been set.
func (o *EngineLicenseEntityCreate) GetNodeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeName, true
}

// SetNodeName sets field value.
func (o *EngineLicenseEntityCreate) SetNodeName(v string) {
	o.NodeName = v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *EngineLicenseEntityCreate) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineLicenseEntityCreate) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *EngineLicenseEntityCreate) HasKey() bool {
	return o != nil && o.Key != nil
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *EngineLicenseEntityCreate) SetKey(v string) {
	o.Key = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EngineLicenseEntityCreate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineLicenseEntityCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EngineLicenseEntityCreate) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EngineLicenseEntityCreate) SetDescription(v string) {
	o.Description = &v
}

// GetLicenseFile returns the LicenseFile field value.
func (o *EngineLicenseEntityCreate) GetLicenseFile() _io.Reader {
	if o == nil {
		var ret _io.Reader
		return ret
	}
	return o.LicenseFile
}

// GetLicenseFileOk returns a tuple with the LicenseFile field value
// and a boolean to check if the value has been set.
func (o *EngineLicenseEntityCreate) GetLicenseFileOk() (*_io.Reader, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LicenseFile, true
}

// SetLicenseFile sets field value.
func (o *EngineLicenseEntityCreate) SetLicenseFile(v _io.Reader) {
	o.LicenseFile = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EngineLicenseEntityCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["licenseId"] = o.LicenseId
	toSerialize["nodeName"] = o.NodeName
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	toSerialize["licenseFile"] = o.LicenseFile

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EngineLicenseEntityCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		LicenseId   *string     `json:"licenseId"`
		NodeName    *string     `json:"nodeName"`
		Key         *string     `json:"key,omitempty"`
		Description *string     `json:"description,omitempty"`
		LicenseFile *_io.Reader `json:"licenseFile"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.LicenseId == nil {
		return fmt.Errorf("required field licenseId missing")
	}
	if all.NodeName == nil {
		return fmt.Errorf("required field nodeName missing")
	}
	if all.LicenseFile == nil {
		return fmt.Errorf("required field licenseFile missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"licenseId", "nodeName", "key", "description", "licenseFile"})
	} else {
		return err
	}
	o.LicenseId = *all.LicenseId
	o.NodeName = *all.NodeName
	o.Key = all.Key
	o.Description = all.Description
	o.LicenseFile = *all.LicenseFile

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
