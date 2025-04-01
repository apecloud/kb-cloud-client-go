// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// Environment_storageConfig Storage config for environment
type Environment_storageConfig struct {
	// these storages will be created
	Storages []EnvironmentStorage `json:"storages"`
	// the storage config for log
	Log Environment_storageConfigLog `json:"log"`
	// the storage config for backup
	Backup Environment_storageConfigBackup `json:"backup"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironment_storageConfig instantiates a new Environment_storageConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironment_storageConfig(storages []EnvironmentStorage, log Environment_storageConfigLog, backup Environment_storageConfigBackup) *Environment_storageConfig {
	this := Environment_storageConfig{}
	this.Storages = storages
	this.Log = log
	this.Backup = backup
	return &this
}

// NewEnvironment_storageConfigWithDefaults instantiates a new Environment_storageConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironment_storageConfigWithDefaults() *Environment_storageConfig {
	this := Environment_storageConfig{}
	return &this
}

// GetStorages returns the Storages field value.
func (o *Environment_storageConfig) GetStorages() []EnvironmentStorage {
	if o == nil {
		var ret []EnvironmentStorage
		return ret
	}
	return o.Storages
}

// GetStoragesOk returns a tuple with the Storages field value
// and a boolean to check if the value has been set.
func (o *Environment_storageConfig) GetStoragesOk() (*[]EnvironmentStorage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Storages, true
}

// SetStorages sets field value.
func (o *Environment_storageConfig) SetStorages(v []EnvironmentStorage) {
	o.Storages = v
}

// GetLog returns the Log field value.
func (o *Environment_storageConfig) GetLog() Environment_storageConfigLog {
	if o == nil {
		var ret Environment_storageConfigLog
		return ret
	}
	return o.Log
}

// GetLogOk returns a tuple with the Log field value
// and a boolean to check if the value has been set.
func (o *Environment_storageConfig) GetLogOk() (*Environment_storageConfigLog, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Log, true
}

// SetLog sets field value.
func (o *Environment_storageConfig) SetLog(v Environment_storageConfigLog) {
	o.Log = v
}

// GetBackup returns the Backup field value.
func (o *Environment_storageConfig) GetBackup() Environment_storageConfigBackup {
	if o == nil {
		var ret Environment_storageConfigBackup
		return ret
	}
	return o.Backup
}

// GetBackupOk returns a tuple with the Backup field value
// and a boolean to check if the value has been set.
func (o *Environment_storageConfig) GetBackupOk() (*Environment_storageConfigBackup, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Backup, true
}

// SetBackup sets field value.
func (o *Environment_storageConfig) SetBackup(v Environment_storageConfigBackup) {
	o.Backup = v
}

// MarshalJSON serializes the struct using spec logic.
func (o Environment_storageConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["storages"] = o.Storages
	toSerialize["log"] = o.Log
	toSerialize["backup"] = o.Backup

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Environment_storageConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Storages *[]EnvironmentStorage            `json:"storages"`
		Log      *Environment_storageConfigLog    `json:"log"`
		Backup   *Environment_storageConfigBackup `json:"backup"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Storages == nil {
		return fmt.Errorf("required field storages missing")
	}
	if all.Log == nil {
		return fmt.Errorf("required field log missing")
	}
	if all.Backup == nil {
		return fmt.Errorf("required field backup missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"storages", "log", "backup"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Storages = *all.Storages
	if all.Log.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Log = *all.Log
	if all.Backup.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Backup = *all.Backup

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
