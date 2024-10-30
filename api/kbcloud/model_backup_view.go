// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package kbcloud

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)


 
type BackupView struct {
	// the paths of file to view
	Filepaths []string `json:"filepaths,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewBackupView instantiates a new BackupView object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBackupView() *BackupView {
	this := BackupView{}
	return &this
}

// NewBackupViewWithDefaults instantiates a new BackupView object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBackupViewWithDefaults() *BackupView {
	this := BackupView{}
	return &this
}
// GetFilepaths returns the Filepaths field value if set, zero value otherwise.
func (o *BackupView) GetFilepaths() []string {
	if o == nil || o.Filepaths == nil {
		var ret []string
		return ret
	}
	return o.Filepaths
}

// GetFilepathsOk returns a tuple with the Filepaths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupView) GetFilepathsOk() (*[]string, bool) {
	if o == nil || o.Filepaths == nil {
		return nil, false
	}
	return &o.Filepaths, true
}

// HasFilepaths returns a boolean if a field has been set.
func (o *BackupView) HasFilepaths() bool {
	return o != nil && o.Filepaths != nil
}

// SetFilepaths gets a reference to the given []string and assigns it to the Filepaths field.
func (o *BackupView) SetFilepaths(v []string) {
	o.Filepaths = v
}



// MarshalJSON serializes the struct using spec logic.
func (o BackupView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Filepaths != nil {
		toSerialize["filepaths"] = o.Filepaths
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BackupView) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Filepaths []string `json:"filepaths,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{ "filepaths",  })
	} else {
		return err
	}
	o.Filepaths = all.Filepaths

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
