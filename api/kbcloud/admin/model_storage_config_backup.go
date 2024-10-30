// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package admin

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)



// StorageConfigBackup the storage config for backup 
type StorageConfigBackup struct {
	// the name of storage
	StorageName string `json:"storageName"`
	// the backup repo name which will be init
	BackupRepoName string `json:"backupRepoName"`
	// judge whether to set the backup is default which will be init
	DefaultBackupRepo bool `json:"defaultBackupRepo"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewStorageConfigBackup instantiates a new StorageConfigBackup object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStorageConfigBackup(storageName string, backupRepoName string, defaultBackupRepo bool) *StorageConfigBackup {
	this := StorageConfigBackup{}
	this.StorageName = storageName
	this.BackupRepoName = backupRepoName
	this.DefaultBackupRepo = defaultBackupRepo
	return &this
}

// NewStorageConfigBackupWithDefaults instantiates a new StorageConfigBackup object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStorageConfigBackupWithDefaults() *StorageConfigBackup {
	this := StorageConfigBackup{}
	return &this
}
// GetStorageName returns the StorageName field value.
func (o *StorageConfigBackup) GetStorageName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.StorageName
}

// GetStorageNameOk returns a tuple with the StorageName field value
// and a boolean to check if the value has been set.
func (o *StorageConfigBackup) GetStorageNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageName, true
}

// SetStorageName sets field value.
func (o *StorageConfigBackup) SetStorageName(v string) {
	o.StorageName = v
}


// GetBackupRepoName returns the BackupRepoName field value.
func (o *StorageConfigBackup) GetBackupRepoName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.BackupRepoName
}

// GetBackupRepoNameOk returns a tuple with the BackupRepoName field value
// and a boolean to check if the value has been set.
func (o *StorageConfigBackup) GetBackupRepoNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BackupRepoName, true
}

// SetBackupRepoName sets field value.
func (o *StorageConfigBackup) SetBackupRepoName(v string) {
	o.BackupRepoName = v
}


// GetDefaultBackupRepo returns the DefaultBackupRepo field value.
func (o *StorageConfigBackup) GetDefaultBackupRepo() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.DefaultBackupRepo
}

// GetDefaultBackupRepoOk returns a tuple with the DefaultBackupRepo field value
// and a boolean to check if the value has been set.
func (o *StorageConfigBackup) GetDefaultBackupRepoOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultBackupRepo, true
}

// SetDefaultBackupRepo sets field value.
func (o *StorageConfigBackup) SetDefaultBackupRepo(v bool) {
	o.DefaultBackupRepo = v
}



// MarshalJSON serializes the struct using spec logic.
func (o StorageConfigBackup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["storageName"] = o.StorageName
	toSerialize["backupRepoName"] = o.BackupRepoName
	toSerialize["defaultBackupRepo"] = o.DefaultBackupRepo

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StorageConfigBackup) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		StorageName *string `json:"storageName"`
		BackupRepoName *string `json:"backupRepoName"`
		DefaultBackupRepo *bool `json:"defaultBackupRepo"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.StorageName == nil {
		return fmt.Errorf("required field storageName missing")
	}
	if all.BackupRepoName == nil {
		return fmt.Errorf("required field backupRepoName missing")
	}
	if all.DefaultBackupRepo == nil {
		return fmt.Errorf("required field defaultBackupRepo missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{ "storageName", "backupRepoName", "defaultBackupRepo",  })
	} else {
		return err
	}
	o.StorageName = *all.StorageName
	o.BackupRepoName = *all.BackupRepoName
	o.DefaultBackupRepo = *all.DefaultBackupRepo

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
