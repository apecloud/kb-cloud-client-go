// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type BackupDownload struct {
	// the paths of file to download
	Filepaths []string `json:"filepaths,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBackupDownload instantiates a new BackupDownload object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBackupDownload() *BackupDownload {
	this := BackupDownload{}
	return &this
}

// NewBackupDownloadWithDefaults instantiates a new BackupDownload object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBackupDownloadWithDefaults() *BackupDownload {
	this := BackupDownload{}
	return &this
}

// GetFilepaths returns the Filepaths field value if set, zero value otherwise.
func (o *BackupDownload) GetFilepaths() []string {
	if o == nil || o.Filepaths == nil {
		var ret []string
		return ret
	}
	return o.Filepaths
}

// GetFilepathsOk returns a tuple with the Filepaths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupDownload) GetFilepathsOk() (*[]string, bool) {
	if o == nil || o.Filepaths == nil {
		return nil, false
	}
	return &o.Filepaths, true
}

// HasFilepaths returns a boolean if a field has been set.
func (o *BackupDownload) HasFilepaths() bool {
	return o != nil && o.Filepaths != nil
}

// SetFilepaths gets a reference to the given []string and assigns it to the Filepaths field.
func (o *BackupDownload) SetFilepaths(v []string) {
	o.Filepaths = v
}

// MarshalJSON serializes the struct using spec logic.
func (o BackupDownload) MarshalJSON() ([]byte, error) {
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
func (o *BackupDownload) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Filepaths []string `json:"filepaths,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"filepaths"})
	} else {
		return err
	}
	o.Filepaths = all.Filepaths

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
