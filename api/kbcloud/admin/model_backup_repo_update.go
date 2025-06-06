// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// BackupRepoUpdate BackupRepoUpdate is the payload to update a KubeBlocks cluster backup repo
type BackupRepoUpdate struct {
	// default specifies whether the backupRepo is the default backupRepo
	Default bool `json:"default"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBackupRepoUpdate instantiates a new BackupRepoUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBackupRepoUpdate(defaultVar bool) *BackupRepoUpdate {
	this := BackupRepoUpdate{}
	this.Default = defaultVar
	return &this
}

// NewBackupRepoUpdateWithDefaults instantiates a new BackupRepoUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBackupRepoUpdateWithDefaults() *BackupRepoUpdate {
	this := BackupRepoUpdate{}
	return &this
}

// GetDefault returns the Default field value.
func (o *BackupRepoUpdate) GetDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value
// and a boolean to check if the value has been set.
func (o *BackupRepoUpdate) GetDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Default, true
}

// SetDefault sets field value.
func (o *BackupRepoUpdate) SetDefault(v bool) {
	o.Default = v
}

// MarshalJSON serializes the struct using spec logic.
func (o BackupRepoUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["default"] = o.Default

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BackupRepoUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Default *bool `json:"default"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Default == nil {
		return fmt.Errorf("required field default missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"default"})
	} else {
		return err
	}
	o.Default = *all.Default

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
