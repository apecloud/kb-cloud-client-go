// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// StorageConfigBackup the storage config for backup
type StorageConfigBackup struct {
	// Zero-based index of the default backup storage. Required for multiple storages when storageName is omitted.
	StorageIndex common.NullableInt64 `json:"storageIndex,omitempty"`
	// the name of storage
	StorageName common.NullableString `json:"storageName,omitempty"`
	// the backup repo name which will be init
	BackupRepoName common.NullableString `json:"backupRepoName,omitempty"`
	// judge whether to set the backup is default which will be init
	DefaultBackupRepo common.NullableBool `json:"defaultBackupRepo,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStorageConfigBackup instantiates a new StorageConfigBackup object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStorageConfigBackup() *StorageConfigBackup {
	this := StorageConfigBackup{}
	return &this
}

// NewStorageConfigBackupWithDefaults instantiates a new StorageConfigBackup object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStorageConfigBackupWithDefaults() *StorageConfigBackup {
	this := StorageConfigBackup{}
	return &this
}

// GetStorageIndex returns the StorageIndex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageConfigBackup) GetStorageIndex() int64 {
	if o == nil || o.StorageIndex.Get() == nil {
		var ret int64
		return ret
	}
	return *o.StorageIndex.Get()
}

// GetStorageIndexOk returns a tuple with the StorageIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *StorageConfigBackup) GetStorageIndexOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.StorageIndex.Get(), o.StorageIndex.IsSet()
}

// HasStorageIndex returns a boolean if a field has been set.
func (o *StorageConfigBackup) HasStorageIndex() bool {
	return o != nil && o.StorageIndex.IsSet()
}

// SetStorageIndex gets a reference to the given common.NullableInt64 and assigns it to the StorageIndex field.
func (o *StorageConfigBackup) SetStorageIndex(v int64) {
	o.StorageIndex.Set(&v)
}

// SetStorageIndexNil sets the value for StorageIndex to be an explicit nil.
func (o *StorageConfigBackup) SetStorageIndexNil() {
	o.StorageIndex.Set(nil)
}

// UnsetStorageIndex ensures that no value is present for StorageIndex, not even an explicit nil.
func (o *StorageConfigBackup) UnsetStorageIndex() {
	o.StorageIndex.Unset()
}

// GetStorageName returns the StorageName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageConfigBackup) GetStorageName() string {
	if o == nil || o.StorageName.Get() == nil {
		var ret string
		return ret
	}
	return *o.StorageName.Get()
}

// GetStorageNameOk returns a tuple with the StorageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *StorageConfigBackup) GetStorageNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StorageName.Get(), o.StorageName.IsSet()
}

// HasStorageName returns a boolean if a field has been set.
func (o *StorageConfigBackup) HasStorageName() bool {
	return o != nil && o.StorageName.IsSet()
}

// SetStorageName gets a reference to the given common.NullableString and assigns it to the StorageName field.
func (o *StorageConfigBackup) SetStorageName(v string) {
	o.StorageName.Set(&v)
}

// SetStorageNameNil sets the value for StorageName to be an explicit nil.
func (o *StorageConfigBackup) SetStorageNameNil() {
	o.StorageName.Set(nil)
}

// UnsetStorageName ensures that no value is present for StorageName, not even an explicit nil.
func (o *StorageConfigBackup) UnsetStorageName() {
	o.StorageName.Unset()
}

// GetBackupRepoName returns the BackupRepoName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageConfigBackup) GetBackupRepoName() string {
	if o == nil || o.BackupRepoName.Get() == nil {
		var ret string
		return ret
	}
	return *o.BackupRepoName.Get()
}

// GetBackupRepoNameOk returns a tuple with the BackupRepoName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *StorageConfigBackup) GetBackupRepoNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BackupRepoName.Get(), o.BackupRepoName.IsSet()
}

// HasBackupRepoName returns a boolean if a field has been set.
func (o *StorageConfigBackup) HasBackupRepoName() bool {
	return o != nil && o.BackupRepoName.IsSet()
}

// SetBackupRepoName gets a reference to the given common.NullableString and assigns it to the BackupRepoName field.
func (o *StorageConfigBackup) SetBackupRepoName(v string) {
	o.BackupRepoName.Set(&v)
}

// SetBackupRepoNameNil sets the value for BackupRepoName to be an explicit nil.
func (o *StorageConfigBackup) SetBackupRepoNameNil() {
	o.BackupRepoName.Set(nil)
}

// UnsetBackupRepoName ensures that no value is present for BackupRepoName, not even an explicit nil.
func (o *StorageConfigBackup) UnsetBackupRepoName() {
	o.BackupRepoName.Unset()
}

// GetDefaultBackupRepo returns the DefaultBackupRepo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageConfigBackup) GetDefaultBackupRepo() bool {
	if o == nil || o.DefaultBackupRepo.Get() == nil {
		var ret bool
		return ret
	}
	return *o.DefaultBackupRepo.Get()
}

// GetDefaultBackupRepoOk returns a tuple with the DefaultBackupRepo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *StorageConfigBackup) GetDefaultBackupRepoOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultBackupRepo.Get(), o.DefaultBackupRepo.IsSet()
}

// HasDefaultBackupRepo returns a boolean if a field has been set.
func (o *StorageConfigBackup) HasDefaultBackupRepo() bool {
	return o != nil && o.DefaultBackupRepo.IsSet()
}

// SetDefaultBackupRepo gets a reference to the given common.NullableBool and assigns it to the DefaultBackupRepo field.
func (o *StorageConfigBackup) SetDefaultBackupRepo(v bool) {
	o.DefaultBackupRepo.Set(&v)
}

// SetDefaultBackupRepoNil sets the value for DefaultBackupRepo to be an explicit nil.
func (o *StorageConfigBackup) SetDefaultBackupRepoNil() {
	o.DefaultBackupRepo.Set(nil)
}

// UnsetDefaultBackupRepo ensures that no value is present for DefaultBackupRepo, not even an explicit nil.
func (o *StorageConfigBackup) UnsetDefaultBackupRepo() {
	o.DefaultBackupRepo.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o StorageConfigBackup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.StorageIndex.IsSet() {
		toSerialize["storageIndex"] = o.StorageIndex.Get()
	}
	if o.StorageName.IsSet() {
		toSerialize["storageName"] = o.StorageName.Get()
	}
	if o.BackupRepoName.IsSet() {
		toSerialize["backupRepoName"] = o.BackupRepoName.Get()
	}
	if o.DefaultBackupRepo.IsSet() {
		toSerialize["defaultBackupRepo"] = o.DefaultBackupRepo.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StorageConfigBackup) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		StorageIndex      common.NullableInt64  `json:"storageIndex,omitempty"`
		StorageName       common.NullableString `json:"storageName,omitempty"`
		BackupRepoName    common.NullableString `json:"backupRepoName,omitempty"`
		DefaultBackupRepo common.NullableBool   `json:"defaultBackupRepo,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"storageIndex", "storageName", "backupRepoName", "defaultBackupRepo"})
	} else {
		return err
	}
	o.StorageIndex = all.StorageIndex
	o.StorageName = all.StorageName
	o.BackupRepoName = all.BackupRepoName
	o.DefaultBackupRepo = all.DefaultBackupRepo

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
