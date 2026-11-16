// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// StorageConfig Storage config for environment
type StorageConfig struct {
	// these storages will be created
	Storages []EnvironmentStorage `json:"storages"`
	// the storage config for backup
	Backup *StorageConfigBackup `json:"backup,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStorageConfig instantiates a new StorageConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStorageConfig(storages []EnvironmentStorage) *StorageConfig {
	this := StorageConfig{}
	this.Storages = storages
	return &this
}

// NewStorageConfigWithDefaults instantiates a new StorageConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStorageConfigWithDefaults() *StorageConfig {
	this := StorageConfig{}
	return &this
}

// GetStorages returns the Storages field value.
func (o *StorageConfig) GetStorages() []EnvironmentStorage {
	if o == nil {
		var ret []EnvironmentStorage
		return ret
	}
	return o.Storages
}

// GetStoragesOk returns a tuple with the Storages field value
// and a boolean to check if the value has been set.
func (o *StorageConfig) GetStoragesOk() (*[]EnvironmentStorage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Storages, true
}

// SetStorages sets field value.
func (o *StorageConfig) SetStorages(v []EnvironmentStorage) {
	o.Storages = v
}

// GetBackup returns the Backup field value if set, zero value otherwise.
func (o *StorageConfig) GetBackup() StorageConfigBackup {
	if o == nil || o.Backup == nil {
		var ret StorageConfigBackup
		return ret
	}
	return *o.Backup
}

// GetBackupOk returns a tuple with the Backup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageConfig) GetBackupOk() (*StorageConfigBackup, bool) {
	if o == nil || o.Backup == nil {
		return nil, false
	}
	return o.Backup, true
}

// HasBackup returns a boolean if a field has been set.
func (o *StorageConfig) HasBackup() bool {
	return o != nil && o.Backup != nil
}

// SetBackup gets a reference to the given StorageConfigBackup and assigns it to the Backup field.
func (o *StorageConfig) SetBackup(v StorageConfigBackup) {
	o.Backup = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o StorageConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["storages"] = o.Storages
	if o.Backup != nil {
		toSerialize["backup"] = o.Backup
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StorageConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Storages *[]EnvironmentStorage `json:"storages"`
		Backup   *StorageConfigBackup  `json:"backup,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Storages == nil {
		return fmt.Errorf("required field storages missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"storages", "backup"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Storages = *all.Storages
	if all.Backup != nil && all.Backup.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Backup = all.Backup

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
